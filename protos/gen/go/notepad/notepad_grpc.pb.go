// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: notepad/notepad.proto

package ssov1

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

// NotepadClient is the client API for Notepad service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotepadClient interface {
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error)
	GetOne(ctx context.Context, in *GetOneRequest, opts ...grpc.CallOption) (*GetOneResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	IsOwner(ctx context.Context, in *IsOwnerRequest, opts ...grpc.CallOption) (*IsOwnerResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type notepadClient struct {
	cc grpc.ClientConnInterface
}

func NewNotepadClient(cc grpc.ClientConnInterface) NotepadClient {
	return &notepadClient{cc}
}

func (c *notepadClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error) {
	out := new(SetResponse)
	err := c.cc.Invoke(ctx, "/notepad.Notepad/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notepadClient) GetOne(ctx context.Context, in *GetOneRequest, opts ...grpc.CallOption) (*GetOneResponse, error) {
	out := new(GetOneResponse)
	err := c.cc.Invoke(ctx, "/notepad.Notepad/GetOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notepadClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/notepad.Notepad/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notepadClient) IsOwner(ctx context.Context, in *IsOwnerRequest, opts ...grpc.CallOption) (*IsOwnerResponse, error) {
	out := new(IsOwnerResponse)
	err := c.cc.Invoke(ctx, "/notepad.Notepad/IsOwner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notepadClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/notepad.Notepad/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotepadServer is the server API for Notepad service.
// All implementations must embed UnimplementedNotepadServer
// for forward compatibility
type NotepadServer interface {
	Set(context.Context, *SetRequest) (*SetResponse, error)
	GetOne(context.Context, *GetOneRequest) (*GetOneResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	IsOwner(context.Context, *IsOwnerRequest) (*IsOwnerResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedNotepadServer()
}

// UnimplementedNotepadServer must be embedded to have forward compatible implementations.
type UnimplementedNotepadServer struct {
}

func (UnimplementedNotepadServer) Set(context.Context, *SetRequest) (*SetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedNotepadServer) GetOne(context.Context, *GetOneRequest) (*GetOneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOne not implemented")
}
func (UnimplementedNotepadServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedNotepadServer) IsOwner(context.Context, *IsOwnerRequest) (*IsOwnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsOwner not implemented")
}
func (UnimplementedNotepadServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedNotepadServer) mustEmbedUnimplementedNotepadServer() {}

// UnsafeNotepadServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotepadServer will
// result in compilation errors.
type UnsafeNotepadServer interface {
	mustEmbedUnimplementedNotepadServer()
}

func RegisterNotepadServer(s grpc.ServiceRegistrar, srv NotepadServer) {
	s.RegisterService(&Notepad_ServiceDesc, srv)
}

func _Notepad_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotepadServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notepad.Notepad/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotepadServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notepad_GetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotepadServer).GetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notepad.Notepad/GetOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotepadServer).GetOne(ctx, req.(*GetOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notepad_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotepadServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notepad.Notepad/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotepadServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notepad_IsOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotepadServer).IsOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notepad.Notepad/IsOwner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotepadServer).IsOwner(ctx, req.(*IsOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notepad_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotepadServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notepad.Notepad/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotepadServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Notepad_ServiceDesc is the grpc.ServiceDesc for Notepad service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Notepad_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notepad.Notepad",
	HandlerType: (*NotepadServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _Notepad_Set_Handler,
		},
		{
			MethodName: "GetOne",
			Handler:    _Notepad_GetOne_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _Notepad_GetAll_Handler,
		},
		{
			MethodName: "IsOwner",
			Handler:    _Notepad_IsOwner_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Notepad_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notepad/notepad.proto",
}
