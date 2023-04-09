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

// ExerciseServiceClient is the client API for ExerciseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExerciseServiceClient interface {
	CreateExercise(ctx context.Context, in *CreateExerciseRequest, opts ...grpc.CallOption) (*CreateExerciseResponse, error)
	GetExercises(ctx context.Context, in *GetExercisesRequest, opts ...grpc.CallOption) (*GetExercisesResponse, error)
}

type exerciseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExerciseServiceClient(cc grpc.ClientConnInterface) ExerciseServiceClient {
	return &exerciseServiceClient{cc}
}

func (c *exerciseServiceClient) CreateExercise(ctx context.Context, in *CreateExerciseRequest, opts ...grpc.CallOption) (*CreateExerciseResponse, error) {
	out := new(CreateExerciseResponse)
	err := c.cc.Invoke(ctx, "/pb.ExerciseService/CreateExercise", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exerciseServiceClient) GetExercises(ctx context.Context, in *GetExercisesRequest, opts ...grpc.CallOption) (*GetExercisesResponse, error) {
	out := new(GetExercisesResponse)
	err := c.cc.Invoke(ctx, "/pb.ExerciseService/GetExercises", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExerciseServiceServer is the server API for ExerciseService service.
// All implementations must embed UnimplementedExerciseServiceServer
// for forward compatibility
type ExerciseServiceServer interface {
	CreateExercise(context.Context, *CreateExerciseRequest) (*CreateExerciseResponse, error)
	GetExercises(context.Context, *GetExercisesRequest) (*GetExercisesResponse, error)
	mustEmbedUnimplementedExerciseServiceServer()
}

// UnimplementedExerciseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExerciseServiceServer struct {
}

func (UnimplementedExerciseServiceServer) CreateExercise(context.Context, *CreateExerciseRequest) (*CreateExerciseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExercise not implemented")
}
func (UnimplementedExerciseServiceServer) GetExercises(context.Context, *GetExercisesRequest) (*GetExercisesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExercises not implemented")
}
func (UnimplementedExerciseServiceServer) mustEmbedUnimplementedExerciseServiceServer() {}

// UnsafeExerciseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExerciseServiceServer will
// result in compilation errors.
type UnsafeExerciseServiceServer interface {
	mustEmbedUnimplementedExerciseServiceServer()
}

func RegisterExerciseServiceServer(s grpc.ServiceRegistrar, srv ExerciseServiceServer) {
	s.RegisterService(&ExerciseService_ServiceDesc, srv)
}

func _ExerciseService_CreateExercise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExerciseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExerciseServiceServer).CreateExercise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ExerciseService/CreateExercise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExerciseServiceServer).CreateExercise(ctx, req.(*CreateExerciseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExerciseService_GetExercises_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExercisesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExerciseServiceServer).GetExercises(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ExerciseService/GetExercises",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExerciseServiceServer).GetExercises(ctx, req.(*GetExercisesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExerciseService_ServiceDesc is the grpc.ServiceDesc for ExerciseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExerciseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ExerciseService",
	HandlerType: (*ExerciseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateExercise",
			Handler:    _ExerciseService_CreateExercise_Handler,
		},
		{
			MethodName: "GetExercises",
			Handler:    _ExerciseService_GetExercises_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_gymbro.proto",
}
