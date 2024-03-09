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

	"github.com/nedpals/supabase-go"
	supabaseCommunityGo "github.com/supabase-community/supabase-go"
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
	if columnName, ok := ctx.Value(GetterContextKey("columnToSearch")).(string); ok {
		columnToSearch = columnName
	}

	if value, ok := ctx.Value(GetterContextKey("id")).(string); ok {
		id = value
	}

	return
}

func checkIfHasValidRequestParams(_funcName string, req *_Request, _type string) (*_Response, error) {
	if res, err := checkRequestForMissingParameters(
		req,
		messages.MessageNoRequestBodyProvided(_funcName),
	); err != nil {
		return res, err
	}

	switch _type {
	case "getter":
		getterReq := req.GetterRequest

		if res, err := checkRequestForMissingParameters(
			getterReq,
			messages.MessageNoRequestBodyProvided(_funcName),
		); err != nil {
			return res, err
		}

	case "setter":
		setterReq := req.SetterRequest

		if res, err := checkRequestForMissingParameters(
			setterReq,
			messages.MessageNoRequestBodyProvided(_funcName),
		); err != nil {
			return res, err
		}
	}

	return nil, nil
}

func setupErrorResponse(passedError error, code pb.CoreServiceResponse_CoreServiceResponseCode, typ string) (res *_Response, err error) {
	errMsg := passedError.Error()

	switch typ {
	case "getter":
		return &_Response{
			Code: code,
			GetterResponse: &pb.GetterResponse{
				ErrorMessage: &errMsg,
			},
		}, passedError
	case "setter":
		return &_Response{
			Code: code,
			SetterResponse: &pb.SetterResponse{
				ErrorMessage: &errMsg,
			},
		}, passedError
	}

	return nil, nil
}

func (srv *CoreServiceServer) dependencies(req *_Request) (
	*supabase.Client,
	*supabaseCommunityGo.Client,
	error,
) {

	if _, err := srv.GetClientFromRequest(req); err != nil {
		return nil, nil, err
	}

	return srv.GetSupabaseClient(), srv.GetSupabaseCommunityClient(), nil
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
