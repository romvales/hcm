package hcmcore

import (
	"context"
	"goServer/internal/core/converters"
	"goServer/internal/core/pb"
	"reflect"
	"time"

	"github.com/google/uuid"
	supabaseCommunityGo "github.com/supabase-community/supabase-go"
)

func (srv *CoreServiceServer) saveUpsertDataToTable(req *_Request, target any, resp any, tableName string) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.saveUpsertDataToTable()"

	return srv.queryItemById(context.Background(), CoreServiceGetQueryParams{
		Query:    Q_SETTER,
		FuncName: _funcName,
		Req:      req,
		Resp:     resp,

		// Disable the default req parameters check for this query.
		DisableReqParamsCheck: true,
		Callback: DatabaseActionCallback{
			UseSupabaseCommunityClient: true,
			SupabaseCallback: func(ctx _Context, req *_Request, resp, client any) (*_Response, error) {
				supabaseClient := client.(*supabaseCommunityGo.Client)

				ref := reflect.ValueOf(target)
				uuidField := ref.Elem().FieldByName("Uuid")
				lastUpdatedAtField := ref.Elem().FieldByName("LastUpdatedAt")

				// When the uuid column is empty, make sure to set a uuid for it.
				// It is assumed that when the uuid column is empty, it is a new data.
				if uuidField.CanSet() && uuidField.IsZero() {
					createdAt := ref.Elem().FieldByName("CreatedAt")
					createdAt.SetString(time.Now().UTC().Format(time.RFC3339Nano))

					uuidField.SetString(uuid.NewString())
				}

				if lastUpdatedAtField.CanSet() {
					lastUpdatedAtField.SetString(time.Now().UTC().Format(time.RFC3339Nano))
				}

				query := supabaseClient.From(tableName).Upsert(target, "id", "", "planned").Single()

				if _, err = query.ExecuteTo(resp); err != nil {
					return setupErrorResponse(err, pb.CoreServiceResponse_C_DBERROR, Q_SETTER)
				}

				return nil, nil
			},
		},
	})
}

