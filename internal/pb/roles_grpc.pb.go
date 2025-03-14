// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: internal/pb/roles.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RolesService_AssignRole_FullMethodName      = "/roles.RolesService/AssignRole"
	RolesService_RevokeRole_FullMethodName      = "/roles.RolesService/RevokeRole"
	RolesService_CheckPermission_FullMethodName = "/roles.RolesService/CheckPermission"
)

// RolesServiceClient is the client API for RolesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RolesServiceClient interface {
	AssignRole(ctx context.Context, in *AssignRoleRequest, opts ...grpc.CallOption) (*AssignRoleResponse, error)
	RevokeRole(ctx context.Context, in *RevokeRoleRequest, opts ...grpc.CallOption) (*RevokeRoleResponse, error)
	CheckPermission(ctx context.Context, in *CheckPermissionRequest, opts ...grpc.CallOption) (*CheckPermissionResponse, error)
}

type rolesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRolesServiceClient(cc grpc.ClientConnInterface) RolesServiceClient {
	return &rolesServiceClient{cc}
}

func (c *rolesServiceClient) AssignRole(ctx context.Context, in *AssignRoleRequest, opts ...grpc.CallOption) (*AssignRoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AssignRoleResponse)
	err := c.cc.Invoke(ctx, RolesService_AssignRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesServiceClient) RevokeRole(ctx context.Context, in *RevokeRoleRequest, opts ...grpc.CallOption) (*RevokeRoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RevokeRoleResponse)
	err := c.cc.Invoke(ctx, RolesService_RevokeRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesServiceClient) CheckPermission(ctx context.Context, in *CheckPermissionRequest, opts ...grpc.CallOption) (*CheckPermissionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckPermissionResponse)
	err := c.cc.Invoke(ctx, RolesService_CheckPermission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RolesServiceServer is the server API for RolesService service.
// All implementations must embed UnimplementedRolesServiceServer
// for forward compatibility.
type RolesServiceServer interface {
	AssignRole(context.Context, *AssignRoleRequest) (*AssignRoleResponse, error)
	RevokeRole(context.Context, *RevokeRoleRequest) (*RevokeRoleResponse, error)
	CheckPermission(context.Context, *CheckPermissionRequest) (*CheckPermissionResponse, error)
	mustEmbedUnimplementedRolesServiceServer()
}

// UnimplementedRolesServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRolesServiceServer struct{}

func (UnimplementedRolesServiceServer) AssignRole(context.Context, *AssignRoleRequest) (*AssignRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignRole not implemented")
}
func (UnimplementedRolesServiceServer) RevokeRole(context.Context, *RevokeRoleRequest) (*RevokeRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeRole not implemented")
}
func (UnimplementedRolesServiceServer) CheckPermission(context.Context, *CheckPermissionRequest) (*CheckPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPermission not implemented")
}
func (UnimplementedRolesServiceServer) mustEmbedUnimplementedRolesServiceServer() {}
func (UnimplementedRolesServiceServer) testEmbeddedByValue()                      {}

// UnsafeRolesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RolesServiceServer will
// result in compilation errors.
type UnsafeRolesServiceServer interface {
	mustEmbedUnimplementedRolesServiceServer()
}

func RegisterRolesServiceServer(s grpc.ServiceRegistrar, srv RolesServiceServer) {
	// If the following call pancis, it indicates UnimplementedRolesServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RolesService_ServiceDesc, srv)
}

func _RolesService_AssignRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServiceServer).AssignRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolesService_AssignRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServiceServer).AssignRole(ctx, req.(*AssignRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolesService_RevokeRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServiceServer).RevokeRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolesService_RevokeRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServiceServer).RevokeRole(ctx, req.(*RevokeRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolesService_CheckPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServiceServer).CheckPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolesService_CheckPermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServiceServer).CheckPermission(ctx, req.(*CheckPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RolesService_ServiceDesc is the grpc.ServiceDesc for RolesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RolesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "roles.RolesService",
	HandlerType: (*RolesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AssignRole",
			Handler:    _RolesService_AssignRole_Handler,
		},
		{
			MethodName: "RevokeRole",
			Handler:    _RolesService_RevokeRole_Handler,
		},
		{
			MethodName: "CheckPermission",
			Handler:    _RolesService_CheckPermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/pb/roles.proto",
}
