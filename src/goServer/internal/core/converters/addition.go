package converters

import "goServer/internal/core/pb"

func ConvertMapToAdditionProto(mp *Addition) *pb.Addition {
	result := &pb.Addition{
		Id:            mp.Id,
		CreatedById:   mp.CreatedById,
		UpdatedByid:   mp.UpdatedByid,
		WorkerId:      mp.WorkerId,
		CreatedAt:     convertIso8601ToTimestamppb(mp.CreatedAt),
		LastUpdatedAt: convertIso8601ToTimestamppb(mp.LastUpdatedAt),
		EffectiveAt:   convertIso8601ToTimestamppb(mp.EffectiveAt),
		Name:          mp.Name,
		Value:         mp.Value,
		Flags:         mp.Flags,
		Uuid:          mp.Uuid,
	}

	return result
}
