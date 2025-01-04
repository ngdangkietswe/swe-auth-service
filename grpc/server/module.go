package server

import "go.uber.org/fx"

var Module = fx.Provide(
	NewAuthGrpcServer,
	NewPermissionGrpcServer,
	NewPermissionInternalGrpcServer,
	NewGrpcServer,
)
