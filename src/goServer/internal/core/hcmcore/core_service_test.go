package hcmcore_test

import (
	"context"
	"errors"
	"goServer/internal/core/hcmcore"
	"goServer/internal/core/pb"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	ErrCoreServiceServerUnimplementedMethod = errors.New("CoreServiceServer: Not implemented")
)

const (
	TestSaveWorkerParams_N_WORKERS = 10
)

func TestGetWorkerById(t *testing.T) {
	coreService := hcmcore.NewCoreServiceServer()

	workerId := int64(0)

	_, err := coreService.GetWorkerById(context.TODO(), &pb.CoreServiceRequest{TargetId: &workerId})
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

	t.Run("create a new worker for each faked data generated", func(t *testing.T) {
		faker := faker.New()

		for i := 0; i < TestSaveWorkerParams_N_WORKERS; i++ {
			fakeAddress := faker.Address()
			middleName := faker.Person().LastName()
			worker := &pb.Worker{
				Uuid:       uuid.NewString(),
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

			res, err := coreService.SaveWorker(context.Background(), &pb.CoreServiceRequest{
				UsedClient:       pb.CoreServiceRequest_C_SUPABASE,
				SaveWorkerTarget: worker,
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
				res.SaveWorkerResult,
				"did not return proper response",
			)

			saveWorkerResult := res.GetSaveWorkerResult()
			updatedWorker := saveWorkerResult.GetUpdatedTarget()

			assert.NotEmpty(
				updatedWorker,
				"did not return the updated target",
			)

		}

	})

	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func TestSaveOrganization(t *testing.T) {
	// coreService := hcmcore.NewCoreServiceServer()

	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func TestSaveRole(t *testing.T) {
	// coreService := hcmcore.NewCoreServiceServer()

	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func TestSaveTeam(t *testing.T) {
	// coreService := hcmcore.NewCoreServiceServer()

	t.Error(ErrCoreServiceServerUnimplementedMethod)
}

func TestDeleteWorkerById(t *testing.T) {
	t.Error(ErrCoreServiceServerUnimplementedMethod)
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
