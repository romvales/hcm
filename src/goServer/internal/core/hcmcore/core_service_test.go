package hcmcore_test

import (
	"context"
	"errors"
	"goServer/internal/core/hcmcore"
	"goServer/internal/core/pb"
	"strconv"
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	ErrCoreServiceServerUnimplementedMethod = errors.New("CoreServiceServer: Not implemented")
)

const (
	TestSaveWorkerParams_N_WORKERS      = 10
	TestSaveOrganization_N_ORGANIZATION = 10
)

func TestGetWorkerById(t *testing.T) {
	coreService := hcmcore.NewCoreServiceServer()

	workerId := int64(0)

	_, err := coreService.GetWorkerById(context.TODO(), &pb.CoreServiceRequest{
		GetterRequest: &pb.GetterRequest{
			TargetId: &workerId,
		},
	})

	if err != nil {
		t.Log(err)
	}

	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func TestGetOrganizationById(t *testing.T) {
	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func TestGetOrganizationByCreatorId(t *testing.T) {
	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func TestGetRoleById(t *testing.T) {
	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func TestGetTeamById(t *testing.T) {
	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func TestGetRolesFromOrganization(t *testing.T) {
	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func TestGetTeamsFromOrganization(t *testing.T) {
	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func TestGetOrganizationJoinRequests(t *testing.T) {
	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func TestGetWorkerJoinRequests(t *testing.T) {
	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func TestGetJoinRequestById(t *testing.T) {
	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func TestGetCompensationById(t *testing.T) {
	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func TestGetAdditionById(t *testing.T) {
	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func TestGetDeductionById(t *testing.T) {
	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func TestSaveWorker(t *testing.T) {
	assert := assert.New(t)
	coreService := hcmcore.NewCoreServiceServer()
	client := coreService.GetSupabaseCommunityClient()

	t.Run("create a new worker for each faked data generated", func(t *testing.T) {
		for i := 0; i < TestSaveWorkerParams_N_WORKERS; i++ {
			worker := createMockWorker(t, false)

			res, err := coreService.SaveWorker(context.Background(), &pb.CoreServiceRequest{
				UsedClient: pb.CoreServiceRequest_C_SUPABASE,
				SetterRequest: &pb.SetterRequest{
					WorkerTarget: worker,
				},
			})

			if err != nil {
				t.Error(err)
			}

			assert.Equal(
				pb.CoreServiceResponse_C_NOERROR,
				res.Code,
				"must not return an error",
			)

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

			// cleanup
			_, err = client.From("workers").Delete("", "").Eq("id", strconv.FormatInt(updatedWorker.Id, 10)).ExecuteTo(nil)
			if err != nil {
				t.Log(err)
			}

		}

	})

}

func TestSaveOrganization(t *testing.T) {
	assert := assert.New(t)
	coreService := hcmcore.NewCoreServiceServer()
	client := coreService.GetSupabaseCommunityClient()

	t.Run("create an organization using faked data", func(t *testing.T) {
		faker := faker.New()

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

			if err != nil {
				t.Error(err)
			}

			assert.NotEmpty(
				res,
				"no response was returned by the function",
			)

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
			_, err = client.From("organizations").Delete("", "").Eq("id", strconv.FormatInt(updatedOrganization.GetId(), 10)).ExecuteTo(nil)
			if err != nil {
				t.Log(err)
			}
		}

	})

}

func TestSaveRole(t *testing.T) {
	assert := assert.New(t)
	coreService := hcmcore.NewCoreServiceServer()
	client := coreService.GetSupabaseCommunityClient()

	t.Run("create a new role for an organization", func(t *testing.T) {
		organization := createMockOrganization(t)
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

			role := createMockRole(t, roleName, organization.Id)

			res, err := coreService.SaveRole(context.Background(), &pb.CoreServiceRequest{
				UsedClient: pb.CoreServiceRequest_C_SUPABASE,
				SetterRequest: &pb.SetterRequest{
					RoleTarget: role,
				},
			})

			if err != nil {
				t.Log(err)
			}

			assert.NotEmpty(
				res,
				"did not return a proper response",
			)

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

	t.Run("create a new team for an organization", func(t *testing.T) {
		organization := createMockOrganization(t)

		mockTeams := []string{
			"English Department",
			"Engineering Department (IT)",
			"Engineering Department (Civil)",
			"Medical Department",
			"Mathematics Department",
		}

		for _, teamName := range mockTeams {
			team := createMockTeam(t, teamName, organization.GetId())

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
	client := coreService.GetSupabaseCommunityClient()

	t.Run("should create a new identity card for a mock worker", func(t *testing.T) {
		faker := faker.New()
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

			// cleanup
			_, err = client.From("workerIdentityCards").Delete("", "").Eq("id", strconv.FormatInt(updatedCard.Id, 10)).ExecuteTo(nil)
			if err != nil {
				t.Error(err)
			}
		}

		// cleanup
		_, err := coreService.DeleteWorkerById(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			SetterRequest: &pb.SetterRequest{
				TargetId: &worker.Id,
			},
		})
		if err != nil {
			t.Fatal(err)
		}
	})

}

func TestSaveMember(t *testing.T) {
	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func TestSaveCompensation(t *testing.T) {
	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func TestSaveAddition(t *testing.T) {
	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func TestSaveDeduction(t *testing.T) {
	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func TestSaveShift(t *testing.T) {
	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func TestDeleteWorkerById(t *testing.T) {
	assert := assert.New(t)
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should delete the mock user persisted on the database", func(t *testing.T) {
		worker := createMockWorker(t, true)

		res, err := coreService.DeleteWorkerById(context.Background(), &pb.CoreServiceRequest{
			SetterRequest: &pb.SetterRequest{
				TargetId: &worker.Id,
			},
		})
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(
			res.GetCode(),
			pb.CoreServiceResponse_C_NOERROR,
			"did not successfully returned a C_NOERROR code",
		)
	})
}

func TestDeleteOrganizationById(t *testing.T) {
	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func TestDeleteRoleById(t *testing.T) {
	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func TestDeleteTeamById(t *testing.T) {
	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func TestDeleteWorkerIdentityCardById(t *testing.T) {
	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func TestDeleteCompensationById(t *testing.T) {
	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func TestDeleteMemberById(t *testing.T) {
	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func TestDeletePayrollById(t *testing.T) {
	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func TestDeleteAdditionById(t *testing.T) {
	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func TestDeleteDeductionById(t *testing.T) {
	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func TestDeleteShiftById(t *testing.T) {
	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func createMockOrganization(t *testing.T) *pb.Organization {
	coreService := hcmcore.NewCoreServiceServer()
	faker := faker.New()

	// Create a mock organization
	res, err := coreService.SaveOrganization(context.Background(), &pb.CoreServiceRequest{
		SetterRequest: &pb.SetterRequest{
			OrganizationTarget: &pb.Organization{
				Name:     faker.Company().Name(),
				Industry: pb.Organization_EDUCATION,
				Flags:    uint32(pb.Organization_UNKNOWN),
			},
		},
	})

	if err != nil {
		t.Log(err)
	}

	return res.GetSetterResponse().GetUpdatedOrganizationTarget()
}

func createMockTeam(_ *testing.T, name string, organizationId int64) *pb.Team {
	team := &pb.Team{
		Name:           name,
		OrganizationId: organizationId,
		Flags:          uint32(pb.Team_UNKNOWN),
	}

	return team
}

func createMockRole(_ *testing.T, name string, organizationId int64) *pb.Role {
	role := &pb.Role{
		Name:           name,
		OrganizationId: organizationId,
		Flags:          uint32(pb.Role_UNKNOWN),
	}

	return role
}

func createMockWorker(t *testing.T, persist bool) *pb.Worker {
	coreService := hcmcore.NewCoreServiceServer()
	faker := faker.New()

	fakeAddress := faker.Address()
	middleName := faker.Person().LastName()

	worker := &pb.Worker{
		CreatedAt:  timestamppb.Now(),
		FirstName:  faker.Person().FirstNameMale(),
		MiddleName: &middleName,
		LastName:   faker.Person().LastName(),
		Birthdate:  timestamppb.New(faker.Time().Time(time.Now())),
		Email:      faker.Internet().Email(),
		Username:   faker.Internet().User(),
		Addresses: []*structpb.Struct{
			{
				Fields: map[string]*structpb.Value{
					"addressLine1": structpb.NewStringValue(fakeAddress.Address()),
					"addressLine2": structpb.NewStringValue(fakeAddress.Address()),
					"city":         structpb.NewStringValue(fakeAddress.City()),
					"state":        structpb.NewStringValue(fakeAddress.State()),
					"country":      structpb.NewStringValue(fakeAddress.Country()),
				},
			},
		},
		Flags: uint32(pb.Worker_G_MALE),
	}

	if persist {
		res, err := coreService.SaveWorker(context.Background(), &pb.CoreServiceRequest{
			UsedClient: pb.CoreServiceRequest_C_SUPABASE,
			SetterRequest: &pb.SetterRequest{
				WorkerTarget: worker,
			},
		})

		if err != nil {
			t.Fatal(err)
		}

		worker = res.SetterResponse.UpdatedWorkerTarget
	}

	return worker
}
