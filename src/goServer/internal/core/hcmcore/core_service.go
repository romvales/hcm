package hcmcore

import (
	"context"
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
	if req == nil {
		err := goServerErrors.ErrMissingRequestParameter(errMsg)
		msg := err.Error()

		return &_Response{
			Code: pb.CoreServiceResponse_C_MISSINGPARAMETERS,
			SetterResponse: &pb.SetterResponse{
				ErrorMessage: &msg,
			},
		}, err
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

func checkIfHasValidRequestParams(_funcName string, req *_Request, _type string) (*_Response, error) {
	if res, err := checkRequestForMissingParameters(
		req,
		messages.MessageNoRequestBodyProvided(_funcName),
	); err != nil {
		return res, err
	}

	switch _type {
	case "getter":
		if res, err := checkRequestForMissingParameters(
			req.GetterRequest,
			messages.MessageNoRequestBodyProvided(_funcName),
		); err != nil {
			return res, err
		}

	case "setter":
		if res, err := checkRequestForMissingParameters(
			req.SetterRequest,
			messages.MessageNoRequestBodyProvided(_funcName),
		); err != nil {
			return res, err
		}
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
