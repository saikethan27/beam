// Code generated by protoc-gen-go. DO NOT EDIT.
// source: beam_provision_api.proto

package fnexecution_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _struct "github.com/golang/protobuf/ptypes/struct"

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

// A request to get the provision info of a SDK harness worker instance.
type GetProvisionInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProvisionInfoRequest) Reset()         { *m = GetProvisionInfoRequest{} }
func (m *GetProvisionInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetProvisionInfoRequest) ProtoMessage()    {}
func (*GetProvisionInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_beam_provision_api_e6401c027f28403b, []int{0}
}
func (m *GetProvisionInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProvisionInfoRequest.Unmarshal(m, b)
}
func (m *GetProvisionInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProvisionInfoRequest.Marshal(b, m, deterministic)
}
func (dst *GetProvisionInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProvisionInfoRequest.Merge(dst, src)
}
func (m *GetProvisionInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetProvisionInfoRequest.Size(m)
}
func (m *GetProvisionInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProvisionInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProvisionInfoRequest proto.InternalMessageInfo

// A response containing the provision info of a SDK harness worker instance.
type GetProvisionInfoResponse struct {
	Info                 *ProvisionInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetProvisionInfoResponse) Reset()         { *m = GetProvisionInfoResponse{} }
func (m *GetProvisionInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetProvisionInfoResponse) ProtoMessage()    {}
func (*GetProvisionInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_beam_provision_api_e6401c027f28403b, []int{1}
}
func (m *GetProvisionInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProvisionInfoResponse.Unmarshal(m, b)
}
func (m *GetProvisionInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProvisionInfoResponse.Marshal(b, m, deterministic)
}
func (dst *GetProvisionInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProvisionInfoResponse.Merge(dst, src)
}
func (m *GetProvisionInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetProvisionInfoResponse.Size(m)
}
func (m *GetProvisionInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProvisionInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetProvisionInfoResponse proto.InternalMessageInfo

func (m *GetProvisionInfoResponse) GetInfo() *ProvisionInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

// Runtime provisioning information for a SDK harness worker instance,
// such as pipeline options, resource constraints and other job metadata
type ProvisionInfo struct {
	// (required) The job ID.
	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	// (required) The job name.
	JobName string `protobuf:"bytes,2,opt,name=job_name,json=jobName,proto3" json:"job_name,omitempty"`
	// (required) The worker ID. Often this will be the hostname.
	//
	// This is independent of the id passed to the SDK harness via the 'id'
	// argument in the Beam container contract.
	WorkerId string `protobuf:"bytes,5,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`
	// (required) Pipeline options. For non-template jobs, the options are
	// identical to what is passed to job submission.
	PipelineOptions *_struct.Struct `protobuf:"bytes,3,opt,name=pipeline_options,json=pipelineOptions,proto3" json:"pipeline_options,omitempty"`
	// (optional) Resource limits that the SDK harness worker should respect.
	// Runners may -- but are not required to -- enforce any limits provided.
	ResourceLimits *Resources `protobuf:"bytes,4,opt,name=resource_limits,json=resourceLimits,proto3" json:"resource_limits,omitempty"`
	// (required) The artifact retrieval token produced by
	// ArtifactStagingService.CommitManifestResponse.
	RetrievalToken       string   `protobuf:"bytes,6,opt,name=retrieval_token,json=retrievalToken,proto3" json:"retrieval_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProvisionInfo) Reset()         { *m = ProvisionInfo{} }
func (m *ProvisionInfo) String() string { return proto.CompactTextString(m) }
func (*ProvisionInfo) ProtoMessage()    {}
func (*ProvisionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_beam_provision_api_e6401c027f28403b, []int{2}
}
func (m *ProvisionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProvisionInfo.Unmarshal(m, b)
}
func (m *ProvisionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProvisionInfo.Marshal(b, m, deterministic)
}
func (dst *ProvisionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProvisionInfo.Merge(dst, src)
}
func (m *ProvisionInfo) XXX_Size() int {
	return xxx_messageInfo_ProvisionInfo.Size(m)
}
func (m *ProvisionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProvisionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProvisionInfo proto.InternalMessageInfo

func (m *ProvisionInfo) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *ProvisionInfo) GetJobName() string {
	if m != nil {
		return m.JobName
	}
	return ""
}

func (m *ProvisionInfo) GetWorkerId() string {
	if m != nil {
		return m.WorkerId
	}
	return ""
}

func (m *ProvisionInfo) GetPipelineOptions() *_struct.Struct {
	if m != nil {
		return m.PipelineOptions
	}
	return nil
}

func (m *ProvisionInfo) GetResourceLimits() *Resources {
	if m != nil {
		return m.ResourceLimits
	}
	return nil
}

func (m *ProvisionInfo) GetRetrievalToken() string {
	if m != nil {
		return m.RetrievalToken
	}
	return ""
}

// Resources specify limits for local resources, such memory and cpu. It
// is used to inform SDK harnesses of their allocated footprint.
type Resources struct {
	// (optional) Memory usage limits. SDKs can use this value to configure
	// internal buffer sizes and language specific sizes.
	Memory *Resources_Memory `protobuf:"bytes,1,opt,name=memory,proto3" json:"memory,omitempty"`
	// (optional) CPU usage limits.
	Cpu *Resources_Cpu `protobuf:"bytes,2,opt,name=cpu,proto3" json:"cpu,omitempty"`
	// (optional) Disk size limits for the semi-persistent location.
	SemiPersistentDisk   *Resources_Disk `protobuf:"bytes,3,opt,name=semi_persistent_disk,json=semiPersistentDisk,proto3" json:"semi_persistent_disk,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Resources) Reset()         { *m = Resources{} }
func (m *Resources) String() string { return proto.CompactTextString(m) }
func (*Resources) ProtoMessage()    {}
func (*Resources) Descriptor() ([]byte, []int) {
	return fileDescriptor_beam_provision_api_e6401c027f28403b, []int{3}
}
func (m *Resources) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resources.Unmarshal(m, b)
}
func (m *Resources) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resources.Marshal(b, m, deterministic)
}
func (dst *Resources) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resources.Merge(dst, src)
}
func (m *Resources) XXX_Size() int {
	return xxx_messageInfo_Resources.Size(m)
}
func (m *Resources) XXX_DiscardUnknown() {
	xxx_messageInfo_Resources.DiscardUnknown(m)
}

