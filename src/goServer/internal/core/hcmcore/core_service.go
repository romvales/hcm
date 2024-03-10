package hcmcore

import (
	"context"
	"errors"
	"goServer/internal/apiClient"
	"goServer/internal/core/pb"
	goServerErrors "goServer/internal/errors"
	"goServer/internal/messages"
	"reflect"
	"regexp"
	"strconv"
)

type CoreServiceServer struct {
	*pb.UnimplementedCoreServiceServer
	apiClient.RequestUsedClient
}

type _Request = pb.CoreServiceRequest
type _Response = pb.CoreServiceResponse
type _Context = context.Context

func NewCoreServiceServer() *CoreServiceServer {
	return &CoreServiceServer{}
}

func checkRequestForMissingParameters(req any, errMsg string) (*_Response, error) {
	constructResponse := func() (*_Response, error) {
		err := goServerErrors.ErrMissingRequestParameter(errMsg)
		msg := err.Error()

		return &_Response{
			Code: pb.CoreServiceResponse_C_MISSINGPARAMETERS,
			SetterResponse: &pb.SetterResponse{
				ErrorMessage: &msg,
			},
		}, err
	}

	if req == nil {
		return constructResponse()
	}

	if reflect.ValueOf(req).IsNil() {
		return constructResponse()
	}

	return nil, nil
}

func checkIfRequestTargetIsMissing(_funcName string, target any) (*_Response, error) {
	if res, err := checkRequestForMissingParameters(
		target,
		messages.MessageNoTargetProvided(_funcName),
	); err != nil {
		return res, err
	}

	return nil, nil
}

func getParameterExpectedId(columns []string, params map[string]any) (columnToSearch string, id string) {
	patt := regexp.MustCompile("(userId|uuid)")

	if columnToSearch = columns[0]; patt.MatchString(columnToSearch) {
		id = *params[columnToSearch].(*string)
	} else {
		id = strconv.FormatInt(*params[columnToSearch].(*int64), 10)
	}

	return
}

func getParameterExpectedFromContext(ctx _Context) (columnToSearch string, id string) {
	if columnName, ok := ctx.Value(QueryContextKey("columnToSearch")).(string); ok {
		columnToSearch = columnName
	}

	if value, ok := ctx.Value(QueryContextKey("id")).(string); ok {
		id = value
	}

	return
}

func checkIfHasValidRequestParams(_funcName string, req *_Request, _type CoreServiceQueryType) (*_Response, error) {
	noReqBodyMsg := messages.MessageNoRequestBodyProvided(_funcName)

	if res, err := checkRequestForMissingParameters(req, noReqBodyMsg); err != nil {
		return res, err
	}

	switch _type {
	case Q_GETTER:
		getterReq := req.GetterRequest

		if res, err := checkRequestForMissingParameters(getterReq, noReqBodyMsg); err != nil {
			return res, err
		}

	case Q_SETTER:
		setterReq := req.SetterRequest

		if res, err := checkRequestForMissingParameters(setterReq, noReqBodyMsg); err != nil {
			return res, err
		}
	}

	return nil, nil
}

func setupErrorResponse(passedError error, code pb.CoreServiceResponse_CoreServiceResponseCode, typ CoreServiceQueryType) (res *_Response, err error) {
	errMsg := passedError.Error()

	switch typ {
	case Q_GETTER:
		return &_Response{
			Code: code,
			GetterResponse: &pb.GetterResponse{
				ErrorMessage: &errMsg,
			},
		}, passedError
	case Q_SETTER:
		return &_Response{
			Code: code,
			SetterResponse: &pb.SetterResponse{
				ErrorMessage: &errMsg,
			},
		}, passedError
	}

	return nil, nil
}

func (srv *CoreServiceServer) countEmptyParameters(params map[string]any) (nonEmptyParams []string, count int) {
	for column, reqValue := range params {
		if !reflect.ValueOf(reqValue).IsNil() {
			nonEmptyParams = append(nonEmptyParams, column)
			count++
		}
	}

	return
}

func (srv *CoreServiceServer) someRequiredFieldAreNotProvided(_funcName, fields string) (*_Response, error) {
	errMsg := messages.MessageRequiredFieldNotProvided(_funcName, fields)

	return &_Response{
		Code: pb.CoreServiceResponse_C_MISSINGPARAMETERS,
		SetterResponse: &pb.SetterResponse{
			ErrorMessage: &errMsg,
		},
	}, errors.New(errMsg)
}

func (srv *CoreServiceServer) queryItemById(ctx _Context, params CoreServiceGetQueryParams) (res *_Response, err error) {
	req := params.Req
	resp := params.Resp
	callback := params.Callback
	_funcName := params.FuncName
	queryTyp := params.Query
	requiredParams := []string{"userId", "targetId", "targetUuid"}

	if queryTyp == "" {
		queryTyp = Q_GETTER
	}

	if res, err := checkIfHasValidRequestParams(_funcName, req, queryTyp); err != nil {
		return res, err
	}

	if !params.DisableReqParamsCheck {
		queryParams := map[string]any{}

		switch queryTyp {
		case Q_GETTER:
			queryParams["userId"] = req.GetterRequest.UserId
			queryParams["id"] = req.GetterRequest.TargetId
			queryParams["uuid"] = req.GetterRequest.TargetUuid
		case Q_SETTER:
			queryParams["userId"] = req.SetterRequest.UserId
			queryParams["id"] = req.SetterRequest.TargetId
			queryParams["uuid"] = req.SetterRequest.TargetUuid
		}

		for key, value := range params.AdditionalReqParams {
			queryParams[key] = value
			requiredParams = append(requiredParams, key)
		}

		if columns, count := srv.countEmptyParameters(queryParams); count > 1 || count == 0 {
			errMsg := messages.MessageProvideAtleastOneOfTheFollowing(_funcName, requiredParams)
			return setupErrorResponse(goServerErrors.ErrMissingRequestParameter(errMsg), pb.CoreServiceResponse_C_CLIENTERROR, queryTyp)
		} else {
			columnToSearch, id := getParameterExpectedId(columns, queryParams)
			ctx = context.WithValue(ctx, QueryContextKey("columnToSearch"), columnToSearch)
			ctx = context.WithValue(ctx, QueryContextKey("id"), id)
		}
	}

	switch req.GetUsedClient() {
	case pb.CoreServiceRequest_C_SUPABASE:
		var client any

		if callback.UseSupabaseCommunityClient {
			client = srv.GetSupabaseCommunityClient()
		} else {
			client = srv.GetSupabaseClient()
		}

		if res, err = callback.SupabaseCallback(ctx, req, resp, client); err != nil {
			return
		}
	default:
		return setupErrorResponse(goServerErrors.ErrInvalidClientFromRequestUnimplemented, pb.CoreServiceResponse_C_CLIENTERROR, queryTyp)
	}

	return nil, nil
}
