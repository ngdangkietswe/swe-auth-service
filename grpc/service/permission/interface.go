package permission

import (
	"context"
	"github.com/google/wire"
	"github.com/ngdangkietswe/swe-auth-service/cache"
	"github.com/ngdangkietswe/swe-auth-service/data/repository/action"
	"github.com/ngdangkietswe/swe-auth-service/data/repository/permission"
	"github.com/ngdangkietswe/swe-auth-service/data/repository/resource"
	"github.com/ngdangkietswe/swe-auth-service/data/repository/user_permissions"
	validator "github.com/ngdangkietswe/swe-auth-service/grpc/validator/permission"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/common"
)

type IPermissionService interface {
	UpsertPermission(ctx context.Context, req *auth.UpsertPermissionReq) (*common.UpsertResp, error)
	ListPermissions(ctx context.Context, req *auth.ListPermissionsReq) (*auth.ListPermissionsResp, error)
	AssignPermissions(ctx context.Context, req *auth.AssignPermissionsReq) (*common.EmptyResp, error)
	PermissionOfUser(ctx context.Context, req *common.IdReq) (*auth.PermissionOfUserResp, error)
}

// ProvidePermissionService is a function to provide a permission service
func ProvidePermissionService() (p IPermissionService) {
	wire.Build(
		action.ProvideActionRepository(),
		resource.ProvideResourceRepository(),
		permission.ProvidePermissionRepository(),
		user_permissions.ProvideUserPermissionsRepository(),
		validator.ProvidePermissionValidator,
		cache.ProvideRedisCache(),
		NewPermissionGrpcService,
	)
	return
}
