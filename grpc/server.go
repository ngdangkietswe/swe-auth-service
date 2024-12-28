package grpc

import (
	"fmt"
	"github.com/ngdangkietswe/swe-auth-service/data/datasource"
	"github.com/ngdangkietswe/swe-auth-service/grpc/server/auth"
	"github.com/ngdangkietswe/swe-auth-service/grpc/server/permission"
	"github.com/ngdangkietswe/swe-go-common-shared/config"
	"github.com/ngdangkietswe/swe-go-common-shared/middleware"
	"google.golang.org/grpc"
	"log"
	"net"
)

// NewGrpcServer function is used to create a new gRPC server. It listens on the gRPC port and serves the gRPC server.
func NewGrpcServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GetInt("GRPC_PORT", 7020)))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	entClient := datasource.NewEntClient()
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.AuthMiddleware),
	)

	auth.NewGrpcHandler(entClient).RegisterGrpcServer(grpcServer)
	permission.NewGrpcHandler(entClient).RegisterGrpcServer(grpcServer)

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
