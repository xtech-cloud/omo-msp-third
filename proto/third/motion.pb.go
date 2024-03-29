// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/third/motion.proto

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

type MotionInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Type                 uint32   `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Created              int64    `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64    `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string   `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string   `protobuf:"bytes,7,opt,name=operator,proto3" json:"operator,omitempty"`
	Name                 string   `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	App                  string   `protobuf:"bytes,9,opt,name=app,proto3" json:"app,omitempty"`
	Sn                   string   `protobuf:"bytes,10,opt,name=sn,proto3" json:"sn,omitempty"`
	User                 string   `protobuf:"bytes,11,opt,name=user,proto3" json:"user,omitempty"`
	Event                string   `protobuf:"bytes,12,opt,name=event,proto3" json:"event,omitempty"`
	Scene                string   `protobuf:"bytes,13,opt,name=scene,proto3" json:"scene,omitempty"`
	Count                uint32   `protobuf:"varint,14,opt,name=count,proto3" json:"count,omitempty"`
	Content              string   `protobuf:"bytes,15,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MotionInfo) Reset()         { *m = MotionInfo{} }
func (m *MotionInfo) String() string { return proto.CompactTextString(m) }
func (*MotionInfo) ProtoMessage()    {}
func (*MotionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4acdfdad1db9312c, []int{0}
}

func (m *MotionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MotionInfo.Unmarshal(m, b)
}
func (m *MotionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MotionInfo.Marshal(b, m, deterministic)
}
func (m *MotionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MotionInfo.Merge(m, src)
}
func (m *MotionInfo) XXX_Size() int {
	return xxx_messageInfo_MotionInfo.Size(m)
}
func (m *MotionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MotionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MotionInfo proto.InternalMessageInfo

func (m *MotionInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *MotionInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MotionInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *MotionInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *MotionInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *MotionInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MotionInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *MotionInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MotionInfo) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

func (m *MotionInfo) GetSn() string {
	if m != nil {
		return m.Sn
	}
	return ""
}

func (m *MotionInfo) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *MotionInfo) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *MotionInfo) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

func (m *MotionInfo) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *MotionInfo) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type ReqMotionAdd struct {
	App                  string   `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	Sn                   string   `protobuf:"bytes,2,opt,name=sn,proto3" json:"sn,omitempty"`
	Event                string   `protobuf:"bytes,3,opt,name=event,proto3" json:"event,omitempty"`
	Scene                string   `protobuf:"bytes,4,opt,name=scene,proto3" json:"scene,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	Count                uint32   `protobuf:"varint,6,opt,name=count,proto3" json:"count,omitempty"`
	Content              string   `protobuf:"bytes,7,opt,name=content,proto3" json:"content,omitempty"`
	Type                 uint32   `protobuf:"varint,8,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqMotionAdd) Reset()         { *m = ReqMotionAdd{} }
func (m *ReqMotionAdd) String() string { return proto.CompactTextString(m) }
func (*ReqMotionAdd) ProtoMessage()    {}
func (*ReqMotionAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_4acdfdad1db9312c, []int{1}
}

func (m *ReqMotionAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqMotionAdd.Unmarshal(m, b)
}
func (m *ReqMotionAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqMotionAdd.Marshal(b, m, deterministic)
}
func (m *ReqMotionAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqMotionAdd.Merge(m, src)
}
func (m *ReqMotionAdd) XXX_Size() int {
	return xxx_messageInfo_ReqMotionAdd.Size(m)
}
func (m *ReqMotionAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqMotionAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqMotionAdd proto.InternalMessageInfo

func (m *ReqMotionAdd) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

func (m *ReqMotionAdd) GetSn() string {
	if m != nil {
		return m.Sn
	}
	return ""
}

func (m *ReqMotionAdd) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *ReqMotionAdd) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

func (m *ReqMotionAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqMotionAdd) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReqMotionAdd) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *ReqMotionAdd) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type ReplyMotionInfo struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Info                 *MotionInfo  `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyMotionInfo) Reset()         { *m = ReplyMotionInfo{} }
func (m *ReplyMotionInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyMotionInfo) ProtoMessage()    {}
func (*ReplyMotionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4acdfdad1db9312c, []int{2}
}

