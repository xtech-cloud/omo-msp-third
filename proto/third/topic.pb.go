// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/third/topic.proto

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

type TopicInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Type                 uint32   `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Created              int64    `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64    `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string   `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string   `protobuf:"bytes,7,opt,name=operator,proto3" json:"operator,omitempty"`
	Name                 string   `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,9,opt,name=remark,proto3" json:"remark,omitempty"`
	Scene                string   `protobuf:"bytes,10,opt,name=scene,proto3" json:"scene,omitempty"`
	Time                 uint32   `protobuf:"varint,11,opt,name=time,proto3" json:"time,omitempty"`
	Compere              string   `protobuf:"bytes,12,opt,name=compere,proto3" json:"compere,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicInfo) Reset()         { *m = TopicInfo{} }
func (m *TopicInfo) String() string { return proto.CompactTextString(m) }
func (*TopicInfo) ProtoMessage()    {}
func (*TopicInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac6e87e37c521a00, []int{0}
}

func (m *TopicInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicInfo.Unmarshal(m, b)
}
func (m *TopicInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicInfo.Marshal(b, m, deterministic)
}
func (m *TopicInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicInfo.Merge(m, src)
}
func (m *TopicInfo) XXX_Size() int {
	return xxx_messageInfo_TopicInfo.Size(m)
}
func (m *TopicInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TopicInfo proto.InternalMessageInfo

func (m *TopicInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *TopicInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TopicInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *TopicInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *TopicInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *TopicInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *TopicInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *TopicInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TopicInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *TopicInfo) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

func (m *TopicInfo) GetTime() uint32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *TopicInfo) GetCompere() string {
	if m != nil {
		return m.Compere
	}
	return ""
}

type TopicRecord struct {
	Uid                  string         `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64         `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64          `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Operator             string         `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	Topic                string         `protobuf:"bytes,5,opt,name=topic,proto3" json:"topic,omitempty"`
	Date                 int64          `protobuf:"varint,6,opt,name=date,proto3" json:"date,omitempty"`
	State                uint32         `protobuf:"varint,7,opt,name=state,proto3" json:"state,omitempty"`
	Scene                string         `protobuf:"bytes,8,opt,name=scene,proto3" json:"scene,omitempty"`
	Sn                   string         `protobuf:"bytes,9,opt,name=sn,proto3" json:"sn,omitempty"`
	Compere              string         `protobuf:"bytes,10,opt,name=compere,proto3" json:"compere,omitempty"`
	Results              []*TopicResult `protobuf:"bytes,21,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TopicRecord) Reset()         { *m = TopicRecord{} }
func (m *TopicRecord) String() string { return proto.CompactTextString(m) }
func (*TopicRecord) ProtoMessage()    {}
func (*TopicRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac6e87e37c521a00, []int{1}
}

func (m *TopicRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicRecord.Unmarshal(m, b)
}
func (m *TopicRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicRecord.Marshal(b, m, deterministic)
}
func (m *TopicRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicRecord.Merge(m, src)
}
func (m *TopicRecord) XXX_Size() int {
	return xxx_messageInfo_TopicRecord.Size(m)
}
func (m *TopicRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicRecord.DiscardUnknown(m)
}

var xxx_messageInfo_TopicRecord proto.InternalMessageInfo

func (m *TopicRecord) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *TopicRecord) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TopicRecord) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *TopicRecord) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *TopicRecord) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *TopicRecord) GetDate() int64 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *TopicRecord) GetState() uint32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *TopicRecord) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

func (m *TopicRecord) GetSn() string {
	if m != nil {
		return m.Sn
	}
	return ""
}

func (m *TopicRecord) GetCompere() string {
	if m != nil {
		return m.Compere
	}
	return ""
}

