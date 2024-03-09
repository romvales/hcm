package hcmcore

import (
	"context"
	"goServer/internal/core/converters"
	"goServer/internal/core/pb"
	"goServer/internal/messages"
	"strconv"

	goServerErrors "goServer/internal/errors"

	"github.com/nedpals/supabase-go"
)

type CoreServiceQueryType string

const (
	Q_GETTER CoreServiceQueryType = "getter"
	Q_SETTER CoreServiceQueryType = "setter"
)

type DatabaseActionCallback struct {
	UseSupabaseCommunityClient bool
	SupabaseCallback           func(ctx _Context, req *_Request, resp any, client any) (*_Response, error)
}

type CoreServiceGetQueryParams struct {
	Query               CoreServiceQueryType
	FuncName            string
	Req                 *_Request
	Resp                any
	AdditionalReqParams map[string]any
	Callback            DatabaseActionCallback
}

type GetterContextKey string

func supabaseGetItemByIdCallback(tableName string) func(ctx _Context, req *_Request, resp, client any) (*_Response, error) {
	return func(ctx _Context, req *_Request, resp, client any) (res *_Response, err error) {
		columnToSearch, id := getParameterExpectedFromContext(ctx)
		supabaseClient := client.(*supabase.Client)

		err = supabaseClient.DB.From(tableName).Select().Single().Eq(columnToSearch, id).Execute(&resp)
		if err != nil {
			return setupErrorResponse(err, pb.CoreServiceResponse_C_DBERROR, string(Q_GETTER))
		}

		return nil, nil
	}
}

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
		return setupErrorResponse(goServerErrors.ErrMissingRequestParameter(errMsg), pb.CoreServiceResponse_C_CLIENTERROR, string(params.Query))
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
		return setupErrorResponse(goServerErrors.ErrInvalidClientFromRequestUnimplemented, pb.CoreServiceResponse_C_CLIENTERROR, "setter")
	}

	return nil, nil
}

