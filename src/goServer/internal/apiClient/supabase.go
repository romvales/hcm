package apiClient

import (
	"errors"
	"log"
	"os"

	supabase "github.com/nedpals/supabase-go"
)

var (
	SUPABASE_URL                    = os.Getenv("SUPABASE_URL")
	SUPABASE_LOCAL_SERVICE_ROLE_KEY = os.Getenv("SUPABASE_LOCAL_SERVICE_ROLE_KEY")
)

var (
	ErrSupabaseUrlVariableNotSet         = errors.New("SUPABASE_URL was not found in the environment")
	ErrSupabaseLocalServiceRoleKeyNotSet = errors.New("SUPABASE_LOCAL_SERVICE_ROLE_KEY was not found in the environment")
)

func init() {

	if SUPABASE_URL == "" {
		log.Fatal(ErrSupabaseUrlVariableNotSet)
	}

	if SUPABASE_LOCAL_SERVICE_ROLE_KEY == "" {
		log.Fatal(ErrSupabaseLocalServiceRoleKeyNotSet)
	}

}

func NewSupabaseClient() (client *supabase.Client) {
	url, serviceRoleKey := SUPABASE_URL, SUPABASE_LOCAL_SERVICE_ROLE_KEY
	client = supabase.CreateClient(url, serviceRoleKey)
	return
}
