// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/third/holiday.proto

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

type HolidayInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Type                 uint32   `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Created              int64    `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64    `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string   `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string   `protobuf:"bytes,7,opt,name=operator,proto3" json:"operator,omitempty"`
	Name                 string   `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,9,opt,name=remark,proto3" json:"remark,omitempty"`
	Owner                string   `protobuf:"bytes,10,opt,name=owner,proto3" json:"owner,omitempty"`
	From                 int64    `protobuf:"varint,11,opt,name=from,proto3" json:"from,omitempty"`
	End                  int64    `protobuf:"varint,12,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HolidayInfo) Reset()         { *m = HolidayInfo{} }
func (m *HolidayInfo) String() string { return proto.CompactTextString(m) }
func (*HolidayInfo) ProtoMessage()    {}
func (*HolidayInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_933eba8bbce86885, []int{0}
}

func (m *HolidayInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HolidayInfo.Unmarshal(m, b)
}
func (m *HolidayInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HolidayInfo.Marshal(b, m, deterministic)
}
func (m *HolidayInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HolidayInfo.Merge(m, src)
}
func (m *HolidayInfo) XXX_Size() int {
	return xxx_messageInfo_HolidayInfo.Size(m)
}
func (m *HolidayInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_HolidayInfo.DiscardUnknown(m)
}

var xxx_messageInfo_HolidayInfo proto.InternalMessageInfo

func (m *HolidayInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *HolidayInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *HolidayInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *HolidayInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *HolidayInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *HolidayInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *HolidayInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *HolidayInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HolidayInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *HolidayInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *HolidayInfo) GetFrom() int64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *HolidayInfo) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

type ReqHolidayAdd struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Owner                string   `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	From                 int64    `protobuf:"varint,6,opt,name=from,proto3" json:"from,omitempty"`
	End                  int64    `protobuf:"varint,7,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqHolidayAdd) Reset()         { *m = ReqHolidayAdd{} }
func (m *ReqHolidayAdd) String() string { return proto.CompactTextString(m) }
func (*ReqHolidayAdd) ProtoMessage()    {}
func (*ReqHolidayAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_933eba8bbce86885, []int{1}
}

func (m *ReqHolidayAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqHolidayAdd.Unmarshal(m, b)
}
func (m *ReqHolidayAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqHolidayAdd.Marshal(b, m, deterministic)
}
func (m *ReqHolidayAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqHolidayAdd.Merge(m, src)
}
func (m *ReqHolidayAdd) XXX_Size() int {
	return xxx_messageInfo_ReqHolidayAdd.Size(m)
}
func (m *ReqHolidayAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqHolidayAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqHolidayAdd proto.InternalMessageInfo

func (m *ReqHolidayAdd) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqHolidayAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqHolidayAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqHolidayAdd) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqHolidayAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqHolidayAdd) GetFrom() int64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *ReqHolidayAdd) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

type ReqHolidayUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	From                 int64    `protobuf:"varint,6,opt,name=from,proto3" json:"from,omitempty"`
	End                  int64    `protobuf:"varint,7,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqHolidayUpdate) Reset()         { *m = ReqHolidayUpdate{} }
func (m *ReqHolidayUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqHolidayUpdate) ProtoMessage()    {}
func (*ReqHolidayUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_933eba8bbce86885, []int{2}
}

func (m *ReqHolidayUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqHolidayUpdate.Unmarshal(m, b)
}
func (m *ReqHolidayUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqHolidayUpdate.Marshal(b, m, deterministic)
}
func (m *ReqHolidayUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqHolidayUpdate.Merge(m, src)
}
func (m *ReqHolidayUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqHolidayUpdate.Size(m)
}
func (m *ReqHolidayUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqHolidayUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqHolidayUpdate proto.InternalMessageInfo

func (m *ReqHolidayUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqHolidayUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqHolidayUpdate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqHolidayUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqHolidayUpdate) GetFrom() int64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *ReqHolidayUpdate) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

type ReplyHolidayInfo struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Info                 *HolidayInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyHolidayInfo) Reset()         { *m = ReplyHolidayInfo{} }
func (m *ReplyHolidayInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyHolidayInfo) ProtoMessage()    {}
func (*ReplyHolidayInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_933eba8bbce86885, []int{3}
}

