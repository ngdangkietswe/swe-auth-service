package resource

import (
	"context"
	"github.com/google/uuid"
	"github.com/google/wire"
	"github.com/ngdangkietswe/swe-auth-service/data/datasource"
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
)

type IResourceRepository interface {
	ExistsById(ctx context.Context, id string) (bool, error)
	FindAllByIds(ctx context.Context, ids []uuid.UUID) ([]*ent.Resource, error)
}

// ProvideResourceRepository is a function to provide a resource repository
func ProvideResourceRepository() (r IResourceRepository) {
	wire.Build(
		datasource.ProvideEntClient,
		NewResourceRepository,
	)
	return
}
