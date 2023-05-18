// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: es.proto

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

const (
	Es_SearchForFiles_FullMethodName     = "/pb.es/SearchForFiles"
	Es_SearchForPosts_FullMethodName     = "/pb.es/SearchForPosts"
	Es_SearchForFileRank_FullMethodName  = "/pb.es/SearchForFileRank"
	Es_SearchForPostsRank_FullMethodName = "/pb.es/SearchForPostsRank"
	Es_SearchForFilesById_FullMethodName = "/pb.es/SearchForFilesById"
	Es_SearchForPostsById_FullMethodName = "/pb.es/SearchForPostsById"
	Es_UpdateFiles_FullMethodName        = "/pb.es/UpdateFiles"
	Es_UpdatePosts_FullMethodName        = "/pb.es/UpdatePosts"
	Es_InsertFile_FullMethodName         = "/pb.es/InsertFile"
	Es_InsertPost_FullMethodName         = "/pb.es/InsertPost"
)

// EsClient is the client API for Es service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EsClient interface {
	SearchForFiles(ctx context.Context, in *SearchForFilesReq, opts ...grpc.CallOption) (*SearchForFilesResp, error)
	SearchForPosts(ctx context.Context, in *SearchForPostsReq, opts ...grpc.CallOption) (*SearchForPostsResp, error)
	SearchForFileRank(ctx context.Context, in *SearchForFileRankReq, opts ...grpc.CallOption) (*SearchForFileRankResp, error)
	SearchForPostsRank(ctx context.Context, in *SearchForPostsRankReq, opts ...grpc.CallOption) (*SearchForPostsRankResp, error)
	SearchForFilesById(ctx context.Context, in *SearchForFilesByIdReq, opts ...grpc.CallOption) (*SearchForFilesByIdResp, error)
	SearchForPostsById(ctx context.Context, in *SearchForPostsByIdReq, opts ...grpc.CallOption) (*SearchForPostsByIdResp, error)
	UpdateFiles(ctx context.Context, in *UpdateFilesReq, opts ...grpc.CallOption) (*UpdateFilesResp, error)
	UpdatePosts(ctx context.Context, in *UpdatePostsReq, opts ...grpc.CallOption) (*UpdatePostsResp, error)
	InsertFile(ctx context.Context, in *InsertFileReq, opts ...grpc.CallOption) (*InsertFileResp, error)
	InsertPost(ctx context.Context, in *InsertPostReq, opts ...grpc.CallOption) (*InsertPostResp, error)
}

type esClient struct {
	cc grpc.ClientConnInterface
}

func NewEsClient(cc grpc.ClientConnInterface) EsClient {
	return &esClient{cc}
}

