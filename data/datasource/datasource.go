package datasource

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
	"github.com/ngdangkietswe/swe-go-common-shared/config"
	"log"

	_ "github.com/lib/pq"
)

func NewEntClient() *ent.Client {
	client, err := ent.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
			config.GetString("DB_HOST", "localhost"),
			config.GetInt("DB_PORT", 5432),
			config.GetString("DB_USER", "postgres"),
			config.GetString("DB_NAME", "SweAuth"),
			config.GetString("DB_PASSWORD", "123456"),
			config.GetString("DB_SSL_MODE", "disable")),
	)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}

// ProvideEntClient is a function to provide an ent client
func ProvideEntClient() (e *ent.Client) {
	wire.Build(
		NewEntClient,
	)
	return
}
