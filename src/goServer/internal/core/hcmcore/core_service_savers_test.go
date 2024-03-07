package hcmcore_test

import (
	"context"
	"goServer/internal/core/hcmcore"
	"goServer/internal/core/pb"
	"strconv"
	"sync"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestSaveWorker(t *testing.T) {
	assert := assert.New(t)
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should create a new worker for each faked data generated", func(t *testing.T) {
		for i := 0; i < TestSaveWorkerParams_N_WORKERS; i++ {
			worker := createMockWorker(t, false)

			res, err := coreService.SaveWorker(context.Background(), &pb.CoreServiceRequest{
				UsedClient: pb.CoreServiceRequest_C_SUPABASE,
				SetterRequest: &pb.SetterRequest{
					WorkerTarget: worker,
				},
			})

			assertCheckForMissingResponse(t, err, res)

			assert.NotEmpty(
				res.SetterResponse,
				"did not return proper response",
			)

			saveWorkerResult := res.GetSetterResponse()
			updatedWorker := saveWorkerResult.GetUpdatedWorkerTarget()

			assert.NotEmpty(
				updatedWorker,
				"did not return the updated target",
			)

			assert.NotEmpty(
				updatedWorker.Id,
				"did not returned an id",
			)

			assert.Equal(
				worker.GetFirstName(),
				updatedWorker.GetFirstName(),
				"did not match the expected name",
			)

			cleanCreatedMockDataInTableNameById(t, "workers", updatedWorker.Id)
		}

	})

}

func TestSaveOrganization(t *testing.T) {
	assert := assert.New(t)
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should create an organization using faked data", func(t *testing.T) {
		faker := DefaultFaker

		for i := 0; i < TestSaveOrganization_N_ORGANIZATION; i++ {
			organization := &pb.Organization{
				Name: faker.Company().Name(),
				Industry: pb.Organization_Industry(faker.RandomIntElement([]int{
					int(pb.Organization_AGRICULTURE),
					int(pb.Organization_CHEMICAL),
					int(pb.Organization_COMMERCE),
					int(pb.Organization_CONSTRUCTION),
					int(pb.Organization_EDUCATION),
					int(pb.Organization_FINANCIAL),
					int(pb.Organization_FORESTRY),
					int(pb.Organization_HEALTH),
				})),
				Flags: uint32(pb.Organization_UNKNOWN),
			}

			res, err := coreService.SaveOrganization(context.Background(), &pb.CoreServiceRequest{
				UsedClient: pb.CoreServiceRequest_C_SUPABASE,
				SetterRequest: &pb.SetterRequest{
					OrganizationTarget: organization,
				},
			})

			assertCheckForMissingResponse(t, err, res)

			assert.NotEmpty(
				res.SetterResponse,
				"did not return a proper response",
			)

			updatedOrganization := res.SetterResponse.UpdatedOrganizationTarget

			assert.NotEmpty(
				updatedOrganization,
				"did not returned the updated version of the organization",
			)

			assert.Equal(
				organization.GetName(),
				updatedOrganization.GetName(),
				"did not match the expected organization name",
			)

			// cleanup
			cleanCreatedMockDataInTableNameById(t, "organizations", updatedOrganization.Id)
		}

	})

}

func TestSaveRole(t *testing.T) {
	assert := assert.New(t)
	coreService := hcmcore.NewCoreServiceServer()
	client := coreService.GetSupabaseCommunityClient()

	t.Run("should create a new role for an organization", func(t *testing.T) {
		organization := createMockOrganization(t, true)
		mockRoles := []string{
			"Chief Executive Officer",
			"Chief Operations Manager",
			"General Manager",
			"Software Engineer",
			"Project Manager",
			"Accountant",
			"Financial Manager",
			"English Teacher",
		}

		for _, roleName := range mockRoles {

			role := createMockRole(t, roleName, organization.Id, false)

			res, err := coreService.SaveRole(context.Background(), &pb.CoreServiceRequest{
				UsedClient: pb.CoreServiceRequest_C_SUPABASE,
				SetterRequest: &pb.SetterRequest{
					RoleTarget: role,
				},
			})

			assertCheckForMissingResponse(t, err, res)

			updatedRole := res.SetterResponse.GetUpdatedRoleTarget()

			assert.NotEmpty(
				updatedRole,
				"did not return the updated role",
			)

			// cleanup
			_, err = client.From("roles").Delete("", "").Eq("id", strconv.FormatInt(updatedRole.Id, 10)).ExecuteTo(nil)
			if err != nil {
				t.Log(err)
			}
		}

		// cleanup
		_, err := client.From("organizations").Delete("", "").Eq("id", strconv.FormatInt(organization.GetId(), 10)).ExecuteTo(nil)
		if err != nil {
			t.Log(err)
		}

	})

}

