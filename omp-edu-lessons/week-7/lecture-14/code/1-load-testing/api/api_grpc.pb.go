// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TestAPIClient is the client API for TestAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestAPIClient interface {
	TestAPIHandler(ctx context.Context, in *TestAPIHandlerRequest, opts ...grpc.CallOption) (*TestAPIHandlerResponse, error)
}

type testAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewTestAPIClient(cc grpc.ClientConnInterface) TestAPIClient {
	return &testAPIClient{cc}
}

func (c *testAPIClient) TestAPIHandler(ctx context.Context, in *TestAPIHandlerRequest, opts ...grpc.CallOption) (*TestAPIHandlerResponse, error) {
	out := new(TestAPIHandlerResponse)
	err := c.cc.Invoke(ctx, "/api.TestAPI/TestAPIHandler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestAPIServer is the server API for TestAPI service.
// All implementations must embed UnimplementedTestAPIServer
// for forward compatibility
type TestAPIServer interface {
	TestAPIHandler(context.Context, *TestAPIHandlerRequest) (*TestAPIHandlerResponse, error)
	mustEmbedUnimplementedTestAPIServer()
}

// UnimplementedTestAPIServer must be embedded to have forward compatible implementations.
type UnimplementedTestAPIServer struct {
}

func (UnimplementedTestAPIServer) TestAPIHandler(context.Context, *TestAPIHandlerRequest) (*TestAPIHandlerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestAPIHandler not implemented")
}
func (UnimplementedTestAPIServer) mustEmbedUnimplementedTestAPIServer() {}

// UnsafeTestAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestAPIServer will
// result in compilation errors.
type UnsafeTestAPIServer interface {
	mustEmbedUnimplementedTestAPIServer()
}

func RegisterTestAPIServer(s grpc.ServiceRegistrar, srv TestAPIServer) {
	s.RegisterService(&TestAPI_ServiceDesc, srv)
}

func _TestAPI_TestAPIHandler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestAPIHandlerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestAPIServer).TestAPIHandler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TestAPI/TestAPIHandler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestAPIServer).TestAPIHandler(ctx, req.(*TestAPIHandlerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TestAPI_ServiceDesc is the grpc.ServiceDesc for TestAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.TestAPI",
	HandlerType: (*TestAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestAPIHandler",
			Handler:    _TestAPI_TestAPIHandler_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}