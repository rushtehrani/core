// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package gen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// WorkflowTemplateServiceClient is the client API for WorkflowTemplateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkflowTemplateServiceClient interface {
	// Get the generated WorkflowTemplate, applying any modifications based on the content
	GenerateWorkflowTemplate(ctx context.Context, in *GenerateWorkflowTemplateRequest, opts ...grpc.CallOption) (*WorkflowTemplate, error)
	CreateWorkflowTemplate(ctx context.Context, in *CreateWorkflowTemplateRequest, opts ...grpc.CallOption) (*WorkflowTemplate, error)
	CreateWorkflowTemplateVersion(ctx context.Context, in *CreateWorkflowTemplateRequest, opts ...grpc.CallOption) (*WorkflowTemplate, error)
	GetWorkflowTemplate(ctx context.Context, in *GetWorkflowTemplateRequest, opts ...grpc.CallOption) (*WorkflowTemplate, error)
	ListWorkflowTemplateVersions(ctx context.Context, in *ListWorkflowTemplateVersionsRequest, opts ...grpc.CallOption) (*ListWorkflowTemplateVersionsResponse, error)
	ListWorkflowTemplates(ctx context.Context, in *ListWorkflowTemplatesRequest, opts ...grpc.CallOption) (*ListWorkflowTemplatesResponse, error)
	CloneWorkflowTemplate(ctx context.Context, in *CloneWorkflowTemplateRequest, opts ...grpc.CallOption) (*WorkflowTemplate, error)
	ArchiveWorkflowTemplate(ctx context.Context, in *ArchiveWorkflowTemplateRequest, opts ...grpc.CallOption) (*ArchiveWorkflowTemplateResponse, error)
	ListWorkflowTemplatesField(ctx context.Context, in *ListWorkflowTemplatesFieldRequest, opts ...grpc.CallOption) (*ListWorkflowTemplatesFieldResponse, error)
}

type workflowTemplateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkflowTemplateServiceClient(cc grpc.ClientConnInterface) WorkflowTemplateServiceClient {
	return &workflowTemplateServiceClient{cc}
}

