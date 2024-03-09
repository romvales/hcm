package hcmcore

import (
	"goServer/internal/core/converters"
	"goServer/internal/core/pb"
	"time"
)

func (srv *CoreServiceServer) SendJoinRequest(ctx _Context, req *_Request) (res *_Response, err error) {
	var _funcName = "CoreServiceServer.SendJoinRequest()"

	if res, err = checkIfHasValidRequestParams(_funcName, req, Q_SETTER); err != nil {
		return
	}

	setterReq := req.SetterRequest

	senderId := setterReq.GetRequestSenderId()
	targetId := setterReq.GetTargetId()

	if senderId == 0 || targetId == 0 {
		return srv.someRequiredFieldAreNotProvided(_funcName, "requestSenderId, targetId")
	}

	// TODO: After adding a configuration system, make sure to set the expiration of the request
	//       based on the set maximum expiration request time of the organization.
	target := &converters.JoinRequest{
		ExpiredAt: time.Now().Add((time.Hour * 24) * 3 /* 3 days */).Format(time.RFC3339Nano),
		Flags:     uint32(pb.JoinRequest_S_WAIT),
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
