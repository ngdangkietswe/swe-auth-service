package action

import (
	"context"
	"github.com/google/uuid"
	"github.com/google/wire"
	"github.com/ngdangkietswe/swe-auth-service/data/datasource"
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
)

type IActionRepository interface {
	ExistsById(ctx context.Context, id string) (bool, error)
	FindAllByIds(ctx context.Context, ids []uuid.UUID) ([]*ent.Action, error)
}

// ProvideActionRepository is a function to provide an action repository
func ProvideActionRepository() (a IActionRepository) {
	wire.Build(
		datasource.ProvideEntClient,
		NewActionRepository,
	)
	return
}
