package hcmcore

import (
	"context"
	"goServer/internal/core/converters"
	"goServer/internal/core/pb"
	"goServer/internal/messages"

	goServerErrors "goServer/internal/errors"

	"github.com/nedpals/supabase-go"
)

type DatabaseActionCallback struct {
	UseSupabaseCommunityClient bool
	SupabaseCallback           func(ctx _Context, req *_Request, resp any, client any)
}

type CoreServiceGetQueryParams struct {
	FuncName            string
	Req                 *_Request
	Resp                any
	AdditionalReqParams map[string]any
	Callback            DatabaseActionCallback
}

type GetterContextKey string

func (srv *CoreServiceServer) queryItemById(ctx _Context, params CoreServiceGetQueryParams) (res *_Response, err error) {
	req := params.Req
	resp := params.Resp
	callback := params.Callback
	_funcName := params.FuncName
	requiredParams := []string{"userId", "targetId", "targetUuid"}

	if res, err := checkIfHasValidRequestParams(_funcName, req, "getter"); err != nil {
		return res, err
	}

	getterRequest := req.GetterRequest

	getterParams := map[string]any{
		"userId": getterRequest.UserId,
		"id":     getterRequest.TargetId,
		"uuid":   getterRequest.TargetUuid,
	}

	for key, value := range params.AdditionalReqParams {
		getterParams[key] = value
		requiredParams = append(requiredParams, key)
	}

	if columns, count := srv.countEmptyParameters(getterParams); count > 1 || count == 0 {
		errMsg := messages.MessageProvideAtleastOneOfTheFollowing(_funcName, requiredParams)

		return setupClientErrorResponse(
			goServerErrors.ErrMissingRequestParameter(errMsg),
		)
	} else {
		columnToSearch, id := getParameterExpectedId(columns, getterParams)
		ctx = context.WithValue(ctx, GetterContextKey("columnToSearch"), columnToSearch)
		ctx = context.WithValue(ctx, GetterContextKey("id"), id)
	}

	switch req.GetUsedClient() {
	case pb.CoreServiceRequest_C_SUPABASE:
		if callback.UseSupabaseCommunityClient {
			client := srv.GetSupabaseCommunityClient()
			callback.SupabaseCallback(ctx, req, resp, client)
		} else {
			client := srv.GetSupabaseClient()
			callback.SupabaseCallback(ctx, req, resp, client)
		}
	default:
		errMsg := goServerErrors.ErrInvalidClientFromRequestUnimplemented.Error()

		return &_Response{
			Code: pb.CoreServiceResponse_C_CLIENTERROR,
			GetterResponse: &pb.GetterResponse{
				ErrorMessage: &errMsg,
			},
		}, goServerErrors.ErrInvalidClientFromRequestUnimplemented
	}

	return nil, nil
}

func (srv *CoreServiceServer) GetWorkerById(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.GetWorkerById()"

	resp := &converters.Worker{}

	res, err = srv.queryItemById(ctx, CoreServiceGetQueryParams{
		FuncName: _funcName,
		Resp:     resp,
		Req:      req,
		Callback: DatabaseActionCallback{
			SupabaseCallback: func(ctx _Context, req *_Request, resp, client any) {
				columnToSearch, id := getParameterExpectedFromContext(ctx)
				supabaseClient := client.(*supabase.Client)
				supabaseClient.DB.From("workers").Select().Single().Eq(columnToSearch, id).Execute(&resp)
			},
		},
	})

	if err != nil {
		return
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
		GetterResponse: &pb.GetterResponse{
			WorkerResult: converters.ConvertMapToWorkerProto(resp),
		},
	}, nil
}

func (srv *CoreServiceServer) GetOrganizationById(ctx _Context, req *_Request) (res *_Response, err error) {

	return &_Response{
		Code:           pb.CoreServiceResponse_C_NOERROR,
		GetterResponse: &pb.GetterResponse{},
	}, nil
}

func (srv *CoreServiceServer) GetRoleById(ctx _Context, req *_Request) (res *_Response, err error) {

	return &_Response{
		Code:           pb.CoreServiceResponse_C_NOERROR,
		GetterResponse: &pb.GetterResponse{},
	}, nil
}

func (srv *CoreServiceServer) GetTeamById(ctx _Context, req *_Request) (res *_Response, err error) {

	return &_Response{
		Code:           pb.CoreServiceResponse_C_NOERROR,
		GetterResponse: &pb.GetterResponse{},
	}, nil
}

func (srv *CoreServiceServer) GetRolesFromOrganization(ctx _Context, req *_Request) (res *_Response, err error) {

	return &_Response{
		Code:           pb.CoreServiceResponse_C_NOERROR,
		GetterResponse: &pb.GetterResponse{},
	}, nil
}

func (srv *CoreServiceServer) GetTeamsFromOrganization(ctx _Context, req *_Request) (res *_Response, err error) {

	return &_Response{
		Code:           pb.CoreServiceResponse_C_NOERROR,
		GetterResponse: &pb.GetterResponse{},
	}, nil
}

func (srv *CoreServiceServer) GetOrganizationMembers(ctx _Context, req *_Request) (res *_Response, err error) {

	return &_Response{
		Code:           pb.CoreServiceResponse_C_NOERROR,
		GetterResponse: &pb.GetterResponse{},
	}, nil
}

func (srv *CoreServiceServer) GetOrganizationJoinRequests(ctx _Context, req *_Request) (res *_Response, err error) {

	return &_Response{
		Code:           pb.CoreServiceResponse_C_NOERROR,
		GetterResponse: &pb.GetterResponse{},
	}, nil
}

func (srv *CoreServiceServer) GetWorkerJoinRequests(ctx _Context, req *_Request) (res *_Response, err error) {

	return &_Response{
		Code:           pb.CoreServiceResponse_C_NOERROR,
		GetterResponse: &pb.GetterResponse{},
	}, nil
}

func (srv *CoreServiceServer) GetJoinRequestById(ctx _Context, req *_Request) (res *_Response, err error) {

	return &_Response{
		Code:           pb.CoreServiceResponse_C_NOERROR,
		GetterResponse: &pb.GetterResponse{},
	}, nil
}

func (srv *CoreServiceServer) GetCompensationById(ctx _Context, req *_Request) (res *_Response, err error) {

	return &_Response{
		Code:           pb.CoreServiceResponse_C_NOERROR,
		GetterResponse: &pb.GetterResponse{},
	}, nil
}

func (srv *CoreServiceServer) GetAdditionById(ctx _Context, req *_Request) (res *_Response, err error) {

	return &_Response{
		Code:           pb.CoreServiceResponse_C_NOERROR,
		GetterResponse: &pb.GetterResponse{},
	}, nil
}

func (srv *CoreServiceServer) GetDeductionById(ctx _Context, req *_Request) (res *_Response, err error) {

	return &_Response{
		Code:           pb.CoreServiceResponse_C_NOERROR,
		GetterResponse: &pb.GetterResponse{},
	}, nil
}

func (srv *CoreServiceServer) GetShiftById(ctx _Context, req *_Request) (res *_Response, err error) {

	return &_Response{
		Code:           pb.CoreServiceResponse_C_NOERROR,
		GetterResponse: &pb.GetterResponse{},
	}, nil
}

func (srv *CoreServiceServer) GetAttendanceByid(ctx _Context, req *_Request) (res *_Response, err error) {

	return &_Response{
		Code:           pb.CoreServiceResponse_C_NOERROR,
		GetterResponse: &pb.GetterResponse{},
	}, nil
}
