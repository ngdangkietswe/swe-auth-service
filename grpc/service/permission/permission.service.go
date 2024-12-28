package permission

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
	"github.com/ngdangkietswe/swe-auth-service/data/repository"
	"github.com/ngdangkietswe/swe-auth-service/grpc/mapper"
	"github.com/ngdangkietswe/swe-auth-service/grpc/utils"
	validator "github.com/ngdangkietswe/swe-auth-service/grpc/validator/permission"
	commonutil "github.com/ngdangkietswe/swe-go-common-shared/util"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/common"
	"github.com/samber/lo"
	"log"
)

type permissionSvc struct {
	actionRepo          repository.IActionRepository
	resourceRepo        repository.IResourceRepository
	permissionRepo      repository.IPermissionRepository
	userPermissionsRepo repository.IUserPermissionsRepository
	authRepo            repository.IAuthRepository
	permissionValidator validator.IPermissionValidator
}

// UpsertPermission is a function that upserts a permission.
func (p permissionSvc) UpsertPermission(ctx context.Context, req *auth.UpsertPermissionReq) (*common.UpsertResp, error) {
	err := p.permissionValidator.ValidateUpsertPermission(req)
	if err != nil {
		log.Printf("error validating upsert permission: %v", err)
		return nil, err
	}

	if req.Id != nil {
		exists, err := p.permissionRepo.ExistsById(ctx, *req.Id)
		if err != nil {
			log.Printf("error finding permission by id: %v", err)
			return nil, err
		} else if !exists {
			log.Printf("permission with id %s does not exist", *req.Id)
			return nil, errors.New("permission does not exist")
		}
	}

	permission, err := p.permissionRepo.UpsertPermission(ctx, req)
	if err != nil {
		log.Printf("error upserting permission: %v", err)
		return nil, err
	}

	return &common.UpsertResp{
		Success: true,
		Resp: &common.UpsertResp_Data_{
			Data: &common.UpsertResp_Data{
				Id: permission.ID.String(),
			},
		},
	}, nil
}

// ListPermissions is a function that lists permissions.
func (p permissionSvc) ListPermissions(ctx context.Context, req *auth.ListPermissionsReq) (*auth.ListPermissionsResp, error) {
	err := p.permissionValidator.ValidateListPermissions(req)
	if err != nil {
		log.Printf("error validating list permissions: %v", err)
		return nil, err
	}

	normalizePageable := utils.NormalizePageable(req.Pageable)
	permissions, err := p.permissionRepo.ListPermissions(ctx, req, normalizePageable)
	if err != nil {
		log.Printf("error listing permissions: %v", err)
		return nil, err
	}

	actionIds := lo.Map(permissions, func(permission *ent.Permission, _ int) uuid.UUID {
		return permission.ActionID
	})
	resourceIds := lo.Map(permissions, func(permission *ent.Permission, _ int) uuid.UUID {
		return permission.ResourceID
	})

	actions, err := p.actionRepo.FindAllByIds(ctx, lo.Uniq(actionIds))
	if err != nil {
		log.Printf("error finding actions by ids: %v", err)
		return nil, err
	}
	actionMap := commonutil.Convert2Map(actions, func(action *ent.Action) uuid.UUID {
		return action.ID
	})

	resources, err := p.resourceRepo.FindAllByIds(ctx, lo.Uniq(resourceIds))
	if err != nil {
		log.Printf("error finding resources by ids: %v", err)
		return nil, err
	}
	resourceMap := commonutil.Convert2Map(resources, func(resource *ent.Resource) uuid.UUID {
		return resource.ID
	})

	return &auth.ListPermissionsResp{
		Success: true,
		Resp: &auth.ListPermissionsResp_Data_{
			Data: &auth.ListPermissionsResp_Data{
				Permissions:  mapper.AsListPermission(permissions, actionMap, resourceMap),
				PageMetaData: utils.AsPageMetaData(normalizePageable, int64(len(permissions))),
			},
		},
	}, nil
}

// AssignPermissions is a function that assigns permissions to a user.
func (p permissionSvc) AssignPermissions(ctx context.Context, req *auth.AssignPermissionsReq) (*common.EmptyResp, error) {
	err := p.permissionValidator.ValidateAssignPermissions(ctx, req)
	if err != nil {
		log.Printf("error validating assign permissions: %v", err)
		return nil, err
	}

	err = p.userPermissionsRepo.CreateUserPermissions(ctx, req.UserId, req.PermissionIds)
	if err != nil {
		log.Printf("error creating user permissions: %v", err)
		return nil, err
	}

	return &common.EmptyResp{
		Success: true,
	}, nil
}

func NewPermissionGrpcService(
	actionRepo repository.IActionRepository,
	resourceRepo repository.IResourceRepository,
	permissionRepo repository.IPermissionRepository,
	userPermissionsRepo repository.IUserPermissionsRepository,
	authRepo repository.IAuthRepository,
	permissionValidator validator.IPermissionValidator) IPermissionService {
	return &permissionSvc{
		permissionRepo: permissionRepo,
		actionRepo:     actionRepo,
		resourceRepo:   resourceRepo,
		authRepo:       authRepo, userPermissionsRepo: userPermissionsRepo,
		permissionValidator: permissionValidator,
	}
}
