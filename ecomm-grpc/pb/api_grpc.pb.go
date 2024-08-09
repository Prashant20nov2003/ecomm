// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: api.proto

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

// EcommClient is the client API for Ecomm service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EcommClient interface {
	CreateProduct(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductRes, error)
	GetProduct(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductRes, error)
	ListProducts(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ListProductRes, error)
	UpdateProduct(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductRes, error)
	DeleteProduct(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductRes, error)
	CreateOrder(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderRes, error)
	GetOrder(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderRes, error)
	ListOrders(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*ListOrderRes, error)
	DeleteOrder(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderRes, error)
	CreateUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRes, error)
	GetUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRes, error)
	ListUsers(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*ListUserRes, error)
	UpdateUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRes, error)
	DeleteUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRes, error)
	CreateSession(ctx context.Context, in *SessionReq, opts ...grpc.CallOption) (*SessionRes, error)
	GetSession(ctx context.Context, in *SessionReq, opts ...grpc.CallOption) (*SessionRes, error)
	RevokeSession(ctx context.Context, in *SessionReq, opts ...grpc.CallOption) (*SessionRes, error)
	DeleteSession(ctx context.Context, in *SessionReq, opts ...grpc.CallOption) (*SessionRes, error)
}

type ecommClient struct {
	cc grpc.ClientConnInterface
}

func NewEcommClient(cc grpc.ClientConnInterface) EcommClient {
	return &ecommClient{cc}
}

