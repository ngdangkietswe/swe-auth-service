package grpc

import (
	"fmt"
	"github.com/ngdangkietswe/swe-auth-service/configs"
	"github.com/ngdangkietswe/swe-auth-service/data/datasource"
	"github.com/ngdangkietswe/swe-auth-service/grpc/middleware"
	"github.com/ngdangkietswe/swe-auth-service/grpc/server/auth"
	"google.golang.org/grpc"
	"log"
	"net"
)

// NewGrpcServer function is used to create a new gRPC server. It listens on the gRPC port and serves the gRPC server.
func NewGrpcServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", configs.GlobalConfig.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	entClient := datasource.NewEntClient()
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.AuthMiddleware),
	)
	auth.NewGrpcHandler(entClient).RegisterGrpcServer(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
