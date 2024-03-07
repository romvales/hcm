package hcmcore_test

import (
	"context"
	"goServer/internal/core/hcmcore"
	"goServer/internal/core/pb"
	"testing"

	"github.com/stretchr/testify/assert"
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

		assert.NotEmpty(workerResult, "did not return a result in the response")
		assert.Equal(worker.Id, workerResult.Id, "did not match the expected id")
		assert.Equal(worker.FirstName, workerResult.FirstName, "did not match the expected first name")
		assert.Equal(worker.LastName, workerResult.LastName, "did not match the expected last name")

		cleanCreatedMockDataInTableNameById(t, "workers", worker.Id)
	})

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
