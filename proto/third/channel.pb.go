// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/third/channel.proto

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

type ChannelInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64    `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64    `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string   `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string   `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Name                 string   `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,8,opt,name=remark,proto3" json:"remark,omitempty"`
	Phone                string   `protobuf:"bytes,9,opt,name=phone,proto3" json:"phone,omitempty"`
	Secret               string   `protobuf:"bytes,10,opt,name=secret,proto3" json:"secret,omitempty"`
	Tags                 []string `protobuf:"bytes,11,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChannelInfo) Reset()         { *m = ChannelInfo{} }
func (m *ChannelInfo) String() string { return proto.CompactTextString(m) }
func (*ChannelInfo) ProtoMessage()    {}
func (*ChannelInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_76e36d889a4cd775, []int{0}
}

func (m *ChannelInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelInfo.Unmarshal(m, b)
}
func (m *ChannelInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelInfo.Marshal(b, m, deterministic)
}
func (m *ChannelInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelInfo.Merge(m, src)
}
func (m *ChannelInfo) XXX_Size() int {
	return xxx_messageInfo_ChannelInfo.Size(m)
}
func (m *ChannelInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelInfo proto.InternalMessageInfo

func (m *ChannelInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ChannelInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ChannelInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *ChannelInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *ChannelInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ChannelInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ChannelInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChannelInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ChannelInfo) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *ChannelInfo) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *ChannelInfo) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type ReqChannelAdd struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Phone                string   `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	Remark               string   `protobuf:"bytes,6,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,8,opt,name=operator,proto3" json:"operator,omitempty"`
	Passwords            string   `protobuf:"bytes,9,opt,name=passwords,proto3" json:"passwords,omitempty"`
	Entity               string   `protobuf:"bytes,10,opt,name=entity,proto3" json:"entity,omitempty"`
	Portrait             string   `protobuf:"bytes,11,opt,name=portrait,proto3" json:"portrait,omitempty"`
	Tags                 []string `protobuf:"bytes,13,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqChannelAdd) Reset()         { *m = ReqChannelAdd{} }
func (m *ReqChannelAdd) String() string { return proto.CompactTextString(m) }
func (*ReqChannelAdd) ProtoMessage()    {}
func (*ReqChannelAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_76e36d889a4cd775, []int{1}
}

func (m *ReqChannelAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqChannelAdd.Unmarshal(m, b)
}
func (m *ReqChannelAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqChannelAdd.Marshal(b, m, deterministic)
}
func (m *ReqChannelAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqChannelAdd.Merge(m, src)
}
func (m *ReqChannelAdd) XXX_Size() int {
	return xxx_messageInfo_ReqChannelAdd.Size(m)
}
func (m *ReqChannelAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqChannelAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqChannelAdd proto.InternalMessageInfo

func (m *ReqChannelAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqChannelAdd) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *ReqChannelAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqChannelAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqChannelAdd) GetPasswords() string {
	if m != nil {
		return m.Passwords
	}
	return ""
}

func (m *ReqChannelAdd) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *ReqChannelAdd) GetPortrait() string {
	if m != nil {
		return m.Portrait
	}
	return ""
}

func (m *ReqChannelAdd) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type ReqChannelUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	NickName             string   `protobuf:"bytes,3,opt,name=nickName,proto3" json:"nickName,omitempty"`
	Phone                string   `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Job                  string   `protobuf:"bytes,5,opt,name=job,proto3" json:"job,omitempty"`
	Remark               string   `protobuf:"bytes,6,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,8,opt,name=operator,proto3" json:"operator,omitempty"`
	Portrait             string   `protobuf:"bytes,9,opt,name=portrait,proto3" json:"portrait,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqChannelUpdate) Reset()         { *m = ReqChannelUpdate{} }
func (m *ReqChannelUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqChannelUpdate) ProtoMessage()    {}
func (*ReqChannelUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_76e36d889a4cd775, []int{2}
}

func (m *ReqChannelUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqChannelUpdate.Unmarshal(m, b)
}
func (m *ReqChannelUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqChannelUpdate.Marshal(b, m, deterministic)
}
func (m *ReqChannelUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqChannelUpdate.Merge(m, src)
}
func (m *ReqChannelUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqChannelUpdate.Size(m)
}
func (m *ReqChannelUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqChannelUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqChannelUpdate proto.InternalMessageInfo

func (m *ReqChannelUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqChannelUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqChannelUpdate) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *ReqChannelUpdate) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *ReqChannelUpdate) GetJob() string {
	if m != nil {
		return m.Job
	}
	return ""
}

func (m *ReqChannelUpdate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqChannelUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqChannelUpdate) GetPortrait() string {
	if m != nil {
		return m.Portrait
	}
	return ""
}

type ReqChannelTags struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Operator             string   `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	Tags                 []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqChannelTags) Reset()         { *m = ReqChannelTags{} }
func (m *ReqChannelTags) String() string { return proto.CompactTextString(m) }
func (*ReqChannelTags) ProtoMessage()    {}
func (*ReqChannelTags) Descriptor() ([]byte, []int) {
	return fileDescriptor_76e36d889a4cd775, []int{3}
}

