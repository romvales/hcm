package hcmcore

import (
	"context"
	"errors"
	"goServer/internal/core/pb"
	"goServer/internal/messages"

	supabaseCommunityGo "github.com/supabase-community/supabase-go"
)

func (srv *CoreServiceServer) deleteItemById(_ context.Context, req *_Request, _funcName, tableName string) (res *_Response, err error) {
	var id, columnToSearch string

	if res, err := checkIfHasValidRequestParams(_funcName, req, "setter"); err != nil {
		return res, err
	}

	deleteReq := req.SetterRequest
	params := map[string]interface{}{
		"id":   deleteReq.TargetId,
		"uuid": deleteReq.TargetUuid,
	}

	usedParams := []string{"targetId", "targetUuid"}

	switch tableName {
	case "workers":
		params["userId"] = deleteReq.UserId
		usedParams = append(usedParams, "userId")
	case "organizations":

	case "roles":

	case "teams":
	}

	if columns, count := srv.countEmptyParameters(params); count > 1 || count == 0 {
		errMsg := messages.MessageProvideAtleastOneOfTheFollowing(_funcName, usedParams)
		return &_Response{
			Code: pb.CoreServiceResponse_C_MISSINGPARAMETERS,
			SetterResponse: &pb.SetterResponse{
				ErrorMessage: &errMsg,
			},
		}, errors.New(errMsg)
	} else {
		columnToSearch, id = getParameterExpectedId(columns, params)
	}

	switch req.GetUsedClient() {
	case pb.CoreServiceRequest_C_SUPABASE:
		var client *supabaseCommunityGo.Client

		if _, client, err = srv.dependencies(req); err != nil {
			errMsg := err.Error()
			return &pb.CoreServiceResponse{
				Code: pb.CoreServiceResponse_C_CLIENTERROR,
				SetterResponse: &pb.SetterResponse{
					ErrorMessage: &errMsg,
				},
			}, err
		}

		if deleteReq.GetSoftDeleteOp() {
			// TODO: Handle soft deletion
		} else {
			query := client.From(tableName).Delete("", "").Eq(columnToSearch, id)

			if _, err := query.ExecuteTo(nil); err != nil {
				errMsg := err.Error()
				return &_Response{
					Code: pb.CoreServiceResponse_C_DBERROR,
					SetterResponse: &pb.SetterResponse{
						ErrorMessage: &errMsg,
					},
				}, err
			}
		}
	default:
	}

	return &_Response{
		Code:           pb.CoreServiceResponse_C_NOERROR,
		SetterResponse: &pb.SetterResponse{},
	}, nil
}

func (srv *CoreServiceServer) DeleteWorkerById(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.DeleteWorkerById()"
	return srv.deleteItemById(ctx, req, _funcName, "workers")
}

func (srv *CoreServiceServer) DeleteOrganizationById(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.DeleteOrganizationById()"
	return srv.deleteItemById(ctx, req, _funcName, "organizations")
}

func (srv *CoreServiceServer) DeleteRoleById(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.DeleteRoleById()"
	return srv.deleteItemById(ctx, req, _funcName, "roles")
}

func (srv *CoreServiceServer) DeleteTeamById(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.DeleteTeamById()"
	return srv.deleteItemById(ctx, req, _funcName, "teams")
}

func (srv *CoreServiceServer) DeleteWorkerIdentityCardById(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.DeleteWorkerIdentityCardById()"
	return srv.deleteItemById(ctx, req, _funcName, "workerIdentityCards")
}

func (srv *CoreServiceServer) DeleteMemberById(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.DeleteMemberById()"
	return srv.deleteItemById(ctx, req, _funcName, "organizationsMembers")
}

func (srv *CoreServiceServer) DeleteCompensationById(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.DeleteCompensationById()"
	return srv.deleteItemById(ctx, req, _funcName, "compensations")
}

func (srv *CoreServiceServer) DeleteAdditionById(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.DeleteAdditionById()"
	return srv.deleteItemById(ctx, req, _funcName, "additions")
}

func (srv *CoreServiceServer) DeleteDeductionById(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.DeleteDeductionById()"
	return srv.deleteItemById(ctx, req, _funcName, "deductions")
}

func (srv *CoreServiceServer) DeletePayrollById(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.DeletePayrollById()"
	return srv.deleteItemById(ctx, req, _funcName, "payrolls")
}

func (srv *CoreServiceServer) DeleteShiftById(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.DeleteShiftById()"
	var tableName = "shifts"

	if res, err := checkIfHasValidRequestParams(_funcName, req, "setter"); err != nil {
		return res, err
	}

	deleteReq := req.SetterRequest

	if deleteReq.GetTargetShiftType() == pb.SetterRequest_T_OVERRIDESHIFT {
		tableName = "overrideShifts"
	}

	return srv.deleteItemById(ctx, req, _funcName, tableName)
}