func (c *ecommClient) CreateProduct(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductRes, error) {
	out := new(ProductRes)
	err := c.cc.Invoke(ctx, "/pb.ecomm/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) GetProduct(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductRes, error) {
	out := new(ProductRes)
	err := c.cc.Invoke(ctx, "/pb.ecomm/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) ListProducts(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ListProductRes, error) {
	out := new(ListProductRes)
	err := c.cc.Invoke(ctx, "/pb.ecomm/ListProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) UpdateProduct(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductRes, error) {
	out := new(ProductRes)
	err := c.cc.Invoke(ctx, "/pb.ecomm/UpdateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) DeleteProduct(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductRes, error) {
	out := new(ProductRes)
	err := c.cc.Invoke(ctx, "/pb.ecomm/DeleteProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) CreateOrder(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderRes, error) {
	out := new(OrderRes)
	err := c.cc.Invoke(ctx, "/pb.ecomm/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) GetOrder(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderRes, error) {
	out := new(OrderRes)
	err := c.cc.Invoke(ctx, "/pb.ecomm/GetOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) ListOrders(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*ListOrderRes, error) {
	out := new(ListOrderRes)
	err := c.cc.Invoke(ctx, "/pb.ecomm/ListOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) DeleteOrder(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderRes, error) {
	out := new(OrderRes)
	err := c.cc.Invoke(ctx, "/pb.ecomm/DeleteOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) CreateUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRes, error) {
	out := new(UserRes)
	err := c.cc.Invoke(ctx, "/pb.ecomm/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) GetUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRes, error) {
	out := new(UserRes)
	err := c.cc.Invoke(ctx, "/pb.ecomm/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) ListUsers(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*ListUserRes, error) {
	out := new(ListUserRes)
	err := c.cc.Invoke(ctx, "/pb.ecomm/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) UpdateUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRes, error) {
	out := new(UserRes)
	err := c.cc.Invoke(ctx, "/pb.ecomm/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) DeleteUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRes, error) {
	out := new(UserRes)
	err := c.cc.Invoke(ctx, "/pb.ecomm/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) CreateSession(ctx context.Context, in *SessionReq, opts ...grpc.CallOption) (*SessionRes, error) {
	out := new(SessionRes)
	err := c.cc.Invoke(ctx, "/pb.ecomm/CreateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) GetSession(ctx context.Context, in *SessionReq, opts ...grpc.CallOption) (*SessionRes, error) {
	out := new(SessionRes)
	err := c.cc.Invoke(ctx, "/pb.ecomm/GetSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) RevokeSession(ctx context.Context, in *SessionReq, opts ...grpc.CallOption) (*SessionRes, error) {
	out := new(SessionRes)
	err := c.cc.Invoke(ctx, "/pb.ecomm/RevokeSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommClient) DeleteSession(ctx context.Context, in *SessionReq, opts ...grpc.CallOption) (*SessionRes, error) {
	out := new(SessionRes)
	err := c.cc.Invoke(ctx, "/pb.ecomm/DeleteSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EcommServer is the server API for Ecomm service.
// All implementations must embed UnimplementedEcommServer
// for forward compatibility
type EcommServer interface {
	CreateProduct(context.Context, *ProductReq) (*ProductRes, error)
	GetProduct(context.Context, *ProductReq) (*ProductRes, error)
	ListProducts(context.Context, *ProductReq) (*ListProductRes, error)
	UpdateProduct(context.Context, *ProductReq) (*ProductRes, error)
	DeleteProduct(context.Context, *ProductReq) (*ProductRes, error)
	CreateOrder(context.Context, *OrderReq) (*OrderRes, error)
	GetOrder(context.Context, *OrderReq) (*OrderRes, error)
	ListOrders(context.Context, *OrderReq) (*ListOrderRes, error)
	DeleteOrder(context.Context, *OrderReq) (*OrderRes, error)
	CreateUser(context.Context, *UserReq) (*UserRes, error)
	GetUser(context.Context, *UserReq) (*UserRes, error)
	ListUsers(context.Context, *UserReq) (*ListUserRes, error)
	UpdateUser(context.Context, *UserReq) (*UserRes, error)
	DeleteUser(context.Context, *UserReq) (*UserRes, error)
	CreateSession(context.Context, *SessionReq) (*SessionRes, error)
	GetSession(context.Context, *SessionReq) (*SessionRes, error)
	RevokeSession(context.Context, *SessionReq) (*SessionRes, error)
	DeleteSession(context.Context, *SessionReq) (*SessionRes, error)
	mustEmbedUnimplementedEcommServer()
}

// UnimplementedEcommServer must be embedded to have forward compatible implementations.
type UnimplementedEcommServer struct {
}

func (UnimplementedEcommServer) CreateProduct(context.Context, *ProductReq) (*ProductRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedEcommServer) GetProduct(context.Context, *ProductReq) (*ProductRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedEcommServer) ListProducts(context.Context, *ProductReq) (*ListProductRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProducts not implemented")
}
func (UnimplementedEcommServer) UpdateProduct(context.Context, *ProductReq) (*ProductRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedEcommServer) DeleteProduct(context.Context, *ProductReq) (*ProductRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedEcommServer) CreateOrder(context.Context, *OrderReq) (*OrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedEcommServer) GetOrder(context.Context, *OrderReq) (*OrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedEcommServer) ListOrders(context.Context, *OrderReq) (*ListOrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrders not implemented")
}
func (UnimplementedEcommServer) DeleteOrder(context.Context, *OrderReq) (*OrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrder not implemented")
}
func (UnimplementedEcommServer) CreateUser(context.Context, *UserReq) (*UserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedEcommServer) GetUser(context.Context, *UserReq) (*UserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedEcommServer) ListUsers(context.Context, *UserReq) (*ListUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedEcommServer) UpdateUser(context.Context, *UserReq) (*UserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedEcommServer) DeleteUser(context.Context, *UserReq) (*UserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedEcommServer) CreateSession(context.Context, *SessionReq) (*SessionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (UnimplementedEcommServer) GetSession(context.Context, *SessionReq) (*SessionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSession not implemented")
}
func (UnimplementedEcommServer) RevokeSession(context.Context, *SessionReq) (*SessionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeSession not implemented")
}
func (UnimplementedEcommServer) DeleteSession(context.Context, *SessionReq) (*SessionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSession not implemented")
}
func (UnimplementedEcommServer) mustEmbedUnimplementedEcommServer() {}

// UnsafeEcommServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EcommServer will
// result in compilation errors.
type UnsafeEcommServer interface {
	mustEmbedUnimplementedEcommServer()
}

func RegisterEcommServer(s grpc.ServiceRegistrar, srv EcommServer) {
	s.RegisterService(&Ecomm_ServiceDesc, srv)
}

func _Ecomm_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ecomm/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).CreateProduct(ctx, req.(*ProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ecomm/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).GetProduct(ctx, req.(*ProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_ListProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).ListProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ecomm/ListProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).ListProducts(ctx, req.(*ProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ecomm/UpdateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).UpdateProduct(ctx, req.(*ProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ecomm/DeleteProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).DeleteProduct(ctx, req.(*ProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ecomm/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).CreateOrder(ctx, req.(*OrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ecomm/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).GetOrder(ctx, req.(*OrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_ListOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).ListOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ecomm/ListOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).ListOrders(ctx, req.(*OrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_DeleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).DeleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ecomm/DeleteOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).DeleteOrder(ctx, req.(*OrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ecomm/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).CreateUser(ctx, req.(*UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ecomm/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).GetUser(ctx, req.(*UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ecomm/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).ListUsers(ctx, req.(*UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ecomm/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).UpdateUser(ctx, req.(*UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ecomm/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).DeleteUser(ctx, req.(*UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ecomm/CreateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).CreateSession(ctx, req.(*SessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_GetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).GetSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ecomm/GetSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).GetSession(ctx, req.(*SessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_RevokeSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).RevokeSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ecomm/RevokeSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).RevokeSession(ctx, req.(*SessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ecomm_DeleteSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommServer).DeleteSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ecomm/DeleteSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommServer).DeleteSession(ctx, req.(*SessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Ecomm_ServiceDesc is the grpc.ServiceDesc for Ecomm service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ecomm_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ecomm",
	HandlerType: (*EcommServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _Ecomm_CreateProduct_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _Ecomm_GetProduct_Handler,
		},
		{
			MethodName: "ListProducts",
			Handler:    _Ecomm_ListProducts_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _Ecomm_UpdateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _Ecomm_DeleteProduct_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _Ecomm_CreateOrder_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _Ecomm_GetOrder_Handler,
		},
		{
			MethodName: "ListOrders",
			Handler:    _Ecomm_ListOrders_Handler,
		},
		{
			MethodName: "DeleteOrder",
			Handler:    _Ecomm_DeleteOrder_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Ecomm_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Ecomm_GetUser_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _Ecomm_ListUsers_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Ecomm_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Ecomm_DeleteUser_Handler,
		},
		{
			MethodName: "CreateSession",
			Handler:    _Ecomm_CreateSession_Handler,
		},
		{
			MethodName: "GetSession",
			Handler:    _Ecomm_GetSession_Handler,
		},
		{
			MethodName: "RevokeSession",
			Handler:    _Ecomm_RevokeSession_Handler,
		},
		{
			MethodName: "DeleteSession",
			Handler:    _Ecomm_DeleteSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
