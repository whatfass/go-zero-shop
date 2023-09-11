// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.2
// source: order.proto

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

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderClient interface {
	CreateProductOrder(ctx context.Context, in *CreateProductOrderReq, opts ...grpc.CallOption) (*CreateOrderResp, error)
	CreateSeckillOrder(ctx context.Context, in *CreateSeckillOrderReq, opts ...grpc.CallOption) (*CreateOrderResp, error)
	OrderList(ctx context.Context, in *OrderListReq, opts ...grpc.CallOption) (*OrderListResp, error)
	OrderDetail(ctx context.Context, in *OrderDetailReq, opts ...grpc.CallOption) (*OrderDetailResp, error)
	DeleteOrder(ctx context.Context, in *DeleteOrderReq, opts ...grpc.CallOption) (*DeleteOrderResp, error)
	UpdateOrderStatus(ctx context.Context, in *UpdateOrderStatusReq, opts ...grpc.CallOption) (*UpdateOrderStatusResp, error)
	// other
	GetOrderOnlyDetail(ctx context.Context, in *GetOrderOnlyDetailReq, opts ...grpc.CallOption) (*GetOrderOnlyDetailResp, error)
}

type orderClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderClient(cc grpc.ClientConnInterface) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) CreateProductOrder(ctx context.Context, in *CreateProductOrderReq, opts ...grpc.CallOption) (*CreateOrderResp, error) {
	out := new(CreateOrderResp)
	err := c.cc.Invoke(ctx, "/pb.order/createProductOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) CreateSeckillOrder(ctx context.Context, in *CreateSeckillOrderReq, opts ...grpc.CallOption) (*CreateOrderResp, error) {
	out := new(CreateOrderResp)
	err := c.cc.Invoke(ctx, "/pb.order/createSeckillOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) OrderList(ctx context.Context, in *OrderListReq, opts ...grpc.CallOption) (*OrderListResp, error) {
	out := new(OrderListResp)
	err := c.cc.Invoke(ctx, "/pb.order/orderList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) OrderDetail(ctx context.Context, in *OrderDetailReq, opts ...grpc.CallOption) (*OrderDetailResp, error) {
	out := new(OrderDetailResp)
	err := c.cc.Invoke(ctx, "/pb.order/orderDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) DeleteOrder(ctx context.Context, in *DeleteOrderReq, opts ...grpc.CallOption) (*DeleteOrderResp, error) {
	out := new(DeleteOrderResp)
	err := c.cc.Invoke(ctx, "/pb.order/deleteOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) UpdateOrderStatus(ctx context.Context, in *UpdateOrderStatusReq, opts ...grpc.CallOption) (*UpdateOrderStatusResp, error) {
	out := new(UpdateOrderStatusResp)
	err := c.cc.Invoke(ctx, "/pb.order/updateOrderStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) GetOrderOnlyDetail(ctx context.Context, in *GetOrderOnlyDetailReq, opts ...grpc.CallOption) (*GetOrderOnlyDetailResp, error) {
	out := new(GetOrderOnlyDetailResp)
	err := c.cc.Invoke(ctx, "/pb.order/getOrderOnlyDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
// All implementations must embed UnimplementedOrderServer
// for forward compatibility
type OrderServer interface {
	CreateProductOrder(context.Context, *CreateProductOrderReq) (*CreateOrderResp, error)
	CreateSeckillOrder(context.Context, *CreateSeckillOrderReq) (*CreateOrderResp, error)
	OrderList(context.Context, *OrderListReq) (*OrderListResp, error)
	OrderDetail(context.Context, *OrderDetailReq) (*OrderDetailResp, error)
	DeleteOrder(context.Context, *DeleteOrderReq) (*DeleteOrderResp, error)
	UpdateOrderStatus(context.Context, *UpdateOrderStatusReq) (*UpdateOrderStatusResp, error)
	// other
	GetOrderOnlyDetail(context.Context, *GetOrderOnlyDetailReq) (*GetOrderOnlyDetailResp, error)
	mustEmbedUnimplementedOrderServer()
}

// UnimplementedOrderServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (UnimplementedOrderServer) CreateProductOrder(context.Context, *CreateProductOrderReq) (*CreateOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProductOrder not implemented")
}
func (UnimplementedOrderServer) CreateSeckillOrder(context.Context, *CreateSeckillOrderReq) (*CreateOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSeckillOrder not implemented")
}
func (UnimplementedOrderServer) OrderList(context.Context, *OrderListReq) (*OrderListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderList not implemented")
}
func (UnimplementedOrderServer) OrderDetail(context.Context, *OrderDetailReq) (*OrderDetailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderDetail not implemented")
}
func (UnimplementedOrderServer) DeleteOrder(context.Context, *DeleteOrderReq) (*DeleteOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrder not implemented")
}
func (UnimplementedOrderServer) UpdateOrderStatus(context.Context, *UpdateOrderStatusReq) (*UpdateOrderStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrderStatus not implemented")
}
func (UnimplementedOrderServer) GetOrderOnlyDetail(context.Context, *GetOrderOnlyDetailReq) (*GetOrderOnlyDetailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderOnlyDetail not implemented")
}
func (UnimplementedOrderServer) mustEmbedUnimplementedOrderServer() {}

// UnsafeOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServer will
// result in compilation errors.
type UnsafeOrderServer interface {
	mustEmbedUnimplementedOrderServer()
}

func RegisterOrderServer(s grpc.ServiceRegistrar, srv OrderServer) {
	s.RegisterService(&Order_ServiceDesc, srv)
}

func _Order_CreateProductOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CreateProductOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.order/createProductOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CreateProductOrder(ctx, req.(*CreateProductOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_CreateSeckillOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSeckillOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CreateSeckillOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.order/createSeckillOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CreateSeckillOrder(ctx, req.(*CreateSeckillOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_OrderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).OrderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.order/orderList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).OrderList(ctx, req.(*OrderListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_OrderDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).OrderDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.order/orderDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).OrderDetail(ctx, req.(*OrderDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_DeleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).DeleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.order/deleteOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).DeleteOrder(ctx, req.(*DeleteOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_UpdateOrderStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).UpdateOrderStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.order/updateOrderStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).UpdateOrderStatus(ctx, req.(*UpdateOrderStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_GetOrderOnlyDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderOnlyDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).GetOrderOnlyDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.order/getOrderOnlyDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).GetOrderOnlyDetail(ctx, req.(*GetOrderOnlyDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Order_ServiceDesc is the grpc.ServiceDesc for Order service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Order_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createProductOrder",
			Handler:    _Order_CreateProductOrder_Handler,
		},
		{
			MethodName: "createSeckillOrder",
			Handler:    _Order_CreateSeckillOrder_Handler,
		},
		{
			MethodName: "orderList",
			Handler:    _Order_OrderList_Handler,
		},
		{
			MethodName: "orderDetail",
			Handler:    _Order_OrderDetail_Handler,
		},
		{
			MethodName: "deleteOrder",
			Handler:    _Order_DeleteOrder_Handler,
		},
		{
			MethodName: "updateOrderStatus",
			Handler:    _Order_UpdateOrderStatus_Handler,
		},
		{
			MethodName: "getOrderOnlyDetail",
			Handler:    _Order_GetOrderOnlyDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}