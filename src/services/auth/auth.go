package auth

import (
	authPb "app/proto/auth"
	"app/src/config"
	"context"
	"flag"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RegisterHandlerFromEndPoint(ctx context.Context, mux *runtime.ServeMux) error {

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	grpcServerEndpoint := flag.String("grpc-server-endpoint", config.AppConfiguration.AuthServiceUri, "gRPC server endpoint")

	err := authPb.RegisterAuthControllerHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	return err
}
