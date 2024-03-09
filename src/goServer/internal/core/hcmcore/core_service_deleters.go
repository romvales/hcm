package hcmcore

import (
	"context"
	"goServer/internal/core/pb"

	supabaseCommunityGo "github.com/supabase-community/supabase-go"
)

func (srv *CoreServiceServer) deleteItemById(ctx context.Context, req *_Request, _funcName, tableName string) (res *_Response, err error) {
	res, err = srv.queryItemById(ctx, CoreServiceGetQueryParams{
		Query:    Q_SETTER,
		FuncName: _funcName,
		Req:      req,
		Callback: DatabaseActionCallback{
			UseSupabaseCommunityClient: true,
			SupabaseCallback: func(ctx _Context, req *_Request, resp, client any) (*_Response, error) {
				columnToSearch, id := getParameterExpectedFromContext(ctx)
				supabaseClient := client.(*supabaseCommunityGo.Client)

				deleteReq := req.GetSetterRequest()

				query := supabaseClient.From(tableName).Delete("", "").Eq(columnToSearch, id)

				// TODO: Handle soft delete operations
				if deleteReq.GetSoftDeleteOp() {

					return nil, nil
				}

				if _, err := query.ExecuteTo(resp); err != nil {
					return setupErrorResponse(err, pb.CoreServiceResponse_C_DBERROR, Q_SETTER)
				}

				return nil, nil
			},
		},
	})

	if err != nil {
		return
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
	var tableName = "standardShifts"

	if res, err := checkIfHasValidRequestParams(_funcName, req, Q_SETTER); err != nil {
		return res, err
	}

	deleteReq := req.SetterRequest

	if deleteReq.GetTargetShiftType() == pb.SetterRequest_T_OVERRIDESHIFT {
		tableName = "overrideShifts"
	}

	return srv.deleteItemById(ctx, req, _funcName, tableName)
}

func (srv *CoreServiceServer) DeleteAttendanceById(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.DeleteAttendanceById()"
	return srv.deleteItemById(ctx, req, _funcName, "attendances")
}
