package converters

import (
	"goServer/internal/core/pb"
	"log"

	"google.golang.org/protobuf/types/known/structpb"
)

func ConvertMapToWorkerIdentityCardProto(mp *WorkerIdentityCard) *pb.WorkerIdentityCard {
	result := &pb.WorkerIdentityCard{
		Id:            mp.Id,
		WorkerId:      mp.WorkerId,
		CreatedById:   mp.CreatedById,
		UpdatedById:   mp.UpdatedById,
		CreatedAt:     convertIso8601ToTimestamppb(mp.CreatedAt),
		LastUpdatedAt: convertIso8601ToTimestamppb(mp.LastUpdatedAt),
		FrontImageUrl: mp.FrontImageUrl,
		BackImageUrl:  mp.BackImageUrl,
		Name:          mp.Name,
		Flags:         mp.Flags,
		Uuid:          mp.Uuid,
	}

	extractedInfo, parseError := structpb.NewStruct(mp.ExtractedInfo)

	if parseError != nil {
		log.Panicln(parseError)
	}

	result.ExtractedInfo = extractedInfo

	return result
}

func ConvertMapToVerifiedWorkerIdentityCardRelation(mp *VerifiedWorkerIdentityCardRelation) *pb.VerifiedWorkerIdentityCardRelation {
	result := &pb.VerifiedWorkerIdentityCardRelation{
		Id:           mp.Id,
		OwnerId:      mp.OwnerId,
		VerifiedById: mp.VerifiedById,
		CardId:       mp.CardId,
	}

	return result
}