func (c *esClient) SearchForFiles(ctx context.Context, in *SearchForFilesReq, opts ...grpc.CallOption) (*SearchForFilesResp, error) {
	out := new(SearchForFilesResp)
	err := c.cc.Invoke(ctx, Es_SearchForFiles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *esClient) SearchForPosts(ctx context.Context, in *SearchForPostsReq, opts ...grpc.CallOption) (*SearchForPostsResp, error) {
	out := new(SearchForPostsResp)
	err := c.cc.Invoke(ctx, Es_SearchForPosts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *esClient) SearchForFileRank(ctx context.Context, in *SearchForFileRankReq, opts ...grpc.CallOption) (*SearchForFileRankResp, error) {
	out := new(SearchForFileRankResp)
	err := c.cc.Invoke(ctx, Es_SearchForFileRank_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *esClient) SearchForPostsRank(ctx context.Context, in *SearchForPostsRankReq, opts ...grpc.CallOption) (*SearchForPostsRankResp, error) {
	out := new(SearchForPostsRankResp)
	err := c.cc.Invoke(ctx, Es_SearchForPostsRank_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *esClient) SearchForFilesById(ctx context.Context, in *SearchForFilesByIdReq, opts ...grpc.CallOption) (*SearchForFilesByIdResp, error) {
	out := new(SearchForFilesByIdResp)
	err := c.cc.Invoke(ctx, Es_SearchForFilesById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *esClient) SearchForPostsById(ctx context.Context, in *SearchForPostsByIdReq, opts ...grpc.CallOption) (*SearchForPostsByIdResp, error) {
	out := new(SearchForPostsByIdResp)
	err := c.cc.Invoke(ctx, Es_SearchForPostsById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *esClient) UpdateFiles(ctx context.Context, in *UpdateFilesReq, opts ...grpc.CallOption) (*UpdateFilesResp, error) {
	out := new(UpdateFilesResp)
	err := c.cc.Invoke(ctx, Es_UpdateFiles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *esClient) UpdatePosts(ctx context.Context, in *UpdatePostsReq, opts ...grpc.CallOption) (*UpdatePostsResp, error) {
	out := new(UpdatePostsResp)
	err := c.cc.Invoke(ctx, Es_UpdatePosts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *esClient) InsertFile(ctx context.Context, in *InsertFileReq, opts ...grpc.CallOption) (*InsertFileResp, error) {
	out := new(InsertFileResp)
	err := c.cc.Invoke(ctx, Es_InsertFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *esClient) InsertPost(ctx context.Context, in *InsertPostReq, opts ...grpc.CallOption) (*InsertPostResp, error) {
	out := new(InsertPostResp)
	err := c.cc.Invoke(ctx, Es_InsertPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EsServer is the server API for Es service.
// All implementations must embed UnimplementedEsServer
// for forward compatibility
type EsServer interface {
	SearchForFiles(context.Context, *SearchForFilesReq) (*SearchForFilesResp, error)
	SearchForPosts(context.Context, *SearchForPostsReq) (*SearchForPostsResp, error)
	SearchForFileRank(context.Context, *SearchForFileRankReq) (*SearchForFileRankResp, error)
	SearchForPostsRank(context.Context, *SearchForPostsRankReq) (*SearchForPostsRankResp, error)
	SearchForFilesById(context.Context, *SearchForFilesByIdReq) (*SearchForFilesByIdResp, error)
	SearchForPostsById(context.Context, *SearchForPostsByIdReq) (*SearchForPostsByIdResp, error)
	UpdateFiles(context.Context, *UpdateFilesReq) (*UpdateFilesResp, error)
	UpdatePosts(context.Context, *UpdatePostsReq) (*UpdatePostsResp, error)
	InsertFile(context.Context, *InsertFileReq) (*InsertFileResp, error)
	InsertPost(context.Context, *InsertPostReq) (*InsertPostResp, error)
	mustEmbedUnimplementedEsServer()
}

// UnimplementedEsServer must be embedded to have forward compatible implementations.
type UnimplementedEsServer struct {
}

func (UnimplementedEsServer) SearchForFiles(context.Context, *SearchForFilesReq) (*SearchForFilesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchForFiles not implemented")
}
func (UnimplementedEsServer) SearchForPosts(context.Context, *SearchForPostsReq) (*SearchForPostsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchForPosts not implemented")
}
func (UnimplementedEsServer) SearchForFileRank(context.Context, *SearchForFileRankReq) (*SearchForFileRankResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchForFileRank not implemented")
}
func (UnimplementedEsServer) SearchForPostsRank(context.Context, *SearchForPostsRankReq) (*SearchForPostsRankResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchForPostsRank not implemented")
}
func (UnimplementedEsServer) SearchForFilesById(context.Context, *SearchForFilesByIdReq) (*SearchForFilesByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchForFilesById not implemented")
}
func (UnimplementedEsServer) SearchForPostsById(context.Context, *SearchForPostsByIdReq) (*SearchForPostsByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchForPostsById not implemented")
}
func (UnimplementedEsServer) UpdateFiles(context.Context, *UpdateFilesReq) (*UpdateFilesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFiles not implemented")
}
func (UnimplementedEsServer) UpdatePosts(context.Context, *UpdatePostsReq) (*UpdatePostsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePosts not implemented")
}
func (UnimplementedEsServer) InsertFile(context.Context, *InsertFileReq) (*InsertFileResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertFile not implemented")
}
func (UnimplementedEsServer) InsertPost(context.Context, *InsertPostReq) (*InsertPostResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertPost not implemented")
}
func (UnimplementedEsServer) mustEmbedUnimplementedEsServer() {}

// UnsafeEsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EsServer will
// result in compilation errors.
type UnsafeEsServer interface {
	mustEmbedUnimplementedEsServer()
}

func RegisterEsServer(s grpc.ServiceRegistrar, srv EsServer) {
	s.RegisterService(&Es_ServiceDesc, srv)
}

func _Es_SearchForFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchForFilesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EsServer).SearchForFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Es_SearchForFiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EsServer).SearchForFiles(ctx, req.(*SearchForFilesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Es_SearchForPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchForPostsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EsServer).SearchForPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Es_SearchForPosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EsServer).SearchForPosts(ctx, req.(*SearchForPostsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Es_SearchForFileRank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchForFileRankReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EsServer).SearchForFileRank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Es_SearchForFileRank_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EsServer).SearchForFileRank(ctx, req.(*SearchForFileRankReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Es_SearchForPostsRank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchForPostsRankReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EsServer).SearchForPostsRank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Es_SearchForPostsRank_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EsServer).SearchForPostsRank(ctx, req.(*SearchForPostsRankReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Es_SearchForFilesById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchForFilesByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EsServer).SearchForFilesById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Es_SearchForFilesById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EsServer).SearchForFilesById(ctx, req.(*SearchForFilesByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Es_SearchForPostsById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchForPostsByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EsServer).SearchForPostsById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Es_SearchForPostsById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EsServer).SearchForPostsById(ctx, req.(*SearchForPostsByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Es_UpdateFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFilesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EsServer).UpdateFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Es_UpdateFiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EsServer).UpdateFiles(ctx, req.(*UpdateFilesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Es_UpdatePosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EsServer).UpdatePosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Es_UpdatePosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EsServer).UpdatePosts(ctx, req.(*UpdatePostsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Es_InsertFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertFileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EsServer).InsertFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Es_InsertFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EsServer).InsertFile(ctx, req.(*InsertFileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Es_InsertPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertPostReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EsServer).InsertPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Es_InsertPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EsServer).InsertPost(ctx, req.(*InsertPostReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Es_ServiceDesc is the grpc.ServiceDesc for Es service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Es_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.es",
	HandlerType: (*EsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchForFiles",
			Handler:    _Es_SearchForFiles_Handler,
		},
		{
			MethodName: "SearchForPosts",
			Handler:    _Es_SearchForPosts_Handler,
		},
		{
			MethodName: "SearchForFileRank",
			Handler:    _Es_SearchForFileRank_Handler,
		},
		{
			MethodName: "SearchForPostsRank",
			Handler:    _Es_SearchForPostsRank_Handler,
		},
		{
			MethodName: "SearchForFilesById",
			Handler:    _Es_SearchForFilesById_Handler,
		},
		{
			MethodName: "SearchForPostsById",
			Handler:    _Es_SearchForPostsById_Handler,
		},
		{
			MethodName: "UpdateFiles",
			Handler:    _Es_UpdateFiles_Handler,
		},
		{
			MethodName: "UpdatePosts",
			Handler:    _Es_UpdatePosts_Handler,
		},
		{
			MethodName: "InsertFile",
			Handler:    _Es_InsertFile_Handler,
		},
		{
			MethodName: "InsertPost",
			Handler:    _Es_InsertPost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "es.proto",
}
