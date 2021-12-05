package mwclient

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/ozonmp/week-5-workshop/product-service/internal/config"
)

const (
	appNameHeader    = "x-app-name"
	appVersionHeader = "x-app-version"
)

// AddAppInfoUnary добавляет в единичные запросы информацию о клиенте.
func AddAppInfoUnary(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	ctx = metadata.AppendToOutgoingContext(ctx, appNameHeader, config.AppName)
	ctx = metadata.AppendToOutgoingContext(ctx, appVersionHeader, config.AppVersion)
	return invoker(ctx, method, req, reply, cc, opts...)
}
