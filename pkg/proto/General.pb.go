// Code generated by protoc-gen-go. DO NOT EDIT.
// source: General.proto

package lab2

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

type Book struct {
	Chunks               []*Chunk `protobuf:"bytes,1,rep,name=chunks,proto3" json:"chunks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad40bea04b045104, []int{0}
}

func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (m *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(m, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetChunks() []*Chunk {
	if m != nil {
		return m.Chunks
	}
	return nil
}

type Chunk struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chunk) Reset()         { *m = Chunk{} }
func (m *Chunk) String() string { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()    {}
func (*Chunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad40bea04b045104, []int{1}
}

func (m *Chunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chunk.Unmarshal(m, b)
}
func (m *Chunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chunk.Marshal(b, m, deterministic)
}
func (m *Chunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chunk.Merge(m, src)
}
func (m *Chunk) XXX_Size() int {
	return xxx_messageInfo_Chunk.Size(m)
}
func (m *Chunk) XXX_DiscardUnknown() {
	xxx_messageInfo_Chunk.DiscardUnknown(m)
}

var xxx_messageInfo_Chunk proto.InternalMessageInfo

func (m *Chunk) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Chunk) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Message struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad40bea04b045104, []int{2}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Proposal struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Chunk                *Chunk   `protobuf:"bytes,2,opt,name=chunk,proto3" json:"chunk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Proposal) Reset()         { *m = Proposal{} }
func (m *Proposal) String() string { return proto.CompactTextString(m) }
func (*Proposal) ProtoMessage()    {}
func (*Proposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad40bea04b045104, []int{3}
}

func (m *Proposal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proposal.Unmarshal(m, b)
}
func (m *Proposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proposal.Marshal(b, m, deterministic)
}
func (m *Proposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proposal.Merge(m, src)
}
func (m *Proposal) XXX_Size() int {
	return xxx_messageInfo_Proposal.Size(m)
}
func (m *Proposal) XXX_DiscardUnknown() {
	xxx_messageInfo_Proposal.DiscardUnknown(m)
}

var xxx_messageInfo_Proposal proto.InternalMessageInfo

func (m *Proposal) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Proposal) GetChunk() *Chunk {
	if m != nil {
		return m.Chunk
	}
	return nil
}

func init() {
	proto.RegisterType((*Book)(nil), "lab2.Book")
	proto.RegisterType((*Chunk)(nil), "lab2.Chunk")
	proto.RegisterType((*Message)(nil), "lab2.Message")
	proto.RegisterType((*Proposal)(nil), "lab2.Proposal")
}

func init() {
	proto.RegisterFile("General.proto", fileDescriptor_ad40bea04b045104)
}

