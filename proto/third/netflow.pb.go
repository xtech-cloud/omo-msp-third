// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/third/netflow.proto

package third

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type NetflowInfo struct {
	Uid                  string         `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64         `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Type                 uint32         `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Created              int64          `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Creator              string         `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	Name                 string         `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Quote                string         `protobuf:"bytes,7,opt,name=quote,proto3" json:"quote,omitempty"`
	Scene                string         `protobuf:"bytes,8,opt,name=scene,proto3" json:"scene,omitempty"`
	Size                 uint64         `protobuf:"varint,9,opt,name=size,proto3" json:"size,omitempty"`
	Template             string         `protobuf:"bytes,10,opt,name=template,proto3" json:"template,omitempty"`
	Target               string         `protobuf:"bytes,11,opt,name=target,proto3" json:"target,omitempty"`
	Contents             []*ContentInfo `protobuf:"bytes,21,rep,name=contents,proto3" json:"contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *NetflowInfo) Reset()         { *m = NetflowInfo{} }
func (m *NetflowInfo) String() string { return proto.CompactTextString(m) }
func (*NetflowInfo) ProtoMessage()    {}
func (*NetflowInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb66cca61b7c6432, []int{0}
}

func (m *NetflowInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetflowInfo.Unmarshal(m, b)
}
func (m *NetflowInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetflowInfo.Marshal(b, m, deterministic)
}
func (m *NetflowInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetflowInfo.Merge(m, src)
}
func (m *NetflowInfo) XXX_Size() int {
	return xxx_messageInfo_NetflowInfo.Size(m)
}
func (m *NetflowInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NetflowInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NetflowInfo proto.InternalMessageInfo

func (m *NetflowInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *NetflowInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *NetflowInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *NetflowInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *NetflowInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *NetflowInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetflowInfo) GetQuote() string {
	if m != nil {
		return m.Quote
	}
	return ""
}

func (m *NetflowInfo) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

func (m *NetflowInfo) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *NetflowInfo) GetTemplate() string {
	if m != nil {
		return m.Template
	}
	return ""
}

func (m *NetflowInfo) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *NetflowInfo) GetContents() []*ContentInfo {
	if m != nil {
		return m.Contents
	}
	return nil
}

type ContentInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Size                 uint64   `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Group                string   `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`
	Type                 uint32   `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	Children             []string `protobuf:"bytes,10,rep,name=children,proto3" json:"children,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContentInfo) Reset()         { *m = ContentInfo{} }
func (m *ContentInfo) String() string { return proto.CompactTextString(m) }
func (*ContentInfo) ProtoMessage()    {}
func (*ContentInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb66cca61b7c6432, []int{1}
}

func (m *ContentInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentInfo.Unmarshal(m, b)
}
func (m *ContentInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentInfo.Marshal(b, m, deterministic)
}
func (m *ContentInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentInfo.Merge(m, src)
}
func (m *ContentInfo) XXX_Size() int {
	return xxx_messageInfo_ContentInfo.Size(m)
}
func (m *ContentInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ContentInfo proto.InternalMessageInfo

func (m *ContentInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ContentInfo) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ContentInfo) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *ContentInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ContentInfo) GetChildren() []string {
	if m != nil {
		return m.Children
	}
	return nil
}

