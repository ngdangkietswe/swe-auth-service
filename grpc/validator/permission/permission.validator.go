package permission

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-auth-service/data/repository/action"
	authrepo "github.com/ngdangkietswe/swe-auth-service/data/repository/auth"
	permissionrepo "github.com/ngdangkietswe/swe-auth-service/data/repository/permission"
	"github.com/ngdangkietswe/swe-auth-service/data/repository/resource"
	commonutils "github.com/ngdangkietswe/swe-go-common-shared/util"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
)

type permissionValidator struct {
	actionRepo     action.IActionRepository
	resourceRepo   resource.IResourceRepository
	permissionRepo permissionrepo.IPermissionRepository
	authRepo       authrepo.IAuthRepository
}

// ValidateListPermissions is a function that validates the ListPermissionsReq.
func (p permissionValidator) ValidateListPermissions(req *auth.ListPermissionsReq) error {
	if req.ActionId != nil {
		if err := p.validateActionIdReq(nil, *req.ActionId); err != nil {
			return err
		}
	}

	if req.ResourceId != nil {
		if err := p.validateResourceIdReq(nil, *req.ResourceId); err != nil {
			return err
		}
	}

	return nil
}

// ValidateUpsertPermission is a function that validates the UpsertPermissionReq.
func (p permissionValidator) ValidateUpsertPermission(req *auth.UpsertPermissionReq) error {
	if req.Id != nil && *req.Id != "" {
		if _, err := uuid.Parse(*req.Id); err != nil {
			return errors.New("invalid permission id")
		}
	}

	if err := p.validateActionIdReq(nil, req.ActionId); err != nil {
		return err
	}

	if err := p.validateResourceIdReq(nil, req.ResourceId); err != nil {
		return err
	}

	exists, err := p.permissionRepo.ExistsByActionAndResource(nil, req.ActionId, req.ResourceId)
	if err != nil {
		return err
	} else if exists {
		return errors.New("permission already exists for the action and resource")
	}

	return nil
}

// validateActionIdReq is a function that validates the action id.
func (p permissionValidator) validateActionIdReq(ctx context.Context, actionId string) error {
	if actionId == "" {
		return errors.New("action id is required")
	} else if _, err := uuid.Parse(actionId); err != nil {
		return errors.New("invalid action id")
	} else if exists, err := p.actionRepo.ExistsById(ctx, actionId); err != nil {
		return err
	} else if !exists {
		return errors.New("action does not exist")
	}

	return nil
}

// validateResourceIdReq is a function that validates the resource id.
func (p permissionValidator) validateResourceIdReq(ctx context.Context, resourceId string) error {
	if resourceId == "" {
		return errors.New("resource id is required")
	} else if _, err := uuid.Parse(resourceId); err != nil {
		return errors.New("invalid resource id")
	} else if exists, err := p.resourceRepo.ExistsById(ctx, resourceId); err != nil {
		return err
	} else if !exists {
		return errors.New("resource does not exist")
	}

	return nil
}

// ValidateAssignPermissions is a function that validates the AssignPermissionsReq.
func (p permissionValidator) ValidateAssignPermissions(ctx context.Context, req *auth.AssignPermissionsReq) error {
	if req.UserId == "" {
		return errors.New("user id is required")
	} else if _, err := uuid.Parse(req.UserId); err != nil {
		return errors.New("invalid user id")
	}

	exists, err := p.authRepo.ExistsById(ctx, req.UserId)
	if err != nil {
		return err
	} else if !exists {
		return errors.New("user does not exist")
	}

	if len(req.PermissionIds) == 0 {
		return errors.New("permission ids are required")
	} else if hasInvalid := commonutils.HasAnyInvalidUUID(req.PermissionIds); hasInvalid == true {
		return errors.New("invalid permission ids")
	}

	exists, err = p.permissionRepo.ExistsByIdIn(ctx, req.PermissionIds)
	if err != nil {
		return err
	} else if !exists {
		return errors.New("one or more permissions do not exist")
	}

	return nil
}

func NewPermissionValidator(
	actionRepo action.IActionRepository,
	resourceRepo resource.IResourceRepository,
	permissionRepo permissionrepo.IPermissionRepository,
	authRepo authrepo.IAuthRepository) IPermissionValidator {
	return &permissionValidator{
		actionRepo:     actionRepo,
		resourceRepo:   resourceRepo,
		permissionRepo: permissionRepo,
		authRepo:       authRepo,
	}
}
