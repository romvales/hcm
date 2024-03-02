package converters

import "goServer/internal/core/pb"

func ConvertMapToMemberProto(mp *Member) *pb.Member {
	result := &pb.Member{
		Id:                    mp.Id,
		OrganizationId:        mp.OrganizationId,
		WorkerId:              mp.WorkerId,
		HiredById:             mp.HiredById,
		CreatedAt:             convertIso8601ToTimestamppb(mp.CreatedAt),
		LastUpdatedAt:         convertIso8601ToTimestamppb(mp.LastUpdatedAt),
		HiredAt:               convertIso8601ToTimestamppb(mp.HiredAt),
		SuspendedAt:           convertIso8601ToTimestamppb(mp.SuspendedAt),
		LeaveAt:               convertIso8601ToTimestamppb(mp.LeaveAt),
		TerminatedAt:          convertIso8601ToTimestamppb(mp.TerminatedAt),
		ReturnedAt:            convertIso8601ToTimestamppb(mp.ReturnedAt),
		ScheduledSuspensionAt: convertIso8601ToTimestamppb(mp.ScheduledSuspensionAt),
		Flags:                 mp.Flags,
		Uuid:                  mp.Uuid,
	}

	return result
}
