package hcmcore_test

import (
	"context"
	"goServer/internal/core/hcmcore"
	"goServer/internal/core/pb"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendJoinRequest(t *testing.T) {
	assert := assert.New(t)
	coreService := hcmcore.NewCoreServiceServer()

	t.Run("should be able for a mock worker to send a join request to an organization", func(t *testing.T) {
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

		assert.Equal(res.Code, pb.CoreServiceResponse_C_NOERROR, "expected for the function to execute properly")

		cleanCreatedMockDataInTableNameById(t, "organizations", organization.Id)
		cleanCreatedMockDataInTableNameById(t, "workers", worker.Id)
	})

}
