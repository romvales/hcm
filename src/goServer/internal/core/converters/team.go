package converters

import "goServer/internal/core/pb"

func ConvertMapToTeamProto(mp *Team) *pb.Team {
	result := &pb.Team{
		Id:             mp.Id,
		CreatedById:    mp.CreatedById,
		UpdatedById:    mp.UpdatedById,
		OrganizationId: mp.OrganizationId,
		CreatedAt:      convertIso8601ToTimestamppb(mp.CreatedAt),
		LastUpdatedAt:  convertIso8601ToTimestamppb(mp.LastUpdatedAt),
		Name:           mp.Name,
		Flags:          mp.Flags,
		Uuid:           mp.Uuid,
	}

	return result
}

func ConvertMapToTeamMemberRelationProto(mp *TeamMemberRelation) *pb.TeamMemberRelation {
	result := &pb.TeamMemberRelation{
		Id:       mp.Id,
		TeamId:   mp.TeamId,
		WorkerId: mp.WorkerId,
	}

	return result
}
