package impl

import (
	"context"
	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
	"github.com/ngdangkietswe/swe-auth-service/data/ent/action"
	"github.com/ngdangkietswe/swe-auth-service/data/repository"
)

type actionRepository struct {
	entClient *ent.Client
}

// FindAllByIds is a function that finds all actions by IDs.
func (a actionRepository) FindAllByIds(ctx context.Context, ids []uuid.UUID) ([]*ent.Action, error) {
	return a.entClient.Action.Query().Where(action.IDIn(ids...)).All(ctx)
}

// ExistsById is a function that checks if an action exists by ID.
func (a actionRepository) ExistsById(ctx context.Context, id string) (bool, error) {
	return a.entClient.Action.Query().Where(action.IDEQ(uuid.MustParse(id))).Exist(ctx)
}

func NewActionRepository(entClient *ent.Client) repository.IActionRepository {
	return &actionRepository{entClient: entClient}
}
