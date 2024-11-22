package datasource

import (
	"context"
	"fmt"
	"github.com/ngdangkietswe/swe-auth-service/configs"
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
	"log"
)

func NewEntClient() *ent.Client {
	client, err := ent.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s",
			configs.GlobalConfig.DbHost,
			configs.GlobalConfig.DbPort,
			configs.GlobalConfig.DbUser,
			configs.GlobalConfig.DbName,
			configs.GlobalConfig.DbPassword),
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