func (m *ReplyHolidayInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyHolidayInfo.Unmarshal(m, b)
}
func (m *ReplyHolidayInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyHolidayInfo.Marshal(b, m, deterministic)
}
func (m *ReplyHolidayInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyHolidayInfo.Merge(m, src)
}
func (m *ReplyHolidayInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyHolidayInfo.Size(m)
}
func (m *ReplyHolidayInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyHolidayInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyHolidayInfo proto.InternalMessageInfo

func (m *ReplyHolidayInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyHolidayInfo) GetInfo() *HolidayInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type ReplyHolidayList struct {
	Status               *ReplyStatus   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                uint32         `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Pages                uint32         `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	Page                 uint32         `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	Number               uint32         `protobuf:"varint,5,opt,name=number,proto3" json:"number,omitempty"`
	List                 []*HolidayInfo `protobuf:"bytes,16,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReplyHolidayList) Reset()         { *m = ReplyHolidayList{} }
func (m *ReplyHolidayList) String() string { return proto.CompactTextString(m) }
func (*ReplyHolidayList) ProtoMessage()    {}
func (*ReplyHolidayList) Descriptor() ([]byte, []int) {
	return fileDescriptor_933eba8bbce86885, []int{4}
}

func (m *ReplyHolidayList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyHolidayList.Unmarshal(m, b)
}
func (m *ReplyHolidayList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyHolidayList.Marshal(b, m, deterministic)
}
func (m *ReplyHolidayList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyHolidayList.Merge(m, src)
}
func (m *ReplyHolidayList) XXX_Size() int {
	return xxx_messageInfo_ReplyHolidayList.Size(m)
}
func (m *ReplyHolidayList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyHolidayList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyHolidayList proto.InternalMessageInfo

func (m *ReplyHolidayList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyHolidayList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyHolidayList) GetPages() uint32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *ReplyHolidayList) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ReplyHolidayList) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *ReplyHolidayList) GetList() []*HolidayInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*HolidayInfo)(nil), "omo.msp.third.HolidayInfo")
	proto.RegisterType((*ReqHolidayAdd)(nil), "omo.msp.third.ReqHolidayAdd")
	proto.RegisterType((*ReqHolidayUpdate)(nil), "omo.msp.third.ReqHolidayUpdate")
	proto.RegisterType((*ReplyHolidayInfo)(nil), "omo.msp.third.ReplyHolidayInfo")
	proto.RegisterType((*ReplyHolidayList)(nil), "omo.msp.third.ReplyHolidayList")
}

func init() {
	proto.RegisterFile("proto/third/holiday.proto", fileDescriptor_933eba8bbce86885)
}

