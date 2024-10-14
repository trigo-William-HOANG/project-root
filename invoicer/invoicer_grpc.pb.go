// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.0
// source: invoicer.proto

package invoicer

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
	Invoicer_Create_FullMethodName = "/Invoicer/Create"
)

// InvoicerClient is the client API for Invoicer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InvoicerClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
}

type invoicerClient struct {
	cc grpc.ClientConnInterface
}

func NewInvoicerClient(cc grpc.ClientConnInterface) InvoicerClient {
	return &invoicerClient{cc}
}

func (c *invoicerClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, Invoicer_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InvoicerServer is the server API for Invoicer service.
// All implementations must embed UnimplementedInvoicerServer
// for forward compatibility.
type InvoicerServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	mustEmbedUnimplementedInvoicerServer()
}

// UnimplementedInvoicerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedInvoicerServer struct{}

func (UnimplementedInvoicerServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedInvoicerServer) mustEmbedUnimplementedInvoicerServer() {}
func (UnimplementedInvoicerServer) testEmbeddedByValue()                  {}

// UnsafeInvoicerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InvoicerServer will
// result in compilation errors.
type UnsafeInvoicerServer interface {
	mustEmbedUnimplementedInvoicerServer()
}

func RegisterInvoicerServer(s grpc.ServiceRegistrar, srv InvoicerServer) {
	// If the following call pancis, it indicates UnimplementedInvoicerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Invoicer_ServiceDesc, srv)
}

func _Invoicer_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoicerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Invoicer_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoicerServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Invoicer_ServiceDesc is the grpc.ServiceDesc for Invoicer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Invoicer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Invoicer",
	HandlerType: (*InvoicerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Invoicer_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "invoicer.proto",
}