func (c *workflowTemplateServiceClient) GenerateWorkflowTemplate(ctx context.Context, in *GenerateWorkflowTemplateRequest, opts ...grpc.CallOption) (*WorkflowTemplate, error) {
	out := new(WorkflowTemplate)
	err := c.cc.Invoke(ctx, "/api.WorkflowTemplateService/GenerateWorkflowTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowTemplateServiceClient) CreateWorkflowTemplate(ctx context.Context, in *CreateWorkflowTemplateRequest, opts ...grpc.CallOption) (*WorkflowTemplate, error) {
	out := new(WorkflowTemplate)
	err := c.cc.Invoke(ctx, "/api.WorkflowTemplateService/CreateWorkflowTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowTemplateServiceClient) CreateWorkflowTemplateVersion(ctx context.Context, in *CreateWorkflowTemplateRequest, opts ...grpc.CallOption) (*WorkflowTemplate, error) {
	out := new(WorkflowTemplate)
	err := c.cc.Invoke(ctx, "/api.WorkflowTemplateService/CreateWorkflowTemplateVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowTemplateServiceClient) GetWorkflowTemplate(ctx context.Context, in *GetWorkflowTemplateRequest, opts ...grpc.CallOption) (*WorkflowTemplate, error) {
	out := new(WorkflowTemplate)
	err := c.cc.Invoke(ctx, "/api.WorkflowTemplateService/GetWorkflowTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowTemplateServiceClient) ListWorkflowTemplateVersions(ctx context.Context, in *ListWorkflowTemplateVersionsRequest, opts ...grpc.CallOption) (*ListWorkflowTemplateVersionsResponse, error) {
	out := new(ListWorkflowTemplateVersionsResponse)
	err := c.cc.Invoke(ctx, "/api.WorkflowTemplateService/ListWorkflowTemplateVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowTemplateServiceClient) ListWorkflowTemplates(ctx context.Context, in *ListWorkflowTemplatesRequest, opts ...grpc.CallOption) (*ListWorkflowTemplatesResponse, error) {
	out := new(ListWorkflowTemplatesResponse)
	err := c.cc.Invoke(ctx, "/api.WorkflowTemplateService/ListWorkflowTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowTemplateServiceClient) CloneWorkflowTemplate(ctx context.Context, in *CloneWorkflowTemplateRequest, opts ...grpc.CallOption) (*WorkflowTemplate, error) {
	out := new(WorkflowTemplate)
	err := c.cc.Invoke(ctx, "/api.WorkflowTemplateService/CloneWorkflowTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowTemplateServiceClient) ArchiveWorkflowTemplate(ctx context.Context, in *ArchiveWorkflowTemplateRequest, opts ...grpc.CallOption) (*ArchiveWorkflowTemplateResponse, error) {
	out := new(ArchiveWorkflowTemplateResponse)
	err := c.cc.Invoke(ctx, "/api.WorkflowTemplateService/ArchiveWorkflowTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowTemplateServiceClient) ListWorkflowTemplatesField(ctx context.Context, in *ListWorkflowTemplatesFieldRequest, opts ...grpc.CallOption) (*ListWorkflowTemplatesFieldResponse, error) {
	out := new(ListWorkflowTemplatesFieldResponse)
	err := c.cc.Invoke(ctx, "/api.WorkflowTemplateService/ListWorkflowTemplatesField", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkflowTemplateServiceServer is the server API for WorkflowTemplateService service.
// All implementations must embed UnimplementedWorkflowTemplateServiceServer
// for forward compatibility
type WorkflowTemplateServiceServer interface {
	// Get the generated WorkflowTemplate, applying any modifications based on the content
	GenerateWorkflowTemplate(context.Context, *GenerateWorkflowTemplateRequest) (*WorkflowTemplate, error)
	CreateWorkflowTemplate(context.Context, *CreateWorkflowTemplateRequest) (*WorkflowTemplate, error)
	CreateWorkflowTemplateVersion(context.Context, *CreateWorkflowTemplateRequest) (*WorkflowTemplate, error)
	GetWorkflowTemplate(context.Context, *GetWorkflowTemplateRequest) (*WorkflowTemplate, error)
	ListWorkflowTemplateVersions(context.Context, *ListWorkflowTemplateVersionsRequest) (*ListWorkflowTemplateVersionsResponse, error)
	ListWorkflowTemplates(context.Context, *ListWorkflowTemplatesRequest) (*ListWorkflowTemplatesResponse, error)
	CloneWorkflowTemplate(context.Context, *CloneWorkflowTemplateRequest) (*WorkflowTemplate, error)
	ArchiveWorkflowTemplate(context.Context, *ArchiveWorkflowTemplateRequest) (*ArchiveWorkflowTemplateResponse, error)
	ListWorkflowTemplatesField(context.Context, *ListWorkflowTemplatesFieldRequest) (*ListWorkflowTemplatesFieldResponse, error)
	mustEmbedUnimplementedWorkflowTemplateServiceServer()
}

// UnimplementedWorkflowTemplateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkflowTemplateServiceServer struct {
}

func (UnimplementedWorkflowTemplateServiceServer) GenerateWorkflowTemplate(context.Context, *GenerateWorkflowTemplateRequest) (*WorkflowTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateWorkflowTemplate not implemented")
}
func (UnimplementedWorkflowTemplateServiceServer) CreateWorkflowTemplate(context.Context, *CreateWorkflowTemplateRequest) (*WorkflowTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkflowTemplate not implemented")
}
func (UnimplementedWorkflowTemplateServiceServer) CreateWorkflowTemplateVersion(context.Context, *CreateWorkflowTemplateRequest) (*WorkflowTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkflowTemplateVersion not implemented")
}
func (UnimplementedWorkflowTemplateServiceServer) GetWorkflowTemplate(context.Context, *GetWorkflowTemplateRequest) (*WorkflowTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkflowTemplate not implemented")
}
func (UnimplementedWorkflowTemplateServiceServer) ListWorkflowTemplateVersions(context.Context, *ListWorkflowTemplateVersionsRequest) (*ListWorkflowTemplateVersionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkflowTemplateVersions not implemented")
}
func (UnimplementedWorkflowTemplateServiceServer) ListWorkflowTemplates(context.Context, *ListWorkflowTemplatesRequest) (*ListWorkflowTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkflowTemplates not implemented")
}
func (UnimplementedWorkflowTemplateServiceServer) CloneWorkflowTemplate(context.Context, *CloneWorkflowTemplateRequest) (*WorkflowTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloneWorkflowTemplate not implemented")
}
func (UnimplementedWorkflowTemplateServiceServer) ArchiveWorkflowTemplate(context.Context, *ArchiveWorkflowTemplateRequest) (*ArchiveWorkflowTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArchiveWorkflowTemplate not implemented")
}
func (UnimplementedWorkflowTemplateServiceServer) ListWorkflowTemplatesField(context.Context, *ListWorkflowTemplatesFieldRequest) (*ListWorkflowTemplatesFieldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkflowTemplatesField not implemented")
}
func (UnimplementedWorkflowTemplateServiceServer) mustEmbedUnimplementedWorkflowTemplateServiceServer() {
}

// UnsafeWorkflowTemplateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkflowTemplateServiceServer will
// result in compilation errors.
type UnsafeWorkflowTemplateServiceServer interface {
	mustEmbedUnimplementedWorkflowTemplateServiceServer()
}

func RegisterWorkflowTemplateServiceServer(s grpc.ServiceRegistrar, srv WorkflowTemplateServiceServer) {
	s.RegisterService(&_WorkflowTemplateService_serviceDesc, srv)
}

func _WorkflowTemplateService_GenerateWorkflowTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateWorkflowTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowTemplateServiceServer).GenerateWorkflowTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkflowTemplateService/GenerateWorkflowTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowTemplateServiceServer).GenerateWorkflowTemplate(ctx, req.(*GenerateWorkflowTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowTemplateService_CreateWorkflowTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkflowTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowTemplateServiceServer).CreateWorkflowTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkflowTemplateService/CreateWorkflowTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowTemplateServiceServer).CreateWorkflowTemplate(ctx, req.(*CreateWorkflowTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowTemplateService_CreateWorkflowTemplateVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkflowTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowTemplateServiceServer).CreateWorkflowTemplateVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkflowTemplateService/CreateWorkflowTemplateVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowTemplateServiceServer).CreateWorkflowTemplateVersion(ctx, req.(*CreateWorkflowTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowTemplateService_GetWorkflowTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkflowTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowTemplateServiceServer).GetWorkflowTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkflowTemplateService/GetWorkflowTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowTemplateServiceServer).GetWorkflowTemplate(ctx, req.(*GetWorkflowTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowTemplateService_ListWorkflowTemplateVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkflowTemplateVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowTemplateServiceServer).ListWorkflowTemplateVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkflowTemplateService/ListWorkflowTemplateVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowTemplateServiceServer).ListWorkflowTemplateVersions(ctx, req.(*ListWorkflowTemplateVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowTemplateService_ListWorkflowTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkflowTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowTemplateServiceServer).ListWorkflowTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkflowTemplateService/ListWorkflowTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowTemplateServiceServer).ListWorkflowTemplates(ctx, req.(*ListWorkflowTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowTemplateService_CloneWorkflowTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloneWorkflowTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowTemplateServiceServer).CloneWorkflowTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkflowTemplateService/CloneWorkflowTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowTemplateServiceServer).CloneWorkflowTemplate(ctx, req.(*CloneWorkflowTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowTemplateService_ArchiveWorkflowTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArchiveWorkflowTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowTemplateServiceServer).ArchiveWorkflowTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkflowTemplateService/ArchiveWorkflowTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowTemplateServiceServer).ArchiveWorkflowTemplate(ctx, req.(*ArchiveWorkflowTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowTemplateService_ListWorkflowTemplatesField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkflowTemplatesFieldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowTemplateServiceServer).ListWorkflowTemplatesField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkflowTemplateService/ListWorkflowTemplatesField",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowTemplateServiceServer).ListWorkflowTemplatesField(ctx, req.(*ListWorkflowTemplatesFieldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkflowTemplateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.WorkflowTemplateService",
	HandlerType: (*WorkflowTemplateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateWorkflowTemplate",
			Handler:    _WorkflowTemplateService_GenerateWorkflowTemplate_Handler,
		},
		{
			MethodName: "CreateWorkflowTemplate",
			Handler:    _WorkflowTemplateService_CreateWorkflowTemplate_Handler,
		},
		{
			MethodName: "CreateWorkflowTemplateVersion",
			Handler:    _WorkflowTemplateService_CreateWorkflowTemplateVersion_Handler,
		},
		{
			MethodName: "GetWorkflowTemplate",
			Handler:    _WorkflowTemplateService_GetWorkflowTemplate_Handler,
		},
		{
			MethodName: "ListWorkflowTemplateVersions",
			Handler:    _WorkflowTemplateService_ListWorkflowTemplateVersions_Handler,
		},
		{
			MethodName: "ListWorkflowTemplates",
			Handler:    _WorkflowTemplateService_ListWorkflowTemplates_Handler,
		},
		{
			MethodName: "CloneWorkflowTemplate",
			Handler:    _WorkflowTemplateService_CloneWorkflowTemplate_Handler,
		},
		{
			MethodName: "ArchiveWorkflowTemplate",
			Handler:    _WorkflowTemplateService_ArchiveWorkflowTemplate_Handler,
		},
		{
			MethodName: "ListWorkflowTemplatesField",
			Handler:    _WorkflowTemplateService_ListWorkflowTemplatesField_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workflow_template.proto",
}
