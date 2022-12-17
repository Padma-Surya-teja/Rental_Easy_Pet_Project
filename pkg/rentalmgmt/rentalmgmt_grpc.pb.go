// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: pkg/rentalmgmt/rentalmgmt.proto

package rentalmgmt

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

// Rental_Easy_FunctionalitiesClient is the client API for Rental_Easy_Functionalities service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Rental_Easy_FunctionalitiesClient interface {
	CreateUser(ctx context.Context, in *NewUser, opts ...grpc.CallOption) (*UserId, error)
	CreateItem(ctx context.Context, in *NewItem, opts ...grpc.CallOption) (*ItemId, error)
	GetAllItems(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Items, error)
	GetItemById(ctx context.Context, in *ItemId, opts ...grpc.CallOption) (*DetailedItem, error)
	UpdateItem(ctx context.Context, in *DetailedItem, opts ...grpc.CallOption) (*ItemId, error)
}

type rental_Easy_FunctionalitiesClient struct {
	cc grpc.ClientConnInterface
}

func NewRental_Easy_FunctionalitiesClient(cc grpc.ClientConnInterface) Rental_Easy_FunctionalitiesClient {
	return &rental_Easy_FunctionalitiesClient{cc}
}

func (c *rental_Easy_FunctionalitiesClient) CreateUser(ctx context.Context, in *NewUser, opts ...grpc.CallOption) (*UserId, error) {
	out := new(UserId)
	err := c.cc.Invoke(ctx, "/rentalmgmt.Rental_Easy_Functionalities/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rental_Easy_FunctionalitiesClient) CreateItem(ctx context.Context, in *NewItem, opts ...grpc.CallOption) (*ItemId, error) {
	out := new(ItemId)
	err := c.cc.Invoke(ctx, "/rentalmgmt.Rental_Easy_Functionalities/CreateItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rental_Easy_FunctionalitiesClient) GetAllItems(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Items, error) {
	out := new(Items)
	err := c.cc.Invoke(ctx, "/rentalmgmt.Rental_Easy_Functionalities/GetAllItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rental_Easy_FunctionalitiesClient) GetItemById(ctx context.Context, in *ItemId, opts ...grpc.CallOption) (*DetailedItem, error) {
	out := new(DetailedItem)
	err := c.cc.Invoke(ctx, "/rentalmgmt.Rental_Easy_Functionalities/GetItemById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rental_Easy_FunctionalitiesClient) UpdateItem(ctx context.Context, in *DetailedItem, opts ...grpc.CallOption) (*ItemId, error) {
	out := new(ItemId)
	err := c.cc.Invoke(ctx, "/rentalmgmt.Rental_Easy_Functionalities/UpdateItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Rental_Easy_FunctionalitiesServer is the server API for Rental_Easy_Functionalities service.
// All implementations must embed UnimplementedRental_Easy_FunctionalitiesServer
// for forward compatibility
type Rental_Easy_FunctionalitiesServer interface {
	CreateUser(context.Context, *NewUser) (*UserId, error)
	CreateItem(context.Context, *NewItem) (*ItemId, error)
	GetAllItems(context.Context, *Request) (*Items, error)
	GetItemById(context.Context, *ItemId) (*DetailedItem, error)
	UpdateItem(context.Context, *DetailedItem) (*ItemId, error)
	mustEmbedUnimplementedRental_Easy_FunctionalitiesServer()
}

// UnimplementedRental_Easy_FunctionalitiesServer must be embedded to have forward compatible implementations.
type UnimplementedRental_Easy_FunctionalitiesServer struct {
}

func (UnimplementedRental_Easy_FunctionalitiesServer) CreateUser(context.Context, *NewUser) (*UserId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedRental_Easy_FunctionalitiesServer) CreateItem(context.Context, *NewItem) (*ItemId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateItem not implemented")
}
func (UnimplementedRental_Easy_FunctionalitiesServer) GetAllItems(context.Context, *Request) (*Items, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllItems not implemented")
}
func (UnimplementedRental_Easy_FunctionalitiesServer) GetItemById(context.Context, *ItemId) (*DetailedItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItemById not implemented")
}
func (UnimplementedRental_Easy_FunctionalitiesServer) UpdateItem(context.Context, *DetailedItem) (*ItemId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateItem not implemented")
}
func (UnimplementedRental_Easy_FunctionalitiesServer) mustEmbedUnimplementedRental_Easy_FunctionalitiesServer() {
}

// UnsafeRental_Easy_FunctionalitiesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Rental_Easy_FunctionalitiesServer will
// result in compilation errors.
type UnsafeRental_Easy_FunctionalitiesServer interface {
	mustEmbedUnimplementedRental_Easy_FunctionalitiesServer()
}

func RegisterRental_Easy_FunctionalitiesServer(s grpc.ServiceRegistrar, srv Rental_Easy_FunctionalitiesServer) {
	s.RegisterService(&Rental_Easy_Functionalities_ServiceDesc, srv)
}

func _Rental_Easy_Functionalities_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Rental_Easy_FunctionalitiesServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rentalmgmt.Rental_Easy_Functionalities/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Rental_Easy_FunctionalitiesServer).CreateUser(ctx, req.(*NewUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rental_Easy_Functionalities_CreateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Rental_Easy_FunctionalitiesServer).CreateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rentalmgmt.Rental_Easy_Functionalities/CreateItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Rental_Easy_FunctionalitiesServer).CreateItem(ctx, req.(*NewItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rental_Easy_Functionalities_GetAllItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Rental_Easy_FunctionalitiesServer).GetAllItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rentalmgmt.Rental_Easy_Functionalities/GetAllItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Rental_Easy_FunctionalitiesServer).GetAllItems(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rental_Easy_Functionalities_GetItemById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Rental_Easy_FunctionalitiesServer).GetItemById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rentalmgmt.Rental_Easy_Functionalities/GetItemById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Rental_Easy_FunctionalitiesServer).GetItemById(ctx, req.(*ItemId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rental_Easy_Functionalities_UpdateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailedItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Rental_Easy_FunctionalitiesServer).UpdateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rentalmgmt.Rental_Easy_Functionalities/UpdateItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Rental_Easy_FunctionalitiesServer).UpdateItem(ctx, req.(*DetailedItem))
	}
	return interceptor(ctx, in, info, handler)
}

// Rental_Easy_Functionalities_ServiceDesc is the grpc.ServiceDesc for Rental_Easy_Functionalities service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rental_Easy_Functionalities_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rentalmgmt.Rental_Easy_Functionalities",
	HandlerType: (*Rental_Easy_FunctionalitiesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Rental_Easy_Functionalities_CreateUser_Handler,
		},
		{
			MethodName: "CreateItem",
			Handler:    _Rental_Easy_Functionalities_CreateItem_Handler,
		},
		{
			MethodName: "GetAllItems",
			Handler:    _Rental_Easy_Functionalities_GetAllItems_Handler,
		},
		{
			MethodName: "GetItemById",
			Handler:    _Rental_Easy_Functionalities_GetItemById_Handler,
		},
		{
			MethodName: "UpdateItem",
			Handler:    _Rental_Easy_Functionalities_UpdateItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/rentalmgmt/rentalmgmt.proto",
}
