package hcmcore_test

import (
	"context"
	"goServer/internal/core/hcmcore"
	"goServer/internal/core/pb"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGetWorkerById(t *testing.T) {
	assert := assert.New(t)
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should be able to get the mock worker from the database", func(t *testing.T) {
		worker := createMockWorker(t, true)

		res, err := coreService.GetWorkerById(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			GetterRequest: &pb.GetterRequest{
				TargetId: &worker.Id,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		workerResult := res.GetterResponse.GetWorkerResult()

		assert.Equal(res.Code, pb.CoreServiceResponse_C_NOERROR, "expected for the function to work properly")
		assert.NotEmpty(workerResult, "did not return a result in the response")
		assert.Equal(worker.Id, workerResult.Id, "did not match the expected id")
		assert.Equal(worker.FirstName, workerResult.FirstName, "did not match the expected first name")
		assert.Equal(worker.LastName, workerResult.LastName, "did not match the expected last name")

		cleanCreatedMockDataInTableNameById(t, "workers", worker.Id)
	})

}

func TestGetOrganizationById(t *testing.T) {
	assert := assert.New(t)
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should be able to get the mock worker from the database", func(t *testing.T) {
		organization := createMockOrganization(t, true)

		res, err := coreService.GetOrganizationById(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			GetterRequest: &pb.GetterRequest{
				TargetId: &organization.Id,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		organizationResult := res.GetterResponse.GetOrganizationResult()

		assert.Equal(organization.Id, organizationResult.Id, "did not match the expected id")
		assert.Equal(organization.Name, organizationResult.Name, "did not match the expected organization name")

		cleanCreatedMockDataInTableNameById(t, "organizations", organization.Id)
	})

}

func TestGetOrganizationsByCreatorId(t *testing.T) {
	assert := assert.New(t)
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should get all mock organizations created by a mock worker", func(t *testing.T) {
		worker := createMockWorker(t, true)

		organizationsByWorker := []*pb.Organization{}

		for i := 0; i < TestSaveOrganization_N_ORGANIZATION; i++ {
			org := createMockOrganization(t, false)
			org.CreatedById = &worker.Id
			org = saveMockOrganization(t, org)

			organizationsByWorker = append(organizationsByWorker, org)

			defer func() {
				cleanCreatedMockDataInTableNameById(t, "organizations", org.Id)
			}()
		}

		res, err := coreService.GetOrganizationsByCreatorId(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			GetterRequest: &pb.GetterRequest{
				TargetId: &worker.Id,
			},
		})

		assertCheckForMissingResponse(t, err, res)
		organizationsResult := res.GetterResponse.GetOrganizationsResult()

		assert.Len(organizationsResult, len(organizationsByWorker), "did not match the expected number of organizations the mock worker created")

		cleanCreatedMockDataInTableNameById(t, "workers", worker.Id)
	})

}

func TestGetRoleById(t *testing.T) {
	assert := assert.New(t)
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should get mock roles from the database", func(t *testing.T) {
		organization := createMockOrganization(t, true)
		roles := []string{"Civil Engineer", "Software Engineer", "Financial Assistant"}
		savedMockRoles := []*pb.Role{}

		for _, roleName := range roles {
			mockRole := createMockRole(t, roleName, organization.Id, true)
			savedMockRoles = append(savedMockRoles, mockRole)

			defer func() {
				cleanCreatedMockDataInTableNameById(t, "roles", mockRole.Id)
			}()
		}

		for _, mockRole := range savedMockRoles {

			res, err := coreService.GetRoleById(context.Background(), &pb.CoreServiceRequest{
				UsedClient: pb.CoreServiceRequest_C_SUPABASE,
				GetterRequest: &pb.GetterRequest{
					TargetId: &mockRole.Id,
				},
			})

			assertCheckForMissingResponse(t, err, res)

			expectedRole := res.GetterResponse.GetRoleResult()

			assert.NotEmpty(expectedRole, "did not returned any resulting role")
			assert.Equal(expectedRole.Name, mockRole.Name, "did not match the expected mock role name")
		}

		cleanCreatedMockDataInTableNameById(t, "organizations", organization.Id)
	})

}

