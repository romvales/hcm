package converters

import "goServer/internal/core/pb"

func ConvertMapToRoleProto(mp *Role) *pb.Role {
	result := &pb.Role{
		Id:            mp.Id,
		CreatedById:   mp.CreatedById,
		UpdatedById:   mp.UpdatedById,
		CreatedAt:     convertIso8601ToTimestamppb(mp.CreatedAt),
		LastUpdatedAt: convertIso8601ToTimestamppb(mp.LastUpdatedAt),
		Name:          mp.Name,
		Flags:         mp.Flags,
		Uuid:          mp.Uuid,
	}

	return result
}

func ConvertMapToRoleShiftRelationProto(mp *RoleShiftRelation) *pb.RoleShiftRelation {
	result := &pb.RoleShiftRelation{
		Id:              mp.Id,
		RoleId:          mp.RoleId,
		StandardShiftId: mp.StandardShiftId,
	}
	return result
}
