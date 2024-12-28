package impl

import (
	"context"
	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
	"github.com/ngdangkietswe/swe-auth-service/data/ent/resource"
	"github.com/ngdangkietswe/swe-auth-service/data/repository"
)

type resourceRepository struct {
	entClient *ent.Client
}

// FindAllByIds is a function that finds all resources by IDs
func (r resourceRepository) FindAllByIds(ctx context.Context, ids []uuid.UUID) ([]*ent.Resource, error) {
	return r.entClient.Resource.Query().Where(resource.IDIn(ids...)).All(ctx)
}

// ExistsById is a function that checks if a resource exists by ID
func (r resourceRepository) ExistsById(ctx context.Context, id string) (bool, error) {
	return r.entClient.Resource.Query().Where(resource.IDEQ(uuid.MustParse(id))).Exist(ctx)
}

func NewResourceRepository(entClient *ent.Client) repository.IResourceRepository {
	return &resourceRepository{entClient: entClient}
}