func (m *ReqChannelTags) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqChannelTags.Unmarshal(m, b)
}
func (m *ReqChannelTags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqChannelTags.Marshal(b, m, deterministic)
}
func (m *ReqChannelTags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqChannelTags.Merge(m, src)
}
func (m *ReqChannelTags) XXX_Size() int {
	return xxx_messageInfo_ReqChannelTags.Size(m)
}
func (m *ReqChannelTags) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqChannelTags.DiscardUnknown(m)
}

var xxx_messageInfo_ReqChannelTags proto.InternalMessageInfo

func (m *ReqChannelTags) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqChannelTags) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqChannelTags) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type ReplyChannelInfo struct {
	Info                 *ChannelInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyChannelInfo) Reset()         { *m = ReplyChannelInfo{} }
func (m *ReplyChannelInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyChannelInfo) ProtoMessage()    {}
func (*ReplyChannelInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_76e36d889a4cd775, []int{4}
}

func (m *ReplyChannelInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyChannelInfo.Unmarshal(m, b)
}
func (m *ReplyChannelInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyChannelInfo.Marshal(b, m, deterministic)
}
func (m *ReplyChannelInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyChannelInfo.Merge(m, src)
}
func (m *ReplyChannelInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyChannelInfo.Size(m)
}
func (m *ReplyChannelInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyChannelInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyChannelInfo proto.InternalMessageInfo

func (m *ReplyChannelInfo) GetInfo() *ChannelInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplyChannelInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReqChannelList struct {
	List                 []string `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqChannelList) Reset()         { *m = ReqChannelList{} }
func (m *ReqChannelList) String() string { return proto.CompactTextString(m) }
func (*ReqChannelList) ProtoMessage()    {}
func (*ReqChannelList) Descriptor() ([]byte, []int) {
	return fileDescriptor_76e36d889a4cd775, []int{5}
}

func (m *ReqChannelList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqChannelList.Unmarshal(m, b)
}
func (m *ReqChannelList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqChannelList.Marshal(b, m, deterministic)
}
func (m *ReqChannelList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqChannelList.Merge(m, src)
}
func (m *ReqChannelList) XXX_Size() int {
	return xxx_messageInfo_ReqChannelList.Size(m)
}
func (m *ReqChannelList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqChannelList.DiscardUnknown(m)
}

var xxx_messageInfo_ReqChannelList proto.InternalMessageInfo

func (m *ReqChannelList) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

type ReplyChannelList struct {
	Total                uint64         `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	PageMax              uint32         `protobuf:"varint,2,opt,name=pageMax,proto3" json:"pageMax,omitempty"`
	PageNow              uint32         `protobuf:"varint,3,opt,name=pageNow,proto3" json:"pageNow,omitempty"`
	List                 []*ChannelInfo `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`
	Status               *ReplyStatus   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReplyChannelList) Reset()         { *m = ReplyChannelList{} }
func (m *ReplyChannelList) String() string { return proto.CompactTextString(m) }
func (*ReplyChannelList) ProtoMessage()    {}
func (*ReplyChannelList) Descriptor() ([]byte, []int) {
	return fileDescriptor_76e36d889a4cd775, []int{6}
}

func (m *ReplyChannelList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyChannelList.Unmarshal(m, b)
}
func (m *ReplyChannelList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyChannelList.Marshal(b, m, deterministic)
}
func (m *ReplyChannelList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyChannelList.Merge(m, src)
}
func (m *ReplyChannelList) XXX_Size() int {
	return xxx_messageInfo_ReplyChannelList.Size(m)
}
func (m *ReplyChannelList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyChannelList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyChannelList proto.InternalMessageInfo

func (m *ReplyChannelList) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyChannelList) GetPageMax() uint32 {
	if m != nil {
		return m.PageMax
	}
	return 0
}

func (m *ReplyChannelList) GetPageNow() uint32 {
	if m != nil {
		return m.PageNow
	}
	return 0
}

func (m *ReplyChannelList) GetList() []*ChannelInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ReplyChannelList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*ChannelInfo)(nil), "omo.msp.third.ChannelInfo")
	proto.RegisterType((*ReqChannelAdd)(nil), "omo.msp.third.ReqChannelAdd")
	proto.RegisterType((*ReqChannelUpdate)(nil), "omo.msp.third.ReqChannelUpdate")
	proto.RegisterType((*ReqChannelTags)(nil), "omo.msp.third.ReqChannelTags")
	proto.RegisterType((*ReplyChannelInfo)(nil), "omo.msp.third.ReplyChannelInfo")
	proto.RegisterType((*ReqChannelList)(nil), "omo.msp.third.ReqChannelList")
	proto.RegisterType((*ReplyChannelList)(nil), "omo.msp.third.ReplyChannelList")
}

