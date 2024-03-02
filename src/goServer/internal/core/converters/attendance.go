package converters

import "goServer/internal/core/pb"

func ConvertMapToAttendanceProto(mp *Attendance) *pb.Attendance {
	result := &pb.Attendance{
		Id:            mp.Id,
		WorkerId:      mp.WorkerId,
		ShiftId:       mp.ShiftId,
		OshiftId:      mp.OshiftId,
		CreatedById:   mp.CreatedById,
		UpdatedById:   mp.UpdatedById,
		CreatedAt:     convertIso8601ToTimestamppb(mp.CreatedAt),
		LastUpdatedAt: convertIso8601ToTimestamppb(mp.LastUpdatedAt),
		ClockIn:       convertIso8601ToTimestamppb(mp.ClockIn),
		ClockOut:      convertIso8601ToTimestamppb(mp.ClockOut),
		Computed:      mp.Computed,
		UnderTime:     mp.UnderTime,
		OverTime:      mp.OverTime,
		LateTime:      mp.LateTime,
		BreakTime:     mp.BreakTime,
		Flags:         mp.Flags,
		Uuid:          mp.Uuid,
	}

	return result
}
