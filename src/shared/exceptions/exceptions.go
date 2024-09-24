package exceptions

import (
	"app/src/lib/logger"
	"context"
	"encoding/json"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/status"
)

type CustomErrorResponse struct {
	Code    int32               `json:"code"`
	Message string              `json:"message"`
	Details []map[string]string `json:"details,omitempty"`
}

func ExceptionHandler(_ context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	grpcStatus, _ := status.FromError(err)
	log := logger.NewLogger("ExceptionHandler")

	var details []map[string]string
	log.Error(grpcStatus)

	customErr := CustomErrorResponse{
		Code:    int32(grpcStatus.Code()),
		Message: grpcStatus.Message(),
		Details: details,
	}

	w.WriteHeader(runtime.HTTPStatusFromCode(grpcStatus.Code()))
	w.Header().Set("Content-Type", marshaler.ContentType(customErr))

	log.Error(customErr.Details)
	jsonError, _ := json.Marshal(customErr)
	w.Write(jsonError)
}