func TestGetTeamById(t *testing.T) {
	assert := assert.New(t)
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should get mock teams from the database", func(t *testing.T) {
		organization := createMockOrganization(t, true)
		teams := []string{"Civil Engineering Department", "Systems IT Department", "Finance Deparment"}
		savedMockTeams := []*pb.Team{}

		for _, teamName := range teams {
			mockTeam := createMockTeam(t, teamName, organization.Id, true)
			savedMockTeams = append(savedMockTeams, mockTeam)

			defer func() {
				cleanCreatedMockDataInTableNameById(t, "teams", mockTeam.Id)
			}()
		}

		for _, mockTeam := range savedMockTeams {

			res, err := coreService.GetTeamById(context.Background(), &pb.CoreServiceRequest{
				UsedClient: pb.CoreServiceRequest_C_SUPABASE,
				GetterRequest: &pb.GetterRequest{
					TargetId: &mockTeam.Id,
				},
			})

			assertCheckForMissingResponse(t, err, res)

			expectedTeam := res.GetterResponse.GetTeamResult()

			assert.NotEmpty(expectedTeam, "did not returned any resulting team")
			assert.Equal(expectedTeam.Name, mockTeam.Name, "did not match the expected mock team name")
		}

		cleanCreatedMockDataInTableNameById(t, "organizations", organization.Id)
	})

}

func TestGetRolesFromOrganization(t *testing.T) {
	assert := assert.New(t)
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should get mock roles from the database", func(t *testing.T) {
		organization := createMockOrganization(t, true)
		roles := []string{"Civil Engineer", "Software Engineer", "Financial Assistant"}
		savedMockRoles := []*pb.Role{}

		for _, roleName := range roles {
			mockRole := createMockRole(t, roleName, organization.Id, true)
			savedMockRoles = append(savedMockRoles, mockRole)

			defer func() {
				cleanCreatedMockDataInTableNameById(t, "roles", mockRole.Id)
			}()
		}

		res, err := coreService.GetRolesFromOrganization(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			GetterRequest: &pb.GetterRequest{
				TargetId: &organization.Id,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		organizationRoles := res.GetterResponse.GetOrganizationRolesResult()

		assert.Len(organizationRoles, len(savedMockRoles), "did not match the expected number of mock roles created for the mock organization")

		cleanCreatedMockDataInTableNameById(t, "organizations", organization.Id)
	})

}

func TestGetTeamsFromOrganization(t *testing.T) {
	assert := assert.New(t)
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should get mock teams from the database", func(t *testing.T) {
		organization := createMockOrganization(t, true)
		teams := []string{"Civil Engineering Department", "Systems IT Department", "Finance Deparment"}
		savedMockTeams := []*pb.Team{}

		for _, teamName := range teams {
			mockTeam := createMockTeam(t, teamName, organization.Id, true)
			savedMockTeams = append(savedMockTeams, mockTeam)

			defer func() {
				cleanCreatedMockDataInTableNameById(t, "teams", mockTeam.Id)
			}()
		}

		res, err := coreService.GetTeamsFromOrganization(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			GetterRequest: &pb.GetterRequest{
				TargetId: &organization.Id,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		organizationTeams := res.GetterResponse.GetOrganizationTeamsResult()

		assert.Len(organizationTeams, len(savedMockTeams), "did not match the expected number of mock teams saved in the organization")

		cleanCreatedMockDataInTableNameById(t, "organizations", organization.Id)
	})

}

