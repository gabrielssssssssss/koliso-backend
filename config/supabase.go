package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/supabase-community/supabase-go"
)

func NewSupabaseClient() *supabase.Client {
	SUPABASE_URL := os.Getenv("SUPABASE_URL")
	SUPABASE_ANON_KEY := os.Getenv("SUPABASE_ANON_KEY")

	client, err := supabase.NewClient(SUPABASE_URL, SUPABASE_ANON_KEY, &supabase.ClientOptions{})
	if err != nil {
		fmt.Println("Failed to initalize the client: ", err)
	}

	return client
}

func NewSupabaseContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
