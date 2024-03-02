package converters

import "goServer/internal/core/pb"

func ConvertMapToJoinRequestProto(mp *JoinRequest) *pb.JoinRequest {
	result := &pb.JoinRequest{
		Id:             mp.Id,
		WorkerId:       mp.WorkerId,
		OrganizationId: mp.OrganizationId,
		CreatedAt:      convertIso8601ToTimestamppb(mp.CreatedAt),
		ExpiredAt:      convertIso8601ToTimestamppb(mp.ExpiredAt),
		Flags:          mp.Flags,
		Uuid:           mp.Uuid,
	}

	return result
}

func ConvertMapToOrganizationPendingRequestRelationProto(mp *OrganizationPendingRequestRelation) *pb.OrganizationPendingRequestRelation {
	result := &pb.OrganizationPendingRequestRelation{
		Id:             mp.Id,
		OrganizationId: mp.OrganizationId,
		RequestId:      mp.RequestId,
	}

	return result
}

func ConvertMapToWorkerPendingRequestRelationProto(mp *WorkerPendingRequestRelation) *pb.WorkerPendingRequestRelation {
	result := &pb.WorkerPendingRequestRelation{
		Id:        mp.Id,
		WorkerId:  mp.WorkerId,
		RequestId: mp.RequestId,
	}

	return result
}
