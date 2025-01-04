package user_permissions

import (
	"context"
	"github.com/google/uuid"
	"github.com/google/wire"
	"github.com/ngdangkietswe/swe-auth-service/data/datasource"
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
)

type IUserPermissionsRepository interface {
	CreateUserPermissions(ctx context.Context, userId string, permissionIds []string) error
	DeleteAllByUserId(ctx context.Context, userId string) error
	FindAllByUserId(ctx context.Context, userId uuid.UUID) ([]*ent.UsersPermission, error)
}

// ProvideUserPermissionsRepository is a function to provide a user permission repository
func ProvideUserPermissionsRepository() (u IUserPermissionsRepository) {
	wire.Build(
		datasource.ProvideEntClient,
		NewUserPermissionsRepository,
	)
	return
}