var fileDescriptor_933eba8bbce86885 = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x5d, 0x6a, 0xdb, 0x40,
	0x10, 0x8e, 0x6c, 0x59, 0x49, 0xc6, 0x91, 0x31, 0x4b, 0x28, 0x5b, 0xd3, 0x12, 0xa3, 0x27, 0x3f,
	0x29, 0xe0, 0x9e, 0x20, 0x2e, 0xd4, 0x4d, 0xe9, 0x1f, 0x1b, 0xfa, 0xd2, 0x37, 0xc5, 0x3b, 0x6e,
	0x96, 0x5a, 0x5a, 0x65, 0xb5, 0x76, 0xf1, 0x35, 0x0a, 0x3d, 0x42, 0x7b, 0x9c, 0x9e, 0xa9, 0xec,
	0x48, 0x4e, 0x14, 0xac, 0xc4, 0xa1, 0x79, 0x9b, 0x6f, 0x66, 0xf6, 0x9b, 0xef, 0x1b, 0x0d, 0x82,
	0xe7, 0xb9, 0xd1, 0x56, 0x9f, 0xda, 0x2b, 0x65, 0xe4, 0xe9, 0x95, 0x5e, 0x28, 0x99, 0xac, 0x63,
	0xca, 0xb1, 0x50, 0xa7, 0x3a, 0x4e, 0x8b, 0x3c, 0xa6, 0xe2, 0x80, 0xd7, 0x3b, 0x67, 0x3a, 0x4d,
	0x75, 0x56, 0x36, 0x46, 0xbf, 0x5a, 0xd0, 0x7d, 0x5b, 0x3e, 0x3d, 0xcf, 0xe6, 0x9a, 0xf5, 0xa1,
	0xbd, 0x54, 0x92, 0x7b, 0x43, 0x6f, 0x74, 0x28, 0x5c, 0xc8, 0x7a, 0xd0, 0x52, 0x92, 0xb7, 0x86,
	0xde, 0xc8, 0x17, 0x2d, 0x25, 0x19, 0x03, 0xdf, 0xae, 0x73, 0xe4, 0xed, 0xa1, 0x37, 0x0a, 0x05,
	0xc5, 0x8c, 0xc3, 0xfe, 0xcc, 0x60, 0x62, 0x51, 0x72, 0x7f, 0xe8, 0x8d, 0xda, 0x62, 0x03, 0x5d,
	0x65, 0x99, 0x4b, 0xaa, 0x74, 0xca, 0x4a, 0x05, 0x6f, 0xde, 0x68, 0xc3, 0x03, 0x9a, 0xb6, 0x81,
	0x6c, 0x00, 0x07, 0x3a, 0x47, 0x43, 0xa5, 0x7d, 0x2a, 0xdd, 0x60, 0x37, 0x3d, 0x4b, 0x52, 0xe4,
	0x07, 0x94, 0xa7, 0x98, 0x3d, 0x83, 0xc0, 0x60, 0x9a, 0x98, 0xef, 0xfc, 0x90, 0xb2, 0x15, 0x62,
	0xc7, 0xd0, 0xd1, 0x3f, 0x32, 0x34, 0x1c, 0x28, 0x5d, 0x02, 0xc7, 0x30, 0x37, 0x3a, 0xe5, 0x5d,
	0x92, 0x43, 0xb1, 0x73, 0x8d, 0x99, 0xe4, 0x47, 0x94, 0x72, 0x61, 0xf4, 0xdb, 0x83, 0x50, 0xe0,
	0x75, 0xb5, 0x9a, 0x33, 0x29, 0x1b, 0x36, 0xb3, 0xd1, 0xd2, 0x6a, 0xd4, 0xd2, 0x6e, 0xd6, 0xe2,
	0xd7, 0xb5, 0xd4, 0x9d, 0x76, 0xb6, 0x9d, 0x92, 0xce, 0x60, 0x5b, 0xe7, 0xfe, 0xad, 0xce, 0x9f,
	0x1e, 0xf4, 0x6f, 0x75, 0x7e, 0xa1, 0xdd, 0x3e, 0x51, 0xea, 0xd3, 0x45, 0xad, 0x9c, 0xa6, 0x7c,
	0xb1, 0xae, 0x1f, 0xd6, 0x18, 0x82, 0xc2, 0x26, 0x76, 0x59, 0x90, 0xac, 0xee, 0x78, 0x10, 0xdf,
	0x39, 0xd1, 0x98, 0x1e, 0x5c, 0x50, 0x87, 0xa8, 0x3a, 0x59, 0x0c, 0xbe, 0xca, 0xe6, 0x9a, 0x54,
	0x6f, 0xbf, 0xa8, 0xb1, 0x0b, 0xea, 0x8b, 0xfe, 0x7a, 0x77, 0x07, 0xbf, 0x57, 0x85, 0xfd, 0xaf,
	0xc1, 0xc7, 0xd0, 0xb1, 0xda, 0x26, 0x0b, 0x9a, 0x1c, 0x8a, 0x12, 0xb8, 0x6c, 0x9e, 0x7c, 0xc3,
	0xa2, 0x3a, 0xfd, 0x12, 0xb8, 0x95, 0xb8, 0x80, 0x3e, 0x6c, 0x28, 0x28, 0x76, 0xab, 0xcd, 0x96,
	0xe9, 0x25, 0x96, 0x0b, 0x0c, 0x45, 0x85, 0x9c, 0xa1, 0x85, 0x2a, 0x2c, 0xef, 0x0f, 0xdb, 0xbb,
	0x0c, 0xb9, 0xbe, 0xf1, 0x1f, 0x1f, 0x7a, 0x55, 0xf6, 0x02, 0xcd, 0x4a, 0xcd, 0x90, 0x9d, 0x43,
	0x70, 0x26, 0xe5, 0xa7, 0x0c, 0xd9, 0x8b, 0x2d, 0x23, 0xb5, 0x73, 0x1d, 0x9c, 0x34, 0xd9, 0xac,
	0x4d, 0x88, 0xf6, 0xd8, 0x14, 0x82, 0x29, 0x5a, 0x47, 0xb5, 0xbd, 0x93, 0xeb, 0x25, 0x16, 0xd6,
	0xf5, 0x3d, 0x86, 0xe8, 0x33, 0x40, 0x79, 0x79, 0x93, 0xa4, 0x40, 0x76, 0x72, 0xaf, 0xae, 0xb2,
	0xe9, 0x31, 0x8c, 0x1f, 0xa1, 0x3b, 0x45, 0x3b, 0x59, 0xbf, 0x51, 0x0b, 0x8b, 0xa6, 0xc9, 0xaa,
	0xd3, 0x57, 0x56, 0x1f, 0xe4, 0x73, 0x27, 0x10, 0xed, 0xb1, 0x0f, 0x70, 0x34, 0x45, 0xeb, 0xbe,
	0xb2, 0x2a, 0xac, 0x9a, 0xed, 0x20, 0x7c, 0x79, 0xdf, 0x89, 0xd0, 0xe3, 0x68, 0x8f, 0xbd, 0x83,
	0x5e, 0x65, 0x78, 0x87, 0xc2, 0xca, 0x31, 0x6f, 0x22, 0xac, 0xac, 0xbe, 0x86, 0x43, 0x81, 0xa9,
	0x5e, 0xe1, 0xae, 0x0f, 0xf1, 0x00, 0xc9, 0x24, 0xfc, 0xda, 0xad, 0xfd, 0xe2, 0x2f, 0x03, 0x02,
	0xaf, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xab, 0xd1, 0x3c, 0x0c, 0x22, 0x06, 0x00, 0x00,
}
