// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chaosdaemon.proto

package chaosdaemon

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Rule_Action int32

const (
	Rule_ADD    Rule_Action = 0
	Rule_DELETE Rule_Action = 1
)

var Rule_Action_name = map[int32]string{
	0: "ADD",
	1: "DELETE",
}

var Rule_Action_value = map[string]int32{
	"ADD":    0,
	"DELETE": 1,
}

func (x Rule_Action) String() string {
	return proto.EnumName(Rule_Action_name, int32(x))
}

func (Rule_Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_143136706133b591, []int{5, 0}
}

type Rule_Direction int32

const (
	Rule_INPUT  Rule_Direction = 0
	Rule_OUTPUT Rule_Direction = 1
)

var Rule_Direction_name = map[int32]string{
	0: "INPUT",
	1: "OUTPUT",
}

var Rule_Direction_value = map[string]int32{
	"INPUT":  0,
	"OUTPUT": 1,
}

func (x Rule_Direction) String() string {
	return proto.EnumName(Rule_Direction_name, int32(x))
}

func (Rule_Direction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_143136706133b591, []int{5, 1}
}

type NetemRequest struct {
	Netem                *Netem   `protobuf:"bytes,1,opt,name=netem,proto3" json:"netem,omitempty"`
	ContainerId          string   `protobuf:"bytes,2,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetemRequest) Reset()         { *m = NetemRequest{} }
func (m *NetemRequest) String() string { return proto.CompactTextString(m) }
func (*NetemRequest) ProtoMessage()    {}
func (*NetemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_143136706133b591, []int{0}
}

func (m *NetemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetemRequest.Unmarshal(m, b)
}
func (m *NetemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetemRequest.Marshal(b, m, deterministic)
}
func (m *NetemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetemRequest.Merge(m, src)
}
func (m *NetemRequest) XXX_Size() int {
	return xxx_messageInfo_NetemRequest.Size(m)
}
func (m *NetemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NetemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NetemRequest proto.InternalMessageInfo

func (m *NetemRequest) GetNetem() *Netem {
	if m != nil {
		return m.Netem
	}
	return nil
}

func (m *NetemRequest) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

type Netem struct {
	Time                 uint32   `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Jitter               uint32   `protobuf:"varint,2,opt,name=jitter,proto3" json:"jitter,omitempty"`
	DelayCorr            float32  `protobuf:"fixed32,3,opt,name=delay_corr,json=delayCorr,proto3" json:"delay_corr,omitempty"`
	Limit                uint32   `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	Loss                 float32  `protobuf:"fixed32,5,opt,name=loss,proto3" json:"loss,omitempty"`
	LossCorr             float32  `protobuf:"fixed32,6,opt,name=loss_corr,json=lossCorr,proto3" json:"loss_corr,omitempty"`
	Gap                  uint32   `protobuf:"varint,7,opt,name=gap,proto3" json:"gap,omitempty"`
	Duplicate            float32  `protobuf:"fixed32,8,opt,name=duplicate,proto3" json:"duplicate,omitempty"`
	DuplicateCorr        float32  `protobuf:"fixed32,9,opt,name=duplicate_corr,json=duplicateCorr,proto3" json:"duplicate_corr,omitempty"`
	Reorder              float32  `protobuf:"fixed32,10,opt,name=reorder,proto3" json:"reorder,omitempty"`
	ReorderCorr          float32  `protobuf:"fixed32,11,opt,name=reorder_corr,json=reorderCorr,proto3" json:"reorder_corr,omitempty"`
	Corrupt              float32  `protobuf:"fixed32,12,opt,name=corrupt,proto3" json:"corrupt,omitempty"`
	CorruptCorr          float32  `protobuf:"fixed32,13,opt,name=corrupt_corr,json=corruptCorr,proto3" json:"corrupt_corr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Netem) Reset()         { *m = Netem{} }
func (m *Netem) String() string { return proto.CompactTextString(m) }
func (*Netem) ProtoMessage()    {}
func (*Netem) Descriptor() ([]byte, []int) {
	return fileDescriptor_143136706133b591, []int{1}
}

func (m *Netem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Netem.Unmarshal(m, b)
}
func (m *Netem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Netem.Marshal(b, m, deterministic)
}
func (m *Netem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Netem.Merge(m, src)
}
func (m *Netem) XXX_Size() int {
	return xxx_messageInfo_Netem.Size(m)
}
func (m *Netem) XXX_DiscardUnknown() {
	xxx_messageInfo_Netem.DiscardUnknown(m)
}

var xxx_messageInfo_Netem proto.InternalMessageInfo

func (m *Netem) GetTime() uint32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Netem) GetJitter() uint32 {
	if m != nil {
		return m.Jitter
	}
	return 0
}