var fileDescriptor_ad40bea04b045104 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x52, 0x41, 0x4b, 0xf3, 0x40,
	0x14, 0x6c, 0xfa, 0xa5, 0xfd, 0xe2, 0x6b, 0x5a, 0xca, 0x9e, 0x4a, 0x41, 0xd0, 0xf5, 0x12, 0x50,
	0xa3, 0x46, 0x05, 0x2f, 0x5e, 0x34, 0xe0, 0x45, 0x8b, 0x44, 0xc4, 0xf3, 0xc6, 0x2c, 0x35, 0x34,
	0xe6, 0x85, 0xec, 0x16, 0xfd, 0x4b, 0xfe, 0x1d, 0x7f, 0x91, 0x9b, 0x97, 0x54, 0x6b, 0x11, 0xaa,
	0xb7, 0xb7, 0x6f, 0x66, 0x76, 0x66, 0x96, 0x85, 0xfe, 0x95, 0xcc, 0x65, 0x29, 0x32, 0xbf, 0x28,
	0x51, 0x23, 0xb3, 0x33, 0x11, 0x07, 0x7c, 0x17, 0xec, 0x0b, 0xc4, 0x19, 0xdb, 0x81, 0xee, 0xe3,
	0xd3, 0x3c, 0x9f, 0xa9, 0x91, 0xb5, 0xf5, 0xcf, 0xeb, 0x05, 0x3d, 0xbf, 0x82, 0xfd, 0xcb, 0x6a,
	0x17, 0x35, 0x10, 0x3f, 0x80, 0x0e, 0x2d, 0x18, 0x03, 0x3b, 0x17, 0xcf, 0xd2, 0x70, 0x2d, 0x6f,
	0x23, 0xa2, 0xb9, 0xda, 0x25, 0x42, 0x8b, 0x51, 0xdb, 0xec, 0xdc, 0x88, 0x66, 0xbe, 0x09, 0xff,
	0x6f, 0xa4, 0x52, 0x62, 0x4a, 0xb0, 0x96, 0xaf, 0x7a, 0x21, 0xa9, 0x66, 0x7e, 0x0e, 0xce, 0x6d,
	0x89, 0x05, 0x2a, 0x91, 0xb1, 0x01, 0xb4, 0xd3, 0xa2, 0x41, 0xcd, 0xc4, 0xb6, 0xa1, 0x43, 0xae,
	0x74, 0xdf, 0x4a, 0x9e, 0x1a, 0x09, 0xde, 0x2d, 0x70, 0x42, 0x63, 0x33, 0xc1, 0x44, 0xb2, 0x53,
	0x18, 0x86, 0xa9, 0xd2, 0x65, 0x1a, 0xcf, 0xb5, 0x24, 0x9a, 0x62, 0x83, 0x5a, 0xb4, 0xf0, 0x18,
	0xf7, 0xeb, 0x73, 0x13, 0x89, 0xb7, 0x3c, 0x8b, 0xed, 0x01, 0xdc, 0x17, 0x19, 0x8a, 0x84, 0x5e,
	0x61, 0xd9, 0xe5, 0x27, 0xb6, 0x0f, 0x6e, 0x88, 0x2f, 0xf9, 0x27, 0xff, 0x3b, 0x65, 0xbc, 0x2c,
	0xe7, 0xad, 0x43, 0x8b, 0x1d, 0x81, 0x7b, 0x27, 0xf3, 0xe4, 0xab, 0xe4, 0xda, 0x40, 0xc1, 0x9b,
	0x29, 0x35, 0x31, 0xef, 0x49, 0xa5, 0xf6, 0xc1, 0x79, 0x28, 0x53, 0x2d, 0xaf, 0x71, 0xfa, 0x9b,
	0x32, 0x27, 0x6b, 0xec, 0x56, 0xce, 0x95, 0xc6, 0x84, 0x3c, 0x83, 0x61, 0xfd, 0x33, 0xb4, 0xfc,
	0x9b, 0x32, 0xee, 0xd2, 0x4f, 0x3a, 0xfe, 0x08, 0x00, 0x00, 0xff, 0xff, 0x79, 0x9a, 0x1b, 0xa3,
	0x5a, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DataNodeClient is the client API for DataNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataNodeClient interface {
	DistributeChunks(ctx context.Context, opts ...grpc.CallOption) (DataNode_DistributeChunksClient, error)
	UploadBook(ctx context.Context, opts ...grpc.CallOption) (DataNode_UploadBookClient, error)
	DownloadBook(ctx context.Context, in *Message, opts ...grpc.CallOption) (DataNode_DownloadBookClient, error)
	//Distribuido
	SendProposal(ctx context.Context, opts ...grpc.CallOption) (DataNode_SendProposalClient, error)
}

type dataNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewDataNodeClient(cc grpc.ClientConnInterface) DataNodeClient {
	return &dataNodeClient{cc}
}

func (c *dataNodeClient) DistributeChunks(ctx context.Context, opts ...grpc.CallOption) (DataNode_DistributeChunksClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DataNode_serviceDesc.Streams[0], "/lab2.DataNode/DistributeChunks", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataNodeDistributeChunksClient{stream}
	return x, nil
}

type DataNode_DistributeChunksClient interface {
	Send(*Proposal) error
	CloseAndRecv() (*Message, error)
	grpc.ClientStream
}

type dataNodeDistributeChunksClient struct {
	grpc.ClientStream
}

func (x *dataNodeDistributeChunksClient) Send(m *Proposal) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataNodeDistributeChunksClient) CloseAndRecv() (*Message, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataNodeClient) UploadBook(ctx context.Context, opts ...grpc.CallOption) (DataNode_UploadBookClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DataNode_serviceDesc.Streams[1], "/lab2.DataNode/UploadBook", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataNodeUploadBookClient{stream}
	return x, nil
}

type DataNode_UploadBookClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*Message, error)
	grpc.ClientStream
}

type dataNodeUploadBookClient struct {
	grpc.ClientStream
}