func (m *TopicRecord) GetResults() []*TopicResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type TopicResult struct {
	Percent              uint32   `protobuf:"varint,1,opt,name=percent,proto3" json:"percent,omitempty"`
	Count                uint32   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	State                uint32   `protobuf:"varint,3,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicResult) Reset()         { *m = TopicResult{} }
func (m *TopicResult) String() string { return proto.CompactTextString(m) }
func (*TopicResult) ProtoMessage()    {}
func (*TopicResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac6e87e37c521a00, []int{2}
}

func (m *TopicResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicResult.Unmarshal(m, b)
}
func (m *TopicResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicResult.Marshal(b, m, deterministic)
}
func (m *TopicResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicResult.Merge(m, src)
}
func (m *TopicResult) XXX_Size() int {
	return xxx_messageInfo_TopicResult.Size(m)
}
func (m *TopicResult) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicResult.DiscardUnknown(m)
}

var xxx_messageInfo_TopicResult proto.InternalMessageInfo

func (m *TopicResult) GetPercent() uint32 {
	if m != nil {
		return m.Percent
	}
	return 0
}

func (m *TopicResult) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *TopicResult) GetState() uint32 {
	if m != nil {
		return m.State
	}
	return 0
}

type ReqTopicAdd struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Remark               string   `protobuf:"bytes,2,opt,name=remark,proto3" json:"remark,omitempty"`
	Scene                string   `protobuf:"bytes,3,opt,name=scene,proto3" json:"scene,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	Compere              string   `protobuf:"bytes,5,opt,name=compere,proto3" json:"compere,omitempty"`
	Time                 uint32   `protobuf:"varint,6,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqTopicAdd) Reset()         { *m = ReqTopicAdd{} }
func (m *ReqTopicAdd) String() string { return proto.CompactTextString(m) }
func (*ReqTopicAdd) ProtoMessage()    {}
func (*ReqTopicAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac6e87e37c521a00, []int{3}
}

func (m *ReqTopicAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqTopicAdd.Unmarshal(m, b)
}
func (m *ReqTopicAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqTopicAdd.Marshal(b, m, deterministic)
}
func (m *ReqTopicAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqTopicAdd.Merge(m, src)
}
func (m *ReqTopicAdd) XXX_Size() int {
	return xxx_messageInfo_ReqTopicAdd.Size(m)
}
func (m *ReqTopicAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqTopicAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqTopicAdd proto.InternalMessageInfo

func (m *ReqTopicAdd) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ReqTopicAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqTopicAdd) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

func (m *ReqTopicAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqTopicAdd) GetCompere() string {
	if m != nil {
		return m.Compere
	}
	return ""
}

func (m *ReqTopicAdd) GetTime() uint32 {
	if m != nil {
		return m.Time
	}
	return 0
}

type ReqTopicRecord struct {
	Topic                string         `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Date                 int64          `protobuf:"varint,2,opt,name=date,proto3" json:"date,omitempty"`
	State                uint32         `protobuf:"varint,3,opt,name=state,proto3" json:"state,omitempty"`
	Operator             string         `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	Sn                   string         `protobuf:"bytes,5,opt,name=sn,proto3" json:"sn,omitempty"`
	Results              []*TopicResult `protobuf:"bytes,11,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReqTopicRecord) Reset()         { *m = ReqTopicRecord{} }
func (m *ReqTopicRecord) String() string { return proto.CompactTextString(m) }
func (*ReqTopicRecord) ProtoMessage()    {}
func (*ReqTopicRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac6e87e37c521a00, []int{4}
}

func (m *ReqTopicRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqTopicRecord.Unmarshal(m, b)
}
func (m *ReqTopicRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqTopicRecord.Marshal(b, m, deterministic)
}
func (m *ReqTopicRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqTopicRecord.Merge(m, src)
}
func (m *ReqTopicRecord) XXX_Size() int {
	return xxx_messageInfo_ReqTopicRecord.Size(m)
}
func (m *ReqTopicRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqTopicRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ReqTopicRecord proto.InternalMessageInfo

func (m *ReqTopicRecord) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *ReqTopicRecord) GetDate() int64 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *ReqTopicRecord) GetState() uint32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *ReqTopicRecord) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqTopicRecord) GetSn() string {
	if m != nil {
		return m.Sn
	}
	return ""
}

func (m *ReqTopicRecord) GetResults() []*TopicResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type ReplyTopicInfo struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Info                 *TopicInfo   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyTopicInfo) Reset()         { *m = ReplyTopicInfo{} }
func (m *ReplyTopicInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyTopicInfo) ProtoMessage()    {}
func (*ReplyTopicInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac6e87e37c521a00, []int{5}
}

func (m *ReplyTopicInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyTopicInfo.Unmarshal(m, b)
}
func (m *ReplyTopicInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyTopicInfo.Marshal(b, m, deterministic)
}
func (m *ReplyTopicInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyTopicInfo.Merge(m, src)
}
func (m *ReplyTopicInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyTopicInfo.Size(m)
}
func (m *ReplyTopicInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyTopicInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyTopicInfo proto.InternalMessageInfo

func (m *ReplyTopicInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyTopicInfo) GetInfo() *TopicInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type ReplyTopicList struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                uint32       `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Pages                uint32       `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	Page                 uint32       `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	Number               uint32       `protobuf:"varint,5,opt,name=number,proto3" json:"number,omitempty"`
	List                 []*TopicInfo `protobuf:"bytes,16,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyTopicList) Reset()         { *m = ReplyTopicList{} }
func (m *ReplyTopicList) String() string { return proto.CompactTextString(m) }
func (*ReplyTopicList) ProtoMessage()    {}
func (*ReplyTopicList) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac6e87e37c521a00, []int{6}
}

