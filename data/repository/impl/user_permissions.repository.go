package impl

import (
	"context"
	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
	"github.com/ngdangkietswe/swe-auth-service/data/repository"
)

type UserPermissionsRepository struct {
	entClient *ent.Client
}

// CreateUserPermissions is a function that creates user permissions.
func (u UserPermissionsRepository) CreateUserPermissions(ctx context.Context, userId string, permissionIds []string) error {
	var entPs []*ent.UsersPermissionCreate

	for _, permissionId := range permissionIds {
		entPs = append(entPs, u.entClient.UsersPermission.Create().SetUserID(uuid.MustParse(userId)).SetPermissionID(uuid.MustParse(permissionId)))
	}

	_, err := u.entClient.UsersPermission.CreateBulk(entPs...).Save(ctx)
	return err
}

func NewUserPermissionsRepository(entClient *ent.Client) repository.IUserPermissionsRepository {
	return &UserPermissionsRepository{
		entClient: entClient,
	}
}
