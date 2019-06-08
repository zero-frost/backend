// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metric_service.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MetricsResponse_ErrorCode int32

const (
	MetricsResponse_INTERNAL_ERROR MetricsResponse_ErrorCode = 0
	MetricsResponse_ENCODING_ERROR MetricsResponse_ErrorCode = 1
)

var MetricsResponse_ErrorCode_name = map[int32]string{
	0: "INTERNAL_ERROR",
	1: "ENCODING_ERROR",
}

var MetricsResponse_ErrorCode_value = map[string]int32{
	"INTERNAL_ERROR": 0,
	"ENCODING_ERROR": 1,
}

func (x MetricsResponse_ErrorCode) String() string {
	return proto.EnumName(MetricsResponse_ErrorCode_name, int32(x))
}

func (MetricsResponse_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dd7ab6b9927adfe8, []int{1, 0}
}

type MetricsRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricsRequest) Reset()         { *m = MetricsRequest{} }
func (m *MetricsRequest) String() string { return proto.CompactTextString(m) }
func (*MetricsRequest) ProtoMessage()    {}
func (*MetricsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd7ab6b9927adfe8, []int{0}
}

func (m *MetricsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsRequest.Unmarshal(m, b)
}
func (m *MetricsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsRequest.Marshal(b, m, deterministic)
}
func (m *MetricsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsRequest.Merge(m, src)
}
func (m *MetricsRequest) XXX_Size() int {
	return xxx_messageInfo_MetricsRequest.Size(m)
}
func (m *MetricsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsRequest proto.InternalMessageInfo

func (m *MetricsRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

type MetricsResponse struct {
	ErrorCode            MetricsResponse_ErrorCode `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3,enum=auth.v1.MetricsResponse_ErrorCode" json:"error_code,omitempty"`
	Metrics              []*MetricsResponse_Metric `protobuf:"bytes,2,rep,name=metrics,proto3" json:"metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *MetricsResponse) Reset()         { *m = MetricsResponse{} }
func (m *MetricsResponse) String() string { return proto.CompactTextString(m) }
func (*MetricsResponse) ProtoMessage()    {}
func (*MetricsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd7ab6b9927adfe8, []int{1}
}

func (m *MetricsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsResponse.Unmarshal(m, b)
}
func (m *MetricsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsResponse.Marshal(b, m, deterministic)
}
func (m *MetricsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsResponse.Merge(m, src)
}
func (m *MetricsResponse) XXX_Size() int {
	return xxx_messageInfo_MetricsResponse.Size(m)
}
func (m *MetricsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsResponse proto.InternalMessageInfo

func (m *MetricsResponse) GetErrorCode() MetricsResponse_ErrorCode {
	if m != nil {
		return m.ErrorCode
	}
	return MetricsResponse_INTERNAL_ERROR
}

func (m *MetricsResponse) GetMetrics() []*MetricsResponse_Metric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

type MetricsResponse_Metric struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricsResponse_Metric) Reset()         { *m = MetricsResponse_Metric{} }
func (m *MetricsResponse_Metric) String() string { return proto.CompactTextString(m) }
func (*MetricsResponse_Metric) ProtoMessage()    {}
func (*MetricsResponse_Metric) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd7ab6b9927adfe8, []int{1, 0}
}

func (m *MetricsResponse_Metric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsResponse_Metric.Unmarshal(m, b)
}
func (m *MetricsResponse_Metric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsResponse_Metric.Marshal(b, m, deterministic)
}
func (m *MetricsResponse_Metric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsResponse_Metric.Merge(m, src)
}
func (m *MetricsResponse_Metric) XXX_Size() int {
	return xxx_messageInfo_MetricsResponse_Metric.Size(m)
}
func (m *MetricsResponse_Metric) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsResponse_Metric.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsResponse_Metric proto.InternalMessageInfo

func (m *MetricsResponse_Metric) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MetricsResponse_Metric) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *MetricsResponse_Metric) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterEnum("auth.v1.MetricsResponse_ErrorCode", MetricsResponse_ErrorCode_name, MetricsResponse_ErrorCode_value)
	proto.RegisterType((*MetricsRequest)(nil), "auth.v1.MetricsRequest")
	proto.RegisterType((*MetricsResponse)(nil), "auth.v1.MetricsResponse")
	proto.RegisterType((*MetricsResponse_Metric)(nil), "auth.v1.MetricsResponse.Metric")
}

func init() { proto.RegisterFile("metric_service.proto", fileDescriptor_dd7ab6b9927adfe8) }

var fileDescriptor_dd7ab6b9927adfe8 = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x41, 0x6f, 0xd3, 0x4e,
	0x10, 0xc5, 0xff, 0x76, 0xfa, 0x4f, 0x94, 0x2d, 0x84, 0x68, 0xa9, 0x20, 0x58, 0x48, 0x58, 0x3e,
	0x55, 0x11, 0xf1, 0x36, 0x6e, 0x2e, 0x94, 0x0b, 0xa1, 0x84, 0xaa, 0x12, 0xa4, 0xd2, 0xaa, 0x5c,
	0x38, 0x10, 0x6d, 0x9d, 0xc1, 0x36, 0x4a, 0x76, 0xcd, 0xee, 0xd8, 0x05, 0x8e, 0x9c, 0x39, 0xc1,
	0x8d, 0xaf, 0xc5, 0x8d, 0x23, 0xe2, 0x83, 0x20, 0xaf, 0xdd, 0x08, 0x11, 0xf5, 0xe4, 0x99, 0xb7,
	0xbf, 0xf5, 0x7b, 0x63, 0x0f, 0xd9, 0x5b, 0x03, 0xea, 0x2c, 0x5e, 0x18, 0xd0, 0x65, 0x16, 0x43,
	0x98, 0x6b, 0x85, 0x8a, 0x76, 0x44, 0x81, 0x69, 0x58, 0x8e, 0xbd, 0xfb, 0x89, 0x52, 0xc9, 0x0a,
	0x98, 0xc8, 0x33, 0x26, 0xa4, 0x54, 0x28, 0x30, 0x53, 0xd2, 0xd4, 0x98, 0xf7, 0xd0, 0x3e, 0xe2,
	0x51, 0x02, 0x72, 0x64, 0x2e, 0x45, 0x92, 0x80, 0x66, 0x2a, 0xb7, 0xc4, 0x36, 0x1d, 0x04, 0xa4,
	0xf7, 0xd2, 0x9a, 0x19, 0x0e, 0xef, 0x0b, 0x30, 0x48, 0xfb, 0xa4, 0x25, 0xf2, 0x6c, 0xe0, 0xf8,
	0xce, 0x7e, 0x97, 0x57, 0x65, 0xf0, 0xc5, 0x25, 0xb7, 0x36, 0x90, 0xc9, 0x95, 0x34, 0x40, 0xa7,
	0x84, 0x80, 0xd6, 0x4a, 0x2f, 0x62, 0xb5, 0x04, 0x0b, 0xf7, 0xa2, 0x20, 0x6c, 0x12, 0x86, 0xff,
	0xd0, 0xe1, 0xac, 0x42, 0x8f, 0xd5, 0x12, 0x78, 0x17, 0xae, 0x4a, 0xfa, 0x88, 0x74, 0xea, 0x39,
	0xcd, 0xc0, 0xf5, 0x5b, 0xfb, 0xbb, 0xd1, 0x83, 0x6b, 0xef, 0xd7, 0x3d, 0xbf, 0xe2, 0xbd, 0xe7,
	0xa4, 0x5d, 0x4b, 0x94, 0x92, 0x1d, 0x29, 0xd6, 0xd0, 0xc4, 0xb5, 0x75, 0xa5, 0xe1, 0xc7, 0x1c,
	0x06, 0x6e, 0xad, 0x55, 0x35, 0xdd, 0x23, 0xff, 0x97, 0x62, 0x55, 0xc0, 0xa0, 0x65, 0xc5, 0xba,
	0x09, 0x0e, 0x49, 0x77, 0x13, 0x8d, 0x52, 0xd2, 0x3b, 0x9d, 0x9f, 0xcf, 0xf8, 0x7c, 0xfa, 0x62,
	0x31, 0xe3, 0xfc, 0x8c, 0xf7, 0xff, 0xab, 0xb4, 0xd9, 0xfc, 0xf8, 0xec, 0xd9, 0xe9, 0xfc, 0xa4,
	0xd1, 0x9c, 0xe8, 0xcd, 0xc6, 0xfc, 0x9c, 0x90, 0x13, 0xc0, 0x26, 0x2c, 0xbd, 0xbb, 0x1d, 0xdf,
	0x7e, 0x51, 0x6f, 0x70, 0xdd, 0x5c, 0xc1, 0xed, 0xcf, 0x3f, 0x7e, 0x7f, 0x73, 0x6f, 0xd2, 0x5d,
	0x56, 0x8e, 0x59, 0x33, 0xdc, 0xd3, 0x5f, 0xce, 0xd7, 0xe9, 0x4f, 0x87, 0x22, 0xb9, 0x33, 0x2d,
	0x30, 0x05, 0x89, 0x59, 0x6c, 0x7f, 0x99, 0xdf, 0xac, 0x43, 0xf0, 0x8a, 0xdc, 0xa8, 0x5e, 0x38,
	0x6a, 0x7a, 0x3a, 0x4c, 0x11, 0x73, 0x73, 0xc4, 0x58, 0x92, 0x61, 0x5a, 0x5c, 0x84, 0xb1, 0x5a,
	0xb3, 0x4f, 0xa0, 0xd5, 0xe8, 0xad, 0x56, 0x06, 0xd9, 0xdf, 0xac, 0x77, 0xaf, 0x3a, 0x58, 0xd8,
	0x83, 0x27, 0x76, 0x03, 0xe4, 0x5a, 0x64, 0xab, 0xea, 0x4a, 0xd4, 0x1a, 0x87, 0x07, 0x43, 0xc7,
	0x89, 0xfa, 0x22, 0xcf, 0x57, 0x8d, 0x27, 0x7b, 0x67, 0x94, 0x3c, 0xda, 0x52, 0xf8, 0x63, 0xd2,
	0x9a, 0x1c, 0x4c, 0xe8, 0x84, 0x0c, 0x39, 0x60, 0xa1, 0x25, 0x2c, 0xfd, 0xcb, 0x14, 0xa4, 0x8f,
	0x29, 0xf8, 0x1a, 0x8c, 0x2a, 0x74, 0x0c, 0xfe, 0x52, 0x81, 0xf1, 0xa5, 0x42, 0x1f, 0x3e, 0x64,
	0x06, 0x43, 0xda, 0x26, 0x3b, 0xdf, 0x5d, 0xa7, 0xf3, 0xda, 0x2d, 0xc7, 0x17, 0x6d, 0x6b, 0x7e,
	0xf8, 0x27, 0x00, 0x00, 0xff, 0xff, 0xf6, 0xd6, 0x7f, 0x06, 0xeb, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MetricClient is the client API for Metric service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetricClient interface {
	GetMetrics(ctx context.Context, in *MetricsRequest, opts ...grpc.CallOption) (*MetricsResponse, error)
}

type metricClient struct {
	cc *grpc.ClientConn
}

func NewMetricClient(cc *grpc.ClientConn) MetricClient {
	return &metricClient{cc}
}

func (c *metricClient) GetMetrics(ctx context.Context, in *MetricsRequest, opts ...grpc.CallOption) (*MetricsResponse, error) {
	out := new(MetricsResponse)
	err := c.cc.Invoke(ctx, "/auth.v1.Metric/GetMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricServer is the server API for Metric service.
type MetricServer interface {
	GetMetrics(context.Context, *MetricsRequest) (*MetricsResponse, error)
}

// UnimplementedMetricServer can be embedded to have forward compatible implementations.
type UnimplementedMetricServer struct {
}

func (*UnimplementedMetricServer) GetMetrics(ctx context.Context, req *MetricsRequest) (*MetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetrics not implemented")
}

func RegisterMetricServer(s *grpc.Server, srv MetricServer) {
	s.RegisterService(&_Metric_serviceDesc, srv)
}

func _Metric_GetMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricServer).GetMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.v1.Metric/GetMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricServer).GetMetrics(ctx, req.(*MetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Metric_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.v1.Metric",
	HandlerType: (*MetricServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMetrics",
			Handler:    _Metric_GetMetrics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metric_service.proto",
}