// SaveWorker saves the provided pb.Worker parameter.
func (srv *CoreServiceServer) SaveWorker(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.SaveWorker()"

	if res, err := checkIfHasValidRequestParams(_funcName, req, Q_SETTER); err != nil {
		return res, err
	}

	saveReq := req.SetterRequest
	workerTarget := saveReq.WorkerTarget

	if res, err := checkIfRequestTargetIsMissing(_funcName, workerTarget); err != nil {
		return res, err
	}

	target := (&converters.Worker{}).TranslatePb(workerTarget)
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

	if res, err := checkIfHasValidRequestParams(_funcName, req, Q_SETTER); err != nil {
		return res, err
	}

	saveReq := req.SetterRequest
	organizationTarget := saveReq.OrganizationTarget

	if res, err := checkIfRequestTargetIsMissing(_funcName, organizationTarget); err != nil {
		return res, err
	}

	target := (&converters.Organization{}).TranslatePb(organizationTarget)
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

	if res, err := checkIfHasValidRequestParams(_funcName, req, Q_SETTER); err != nil {
		return res, err
	}

	saveReq := req.SetterRequest
	roleTarget := saveReq.RoleTarget

	if res, err := checkIfRequestTargetIsMissing(_funcName, roleTarget); err != nil {
		return res, err
	}

	if saveReq.RoleTarget.GetOrganizationId() == 0 {
		return srv.someRequiredFieldAreNotProvided(_funcName, "organizationId")
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

	if res, err := checkIfHasValidRequestParams(_funcName, req, Q_SETTER); err != nil {
		return res, err
	}

	saveReq := req.SetterRequest
	teamTarget := saveReq.TeamTarget

	if res, err := checkIfRequestTargetIsMissing(_funcName, teamTarget); err != nil {
		return res, err
	}

	if saveReq.TeamTarget.GetOrganizationId() == 0 {
		return srv.someRequiredFieldAreNotProvided(_funcName, "organizationId")
	}

	target := (&converters.Team{}).TranslatePb(teamTarget)
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

func (srv *CoreServiceServer) SaveWorkerIdentityCard(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.SaveWorkerIdentityCard()"

	if res, err := checkIfHasValidRequestParams(_funcName, req, Q_SETTER); err != nil {
		return res, err
	}

	saveReq := req.SetterRequest

	idTarget := saveReq.IdentityCardTarget

	if res, err := checkIfRequestTargetIsMissing(_funcName, idTarget); err != nil {
		return res, err
	}

	if idTarget.GetWorkerId() == 0 {
		return srv.someRequiredFieldAreNotProvided(_funcName, "workerId")
	}

	if idTarget.GetFrontImageUrl() == "" && idTarget.GetBackImageUrl() == "" {
		return srv.someRequiredFieldAreNotProvided(_funcName, "frontImageUrl, backImageUrl")
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

func (srv *CoreServiceServer) SaveMember(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.SaveMember()"

	if res, err := checkIfHasValidRequestParams(_funcName, req, Q_SETTER); err != nil {
		return res, err
	}

	saveReq := req.SetterRequest
	memberTarget := saveReq.MemberTarget

	if res, err := checkIfRequestTargetIsMissing(_funcName, memberTarget); err != nil {
		return res, err
	}

	if memberTarget.GetOrganizationId() == 0 || memberTarget.GetWorkerId() == 0 {
		return srv.someRequiredFieldAreNotProvided(_funcName, "organizationId, workerId")
	}

	target := (&converters.Member{}).TranslatePb(memberTarget)
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

func (srv *CoreServiceServer) SaveCompensation(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.SaveCompensation()"

	if res, err := checkIfHasValidRequestParams(_funcName, req, Q_SETTER); err != nil {
		return res, err
	}

	saveReq := req.SetterRequest
	compensationTarget := saveReq.CompensationTarget

	if res, err := checkIfRequestTargetIsMissing(_funcName, compensationTarget); err != nil {
		return res, err
	}

	if compensationTarget.GetOrganizationId() == 0 || compensationTarget.GetWorkerId() == 0 {
		return srv.someRequiredFieldAreNotProvided(_funcName, "organizationId, workerId")
	}

	target := (&converters.Compensation{}).TranslatePb(compensationTarget)
	resp := &converters.Compensation{}

	if res, err := srv.saveUpsertDataToTable(req, target, resp, "compensations"); err != nil {
		return res, err
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
		SetterResponse: &pb.SetterResponse{
			UpdatedCompensationTarget: converters.ConvertMapToCompensationProto(resp),
		},
	}, nil
}

func (srv *CoreServiceServer) SaveAddition(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.SaveAddition()"

	if res, err := checkIfHasValidRequestParams(_funcName, req, Q_SETTER); err != nil {
		return res, err
	}

	saveReq := req.SetterRequest
	additionTarget := saveReq.AdditionTarget

	if res, err := checkIfRequestTargetIsMissing(_funcName, additionTarget); err != nil {
		return res, err
	}

	target := (&converters.Addition{}).TranslatePb(additionTarget)
	resp := &converters.Addition{}

	if res, err := srv.saveUpsertDataToTable(req, target, resp, "additions"); err != nil {
		return res, err
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
		SetterResponse: &pb.SetterResponse{
			UpdatedAdditionTarget: converters.ConvertMapToAdditionProto(resp),
		},
	}, nil
}

func (srv *CoreServiceServer) SaveDeduction(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.SaveDeduction()"

	if res, err := checkIfHasValidRequestParams(_funcName, req, Q_SETTER); err != nil {
		return res, err
	}

	saveReq := req.SetterRequest
	deductionTarget := saveReq.DeductionTarget

	if res, err := checkIfRequestTargetIsMissing(_funcName, deductionTarget); err != nil {
		return res, err
	}

	target := (&converters.Deduction{}).TranslatePb(deductionTarget)
	resp := &converters.Deduction{}

	if res, err := srv.saveUpsertDataToTable(req, target, resp, "deductions"); err != nil {
		return res, err
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
		SetterResponse: &pb.SetterResponse{
			UpdatedDeductionTarget: converters.ConvertMapToDeductionProto(resp),
		},
	}, nil
}

func (srv *CoreServiceServer) SavePayroll(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.SavePayroll()"

	if res, err := checkIfHasValidRequestParams(_funcName, req, Q_SETTER); err != nil {
		return res, err
	}

	saveReq := req.SetterRequest
	payrollTarget := saveReq.PayrollTarget

	if res, err := checkIfRequestTargetIsMissing(_funcName, payrollTarget); err != nil {
		return res, err
	}

	if payrollTarget.GetOrganizationId() == 0 {
		return srv.someRequiredFieldAreNotProvided(_funcName, "organizationId")
	}

	target := (&converters.Payroll{}).TranslatePb(payrollTarget)
	resp := &converters.Payroll{}

	if res, err := srv.saveUpsertDataToTable(req, target, resp, "payrolls"); err != nil {
		return res, err
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
		SetterResponse: &pb.SetterResponse{
			UpdatedPayrollTarget: converters.ConvertMapToPayrollProto(resp),
		},
	}, nil
}

func (srv *CoreServiceServer) SaveAttendance(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.SaveAttendance()"

	if res, err := checkIfHasValidRequestParams(_funcName, req, Q_SETTER); err != nil {
		return res, err
	}

	saveReq := req.SetterRequest
	attendanceTarget := saveReq.AttendanceTarget

	if res, err := checkIfRequestTargetIsMissing(_funcName, attendanceTarget); err != nil {
		return res, err
	}

	if attendanceTarget.GetWorkerId() == 0 &&
		(attendanceTarget.ShiftId == nil || attendanceTarget.OshiftId == nil) &&
		attendanceTarget.ClockIn == nil {
		return srv.someRequiredFieldAreNotProvided(_funcName, "workerId, shiftId|oshiftId, clockIn")
	}

	target := (&converters.Attendance{}).TranslatePb(attendanceTarget)
	resp := &converters.Attendance{}

	if res, err = srv.saveUpsertDataToTable(req, target, resp, "attendances"); err != nil {
		return
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
		SetterResponse: &pb.SetterResponse{
			UpdatedAttendanceTarget: converters.ConvertMapToAttendanceProto(resp),
		},
	}, nil
}

func (srv *CoreServiceServer) SaveShift(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.SaveShift()"

	if res, err := checkIfHasValidRequestParams(_funcName, req, Q_SETTER); err != nil {
		return res, err
	}

	saveReq := req.SetterRequest

	if res, err := checkIfRequestTargetIsMissing(_funcName, saveReq); err != nil {
		return res, err
	}

	if saveReq.GetTargetShiftType() == pb.SetterRequest_T_OVERRIDESHIFT {
		overrideShiftTarget := saveReq.OverrideShiftTarget

		if res, err := checkIfRequestTargetIsMissing(_funcName, overrideShiftTarget); err != nil {
			return res, err
		}

		if overrideShiftTarget.GetOverrideClockIn() == nil ||
			overrideShiftTarget.GetOverrideClockOut() == nil ||
			overrideShiftTarget.GetOrganizationId() == 0 {
			return srv.someRequiredFieldAreNotProvided(_funcName, "organizationId, overrideClockIn, overrideClockOut")
		}

		if overrideShiftTarget.GetName() == "" {
			return srv.someRequiredFieldAreNotProvided(_funcName, "name")
		}

		if overrideShiftTarget.GroupId == "" {
			overrideShiftTarget.GroupId = uuid.NewString()
		}

		target := (&converters.OverrideShift{}).TranslatePb(overrideShiftTarget)
		resp := &converters.OverrideShift{}

		if res, err = srv.saveUpsertDataToTable(req, target, resp, "overrideShifts"); err != nil {
			return
		}

		return &_Response{
			Code: pb.CoreServiceResponse_C_NOERROR,
			SetterResponse: &pb.SetterResponse{
				UpdatedOverrideShiftTarget: converters.ConvertMapToOverrideShiftProto(resp),
			},
		}, nil
	} else {
		shiftTarget := saveReq.ShiftTarget

		if res, err := checkIfRequestTargetIsMissing(_funcName, shiftTarget); err != nil {
			return res, err
		}

		if shiftTarget.GetClockIn() == nil ||
			shiftTarget.GetClockOut() == nil ||
			shiftTarget.GetOrganizationId() == 0 {
			return srv.someRequiredFieldAreNotProvided(_funcName, "organizationId, clockIn, clockOut")
		}

		if shiftTarget.GetName() == "" {
			return srv.someRequiredFieldAreNotProvided(_funcName, "name")
		}

		if shiftTarget.GroupId == "" {
			shiftTarget.GroupId = uuid.NewString()
		}

		target := (&converters.Shift{}).TranslatePb(shiftTarget)
		resp := &converters.Shift{}

		if res, err = srv.saveUpsertDataToTable(req, target, resp, "standardShifts"); err != nil {
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

func (srv *CoreServiceServer) SaveJoinRequest(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.SaveJoinRequest()"

	if res, err = checkIfHasValidRequestParams(_funcName, req, Q_SETTER); err != nil {
		return
	}

	setterReq := req.SetterRequest

	senderId := setterReq.GetRequestSenderId()
	targetId := setterReq.GetTargetId()
	senderType := setterReq.GetRequestSenderType()

	if senderId == 0 || targetId == 0 {
		return srv.someRequiredFieldAreNotProvided(_funcName, "requestSenderId, targetId")
	}

	// TODO: After adding a configuration system, make sure to set the expiration of the request
	//       based on the set maximum expiration request time of the organization.
	target := &converters.JoinRequest{
		ExpiredAt:  time.Now().Add((time.Hour * 24) * 3 /* 3 days */).Format(time.RFC3339Nano),
		Flags:      uint32(pb.JoinRequest_S_WAIT),
		SenderType: int32(senderType),
	}

	resp := &converters.JoinRequest{}

	switch setterReq.GetRequestSenderType() {
	case pb.JoinRequest_T_ORGANIZATION:
		target.OrganizationId = setterReq.GetRequestSenderId()
		target.WorkerId = setterReq.GetTargetId()

	case pb.JoinRequest_T_WORKER:
		target.OrganizationId = setterReq.GetTargetId()
		target.WorkerId = setterReq.GetRequestSenderId()
	}

	if res, err = srv.saveUpsertDataToTable(req, target, resp, "pendingJoinRequests"); err != nil {
		return
	}

	return &_Response{
		Code: pb.CoreServiceResponse_C_NOERROR,
		SetterResponse: &pb.SetterResponse{
			JoinRequestResult: converters.ConvertMapToJoinRequestProto(resp),
		},
	}, nil
}