type ReqNetflowAdd struct {
	Type                 uint32         `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Scene                string         `protobuf:"bytes,4,opt,name=scene,proto3" json:"scene,omitempty"`
	Operator             string         `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	Quote                string         `protobuf:"bytes,6,opt,name=quote,proto3" json:"quote,omitempty"`
	Template             string         `protobuf:"bytes,7,opt,name=template,proto3" json:"template,omitempty"`
	Target               string         `protobuf:"bytes,8,opt,name=target,proto3" json:"target,omitempty"`
	Contents             []*ContentInfo `protobuf:"bytes,11,rep,name=contents,proto3" json:"contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReqNetflowAdd) Reset()         { *m = ReqNetflowAdd{} }
func (m *ReqNetflowAdd) String() string { return proto.CompactTextString(m) }
func (*ReqNetflowAdd) ProtoMessage()    {}
func (*ReqNetflowAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb66cca61b7c6432, []int{2}
}

func (m *ReqNetflowAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqNetflowAdd.Unmarshal(m, b)
}
func (m *ReqNetflowAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqNetflowAdd.Marshal(b, m, deterministic)
}
func (m *ReqNetflowAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqNetflowAdd.Merge(m, src)
}
func (m *ReqNetflowAdd) XXX_Size() int {
	return xxx_messageInfo_ReqNetflowAdd.Size(m)
}
func (m *ReqNetflowAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqNetflowAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqNetflowAdd proto.InternalMessageInfo

func (m *ReqNetflowAdd) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqNetflowAdd) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

func (m *ReqNetflowAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqNetflowAdd) GetQuote() string {
	if m != nil {
		return m.Quote
	}
	return ""
}

func (m *ReqNetflowAdd) GetTemplate() string {
	if m != nil {
		return m.Template
	}
	return ""
}

func (m *ReqNetflowAdd) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *ReqNetflowAdd) GetContents() []*ContentInfo {
	if m != nil {
		return m.Contents
	}
	return nil
}

type ReplyNetflowInfo struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Info                 *NetflowInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyNetflowInfo) Reset()         { *m = ReplyNetflowInfo{} }
func (m *ReplyNetflowInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyNetflowInfo) ProtoMessage()    {}
func (*ReplyNetflowInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb66cca61b7c6432, []int{3}
}

func (m *ReplyNetflowInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyNetflowInfo.Unmarshal(m, b)
}
func (m *ReplyNetflowInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyNetflowInfo.Marshal(b, m, deterministic)
}
func (m *ReplyNetflowInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyNetflowInfo.Merge(m, src)
}
func (m *ReplyNetflowInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyNetflowInfo.Size(m)
}
func (m *ReplyNetflowInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyNetflowInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyNetflowInfo proto.InternalMessageInfo

func (m *ReplyNetflowInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyNetflowInfo) GetInfo() *NetflowInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type ReplyNetflowList struct {
	Status               *ReplyStatus   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                uint32         `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Pages                uint32         `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	Page                 uint32         `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	Number               uint32         `protobuf:"varint,5,opt,name=number,proto3" json:"number,omitempty"`
	List                 []*NetflowInfo `protobuf:"bytes,16,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReplyNetflowList) Reset()         { *m = ReplyNetflowList{} }
func (m *ReplyNetflowList) String() string { return proto.CompactTextString(m) }
func (*ReplyNetflowList) ProtoMessage()    {}
func (*ReplyNetflowList) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb66cca61b7c6432, []int{4}
}

func (m *ReplyNetflowList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyNetflowList.Unmarshal(m, b)
}
func (m *ReplyNetflowList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyNetflowList.Marshal(b, m, deterministic)
}
func (m *ReplyNetflowList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyNetflowList.Merge(m, src)
}
func (m *ReplyNetflowList) XXX_Size() int {
	return xxx_messageInfo_ReplyNetflowList.Size(m)
}
func (m *ReplyNetflowList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyNetflowList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyNetflowList proto.InternalMessageInfo

func (m *ReplyNetflowList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyNetflowList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyNetflowList) GetPages() uint32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *ReplyNetflowList) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ReplyNetflowList) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *ReplyNetflowList) GetList() []*NetflowInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*NetflowInfo)(nil), "omo.msp.third.NetflowInfo")
	proto.RegisterType((*ContentInfo)(nil), "omo.msp.third.ContentInfo")
	proto.RegisterType((*ReqNetflowAdd)(nil), "omo.msp.third.ReqNetflowAdd")
	proto.RegisterType((*ReplyNetflowInfo)(nil), "omo.msp.third.ReplyNetflowInfo")
	proto.RegisterType((*ReplyNetflowList)(nil), "omo.msp.third.ReplyNetflowList")
}

func init() {
	proto.RegisterFile("proto/third/netflow.proto", fileDescriptor_bb66cca61b7c6432)
}

var fileDescriptor_bb66cca61b7c6432 = []byte{
	// 596 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x5f, 0x6b, 0xd4, 0x4e,
	0x14, 0x6d, 0xf6, 0x4f, 0xba, 0x7b, 0xf3, 0xdb, 0x52, 0x86, 0xfe, 0x64, 0x0c, 0x8a, 0x21, 0x4f,
	0xfb, 0x94, 0x42, 0x05, 0xdf, 0xdb, 0x82, 0x4b, 0x45, 0x2b, 0x4c, 0xf1, 0xc5, 0xb7, 0x34, 0xb9,
	0x6d, 0x07, 0x92, 0x4c, 0x9a, 0xdc, 0x54, 0xd6, 0x2f, 0xe6, 0xb7, 0xf0, 0x49, 0xf0, 0xeb, 0xc8,
	0x4c, 0xb2, 0x71, 0xca, 0xa6, 0xad, 0xfa, 0x76, 0xcf, 0xfd, 0x73, 0xe6, 0xde, 0x73, 0x42, 0xe0,
	0x79, 0x59, 0x29, 0x52, 0x87, 0x74, 0x23, 0xab, 0xf4, 0xb0, 0x40, 0xba, 0xca, 0xd4, 0x97, 0xc8,
	0xe4, 0xd8, 0x42, 0xe5, 0x2a, 0xca, 0xeb, 0x32, 0x32, 0x45, 0x9f, 0xdb, 0x9d, 0x89, 0xca, 0x73,
	0x55, 0xb4, 0x8d, 0xe1, 0xb7, 0x11, 0x78, 0xe7, 0xed, 0xe8, 0x59, 0x71, 0xa5, 0xd8, 0x3e, 0x8c,
	0x1b, 0x99, 0x72, 0x27, 0x70, 0x96, 0x73, 0xa1, 0x43, 0xb6, 0x07, 0x23, 0x99, 0xf2, 0x51, 0xe0,
	0x2c, 0x27, 0x62, 0x24, 0x53, 0xc6, 0x60, 0x42, 0xeb, 0x12, 0xf9, 0x38, 0x70, 0x96, 0x0b, 0x61,
	0x62, 0xc6, 0x61, 0x37, 0xa9, 0x30, 0x26, 0x4c, 0xf9, 0x24, 0x70, 0x96, 0x63, 0xb1, 0x81, 0x7d,
	0x45, 0x55, 0x7c, 0x6a, 0x38, 0x37, 0x50, 0xf3, 0x14, 0x71, 0x8e, 0xdc, 0x35, 0x69, 0x13, 0xb3,
	0x03, 0x98, 0xde, 0x36, 0x8a, 0x90, 0xef, 0x9a, 0x64, 0x0b, 0x74, 0xb6, 0x4e, 0xb0, 0x40, 0x3e,
	0x6b, 0xb3, 0x06, 0xe8, 0xf9, 0x5a, 0x7e, 0x45, 0x3e, 0x37, 0x9b, 0x99, 0x98, 0xf9, 0x30, 0x23,
	0xcc, 0xcb, 0x2c, 0x26, 0xe4, 0x60, 0x9a, 0x7b, 0xcc, 0x9e, 0x81, 0x4b, 0x71, 0x75, 0x8d, 0xc4,
	0x3d, 0x53, 0xe9, 0x10, 0x7b, 0x03, 0xb3, 0x44, 0x15, 0x84, 0x05, 0xd5, 0xfc, 0xff, 0x60, 0xbc,
	0xf4, 0x8e, 0xfc, 0xe8, 0x9e, 0x7a, 0xd1, 0x69, 0x5b, 0xd6, 0xfa, 0x88, 0xbe, 0x37, 0x5c, 0x83,
	0x67, 0x15, 0x06, 0x84, 0xdb, 0x2c, 0x38, 0xb2, 0x16, 0x3c, 0x80, 0xe9, 0x75, 0xa5, 0x9a, 0xd2,
	0xa8, 0x37, 0x17, 0x2d, 0xe8, 0x25, 0x9d, 0x58, 0x92, 0xfa, 0x30, 0x4b, 0x6e, 0x64, 0x96, 0x56,
	0x58, 0x70, 0x08, 0xc6, 0xfa, 0x94, 0x0d, 0x0e, 0x7f, 0x3a, 0xb0, 0x10, 0x78, 0xdb, 0xf9, 0x76,
	0x9c, 0x0e, 0x9b, 0xd2, 0xcb, 0x36, 0xb1, 0x65, 0xf3, 0x61, 0xa6, 0x4a, 0xac, 0x2c, 0x47, 0x7a,
	0xfc, 0x5b, 0x7e, 0xd7, 0x96, 0xdf, 0x16, 0x75, 0xf7, 0x41, 0x51, 0x67, 0x0f, 0x8a, 0xea, 0xfd,
	0x85, 0xa8, 0x77, 0xb0, 0x2f, 0xb0, 0xcc, 0xd6, 0xf6, 0x27, 0x79, 0x04, 0x6e, 0x4d, 0x31, 0x35,
	0xb5, 0x11, 0x77, 0x9b, 0xc9, 0x0c, 0x5c, 0x98, 0x0e, 0xd1, 0x75, 0xb2, 0x08, 0x26, 0xb2, 0xb8,
	0x52, 0x46, 0xfb, 0xed, 0x09, 0x8b, 0x5d, 0x98, 0xbe, 0xf0, 0xbb, 0x73, 0xff, 0xe1, 0xf7, 0xb2,
	0xa6, 0x7f, 0x7a, 0xf8, 0x00, 0xa6, 0xa4, 0x28, 0xce, 0xcc, 0xcb, 0x0b, 0xd1, 0x02, 0x9d, 0x2d,
	0xe3, 0x6b, 0xac, 0x3b, 0x7f, 0x5a, 0xa0, 0x4d, 0xd3, 0xc1, 0xc6, 0x76, 0x1d, 0x6b, 0x41, 0x8b,
	0x26, 0xbf, 0xc4, 0xd6, 0x9c, 0x85, 0xe8, 0x90, 0x3e, 0x28, 0x93, 0x35, 0xf1, 0xfd, 0x41, 0x31,
	0xef, 0x1d, 0xa4, 0xfb, 0x8e, 0x7e, 0x8c, 0x61, 0xaf, 0xcb, 0x5e, 0x60, 0x75, 0x27, 0x13, 0x64,
	0x67, 0xe0, 0x1e, 0xa7, 0xe9, 0xc7, 0x02, 0xd9, 0x8b, 0xad, 0x43, 0xac, 0x6f, 0xc9, 0x7f, 0x35,
	0x74, 0xa6, 0xf5, 0x42, 0xb8, 0xc3, 0x56, 0xe0, 0xae, 0x90, 0x34, 0xd5, 0xb6, 0x26, 0xb7, 0x0d,
	0xd6, 0xc6, 0xd6, 0x3f, 0x21, 0x3a, 0x07, 0x6f, 0x85, 0x74, 0xb2, 0x7e, 0x2b, 0x33, 0xc2, 0x6a,
	0x68, 0x31, 0xcd, 0xd6, 0x56, 0x1f, 0xe5, 0xd3, 0x86, 0x85, 0x3b, 0xec, 0x03, 0xfc, 0xb7, 0x42,
	0xd2, 0x9e, 0xc8, 0x9a, 0x64, 0xf2, 0x04, 0xe1, 0xcb, 0x87, 0x0c, 0x35, 0xc3, 0xe1, 0x0e, 0x7b,
	0x07, 0x7b, 0x9f, 0xca, 0x34, 0x26, 0x7c, 0x6a, 0xc3, 0xb6, 0xcb, 0xe7, 0x43, 0x84, 0xdd, 0xa9,
	0xa7, 0x30, 0x17, 0x98, 0xab, 0x3b, 0x7c, 0x4a, 0xb6, 0x47, 0x48, 0x4e, 0x16, 0x9f, 0x3d, 0xeb,
	0x57, 0x7e, 0xe9, 0x1a, 0xf0, 0xfa, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x20, 0x08, 0x50,
	0x0a, 0x06, 0x00, 0x00,
}
