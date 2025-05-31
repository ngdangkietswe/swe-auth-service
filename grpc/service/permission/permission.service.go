package grpcservicepermission

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
	"github.com/ngdangkietswe/swe-auth-service/data/repository"
	actionrepo "github.com/ngdangkietswe/swe-auth-service/data/repository/action"
	authrepo "github.com/ngdangkietswe/swe-auth-service/data/repository/auth"
	permissionrepo "github.com/ngdangkietswe/swe-auth-service/data/repository/permission"
	resourcerepo "github.com/ngdangkietswe/swe-auth-service/data/repository/resource"
	userpermissionsrepo "github.com/ngdangkietswe/swe-auth-service/data/repository/user_permissions"
	"github.com/ngdangkietswe/swe-auth-service/grpc/mapper"
	"github.com/ngdangkietswe/swe-auth-service/grpc/utils"
	validator "github.com/ngdangkietswe/swe-auth-service/grpc/validator/permission"
	"github.com/ngdangkietswe/swe-go-common-shared/cache"
	"github.com/ngdangkietswe/swe-go-common-shared/constants"
	grpcutil "github.com/ngdangkietswe/swe-go-common-shared/grpc/util"
	"github.com/ngdangkietswe/swe-go-common-shared/logger"
	commonutil "github.com/ngdangkietswe/swe-go-common-shared/util"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/common"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"log"
)

type permissionSvc struct {
	client     *ent.Client
	logger     *logger.Logger
	redisCache *cache.RedisCache

	actionRepo          actionrepo.IActionRepository
	resourceRepo        resourcerepo.IResourceRepository
	permissionRepo      permissionrepo.IPermissionRepository
	userPermissionsRepo userpermissionsrepo.IUserPermissionsRepository
	authRepo            authrepo.IAuthRepository

	permissionValidator validator.IPermissionValidator
}

