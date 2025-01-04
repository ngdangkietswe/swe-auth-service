package resource

import (
	"context"
	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
)

type IResourceRepository interface {
	ExistsById(ctx context.Context, id string) (bool, error)
	FindAllByIds(ctx context.Context, ids []uuid.UUID) ([]*ent.Resource, error)
}
