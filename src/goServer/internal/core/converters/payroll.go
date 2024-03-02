package converters

import (
	"goServer/internal/core/pb"
)

func ConvertMapToPayrollProto(mp *Payroll) *pb.Payroll {
	result := &pb.Payroll{
		Id:             mp.Id,
		CreatedById:    mp.CreatedById,
		UpdatedById:    mp.UpdatedById,
		VerifiedById:   mp.VerifiedById,
		OrganizationId: mp.OrganizationId,
		CreatedAt:      convertIso8601ToTimestamppb(mp.CreatedAt),
		LastUpdatedAt:  convertIso8601ToTimestamppb(mp.LastUpdatedAt),
		Total:          mp.Total,
		Flags:          mp.Flags,
		Uuid:           mp.Uuid,
	}

	if mp.VerifiedAt != nil {
		verifiedAt := convertIso8601ToTimestamppb(mp.VerifiedAt)
		result.VerifiedAt = verifiedAt
	}

	return result
}

func ConvertMapToPayrollCompensationRelationProto(mp PayrollCompensationRelation) *pb.PayrollCompensationRelation {
	result := &pb.PayrollCompensationRelation{
		Id:             mp.Id,
		PayrollId:      mp.PayrollId,
		CompensationId: mp.CompensationId,
	}

	return result
}