func (x *dataNodeUploadBookClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataNodeUploadBookClient) CloseAndRecv() (*Message, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataNodeClient) DownloadBook(ctx context.Context, in *Message, opts ...grpc.CallOption) (DataNode_DownloadBookClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DataNode_serviceDesc.Streams[2], "/lab2.DataNode/DownloadBook", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataNodeDownloadBookClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataNode_DownloadBookClient interface {
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type dataNodeDownloadBookClient struct {
	grpc.ClientStream
}

func (x *dataNodeDownloadBookClient) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataNodeClient) SendProposal(ctx context.Context, opts ...grpc.CallOption) (DataNode_SendProposalClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DataNode_serviceDesc.Streams[3], "/lab2.DataNode/SendProposal", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataNodeSendProposalClient{stream}
	return x, nil
}

type DataNode_SendProposalClient interface {
	Send(*Proposal) error
	CloseAndRecv() (*Message, error)
	grpc.ClientStream
}

type dataNodeSendProposalClient struct {
	grpc.ClientStream
}

func (x *dataNodeSendProposalClient) Send(m *Proposal) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataNodeSendProposalClient) CloseAndRecv() (*Message, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataNodeServer is the server API for DataNode service.
type DataNodeServer interface {
	DistributeChunks(DataNode_DistributeChunksServer) error
	UploadBook(DataNode_UploadBookServer) error
	DownloadBook(*Message, DataNode_DownloadBookServer) error
	//Distribuido
	SendProposal(DataNode_SendProposalServer) error
}

// UnimplementedDataNodeServer can be embedded to have forward compatible implementations.
type UnimplementedDataNodeServer struct {
}

func (*UnimplementedDataNodeServer) DistributeChunks(srv DataNode_DistributeChunksServer) error {
	return status.Errorf(codes.Unimplemented, "method DistributeChunks not implemented")
}
func (*UnimplementedDataNodeServer) UploadBook(srv DataNode_UploadBookServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadBook not implemented")
}
func (*UnimplementedDataNodeServer) DownloadBook(req *Message, srv DataNode_DownloadBookServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadBook not implemented")
}
func (*UnimplementedDataNodeServer) SendProposal(srv DataNode_SendProposalServer) error {
	return status.Errorf(codes.Unimplemented, "method SendProposal not implemented")
}

func RegisterDataNodeServer(s *grpc.Server, srv DataNodeServer) {
	s.RegisterService(&_DataNode_serviceDesc, srv)
}

func _DataNode_DistributeChunks_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataNodeServer).DistributeChunks(&dataNodeDistributeChunksServer{stream})
}

type DataNode_DistributeChunksServer interface {
	SendAndClose(*Message) error
	Recv() (*Proposal, error)
	grpc.ServerStream
}

type dataNodeDistributeChunksServer struct {
	grpc.ServerStream
}

func (x *dataNodeDistributeChunksServer) SendAndClose(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataNodeDistributeChunksServer) Recv() (*Proposal, error) {
	m := new(Proposal)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DataNode_UploadBook_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataNodeServer).UploadBook(&dataNodeUploadBookServer{stream})
}

type DataNode_UploadBookServer interface {
	SendAndClose(*Message) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type dataNodeUploadBookServer struct {
	grpc.ServerStream
}

func (x *dataNodeUploadBookServer) SendAndClose(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataNodeUploadBookServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DataNode_DownloadBook_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Message)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataNodeServer).DownloadBook(m, &dataNodeDownloadBookServer{stream})
}

type DataNode_DownloadBookServer interface {
	Send(*Chunk) error
	grpc.ServerStream
}

type dataNodeDownloadBookServer struct {
	grpc.ServerStream
}

func (x *dataNodeDownloadBookServer) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

func _DataNode_SendProposal_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataNodeServer).SendProposal(&dataNodeSendProposalServer{stream})
}

type DataNode_SendProposalServer interface {
	SendAndClose(*Message) error
	Recv() (*Proposal, error)
	grpc.ServerStream
}

type dataNodeSendProposalServer struct {
	grpc.ServerStream
}

func (x *dataNodeSendProposalServer) SendAndClose(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataNodeSendProposalServer) Recv() (*Proposal, error) {
	m := new(Proposal)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DataNode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lab2.DataNode",
	HandlerType: (*DataNodeServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DistributeChunks",
			Handler:       _DataNode_DistributeChunks_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "UploadBook",
			Handler:       _DataNode_UploadBook_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadBook",
			Handler:       _DataNode_DownloadBook_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendProposal",
			Handler:       _DataNode_SendProposal_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "General.proto",
}

// NameNodeClient is the client API for NameNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NameNodeClient interface {
	WriteLog(ctx context.Context, opts ...grpc.CallOption) (NameNode_WriteLogClient, error)
	//Centralizado
	SendProposal(ctx context.Context, opts ...grpc.CallOption) (NameNode_SendProposalClient, error)
	GenerateProposal(ctx context.Context, opts ...grpc.CallOption) (NameNode_GenerateProposalClient, error)
}

type nameNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewNameNodeClient(cc grpc.ClientConnInterface) NameNodeClient {
	return &nameNodeClient{cc}
}

func (c *nameNodeClient) WriteLog(ctx context.Context, opts ...grpc.CallOption) (NameNode_WriteLogClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NameNode_serviceDesc.Streams[0], "/lab2.NameNode/WriteLog", opts...)
	if err != nil {
		return nil, err
	}
	x := &nameNodeWriteLogClient{stream}
	return x, nil
}

type NameNode_WriteLogClient interface {
	Send(*Proposal) error
	CloseAndRecv() (*Message, error)
	grpc.ClientStream
}

