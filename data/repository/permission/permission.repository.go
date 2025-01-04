package permission

import (
	"context"
	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
	"github.com/ngdangkietswe/swe-auth-service/data/ent/permission"
	"github.com/ngdangkietswe/swe-auth-service/grpc/utils"
	commonutils "github.com/ngdangkietswe/swe-go-common-shared/util"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/common"
)

type permissionRepository struct {
	entClient *ent.Client
}

// FindAllByIds is a function that finds all permissions by IDs.
func (p permissionRepository) FindAllByIds(ctx context.Context, ids []uuid.UUID) ([]*ent.Permission, error) {
	return p.entClient.Permission.Query().Where(permission.IDIn(ids...)).All(ctx)
}

// ExistsById is a function that checks if a permission exists by ID.
func (p permissionRepository) ExistsById(ctx context.Context, id string) (bool, error) {
	exists, err := p.entClient.Permission.Query().Where(permission.ID(uuid.MustParse(id))).Exist(ctx)
	return exists, err
}

func (p permissionRepository) ExistsAllByIds(ctx context.Context, ids []string) (bool, error) {
	uIds := commonutils.Convert2UUID(ids)
	count, err := p.entClient.Permission.Query().Where(permission.IDIn(uIds...)).Count(ctx)
	return count == len(ids), err
}

func (p permissionRepository) ExistsByActionAndResource(ctx context.Context, actionId string, resourceId string) (bool, error) {
	exists, err := p.entClient.Permission.Query().Where(
		permission.ActionIDEQ(uuid.MustParse(actionId)),
		permission.ResourceIDEQ(uuid.MustParse(resourceId)),
	).Exist(ctx)

	return exists, err
}

// UpsertPermission is a function that upserts a permission.
// If the permission has an ID, it will update the permission.
// Otherwise, it will create a new permission
func (p permissionRepository) UpsertPermission(ctx context.Context, permission *auth.UpsertPermissionReq) (*ent.Permission, error) {
	if permission.Id != nil {
		return p.entClient.Permission.UpdateOneID(uuid.MustParse(*permission.Id)).
			SetActionID(uuid.MustParse(permission.ActionId)).
			SetResourceID(uuid.MustParse(permission.ResourceId)).
			Save(ctx)
	} else {
		return p.entClient.Permission.Create().
			SetActionID(uuid.MustParse(permission.ActionId)).
			SetResourceID(uuid.MustParse(permission.ResourceId)).
			Save(ctx)
	}
}

func (p permissionRepository) FindPermissionById(ctx context.Context, id string) (*ent.Permission, error) {
	//TODO implement me
	panic("implement me")
}

// ListPermissions is a function that lists permissions.
func (p permissionRepository) ListPermissions(ctx context.Context, req *auth.ListPermissionsReq, pageable *common.Pageable) ([]*ent.Permission, int64, error) {
	entPs := p.entClient.Permission.Query()

	if req.ActionId != nil {
		entPs = entPs.Where(permission.ActionIDEQ(uuid.MustParse(*req.ActionId)))
	}
	if req.ResourceId != nil {
		entPs = entPs.Where(permission.ResourceIDEQ(uuid.MustParse(*req.ResourceId)))
	}
	if req.Search != nil && *req.Search != "" {
		entPs = entPs.Where(permission.DescriptionContainsFold(*req.Search))
	}

	count, err := entPs.Count(ctx)
	data, err := entPs.Limit(int(pageable.Size)).Offset(int(utils.AsOffset(pageable.Page, pageable.Size))).All(ctx)
	return data, int64(count), err
}

func (p permissionRepository) AssignPermissions(ctx context.Context, req *auth.AssignPermissionsReq) error {
	//TODO implement me
	panic("implement me")
}

func NewPermissionRepository(entClient *ent.Client) IPermissionRepository {
	return &permissionRepository{
		entClient: entClient,
	}
}