func TestGetOrganizationJoinRequests(t *testing.T) {
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should get all mock join requests send by an organization to a worker", func(t *testing.T) {
		organization := createMockOrganization(t, true)
		worker := createMockWorker(t, true)

		res, err := coreService.SendJoinRequest(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			SetterRequest: &pb.SetterRequest{
				RequestSenderType: pb.JoinRequest_T_ORGANIZATION.Enum(),
				RequestSenderId:   &organization.Id,
				TargetId:          &worker.Id,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		cleanCreatedMockDataInTableNameById(t, "organizations", organization.Id)
		cleanCreatedMockDataInTableNameById(t, "workers", worker.Id)
	})

}

func TestGetWorkerJoinRequests(t *testing.T) {
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should get all mock join requests send by a worker to an organization", func(t *testing.T) {
		organization := createMockOrganization(t, true)
		worker := createMockWorker(t, true)

		res, err := coreService.SendJoinRequest(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			SetterRequest: &pb.SetterRequest{
				RequestSenderType: pb.JoinRequest_T_WORKER.Enum(),
				RequestSenderId:   &worker.Id,
				TargetId:          &organization.Id,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		cleanCreatedMockDataInTableNameById(t, "organizations", organization.Id)
		cleanCreatedMockDataInTableNameById(t, "workers", worker.Id)
	})

}

