package converters

import (
	"goServer/internal/core/pb"

	"google.golang.org/protobuf/types/known/structpb"
)

func ConvertMapToWorkerProto(mp *Worker) *pb.Worker {
	result := &pb.Worker{
		Id:            mp.Id,
		UserId:        mp.UserId,
		CreatedById:   mp.CreatedById,
		UpdatedById:   mp.UpdatedById,
		CreatedAt:     convertIso8601ToTimestamppb(mp.CreatedAt),
		LastUpdatedAt: convertIso8601ToTimestamppb(mp.LastUpdatedAt),
		Birthdate:     convertIso8601ToTimestamppb(mp.Birthdate),
		PictureUrl:    mp.PictureUrl,
		FirstName:     mp.FirstName,
		LastName:      mp.LastName,
		MiddleName:    mp.MiddleName,
		Username:      mp.Username,
		Email:         mp.Email,
		Mobile:        mp.Mobile,
		Uuid:          mp.Uuid,
		Flags:         mp.Flags,
	}

	var structValues []*structpb.Struct

	for _, val := range mp.Addresses {
		structVal, parseError := structpb.NewStruct(val)

		if parseError != nil {
			panic(parseError)
		}

		structValues = append(structValues, structVal)
	}

	result.Addresses = structValues

	return result
}

func ConvertMapToWorkerRoleRelationProto(mp *WorkerRoleRelation) *pb.WorkerRoleRelation {
	result := &pb.WorkerRoleRelation{
		Id:             mp.Id,
		RoleId:         mp.RoleId,
		WorkerId:       mp.WorkerId,
		OrganizationId: mp.OrganizationId,
	}

	return result
}

func ConvertMapToWorkerOrganizationRelationProto(mp *WorkerOrganizationRelation) *pb.WorkerOrganizationRelation {
	result := &pb.WorkerOrganizationRelation{
		Id:             mp.Id,
		WorkerId:       mp.WorkerId,
		OrganizationId: mp.OrganizationId,
	}

	return result
}

func ConvertMapToWorkerIdentityCardRelationProto(mp *WorkerIdentityCardRelation) *pb.WorkerIdentityCardRelation {
	result := &pb.WorkerIdentityCardRelation{
		Id:      mp.Id,
		OwnerId: mp.OwnerId,
		CardId:  mp.OwnerId,
	}

	return result
}
