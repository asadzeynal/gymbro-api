// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: service_gymbro.proto

package pb

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

// GymBroClient is the client API for GymBro service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GymBroClient interface {
	CreateExercise(ctx context.Context, in *CreateExerciseRequest, opts ...grpc.CallOption) (*CreateExerciseResponse, error)
	GetExercisesByInitial(ctx context.Context, in *GetExercisesByInitialRequest, opts ...grpc.CallOption) (*GetExercisesByInitialResponse, error)
}

type gymBroClient struct {
	cc grpc.ClientConnInterface
}

func NewGymBroClient(cc grpc.ClientConnInterface) GymBroClient {
	return &gymBroClient{cc}
}

func (c *gymBroClient) CreateExercise(ctx context.Context, in *CreateExerciseRequest, opts ...grpc.CallOption) (*CreateExerciseResponse, error) {
	out := new(CreateExerciseResponse)
	err := c.cc.Invoke(ctx, "/pb.GymBro/CreateExercise", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gymBroClient) GetExercisesByInitial(ctx context.Context, in *GetExercisesByInitialRequest, opts ...grpc.CallOption) (*GetExercisesByInitialResponse, error) {
	out := new(GetExercisesByInitialResponse)
	err := c.cc.Invoke(ctx, "/pb.GymBro/GetExercisesByInitial", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GymBroServer is the server API for GymBro service.
// All implementations must embed UnimplementedGymBroServer
// for forward compatibility
type GymBroServer interface {
	CreateExercise(context.Context, *CreateExerciseRequest) (*CreateExerciseResponse, error)
	GetExercisesByInitial(context.Context, *GetExercisesByInitialRequest) (*GetExercisesByInitialResponse, error)
	mustEmbedUnimplementedGymBroServer()
}

// UnimplementedGymBroServer must be embedded to have forward compatible implementations.
type UnimplementedGymBroServer struct {
}

func (UnimplementedGymBroServer) CreateExercise(context.Context, *CreateExerciseRequest) (*CreateExerciseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExercise not implemented")
}
func (UnimplementedGymBroServer) GetExercisesByInitial(context.Context, *GetExercisesByInitialRequest) (*GetExercisesByInitialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExercisesByInitial not implemented")
}
func (UnimplementedGymBroServer) mustEmbedUnimplementedGymBroServer() {}

// UnsafeGymBroServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GymBroServer will
// result in compilation errors.
type UnsafeGymBroServer interface {
	mustEmbedUnimplementedGymBroServer()
}

func RegisterGymBroServer(s grpc.ServiceRegistrar, srv GymBroServer) {
	s.RegisterService(&GymBro_ServiceDesc, srv)
}

func _GymBro_CreateExercise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExerciseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GymBroServer).CreateExercise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GymBro/CreateExercise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GymBroServer).CreateExercise(ctx, req.(*CreateExerciseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GymBro_GetExercisesByInitial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExercisesByInitialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GymBroServer).GetExercisesByInitial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GymBro/GetExercisesByInitial",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GymBroServer).GetExercisesByInitial(ctx, req.(*GetExercisesByInitialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GymBro_ServiceDesc is the grpc.ServiceDesc for GymBro service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GymBro_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.GymBro",
	HandlerType: (*GymBroServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateExercise",
			Handler:    _GymBro_CreateExercise_Handler,
		},
		{
			MethodName: "GetExercisesByInitial",
			Handler:    _GymBro_GetExercisesByInitial_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_gymbro.proto",
}
