package impl

import (
	"context"
	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
	"github.com/ngdangkietswe/swe-auth-service/data/ent/userspermission"
	"github.com/ngdangkietswe/swe-auth-service/data/repository"
)

type UserPermissionsRepository struct {
	entClient *ent.Client
}

// FindAllByUserId is a function that finds all user permissions by user id.
func (u UserPermissionsRepository) FindAllByUserId(ctx context.Context, userId uuid.UUID) ([]*ent.UsersPermission, error) {
	return u.entClient.UsersPermission.Query().Where(userspermission.UserIDEQ(userId)).All(ctx)
}

// DeleteAllByUserId is a function that deletes all user permissions by user id.
func (u UserPermissionsRepository) DeleteAllByUserId(ctx context.Context, userId string) error {
	_, err := u.entClient.UsersPermission.Delete().Where(userspermission.UserIDEQ(uuid.MustParse(userId))).Exec(ctx)
	return err
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
