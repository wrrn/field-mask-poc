// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config/config.proto

package config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import field_mask "google.golang.org/genproto/protobuf/field_mask"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Config represents a tagging of an IP and if it is enabled.
type Config struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ConfigName           string   `protobuf:"bytes,2,opt,name=config_name,json=configName,proto3" json:"config_name,omitempty"`
	Ip                   []byte   `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Tags                 []string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_61daf4297fba73b7, []int{0}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (dst *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(dst, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Config) GetConfigName() string {
	if m != nil {
		return m.ConfigName
	}
	return ""
}

func (m *Config) GetIp() []byte {
	if m != nil {
		return m.Ip
	}
	return nil
}

func (m *Config) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

// Get all of the configs in a List.
// The fields in the field_mask will be the only ones that are returned.
type ListConfigsRequest struct {
	FieldMask            *field_mask.FieldMask `protobuf:"bytes,1,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListConfigsRequest) Reset()         { *m = ListConfigsRequest{} }
func (m *ListConfigsRequest) String() string { return proto.CompactTextString(m) }
func (*ListConfigsRequest) ProtoMessage()    {}
func (*ListConfigsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_61daf4297fba73b7, []int{1}
}
func (m *ListConfigsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListConfigsRequest.Unmarshal(m, b)
}
func (m *ListConfigsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListConfigsRequest.Marshal(b, m, deterministic)
}
func (dst *ListConfigsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListConfigsRequest.Merge(dst, src)
}
func (m *ListConfigsRequest) XXX_Size() int {
	return xxx_messageInfo_ListConfigsRequest.Size(m)
}
func (m *ListConfigsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListConfigsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListConfigsRequest proto.InternalMessageInfo

func (m *ListConfigsRequest) GetFieldMask() *field_mask.FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

// The list of configs that were requested
type ListConfigsResponse struct {
	Configs              []*Config `protobuf:"bytes,1,rep,name=configs,proto3" json:"configs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListConfigsResponse) Reset()         { *m = ListConfigsResponse{} }
func (m *ListConfigsResponse) String() string { return proto.CompactTextString(m) }
func (*ListConfigsResponse) ProtoMessage()    {}
func (*ListConfigsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_61daf4297fba73b7, []int{2}
}
func (m *ListConfigsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListConfigsResponse.Unmarshal(m, b)
}
func (m *ListConfigsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListConfigsResponse.Marshal(b, m, deterministic)
}
func (dst *ListConfigsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListConfigsResponse.Merge(dst, src)
}
func (m *ListConfigsResponse) XXX_Size() int {
	return xxx_messageInfo_ListConfigsResponse.Size(m)
}
func (m *ListConfigsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListConfigsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListConfigsResponse proto.InternalMessageInfo

func (m *ListConfigsResponse) GetConfigs() []*Config {
	if m != nil {
		return m.Configs
	}
	return nil
}

// The new config that will be created
type CreateConfigRequest struct {
	Config               *Config  `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateConfigRequest) Reset()         { *m = CreateConfigRequest{} }
func (m *CreateConfigRequest) String() string { return proto.CompactTextString(m) }
func (*CreateConfigRequest) ProtoMessage()    {}
func (*CreateConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_61daf4297fba73b7, []int{3}
}
func (m *CreateConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateConfigRequest.Unmarshal(m, b)
}
func (m *CreateConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateConfigRequest.Marshal(b, m, deterministic)
}
func (dst *CreateConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateConfigRequest.Merge(dst, src)
}
func (m *CreateConfigRequest) XXX_Size() int {
	return xxx_messageInfo_CreateConfigRequest.Size(m)
}
func (m *CreateConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateConfigRequest proto.InternalMessageInfo

func (m *CreateConfigRequest) GetConfig() *Config {
	if m != nil {
		return m.Config
	}
	return nil
}

// The created config. Everything will match the request except for the name.
type CreateConfigResponse struct {
	Config               *Config  `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateConfigResponse) Reset()         { *m = CreateConfigResponse{} }
func (m *CreateConfigResponse) String() string { return proto.CompactTextString(m) }
func (*CreateConfigResponse) ProtoMessage()    {}
func (*CreateConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_61daf4297fba73b7, []int{4}
}
func (m *CreateConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateConfigResponse.Unmarshal(m, b)
}
func (m *CreateConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateConfigResponse.Marshal(b, m, deterministic)
}
func (dst *CreateConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateConfigResponse.Merge(dst, src)
}
func (m *CreateConfigResponse) XXX_Size() int {
	return xxx_messageInfo_CreateConfigResponse.Size(m)
}
func (m *CreateConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateConfigResponse proto.InternalMessageInfo

func (m *CreateConfigResponse) GetConfig() *Config {
	if m != nil {
		return m.Config
	}
	return nil
}

// Update a configuration. Only the fields in the field mask will be updated.
// The name is used to identify the config that should updated.
type UpdateConfigRequest struct {
	Config               *Config               `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateConfigRequest) Reset()         { *m = UpdateConfigRequest{} }
func (m *UpdateConfigRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateConfigRequest) ProtoMessage()    {}
func (*UpdateConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_61daf4297fba73b7, []int{5}
}
func (m *UpdateConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateConfigRequest.Unmarshal(m, b)
}
func (m *UpdateConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateConfigRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateConfigRequest.Merge(dst, src)
}
func (m *UpdateConfigRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateConfigRequest.Size(m)
}
func (m *UpdateConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateConfigRequest proto.InternalMessageInfo

func (m *UpdateConfigRequest) GetConfig() *Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *UpdateConfigRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

// The updated configuration.
type UpdateConfigResponse struct {
	Config               *Config  `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateConfigResponse) Reset()         { *m = UpdateConfigResponse{} }
func (m *UpdateConfigResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateConfigResponse) ProtoMessage()    {}
func (*UpdateConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_61daf4297fba73b7, []int{6}
}
func (m *UpdateConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateConfigResponse.Unmarshal(m, b)
}
func (m *UpdateConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateConfigResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateConfigResponse.Merge(dst, src)
}
func (m *UpdateConfigResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateConfigResponse.Size(m)
}
func (m *UpdateConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateConfigResponse proto.InternalMessageInfo

func (m *UpdateConfigResponse) GetConfig() *Config {
	if m != nil {
		return m.Config
	}
	return nil
}

type DeleteConfigRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteConfigRequest) Reset()         { *m = DeleteConfigRequest{} }
func (m *DeleteConfigRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteConfigRequest) ProtoMessage()    {}
func (*DeleteConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_61daf4297fba73b7, []int{7}
}
func (m *DeleteConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteConfigRequest.Unmarshal(m, b)
}
func (m *DeleteConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteConfigRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteConfigRequest.Merge(dst, src)
}
func (m *DeleteConfigRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteConfigRequest.Size(m)
}
func (m *DeleteConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteConfigRequest proto.InternalMessageInfo

func (m *DeleteConfigRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Config)(nil), "poc.config.Config")
	proto.RegisterType((*ListConfigsRequest)(nil), "poc.config.ListConfigsRequest")
	proto.RegisterType((*ListConfigsResponse)(nil), "poc.config.ListConfigsResponse")
	proto.RegisterType((*CreateConfigRequest)(nil), "poc.config.CreateConfigRequest")
	proto.RegisterType((*CreateConfigResponse)(nil), "poc.config.CreateConfigResponse")
	proto.RegisterType((*UpdateConfigRequest)(nil), "poc.config.UpdateConfigRequest")
	proto.RegisterType((*UpdateConfigResponse)(nil), "poc.config.UpdateConfigResponse")
	proto.RegisterType((*DeleteConfigRequest)(nil), "poc.config.DeleteConfigRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConfigServiceClient is the client API for ConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfigServiceClient interface {
	// List all of the config.
	ListConfigs(ctx context.Context, in *ListConfigsRequest, opts ...grpc.CallOption) (*ListConfigsResponse, error)
	// Create a new config.
	CreateConfig(ctx context.Context, in *CreateConfigRequest, opts ...grpc.CallOption) (*CreateConfigResponse, error)
	// Update a config.
	UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...grpc.CallOption) (*UpdateConfigResponse, error)
	// Delete a config.
	DeleteConfig(ctx context.Context, in *DeleteConfigRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type configServiceClient struct {
	cc *grpc.ClientConn
}

func NewConfigServiceClient(cc *grpc.ClientConn) ConfigServiceClient {
	return &configServiceClient{cc}
}

func (c *configServiceClient) ListConfigs(ctx context.Context, in *ListConfigsRequest, opts ...grpc.CallOption) (*ListConfigsResponse, error) {
	out := new(ListConfigsResponse)
	err := c.cc.Invoke(ctx, "/poc.config.ConfigService/ListConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) CreateConfig(ctx context.Context, in *CreateConfigRequest, opts ...grpc.CallOption) (*CreateConfigResponse, error) {
	out := new(CreateConfigResponse)
	err := c.cc.Invoke(ctx, "/poc.config.ConfigService/CreateConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...grpc.CallOption) (*UpdateConfigResponse, error) {
	out := new(UpdateConfigResponse)
	err := c.cc.Invoke(ctx, "/poc.config.ConfigService/UpdateConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) DeleteConfig(ctx context.Context, in *DeleteConfigRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/poc.config.ConfigService/DeleteConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigServiceServer is the server API for ConfigService service.
type ConfigServiceServer interface {
	// List all of the config.
	ListConfigs(context.Context, *ListConfigsRequest) (*ListConfigsResponse, error)
	// Create a new config.
	CreateConfig(context.Context, *CreateConfigRequest) (*CreateConfigResponse, error)
	// Update a config.
	UpdateConfig(context.Context, *UpdateConfigRequest) (*UpdateConfigResponse, error)
	// Delete a config.
	DeleteConfig(context.Context, *DeleteConfigRequest) (*empty.Empty, error)
}

func RegisterConfigServiceServer(s *grpc.Server, srv ConfigServiceServer) {
	s.RegisterService(&_ConfigService_serviceDesc, srv)
}

func _ConfigService_ListConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).ListConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poc.config.ConfigService/ListConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).ListConfigs(ctx, req.(*ListConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_CreateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).CreateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poc.config.ConfigService/CreateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).CreateConfig(ctx, req.(*CreateConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_UpdateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).UpdateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poc.config.ConfigService/UpdateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).UpdateConfig(ctx, req.(*UpdateConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_DeleteConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).DeleteConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poc.config.ConfigService/DeleteConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).DeleteConfig(ctx, req.(*DeleteConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConfigService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "poc.config.ConfigService",
	HandlerType: (*ConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListConfigs",
			Handler:    _ConfigService_ListConfigs_Handler,
		},
		{
			MethodName: "CreateConfig",
			Handler:    _ConfigService_CreateConfig_Handler,
		},
		{
			MethodName: "UpdateConfig",
			Handler:    _ConfigService_UpdateConfig_Handler,
		},
		{
			MethodName: "DeleteConfig",
			Handler:    _ConfigService_DeleteConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "config/config.proto",
}

func init() { proto.RegisterFile("config/config.proto", fileDescriptor_config_61daf4297fba73b7) }

var fileDescriptor_config_61daf4297fba73b7 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4f, 0x4f, 0xfa, 0x40,
	0x14, 0xa4, 0x85, 0xf0, 0xfb, 0xf1, 0x8a, 0x1e, 0x5e, 0x89, 0x69, 0x6a, 0x22, 0xcd, 0x9e, 0xd0,
	0x98, 0x92, 0xe0, 0xc9, 0x78, 0x12, 0xd4, 0xc4, 0xc4, 0x7f, 0x29, 0xf1, 0xe2, 0x85, 0x14, 0xd8,
	0x36, 0x0d, 0x7f, 0x5a, 0xd9, 0x62, 0xe2, 0xc5, 0x8f, 0xe9, 0xe7, 0x31, 0xdd, 0x57, 0x42, 0x0b,
	0x25, 0x46, 0x3d, 0xf1, 0xb2, 0x33, 0xcc, 0x9b, 0x9d, 0xd9, 0x82, 0x3e, 0x0a, 0xe7, 0x5e, 0xe0,
	0xb7, 0xe9, 0xc7, 0x8e, 0x16, 0x61, 0x1c, 0x22, 0x44, 0xe1, 0xc8, 0xa6, 0x13, 0xd3, 0xf2, 0xc3,
	0xd0, 0x9f, 0xf2, 0xb6, 0x44, 0x86, 0x4b, 0xaf, 0xed, 0x05, 0x7c, 0x3a, 0x1e, 0xcc, 0x5c, 0x31,
	0x21, 0xb6, 0x79, 0xb8, 0xc9, 0xe0, 0xb3, 0x28, 0x7e, 0x27, 0x90, 0xb9, 0x50, 0xed, 0x49, 0x21,
	0x44, 0xa8, 0xcc, 0xdd, 0x19, 0x37, 0x14, 0x4b, 0x69, 0xd5, 0x1c, 0x39, 0x63, 0x13, 0x34, 0x5a,
	0x33, 0x90, 0x90, 0x2a, 0x21, 0xa0, 0xa3, 0x87, 0x84, 0xb0, 0x0f, 0x6a, 0x10, 0x19, 0x65, 0x4b,
	0x69, 0xd5, 0x1d, 0x35, 0x88, 0x12, 0x91, 0xd8, 0xf5, 0x85, 0x51, 0xb1, 0xca, 0x89, 0x48, 0x32,
	0xb3, 0x47, 0xc0, 0xbb, 0x40, 0xc4, 0xb4, 0x46, 0x38, 0xfc, 0x75, 0xc9, 0x45, 0x8c, 0xe7, 0x00,
	0x6b, 0xa7, 0x72, 0xa9, 0xd6, 0x31, 0x6d, 0xb2, 0x6a, 0xaf, 0xac, 0xda, 0x37, 0x09, 0xe5, 0xde,
	0x15, 0x13, 0xa7, 0xe6, 0xad, 0x46, 0xd6, 0x03, 0x3d, 0x27, 0x28, 0xa2, 0x70, 0x2e, 0x38, 0x9e,
	0xc2, 0x3f, 0x72, 0x26, 0x0c, 0xc5, 0x2a, 0xb7, 0xb4, 0x0e, 0xda, 0xeb, 0x9c, 0x6c, 0x62, 0x3b,
	0x2b, 0x0a, 0xbb, 0x04, 0xbd, 0xb7, 0xe0, 0x6e, 0xcc, 0x53, 0x20, 0xb5, 0x75, 0x02, 0x55, 0x62,
	0xa4, 0x96, 0x8a, 0x34, 0x52, 0x06, 0xeb, 0x42, 0x23, 0x2f, 0x91, 0x1a, 0xf9, 0x89, 0xc6, 0x07,
	0xe8, 0xcf, 0xd1, 0xf8, 0x2f, 0x36, 0xf0, 0x02, 0xb4, 0xa5, 0x94, 0xa0, 0x28, 0xd5, 0x6f, 0xa3,
	0x04, 0xa2, 0xcb, 0x2c, 0xbb, 0xd0, 0xc8, 0xef, 0xff, 0xc5, 0x1d, 0x8e, 0x41, 0xbf, 0xe2, 0x53,
	0xbe, 0x79, 0x87, 0x82, 0x07, 0xd5, 0xf9, 0x54, 0x61, 0x8f, 0x58, 0x7d, 0xbe, 0x78, 0x0b, 0x46,
	0x1c, 0x9f, 0x40, 0xcb, 0x94, 0x89, 0x47, 0xd9, 0x3d, 0xdb, 0xcf, 0xc6, 0x6c, 0xee, 0xc4, 0xc9,
	0x38, 0x2b, 0x61, 0x1f, 0xea, 0xd9, 0x5a, 0x30, 0xf7, 0x97, 0x82, 0xce, 0x4d, 0x6b, 0x37, 0x21,
	0x2b, 0x9a, 0xcd, 0x29, 0x2f, 0x5a, 0xd0, 0x60, 0x5e, 0xb4, 0x28, 0x62, 0x56, 0xc2, 0x5b, 0xa8,
	0x67, 0x83, 0xcb, 0x8b, 0x16, 0x44, 0x6a, 0x1e, 0x6c, 0xb5, 0x7a, 0x9d, 0x7c, 0xcb, 0xac, 0xd4,
	0xfd, 0xff, 0x92, 0xb6, 0x31, 0xac, 0x4a, 0xec, 0xec, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x92, 0xf4,
	0xd1, 0x64, 0x3a, 0x04, 0x00, 0x00,
}
