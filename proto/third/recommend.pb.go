// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/third/recommend.proto

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

type RecommendInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64    `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64    `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string   `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string   `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Name                 string   `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Owner                string   `protobuf:"bytes,8,opt,name=owner,proto3" json:"owner,omitempty"`
	Type                 uint32   `protobuf:"varint,9,opt,name=type,proto3" json:"type,omitempty"`
	Targets              []string `protobuf:"bytes,31,rep,name=targets,proto3" json:"targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecommendInfo) Reset()         { *m = RecommendInfo{} }
func (m *RecommendInfo) String() string { return proto.CompactTextString(m) }
func (*RecommendInfo) ProtoMessage()    {}
func (*RecommendInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cef07d596f5e8559, []int{0}
}

func (m *RecommendInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecommendInfo.Unmarshal(m, b)
}
func (m *RecommendInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecommendInfo.Marshal(b, m, deterministic)
}
func (m *RecommendInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecommendInfo.Merge(m, src)
}
func (m *RecommendInfo) XXX_Size() int {
	return xxx_messageInfo_RecommendInfo.Size(m)
}
func (m *RecommendInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RecommendInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RecommendInfo proto.InternalMessageInfo

func (m *RecommendInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RecommendInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RecommendInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *RecommendInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *RecommendInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *RecommendInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *RecommendInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RecommendInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *RecommendInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *RecommendInfo) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

type ReqRecommendAdd struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Owner                string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	Type                 uint32   `protobuf:"varint,5,opt,name=type,proto3" json:"type,omitempty"`
	Targets              []string `protobuf:"bytes,21,rep,name=targets,proto3" json:"targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqRecommendAdd) Reset()         { *m = ReqRecommendAdd{} }
func (m *ReqRecommendAdd) String() string { return proto.CompactTextString(m) }
func (*ReqRecommendAdd) ProtoMessage()    {}
func (*ReqRecommendAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_cef07d596f5e8559, []int{1}
}

func (m *ReqRecommendAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqRecommendAdd.Unmarshal(m, b)
}
func (m *ReqRecommendAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqRecommendAdd.Marshal(b, m, deterministic)
}
func (m *ReqRecommendAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqRecommendAdd.Merge(m, src)
}
func (m *ReqRecommendAdd) XXX_Size() int {
	return xxx_messageInfo_ReqRecommendAdd.Size(m)
}
func (m *ReqRecommendAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqRecommendAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqRecommendAdd proto.InternalMessageInfo

func (m *ReqRecommendAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqRecommendAdd) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqRecommendAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqRecommendAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqRecommendAdd) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqRecommendAdd) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

type ReqRecommendUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	Targets              []string `protobuf:"bytes,21,rep,name=targets,proto3" json:"targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqRecommendUpdate) Reset()         { *m = ReqRecommendUpdate{} }
func (m *ReqRecommendUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqRecommendUpdate) ProtoMessage()    {}
func (*ReqRecommendUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_cef07d596f5e8559, []int{2}
}

func (m *ReqRecommendUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqRecommendUpdate.Unmarshal(m, b)
}
func (m *ReqRecommendUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqRecommendUpdate.Marshal(b, m, deterministic)
}
func (m *ReqRecommendUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqRecommendUpdate.Merge(m, src)
}
func (m *ReqRecommendUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqRecommendUpdate.Size(m)
}
func (m *ReqRecommendUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqRecommendUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqRecommendUpdate proto.InternalMessageInfo

func (m *ReqRecommendUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqRecommendUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqRecommendUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqRecommendUpdate) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

type ReplyRecommendInfo struct {
	Info                 *RecommendInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReplyRecommendInfo) Reset()         { *m = ReplyRecommendInfo{} }
func (m *ReplyRecommendInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyRecommendInfo) ProtoMessage()    {}
func (*ReplyRecommendInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cef07d596f5e8559, []int{3}
}

func (m *ReplyRecommendInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyRecommendInfo.Unmarshal(m, b)
}
func (m *ReplyRecommendInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyRecommendInfo.Marshal(b, m, deterministic)
}
func (m *ReplyRecommendInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyRecommendInfo.Merge(m, src)
}
func (m *ReplyRecommendInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyRecommendInfo.Size(m)
}
func (m *ReplyRecommendInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyRecommendInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyRecommendInfo proto.InternalMessageInfo

func (m *ReplyRecommendInfo) GetInfo() *RecommendInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplyRecommendInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplyRecommendList struct {
	Status               *ReplyStatus     `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                uint64           `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	PageMax              uint32           `protobuf:"varint,3,opt,name=pageMax,proto3" json:"pageMax,omitempty"`
	PageNow              uint32           `protobuf:"varint,4,opt,name=pageNow,proto3" json:"pageNow,omitempty"`
	List                 []*RecommendInfo `protobuf:"bytes,21,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ReplyRecommendList) Reset()         { *m = ReplyRecommendList{} }
func (m *ReplyRecommendList) String() string { return proto.CompactTextString(m) }
func (*ReplyRecommendList) ProtoMessage()    {}
func (*ReplyRecommendList) Descriptor() ([]byte, []int) {
	return fileDescriptor_cef07d596f5e8559, []int{4}
}

func (m *ReplyRecommendList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyRecommendList.Unmarshal(m, b)
}
func (m *ReplyRecommendList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyRecommendList.Marshal(b, m, deterministic)
}
func (m *ReplyRecommendList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyRecommendList.Merge(m, src)
}
func (m *ReplyRecommendList) XXX_Size() int {
	return xxx_messageInfo_ReplyRecommendList.Size(m)
}
func (m *ReplyRecommendList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyRecommendList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyRecommendList proto.InternalMessageInfo

func (m *ReplyRecommendList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyRecommendList) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyRecommendList) GetPageMax() uint32 {
	if m != nil {
		return m.PageMax
	}
	return 0
}

func (m *ReplyRecommendList) GetPageNow() uint32 {
	if m != nil {
		return m.PageNow
	}
	return 0
}

func (m *ReplyRecommendList) GetList() []*RecommendInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*RecommendInfo)(nil), "omo.msp.third.RecommendInfo")
	proto.RegisterType((*ReqRecommendAdd)(nil), "omo.msp.third.ReqRecommendAdd")
	proto.RegisterType((*ReqRecommendUpdate)(nil), "omo.msp.third.ReqRecommendUpdate")
	proto.RegisterType((*ReplyRecommendInfo)(nil), "omo.msp.third.ReplyRecommendInfo")
	proto.RegisterType((*ReplyRecommendList)(nil), "omo.msp.third.ReplyRecommendList")
}

func init() {
	proto.RegisterFile("proto/third/recommend.proto", fileDescriptor_cef07d596f5e8559)
}

var fileDescriptor_cef07d596f5e8559 = []byte{
	// 558 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0x13, 0xc7, 0x69, 0x26, 0xa4, 0x54, 0x2b, 0x40, 0x56, 0x28, 0x10, 0x72, 0xca, 0x29,
	0x45, 0xe1, 0x09, 0x5a, 0x24, 0xa2, 0xa2, 0xa6, 0x54, 0x5b, 0x71, 0xe1, 0x66, 0xb2, 0xd3, 0xb0,
	0x22, 0xf6, 0xba, 0xeb, 0x4d, 0x43, 0x78, 0x17, 0xde, 0x86, 0x37, 0xe2, 0xc2, 0x11, 0xed, 0xd8,
	0x8e, 0xec, 0xc6, 0x21, 0x39, 0x70, 0xf3, 0xec, 0x7c, 0xfb, 0xfd, 0xec, 0x8c, 0x0c, 0xcf, 0x63,
	0xad, 0x8c, 0x3a, 0x35, 0x5f, 0xa5, 0x16, 0xa7, 0x1a, 0xa7, 0x2a, 0x0c, 0x31, 0x12, 0x43, 0x3a,
	0x65, 0x1d, 0x15, 0xaa, 0x61, 0x98, 0xc4, 0x43, 0x6a, 0x77, 0xfd, 0x22, 0xd6, 0x22, 0x55, 0x94,
	0x02, 0xfb, 0xbf, 0x1d, 0xe8, 0xf0, 0xfc, 0xf2, 0x45, 0x74, 0xab, 0xd8, 0x31, 0xd4, 0x17, 0x52,
	0xf8, 0x4e, 0xcf, 0x19, 0xb4, 0xb8, 0xfd, 0x64, 0x47, 0x50, 0x93, 0xc2, 0xaf, 0xf5, 0x9c, 0x81,
	0xcb, 0x6b, 0x52, 0x30, 0x1f, 0x9a, 0x53, 0x8d, 0x81, 0x41, 0xe1, 0xd7, 0x7b, 0xce, 0xa0, 0xce,
	0xf3, 0xd2, 0x76, 0x16, 0xb1, 0xa0, 0x8e, 0x9b, 0x76, 0xb2, 0x72, 0x7d, 0x47, 0x69, 0xbf, 0x41,
	0xcc, 0x79, 0xc9, 0xba, 0x70, 0xa8, 0x62, 0xd4, 0xd4, 0xf2, 0xa8, 0xb5, 0xae, 0x19, 0x03, 0x37,
	0x0a, 0x42, 0xf4, 0x9b, 0x74, 0x4e, 0xdf, 0xec, 0x09, 0x34, 0xd4, 0x32, 0x42, 0xed, 0x1f, 0xd2,
	0x61, 0x5a, 0x58, 0xa4, 0x59, 0xc5, 0xe8, 0xb7, 0x7a, 0xce, 0xa0, 0xc3, 0xe9, 0xdb, 0x6a, 0x9a,
	0x40, 0xcf, 0xd0, 0x24, 0xfe, 0xab, 0x5e, 0xdd, 0x6a, 0x66, 0x65, 0xff, 0xa7, 0x03, 0x8f, 0x39,
	0xde, 0xad, 0x83, 0x9f, 0x09, 0xb1, 0xd6, 0x72, 0xaa, 0xb4, 0x6a, 0x45, 0xad, 0x67, 0xe0, 0x69,
	0x0c, 0x03, 0xfd, 0x8d, 0xe2, 0xb7, 0x78, 0x56, 0x95, 0x92, 0xb8, 0x9b, 0x49, 0xc8, 0x5f, 0xa3,
	0xda, 0xdf, 0xd3, 0xb2, 0xbf, 0x18, 0x58, 0xd1, 0xde, 0x27, 0x7a, 0xc4, 0x8a, 0xc9, 0xe4, 0x9e,
	0x6b, 0x05, 0xcf, 0x45, 0x17, 0xf5, 0x07, 0x2e, 0xb6, 0x2b, 0xfe, 0xb0, 0x8a, 0xf1, 0x7c, 0x55,
	0xde, 0x85, 0x37, 0xe0, 0xca, 0xe8, 0x56, 0x91, 0x64, 0x7b, 0x74, 0x32, 0x2c, 0x6d, 0xd5, 0xb0,
	0x84, 0xe5, 0x84, 0x64, 0x23, 0xf0, 0x12, 0x13, 0x98, 0x45, 0x42, 0x9e, 0xda, 0xa3, 0xee, 0xc6,
	0x9d, 0x78, 0xbe, 0xba, 0x21, 0x04, 0xcf, 0x90, 0xfd, 0x5f, 0xce, 0x43, 0xf1, 0x4b, 0x99, 0x98,
	0x02, 0x95, 0xb3, 0x2f, 0x95, 0x1d, 0x98, 0x51, 0x26, 0x98, 0x67, 0xdb, 0x9a, 0x16, 0x36, 0x76,
	0x1c, 0xcc, 0x70, 0x12, 0x7c, 0xa7, 0x17, 0xe9, 0xf0, 0xbc, 0xcc, 0x3b, 0x57, 0x6a, 0x49, 0x13,
	0xcb, 0x3a, 0x57, 0x6a, 0x69, 0xa3, 0xcf, 0x65, 0x62, 0xe8, 0x9d, 0x76, 0x46, 0xb7, 0xc8, 0xd1,
	0x1f, 0x17, 0x8e, 0xd7, 0xe7, 0x37, 0xa8, 0xef, 0xe5, 0x14, 0xd9, 0x04, 0xbc, 0x33, 0x21, 0x3e,
	0x46, 0xc8, 0x5e, 0x6e, 0x50, 0x94, 0xf6, 0xaf, 0xfb, 0xba, 0x2a, 0x5e, 0x49, 0xa7, 0x7f, 0xc0,
	0x2e, 0xa1, 0x95, 0x2e, 0x83, 0x65, 0xdc, 0x7c, 0x90, 0xbb, 0x05, 0x26, 0xc6, 0x42, 0xf7, 0x63,
	0xbb, 0x00, 0x6f, 0x8c, 0xe6, 0xbf, 0x50, 0x7d, 0x80, 0xe6, 0x18, 0x0d, 0xcd, 0x6d, 0x0b, 0xd7,
	0x75, 0x30, 0xc3, 0x1d, 0x5c, 0xf6, 0x7a, 0xff, 0x80, 0xbd, 0x83, 0x16, 0xc7, 0x50, 0xdd, 0xef,
	0x0c, 0xe9, 0x57, 0xb1, 0x65, 0x86, 0xae, 0xa1, 0x3d, 0x46, 0x73, 0xbe, 0x7a, 0x2f, 0xe7, 0x06,
	0x35, 0x3b, 0xa9, 0xa6, 0x49, 0xbb, 0xfb, 0xd9, 0x9a, 0xc0, 0xa3, 0x31, 0x1a, 0xbb, 0x70, 0x32,
	0x31, 0x72, 0xba, 0x83, 0xf2, 0xc5, 0xb6, 0x6d, 0xa5, 0xcb, 0xf4, 0x62, 0x47, 0xe9, 0x28, 0x77,
	0x79, 0x4c, 0x51, 0xff, 0x0a, 0x7b, 0xde, 0xf9, 0xdc, 0x2e, 0xfc, 0xe1, 0xbf, 0x78, 0x54, 0xbc,
	0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xaf, 0xce, 0x74, 0xe9, 0x23, 0x06, 0x00, 0x00,
}
