package converters

import "goServer/internal/core/pb"

func ConvertMapToOrganizationProto(mp *Organization) *pb.Organization {
	result := &pb.Organization{
		Id:               mp.Id,
		CreatedById:      mp.CreatedById,
		UpdatedById:      mp.UpdatedById,
		CreatedAt:        convertIso8601ToTimestamppb(mp.CreatedAt),
		LastUpdatedAt:    convertIso8601ToTimestamppb(mp.LastUpdatedAt),
		Industry:         pb.Organization_Industry(mp.Industry),
		OverrideIndustry: mp.OverrideIndustry,
		Name:             mp.Name,
		Flags:            mp.Flags,
		Uuid:             mp.Uuid,
	}

	return result
}