// PermissionOfUser is a function that gets permissions of a user.
func (p permissionSvc) PermissionOfUser(ctx context.Context, req *common.IdReq) (*auth.PermissionOfUserResp, error) {
	var userId string
	if req.Id != "" {
		userId = req.Id
	} else {
		userId = grpcutil.GetGrpcPrincipal(ctx).UserId
	}

	userUid, err := uuid.Parse(userId)
	if err != nil {
		log.Printf("error parsing user id: %v", err)
		return nil, err
	}

	userPermissions, err := p.userPermissionsRepo.FindAllByUserId(ctx, userUid)
	if err != nil {
		log.Printf("error finding user permissions: %v", err)
		return nil, err
	}

	permissionIds := lo.Map(userPermissions, func(userPermission *ent.UsersPermission, _ int) uuid.UUID {
		return userPermission.PermissionID
	})

	permissions, err := p.permissionRepo.FindAllByIdIn(ctx, lo.Uniq(permissionIds))
	if err != nil {
		log.Printf("error finding permissions by ids: %v", err)
		return nil, err
	}

	actionMap, err := p.getActionMap(ctx, permissions)
	if err != nil {
		return nil, err
	}

	resourceMap, err := p.getResourceMap(ctx, permissions)
	if err != nil {
		return nil, err
	}

	return &auth.PermissionOfUserResp{
		Success: true,
		Resp: &auth.PermissionOfUserResp_Data_{
			Data: &auth.PermissionOfUserResp_Data{
				Permissions: mapper.AsListPermission(permissions, actionMap, resourceMap),
			},
		},
	}, nil
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

	permission, err := repository.WithTxResult(ctx, p.client, p.logger, func(tx *ent.Tx) (*ent.Permission, error) {
		return p.permissionRepo.Save(ctx, tx, req)
	})
	if err != nil {
		p.logger.Error("error upserting permission", zap.Error(err))
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
	permissions, count, err := p.permissionRepo.FindAll(ctx, req, normalizePageable)
	if err != nil {
		log.Printf("error listing permissions: %v", err)
		return nil, err
	}

	actionMap, err := p.getActionMap(ctx, permissions)
	if err != nil {
		return nil, err
	}

	resourceMap, err := p.getResourceMap(ctx, permissions)
	if err != nil {
		return nil, err
	}

	return &auth.ListPermissionsResp{
		Success: true,
		Resp: &auth.ListPermissionsResp_Data_{
			Data: &auth.ListPermissionsResp_Data{
				Permissions:  mapper.AsListPermission(permissions, actionMap, resourceMap),
				PageMetaData: utils.AsPageMetaData(normalizePageable, count),
			},
		},
	}, nil
}

// getActionMap is a function that gets an action map.
func (p permissionSvc) getActionMap(ctx context.Context, permissions []*ent.Permission) (map[uuid.UUID]*ent.Action, error) {
	actionIds := lo.Map(permissions, func(permission *ent.Permission, _ int) uuid.UUID {
		return permission.ActionID
	})
	actions, err := p.actionRepo.FindAllByIds(ctx, lo.Uniq(actionIds))
	if err != nil {
		log.Printf("error finding actions by ids: %v", err)
		return nil, err
	}
	return commonutil.Convert2Map(actions, func(action *ent.Action) uuid.UUID {
		return action.ID
	}), nil
}

// getResourceMap is a function that gets a resource map.
func (p permissionSvc) getResourceMap(ctx context.Context, permissions []*ent.Permission) (map[uuid.UUID]*ent.Resource, error) {
	resourceIds := lo.Map(permissions, func(permission *ent.Permission, _ int) uuid.UUID {
		return permission.ResourceID
	})
	resources, err := p.resourceRepo.FindAllByIds(ctx, lo.Uniq(resourceIds))
	if err != nil {
		log.Printf("error finding resources by ids: %v", err)
		return nil, err
	}
	return commonutil.Convert2Map(resources, func(resource *ent.Resource) uuid.UUID {
		return resource.ID
	}), nil
}

// AssignPermissions is a function that assigns permissions to a user.
func (p permissionSvc) AssignPermissions(ctx context.Context, req *auth.AssignPermissionsReq) (*common.EmptyResp, error) {
	if err := p.permissionValidator.ValidateAssignPermissions(ctx, req); err != nil {
		log.Printf("error validating assign permissions: %v", err)
		return nil, err
	}

	// Delete old permissions of the user
	if err := p.userPermissionsRepo.DeleteAllByUserId(ctx, req.UserId); err != nil {
		log.Printf("error deleting user permissions: %v", err)
		return nil, err
	}

	if err := p.userPermissionsRepo.CreateUserPermissions(ctx, req.UserId, req.PermissionIds); err != nil {
		log.Printf("error creating user permissions: %v", err)
		return nil, err
	}

	// Delete user permission cache after assigning permissions
	log.Printf("deleting user permission cache after assigning permissions for user %s", req.UserId)
	cacheKey := fmt.Sprintf("%s_%s", constants.UserPermissionCacheKeyPrefix, req.UserId)
	if err := p.redisCache.Delete(cacheKey); err != nil {
		log.Printf("error deleting user permission cache: %v", err)
		return nil, err
	}

	return &common.EmptyResp{
		Success: true,
	}, nil
}

func NewPermissionGrpcService(
	client *ent.Client,
	logger *logger.Logger,
	redisCache *cache.RedisCache,
	actionRepo actionrepo.IActionRepository,
	resourceRepo resourcerepo.IResourceRepository,
	permissionRepo permissionrepo.IPermissionRepository,
	userPermissionsRepo userpermissionsrepo.IUserPermissionsRepository,
	authRepo authrepo.IAuthRepository,
	permissionValidator validator.IPermissionValidator,
) IPermissionService {
	return &permissionSvc{
		client:         client,
		logger:         logger,
		redisCache:     redisCache,
		permissionRepo: permissionRepo,
		actionRepo:     actionRepo,
		resourceRepo:   resourceRepo,
		authRepo:       authRepo, userPermissionsRepo: userPermissionsRepo,
		permissionValidator: permissionValidator,
	}
}
