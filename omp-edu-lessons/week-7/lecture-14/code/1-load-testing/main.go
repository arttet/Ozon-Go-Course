package main

import (
	"context"
	"log"
	"net"

	desc "github.com/ozonmp/omp-edu-lessons/week-6/lecture-14/1-load-testing/api"
	"google.golang.org/grpc"
)

type server struct {
	desc.UnimplementedTestAPIServer
}

func (s *server) TestAPIHandler(ctx context.Context, req *desc.TestAPIHandlerRequest) (*desc.TestAPIHandlerResponse, error)  {
	return &desc.TestAPIHandlerResponse{
		Val: req.Val + 1,
	}, nil
}

func initServer() *server {
	return &server{}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	desc.RegisterTestAPIServer(grpcServer, initServer())
	grpcServer.Serve(lis)
}