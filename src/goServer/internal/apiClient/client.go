package apiClient

import (
	"goServer/internal/core/pb"
	goServerErrors "goServer/internal/errors"

	supabase "github.com/nedpals/supabase-go"
	supabaseCommunity "github.com/supabase-community/supabase-go"
)

var (
	unsafeSupabaseClient          = NewSupabaseClient()
	unsafeSupabaseCommunityClient = NewSupabaseCommunityClient()
)

type RequestUsedClient struct {
	supabaseClient           *supabase.Client
	community_supabaseClient *supabaseCommunity.Client
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
			c.community_supabaseClient = unsafeSupabaseCommunityClient
		} else {
			c.supabaseClient = NewSupabaseClient()
			c.community_supabaseClient = NewSupabaseCommunityClient()
		}
	default:
		return c, goServerErrors.ErrInvalidClientFromRequestUnimplemented
	}

	return c, nil
}

func (c *RequestUsedClient) GetSupabaseClient() *supabase.Client {
	// Revert to using unsafe supabase client
	if c.supabaseClient == nil {
		return unsafeSupabaseClient
	}

	return c.supabaseClient
}

func (c *RequestUsedClient) GetSupabaseCommunityClient() *supabaseCommunity.Client {
	// Revert to using unsafe community client
	if c.community_supabaseClient == nil {
		return unsafeSupabaseCommunityClient
	}

	return c.community_supabaseClient
}
