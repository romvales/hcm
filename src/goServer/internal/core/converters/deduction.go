package converters

import "goServer/internal/core/pb"

func ConvertMapToDeductionProto(mp *Deduction) *pb.Deduction {
	result := &pb.Deduction{
		Id:            mp.Id,
		CreatedById:   mp.CreatedById,
		UpdatedById:   mp.UpdatedById,
		WorkerId:      mp.WorkerId,
		CreatedAt:     convertIso8601ToTimestamppb(mp.CreatedAt),
		LastUpdatedAt: convertIso8601ToTimestamppb(mp.LastUpdatedAt),
		EffectiveAt:   convertIso8601ToTimestamppb(mp.EffectiveAt),
		Name:          mp.Name,
		Value:         float32(mp.Value),
		Flags:         mp.Flags,
		Uuid:          mp.Uuid,
	}

	return result
}