type nameNodeWriteLogClient struct {
	grpc.ClientStream
}

func (x *nameNodeWriteLogClient) Send(m *Proposal) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nameNodeWriteLogClient) CloseAndRecv() (*Message, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nameNodeClient) SendProposal(ctx context.Context, opts ...grpc.CallOption) (NameNode_SendProposalClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NameNode_serviceDesc.Streams[1], "/lab2.NameNode/SendProposal", opts...)
	if err != nil {
		return nil, err
	}
	x := &nameNodeSendProposalClient{stream}
	return x, nil
}

type NameNode_SendProposalClient interface {
	Send(*Proposal) error
	Recv() (*Proposal, error)
	grpc.ClientStream
}

type nameNodeSendProposalClient struct {
	grpc.ClientStream
}

func (x *nameNodeSendProposalClient) Send(m *Proposal) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nameNodeSendProposalClient) Recv() (*Proposal, error) {
	m := new(Proposal)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nameNodeClient) GenerateProposal(ctx context.Context, opts ...grpc.CallOption) (NameNode_GenerateProposalClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NameNode_serviceDesc.Streams[2], "/lab2.NameNode/GenerateProposal", opts...)
	if err != nil {
		return nil, err
	}
	x := &nameNodeGenerateProposalClient{stream}
	return x, nil
}

type NameNode_GenerateProposalClient interface {
	Send(*Proposal) error
	Recv() (*Proposal, error)
	grpc.ClientStream
}

type nameNodeGenerateProposalClient struct {
	grpc.ClientStream
}

func (x *nameNodeGenerateProposalClient) Send(m *Proposal) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nameNodeGenerateProposalClient) Recv() (*Proposal, error) {
	m := new(Proposal)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NameNodeServer is the server API for NameNode service.
type NameNodeServer interface {
	WriteLog(NameNode_WriteLogServer) error
	//Centralizado
	SendProposal(NameNode_SendProposalServer) error
	GenerateProposal(NameNode_GenerateProposalServer) error
}

// UnimplementedNameNodeServer can be embedded to have forward compatible implementations.
type UnimplementedNameNodeServer struct {
}

func (*UnimplementedNameNodeServer) WriteLog(srv NameNode_WriteLogServer) error {
	return status.Errorf(codes.Unimplemented, "method WriteLog not implemented")
}
func (*UnimplementedNameNodeServer) SendProposal(srv NameNode_SendProposalServer) error {
	return status.Errorf(codes.Unimplemented, "method SendProposal not implemented")
}
func (*UnimplementedNameNodeServer) GenerateProposal(srv NameNode_GenerateProposalServer) error {
	return status.Errorf(codes.Unimplemented, "method GenerateProposal not implemented")
}

func RegisterNameNodeServer(s *grpc.Server, srv NameNodeServer) {
	s.RegisterService(&_NameNode_serviceDesc, srv)
}

func _NameNode_WriteLog_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NameNodeServer).WriteLog(&nameNodeWriteLogServer{stream})
}

type NameNode_WriteLogServer interface {
	SendAndClose(*Message) error
	Recv() (*Proposal, error)
	grpc.ServerStream
}

type nameNodeWriteLogServer struct {
	grpc.ServerStream
}

func (x *nameNodeWriteLogServer) SendAndClose(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nameNodeWriteLogServer) Recv() (*Proposal, error) {
	m := new(Proposal)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _NameNode_SendProposal_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NameNodeServer).SendProposal(&nameNodeSendProposalServer{stream})
}

type NameNode_SendProposalServer interface {
	Send(*Proposal) error
	Recv() (*Proposal, error)
	grpc.ServerStream
}

type nameNodeSendProposalServer struct {
	grpc.ServerStream
}

func (x *nameNodeSendProposalServer) Send(m *Proposal) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nameNodeSendProposalServer) Recv() (*Proposal, error) {
	m := new(Proposal)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _NameNode_GenerateProposal_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NameNodeServer).GenerateProposal(&nameNodeGenerateProposalServer{stream})
}

type NameNode_GenerateProposalServer interface {
	Send(*Proposal) error
	Recv() (*Proposal, error)
	grpc.ServerStream
}

type nameNodeGenerateProposalServer struct {
	grpc.ServerStream
}

func (x *nameNodeGenerateProposalServer) Send(m *Proposal) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nameNodeGenerateProposalServer) Recv() (*Proposal, error) {
	m := new(Proposal)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _NameNode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lab2.NameNode",
	HandlerType: (*NameNodeServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WriteLog",
			Handler:       _NameNode_WriteLog_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SendProposal",
			Handler:       _NameNode_SendProposal_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GenerateProposal",
			Handler:       _NameNode_GenerateProposal_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "General.proto",
}