func (m *ReplyTopicList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyTopicList.Unmarshal(m, b)
}
func (m *ReplyTopicList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyTopicList.Marshal(b, m, deterministic)
}
func (m *ReplyTopicList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyTopicList.Merge(m, src)
}
func (m *ReplyTopicList) XXX_Size() int {
	return xxx_messageInfo_ReplyTopicList.Size(m)
}
func (m *ReplyTopicList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyTopicList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyTopicList proto.InternalMessageInfo

func (m *ReplyTopicList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyTopicList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyTopicList) GetPages() uint32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *ReplyTopicList) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ReplyTopicList) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *ReplyTopicList) GetList() []*TopicInfo {
	if m != nil {
		return m.List
	}
	return nil
}

type ReplyTopicRecords struct {
	Status               *ReplyStatus   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                uint32         `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Pages                uint32         `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	Page                 uint32         `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	Number               uint32         `protobuf:"varint,5,opt,name=number,proto3" json:"number,omitempty"`
	List                 []*TopicRecord `protobuf:"bytes,16,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReplyTopicRecords) Reset()         { *m = ReplyTopicRecords{} }
func (m *ReplyTopicRecords) String() string { return proto.CompactTextString(m) }
func (*ReplyTopicRecords) ProtoMessage()    {}
func (*ReplyTopicRecords) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac6e87e37c521a00, []int{7}
}

func (m *ReplyTopicRecords) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyTopicRecords.Unmarshal(m, b)
}
func (m *ReplyTopicRecords) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyTopicRecords.Marshal(b, m, deterministic)
}
func (m *ReplyTopicRecords) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyTopicRecords.Merge(m, src)
}
func (m *ReplyTopicRecords) XXX_Size() int {
	return xxx_messageInfo_ReplyTopicRecords.Size(m)
}
func (m *ReplyTopicRecords) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyTopicRecords.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyTopicRecords proto.InternalMessageInfo

func (m *ReplyTopicRecords) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyTopicRecords) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyTopicRecords) GetPages() uint32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *ReplyTopicRecords) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ReplyTopicRecords) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *ReplyTopicRecords) GetList() []*TopicRecord {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*TopicInfo)(nil), "omo.msp.third.TopicInfo")
	proto.RegisterType((*TopicRecord)(nil), "omo.msp.third.TopicRecord")
	proto.RegisterType((*TopicResult)(nil), "omo.msp.third.TopicResult")
	proto.RegisterType((*ReqTopicAdd)(nil), "omo.msp.third.ReqTopicAdd")
	proto.RegisterType((*ReqTopicRecord)(nil), "omo.msp.third.ReqTopicRecord")
	proto.RegisterType((*ReplyTopicInfo)(nil), "omo.msp.third.ReplyTopicInfo")
	proto.RegisterType((*ReplyTopicList)(nil), "omo.msp.third.ReplyTopicList")
	proto.RegisterType((*ReplyTopicRecords)(nil), "omo.msp.third.ReplyTopicRecords")
}

func init() {
	proto.RegisterFile("proto/third/topic.proto", fileDescriptor_ac6e87e37c521a00)
}

