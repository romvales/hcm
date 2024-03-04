package hcmcore

import (
	"goServer/internal/core/converters"
	"goServer/internal/core/pb"
	"goServer/internal/messages"

	goServerErrors "goServer/internal/errors"

	"github.com/nedpals/supabase-go"
)

func (srv *CoreServiceServer) GetWorkerById(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.GetWorkerById()"

	var id, columnToSearch string
	var resp []*converters.Worker
	var errMsg string

	if res, err := checkIfHasValidRequestParams(_funcName, req, "getter"); err != nil {
		return res, err
	}

	getterRequest := req.GetterRequest

	if res, err := checkRequestForMissingParameters(
		getterRequest,
		messages.MessageNoRequestBodyProvided(_funcName),
	); err != nil {
		return res, err
	}

	params := map[string]any{
		"userId": getterRequest.UserId,
		"id":     getterRequest.TargetId,
		"uuid":   getterRequest.TargetUuid,
	}

	if columns, count := srv.countEmptyParameters(params); count > 1 || count == 0 {
		errMsg = messages.MessageProvideAtleastOneOfTheFollowing(_funcName, []string{"userId", "targetId", "targetUuid"})

		return &_Response{
			Code: pb.CoreServiceResponse_C_MISSINGPARAMETERS,
			GetterResponse: &pb.GetterResponse{
				ErrorMessage: &errMsg,
			},
		}, goServerErrors.ErrMissingRequestParameter(errMsg)
	} else {
		columnToSearch, id = getParameterExpectedId(columns, params)
	}

	switch req.GetUsedClient() {
	case pb.CoreServiceRequest_C_SUPABASE:
		var client *supabase.Client

		if client, _, err = srv.dependencies(req); err != nil {
			errMsg := err.Error()
			return &_Response{
				Code: pb.CoreServiceResponse_C_CLIENTERROR,
				GetterResponse: &pb.GetterResponse{
					ErrorMessage: &errMsg,
				},
			}, err
		} else {
			err = client.DB.From("workers").Select().Limit(1).Eq(columnToSearch, id).Execute(&resp)
		}

		if err != nil {
			errMsg := err.Error()

			return &_Response{
				Code: pb.CoreServiceResponse_C_DBERROR,
				GetterResponse: &pb.GetterResponse{
					ErrorMessage: &errMsg,
				},
			}, err
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

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
	}, nil
}

func (srv *CoreServiceServer) GetOrganizationById(ctx _Context, req *_Request) (res *_Response, err error) {

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
	}, nil
}

func (srv *CoreServiceServer) GetRoleById(ctx _Context, req *_Request) (res *_Response, err error) {

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
	}, nil
}

func (srv *CoreServiceServer) GetTeamById(ctx _Context, req *_Request) (res *_Response, err error) {

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
	}, nil
}

func (srv *CoreServiceServer) GetRolesFromOrganization(ctx _Context, req *_Request) (res *_Response, err error) {

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
	}, nil
}

func (srv *CoreServiceServer) GetTeamsFromOrganization(ctx _Context, req *_Request) (res *_Response, err error) {

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
	}, nil
}

func (srv *CoreServiceServer) GetOrganizationMembers(ctx _Context, req *_Request) (res *_Response, err error) {

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
	}, nil
}

func (srv *CoreServiceServer) GetOrganizationJoinRequests(ctx _Context, req *_Request) (res *_Response, err error) {

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
	}, nil
}

func (srv *CoreServiceServer) GetWorkerJoinRequests(ctx _Context, req *_Request) (res *_Response, err error) {

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
	}, nil
}

func (srv *CoreServiceServer) GetJoinRequestById(ctx _Context, req *_Request) (res *_Response, err error) {

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
	}, nil
}

func (srv *CoreServiceServer) GetCompensationById(ctx _Context, req *_Request) (res *_Response, err error) {

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
	}, nil
}

func (srv *CoreServiceServer) GetAdditionById(ctx _Context, req *_Request) (res *_Response, err error) {

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
	}, nil
}

func (srv *CoreServiceServer) GetDeductionById(ctx _Context, req *_Request) (res *_Response, err error) {

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
	}, nil
}