func init() {
	proto.RegisterFile("proto/third/channel.proto", fileDescriptor_76e36d889a4cd775)
}

var fileDescriptor_76e36d889a4cd775 = []byte{
	// 606 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdf, 0x6e, 0xd3, 0x3e,
	0x14, 0x5e, 0xfe, 0x34, 0x6b, 0x4e, 0xd5, 0x69, 0xb2, 0xa6, 0x9f, 0xfc, 0xab, 0x26, 0x51, 0x45,
	0x5c, 0xec, 0x2a, 0x93, 0xc6, 0x13, 0x8c, 0x5d, 0x8c, 0x21, 0x18, 0x28, 0x83, 0x1b, 0xee, 0xbc,
	0xc6, 0xeb, 0xc2, 0x9a, 0x38, 0xb3, 0xdd, 0x95, 0x3d, 0x01, 0xaf, 0xc5, 0x03, 0x70, 0xc7, 0x03,
	0x81, 0x7c, 0x9c, 0xa4, 0xee, 0x5a, 0xd8, 0x84, 0xb8, 0xf3, 0xe7, 0xf3, 0xf9, 0x3b, 0xdf, 0x77,
	0x1c, 0x2b, 0xf0, 0x7f, 0x2d, 0x85, 0x16, 0x87, 0xfa, 0xba, 0x90, 0xf9, 0xe1, 0xe4, 0x9a, 0x55,
	0x15, 0x9f, 0xa5, 0xb8, 0x47, 0x86, 0xa2, 0x14, 0x69, 0xa9, 0xea, 0x14, 0x8b, 0x23, 0xba, 0xc2,
	0x14, 0x65, 0x29, 0x2a, 0x4b, 0x4c, 0xbe, 0xfa, 0x30, 0x38, 0xb1, 0x47, 0xcf, 0xaa, 0x2b, 0x41,
	0x76, 0x21, 0x98, 0x17, 0x39, 0xf5, 0xc6, 0xde, 0x41, 0x9c, 0x99, 0x25, 0xd9, 0x01, 0xbf, 0xc8,
	0xa9, 0x3f, 0xf6, 0x0e, 0xc2, 0xcc, 0x2f, 0x72, 0x42, 0x61, 0x7b, 0x22, 0x39, 0xd3, 0x3c, 0xa7,
	0xc1, 0xd8, 0x3b, 0x08, 0xb2, 0x16, 0x9a, 0xca, 0xbc, 0xce, 0xb1, 0x12, 0xda, 0x4a, 0x03, 0xbb,
	0x33, 0x42, 0xd2, 0x1e, 0x2a, 0xb7, 0x90, 0x8c, 0xa0, 0x2f, 0x6a, 0x2e, 0xb1, 0x14, 0x61, 0xa9,
	0xc3, 0x84, 0x40, 0x58, 0xb1, 0x92, 0xd3, 0x6d, 0xdc, 0xc7, 0x35, 0xf9, 0x0f, 0x22, 0xc9, 0x4b,
	0x26, 0x6f, 0x68, 0x1f, 0x77, 0x1b, 0x44, 0xf6, 0xa0, 0x57, 0x5f, 0x8b, 0x8a, 0xd3, 0x18, 0xb7,
	0x2d, 0x30, 0x6c, 0xc5, 0x27, 0x92, 0x6b, 0x0a, 0x96, 0x6d, 0x91, 0x51, 0xd6, 0x6c, 0xaa, 0xe8,
	0x60, 0x1c, 0x18, 0x65, 0xb3, 0x4e, 0x7e, 0x78, 0x30, 0xcc, 0xf8, 0x6d, 0x33, 0x8c, 0xe3, 0x3c,
	0xef, 0xfa, 0x7b, 0x4e, 0xff, 0xae, 0x4f, 0xef, 0x41, 0x9f, 0xc6, 0x55, 0xb4, 0xe2, 0xca, 0x4d,
	0xd7, 0x7f, 0x90, 0x6e, 0x1f, 0xe2, 0x9a, 0x29, 0xb5, 0x10, 0x32, 0x57, 0x8d, 0xeb, 0xe5, 0x86,
	0x51, 0xe4, 0x95, 0x2e, 0xf4, 0x7d, 0xeb, 0xdc, 0x22, 0xa3, 0x58, 0x0b, 0xa9, 0x25, 0x2b, 0x34,
	0x1d, 0x58, 0xc5, 0x16, 0x77, 0xa9, 0x86, 0x4e, 0xaa, 0xef, 0x1e, 0xec, 0x2e, 0x53, 0x7d, 0xc4,
	0xfb, 0xd8, 0x70, 0xc9, 0x6d, 0x54, 0xdf, 0x89, 0x3a, 0x82, 0x7e, 0x55, 0x4c, 0x6e, 0xce, 0xcd,
	0x7e, 0x60, 0x5b, 0xb5, 0x78, 0x39, 0x86, 0xd0, 0x1d, 0xc3, 0x2e, 0x04, 0x9f, 0xc5, 0x65, 0x33,
	0x1a, 0xb3, 0xfc, 0xab, 0xc1, 0xb8, 0x11, 0xe3, 0xd5, 0x88, 0x49, 0x06, 0x3b, 0xcb, 0x34, 0x1f,
	0xd8, 0x54, 0x6d, 0xc8, 0xe2, 0x6a, 0xfb, 0xeb, 0x9f, 0x14, 0x8e, 0x28, 0x70, 0x46, 0x74, 0x67,
	0x26, 0x54, 0xcf, 0xee, 0xdd, 0x67, 0x90, 0x42, 0x58, 0x54, 0x57, 0x02, 0x65, 0x07, 0x47, 0xa3,
	0x74, 0xe5, 0x39, 0xa5, 0x0e, 0x33, 0x43, 0x1e, 0x39, 0x82, 0x48, 0x69, 0xa6, 0xe7, 0x0a, 0x3b,
	0xae, 0x9f, 0xc0, 0x06, 0x17, 0xc8, 0xc8, 0x1a, 0x66, 0xf2, 0xdc, 0xcd, 0xf2, 0xa6, 0x50, 0x78,
	0x81, 0xb3, 0x42, 0x69, 0x1a, 0x5a, 0x77, 0x66, 0x9d, 0x7c, 0xf3, 0x56, 0xed, 0x21, 0x71, 0x0f,
	0x7a, 0x5a, 0x68, 0x36, 0x43, 0x7f, 0x61, 0x66, 0x81, 0x79, 0x65, 0x35, 0x9b, 0xf2, 0xb7, 0xec,
	0x0b, 0xba, 0x18, 0x66, 0x2d, 0x6c, 0x2b, 0xe7, 0x62, 0x81, 0x37, 0xd9, 0x54, 0xce, 0xc5, 0xc2,
	0x04, 0xed, 0x5a, 0x3e, 0x12, 0xd4, 0xf0, 0x9c, 0xa0, 0xbd, 0xa7, 0x06, 0x3d, 0xfa, 0xe9, 0xc3,
	0x4e, 0xa3, 0x74, 0xc1, 0xe5, 0x5d, 0x31, 0xe1, 0xe4, 0x0c, 0xa2, 0xe3, 0x3c, 0x7f, 0x57, 0x71,
	0xb2, 0xbf, 0x26, 0xe0, 0x3c, 0xc1, 0xd1, 0xb3, 0x4d, 0xf2, 0x8e, 0xab, 0x64, 0x8b, 0xbc, 0x86,
	0xd8, 0x7e, 0xd6, 0x46, 0x6d, 0xdd, 0xce, 0xed, 0x9c, 0x2b, 0x6d, 0xa8, 0x4f, 0xd1, 0x3a, 0x85,
	0xe8, 0x94, 0xeb, 0x7f, 0x20, 0xf4, 0x0a, 0xb6, 0x4f, 0xb9, 0xc6, 0xbb, 0xfa, 0x8d, 0xd2, 0x7b,
	0x36, 0xe5, 0x7f, 0x54, 0x32, 0x87, 0x93, 0x2d, 0x72, 0x02, 0x71, 0xc6, 0x4b, 0x71, 0xf7, 0x68,
	0x3c, 0xba, 0x49, 0xcb, 0xda, 0x79, 0x39, 0xfc, 0x34, 0x70, 0xfe, 0x00, 0x97, 0x11, 0x82, 0x17,
	0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x03, 0x4b, 0x82, 0x4f, 0x41, 0x06, 0x00, 0x00,
}
