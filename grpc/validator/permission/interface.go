package permission

import (
	"context"
	"github.com/google/wire"
	permissionrepo "github.com/ngdangkietswe/swe-auth-service/data/repository/permission"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
)

type IPermissionValidator interface {
	ValidateUpsertPermission(req *auth.UpsertPermissionReq) error
	ValidateListPermissions(req *auth.ListPermissionsReq) error
	ValidateAssignPermissions(ctx context.Context, req *auth.AssignPermissionsReq) error
}

// ProvidePermissionValidator is a function to provide a permission validator
func ProvidePermissionValidator() (p IPermissionValidator) {
	wire.Build(
		permissionrepo.ProvidePermissionRepository(),
		NewPermissionValidator,
	)
	return
}
