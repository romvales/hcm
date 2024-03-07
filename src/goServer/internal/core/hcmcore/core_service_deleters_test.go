package hcmcore_test

import (
	"context"
	"goServer/internal/core/hcmcore"
	"goServer/internal/core/pb"
	"sync"
	"testing"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestDeleteWorkerById(t *testing.T) {
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should delete the mock user persisted in the database", func(t *testing.T) {
		worker := createMockWorker(t, true)

		res, err := coreService.DeleteWorkerById(context.Background(), &pb.CoreServiceRequest{
			SetterRequest: &pb.SetterRequest{
				TargetId: &worker.Id,
			},
		})

		assertCheckForMissingResponse(t, err, res)
	})
}

func TestDeleteOrganizationById(t *testing.T) {
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should delete the mock organization in the database", func(t *testing.T) {
		organization := createMockOrganization(t, true)

		res, err := coreService.DeleteOrganizationById(context.Background(), &pb.CoreServiceRequest{
			SetterRequest: &pb.SetterRequest{
				TargetId: &organization.Id,
			},
		})

		assertCheckForMissingResponse(t, err, res)
	})

}

func TestDeleteRoleById(t *testing.T) {
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should delete the mock organization in the database", func(t *testing.T) {
		organization := createMockOrganization(t, true)
		mockRole := createMockRole(t, "Financial Manager", organization.Id, true)

		res, err := coreService.DeleteRoleById(context.Background(), &pb.CoreServiceRequest{
			SetterRequest: &pb.SetterRequest{
				TargetId: &mockRole.Id,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		cleanCreatedMockDataInTableNameById(t, "organizations", organization.Id)
	})
}

func TestDeleteTeamById(t *testing.T) {
	coreService := hcmcore.NewCoreServiceServer()
	organization := createMockOrganization(t, true)

	t.Run("should delete the created mock team in an organization", func(t *testing.T) {
		team := createMockTeam(t, "Accounting", organization.Id, true)

		res, err := coreService.DeleteTeamById(context.Background(), &pb.CoreServiceRequest{
			SetterRequest: &pb.SetterRequest{
				TargetId: &team.Id,
			},
		})

		assertCheckForMissingResponse(t, err, res)
	})

	cleanCreatedMockDataInTableNameById(t, "organizations", organization.Id)
}

func TestDeleteWorkerIdentityCardById(t *testing.T) {
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should delete the created mock worker identity card from the database.", func(t *testing.T) {
		worker := createMockWorker(t, true)
		faker := DefaultFaker

		idCard := &pb.WorkerIdentityCard{
			WorkerId:      worker.Id,
			Name:          "National ID",
			FrontImageUrl: faker.Internet().URL(),
			BackImageUrl:  faker.Internet().URL(),
		}

		res, err := coreService.SaveWorkerIdentityCard(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			SetterRequest: &pb.SetterRequest{
				IdentityCardTarget: idCard,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		idCard = res.SetterResponse.GetUpdatedIdentityCardTarget()

		res, err = coreService.DeleteWorkerIdentityCardById(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			SetterRequest: &pb.SetterRequest{
				TargetId: &idCard.Id,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		cleanCreatedMockDataInTableNameById(t, "workers", worker.Id)
	})

}

func TestDeleteCompensationById(t *testing.T) {
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should delete the created mock compensation from the database.", func(t *testing.T) {
		organization := createMockOrganization(t, true)
		worker := createMockWorker(t, true)

		compensation := &pb.Compensation{
			OrganizationId: organization.Id,
			WorkerId:       worker.Id,
		}

		res, err := coreService.SaveCompensation(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			SetterRequest: &pb.SetterRequest{
				CompensationTarget: compensation,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		compensation = res.SetterResponse.GetUpdatedCompensationTarget()

		res, err = coreService.DeleteCompensationById(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			SetterRequest: &pb.SetterRequest{
				TargetId: &compensation.Id,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		cleanCreatedMockDataInTableNameById(t, "workers", worker.Id)
		cleanCreatedMockDataInTableNameById(t, "organizations", organization.Id)
	})

}

func TestDeleteMemberById(t *testing.T) {
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should delete the created mock organization member from the database.", func(t *testing.T) {
		organization := createMockOrganization(t, true)
		worker := createMockWorker(t, true)

		member := &pb.Member{
			OrganizationId: organization.Id,
			WorkerId:       worker.Id,
		}

		res, err := coreService.SaveMember(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			SetterRequest: &pb.SetterRequest{
				MemberTarget: member,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		cleanCreatedMockDataInTableNameById(t, "workers", worker.Id)
		cleanCreatedMockDataInTableNameById(t, "organizations", organization.Id)
	})

}

func TestDeletePayrollById(t *testing.T) {
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should delete the created mock payroll from the database.", func(t *testing.T) {
		organization := createMockOrganization(t, true)

		payroll := &pb.Payroll{
			OrganizationId: organization.Id,
		}

		res, err := coreService.SavePayroll(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			SetterRequest: &pb.SetterRequest{
				PayrollTarget: payroll,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		payroll = res.SetterResponse.GetUpdatedPayrollTarget()

		res, err = coreService.DeletePayrollById(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			SetterRequest: &pb.SetterRequest{
				TargetId: &payroll.Id,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		cleanCreatedMockDataInTableNameById(t, "organizations", organization.Id)
	})

}

func TestDeleteAdditionById(t *testing.T) {
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should delete the created mock addition from the database.", func(t *testing.T) {
		worker := createMockWorker(t, true)

		addition := &pb.Addition{
			WorkerId: &worker.Id,
			Name:     "Overtime Bonus",
			Value:    2_500.00,
		}

		res, err := coreService.SaveAddition(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			SetterRequest: &pb.SetterRequest{
				AdditionTarget: addition,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		addition = res.SetterResponse.GetUpdatedAdditionTarget()

		res, err = coreService.DeleteAdditionById(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			SetterRequest: &pb.SetterRequest{
				TargetId: &addition.Id,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		cleanCreatedMockDataInTableNameById(t, "workers", worker.Id)
	})

}

func TestDeleteDeductionById(t *testing.T) {
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should delete the created mock deduction from the database.", func(t *testing.T) {
		worker := createMockWorker(t, true)

		deduction := &pb.Deduction{
			WorkerId: &worker.Id,
			Name:     "Overtime Bonus",
			Value:    2_500.00,
		}

		res, err := coreService.SaveDeduction(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			SetterRequest: &pb.SetterRequest{
				DeductionTarget: deduction,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		deduction = res.SetterResponse.GetUpdatedDeductionTarget()

		res, err = coreService.DeleteDeductionById(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			SetterRequest: &pb.SetterRequest{
				TargetId: &deduction.Id,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		cleanCreatedMockDataInTableNameById(t, "workers", worker.Id)
	})

}

func TestDeleteShiftById(t *testing.T) {
	coreService := hcmcore.NewCoreServiceServer()
	organization := createMockOrganization(t, true)

	var testRun sync.WaitGroup

	testRun.Add(1)
	t.Run("should delete the created mock shift from the database.", func(t *testing.T) {

		shift := &pb.Shift{
			OrganizationId: organization.Id,
			Name:           "Accountant Friday Shift",
			Day:            pb.ShiftDay_FRI,
			ClockIn:        timestamppb.Now(),
			ClockOut:       timestamppb.Now(),
		}

		res, err := coreService.SaveShift(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			SetterRequest: &pb.SetterRequest{
				ShiftTarget: shift,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		shift = res.SetterResponse.GetUpdatedShiftTarget()

		res, err = coreService.DeleteShiftById(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			SetterRequest: &pb.SetterRequest{
				TargetId: &shift.Id,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		testRun.Done()
	})

	testRun.Add(1)
	t.Run("should delete the created mock override shift from the database.", func(t *testing.T) {

		overrideShift := &pb.OverrideShift{
			OrganizationId:   organization.Id,
			GroupId:          uuid.NewString(),
			Name:             "Rom Vales' Penalty Override Monday Shift",
			Day:              pb.ShiftDay_MON,
			OverrideClockIn:  timestamppb.Now(),
			OverrideClockOut: timestamppb.Now(),
		}

		res, err := coreService.SaveShift(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			SetterRequest: &pb.SetterRequest{
				TargetShiftType:     pb.SetterRequest_T_OVERRIDESHIFT.Enum(),
				OverrideShiftTarget: overrideShift,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		overrideShift = res.SetterResponse.GetUpdatedOverrideShiftTarget()

		res, err = coreService.DeleteShiftById(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			SetterRequest: &pb.SetterRequest{
				TargetId: &overrideShift.Id,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		testRun.Done()
	})

	testRun.Wait()

	cleanCreatedMockDataInTableNameById(t, "organizations", organization.Id)
}

func TestDeleteAttendanceById(t *testing.T) {
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should delete the created mock attendance from the database.", func(t *testing.T) {
		organization := createMockOrganization(t, true)
		worker := createMockWorker(t, true)

		shift := &pb.Shift{
			OrganizationId: organization.Id,
			Name:           "Accountant Friday Shift",
			Day:            pb.ShiftDay_FRI,
			ClockIn:        timestamppb.Now(),
			ClockOut:       timestamppb.Now(),
		}

		res, err := coreService.SaveShift(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			SetterRequest: &pb.SetterRequest{
				ShiftTarget: shift,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		shift = res.SetterResponse.GetUpdatedShiftTarget()

		attendance := &pb.Attendance{
			WorkerId: worker.Id,
			ShiftId:  &shift.Id,
			ClockIn:  timestamppb.Now(),
		}

		res, err = coreService.SaveAttendance(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			SetterRequest: &pb.SetterRequest{
				AttendanceTarget: attendance,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		attendance = res.SetterResponse.GetUpdatedAttendanceTarget()

		res, err = coreService.DeleteAttendanceById(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			SetterRequest: &pb.SetterRequest{
				TargetId: &attendance.Id,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		cleanCreatedMockDataInTableNameById(t, "workers", worker.Id)
		cleanCreatedMockDataInTableNameById(t, "standardShifts", shift.Id)
		cleanCreatedMockDataInTableNameById(t, "organizations", organization.Id)
	})

}
