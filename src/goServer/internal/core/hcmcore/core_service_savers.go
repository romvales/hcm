package hcmcore

import (
	"errors"
	"goServer/internal/core/converters"
	"goServer/internal/core/pb"
	"goServer/internal/messages"
	"reflect"
	"time"

	goServerErrors "goServer/internal/errors"

	"github.com/google/uuid"
	supabaseCommunityGo "github.com/supabase-community/supabase-go"
)

func (srv *CoreServiceServer) saveUpsertDataToTable(
	req *_Request,
	target any,
	resp any,
	tableName string) (res *_Response, err error) {

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

		ref := reflect.ValueOf(target)
		uuidField := ref.Elem().FieldByName("Uuid")

		// When the uuid column is empty, make sure to set a uuid for it.
		if uuidField.IsZero() {
			uuidField.SetString(uuid.NewString())
		}

		lastUpdatedAtField := ref.Elem().FieldByName("LastUpdatedAt")
		lastUpdatedAtField.SetString(time.Now().UTC().Format(time.RFC3339Nano))

		query := client.From(tableName).Upsert(target, "id", "", "planned").Single()

		if _, err = query.ExecuteTo(resp); err != nil {
			errMsg := err.Error()

			return &_Response{
				Code: pb.CoreServiceResponse_C_DBERROR,
				SetterResponse: &pb.SetterResponse{
					ErrorMessage: &errMsg,
				},
			}, err
		}
	default:
		errMsg := goServerErrors.ErrInvalidClientFromRequestUnimplemented.Error()

		return &_Response{
			Code: pb.CoreServiceResponse_C_CLIENTERROR,
			SetterResponse: &pb.SetterResponse{
				ErrorMessage: &errMsg,
			},
		}, goServerErrors.ErrInvalidClientFromRequestUnimplemented
	}

	return nil, nil
}