func TestSaveTeam(t *testing.T) {
	assert := assert.New(t)
	coreService := hcmcore.NewCoreServiceServer()
	client := coreService.GetSupabaseCommunityClient()

	t.Run("should create a new team for an organization", func(t *testing.T) {
		organization := createMockOrganization(t, true)

		mockTeams := []string{
			"English Department",
			"Engineering Department (IT)",
			"Engineering Department (Civil)",
			"Medical Department",
			"Mathematics Department",
		}

		for _, teamName := range mockTeams {
			team := createMockTeam(t, teamName, organization.GetId(), false)

			res, err := coreService.SaveTeam(context.Background(), &pb.CoreServiceRequest{
				UsedClient: pb.CoreServiceRequest_C_SUPABASE,
				SetterRequest: &pb.SetterRequest{
					TeamTarget: team,
				},
			})

			if err != nil {
				t.Log(err)
			}

			assert.NotEmpty(
				res,
				"did not return a proper response",
			)

			updatedTeam := res.SetterResponse.GetUpdatedTeamTarget()

			assert.NotEmpty(
				updatedTeam,
				"did not return the updated data of the mock team",
			)

			// cleanup
			_, err = client.From("teams").Delete("", "").Eq("id", strconv.FormatInt(updatedTeam.Id, 10)).ExecuteTo(nil)
			if err != nil {
				t.Log(err)
			}
		}

		// cleanup
		_, err := client.From("organizations").Delete("", "").Eq("id", strconv.FormatInt(organization.Id, 10)).ExecuteTo(nil)
		if err != nil {
			t.Log(err)
		}
	})
}

func TestSaveWorkerIdentityCard(t *testing.T) {
	assert := assert.New(t)
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should create a new identity card for a mock worker", func(t *testing.T) {
		faker := DefaultFaker
		worker := createMockWorker(t, true)

		identificationCards := []*pb.WorkerIdentityCard{
			{
				Name:          "SSS Identification Card",
				FrontImageUrl: faker.Internet().URL(),
				BackImageUrl:  faker.Internet().URL(),
				ExtractedInfo: &structpb.Struct{},
			},
			{
				Name:          "PhilHealth Card",
				FrontImageUrl: faker.Internet().URL(),
				BackImageUrl:  faker.Internet().URL(),
				ExtractedInfo: &structpb.Struct{},
			},
			{
				Name:          "National ID",
				FrontImageUrl: faker.Internet().URL(),
				BackImageUrl:  faker.Internet().URL(),
				ExtractedInfo: &structpb.Struct{},
			},
		}

		for _, card := range identificationCards {
			card.WorkerId = worker.Id

			res, err := coreService.SaveWorkerIdentityCard(context.Background(), &pb.CoreServiceRequest{
				UsedClient: pb.CoreServiceRequest_C_SUPABASE,
				SetterRequest: &pb.SetterRequest{
					IdentityCardTarget: card,
				},
			})

			assert.NotEmpty(
				res,
				"did not returned a proper response",
			)

			assert.NoError(
				err,
				"expected for SaveWorkerIdentityCard to work properly",
			)

			updatedCard := res.SetterResponse.UpdatedIdentityCardTarget

			assert.NotEmpty(
				updatedCard,
				"did not returned the updated identity card",
			)

			cleanCreatedMockDataInTableNameById(t, "workerIdentityCards", updatedCard.Id)
		}

		cleanCreatedMockDataInTableNameById(t, "workers", worker.Id)
	})

}