func TestGetJoinRequestById(t *testing.T) {
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should get all mock join requests send by an organization to a worker", func(t *testing.T) {
		organization := createMockOrganization(t, true)
		worker := createMockWorker(t, true)

		res, err := coreService.SendJoinRequest(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			SetterRequest: &pb.SetterRequest{
				RequestSenderType: pb.JoinRequest_T_ORGANIZATION.Enum(),
				RequestSenderId:   &organization.Id,
				TargetId:          &worker.Id,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		request := res.SetterResponse.GetJoinRequestResult()

		res, err = coreService.GetJoinRequestById(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			GetterRequest: &pb.GetterRequest{
				TargetId: &request.Id,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		cleanCreatedMockDataInTableNameById(t, "pendingJoinRequests", request.Id)
		cleanCreatedMockDataInTableNameById(t, "organizations", organization.Id)
		cleanCreatedMockDataInTableNameById(t, "workers", worker.Id)
	})
}

func TestGetCompensationById(t *testing.T) {
	assert := assert.New(t)
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should get the created mock compensation from the database", func(t *testing.T) {
		worker := createMockWorker(t, true)
		organization := createMockOrganization(t, true)

		compensation := &pb.Compensation{
			CreatedById:    &worker.Id,
			WorkerId:       worker.Id,
			OrganizationId: organization.Id,
			Flags:          uint32(pb.Compensation_UNKNOWN),
		}

		res, err := coreService.SaveCompensation(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			SetterRequest: &pb.SetterRequest{
				CompensationTarget: compensation,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		savedCompensation := res.SetterResponse.GetUpdatedCompensationTarget()

		res, err = coreService.GetCompensationById(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			GetterRequest: &pb.GetterRequest{
				TargetId: &savedCompensation.Id,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		result := res.GetterResponse.GetCompensationResult()

		assert.NotNil(result, "did not return the expected compensation")
		assert.Equal(savedCompensation.Id, result.Id, "did not match the expected compensation id")

		cleanCreatedMockDataInTableNameById(t, "compensations", result.Id)
		cleanCreatedMockDataInTableNameById(t, "workers", worker.Id)
		cleanCreatedMockDataInTableNameById(t, "organizations", organization.Id)
	})

}

func TestGetAdditionById(t *testing.T) {
	assert := assert.New(t)
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should get the mock addtion persisted in the database", func(t *testing.T) {
		hr := createMockWorker(t, true)
		worker := createMockWorker(t, true)
		addition := &pb.Addition{
			CreatedById: &hr.Id,
			WorkerId:    &worker.Id,
			Name:        "Overtime Bonus (2%)",
			Value:       15_000 * 0.02,
			Flags:       uint32(pb.Addition_UNKNOWN),
		}

		res, err := coreService.SaveAddition(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			SetterRequest: &pb.SetterRequest{
				AdditionTarget: addition,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		savedAddition := res.SetterResponse.GetUpdatedAdditionTarget()

		res, err = coreService.GetAdditionById(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			GetterRequest: &pb.GetterRequest{
				TargetId: &savedAddition.Id,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		result := res.GetterResponse.GetAdditionResult()

		assert.NotNil(result, "expected to return the mock addition")

		cleanCreatedMockDataInTableNameById(t, "additions", savedAddition.Id)
		cleanCreatedMockDataInTableNameById(t, "workers", hr.Id)
		cleanCreatedMockDataInTableNameById(t, "workers", worker.Id)
	})

}

func TestGetDeductionById(t *testing.T) {
	assert := assert.New(t)
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should get the mock deduction persisted in the database", func(t *testing.T) {
		hr := createMockWorker(t, true)
		worker := createMockWorker(t, true)
		deduction := &pb.Deduction{
			CreatedById: &hr.Id,
			WorkerId:    &worker.Id,
			Name:        "Overtime Bonus (2%)",
			Value:       15_000 * 0.02,
			Flags:       uint32(pb.Addition_UNKNOWN),
		}

		res, err := coreService.SaveDeduction(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			SetterRequest: &pb.SetterRequest{
				DeductionTarget: deduction,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		savedDeduction := res.SetterResponse.GetUpdatedDeductionTarget()

		res, err = coreService.GetAdditionById(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			GetterRequest: &pb.GetterRequest{
				TargetId: &savedDeduction.Id,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		result := res.GetterResponse.GetAdditionResult()

		assert.NotNil(result, "expected to return the mock deduction")

		cleanCreatedMockDataInTableNameById(t, "deductions", savedDeduction.Id)
		cleanCreatedMockDataInTableNameById(t, "workers", hr.Id)
		cleanCreatedMockDataInTableNameById(t, "workers", worker.Id)
	})
}

func TestGetAttendanceById(t *testing.T) {
	assert := assert.New(t)
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should get the mock attendance from the database", func(t *testing.T) {
		organization := createMockOrganization(t, true)
		worker := createMockWorker(t, true)

		shift := &pb.Shift{
			OrganizationId: organization.Id,
			Name:           "Mock Worker Shift",
			Day:            pb.ShiftDay_MON,
			ClockIn:        timestamppb.Now(),
			ClockOut:       timestamppb.Now(),
		}

		res, err := coreService.SaveShift(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			SetterRequest: &pb.SetterRequest{
				TargetShiftType: pb.SetterRequest_T_SHIFT.Enum(),
				ShiftTarget:     shift,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		savedShift := res.SetterResponse.GetUpdatedShiftTarget()

		attendance := &pb.Attendance{
			WorkerId: worker.Id,
			ShiftId:  &savedShift.Id,
			ClockIn:  timestamppb.Now(),
		}

		res, err = coreService.SaveAttendance(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			SetterRequest: &pb.SetterRequest{
				AttendanceTarget: attendance,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		savedAttendance := res.SetterResponse.GetUpdatedAttendanceTarget()

		assert.NotNil(savedAttendance, "expected to return the updated attendance from the database")

		cleanCreatedMockDataInTableNameById(t, "attendances", savedAttendance.Id)
		cleanCreatedMockDataInTableNameById(t, "standardShifts", shift.Id)
		cleanCreatedMockDataInTableNameById(t, "workers", worker.Id)
	})

}

func TestGetPayrollById(t *testing.T) {
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should get the mock payroll persisted in the database", func(t *testing.T) {
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

		savedPayroll := res.SetterResponse.GetUpdatedPayrollTarget()

		res, err = coreService.GetPayrollById(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			GetterRequest: &pb.GetterRequest{
				TargetId: &savedPayroll.Id,
			},
		})

		assertCheckForMissingResponse(t, err, res)

		cleanCreatedMockDataInTableNameById(t, "organizations", organization.Id)
	})

}