// SaveWorker saves the provided pb.Worker parameter.
func (srv *CoreServiceServer) SaveWorker(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.SaveWorker()"

	if res, err := checkIfHasValidRequestParams(_funcName, req, "setter"); err != nil {
		return res, err
	}

	saveReq := req.SetterRequest

	if res, err := checkRequestForMissingParameters(
		saveReq.WorkerTarget, messages.MessageNoTargetProvided(_funcName),
	); err != nil {
		return res, err
	}

	target := (&converters.Worker{}).TranslatePb(saveReq.WorkerTarget)
	resp := &converters.Worker{}

	if res, err := srv.saveUpsertDataToTable(req, target, resp, "workers"); err != nil {
		return res, err
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

	if res, err := checkIfHasValidRequestParams(_funcName, req, "setter"); err != nil {
		return res, err
	}

	saveReq := req.SetterRequest

	if res, err := checkRequestForMissingParameters(
		saveReq.OrganizationTarget,
		messages.MessageNoTargetProvided(_funcName),
	); err != nil {
		return res, err
	}

	target := (&converters.Organization{}).TranslatePb(req.SetterRequest.OrganizationTarget)
	resp := &converters.Organization{}

	if res, err := srv.saveUpsertDataToTable(req, target, resp, "organizations"); err != nil {
		return res, err
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

	if res, err := checkIfHasValidRequestParams(_funcName, req, "setter"); err != nil {
		return res, err
	}

	saveReq := req.SetterRequest

	if res, err := checkRequestForMissingParameters(
		saveReq.RoleTarget,
		messages.MessageNoTargetProvided(_funcName),
	); err != nil {
		return res, err
	}

	if saveReq.RoleTarget.GetOrganizationId() == 0 {
		errMsg := messages.MessageRequiredFieldNotProvided(_funcName, "organizationId")

		return &_Response{
			SetterResponse: &pb.SetterResponse{
				ErrorMessage: &errMsg,
			},
		}, errors.New(errMsg)
	}

	target := (&converters.Role{}).TranslatePb(req.SetterRequest.RoleTarget)
	resp := &converters.Role{}

	if res, err := srv.saveUpsertDataToTable(req, target, resp, "roles"); err != nil {
		return res, err
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

	if res, err := checkIfHasValidRequestParams(_funcName, req, "setter"); err != nil {
		return res, err
	}

	saveReq := req.SetterRequest

	if res, err := checkRequestForMissingParameters(
		saveReq.TeamTarget,
		messages.MessageNoTargetProvided(_funcName),
	); err != nil {
		return res, err
	}

	if saveReq.TeamTarget.GetOrganizationId() == 0 {
		errMsg := messages.MessageRequiredFieldNotProvided(_funcName, "organizationId")

		return &_Response{
			Code: pb.CoreServiceResponse_C_MISSINGPARAMETERS,
			SetterResponse: &pb.SetterResponse{
				ErrorMessage: &errMsg,
			},
		}, errors.New(errMsg)
	}

	target := (&converters.Team{}).TranslatePb(saveReq.TeamTarget)
	resp := &converters.Team{}

	if res, err := srv.saveUpsertDataToTable(req, target, resp, "teams"); err != nil {
		return res, err
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
		SetterResponse: &pb.SetterResponse{
			UpdatedTeamTarget: converters.ConvertMapToTeamProto(resp),
		},
	}, nil
}

func (srv *CoreServiceServer) SaveWorkerIdentityCard(cx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.SaveWorkerIdentityCard()"

	if res, err := checkIfHasValidRequestParams(_funcName, req, "setter"); err != nil {
		return res, err
	}

	saveReq := req.SetterRequest

	if res, err := checkRequestForMissingParameters(
		saveReq.IdentityCardTarget,
		messages.MessageNoTargetProvided(_funcName),
	); err != nil {
		return res, err
	}

	idTarget := saveReq.IdentityCardTarget

	if idTarget.GetWorkerId() == 0 {
		errMsg := messages.MessageRequiredFieldNotProvided(_funcName, "workerId")

		return &_Response{
			Code: pb.CoreServiceResponse_C_MISSINGPARAMETERS,
			SetterResponse: &pb.SetterResponse{
				ErrorMessage: &errMsg,
			},
		}, errors.New(errMsg)
	}

	if idTarget.GetFrontImageUrl() == "" && idTarget.GetBackImageUrl() == "" {
		errMsg := messages.MessageRequiredFieldNotProvided(_funcName, "frontImageUrl, backImageUrl")

		return &_Response{
			Code: pb.CoreServiceResponse_C_MISSINGPARAMETERS,
			SetterResponse: &pb.SetterResponse{
				ErrorMessage: &errMsg,
			},
		}, errors.New(errMsg)
	}

	target := (&converters.WorkerIdentityCard{}).TranslatePb(idTarget)
	resp := &converters.WorkerIdentityCard{}

	if res, err := srv.saveUpsertDataToTable(req, target, resp, "workerIdentityCards"); err != nil {
		return res, err
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
		SetterResponse: &pb.SetterResponse{
			UpdatedIdentityCardTarget: converters.ConvertMapToWorkerIdentityCardProto(resp),
		},
	}, nil
}

func (srv *CoreServiceServer) SaveMember(cx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.SaveMember()"

	if res, err := checkIfHasValidRequestParams(_funcName, req, "setter"); err != nil {
		return res, err
	}

	saveReq := req.SetterRequest

	if res, err := checkRequestForMissingParameters(
		saveReq.MemberTarget,
		messages.MessageNoTargetProvided(_funcName),
	); err != nil {
		return res, err
	}

	target := (&converters.Member{}).TranslatePb(saveReq.MemberTarget)
	resp := &converters.Member{}

	if res, err := srv.saveUpsertDataToTable(req, target, resp, "organizationsMembers"); err != nil {
		return res, err
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
		SetterResponse: &pb.SetterResponse{
			UpdatedMemberTarget: converters.ConvertMapToMemberProto(resp),
		},
	}, nil
}

func (srv *CoreServiceServer) SaveCompensation(cx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.SaveCompensation()"

	if res, err := checkIfHasValidRequestParams(_funcName, req, "setter"); err != nil {
		return res, err
	}

	saveReq := req.SetterRequest

	if res, err := checkRequestForMissingParameters(
		saveReq.CompensationTarget,
		messages.MessageNoTargetProvided(_funcName),
	); err != nil {
		return res, err
	}

	target := (&converters.Compensation{}).TranslatePb(saveReq.CompensationTarget)
	resp := &converters.Compensation{}

	if res, err := srv.saveUpsertDataToTable(req, target, resp, "compensations"); err != nil {
		return res, err
	}

	return &_Response{
		Code:           pb.CoreServiceResponse_C_NOERROR,
		SetterResponse: &pb.SetterResponse{},
	}, nil
}

func (srv *CoreServiceServer) SaveAddition(cx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.SaveAddition()"

	if res, err := checkIfHasValidRequestParams(_funcName, req, "setter"); err != nil {
		return res, err
	}

	saveReq := req.SetterRequest

	if res, err := checkRequestForMissingParameters(
		saveReq.AdditionTarget,
		messages.MessageNoTargetProvided(_funcName),
	); err != nil {
		return res, err
	}

	target := (&converters.Addition{}).TranslatePb(saveReq.AdditionTarget)
	resp := &converters.Addition{}

	if res, err := srv.saveUpsertDataToTable(req, target, resp, "additions"); err != nil {
		return res, err
	}

	return &_Response{
		Code:           pb.CoreServiceResponse_C_NOERROR,
		SetterResponse: &pb.SetterResponse{},
	}, nil
}

func (srv *CoreServiceServer) SaveDeduction(cx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.SaveDeduction()"

	if res, err := checkIfHasValidRequestParams(_funcName, req, "setter"); err != nil {
		return res, err
	}

	saveReq := req.SetterRequest

	if res, err := checkRequestForMissingParameters(
		saveReq.DeductionTarget,
		messages.MessageNoTargetProvided(_funcName),
	); err != nil {
		return res, err
	}

	target := (&converters.Deduction{}).TranslatePb(saveReq.DeductionTarget)
	resp := &converters.Deduction{}

	if res, err := srv.saveUpsertDataToTable(req, target, resp, "deductions"); err != nil {
		return res, err
	}

	return &_Response{
		Code:           pb.CoreServiceResponse_C_NOERROR,
		SetterResponse: &pb.SetterResponse{},
	}, nil
}

func (srv *CoreServiceServer) SavePayroll(cx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.SavePayroll()"

	if res, err := checkIfHasValidRequestParams(_funcName, req, "setter"); err != nil {
		return res, err
	}

	saveReq := req.SetterRequest

	if res, err := checkRequestForMissingParameters(
		saveReq.PayrollTarget,
		messages.MessageNoTargetProvided(_funcName),
	); err != nil {
		return res, err
	}

	target := (&converters.Payroll{}).TranslatePb(saveReq.PayrollTarget)
	resp := &converters.Payroll{}

	if res, err := srv.saveUpsertDataToTable(req, target, resp, "payrolls"); err != nil {
		return res, err
	}

	return &_Response{
		Code:           pb.CoreServiceResponse_C_NOERROR,
		SetterResponse: &pb.SetterResponse{},
	}, nil
}

func (srv *CoreServiceServer) SaveShift(cx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.SaveShift()"

	if res, err := checkIfHasValidRequestParams(_funcName, req, "setter"); err != nil {
		return res, err
	}

	saveReq := req.SetterRequest

	if res, err := checkRequestForMissingParameters(
		saveReq.ShiftTarget,
		messages.MessageNoTargetProvided(_funcName),
	); err != nil {
		return res, err
	}

	if saveReq.GetTargetShiftType() == pb.SetterRequest_T_OVERRIDESHIFT {
		target := (&converters.OverrideShift{}).TranslatePb(saveReq.OverrideShiftTarget)
		resp := &converters.OverrideShift{}

		res, err = srv.saveUpsertDataToTable(req, target, resp, "overrideShifts")
		if err != nil {
			return
		}

		return &_Response{
			Code: pb.CoreServiceResponse_C_NOERROR,
			SetterResponse: &pb.SetterResponse{
				UpdatedOverrideShiftTarget: converters.ConvertMapToOverrideShiftProto(resp),
			},
		}, nil
	} else {
		target := (&converters.Shift{}).TranslatePb(saveReq.ShiftTarget)
		resp := &converters.Shift{}

		res, err = srv.saveUpsertDataToTable(req, target, resp, "shifts")
		if err != nil {
			return
		}

		return &_Response{
			Code: pb.CoreServiceResponse_C_NOERROR,
			SetterResponse: &pb.SetterResponse{
				UpdatedShiftTarget: converters.ConvertMapToShiftProto(resp),
			},
		}, nil
	}

	// unreachable
}
