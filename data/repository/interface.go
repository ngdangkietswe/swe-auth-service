package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/common"
)

type (
	IAuthRepository interface {
		UpsertUser(ctx context.Context, user *auth.User) (*ent.User, error)
		FindById(ctx context.Context, id string) (*ent.User, error)
		FindByUsername(ctx context.Context, username string) (*ent.User, error)
		FindByUsernameOrEmail(ctx context.Context, username string, email string) (*ent.User, error)
		EnableOrDisable2FA(ctx context.Context, userId string, enable bool) (*ent.User, error)
		ExistsById(ctx context.Context, id string) (bool, error)
	}

	IPermissionRepository interface {
		UpsertPermission(ctx context.Context, permission *auth.UpsertPermissionReq) (*ent.Permission, error)
		FindPermissionById(ctx context.Context, id string) (*ent.Permission, error)
		FindAllByIds(ctx context.Context, ids []uuid.UUID) ([]*ent.Permission, error)
		ListPermissions(ctx context.Context, req *auth.ListPermissionsReq, pageable *common.Pageable) ([]*ent.Permission, int64, error)
		AssignPermissions(ctx context.Context, req *auth.AssignPermissionsReq) error
		ExistsByActionAndResource(ctx context.Context, actionId string, resourceId string) (bool, error)
		ExistsAllByIds(ctx context.Context, ids []string) (bool, error)
		ExistsById(ctx context.Context, id string) (bool, error)
	}

	IUserPermissionsRepository interface {
		CreateUserPermissions(ctx context.Context, userId string, permissionIds []string) error
		DeleteAllByUserId(ctx context.Context, userId string) error
		FindAllByUserId(ctx context.Context, userId uuid.UUID) ([]*ent.UsersPermission, error)
	}

	IActionRepository interface {
		ExistsById(ctx context.Context, id string) (bool, error)
		FindAllByIds(ctx context.Context, ids []uuid.UUID) ([]*ent.Action, error)
	}

	IResourceRepository interface {
		ExistsById(ctx context.Context, id string) (bool, error)
		FindAllByIds(ctx context.Context, ids []uuid.UUID) ([]*ent.Resource, error)
	}
)
