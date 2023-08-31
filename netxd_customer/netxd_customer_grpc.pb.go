// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: netxd_customer/netxd_customer.proto

package bank_customer_proto_git

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

// Customer_ServiceClient is the client API for Customer_Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Customer_ServiceClient interface {
	CreateCustomer(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*CustomerResponse, error)
}

type customer_ServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomer_ServiceClient(cc grpc.ClientConnInterface) Customer_ServiceClient {
	return &customer_ServiceClient{cc}
}

func (c *customer_ServiceClient) CreateCustomer(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*CustomerResponse, error) {
	out := new(CustomerResponse)
	err := c.cc.Invoke(ctx, "/netxd_customer.Customer_Service/CreateCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Customer_ServiceServer is the server API for Customer_Service service.
// All implementations must embed UnimplementedCustomer_ServiceServer
// for forward compatibility
type Customer_ServiceServer interface {
	CreateCustomer(context.Context, *Customer) (*CustomerResponse, error)
	mustEmbedUnimplementedCustomer_ServiceServer()
}

// UnimplementedCustomer_ServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomer_ServiceServer struct {
}

func (UnimplementedCustomer_ServiceServer) CreateCustomer(context.Context, *Customer) (*CustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomer not implemented")
}
func (UnimplementedCustomer_ServiceServer) mustEmbedUnimplementedCustomer_ServiceServer() {}

// UnsafeCustomer_ServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Customer_ServiceServer will
// result in compilation errors.
type UnsafeCustomer_ServiceServer interface {
	mustEmbedUnimplementedCustomer_ServiceServer()
}

func RegisterCustomer_ServiceServer(s grpc.ServiceRegistrar, srv Customer_ServiceServer) {
	s.RegisterService(&Customer_Service_ServiceDesc, srv)
}

func _Customer_Service_CreateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Customer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Customer_ServiceServer).CreateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/netxd_customer.Customer_Service/CreateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Customer_ServiceServer).CreateCustomer(ctx, req.(*Customer))
	}
	return interceptor(ctx, in, info, handler)
}

// Customer_Service_ServiceDesc is the grpc.ServiceDesc for Customer_Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Customer_Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "netxd_customer.Customer_Service",
	HandlerType: (*Customer_ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCustomer",
			Handler:    _Customer_Service_CreateCustomer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "netxd_customer/netxd_customer.proto",
}
