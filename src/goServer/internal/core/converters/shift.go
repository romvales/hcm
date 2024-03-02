package converters

import "goServer/internal/core/pb"

func ConvertMapToShiftProto(mp *Shift) *pb.Shift {
	result := &pb.Shift{
		Id:             mp.Id,
		CreatedById:    mp.CreatedById,
		UpdatedById:    mp.UpdatedById,
		OrganizationId: mp.OrganizationId,
		CreatedAt:      convertIso8601ToTimestamppb(mp.CreatedAt),
		LastUpdatedAt:  convertIso8601ToTimestamppb(mp.LastUpdatedAt),
		Name:           mp.Name,
		Day:            pb.ShiftDay(mp.Day),
		ClockIn:        convertIso8601ToTimestamppb(mp.ClockIn),
		ClockOut:       convertIso8601ToTimestamppb(mp.ClockOut),
	}

	return result
}

func ConvertMapToOverrideShiftProto(mp *OverrideShift) *pb.OverrideShift {
	result := &pb.OverrideShift{
		Id:               mp.Id,
		CreatedById:      mp.CreatedById,
		UpdatedById:      mp.UpdatedById,
		WorkerId:         mp.WorkerId,
		CreatedAt:        convertIso8601ToTimestamppb(mp.CreatedAt),
		LastUpdatedAt:    convertIso8601ToTimestamppb(mp.LastUpdatedAt),
		VerifiedAt:       convertIso8601ToTimestamppb(mp.VerifiedAt),
		CompletedAt:      convertIso8601ToTimestamppb(mp.CompletedAt),
		StartsOn:         convertIso8601ToTimestamppb(mp.StartsOn),
		EndsOn:           convertIso8601ToTimestamppb(mp.EndsOn),
		Name:             mp.Name,
		Day:              pb.ShiftDay(mp.Day),
		OverrideClockIn:  convertIso8601ToTimestamppb(mp.OverrideClockIn),
		OverrideClockOut: convertIso8601ToTimestamppb(mp.OverrideClockOut),
		GroupId:          mp.GroupId,
		Flags:            mp.Flags,
		Uuid:             mp.Uuid,
	}

	return result
}
