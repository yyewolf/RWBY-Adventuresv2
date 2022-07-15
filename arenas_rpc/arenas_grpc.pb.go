// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: arenas.proto

package arenapc

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

// ArenaClient is the client API for Arena service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArenaClient interface {
	CreateArena(ctx context.Context, in *CreateArenaReq, opts ...grpc.CallOption) (*CreateArenaRep, error)
	Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingRep, error)
}

type arenaClient struct {
	cc grpc.ClientConnInterface
}

func NewArenaClient(cc grpc.ClientConnInterface) ArenaClient {
	return &arenaClient{cc}
}

func (c *arenaClient) CreateArena(ctx context.Context, in *CreateArenaReq, opts ...grpc.CallOption) (*CreateArenaRep, error) {
	out := new(CreateArenaRep)
	err := c.cc.Invoke(ctx, "/arenapc.Arena/CreateArena", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arenaClient) Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingRep, error) {
	out := new(PingRep)
	err := c.cc.Invoke(ctx, "/arenapc.Arena/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArenaServer is the server API for Arena service.
// All implementations must embed UnimplementedArenaServer
// for forward compatibility
type ArenaServer interface {
	CreateArena(context.Context, *CreateArenaReq) (*CreateArenaRep, error)
	Ping(context.Context, *PingReq) (*PingRep, error)
	mustEmbedUnimplementedArenaServer()
}

// UnimplementedArenaServer must be embedded to have forward compatible implementations.
type UnimplementedArenaServer struct {
}

func (UnimplementedArenaServer) CreateArena(context.Context, *CreateArenaReq) (*CreateArenaRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArena not implemented")
}
func (UnimplementedArenaServer) Ping(context.Context, *PingReq) (*PingRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedArenaServer) mustEmbedUnimplementedArenaServer() {}

// UnsafeArenaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArenaServer will
// result in compilation errors.
type UnsafeArenaServer interface {
	mustEmbedUnimplementedArenaServer()
}

func RegisterArenaServer(s grpc.ServiceRegistrar, srv ArenaServer) {
	s.RegisterService(&Arena_ServiceDesc, srv)
}

func _Arena_CreateArena_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArenaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArenaServer).CreateArena(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arenapc.Arena/CreateArena",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArenaServer).CreateArena(ctx, req.(*CreateArenaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Arena_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArenaServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arenapc.Arena/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArenaServer).Ping(ctx, req.(*PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Arena_ServiceDesc is the grpc.ServiceDesc for Arena service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Arena_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "arenapc.Arena",
	HandlerType: (*ArenaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateArena",
			Handler:    _Arena_CreateArena_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Arena_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "arenas.proto",
}