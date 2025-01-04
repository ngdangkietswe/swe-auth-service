package grpc

import (
	"github.com/ngdangkietswe/swe-auth-service/grpc/server"
	"github.com/ngdangkietswe/swe-auth-service/grpc/service"
	"github.com/ngdangkietswe/swe-auth-service/grpc/validator"
	"go.uber.org/fx"
)

var Module = fx.Options(
	validator.Module,
	service.Module,
	server.Module,
)
