package mwserver

import (
	"context"
	"fmt"
	"strings"

	"github.com/ozonmp/week-5-workshop/category-service/internal/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const authMetadataKey = "x-app-name"

// GRPCUnauthenticatedRequest проверяет что в запросе есть аутентификационный metadata параметр.
func GRPCUnauthenticatedRequest(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "no metadata")
	}

	for key, vals := range md {
		if isAuthMD(key, vals) {
			return handler(ctx, req)
		}
	}

	errMsg := fmt.Sprintf(`got empty %q metadata`, authMetadataKey)
	logger.ErrorKV(ctx, errMsg)
	return nil, status.Error(codes.InvalidArgument, errMsg)
}

func isAuthMD(key string, vals []string) bool {
	if strings.ToLower(key) != authMetadataKey {
		return false
	}
	if len(vals) == 0 { // Нет значений
		return false
	}
	if vals[0] == "" { // Смотрим только на первое значение
		return false
	}
	return true
}
