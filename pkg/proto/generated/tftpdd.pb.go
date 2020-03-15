// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tftpdd.proto

package TFTPDD

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type TFTPD struct {
	Device               string   `protobuf:"bytes,1,opt,name=Device,proto3" json:"Device,omitempty"`
	Port                 int64    `protobuf:"varint,2,opt,name=Port,proto3" json:"Port,omitempty"`
	PXEPackageURL        string   `protobuf:"bytes,3,opt,name=PXEPackageURL,proto3" json:"PXEPackageURL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TFTPD) Reset()         { *m = TFTPD{} }
func (m *TFTPD) String() string { return proto.CompactTextString(m) }
func (*TFTPD) ProtoMessage()    {}
func (*TFTPD) Descriptor() ([]byte, []int) {
	return fileDescriptor_e89185584b3bfa22, []int{0}
}

func (m *TFTPD) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TFTPD.Unmarshal(m, b)
}
func (m *TFTPD) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TFTPD.Marshal(b, m, deterministic)
}
func (m *TFTPD) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TFTPD.Merge(m, src)
}
func (m *TFTPD) XXX_Size() int {
	return xxx_messageInfo_TFTPD.Size(m)
}
func (m *TFTPD) XXX_DiscardUnknown() {
	xxx_messageInfo_TFTPD.DiscardUnknown(m)
}

var xxx_messageInfo_TFTPD proto.InternalMessageInfo

func (m *TFTPD) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *TFTPD) GetPort() int64 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *TFTPD) GetPXEPackageURL() string {
	if m != nil {
		return m.PXEPackageURL
	}
	return ""
}

type TFTPDId struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TFTPDId) Reset()         { *m = TFTPDId{} }
func (m *TFTPDId) String() string { return proto.CompactTextString(m) }
func (*TFTPDId) ProtoMessage()    {}
func (*TFTPDId) Descriptor() ([]byte, []int) {
	return fileDescriptor_e89185584b3bfa22, []int{1}
}

func (m *TFTPDId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TFTPDId.Unmarshal(m, b)
}
func (m *TFTPDId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TFTPDId.Marshal(b, m, deterministic)
}
func (m *TFTPDId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TFTPDId.Merge(m, src)
}
func (m *TFTPDId) XXX_Size() int {
	return xxx_messageInfo_TFTPDId.Size(m)
}
func (m *TFTPDId) XXX_DiscardUnknown() {
	xxx_messageInfo_TFTPDId.DiscardUnknown(m)
}

var xxx_messageInfo_TFTPDId proto.InternalMessageInfo

func (m *TFTPDId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*TFTPD)(nil), "TFTPDD.TFTPD")
	proto.RegisterType((*TFTPDId)(nil), "TFTPDD.TFTPDId")
}

func init() {
	proto.RegisterFile("tftpdd.proto", fileDescriptor_e89185584b3bfa22)
}

var fileDescriptor_e89185584b3bfa22 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x49, 0x2b, 0x29,
	0x48, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x0b, 0x71, 0x0b, 0x09, 0x70, 0x71,
	0x51, 0x8a, 0xe4, 0x62, 0x05, 0xb3, 0x84, 0xc4, 0xb8, 0xd8, 0x5c, 0x52, 0xcb, 0x32, 0x93, 0x53,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xa0, 0x3c, 0x21, 0x21, 0x2e, 0x96, 0x80, 0xfc, 0xa2,
	0x12, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x30, 0x5b, 0x48, 0x85, 0x8b, 0x37, 0x20, 0xc2,
	0x35, 0x20, 0x31, 0x39, 0x3b, 0x31, 0x3d, 0x35, 0x34, 0xc8, 0x47, 0x82, 0x19, 0xac, 0x05, 0x55,
	0x50, 0x49, 0x92, 0x8b, 0x1d, 0x6c, 0xb4, 0x67, 0x8a, 0x10, 0x1f, 0x17, 0x93, 0x67, 0x0a, 0xd4,
	0x60, 0x26, 0xcf, 0x14, 0x23, 0x6b, 0x2e, 0x5e, 0x88, 0xfd, 0xbe, 0x89, 0x79, 0x89, 0xe9, 0xa9,
	0x45, 0x42, 0x5a, 0x5c, 0x6c, 0xce, 0x45, 0xa9, 0x89, 0x25, 0xa9, 0x42, 0xbc, 0x7a, 0x10, 0x19,
	0x08, 0x25, 0xc5, 0x8f, 0xc2, 0xf5, 0x4c, 0x51, 0x62, 0x48, 0x62, 0x03, 0xfb, 0xc0, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x16, 0x4b, 0xae, 0x99, 0xd1, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TFTPDDManagerClient is the client API for TFTPDDManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TFTPDDManagerClient interface {
	Create(ctx context.Context, in *TFTPD, opts ...grpc.CallOption) (*TFTPDId, error)
}

type tFTPDDManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewTFTPDDManagerClient(cc grpc.ClientConnInterface) TFTPDDManagerClient {
	return &tFTPDDManagerClient{cc}
}

func (c *tFTPDDManagerClient) Create(ctx context.Context, in *TFTPD, opts ...grpc.CallOption) (*TFTPDId, error) {
	out := new(TFTPDId)
	err := c.cc.Invoke(ctx, "/TFTPDD.TFTPDDManager/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TFTPDDManagerServer is the server API for TFTPDDManager service.
type TFTPDDManagerServer interface {
	Create(context.Context, *TFTPD) (*TFTPDId, error)
}

// UnimplementedTFTPDDManagerServer can be embedded to have forward compatible implementations.
type UnimplementedTFTPDDManagerServer struct {
}

func (*UnimplementedTFTPDDManagerServer) Create(ctx context.Context, req *TFTPD) (*TFTPDId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func RegisterTFTPDDManagerServer(s *grpc.Server, srv TFTPDDManagerServer) {
	s.RegisterService(&_TFTPDDManager_serviceDesc, srv)
}

func _TFTPDDManager_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TFTPD)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TFTPDDManagerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TFTPDD.TFTPDDManager/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TFTPDDManagerServer).Create(ctx, req.(*TFTPD))
	}
	return interceptor(ctx, in, info, handler)
}

var _TFTPDDManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "TFTPDD.TFTPDDManager",
	HandlerType: (*TFTPDDManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TFTPDDManager_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tftpdd.proto",
}