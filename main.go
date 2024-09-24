package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"app/src/config"
	"app/src/lib/logger"
	authSrv "app/src/services/auth"
	"app/src/shared/exceptions"
	middleware "app/src/shared/middlewares"
)

var log = logger.NewLogger("Main")

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(
		runtime.WithErrorHandler(exceptions.ExceptionHandler),
	)

	authSrv.RegisterHandlerFromEndPoint(ctx, mux)

	logMw := middleware.LoggingMiddleware(mux)

	return http.ListenAndServe(fmt.Sprintf(":%v", config.AppConfiguration.AppPort), logMw)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	log.Printf("Starting gRPC-Gateway for multiple services on http://localhost:%v", config.AppConfiguration.AppPort)
	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
