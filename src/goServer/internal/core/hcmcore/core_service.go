package hcmcore

import (
	"context"
	"encoding/json"
	"goServer/internal/apiClient"
	"goServer/internal/core/converters"
	"goServer/internal/core/pb"
	goServerErrors "goServer/internal/errors"
	"log"
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
	var id, columnToSearch string
	var resp []*converters.Worker

	params := map[string]any{"userId": req.UserId, "id": req.TargetId, "uuid": req.TargetUuid}
	errMsg := "CoreServiceServer.GetWorkerById(): provide at least one of the following `userId, workerId, workerUuid`"

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
			return &_Response{}, err
		} else {
			err = client.DB.From("workers").Select().Limit(1).Eq(columnToSearch, id).Execute(&resp)
		}

	default:
		return &_Response{}, goServerErrors.ErrInvalidClientFromRequestUnimplemented
	}

	if err != nil {
		return &_Response{}, err
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

func (srv *CoreServiceServer) SaveWorker(ctx _Context, req *_Request) (res *_Response, err error) {
	errMsg := "CoreServiceServer.SaveWorker(): no target provided"

	if req.SaveWorkerTarget == nil {
		return &_Response{}, goServerErrors.ErrMissingRequestParameter(errMsg)
	}

	target := (&converters.Worker{}).TranslatePb(req.SaveWorkerTarget)
	resp := converters.Worker{}

	switch req.GetUsedClient() {
	case pb.CoreServiceRequest_C_SUPABASE:
		var client *supabaseCommunityGo.Client

		if _, client, err = srv.dependencies(req); err != nil {
			return &_Response{}, err
		}

		var b []byte

		target.LastUpdatedAt = time.Now().UTC().Format(time.RFC3339Nano)

		b, _, err = client.From("workers").Upsert(target, "username", "", "1").Single().Execute()

		if err := json.Unmarshal(b, &resp); err != nil {
			log.Panic(err)
		}
	}

	// fmt.Printf("%#v", resp.TranslatePb(&pb.Worker{}))

	if err != nil {
		return &_Response{}, err
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
		SaveWorkerResult: &pb.SaveWorkerResponse{
			UpdatedTarget: converters.ConvertMapToWorkerProto(&resp),
		},
	}, nil
}

func (srv *CoreServiceServer) SaveOrganization(ctx _Context, req *_Request) (res *_Response, err error) {
	errMsg := "CoreServiceServer.SaveOrganization(): no target provided"

	if req.SaveOrganizationTarget == nil {
		return &_Response{}, goServerErrors.ErrMissingRequestParameter(errMsg)
	}

	target := (&converters.Organization{}).TranslatePb(req.SaveOrganizationTarget)

	switch req.GetUsedClient() {
	case pb.CoreServiceRequest_C_SUPABASE:
		var client *supabaseCommunityGo.Client

		if _, client, err = srv.dependencies(req); err != nil {
			return &_Response{}, err
		}

		target.LastUpdatedAt = time.Now().UTC().Format(time.RFC3339Nano)

		_, _, err = client.From("organizations").Upsert(target, "name", "", "1").Execute()
	}

	if err != nil {
		return &_Response{}, err
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
	}, nil
}

func (srv *CoreServiceServer) SaveRole(ctx _Context, req *_Request) (res *_Response, err error) {
	errMsg := "CoreServiceServer.SaveRole(): no target provided"

	if req.SaveRoleTarget == nil {
		return &_Response{}, goServerErrors.ErrMissingRequestParameter(errMsg)
	}

	target := (&converters.Role{}).TranslatePb(req.SaveRoleTarget)

	switch req.GetUsedClient() {
	case pb.CoreServiceRequest_C_SUPABASE:
		var client *supabaseCommunityGo.Client

		if _, client, err = srv.dependencies(req); err != nil {
			return &_Response{}, err
		}

		target.LastUpdatedAt = time.Now().UTC().Format(time.RFC3339Nano)
		_, _, err = client.From("roles").Upsert(target, "uuid", "", "1").Execute()
	}

	if err != nil {
		return &_Response{}, err
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
	}, nil
}

func (srv *CoreServiceServer) SaveTeam(ctx _Context, req *_Request) (res *_Response, err error) {
	errMsg := "CoreServiceServer.SaveTeam(): no target provided"

	if req.SaveTeamTarget == nil {
		return &_Response{}, goServerErrors.ErrMissingRequestParameter(errMsg)
	}

	target := (&converters.Team{}).TranslatePb(req.SaveTeamTarget)

	switch req.GetUsedClient() {
	case pb.CoreServiceRequest_C_SUPABASE:
		var client *supabaseCommunityGo.Client

		if _, client, err = srv.dependencies(req); err != nil {
			return &_Response{}, err
		}

		target.LastUpdatedAt = time.Now().UTC().Format(time.RFC3339Nano)
		_, _, err = client.From("teams").Upsert(target, "uuid", "", "1").Execute()
	}

	if err != nil {
		return &_Response{}, err
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
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