func (m *ReplyMotionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyMotionInfo.Unmarshal(m, b)
}
func (m *ReplyMotionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyMotionInfo.Marshal(b, m, deterministic)
}
func (m *ReplyMotionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyMotionInfo.Merge(m, src)
}
func (m *ReplyMotionInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyMotionInfo.Size(m)
}
func (m *ReplyMotionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyMotionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyMotionInfo proto.InternalMessageInfo

func (m *ReplyMotionInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyMotionInfo) GetInfo() *MotionInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type ReplyMotionList struct {
	Status               *ReplyStatus  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                uint32        `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Pages                uint32        `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	Page                 uint32        `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	Number               uint32        `protobuf:"varint,5,opt,name=number,proto3" json:"number,omitempty"`
	List                 []*MotionInfo `protobuf:"bytes,16,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReplyMotionList) Reset()         { *m = ReplyMotionList{} }
func (m *ReplyMotionList) String() string { return proto.CompactTextString(m) }
func (*ReplyMotionList) ProtoMessage()    {}
func (*ReplyMotionList) Descriptor() ([]byte, []int) {
	return fileDescriptor_4acdfdad1db9312c, []int{3}
}

func (m *ReplyMotionList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyMotionList.Unmarshal(m, b)
}
func (m *ReplyMotionList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyMotionList.Marshal(b, m, deterministic)
}
func (m *ReplyMotionList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyMotionList.Merge(m, src)
}
func (m *ReplyMotionList) XXX_Size() int {
	return xxx_messageInfo_ReplyMotionList.Size(m)
}
func (m *ReplyMotionList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyMotionList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyMotionList proto.InternalMessageInfo

func (m *ReplyMotionList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyMotionList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyMotionList) GetPages() uint32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *ReplyMotionList) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ReplyMotionList) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *ReplyMotionList) GetList() []*MotionInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*MotionInfo)(nil), "omo.msp.third.MotionInfo")
	proto.RegisterType((*ReqMotionAdd)(nil), "omo.msp.third.ReqMotionAdd")
	proto.RegisterType((*ReplyMotionInfo)(nil), "omo.msp.third.ReplyMotionInfo")
	proto.RegisterType((*ReplyMotionList)(nil), "omo.msp.third.ReplyMotionList")
}

func init() {
	proto.RegisterFile("proto/third/motion.proto", fileDescriptor_4acdfdad1db9312c)
}