var xxx_messageInfo_Resources proto.InternalMessageInfo

func (m *Resources) GetMemory() *Resources_Memory {
	if m != nil {
		return m.Memory
	}
	return nil
}

func (m *Resources) GetCpu() *Resources_Cpu {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *Resources) GetSemiPersistentDisk() *Resources_Disk {
	if m != nil {
		return m.SemiPersistentDisk
	}
	return nil
}

// Memory limits.
type Resources_Memory struct {
	// (optional) Hard limit in bytes. A zero value means unspecified.
	Size                 uint64   `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resources_Memory) Reset()         { *m = Resources_Memory{} }
func (m *Resources_Memory) String() string { return proto.CompactTextString(m) }
func (*Resources_Memory) ProtoMessage()    {}
func (*Resources_Memory) Descriptor() ([]byte, []int) {
	return fileDescriptor_beam_provision_api_e6401c027f28403b, []int{3, 0}
}
func (m *Resources_Memory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resources_Memory.Unmarshal(m, b)
}
func (m *Resources_Memory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resources_Memory.Marshal(b, m, deterministic)
}
func (dst *Resources_Memory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resources_Memory.Merge(dst, src)
}
func (m *Resources_Memory) XXX_Size() int {
	return xxx_messageInfo_Resources_Memory.Size(m)
}
func (m *Resources_Memory) XXX_DiscardUnknown() {
	xxx_messageInfo_Resources_Memory.DiscardUnknown(m)
}

var xxx_messageInfo_Resources_Memory proto.InternalMessageInfo

func (m *Resources_Memory) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

// CPU limits.
type Resources_Cpu struct {
	// (optional) Shares of a cpu to use. Fractional values, such as "0.2"
	// or "2.5", are fine. Any value <= 0 means unspecified.
	Shares               float32  `protobuf:"fixed32,1,opt,name=shares,proto3" json:"shares,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resources_Cpu) Reset()         { *m = Resources_Cpu{} }
func (m *Resources_Cpu) String() string { return proto.CompactTextString(m) }
func (*Resources_Cpu) ProtoMessage()    {}
func (*Resources_Cpu) Descriptor() ([]byte, []int) {
	return fileDescriptor_beam_provision_api_e6401c027f28403b, []int{3, 1}
}
func (m *Resources_Cpu) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resources_Cpu.Unmarshal(m, b)
}
func (m *Resources_Cpu) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resources_Cpu.Marshal(b, m, deterministic)
}
func (dst *Resources_Cpu) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resources_Cpu.Merge(dst, src)
}
func (m *Resources_Cpu) XXX_Size() int {
	return xxx_messageInfo_Resources_Cpu.Size(m)
}
func (m *Resources_Cpu) XXX_DiscardUnknown() {
	xxx_messageInfo_Resources_Cpu.DiscardUnknown(m)
}

