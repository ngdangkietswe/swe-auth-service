package datasource

import (
	"context"
	"fmt"
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
	"github.com/ngdangkietswe/swe-go-common-shared/config"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// NewEntClient initializes and returns an Ent client connected to the configured Postgres database.
func NewEntClient() *ent.Client {
	// Load DB config
	dbHost := config.GetString("DB_HOST", "localhost")
	dbPort := config.GetInt("DB_PORT", 5432)
	dbUser := config.GetString("DB_USER", "postgres")
	dbName := config.GetString("DB_NAME", "SweIntegration")
	dbPassword := config.GetString("DB_PASSWORD", "123456")
	dbSSLMode := config.GetString("DB_SSL_MODE", "disable")
	dbSearchPath := config.GetString("DB_SEARCH_PATH", "sweauth")
	enableDebug := config.GetBool("ENT_DEBUG", false)

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s search_path=%s",
		dbHost, dbPort, dbUser, dbName, dbPassword, dbSSLMode, dbSearchPath,
	)

	// Build Ent client options
	opts := []ent.Option{
		ent.Log(func(args ...any) {
			log.Printf("[ENT DEBUG] %v", fmt.Sprint(args...))
		}),
	}

	if enableDebug {
		opts = append(opts, ent.Debug())
	}

	// Open connection
	client, err := ent.Open("postgres", dsn, opts...)
	if err != nil {
		log.Fatalf("[ERROR] Failed opening connection to postgres: %v", err)
	}

	// Use a context with timeout for schema migration
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = client.Schema.Create(ctx); err != nil {
		log.Fatalf("[ERROR] Failed creating schema resources: %v", err)
	}

	log.Println("[INFO] Ent client connected and schema initialized successfully.")
	return client
}
