package hcmcore_test

import (
	"context"
	"errors"
	"goServer/internal/core/hcmcore"
	"goServer/internal/core/pb"
	"strconv"
	"testing"
	"time"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	ErrCoreServiceServerUnimplementedMethod = errors.New("CoreServiceServer: Not implemented")
	DefaultFaker                            = faker.New()
)

const (
	TestSaveWorkerParams_N_WORKERS      = 10
	TestSaveOrganization_N_ORGANIZATION = 10
)

func createMockOrganization(t *testing.T, persist bool) (organization *pb.Organization) {
	organization = &pb.Organization{
		Name:     DefaultFaker.Company().Name(),
		Industry: pb.Organization_EDUCATION,
		Flags:    uint32(pb.Organization_UNKNOWN),
	}

	if persist {
		organization = saveMockOrganization(t, organization)
	}

	return
}

func saveMockOrganization(t *testing.T, mock *pb.Organization) *pb.Organization {
	coreService := hcmcore.NewCoreServiceServer()
	res, err := coreService.SaveOrganization(context.Background(), &pb.CoreServiceRequest{
		SetterRequest: &pb.SetterRequest{
			OrganizationTarget: mock,
		},
	})

	if err != nil {
		t.Log(err)
	}

	return res.SetterResponse.GetUpdatedOrganizationTarget()
}

func createMockTeam(t *testing.T, name string, organizationId int64, persist bool) *pb.Team {
	coreService := hcmcore.NewCoreServiceServer()
	team := &pb.Team{
		Name:           name,
		OrganizationId: organizationId,
		Flags:          uint32(pb.Team_UNKNOWN),
	}

	if persist {
		res, err := coreService.SaveTeam(context.Background(), &pb.CoreServiceRequest{
			SetterRequest: &pb.SetterRequest{
				TeamTarget: team,
			},
		})

		if err != nil {
			t.Error(err)
		}

		team = res.SetterResponse.GetUpdatedTeamTarget()
	}

	return team
}

func createMockRole(t *testing.T, name string, organizationId int64, persist bool) *pb.Role {
	coreService := hcmcore.NewCoreServiceServer()
	role := &pb.Role{
		Name:           name,
		OrganizationId: organizationId,
		Flags:          uint32(pb.Role_UNKNOWN),
	}

	if persist {
		res, err := coreService.SaveRole(context.Background(), &pb.CoreServiceRequest{
			SetterRequest: &pb.SetterRequest{
				RoleTarget: role,
			},
		})

		if err != nil {
			t.Error(err)
		}

		role = res.SetterResponse.GetUpdatedRoleTarget()
	}

	return role
}

func createMockShift(t *testing.T, organizationId int64, typ pb.SetterRequest_ShiftOpType, persist bool) any {
	coreService := hcmcore.NewCoreServiceServer()

	var resp any

	coreServiceReq := &pb.CoreServiceRequest{
		SetterRequest: &pb.SetterRequest{
			TargetShiftType: typ.Enum(),
		},
	}

	switch typ {
	case pb.SetterRequest_T_OVERRIDESHIFT:
		overrideShift := &pb.OverrideShift{
			OrganizationId:   organizationId,
			Name:             "Rom Vales' Override Shift",
			OverrideClockIn:  timestamppb.Now(),
			OverrideClockOut: timestamppb.Now(),
		}

		coreServiceReq.SetterRequest.OverrideShiftTarget = overrideShift
		resp = overrideShift
	case pb.SetterRequest_T_SHIFT:
		shift := &pb.Shift{
			OrganizationId: organizationId,
			Name:           "Finance Shift",
			ClockIn:        timestamppb.Now(),
			ClockOut:       timestamppb.Now(),
		}

		coreServiceReq.SetterRequest.ShiftTarget = shift
		resp = shift
	}

	if !persist {
		return resp
	}

	res, err := coreService.SaveShift(context.Background(), coreServiceReq)
	if err != nil {
		t.Fatal(err)
	}

	switch typ {
	case pb.SetterRequest_T_OVERRIDESHIFT:
		resp = res.SetterResponse.GetUpdatedOverrideShiftTarget()
	case pb.SetterRequest_T_SHIFT:
		resp = res.SetterResponse.GetUpdatedShiftTarget()
	}

	return resp
}

func createMockWorker(t *testing.T, persist bool) *pb.Worker {
	coreService := hcmcore.NewCoreServiceServer()
	faker := DefaultFaker

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

func assertCheckForMissingResponse(t *testing.T, err error, res *pb.CoreServiceResponse) {
	assert := assert.New(t)

	assert.NoError(
		err,
		"expected for the function to work properly",
	)

	assert.NotEmpty(
		res,
		"expected to return a proper response",
	)

	assert.Equal(
		pb.CoreServiceResponse_C_NOERROR,
		res.GetCode(),
		"did not successfully returned a C_NOERROR code",
	)
}

func cleanCreatedMockDataInTableNameById(t *testing.T, tableName string, id int64) {
	client := hcmcore.NewCoreServiceServer().GetSupabaseCommunityClient()

	_, err := client.From(tableName).Delete("", "").Eq("id", strconv.FormatInt(id, 10)).ExecuteTo(nil)
	if err != nil {
		t.Fatal(err)
	}
}
