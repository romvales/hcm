package apiClient

import (
	"goServer/internal/core/pb"
	goServerErrors "goServer/internal/errors"

	supabase "github.com/nedpals/supabase-go"
)

var (
	unsafeSupabaseClient = NewSupabaseClient()
)

type RequestUsedClient struct {
	supabaseClient *supabase.Client
}

type RequestRequiredFields interface {
	GetUsedClient() pb.CoreServiceRequest_CoreServiceRequestClient
	GetClientUnsafe() bool
}

func (c *RequestUsedClient) GetClientFromRequest(req RequestRequiredFields) (*RequestUsedClient, error) {
	usedClient := req.GetUsedClient()
	isUnsafe := req.GetClientUnsafe()

	switch usedClient {
	case pb.CoreServiceRequest_C_SUPABASE:
		if isUnsafe {
			c.supabaseClient = unsafeSupabaseClient
		} else {
			c.supabaseClient = NewSupabaseClient()
		}
	default:
		return c, goServerErrors.ErrInvalidClientFromRequestUnimplemented
	}

	return c, nil
}

func (c *RequestUsedClient) GetSupabaseClient() *supabase.Client {
	return c.supabaseClient
}