func (srv *CoreServiceServer) GetWorkerById(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.GetWorkerById()"

	resp := &converters.Worker{}

	res, err = srv.queryItemById(ctx, CoreServiceGetQueryParams{
		Query:    Q_GETTER,
		FuncName: _funcName,
		Resp:     resp,
		Req:      req,
		Callback: DatabaseActionCallback{
			SupabaseCallback: supabaseGetItemByIdCallback("workers"),
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
	var _funcName = "CoreServiceServer.GetOrganizationById()"

	resp := &converters.Organization{}

	res, err = srv.queryItemById(ctx, CoreServiceGetQueryParams{
		Query:    Q_GETTER,
		FuncName: _funcName,
		Resp:     resp,
		Req:      req,
		Callback: DatabaseActionCallback{
			SupabaseCallback: supabaseGetItemByIdCallback("organizations"),
		},
	})
	if err != nil {
		return
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
		GetterResponse: &pb.GetterResponse{
			OrganizationResult: converters.ConvertMapToOrganizationProto(resp),
		},
	}, nil
}

func (srv *CoreServiceServer) GetOrganizationsByCreatorId(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.GetOrganizationByCreatorId()"

	result := []*pb.Organization{}

	res, err = srv.queryItemById(ctx, CoreServiceGetQueryParams{
		Query:    Q_GETTER,
		FuncName: _funcName,
		Req:      req,
		Callback: DatabaseActionCallback{
			SupabaseCallback: func(ctx _Context, _ *_Request, _, client any) (*_Response, error) {
				supabaseClient := client.(*supabase.Client)
				workerId := strconv.FormatInt(req.GetterRequest.GetTargetId(), 10)

				resp := []*converters.Organization{}

				err = supabaseClient.DB.From("organizations").Select().Eq("createdById", workerId).Execute(&resp)
				if err != nil {
					return setupErrorResponse(err, pb.CoreServiceResponse_C_DBERROR, "getter")
				}

				for _, organization := range resp {
					result = append(result, converters.ConvertMapToOrganizationProto(organization))
				}

				return nil, nil
			},
		},
	})

	if err != nil {
		return
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
		GetterResponse: &pb.GetterResponse{
			OrganizationsResult: result,
		},
	}, nil
}

func (srv *CoreServiceServer) GetRoleById(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.GetRoleById()"

	resp := &converters.Role{}

	res, err = srv.queryItemById(ctx, CoreServiceGetQueryParams{
		Query:    Q_GETTER,
		FuncName: _funcName,
		Resp:     resp,
		Req:      req,
		Callback: DatabaseActionCallback{
			SupabaseCallback: supabaseGetItemByIdCallback("roles"),
		},
	})
	if err != nil {
		return
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
		GetterResponse: &pb.GetterResponse{
			RoleResult: converters.ConvertMapToRoleProto(resp),
		},
	}, nil
}

func (srv *CoreServiceServer) GetTeamById(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.GetTeamById()"

	resp := &converters.Team{}

	res, err = srv.queryItemById(ctx, CoreServiceGetQueryParams{
		Query:    Q_GETTER,
		FuncName: _funcName,
		Resp:     resp,
		Req:      req,
		Callback: DatabaseActionCallback{
			SupabaseCallback: supabaseGetItemByIdCallback("teams"),
		},
	})
	if err != nil {
		return
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
		GetterResponse: &pb.GetterResponse{
			TeamResult: converters.ConvertMapToTeamProto(resp),
		},
	}, nil
}

func (srv *CoreServiceServer) GetRolesFromOrganization(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.GetRolesFromOrganization()"

	result := []*pb.Role{}

	res, err = srv.queryItemById(ctx, CoreServiceGetQueryParams{
		Query:    Q_GETTER,
		FuncName: _funcName,
		Req:      req,
		Callback: DatabaseActionCallback{
			SupabaseCallback: func(ctx _Context, req *_Request, _, client any) (*_Response, error) {
				_, organizationId := getParameterExpectedFromContext(ctx)
				supabaseClient := client.(*supabase.Client)
				resp := []*converters.Role{}

				err = supabaseClient.DB.From("roles").Select().Eq("organizationId", organizationId).Execute(&resp)
				if err != nil {
					return setupErrorResponse(err, pb.CoreServiceResponse_C_DBERROR, "getter")
				}

				for _, role := range resp {
					result = append(result, converters.ConvertMapToRoleProto(role))
				}

				return nil, nil
			},
		},
	})

	if err != nil {
		return
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
		GetterResponse: &pb.GetterResponse{
			OrganizationRolesResult: result,
		},
	}, nil
}

func (srv *CoreServiceServer) GetTeamsFromOrganization(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.GetTeamsFromOrganization()"

	result := []*pb.Team{}

	res, err = srv.queryItemById(ctx, CoreServiceGetQueryParams{
		Query:    Q_GETTER,
		FuncName: _funcName,
		Req:      req,
		Callback: DatabaseActionCallback{
			SupabaseCallback: func(ctx _Context, req *_Request, _, client any) (*_Response, error) {
				_, organizationId := getParameterExpectedFromContext(ctx)
				supabaseClient := client.(*supabase.Client)
				resp := []*converters.Team{}

				err = supabaseClient.DB.From("teams").Select().Eq("organizationId", organizationId).Execute(&resp)
				if err != nil {
					return setupErrorResponse(err, pb.CoreServiceResponse_C_DBERROR, "getter")
				}

				for _, team := range resp {
					result = append(result, converters.ConvertMapToTeamProto(team))
				}

				return nil, nil
			},
		},
	})

	if err != nil {
		return
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
		GetterResponse: &pb.GetterResponse{
			OrganizationTeamsResult: result,
		},
	}, nil
}

func (srv *CoreServiceServer) GetOrganizationMembers(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.GetOrganizationMembers()"

	result := []*pb.Member{}

	res, err = srv.queryItemById(ctx, CoreServiceGetQueryParams{
		Query:    Q_GETTER,
		FuncName: _funcName,
		Req:      req,
		Callback: DatabaseActionCallback{
			SupabaseCallback: func(ctx _Context, req *_Request, _, client any) (*_Response, error) {
				_, organizationId := getParameterExpectedFromContext(ctx)
				supabaseClient := client.(*supabase.Client)

				resp := []*converters.Member{}

				err = supabaseClient.DB.From("organizationsMembers").Select().Eq("organizationId", organizationId).Execute(&resp)
				if err != nil {
					return setupErrorResponse(err, pb.CoreServiceResponse_C_DBERROR, "getter")
				}

				for _, member := range resp {
					result = append(result, converters.ConvertMapToMemberProto(member))
				}

				return nil, nil
			},
		},
	})

	if err != nil {
		return
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
		GetterResponse: &pb.GetterResponse{
			MembersResult: result,
		},
	}, nil
}

func (srv *CoreServiceServer) GetOrganizationJoinRequests(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.GetOrganizationJoinRequests()"

	result := []*pb.OrganizationPendingRequestRelation{}

	res, err = srv.queryItemById(ctx, CoreServiceGetQueryParams{
		Query:    Q_GETTER,
		FuncName: _funcName,
		Req:      req,
		Callback: DatabaseActionCallback{
			SupabaseCallback: func(ctx _Context, req *_Request, _, client any) (*_Response, error) {
				_, organizationId := getParameterExpectedFromContext(ctx)
				supabaseClient := client.(*supabase.Client)

				resp := []*converters.OrganizationPendingRequestRelation{}

				if err = supabaseClient.DB.From("organizationsPendingRequests").Select().Eq("organizationId", organizationId).Execute(&resp); err != nil {
					return setupErrorResponse(err, pb.CoreServiceResponse_C_DBERROR, "getter")
				}

				for _, joinRequest := range resp {
					result = append(result, converters.ConvertMapToOrganizationPendingRequestRelationProto(joinRequest))
				}

				return nil, nil
			},
		},
	})

	if err != nil {
		return
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
		GetterResponse: &pb.GetterResponse{
			OrganizationPendingRequestsResult: result,
		},
	}, nil
}

func (srv *CoreServiceServer) GetWorkerJoinRequests(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.GetWorkerJoinRequests()"

	result := []*pb.WorkerPendingRequestRelation{}

	res, err = srv.queryItemById(ctx, CoreServiceGetQueryParams{
		Query:    Q_GETTER,
		FuncName: _funcName,
		Req:      req,
		Callback: DatabaseActionCallback{
			SupabaseCallback: func(ctx _Context, req *_Request, _, client any) (*_Response, error) {
				_, workerId := getParameterExpectedFromContext(ctx)
				supabaseClient := client.(*supabase.Client)

				resp := []*converters.WorkerPendingRequestRelation{}

				if err = supabaseClient.DB.From("workersPendingRequests").Select().Eq("workerId", workerId).Execute(&resp); err != nil {
					return setupErrorResponse(err, pb.CoreServiceResponse_C_DBERROR, "getter")
				}

				for _, joinRequest := range resp {
					result = append(result, converters.ConvertMapToWorkerPendingRequestRelationProto(joinRequest))
				}

				return nil, nil
			},
		},
	})

	if err != nil {
		return
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
		GetterResponse: &pb.GetterResponse{
			WorkerPendingRequestsResult: result,
		},
	}, nil
}

func (srv *CoreServiceServer) GetJoinRequestById(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.GetJoinRequestById()"

	resp := &converters.JoinRequest{}

	res, err = srv.queryItemById(ctx, CoreServiceGetQueryParams{
		Query:    Q_GETTER,
		FuncName: _funcName,
		Resp:     resp,
		Req:      req,
		Callback: DatabaseActionCallback{
			SupabaseCallback: supabaseGetItemByIdCallback("pendingJoinRequests"),
		},
	})
	if err != nil {
		return
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
		GetterResponse: &pb.GetterResponse{
			JoinRequestResult: converters.ConvertMapToJoinRequestProto(resp),
		},
	}, nil
}

func (srv *CoreServiceServer) GetPayrollById(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.GetPayrollById()"

	resp := &converters.Payroll{}

	res, err = srv.queryItemById(ctx, CoreServiceGetQueryParams{
		Query:    Q_GETTER,
		FuncName: _funcName,
		Resp:     resp,
		Req:      req,
		Callback: DatabaseActionCallback{
			SupabaseCallback: supabaseGetItemByIdCallback("payrolls"),
		},
	})
	if err != nil {
		return
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
		GetterResponse: &pb.GetterResponse{
			PayrollResult: converters.ConvertMapToPayrollProto(resp),
		},
	}, nil
}

func (srv *CoreServiceServer) GetCompensationById(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.GetCompensationById()"

	resp := &converters.Compensation{}

	res, err = srv.queryItemById(ctx, CoreServiceGetQueryParams{
		Query:    Q_GETTER,
		FuncName: _funcName,
		Resp:     resp,
		Req:      req,
		Callback: DatabaseActionCallback{
			SupabaseCallback: supabaseGetItemByIdCallback("compensations"),
		},
	})
	if err != nil {
		return
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
		GetterResponse: &pb.GetterResponse{
			CompensationResult: converters.ConvertMapToCompensationProto(resp),
		},
	}, nil
}

func (srv *CoreServiceServer) GetAdditionById(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.GetAdditionById()"

	resp := &converters.Addition{}

	res, err = srv.queryItemById(ctx, CoreServiceGetQueryParams{
		Query:    Q_GETTER,
		FuncName: _funcName,
		Resp:     resp,
		Req:      req,
		Callback: DatabaseActionCallback{
			SupabaseCallback: supabaseGetItemByIdCallback("additions"),
		},
	})
	if err != nil {
		return
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
		GetterResponse: &pb.GetterResponse{
			AdditionResult: converters.ConvertMapToAdditionProto(resp),
		},
	}, nil
}

func (srv *CoreServiceServer) GetDeductionById(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.GetDeductionById()"

	resp := &converters.Deduction{}

	res, err = srv.queryItemById(ctx, CoreServiceGetQueryParams{
		Query:    Q_GETTER,
		FuncName: _funcName,
		Resp:     resp,
		Req:      req,
		Callback: DatabaseActionCallback{
			SupabaseCallback: supabaseGetItemByIdCallback("deductions"),
		},
	})
	if err != nil {
		return
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
		GetterResponse: &pb.GetterResponse{
			DeductionResult: converters.ConvertMapToDeductionProto(resp),
		},
	}, nil
}

func (srv *CoreServiceServer) GetShiftById(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.GetShiftById()"
	var tableName string

	resp := &converters.Role{}

	switch req.GetterRequest.GetTargetShiftType() {
	case pb.SetterRequest_T_OVERRIDESHIFT:
		tableName = "overrideShifts"
	case pb.SetterRequest_T_SHIFT:
		tableName = "standardShifts"
	}

	res, err = srv.queryItemById(ctx, CoreServiceGetQueryParams{
		Query:    Q_GETTER,
		FuncName: _funcName,
		Resp:     resp,
		Req:      req,
		Callback: DatabaseActionCallback{
			SupabaseCallback: supabaseGetItemByIdCallback(tableName),
		},
	})
	if err != nil {
		return
	}

	return &_Response{
		Code:           pb.CoreServiceResponse_C_NOERROR,
		GetterResponse: &pb.GetterResponse{},
	}, nil
}

func (srv *CoreServiceServer) GetAttendanceById(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.GetAttendanceById()"

	resp := &converters.Role{}

	res, err = srv.queryItemById(ctx, CoreServiceGetQueryParams{
		Query:    Q_GETTER,
		FuncName: _funcName,
		Resp:     resp,
		Req:      req,
		Callback: DatabaseActionCallback{
			SupabaseCallback: supabaseGetItemByIdCallback("attendances"),
		},
	})
	if err != nil {
		return
	}

	return &_Response{
		Code:           pb.CoreServiceResponse_C_NOERROR,
		GetterResponse: &pb.GetterResponse{},
	}, nil
}
