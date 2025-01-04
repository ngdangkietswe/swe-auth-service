package validator

import (
	"github.com/ngdangkietswe/swe-auth-service/grpc/validator/auth"
	"github.com/ngdangkietswe/swe-auth-service/grpc/validator/permission"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		auth.NewAuthValidator,
		permission.NewPermissionValidator,
	),
)
