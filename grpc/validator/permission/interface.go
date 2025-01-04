package permission

import (
	"context"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
)

type IPermissionValidator interface {
	ValidateUpsertPermission(req *auth.UpsertPermissionReq) error
	ValidateListPermissions(req *auth.ListPermissionsReq) error
	ValidateAssignPermissions(ctx context.Context, req *auth.AssignPermissionsReq) error
}