var fileDescriptor_ac6e87e37c521a00 = []byte{
	// 733 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x96, 0xdf, 0x6e, 0xd3, 0x3e,
	0x14, 0xc7, 0x9b, 0xa4, 0x7f, 0x56, 0x67, 0xad, 0xf6, 0xb3, 0xf6, 0x03, 0xab, 0x62, 0x52, 0x95,
	0xab, 0x5e, 0xa0, 0x4e, 0x2a, 0xbc, 0xc0, 0x86, 0x58, 0x05, 0x1a, 0x20, 0x79, 0x70, 0xc3, 0x5d,
	0x97, 0x9c, 0x81, 0x45, 0x13, 0x07, 0xc7, 0x99, 0xb4, 0x67, 0xd9, 0x5b, 0xf0, 0x18, 0xdc, 0xf0,
	0x0e, 0x3c, 0x09, 0xf2, 0x71, 0x92, 0xa6, 0x2c, 0xcd, 0xa6, 0x5d, 0x71, 0xe7, 0xaf, 0x7d, 0x7c,
	0x7c, 0xbe, 0x9f, 0x73, 0xb2, 0x95, 0x3c, 0x4d, 0x95, 0xd4, 0xf2, 0x58, 0x7f, 0x15, 0x2a, 0x3a,
	0xd6, 0x32, 0x15, 0xe1, 0x1c, 0x77, 0xe8, 0x48, 0xc6, 0x72, 0x1e, 0x67, 0xe9, 0x1c, 0x8f, 0x26,
	0xac, 0x1e, 0x17, 0xca, 0x38, 0x96, 0x89, 0x0d, 0x0c, 0x6e, 0x5d, 0x32, 0xfc, 0x68, 0x2e, 0xbe,
	0x49, 0xae, 0x24, 0x3d, 0x20, 0x5e, 0x2e, 0x22, 0xe6, 0x4c, 0x9d, 0xd9, 0x90, 0x9b, 0x25, 0x1d,
	0x13, 0x57, 0x44, 0xcc, 0x9d, 0x3a, 0xb3, 0x2e, 0x77, 0x45, 0x44, 0x29, 0xe9, 0xea, 0x9b, 0x14,
	0x98, 0x37, 0x75, 0x66, 0x23, 0x8e, 0x6b, 0xca, 0xc8, 0x20, 0x54, 0xb0, 0xd2, 0x10, 0xb1, 0xee,
	0xd4, 0x99, 0x79, 0xbc, 0x94, 0xe6, 0x24, 0x4f, 0x23, 0x3c, 0xe9, 0xd9, 0x93, 0x42, 0x56, 0x77,
	0xa4, 0x62, 0x7d, 0x7c, 0xad, 0x94, 0x74, 0x42, 0xf6, 0x64, 0x0a, 0x0a, 0x8f, 0x06, 0x78, 0x54,
	0x69, 0xf3, 0x7a, 0xb2, 0x8a, 0x81, 0xed, 0xe1, 0x3e, 0xae, 0xe9, 0x13, 0xd2, 0x57, 0x10, 0xaf,
	0xd4, 0x37, 0x36, 0xc4, 0xdd, 0x42, 0xd1, 0x43, 0xd2, 0xcb, 0x42, 0x48, 0x80, 0x11, 0xdc, 0xb6,
	0x02, 0xeb, 0x17, 0x31, 0x30, 0xbf, 0xa8, 0x5f, 0xc4, 0xb6, 0x7e, 0x19, 0xa7, 0xa0, 0x80, 0xed,
	0x17, 0xb5, 0x58, 0x69, 0xe8, 0xf8, 0x48, 0x87, 0x43, 0x28, 0x55, 0xf4, 0x00, 0x3e, 0x35, 0x16,
	0xde, 0x36, 0x8b, 0xba, 0xaf, 0xee, 0x5f, 0xbe, 0x0e, 0x49, 0x0f, 0xbb, 0x87, 0x94, 0x86, 0xdc,
	0x0a, 0x53, 0xab, 0x81, 0x85, 0x80, 0x3c, 0x8e, 0x6b, 0x74, 0xa5, 0xcd, 0xe6, 0x00, 0x0d, 0x58,
	0xb1, 0xf1, 0xba, 0x57, 0xf7, 0x3a, 0x26, 0x6e, 0x96, 0x14, 0x54, 0xdc, 0x2c, 0xa9, 0xfb, 0x24,
	0x5b, 0x3e, 0xe9, 0x4b, 0x32, 0x50, 0x90, 0xe5, 0x6b, 0x9d, 0xb1, 0xff, 0xa7, 0xde, 0xcc, 0x5f,
	0x4c, 0xe6, 0x5b, 0x03, 0x34, 0x2f, 0x20, 0x98, 0x10, 0x5e, 0x86, 0x06, 0x17, 0x15, 0x1c, 0xa3,
	0x4d, 0xfa, 0x14, 0x54, 0x08, 0x89, 0x46, 0x40, 0x23, 0x5e, 0x4a, 0x53, 0x5e, 0x28, 0xf3, 0x44,
	0x23, 0xa7, 0x11, 0xb7, 0x62, 0x63, 0xc5, 0xab, 0x59, 0x09, 0x6e, 0x1d, 0xe2, 0x73, 0xf8, 0x8e,
	0x89, 0x4f, 0xa2, 0x08, 0xd1, 0x08, 0xbd, 0x86, 0x02, 0xba, 0x15, 0xb5, 0xa6, 0xbb, 0xcd, 0x4d,
	0xf7, 0xea, 0x20, 0xda, 0xd0, 0xd7, 0xa0, 0xf4, 0xb6, 0xa1, 0x94, 0xa3, 0xd2, 0xdf, 0x8c, 0x4a,
	0xf0, 0xc3, 0x21, 0xe3, 0xb2, 0xba, 0x62, 0x26, 0xaa, 0xde, 0x39, 0x4d, 0xbd, 0x73, 0x9b, 0x7a,
	0x57, 0x37, 0xdc, 0x5a, 0x9c, 0xed, 0x60, 0xaf, 0xea, 0x60, 0xad, 0x4f, 0xfe, 0xc3, 0xfb, 0xa4,
	0x4c, 0xcd, 0xe9, 0xfa, 0x66, 0xf3, 0x9d, 0x2f, 0x48, 0xdf, 0x3c, 0x9e, 0x67, 0x58, 0xf4, 0xdd,
	0x34, 0x18, 0x7e, 0x81, 0x11, 0xbc, 0x88, 0xa4, 0xcf, 0x49, 0x57, 0x24, 0x57, 0x12, 0x1d, 0xf9,
	0x0b, 0xd6, 0xf4, 0xb0, 0xc9, 0xcd, 0x31, 0x2a, 0xf8, 0xe9, 0xd4, 0x1f, 0x3d, 0x17, 0x99, 0x7e,
	0xd4, 0xa3, 0x08, 0x57, 0xaf, 0xd6, 0xe5, 0xe4, 0xa0, 0x30, 0xbb, 0xe9, 0xea, 0x0b, 0x64, 0x25,
	0x48, 0x14, 0x06, 0xb9, 0x59, 0x20, 0xc4, 0x11, 0xc7, 0xb5, 0x99, 0x93, 0x24, 0x8f, 0x2f, 0x41,
	0x21, 0xc4, 0x11, 0x2f, 0x94, 0x31, 0xb3, 0x16, 0x99, 0x66, 0x07, 0x48, 0xb1, 0xc5, 0x8c, 0x89,
	0x0a, 0x7e, 0x39, 0xe4, 0xbf, 0x8d, 0x19, 0xdb, 0xf7, 0xec, 0x9f, 0xf2, 0x33, 0xdf, 0xf2, 0xb3,
	0x63, 0x2a, 0x4c, 0xd9, 0xd6, 0xd1, 0xe2, 0x77, 0x97, 0xec, 0xe3, 0xee, 0x05, 0xa8, 0x6b, 0x11,
	0x02, 0x7d, 0x4d, 0xfa, 0x27, 0x51, 0xf4, 0xc1, 0x7c, 0x2c, 0x77, 0x6c, 0x54, 0x1f, 0xe3, 0xe4,
	0xa8, 0xc9, 0x62, 0x45, 0x2b, 0xe8, 0x98, 0x34, 0x4b, 0xd0, 0x3b, 0xd2, 0xe4, 0x90, 0x69, 0x13,
	0x77, 0x7f, 0x9a, 0x73, 0xe2, 0x2f, 0x41, 0x9f, 0xde, 0x9c, 0x89, 0xb5, 0x06, 0x45, 0x9f, 0x35,
	0xe7, 0xb2, 0xa7, 0x2d, 0xd9, 0xcc, 0xd8, 0x05, 0x1d, 0xfa, 0x8e, 0xec, 0x2f, 0x41, 0x9b, 0x4e,
	0x88, 0x4c, 0x8b, 0xf0, 0x31, 0xe9, 0xaa, 0xcb, 0x41, 0x87, 0xbe, 0x25, 0xe3, 0x4f, 0xf8, 0x5f,
	0xec, 0xbe, 0xfa, 0x6c, 0xd4, 0x84, 0x35, 0x25, 0x2c, 0x8c, 0xbe, 0x22, 0x43, 0x0e, 0xb1, 0xbc,
	0x86, 0xfb, 0x90, 0xb5, 0x25, 0x39, 0x23, 0xc3, 0x93, 0x28, 0x2a, 0xfe, 0x1c, 0x1d, 0xed, 0x68,
	0x9f, 0x3d, 0x6e, 0xcd, 0xf3, 0x9e, 0x90, 0x25, 0xe8, 0x72, 0xbc, 0xdb, 0x29, 0x4d, 0x77, 0x42,
	0x2f, 0xee, 0x07, 0x9d, 0xd3, 0xd1, 0x67, 0xbf, 0xf6, 0xbb, 0xe3, 0xb2, 0x8f, 0xe2, 0xc5, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x53, 0x47, 0xb0, 0x92, 0xb5, 0x08, 0x00, 0x00,
}