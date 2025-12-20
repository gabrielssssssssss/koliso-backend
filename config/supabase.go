package config

import (
	"context"
	"os"
	"time"

	"github.com/supabase-community/supabase-go"
)

func NewSupabaseDatabase() *supabase.Client {
	SUPABASE_URL := os.Getenv("SUPABASE_URL")
	SUPABASE_ANON_KEY := os.Getenv("SUPABASE_ANON_KEY")

	client, err := supabase.NewClient(SUPABASE_URL, SUPABASE_ANON_KEY, &supabase.ClientOptions{})
	if err != nil {
		return nil
	}

	return client
}

func NewSupabaseContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
