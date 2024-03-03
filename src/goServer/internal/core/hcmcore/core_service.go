package hcmcore

import (
	"context"
	"errors"
	"goServer/internal/apiClient"
	"goServer/internal/core/converters"
	"goServer/internal/core/pb"
	goServerErrors "goServer/internal/errors"
	"goServer/internal/messages"
	"reflect"
	"regexp"
	"strconv"
	"time"

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
			SetterResponse: &pb.SetterResponse{
				ErrorMessage: &msg,
			},
		}, err
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

func (srv *CoreServiceServer) GetWorkerById(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.GetWorkerById()"
	var id, columnToSearch string
	var resp []*converters.Worker
	var errMsg string

	getterRequest := req.GetterRequest

	if res, err := checkRequestForMissingParameters(
		getterRequest,
		messages.MessageNoRequestBodyProvided(_funcName),
	); err != nil {
		return res, err
	}

	errMsg = messages.MessageProvideAtleastOneOfTheFollowing(_funcName, []string{"userId", "targetId", "targetUuid"})
	params := map[string]any{
		"userId": getterRequest.UserId,
		"id":     getterRequest.TargetId,
		"uuid":   getterRequest.TargetUuid,
	}

	if columns, count := srv.countEmptyParameters(params); count > 1 && count == 0 {
		return &_Response{}, goServerErrors.ErrMissingRequestParameter(errMsg)
	} else {
		stringsParamPatt := regexp.MustCompile("(userId|uuid)")

		if columnToSearch = columns[0]; stringsParamPatt.MatchString(columnToSearch) {
			id = *params[columnToSearch].(*string)
		} else {
			id = strconv.FormatInt(*params[columnToSearch].(*int64), 10)
		}
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
		return &_Response{}, goServerErrors.ErrInvalidClientFromRequestUnimplemented
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

// SaveWorker saves the provided pb.Worker parameter.
func (srv *CoreServiceServer) SaveWorker(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer._SaveWorker()"

	saveReq := req.SetterRequest

	if res, err := checkRequestForMissingParameters(
		saveReq, messages.MessageNoRequestBodyProvided(_funcName),
	); err != nil {
		return res, err
	}

	if res, err := checkRequestForMissingParameters(
		saveReq, messages.MessageNoTargetProvided(_funcName),
	); err != nil {
		return res, err
	}

	target := (&converters.Worker{}).TranslatePb(saveReq.WorkerTarget)
	resp := &converters.Worker{}

	switch req.GetUsedClient() {
	case pb.CoreServiceRequest_C_SUPABASE:
		var client *supabaseCommunityGo.Client

		if _, client, err = srv.dependencies(req); err != nil {
			errMsg := err.Error()

			return &_Response{
				Code: pb.CoreServiceResponse_C_CLIENTERROR,
				SetterResponse: &pb.SetterResponse{
					ErrorMessage: &errMsg,
				},
			}, err
		}

		target.LastUpdatedAt = time.Now().UTC().Format(time.RFC3339Nano)
		query := client.From("workers").Upsert(target, "id", "", "planned").Single()

		if _, err = query.ExecuteTo(resp); err != nil {
			errMsg := err.Error()

			return &_Response{
				Code: pb.CoreServiceResponse_C_DBERROR,
				SetterResponse: &pb.SetterResponse{
					ErrorMessage: &errMsg,
				},
			}, err
		}
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
		SetterResponse: &pb.SetterResponse{
			UpdatedWorkerTarget: converters.ConvertMapToWorkerProto(resp),
		},
	}, nil
}

func (srv *CoreServiceServer) SaveOrganization(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.SaveOrganization()"

	saveReq := req.SetterRequest

	if res, err := checkRequestForMissingParameters(
		saveReq,
		messages.MessageNoRequestBodyProvided(_funcName),
	); err != nil {
		return res, err
	}

	if res, err := checkRequestForMissingParameters(
		saveReq.OrganizationTarget,
		messages.MessageNoTargetProvided(_funcName),
	); err != nil {
		return res, err
	}

	target := (&converters.Organization{}).TranslatePb(req.SetterRequest.OrganizationTarget)
	resp := &converters.Organization{}

	switch req.GetUsedClient() {
	case pb.CoreServiceRequest_C_SUPABASE:
		var client *supabaseCommunityGo.Client

		if _, client, err = srv.dependencies(req); err != nil {
			errMsg := err.Error()

			return &_Response{
				Code: pb.CoreServiceResponse_C_CLIENTERROR,
				SetterResponse: &pb.SetterResponse{
					ErrorMessage: &errMsg,
				},
			}, err
		}

		target.LastUpdatedAt = time.Now().UTC().Format(time.RFC3339Nano)
		query := client.From("organizations").Upsert(target, "id", "", "planned").Single()

		if _, err = query.ExecuteTo(resp); err != nil {
			errMsg := err.Error()

			return &_Response{
				Code: pb.CoreServiceResponse_C_DBERROR,
				SetterResponse: &pb.SetterResponse{
					ErrorMessage: &errMsg,
				},
			}, err
		}
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
		SetterResponse: &pb.SetterResponse{
			UpdatedOrganizationTarget: converters.ConvertMapToOrganizationProto(resp),
		},
	}, nil
}

func (srv *CoreServiceServer) SaveRole(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.SaveRole()"

	saveReq := req.SetterRequest

	if res, err := checkRequestForMissingParameters(
		saveReq,
		messages.MessageNoRequestBodyProvided(_funcName),
	); err != nil {
		return res, err
	}

	if res, err := checkRequestForMissingParameters(
		saveReq.RoleTarget,
		messages.MessageNoTargetProvided(_funcName),
	); err != nil {
		return res, err
	}

	errMsg := messages.MessageRequiredFieldNotProvided(_funcName, "organizationId")
	if saveReq.RoleTarget.GetOrganizationId() == 0 {
		return &_Response{
			SetterResponse: &pb.SetterResponse{
				ErrorMessage: &errMsg,
			},
		}, errors.New(errMsg)
	}

	target := (&converters.Role{}).TranslatePb(req.SetterRequest.RoleTarget)
	resp := &converters.Role{}

	switch req.GetUsedClient() {
	case pb.CoreServiceRequest_C_SUPABASE:
		var client *supabaseCommunityGo.Client

		if _, client, err = srv.dependencies(req); err != nil {
			errMsg := err.Error()

			return &_Response{
				Code: pb.CoreServiceResponse_C_CLIENTERROR,
				SetterResponse: &pb.SetterResponse{
					ErrorMessage: &errMsg,
				},
			}, err
		}

		target.LastUpdatedAt = time.Now().UTC().Format(time.RFC3339Nano)

		query := client.From("roles").Upsert(target, "id", "", "planned").Single()

		if _, err = query.ExecuteTo(resp); err != nil {
			errMsg := err.Error()

			return &_Response{
				Code: pb.CoreServiceResponse_C_DBERROR,
				SetterResponse: &pb.SetterResponse{
					ErrorMessage: &errMsg,
				},
			}, err
		}
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
		SetterResponse: &pb.SetterResponse{
			UpdatedRoleTarget: converters.ConvertMapToRoleProto(resp),
		},
	}, nil
}

func (srv *CoreServiceServer) SaveTeam(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.SaveTeam()"

	saveReq := req.SetterRequest

	if res, err := checkRequestForMissingParameters(
		saveReq,
		messages.MessageNoRequestBodyProvided(_funcName),
	); err != nil {
		return res, err
	}

	if res, err := checkRequestForMissingParameters(
		saveReq.TeamTarget,
		messages.MessageNoTargetProvided(_funcName),
	); err != nil {
		return res, err
	}

	errMsg := messages.MessageRequiredFieldNotProvided(_funcName, "organizationId")
	if saveReq.TeamTarget.GetOrganizationId() == 0 {
		return &_Response{
			SetterResponse: &pb.SetterResponse{
				ErrorMessage: &errMsg,
			},
		}, errors.New(errMsg)
	}

	target := (&converters.Team{}).TranslatePb(saveReq.TeamTarget)
	resp := &converters.Team{}

	switch req.GetUsedClient() {
	case pb.CoreServiceRequest_C_SUPABASE:
		var client *supabaseCommunityGo.Client

		if _, client, err = srv.dependencies(req); err != nil {
			errMsg := err.Error()

			return &_Response{
				Code: pb.CoreServiceResponse_C_CLIENTERROR,
				SetterResponse: &pb.SetterResponse{
					ErrorMessage: &errMsg,
				},
			}, err
		}

		target.LastUpdatedAt = time.Now().UTC().Format(time.RFC3339Nano)
		query := client.From("teams").Upsert(target, "id", "", "planned").Single()

		if _, err := query.ExecuteTo(resp); err != nil {
			errMsg := err.Error()

			return &_Response{
				Code: pb.CoreServiceResponse_C_DBERROR,
				SetterResponse: &pb.SetterResponse{
					ErrorMessage: &errMsg,
				},
			}, err
		}
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
		SetterResponse: &pb.SetterResponse{
			UpdatedTeamTarget: converters.ConvertMapToTeamProto(resp),
		},
	}, nil
}

func (srv *CoreServiceServer) DeleteWorkerById(ctx _Context, req *_Request) (res *_Response, err error) {

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
	}, nil
}

func (srv *CoreServiceServer) DeleteOrganizationById(ctx _Context, req *_Request) (res *_Response, err error) {

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
	}, nil
}

func (srv *CoreServiceServer) DeleteRoleById(ctx _Context, req *_Request) (res *_Response, err error) {

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
	}, nil
}

func (srv *CoreServiceServer) DeleteTeamById(ctx _Context, req *_Request) (res *_Response, err error) {

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
	}, nil
}
