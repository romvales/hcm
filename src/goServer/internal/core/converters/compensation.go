package converters

import "goServer/internal/core/pb"

func ConvertMapToCompensationProto(mp *Compensation) *pb.Compensation {
	result := &pb.Compensation{
		Id:            mp.Id,
		CreatedById:   mp.CreatedById,
		UpdatedById:   mp.UpdatedById,
		CreatedAt:     convertIso8601ToTimestamppb(mp.CreatedAt),
		LastUpdatedAt: convertIso8601ToTimestamppb(mp.LastUpdatedAt),
		PaidAt:        convertIso8601ToTimestamppb(mp.PaidAt),
		ApprovedAt:    convertIso8601ToTimestamppb(mp.ApprovedAt),
		RejectedAt:    convertIso8601ToTimestamppb(mp.RejectedAt),
		PeriodStart:   convertIso8601ToTimestamppb(mp.PeriodStart),
		PeriodEnd:     convertIso8601ToTimestamppb(mp.PeriodEnd),
		Gvalue:        float32(mp.Gvalue),
		Avalue:        float32(mp.Avalue),
		Dvalue:        float32(mp.Dvalue),
		Value:         float32(mp.Value),
		Flags:         mp.Flags,
		Uuid:          mp.Uuid,
	}

	return result
}

func ConvertMapToCompensationAdditionRelationProto(mp *CompensationAdditionRelation) *pb.CompensationAdditionRelation {
	result := &pb.CompensationAdditionRelation{
		Id:             mp.Id,
		CompensationId: mp.CompensationId,
		AdditionId:     mp.AdditionId,
	}

	return result
}

func ConvertMapToCompensationDeductionRelationProto(mp *CompensationDeductionRelation) *pb.CompensationDeductionRelation {
	result := &pb.CompensationDeductionRelation{
		Id:             mp.Id,
		CompensationId: mp.CompensationId,
		DeductionId:    mp.DeductionId,
	}

	return result
}
