package server

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/ozonmp/week-5-workshop/product-service/internal/pkg/logger"
	desc "github.com/ozonmp/week-5-workshop/product-service/pkg/product-service"
)

func createGatewayServer(grpcAddr, gatewayAddr string, allowedOrigins []string) *http.Server {
	ctx := context.Background()

	// Create a client connection to the gRPC Server we just started.
	// This is where the gRPC-Gateway proxies the requests.
	conn, err := grpc.DialContext(
		ctx,
		grpcAddr,
		grpc.WithInsecure(),
	)
	if err != nil {
		logger.FatalKV(ctx, "Failed to dial server", "err", err)
	}

	mux := runtime.NewServeMux()
	if err := desc.RegisterProductServiceHandler(ctx, mux, conn); err != nil {
		logger.FatalKV(ctx, "Failed registration handler", "err", err)
	}

	gatewayServer := &http.Server{
		Addr:    gatewayAddr,
		Handler: cors(mux, allowedOrigins),
	}

	return gatewayServer
}

func cors(h http.Handler, allowedOrigins []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		providedOrigin := r.Header.Get("Origin")
		matches := false
		for _, allowedOrigin := range allowedOrigins {
			if providedOrigin == allowedOrigin {
				matches = true
				break
			}
		}

		if matches {
			w.Header().Set("Access-Control-Allow-Origin", providedOrigin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType, grpc-metadata-log-level")
		}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}
