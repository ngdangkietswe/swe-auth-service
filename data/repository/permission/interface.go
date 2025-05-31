package permission

import (
	"context"
	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/common"
)

type IPermissionRepository interface {
	Save(ctx context.Context, tx *ent.Tx, permission *auth.UpsertPermissionReq) (*ent.Permission, error)
	FindById(ctx context.Context, id string) (*ent.Permission, error)
	FindAllByIdIn(ctx context.Context, ids []uuid.UUID) ([]*ent.Permission, error)
	FindAll(ctx context.Context, req *auth.ListPermissionsReq, pageable *common.Pageable) ([]*ent.Permission, int64, error)
	AssignPermissions(ctx context.Context, tx *ent.Tx, req *auth.AssignPermissionsReq) error
	ExistsByActionAndResource(ctx context.Context, actionId string, resourceId string) (bool, error)
	ExistsByIdIn(ctx context.Context, ids []string) (bool, error)
	ExistsById(ctx context.Context, id string) (bool, error)
}
