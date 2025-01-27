package main

import (
	"github.com/ngdangkietswe/swe-auth-service/cache"
	"github.com/ngdangkietswe/swe-auth-service/data/repository"
	"github.com/ngdangkietswe/swe-auth-service/grpc"
	"github.com/ngdangkietswe/swe-auth-service/kafka"
	"github.com/ngdangkietswe/swe-auth-service/logger"
	"github.com/ngdangkietswe/swe-go-common-shared/config"
	"go.uber.org/fx"
	grpcserver "google.golang.org/grpc"
)

func main() {
	config.Init()
	app := fx.New(
		logger.Module,
		repository.Module,
		cache.Module,
		kafka.Module,
		grpc.Module,
		fx.Invoke(func(*grpcserver.Server) {}),
	)
	app.Run()
}