func TestSaveMember(t *testing.T) {
	assert := assert.New(t)
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should create new member in a mock organization", func(t *testing.T) {
		worker := createMockWorker(t, true)
		organization := createMockOrganization(t, true)

		member := &pb.Member{
			OrganizationId: organization.Id,
			WorkerId:       worker.Id,
			Flags:          uint32(pb.Member_UNKNOWN),
		}

		res, err := coreService.SaveMember(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			SetterRequest: &pb.SetterRequest{
				MemberTarget: member,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		saveRes := res.SetterResponse

		assert.NotEmpty(
			saveRes,
			"did not return a proper response",
		)

		updatedMember := saveRes.UpdatedMemberTarget

		assert.Equal(
			member.OrganizationId,
			updatedMember.OrganizationId,
			"updatedMember did not match the specified organization id",
		)

		cleanCreatedMockDataInTableNameById(t, "workers", worker.Id)
		cleanCreatedMockDataInTableNameById(t, "organizations", organization.Id)
	})

}

func TestSaveCompensation(t *testing.T) {
	assert := assert.New(t)
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should create a new mock compensation", func(t *testing.T) {
		worker := createMockWorker(t, true)
		organization := createMockOrganization(t, true)

		compensation := &pb.Compensation{
			CreatedById:    &worker.Id,
			WorkerId:       worker.Id,
			OrganizationId: organization.Id,
		}

		res, err := coreService.SaveCompensation(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			SetterRequest: &pb.SetterRequest{
				CompensationTarget: compensation,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		assert.NotEmpty(
			res.SetterResponse,
			"did not properly returned a response",
		)

		compensation = res.SetterResponse.GetUpdatedCompensationTarget()

		assert.NotNil(compensation, "no updated compensation was returned")

		cleanCreatedMockDataInTableNameById(t, "compensations", compensation.Id)
		cleanCreatedMockDataInTableNameById(t, "workers", worker.Id)
		cleanCreatedMockDataInTableNameById(t, "organizations", organization.Id)
	})

}

func TestSaveAddition(t *testing.T) {
	assert := assert.New(t)
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should create series of addition data that can be used by a compensation", func(t *testing.T) {
		worker := createMockWorker(t, true)

		additions := []*pb.Addition{
			{
				CreatedById: &worker.Id,
				Name:        "Overtime Bonus",
				Value:       300.00,
				Flags:       uint32(pb.Addition_UNKNOWN),
			},
			{
				CreatedById: &worker.Id,
				Name:        "IT Department Onboarding Bonus",
				Value:       250.00,
				Flags:       uint32(pb.Addition_UNKNOWN),
			},
		}

		for _, add := range additions {

			res, err := coreService.SaveAddition(context.Background(), &pb.CoreServiceRequest{
				SetterRequest: &pb.SetterRequest{
					AdditionTarget: add,
				},
			})

			assertCheckForMissingResponse(t, err, res)

			updatedAdd := res.SetterResponse.GetUpdatedAdditionTarget()

			assert.NotEmpty(
				updatedAdd,
				"did not returned the updated addition target",
			)

			cleanCreatedMockDataInTableNameById(t, "additions", updatedAdd.Id)
		}

		cleanCreatedMockDataInTableNameById(t, "workers", worker.Id)
	})

}

func TestSaveDeduction(t *testing.T) {
	assert := assert.New(t)
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should create series of deduction data that can be used by a compensation", func(t *testing.T) {
		worker := createMockWorker(t, true)

		deductions := []*pb.Deduction{
			{
				CreatedById: &worker.Id,
				Name:        "SSS Monthly Contribution",
				Value:       300.00,
				Flags:       uint32(pb.Addition_UNKNOWN),
			},
			{
				CreatedById: &worker.Id,
				Name:        "PhilHealth Contribution",
				Value:       250.00,
				Flags:       uint32(pb.Addition_UNKNOWN),
			},
		}

		for _, deduct := range deductions {

			res, err := coreService.SaveDeduction(context.Background(), &pb.CoreServiceRequest{
				SetterRequest: &pb.SetterRequest{
					DeductionTarget: deduct,
				},
			})

			assertCheckForMissingResponse(t, err, res)

			updatedAdd := res.SetterResponse.GetUpdatedDeductionTarget()

			assert.NotEmpty(
				updatedAdd,
				"did not returned the updated deduction target",
			)

			cleanCreatedMockDataInTableNameById(t, "deductions", updatedAdd.Id)
		}

		cleanCreatedMockDataInTableNameById(t, "workers", worker.Id)
	})

}

