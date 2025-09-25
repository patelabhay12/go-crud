package initializers

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

// Exported variable accessible from other packages
var DB *pgx.Conn

func ConnectToDB() {
	dbURL := os.Getenv("SUPABASE_URL")
	if dbURL == "" {
		log.Fatal("Please set the SUPABASE_URL environment variable")
	}

	var err error
	DB, err = pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	log.Println("Connected to Supabase successfully!")
}
 