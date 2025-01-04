package permission

import (
	"context"
	"github.com/google/uuid"
	"github.com/google/wire"
	"github.com/ngdangkietswe/swe-auth-service/data/datasource"
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/common"
)

type IPermissionRepository interface {
	UpsertPermission(ctx context.Context, permission *auth.UpsertPermissionReq) (*ent.Permission, error)
	FindPermissionById(ctx context.Context, id string) (*ent.Permission, error)
	FindAllByIds(ctx context.Context, ids []uuid.UUID) ([]*ent.Permission, error)
	ListPermissions(ctx context.Context, req *auth.ListPermissionsReq, pageable *common.Pageable) ([]*ent.Permission, int64, error)
	AssignPermissions(ctx context.Context, req *auth.AssignPermissionsReq) error
	ExistsByActionAndResource(ctx context.Context, actionId string, resourceId string) (bool, error)
	ExistsAllByIds(ctx context.Context, ids []string) (bool, error)
	ExistsById(ctx context.Context, id string) (bool, error)
}

// ProvidePermissionRepository is a function to provide a permission repository
func ProvidePermissionRepository() (p IPermissionRepository) {
	wire.Build(
		datasource.ProvideEntClient,
		NewPermissionRepository,
	)
	return
}