func TestSavePayroll(t *testing.T) {
	assert := assert.New(t)
	coreService := hcmcore.NewCoreServiceServer()
	organization := createMockOrganization(t, true)

	t.Run("should create a new empty payroll for an organization", func(t *testing.T) {
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

		assert.NotNil(
			payroll,
			"did not returned the updated payroll target",
		)

		cleanCreatedMockDataInTableNameById(t, "payrolls", payroll.Id)
	})

	cleanCreatedMockDataInTableNameById(t, "organizations", organization.Id)
}

func TestSaveShift(t *testing.T) {
	assert := assert.New(t)
	coreService := hcmcore.NewCoreServiceServer()
	organization := createMockOrganization(t, true)

	var testRun sync.WaitGroup

	testRun.Add(1)
	t.Run("should create a new shift for an organization", func(t *testing.T) {
		shift := &pb.Shift{
			OrganizationId: organization.Id,
			Name:           "Software Engineer Monday Shift",
			Day:            pb.ShiftDay_MON,
			ClockIn:        timestamppb.Now(),
			ClockOut:       timestamppb.Now(),
		}

		res, err := coreService.SaveShift(context.Background(), &pb.CoreServiceRequest{
			SetterRequest: &pb.SetterRequest{
				ShiftTarget: shift,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		updatedShift := res.SetterResponse.UpdatedShiftTarget

		assert.NotNil(
			updatedShift,
			"did not returned the updated shift in the response",
		)

		testRun.Done()
	})

	testRun.Add(1)
	t.Run("should create a new override shift for a worker in an organization", func(t *testing.T) {
		overrideShift := &pb.OverrideShift{
			OrganizationId:   organization.Id,
			Name:             "Penalty Override (Tue)",
			Day:              pb.ShiftDay_TUE,
			OverrideClockIn:  timestamppb.Now(),
			OverrideClockOut: timestamppb.Now(),
			GroupId:          uuid.NewString(),
		}

		res, err := coreService.SaveShift(context.Background(), &pb.CoreServiceRequest{
			SetterRequest: &pb.SetterRequest{
				TargetShiftType:     pb.SetterRequest_T_OVERRIDESHIFT.Enum(),
				OverrideShiftTarget: overrideShift,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		updatedShift := res.SetterResponse.UpdatedOverrideShiftTarget

		assert.NotNil(
			updatedShift,
			"did not returned the updated shift in the response",
		)

		testRun.Done()
	})

	testRun.Wait()

	cleanCreatedMockDataInTableNameById(t, "organizations", organization.Id)
}

func TestSaveAttendance(t *testing.T) {
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should persist an attendance for a mock user in the database", func(t *testing.T) {
		organization := createMockOrganization(t, true)
		worker := createMockWorker(t, true)

		res, err := coreService.SaveShift(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			SetterRequest: &pb.SetterRequest{
				ShiftTarget: &pb.Shift{
					Name:           "Accounting Shift (Mon)",
					OrganizationId: organization.Id,
					Day:            pb.ShiftDay_MON,
					ClockIn:        timestamppb.Now(),
					ClockOut:       timestamppb.Now(),
				},
			},
		})

		assertCheckForMissingResponse(t, err, res)

		shift := res.SetterResponse.GetUpdatedShiftTarget()

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

		cleanCreatedMockDataInTableNameById(t, "workers", worker.Id)
		cleanCreatedMockDataInTableNameById(t, "organizations", organization.Id)
	})

}
