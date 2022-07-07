// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: api/dict/v1/dict.proto

package v1

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

// DictClient is the client API for Dict service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DictClient interface {
	Get(ctx context.Context, in *DictRequest, opts ...grpc.CallOption) (*DictResponse, error)
}

type dictClient struct {
	cc grpc.ClientConnInterface
}

func NewDictClient(cc grpc.ClientConnInterface) DictClient {
	return &dictClient{cc}
}

func (c *dictClient) Get(ctx context.Context, in *DictRequest, opts ...grpc.CallOption) (*DictResponse, error) {
	out := new(DictResponse)
	err := c.cc.Invoke(ctx, "/api.dict.v1.Dict/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DictServer is the server API for Dict service.
// All implementations must embed UnimplementedDictServer
// for forward compatibility
type DictServer interface {
	Get(context.Context, *DictRequest) (*DictResponse, error)
	mustEmbedUnimplementedDictServer()
}

// UnimplementedDictServer must be embedded to have forward compatible implementations.
type UnimplementedDictServer struct {
}

func (UnimplementedDictServer) Get(context.Context, *DictRequest) (*DictResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDictServer) mustEmbedUnimplementedDictServer() {}

// UnsafeDictServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DictServer will
// result in compilation errors.
type UnsafeDictServer interface {
	mustEmbedUnimplementedDictServer()
}

func RegisterDictServer(s grpc.ServiceRegistrar, srv DictServer) {
	s.RegisterService(&Dict_ServiceDesc, srv)
}

func _Dict_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.dict.v1.Dict/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServer).Get(ctx, req.(*DictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Dict_ServiceDesc is the grpc.ServiceDesc for Dict service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dict_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.dict.v1.Dict",
	HandlerType: (*DictServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Dict_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/dict/v1/dict.proto",
}