var fileDescriptor_4acdfdad1db9312c = []byte{
	// 562 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x6a, 0xdb, 0x30,
	0x14, 0xae, 0xed, 0xc4, 0x49, 0x94, 0x38, 0x2d, 0xa2, 0x0c, 0x2d, 0xfb, 0x21, 0xf8, 0x2a, 0x37,
	0x4b, 0x21, 0x7b, 0x82, 0x76, 0xd0, 0xb0, 0xb1, 0x32, 0x50, 0xd9, 0xcd, 0xee, 0x5c, 0xfb, 0x74,
	0x13, 0xc4, 0x92, 0x6b, 0xc9, 0x81, 0xbc, 0xdb, 0x1e, 0x61, 0x30, 0xd8, 0x13, 0x0d, 0x1d, 0x39,
	0xce, 0x4f, 0xd3, 0x66, 0xec, 0xee, 0x7c, 0xe7, 0xe7, 0xd3, 0xf7, 0x1d, 0x09, 0x11, 0x56, 0x94,
	0xca, 0xa8, 0x0b, 0xf3, 0x43, 0x94, 0xd9, 0x45, 0xae, 0x8c, 0x50, 0x72, 0x8a, 0x29, 0x1a, 0xa9,
	0x5c, 0x4d, 0x73, 0x5d, 0x4c, 0xb1, 0x36, 0xda, 0x69, 0x4c, 0x55, 0x9e, 0xaf, 0x1b, 0xe3, 0x3f,
	0x3e, 0x21, 0x37, 0x38, 0xf9, 0x51, 0xde, 0x2b, 0x7a, 0x46, 0x82, 0x4a, 0x64, 0xcc, 0x1b, 0x7b,
	0x93, 0x1e, 0xb7, 0x21, 0x1d, 0x12, 0x5f, 0x64, 0xcc, 0x1f, 0x7b, 0x93, 0x16, 0xf7, 0x45, 0x46,
	0x29, 0x69, 0x99, 0x55, 0x01, 0x2c, 0x18, 0x7b, 0x93, 0x88, 0x63, 0x4c, 0x19, 0xe9, 0xa4, 0x25,
	0x24, 0x06, 0x32, 0xd6, 0x1a, 0x7b, 0x93, 0x80, 0xaf, 0xa1, 0xad, 0x54, 0x45, 0x86, 0x95, 0xb6,
	0xab, 0xd4, 0xb0, 0x99, 0x51, 0x25, 0x0b, 0xf1, 0xb4, 0x35, 0xa4, 0x23, 0xd2, 0x55, 0x05, 0x94,
	0x58, 0xea, 0x60, 0xa9, 0xc1, 0xf6, 0x74, 0x99, 0xe4, 0xc0, 0xba, 0x98, 0xc7, 0xd8, 0x6a, 0x4e,
	0x8a, 0x82, 0xf5, 0x9c, 0xe6, 0xa4, 0x28, 0xac, 0x66, 0x2d, 0x19, 0xc1, 0x84, 0xaf, 0xa5, 0x9d,
	0xaa, 0x34, 0x94, 0xac, 0xef, 0xa6, 0x6c, 0x4c, 0xcf, 0x49, 0x1b, 0x96, 0x20, 0x0d, 0x1b, 0x60,
	0xd2, 0x01, 0x9b, 0xd5, 0x29, 0x48, 0x60, 0x91, 0xcb, 0x22, 0xb0, 0xd9, 0x54, 0x55, 0xd2, 0xb0,
	0x21, 0x9a, 0x76, 0x00, 0x1d, 0x28, 0x69, 0x2c, 0xc7, 0x69, 0xed, 0xc0, 0xc1, 0xf8, 0xa7, 0x47,
	0x06, 0x1c, 0x1e, 0xdc, 0x5e, 0x2f, 0xb3, 0x6c, 0x2d, 0xd1, 0xdb, 0x97, 0xe8, 0x37, 0x12, 0x1b,
	0x39, 0xc1, 0x41, 0x39, 0xad, 0x6d, 0x39, 0xdb, 0x0b, 0x6a, 0xef, 0x2d, 0xa8, 0x91, 0x1a, 0x3e,
	0x21, 0xb5, 0xb3, 0x23, 0xb5, 0xb9, 0xce, 0xee, 0xe6, 0x3a, 0x63, 0x43, 0x4e, 0x39, 0x14, 0x8b,
	0xd5, 0xd6, 0xbb, 0x98, 0x91, 0x50, 0x9b, 0xc4, 0x54, 0x1a, 0x3d, 0xf4, 0x67, 0xa3, 0xe9, 0xce,
	0x03, 0x9b, 0x62, 0xff, 0x2d, 0x76, 0xf0, 0xba, 0x93, 0xbe, 0x23, 0x2d, 0x21, 0xef, 0x15, 0x9a,
	0xec, 0xcf, 0x5e, 0xee, 0x4d, 0x6c, 0xc8, 0x39, 0xb6, 0xc5, 0xbf, 0xbc, 0x9d, 0x63, 0x3f, 0x0b,
	0x6d, 0xfe, 0xeb, 0xd8, 0x73, 0xd2, 0x36, 0xca, 0x24, 0x0b, 0x3c, 0x37, 0xe2, 0x0e, 0xd8, 0x6c,
	0x91, 0x7c, 0x07, 0x5d, 0xbf, 0x5b, 0x07, 0xac, 0x7b, 0x1b, 0xe0, 0x7a, 0x23, 0x8e, 0x31, 0x7d,
	0x41, 0x42, 0x59, 0xe5, 0x77, 0xe0, 0x76, 0x1b, 0xf1, 0x1a, 0x59, 0x3b, 0x0b, 0xa1, 0x0d, 0x3b,
	0x1b, 0x07, 0x47, 0xec, 0xd8, 0xb6, 0xd9, 0xef, 0x80, 0x44, 0x2e, 0x79, 0x0b, 0xe5, 0x52, 0xa4,
	0x40, 0xe7, 0x24, 0xbc, 0xcc, 0xb2, 0x2f, 0x12, 0xe8, 0xab, 0x47, 0x36, 0x36, 0x6f, 0x65, 0xf4,
	0xf6, 0x90, 0xc7, 0x0d, 0x7d, 0x7c, 0x42, 0xaf, 0x49, 0x38, 0x07, 0x63, 0x89, 0x1e, 0xef, 0xe3,
	0xa1, 0x02, 0x6d, 0x6c, 0xdf, 0x3f, 0xf0, 0xdc, 0x90, 0xfe, 0x1c, 0xcc, 0xd5, 0xea, 0x5a, 0x2c,
	0x0c, 0x94, 0xf4, 0xf5, 0x61, 0x32, 0x57, 0x7d, 0x8e, 0xce, 0x5e, 0x15, 0xd2, 0x0d, 0xe6, 0x60,
	0xec, 0x6d, 0x08, 0x6d, 0x44, 0x7a, 0x84, 0xef, 0xcd, 0x53, 0x57, 0x89, 0xc3, 0xf1, 0x09, 0xfd,
	0x44, 0x86, 0x5f, 0xf1, 0xaf, 0x38, 0x26, 0xd0, 0x75, 0x8d, 0xd8, 0x21, 0xc2, 0xda, 0xe9, 0x07,
	0xd2, 0xe3, 0x90, 0xab, 0x25, 0x1c, 0x5b, 0xda, 0x33, 0x24, 0x57, 0xd1, 0xb7, 0xfe, 0xd6, 0x37,
	0x7a, 0x17, 0x22, 0x78, 0xff, 0x37, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xe1, 0x53, 0x67, 0x85, 0x05,
	0x00, 0x00,
}
