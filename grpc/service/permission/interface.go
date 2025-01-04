package grpcservicepermission

import (
	"context"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/common"
)

type IPermissionService interface {
	UpsertPermission(ctx context.Context, req *auth.UpsertPermissionReq) (*common.UpsertResp, error)
	ListPermissions(ctx context.Context, req *auth.ListPermissionsReq) (*auth.ListPermissionsResp, error)
	AssignPermissions(ctx context.Context, req *auth.AssignPermissionsReq) (*common.EmptyResp, error)
	PermissionOfUser(ctx context.Context, req *common.IdReq) (*auth.PermissionOfUserResp, error)
}