func (m *Netem) GetDelayCorr() float32 {
	if m != nil {
		return m.DelayCorr
	}
	return 0
}

func (m *Netem) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *Netem) GetLoss() float32 {
	if m != nil {
		return m.Loss
	}
	return 0
}

func (m *Netem) GetLossCorr() float32 {
	if m != nil {
		return m.LossCorr
	}
	return 0
}

func (m *Netem) GetGap() uint32 {
	if m != nil {
		return m.Gap
	}
	return 0
}

func (m *Netem) GetDuplicate() float32 {
	if m != nil {
		return m.Duplicate
	}
	return 0
}

func (m *Netem) GetDuplicateCorr() float32 {
	if m != nil {
		return m.DuplicateCorr
	}
	return 0
}

func (m *Netem) GetReorder() float32 {
	if m != nil {
		return m.Reorder
	}
	return 0
}

func (m *Netem) GetReorderCorr() float32 {
	if m != nil {
		return m.ReorderCorr
	}
	return 0
}

func (m *Netem) GetCorrupt() float32 {
	if m != nil {
		return m.Corrupt
	}
	return 0
}

func (m *Netem) GetCorruptCorr() float32 {
	if m != nil {
		return m.CorruptCorr
	}
	return 0
}

type IpSetRequest struct {
	Ipset                *IpSet   `protobuf:"bytes,1,opt,name=ipset,proto3" json:"ipset,omitempty"`
	ContainerId          string   `protobuf:"bytes,2,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpSetRequest) Reset()         { *m = IpSetRequest{} }
func (m *IpSetRequest) String() string { return proto.CompactTextString(m) }
func (*IpSetRequest) ProtoMessage()    {}
func (*IpSetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_143136706133b591, []int{2}
}

func (m *IpSetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpSetRequest.Unmarshal(m, b)
}
func (m *IpSetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpSetRequest.Marshal(b, m, deterministic)
}
func (m *IpSetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpSetRequest.Merge(m, src)
}
func (m *IpSetRequest) XXX_Size() int {
	return xxx_messageInfo_IpSetRequest.Size(m)
}
func (m *IpSetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IpSetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IpSetRequest proto.InternalMessageInfo

func (m *IpSetRequest) GetIpset() *IpSet {
	if m != nil {
		return m.Ipset
	}
	return nil
}

func (m *IpSetRequest) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

type IpSet struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Ips                  []string `protobuf:"bytes,2,rep,name=ips,proto3" json:"ips,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpSet) Reset()         { *m = IpSet{} }
func (m *IpSet) String() string { return proto.CompactTextString(m) }
func (*IpSet) ProtoMessage()    {}
func (*IpSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_143136706133b591, []int{3}
}

func (m *IpSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpSet.Unmarshal(m, b)
}
func (m *IpSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpSet.Marshal(b, m, deterministic)
}
func (m *IpSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpSet.Merge(m, src)
}
func (m *IpSet) XXX_Size() int {
	return xxx_messageInfo_IpSet.Size(m)
}
func (m *IpSet) XXX_DiscardUnknown() {
	xxx_messageInfo_IpSet.DiscardUnknown(m)
}

var xxx_messageInfo_IpSet proto.InternalMessageInfo

func (m *IpSet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *IpSet) GetIps() []string {
	if m != nil {
		return m.Ips
	}
	return nil
}

type IpTablesRequest struct {
	Rule                 *Rule    `protobuf:"bytes,1,opt,name=rule,proto3" json:"rule,omitempty"`
	ContainerId          string   `protobuf:"bytes,2,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpTablesRequest) Reset()         { *m = IpTablesRequest{} }
func (m *IpTablesRequest) String() string { return proto.CompactTextString(m) }
func (*IpTablesRequest) ProtoMessage()    {}
func (*IpTablesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_143136706133b591, []int{4}
}

func (m *IpTablesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpTablesRequest.Unmarshal(m, b)
}
func (m *IpTablesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpTablesRequest.Marshal(b, m, deterministic)
}
func (m *IpTablesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpTablesRequest.Merge(m, src)
}
func (m *IpTablesRequest) XXX_Size() int {
	return xxx_messageInfo_IpTablesRequest.Size(m)
}
func (m *IpTablesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IpTablesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IpTablesRequest proto.InternalMessageInfo

func (m *IpTablesRequest) GetRule() *Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (m *IpTablesRequest) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

type Rule struct {
	Action               Rule_Action    `protobuf:"varint,1,opt,name=action,proto3,enum=chaosdaemon.Rule_Action" json:"action,omitempty"`
	Direction            Rule_Direction `protobuf:"varint,2,opt,name=direction,proto3,enum=chaosdaemon.Rule_Direction" json:"direction,omitempty"`
	Set                  string         `protobuf:"bytes,3,opt,name=set,proto3" json:"set,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Rule) Reset()         { *m = Rule{} }
func (m *Rule) String() string { return proto.CompactTextString(m) }
func (*Rule) ProtoMessage()    {}
func (*Rule) Descriptor() ([]byte, []int) {
	return fileDescriptor_143136706133b591, []int{5}
}

func (m *Rule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rule.Unmarshal(m, b)
}
func (m *Rule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rule.Marshal(b, m, deterministic)
}
func (m *Rule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rule.Merge(m, src)
}
func (m *Rule) XXX_Size() int {
	return xxx_messageInfo_Rule.Size(m)
}
func (m *Rule) XXX_DiscardUnknown() {
	xxx_messageInfo_Rule.DiscardUnknown(m)
}

var xxx_messageInfo_Rule proto.InternalMessageInfo

func (m *Rule) GetAction() Rule_Action {
	if m != nil {
		return m.Action
	}
	return Rule_ADD
}

func (m *Rule) GetDirection() Rule_Direction {
	if m != nil {
		return m.Direction
	}
	return Rule_INPUT
}

func (m *Rule) GetSet() string {
	if m != nil {
		return m.Set
	}
	return ""
}

type TimeRequest struct {
	ContainerId          string   `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	Sec                  int64    `protobuf:"varint,2,opt,name=sec,proto3" json:"sec,omitempty"`
	Nsec                 int64    `protobuf:"varint,3,opt,name=nsec,proto3" json:"nsec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeRequest) Reset()         { *m = TimeRequest{} }
func (m *TimeRequest) String() string { return proto.CompactTextString(m) }
func (*TimeRequest) ProtoMessage()    {}
func (*TimeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_143136706133b591, []int{6}
}

func (m *TimeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeRequest.Unmarshal(m, b)
}
func (m *TimeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeRequest.Marshal(b, m, deterministic)
}
func (m *TimeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeRequest.Merge(m, src)
}
func (m *TimeRequest) XXX_Size() int {
	return xxx_messageInfo_TimeRequest.Size(m)
}
func (m *TimeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TimeRequest proto.InternalMessageInfo

func (m *TimeRequest) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

func (m *TimeRequest) GetSec() int64 {
	if m != nil {
		return m.Sec
	}
	return 0
}

func (m *TimeRequest) GetNsec() int64 {
	if m != nil {
		return m.Nsec
	}
	return 0
}

func init() {
	proto.RegisterEnum("chaosdaemon.Rule_Action", Rule_Action_name, Rule_Action_value)
	proto.RegisterEnum("chaosdaemon.Rule_Direction", Rule_Direction_name, Rule_Direction_value)
	proto.RegisterType((*NetemRequest)(nil), "chaosdaemon.NetemRequest")
	proto.RegisterType((*Netem)(nil), "chaosdaemon.Netem")
	proto.RegisterType((*IpSetRequest)(nil), "chaosdaemon.IpSetRequest")
	proto.RegisterType((*IpSet)(nil), "chaosdaemon.IpSet")
	proto.RegisterType((*IpTablesRequest)(nil), "chaosdaemon.IpTablesRequest")
	proto.RegisterType((*Rule)(nil), "chaosdaemon.Rule")
	proto.RegisterType((*TimeRequest)(nil), "chaosdaemon.TimeRequest")
}

func init() { proto.RegisterFile("chaosdaemon.proto", fileDescriptor_143136706133b591) }

var fileDescriptor_143136706133b591 = []byte{
	// 636 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdd, 0x4e, 0xdb, 0x4c,
	0x10, 0xc5, 0x71, 0x12, 0xf0, 0x98, 0xf0, 0xc1, 0xea, 0x13, 0x5a, 0xfe, 0xa4, 0xd4, 0x12, 0x52,
	0x6e, 0x1a, 0x2a, 0x7a, 0xd5, 0xab, 0x8a, 0x92, 0xb4, 0x8a, 0x54, 0x41, 0xb5, 0x84, 0xde, 0x70,
	0x81, 0x8c, 0x3d, 0xc0, 0x56, 0x4e, 0xd6, 0x5d, 0xaf, 0x2b, 0xf1, 0x4e, 0x7d, 0x90, 0x3e, 0x4b,
	0x9f, 0xa2, 0xda, 0x59, 0x3b, 0x0d, 0x10, 0x51, 0x24, 0xae, 0x3c, 0x73, 0xe6, 0x9c, 0xe3, 0xd9,
	0x99, 0xb5, 0x61, 0x23, 0xb9, 0x8d, 0x55, 0x91, 0xc6, 0x38, 0x51, 0xd3, 0x7e, 0xae, 0x95, 0x51,
	0x2c, 0x9c, 0x83, 0xb6, 0x77, 0x6e, 0x94, 0xba, 0xc9, 0xf0, 0x80, 0x4a, 0x57, 0xe5, 0xf5, 0x01,
	0x4e, 0x72, 0x73, 0xe7, 0x98, 0xd1, 0x05, 0xac, 0x9e, 0xa0, 0xc1, 0x89, 0xc0, 0xef, 0x25, 0x16,
	0x86, 0xf5, 0xa0, 0x35, 0xb5, 0x39, 0xf7, 0xba, 0x5e, 0x2f, 0x3c, 0x64, 0xfd, 0x79, 0x73, 0xc7,
	0x74, 0x04, 0xf6, 0x0a, 0x56, 0x13, 0x35, 0x35, 0xb1, 0x9c, 0xa2, 0xbe, 0x94, 0x29, 0x6f, 0x74,
	0xbd, 0x5e, 0x20, 0xc2, 0x19, 0x36, 0x4a, 0xa3, 0xdf, 0x0d, 0x68, 0x91, 0x86, 0x31, 0x68, 0x1a,
	0x39, 0x41, 0x72, 0xed, 0x08, 0x8a, 0xd9, 0x26, 0xb4, 0xbf, 0x49, 0x63, 0x50, 0x93, 0xb4, 0x23,
	0xaa, 0x8c, 0xed, 0x01, 0xa4, 0x98, 0xc5, 0x77, 0x97, 0x89, 0xd2, 0x9a, 0xfb, 0x5d, 0xaf, 0xd7,
	0x10, 0x01, 0x21, 0xc7, 0x4a, 0x6b, 0xf6, 0x3f, 0xb4, 0x32, 0x39, 0x91, 0x86, 0x37, 0x49, 0xe5,
	0x12, 0xfb, 0x82, 0x4c, 0x15, 0x05, 0x6f, 0x11, 0x9d, 0x62, 0xb6, 0x03, 0x81, 0x7d, 0x3a, 0x9f,
	0x36, 0x15, 0x56, 0x2c, 0x40, 0x36, 0xeb, 0xe0, 0xdf, 0xc4, 0x39, 0x5f, 0x26, 0x13, 0x1b, 0xb2,
	0x5d, 0x08, 0xd2, 0x32, 0xcf, 0x64, 0x12, 0x1b, 0xe4, 0x2b, 0xd5, 0x6b, 0x6b, 0x80, 0xed, 0xc3,
	0xda, 0x2c, 0x71, 0x8e, 0x01, 0x51, 0x3a, 0x33, 0x94, 0x6c, 0x39, 0x2c, 0x6b, 0x54, 0x3a, 0x45,
	0xcd, 0x81, 0xea, 0x75, 0x6a, 0xe7, 0x55, 0x85, 0x4e, 0x1e, 0x52, 0x39, 0xac, 0xb0, 0x5a, 0x6c,
	0x4b, 0x65, 0x6e, 0xf8, 0xaa, 0x13, 0x57, 0xa9, 0x1b, 0x36, 0x85, 0x4e, 0xdc, 0x71, 0xe2, 0x0a,
	0xb3, 0x62, 0xbb, 0xc9, 0x51, 0x7e, 0x86, 0x66, 0x6e, 0x93, 0x32, 0x2f, 0xd0, 0x2c, 0xdc, 0xa4,
	0x63, 0x3a, 0xc2, 0x73, 0x36, 0xf9, 0x1a, 0x5a, 0x24, 0xb1, 0x73, 0x9e, 0xc6, 0xd5, 0x22, 0x03,
	0x41, 0xb1, 0x1d, 0xa5, 0xcc, 0x0b, 0xde, 0xe8, 0xfa, 0xbd, 0x40, 0xd8, 0x30, 0xba, 0x80, 0xff,
	0x46, 0xf9, 0x38, 0xbe, 0xca, 0xb0, 0xa8, 0xdb, 0xd9, 0x87, 0xa6, 0x2e, 0x33, 0xac, 0xba, 0xd9,
	0xb8, 0xd7, 0x8d, 0x28, 0x33, 0x14, 0x54, 0x7e, 0x4e, 0x2f, 0xbf, 0x3c, 0x68, 0x5a, 0x05, 0x7b,
	0x03, 0xed, 0x38, 0x31, 0x52, 0x4d, 0xc9, 0x74, 0xed, 0x90, 0x3f, 0x32, 0xed, 0x1f, 0x51, 0x5d,
	0x54, 0x3c, 0xf6, 0x0e, 0x82, 0x54, 0x6a, 0x74, 0xa2, 0x06, 0x89, 0x76, 0x1e, 0x8b, 0x06, 0x35,
	0x45, 0xfc, 0x65, 0xdb, 0x43, 0xda, 0x61, 0xfa, 0xd4, 0x8f, 0x0d, 0xa3, 0x3d, 0x68, 0x3b, 0x7b,
	0xb6, 0x0c, 0xfe, 0xd1, 0x60, 0xb0, 0xbe, 0xc4, 0x00, 0xda, 0x83, 0xe1, 0xe7, 0xe1, 0x78, 0xb8,
	0xee, 0x45, 0x11, 0x04, 0x33, 0x23, 0x16, 0x40, 0x6b, 0x74, 0xf2, 0xe5, 0x7c, 0xec, 0x38, 0xa7,
	0xe7, 0x63, 0x1b, 0x7b, 0xd1, 0x57, 0x08, 0xc7, 0x72, 0x82, 0xf5, 0x8c, 0x1e, 0x1e, 0xde, 0x7b,
	0x74, 0x78, 0xd7, 0x46, 0x42, 0xbd, 0xfb, 0xb6, 0x8d, 0x84, 0x36, 0x62, 0x21, 0x9f, 0x20, 0x8a,
	0x0f, 0x7f, 0xfa, 0x10, 0x1e, 0xdb, 0x63, 0x0d, 0xe8, 0x58, 0xec, 0x3d, 0xac, 0x9c, 0xa1, 0x71,
	0x9f, 0xe2, 0xd6, 0x82, 0x4f, 0xda, 0xbd, 0x7f, 0x7b, 0xb3, 0xef, 0x7e, 0x15, 0xfd, 0xfa, 0x57,
	0xd1, 0x1f, 0xda, 0x5f, 0x45, 0xb4, 0xc4, 0x3e, 0x40, 0x38, 0xc0, 0x0c, 0x0d, 0xbe, 0xc0, 0xe3,
	0x08, 0xe0, 0x63, 0x56, 0x16, 0xb7, 0xee, 0x22, 0x6d, 0x2d, 0xb8, 0x8f, 0xff, 0xb4, 0xf8, 0x04,
	0x9d, 0xca, 0xc2, 0xd0, 0xe5, 0x62, 0xbb, 0x0f, 0x5c, 0xee, 0xdd, 0xb9, 0x27, 0x8c, 0x8e, 0xa1,
	0x73, 0x86, 0xc6, 0xce, 0xfe, 0xf4, 0xfa, 0xda, 0x7e, 0x03, 0xf7, 0xef, 0xce, 0xdc, 0x52, 0x9e,
	0xec, 0x66, 0x43, 0x60, 0xa2, 0x7e, 0xa0, 0x7e, 0x99, 0xd1, 0x55, 0x9b, 0x90, 0xb7, 0x7f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x36, 0xef, 0x72, 0x24, 0xca, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChaosDaemonClient is the client API for ChaosDaemon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChaosDaemonClient interface {
	SetNetem(ctx context.Context, in *NetemRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteNetem(ctx context.Context, in *NetemRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	FlushIpSet(ctx context.Context, in *IpSetRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	FlushIptables(ctx context.Context, in *IpTablesRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	SetTimeOffset(ctx context.Context, in *TimeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	RecoverTimeOffset(ctx context.Context, in *TimeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type chaosDaemonClient struct {
	cc *grpc.ClientConn
}

func NewChaosDaemonClient(cc *grpc.ClientConn) ChaosDaemonClient {
	return &chaosDaemonClient{cc}
}

func (c *chaosDaemonClient) SetNetem(ctx context.Context, in *NetemRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/chaosdaemon.ChaosDaemon/SetNetem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaosDaemonClient) DeleteNetem(ctx context.Context, in *NetemRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/chaosdaemon.ChaosDaemon/DeleteNetem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaosDaemonClient) FlushIpSet(ctx context.Context, in *IpSetRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/chaosdaemon.ChaosDaemon/FlushIpSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaosDaemonClient) FlushIptables(ctx context.Context, in *IpTablesRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/chaosdaemon.ChaosDaemon/FlushIptables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaosDaemonClient) SetTimeOffset(ctx context.Context, in *TimeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/chaosdaemon.ChaosDaemon/SetTimeOffset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaosDaemonClient) RecoverTimeOffset(ctx context.Context, in *TimeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/chaosdaemon.ChaosDaemon/RecoverTimeOffset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChaosDaemonServer is the server API for ChaosDaemon service.
type ChaosDaemonServer interface {
	SetNetem(context.Context, *NetemRequest) (*empty.Empty, error)
	DeleteNetem(context.Context, *NetemRequest) (*empty.Empty, error)
	FlushIpSet(context.Context, *IpSetRequest) (*empty.Empty, error)
	FlushIptables(context.Context, *IpTablesRequest) (*empty.Empty, error)
	SetTimeOffset(context.Context, *TimeRequest) (*empty.Empty, error)
	RecoverTimeOffset(context.Context, *TimeRequest) (*empty.Empty, error)
}

// UnimplementedChaosDaemonServer can be embedded to have forward compatible implementations.
type UnimplementedChaosDaemonServer struct {
}

func (*UnimplementedChaosDaemonServer) SetNetem(ctx context.Context, req *NetemRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetNetem not implemented")
}
func (*UnimplementedChaosDaemonServer) DeleteNetem(ctx context.Context, req *NetemRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNetem not implemented")
}
func (*UnimplementedChaosDaemonServer) FlushIpSet(ctx context.Context, req *IpSetRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FlushIpSet not implemented")
}
func (*UnimplementedChaosDaemonServer) FlushIptables(ctx context.Context, req *IpTablesRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FlushIptables not implemented")
}
func (*UnimplementedChaosDaemonServer) SetTimeOffset(ctx context.Context, req *TimeRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTimeOffset not implemented")
}
func (*UnimplementedChaosDaemonServer) RecoverTimeOffset(ctx context.Context, req *TimeRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecoverTimeOffset not implemented")
}

func RegisterChaosDaemonServer(s *grpc.Server, srv ChaosDaemonServer) {
	s.RegisterService(&_ChaosDaemon_serviceDesc, srv)
}

func _ChaosDaemon_SetNetem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaosDaemonServer).SetNetem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chaosdaemon.ChaosDaemon/SetNetem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaosDaemonServer).SetNetem(ctx, req.(*NetemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChaosDaemon_DeleteNetem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaosDaemonServer).DeleteNetem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chaosdaemon.ChaosDaemon/DeleteNetem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaosDaemonServer).DeleteNetem(ctx, req.(*NetemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChaosDaemon_FlushIpSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IpSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaosDaemonServer).FlushIpSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chaosdaemon.ChaosDaemon/FlushIpSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaosDaemonServer).FlushIpSet(ctx, req.(*IpSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChaosDaemon_FlushIptables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IpTablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaosDaemonServer).FlushIptables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chaosdaemon.ChaosDaemon/FlushIptables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaosDaemonServer).FlushIptables(ctx, req.(*IpTablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChaosDaemon_SetTimeOffset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaosDaemonServer).SetTimeOffset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chaosdaemon.ChaosDaemon/SetTimeOffset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaosDaemonServer).SetTimeOffset(ctx, req.(*TimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChaosDaemon_RecoverTimeOffset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaosDaemonServer).RecoverTimeOffset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chaosdaemon.ChaosDaemon/RecoverTimeOffset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaosDaemonServer).RecoverTimeOffset(ctx, req.(*TimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChaosDaemon_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chaosdaemon.ChaosDaemon",
	HandlerType: (*ChaosDaemonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetNetem",
			Handler:    _ChaosDaemon_SetNetem_Handler,
		},
		{
			MethodName: "DeleteNetem",
			Handler:    _ChaosDaemon_DeleteNetem_Handler,
		},
		{
			MethodName: "FlushIpSet",
			Handler:    _ChaosDaemon_FlushIpSet_Handler,
		},
		{
			MethodName: "FlushIptables",
			Handler:    _ChaosDaemon_FlushIptables_Handler,
		},
		{
			MethodName: "SetTimeOffset",
			Handler:    _ChaosDaemon_SetTimeOffset_Handler,
		},
		{
			MethodName: "RecoverTimeOffset",
			Handler:    _ChaosDaemon_RecoverTimeOffset_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chaosdaemon.proto",
}