var xxx_messageInfo_Resources_Cpu proto.InternalMessageInfo

func (m *Resources_Cpu) GetShares() float32 {
	if m != nil {
		return m.Shares
	}
	return 0
}

// Disk limits.
type Resources_Disk struct {
	// (optional) Hard limit in bytes. A zero value means unspecified.
	Size                 uint64   `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resources_Disk) Reset()         { *m = Resources_Disk{} }
func (m *Resources_Disk) String() string { return proto.CompactTextString(m) }
func (*Resources_Disk) ProtoMessage()    {}
func (*Resources_Disk) Descriptor() ([]byte, []int) {
	return fileDescriptor_beam_provision_api_e6401c027f28403b, []int{3, 2}
}
func (m *Resources_Disk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resources_Disk.Unmarshal(m, b)
}
func (m *Resources_Disk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resources_Disk.Marshal(b, m, deterministic)
}
func (dst *Resources_Disk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resources_Disk.Merge(dst, src)
}
func (m *Resources_Disk) XXX_Size() int {
	return xxx_messageInfo_Resources_Disk.Size(m)
}
func (m *Resources_Disk) XXX_DiscardUnknown() {
	xxx_messageInfo_Resources_Disk.DiscardUnknown(m)
}

var xxx_messageInfo_Resources_Disk proto.InternalMessageInfo

func (m *Resources_Disk) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func init() {
	proto.RegisterType((*GetProvisionInfoRequest)(nil), "org.apache.beam.model.fn_execution.v1.GetProvisionInfoRequest")
	proto.RegisterType((*GetProvisionInfoResponse)(nil), "org.apache.beam.model.fn_execution.v1.GetProvisionInfoResponse")
	proto.RegisterType((*ProvisionInfo)(nil), "org.apache.beam.model.fn_execution.v1.ProvisionInfo")
	proto.RegisterType((*Resources)(nil), "org.apache.beam.model.fn_execution.v1.Resources")
	proto.RegisterType((*Resources_Memory)(nil), "org.apache.beam.model.fn_execution.v1.Resources.Memory")
	proto.RegisterType((*Resources_Cpu)(nil), "org.apache.beam.model.fn_execution.v1.Resources.Cpu")
	proto.RegisterType((*Resources_Disk)(nil), "org.apache.beam.model.fn_execution.v1.Resources.Disk")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProvisionServiceClient is the client API for ProvisionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProvisionServiceClient interface {
	// Get provision information for the SDK harness worker instance.
	GetProvisionInfo(ctx context.Context, in *GetProvisionInfoRequest, opts ...grpc.CallOption) (*GetProvisionInfoResponse, error)
}

type provisionServiceClient struct {
	cc *grpc.ClientConn
}

func NewProvisionServiceClient(cc *grpc.ClientConn) ProvisionServiceClient {
	return &provisionServiceClient{cc}
}

func (c *provisionServiceClient) GetProvisionInfo(ctx context.Context, in *GetProvisionInfoRequest, opts ...grpc.CallOption) (*GetProvisionInfoResponse, error) {
	out := new(GetProvisionInfoResponse)
	err := c.cc.Invoke(ctx, "/org.apache.beam.model.fn_execution.v1.ProvisionService/GetProvisionInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProvisionServiceServer is the server API for ProvisionService service.
type ProvisionServiceServer interface {
	// Get provision information for the SDK harness worker instance.
	GetProvisionInfo(context.Context, *GetProvisionInfoRequest) (*GetProvisionInfoResponse, error)
}

func RegisterProvisionServiceServer(s *grpc.Server, srv ProvisionServiceServer) {
	s.RegisterService(&_ProvisionService_serviceDesc, srv)
}

func _ProvisionService_GetProvisionInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProvisionInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvisionServiceServer).GetProvisionInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.apache.beam.model.fn_execution.v1.ProvisionService/GetProvisionInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvisionServiceServer).GetProvisionInfo(ctx, req.(*GetProvisionInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProvisionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "org.apache.beam.model.fn_execution.v1.ProvisionService",
	HandlerType: (*ProvisionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProvisionInfo",
			Handler:    _ProvisionService_GetProvisionInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "beam_provision_api.proto",
}

func init() {
	proto.RegisterFile("beam_provision_api.proto", fileDescriptor_beam_provision_api_e6401c027f28403b)
}

var fileDescriptor_beam_provision_api_e6401c027f28403b = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x95, 0x43, 0x4d, 0x33, 0x40, 0x1a, 0xad, 0x80, 0xba, 0xa6, 0x48, 0x28, 0x02, 0xc1,
	0xd5, 0x96, 0x16, 0x10, 0x77, 0x20, 0xd2, 0x0a, 0x88, 0x04, 0xb4, 0x72, 0xb9, 0x81, 0x1b, 0xcb,
	0x87, 0x49, 0xba, 0x49, 0xec, 0x59, 0x76, 0xd7, 0xe1, 0xf0, 0x1a, 0xbc, 0x03, 0xe2, 0xc9, 0x78,
	0x0e, 0xb4, 0xeb, 0xc4, 0xb4, 0x40, 0xa5, 0x94, 0x3b, 0xfb, 0x9f, 0xfd, 0x3f, 0xcd, 0xfc, 0xbb,
	0x03, 0x7e, 0x82, 0x71, 0x1e, 0x49, 0x45, 0x73, 0xa1, 0x05, 0x15, 0x51, 0x2c, 0x05, 0x97, 0x8a,
	0x0c, 0xb1, 0xbb, 0xa4, 0xc6, 0x3c, 0x96, 0x71, 0x7a, 0x82, 0xdc, 0x1e, 0xe2, 0x39, 0x65, 0x38,
	0xe3, 0xa3, 0x22, 0xc2, 0xcf, 0x98, 0x96, 0x46, 0x50, 0xc1, 0xe7, 0xbb, 0xc1, 0xf6, 0x98, 0x68,
	0x3c, 0xc3, 0x1d, 0x67, 0x4a, 0xca, 0xd1, 0x8e, 0x36, 0xaa, 0x4c, 0x4d, 0x05, 0xe9, 0x6f, 0xc1,
	0xe6, 0x4b, 0x34, 0x47, 0x4b, 0xfc, 0xb0, 0x18, 0x51, 0x88, 0x1f, 0x4b, 0xd4, 0xa6, 0x9f, 0x81,
	0xff, 0x77, 0x49, 0x4b, 0x2a, 0x34, 0xb2, 0x57, 0xd0, 0x16, 0xc5, 0x88, 0xfc, 0xc6, 0xed, 0xc6,
	0xfd, 0xcb, 0x7b, 0x8f, 0xf8, 0x4a, 0xad, 0xf0, 0xb3, 0x2c, 0x47, 0xe8, 0x7f, 0x6f, 0xc2, 0xd5,
	0x33, 0x3a, 0xbb, 0x0e, 0xde, 0x84, 0x92, 0x48, 0x64, 0x8e, 0xde, 0x09, 0xd7, 0x26, 0x94, 0x0c,
	0x33, 0xb6, 0x05, 0xeb, 0x56, 0x2e, 0xe2, 0x1c, 0xfd, 0xa6, 0x2b, 0x5c, 0x9a, 0x50, 0xf2, 0x36,
	0xce, 0x91, 0xdd, 0x84, 0xce, 0x27, 0x52, 0x53, 0x54, 0xd6, 0xb4, 0xe6, 0x6a, 0xeb, 0x95, 0x30,
	0xcc, 0xd8, 0x00, 0x7a, 0x52, 0x48, 0x9c, 0x89, 0x02, 0x23, 0x92, 0xb6, 0x15, 0xed, 0xb7, 0x5c,
	0xdb, 0x9b, 0xbc, 0x8a, 0x86, 0x2f, 0xa3, 0xe1, 0xc7, 0x2e, 0x9a, 0x70, 0x63, 0x69, 0x38, 0xac,
	0xce, 0xb3, 0xf7, 0xb0, 0xa1, 0x50, 0x53, 0xa9, 0x52, 0x8c, 0x66, 0x22, 0x17, 0x46, 0xfb, 0x6d,
	0x87, 0x78, 0xb0, 0xe2, 0xe4, 0xe1, 0xc2, 0xad, 0xc3, 0xee, 0x12, 0xf4, 0xda, 0x71, 0xd8, 0x3d,
	0x8b, 0x36, 0x4a, 0xe0, 0x3c, 0x9e, 0x45, 0x86, 0xa6, 0x58, 0xf8, 0x9e, 0x9b, 0xa0, 0x5b, 0xcb,
	0xef, 0xac, 0xda, 0xff, 0xd9, 0x84, 0x4e, 0x8d, 0x61, 0x87, 0xe0, 0xe5, 0x98, 0x93, 0xfa, 0xb2,
	0xb8, 0x82, 0x27, 0x17, 0x6d, 0x84, 0xbf, 0x71, 0xf6, 0x70, 0x81, 0x61, 0x2f, 0xa0, 0x95, 0xca,
	0xd2, 0x25, 0xbb, 0xfa, 0x85, 0xfe, 0xa6, 0xed, 0xcb, 0x32, 0xb4, 0x00, 0x36, 0x86, 0x6b, 0x1a,
	0x73, 0x11, 0x49, 0x54, 0x5a, 0x68, 0x83, 0x85, 0x89, 0x32, 0xa1, 0xa7, 0x8b, 0xc8, 0x1f, 0x5f,
	0x18, 0x7c, 0x20, 0xf4, 0x34, 0x64, 0x16, 0x79, 0x54, 0x13, 0xad, 0x16, 0x6c, 0x83, 0x57, 0x8d,
	0xc0, 0x18, 0xb4, 0xb5, 0xf8, 0x8a, 0x2e, 0x89, 0x76, 0xe8, 0xbe, 0x83, 0x5b, 0xd0, 0xda, 0x97,
	0x25, 0xbb, 0x01, 0x9e, 0x3e, 0x89, 0x15, 0x6a, 0x57, 0x6c, 0x86, 0x8b, 0xbf, 0x20, 0x80, 0xb6,
	0x85, 0xfc, 0xcb, 0xba, 0xf7, 0xa3, 0x01, 0xbd, 0xfa, 0x45, 0x1e, 0xa3, 0x9a, 0x8b, 0x14, 0xd9,
	0xb7, 0x06, 0xf4, 0xfe, 0xdc, 0x06, 0xf6, 0x74, 0xc5, 0x69, 0xce, 0xd9, 0xb0, 0xe0, 0xd9, 0x7f,
	0xfb, 0xab, 0x35, 0x1c, 0x1c, 0xc0, 0x9d, 0xf3, 0x08, 0xa7, 0x01, 0x83, 0x2b, 0xb5, 0xfd, 0xb9,
	0x14, 0x1f, 0xba, 0xa7, 0xaa, 0xd1, 0x7c, 0x37, 0xf1, 0xdc, 0xfb, 0x7f, 0xf8, 0x2b, 0x00, 0x00,
	0xff, 0xff, 0xb1, 0xb0, 0xe9, 0xeb, 0x6b, 0x04, 0x00, 0x00,
}
