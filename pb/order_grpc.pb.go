// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: pb/order.proto

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

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	Create(ctx context.Context, in *OrderCreateRequest, opts ...grpc.CallOption) (*Empty, error)
	ChangeStatus(ctx context.Context, in *OrderChangeStatus, opts ...grpc.CallOption) (*OperationResponse, error)
	FindOne(ctx context.Context, in *OrderFindOneRequest, opts ...grpc.CallOption) (*OrderFindOneResponse, error)
	FindAll(ctx context.Context, in *OrderFindAllRequest, opts ...grpc.CallOption) (*OrderFindAllResponse, error)
	SumIncome(ctx context.Context, in *OrderSumIncomeRequest, opts ...grpc.CallOption) (*OrderSumResponse, error)
	Cancel(ctx context.Context, in *OrderCancelRequest, opts ...grpc.CallOption) (*OperationResponse, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) Create(ctx context.Context, in *OrderCreateRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/OrderService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) ChangeStatus(ctx context.Context, in *OrderChangeStatus, opts ...grpc.CallOption) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/OrderService/ChangeStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) FindOne(ctx context.Context, in *OrderFindOneRequest, opts ...grpc.CallOption) (*OrderFindOneResponse, error) {
	out := new(OrderFindOneResponse)
	err := c.cc.Invoke(ctx, "/OrderService/FindOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) FindAll(ctx context.Context, in *OrderFindAllRequest, opts ...grpc.CallOption) (*OrderFindAllResponse, error) {
	out := new(OrderFindAllResponse)
	err := c.cc.Invoke(ctx, "/OrderService/FindAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) SumIncome(ctx context.Context, in *OrderSumIncomeRequest, opts ...grpc.CallOption) (*OrderSumResponse, error) {
	out := new(OrderSumResponse)
	err := c.cc.Invoke(ctx, "/OrderService/SumIncome", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) Cancel(ctx context.Context, in *OrderCancelRequest, opts ...grpc.CallOption) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/OrderService/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations must embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	Create(context.Context, *OrderCreateRequest) (*Empty, error)
	ChangeStatus(context.Context, *OrderChangeStatus) (*OperationResponse, error)
	FindOne(context.Context, *OrderFindOneRequest) (*OrderFindOneResponse, error)
	FindAll(context.Context, *OrderFindAllRequest) (*OrderFindAllResponse, error)
	SumIncome(context.Context, *OrderSumIncomeRequest) (*OrderSumResponse, error)
	Cancel(context.Context, *OrderCancelRequest) (*OperationResponse, error)
	mustEmbedUnimplementedOrderServiceServer()
}

// UnimplementedOrderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) Create(context.Context, *OrderCreateRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedOrderServiceServer) ChangeStatus(context.Context, *OrderChangeStatus) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeStatus not implemented")
}
func (UnimplementedOrderServiceServer) FindOne(context.Context, *OrderFindOneRequest) (*OrderFindOneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOne not implemented")
}
func (UnimplementedOrderServiceServer) FindAll(context.Context, *OrderFindAllRequest) (*OrderFindAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAll not implemented")
}
func (UnimplementedOrderServiceServer) SumIncome(context.Context, *OrderSumIncomeRequest) (*OrderSumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SumIncome not implemented")
}
func (UnimplementedOrderServiceServer) Cancel(context.Context, *OrderCancelRequest) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).Create(ctx, req.(*OrderCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_ChangeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderChangeStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ChangeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/ChangeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ChangeStatus(ctx, req.(*OrderChangeStatus))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_FindOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderFindOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).FindOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/FindOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).FindOne(ctx, req.(*OrderFindOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_FindAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderFindAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).FindAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/FindAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).FindAll(ctx, req.(*OrderFindAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_SumIncome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderSumIncomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).SumIncome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/SumIncome",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).SumIncome(ctx, req.(*OrderSumIncomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderCancelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).Cancel(ctx, req.(*OrderCancelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _OrderService_Create_Handler,
		},
		{
			MethodName: "ChangeStatus",
			Handler:    _OrderService_ChangeStatus_Handler,
		},
		{
			MethodName: "FindOne",
			Handler:    _OrderService_FindOne_Handler,
		},
		{
			MethodName: "FindAll",
			Handler:    _OrderService_FindAll_Handler,
		},
		{
			MethodName: "SumIncome",
			Handler:    _OrderService_SumIncome_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _OrderService_Cancel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/order.proto",
}
