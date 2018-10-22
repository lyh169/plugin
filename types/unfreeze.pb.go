// Code generated by protoc-gen-go. DO NOT EDIT.
// source: unfreeze.proto

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	unfreeze.proto

It has these top-level messages:
	Unfreeze
	UnfreezeAction
	UnfreezeCreate
	UnfreezeWithdraw
	UnfreezeTerminate
	ReceiptUnfreeze
	QueryUnfreezeWithdraw
	ReplyQueryUnfreezeWithdraw
*/
package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Unfreeze struct {
	// 解冻交易ID（唯一识别码）
	UnfreezeID string `protobuf:"bytes,1,opt,name=unfreezeID" json:"unfreezeID,omitempty"`
	// 开始时间
	StartTime int64 `protobuf:"varint,2,opt,name=startTime" json:"startTime,omitempty"`
	// 币种
	TokenName string `protobuf:"bytes,3,opt,name=tokenName" json:"tokenName,omitempty"`
	// 冻结总额
	TotalCount int64 `protobuf:"varint,4,opt,name=totalCount" json:"totalCount,omitempty"`
	// 发币人地址
	Initiator string `protobuf:"bytes,5,opt,name=initiator" json:"initiator,omitempty"`
	// 收币人地址
	Beneficiary string `protobuf:"bytes,6,opt,name=beneficiary" json:"beneficiary,omitempty"`
	// 解冻间隔
	Period int64 `protobuf:"varint,7,opt,name=period" json:"period,omitempty"`
	// 解冻方式（百分比；固额） 1 百分比 -> 2 固额
	Means int32 `protobuf:"varint,8,opt,name=means" json:"means,omitempty"`
	// 解冻数量：若为百分比解冻方式该字段值为百分比乘以100，若为固额该字段值为币数量
	Amount int64 `protobuf:"varint,9,opt,name=amount" json:"amount,omitempty"`
	// 已解冻次数
	WithdrawTimes int32 `protobuf:"varint,10,opt,name=withdrawTimes" json:"withdrawTimes,omitempty"`
	// 解冻剩余币数
	Remaining int64 `protobuf:"varint,11,opt,name=remaining" json:"remaining,omitempty"`
}

func (m *Unfreeze) Reset()                    { *m = Unfreeze{} }
func (m *Unfreeze) String() string            { return proto.CompactTextString(m) }
func (*Unfreeze) ProtoMessage()               {}
func (*Unfreeze) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Unfreeze) GetUnfreezeID() string {
	if m != nil {
		return m.UnfreezeID
	}
	return ""
}

func (m *Unfreeze) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *Unfreeze) GetTokenName() string {
	if m != nil {
		return m.TokenName
	}
	return ""
}

func (m *Unfreeze) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *Unfreeze) GetInitiator() string {
	if m != nil {
		return m.Initiator
	}
	return ""
}

func (m *Unfreeze) GetBeneficiary() string {
	if m != nil {
		return m.Beneficiary
	}
	return ""
}

func (m *Unfreeze) GetPeriod() int64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *Unfreeze) GetMeans() int32 {
	if m != nil {
		return m.Means
	}
	return 0
}

func (m *Unfreeze) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Unfreeze) GetWithdrawTimes() int32 {
	if m != nil {
		return m.WithdrawTimes
	}
	return 0
}

func (m *Unfreeze) GetRemaining() int64 {
	if m != nil {
		return m.Remaining
	}
	return 0
}

// message for execs.unfreeze
type UnfreezeAction struct {
	// Types that are valid to be assigned to Value:
	//	*UnfreezeAction_Create
	//	*UnfreezeAction_Withdraw
	//	*UnfreezeAction_Terminate
	Value isUnfreezeAction_Value `protobuf_oneof:"value"`
	Ty    int32                  `protobuf:"varint,4,opt,name=ty" json:"ty,omitempty"`
}

func (m *UnfreezeAction) Reset()                    { *m = UnfreezeAction{} }
func (m *UnfreezeAction) String() string            { return proto.CompactTextString(m) }
func (*UnfreezeAction) ProtoMessage()               {}
func (*UnfreezeAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isUnfreezeAction_Value interface {
	isUnfreezeAction_Value()
}

type UnfreezeAction_Create struct {
	Create *UnfreezeCreate `protobuf:"bytes,1,opt,name=create,oneof"`
}
type UnfreezeAction_Withdraw struct {
	Withdraw *UnfreezeWithdraw `protobuf:"bytes,2,opt,name=withdraw,oneof"`
}
type UnfreezeAction_Terminate struct {
	Terminate *UnfreezeTerminate `protobuf:"bytes,3,opt,name=terminate,oneof"`
}

func (*UnfreezeAction_Create) isUnfreezeAction_Value()    {}
func (*UnfreezeAction_Withdraw) isUnfreezeAction_Value()  {}
func (*UnfreezeAction_Terminate) isUnfreezeAction_Value() {}

func (m *UnfreezeAction) GetValue() isUnfreezeAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *UnfreezeAction) GetCreate() *UnfreezeCreate {
	if x, ok := m.GetValue().(*UnfreezeAction_Create); ok {
		return x.Create
	}
	return nil
}

func (m *UnfreezeAction) GetWithdraw() *UnfreezeWithdraw {
	if x, ok := m.GetValue().(*UnfreezeAction_Withdraw); ok {
		return x.Withdraw
	}
	return nil
}

func (m *UnfreezeAction) GetTerminate() *UnfreezeTerminate {
	if x, ok := m.GetValue().(*UnfreezeAction_Terminate); ok {
		return x.Terminate
	}
	return nil
}

func (m *UnfreezeAction) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*UnfreezeAction) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _UnfreezeAction_OneofMarshaler, _UnfreezeAction_OneofUnmarshaler, _UnfreezeAction_OneofSizer, []interface{}{
		(*UnfreezeAction_Create)(nil),
		(*UnfreezeAction_Withdraw)(nil),
		(*UnfreezeAction_Terminate)(nil),
	}
}

func _UnfreezeAction_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*UnfreezeAction)
	// value
	switch x := m.Value.(type) {
	case *UnfreezeAction_Create:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Create); err != nil {
			return err
		}
	case *UnfreezeAction_Withdraw:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Withdraw); err != nil {
			return err
		}
	case *UnfreezeAction_Terminate:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Terminate); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("UnfreezeAction.Value has unexpected type %T", x)
	}
	return nil
}

func _UnfreezeAction_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*UnfreezeAction)
	switch tag {
	case 1: // value.create
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UnfreezeCreate)
		err := b.DecodeMessage(msg)
		m.Value = &UnfreezeAction_Create{msg}
		return true, err
	case 2: // value.withdraw
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UnfreezeWithdraw)
		err := b.DecodeMessage(msg)
		m.Value = &UnfreezeAction_Withdraw{msg}
		return true, err
	case 3: // value.terminate
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UnfreezeTerminate)
		err := b.DecodeMessage(msg)
		m.Value = &UnfreezeAction_Terminate{msg}
		return true, err
	default:
		return false, nil
	}
}

func _UnfreezeAction_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*UnfreezeAction)
	// value
	switch x := m.Value.(type) {
	case *UnfreezeAction_Create:
		s := proto.Size(x.Create)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UnfreezeAction_Withdraw:
		s := proto.Size(x.Withdraw)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UnfreezeAction_Terminate:
		s := proto.Size(x.Terminate)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// action
type UnfreezeCreate struct {
	StartTime   int64  `protobuf:"varint,1,opt,name=startTime" json:"startTime,omitempty"`
	TokenName   string `protobuf:"bytes,2,opt,name=tokenName" json:"tokenName,omitempty"`
	TotalCount  int64  `protobuf:"varint,3,opt,name=totalCount" json:"totalCount,omitempty"`
	Beneficiary string `protobuf:"bytes,4,opt,name=beneficiary" json:"beneficiary,omitempty"`
	Period      int64  `protobuf:"varint,5,opt,name=period" json:"period,omitempty"`
	Means       int32  `protobuf:"varint,6,opt,name=means" json:"means,omitempty"`
	Amount      int64  `protobuf:"varint,7,opt,name=amount" json:"amount,omitempty"`
}

func (m *UnfreezeCreate) Reset()                    { *m = UnfreezeCreate{} }
func (m *UnfreezeCreate) String() string            { return proto.CompactTextString(m) }
func (*UnfreezeCreate) ProtoMessage()               {}
func (*UnfreezeCreate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UnfreezeCreate) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *UnfreezeCreate) GetTokenName() string {
	if m != nil {
		return m.TokenName
	}
	return ""
}

func (m *UnfreezeCreate) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *UnfreezeCreate) GetBeneficiary() string {
	if m != nil {
		return m.Beneficiary
	}
	return ""
}

func (m *UnfreezeCreate) GetPeriod() int64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *UnfreezeCreate) GetMeans() int32 {
	if m != nil {
		return m.Means
	}
	return 0
}

func (m *UnfreezeCreate) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type UnfreezeWithdraw struct {
	UnfreezeID string `protobuf:"bytes,1,opt,name=unfreezeID" json:"unfreezeID,omitempty"`
}

func (m *UnfreezeWithdraw) Reset()                    { *m = UnfreezeWithdraw{} }
func (m *UnfreezeWithdraw) String() string            { return proto.CompactTextString(m) }
func (*UnfreezeWithdraw) ProtoMessage()               {}
func (*UnfreezeWithdraw) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UnfreezeWithdraw) GetUnfreezeID() string {
	if m != nil {
		return m.UnfreezeID
	}
	return ""
}

type UnfreezeTerminate struct {
	UnfreezeID string `protobuf:"bytes,1,opt,name=unfreezeID" json:"unfreezeID,omitempty"`
}

func (m *UnfreezeTerminate) Reset()                    { *m = UnfreezeTerminate{} }
func (m *UnfreezeTerminate) String() string            { return proto.CompactTextString(m) }
func (*UnfreezeTerminate) ProtoMessage()               {}
func (*UnfreezeTerminate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UnfreezeTerminate) GetUnfreezeID() string {
	if m != nil {
		return m.UnfreezeID
	}
	return ""
}

// receipt
type ReceiptUnfreeze struct {
	UnfreezeID  string `protobuf:"bytes,1,opt,name=unfreezeID" json:"unfreezeID,omitempty"`
	Initiator   string `protobuf:"bytes,2,opt,name=initiator" json:"initiator,omitempty"`
	Beneficiary string `protobuf:"bytes,3,opt,name=beneficiary" json:"beneficiary,omitempty"`
	TokenName   string `protobuf:"bytes,4,opt,name=tokenName" json:"tokenName,omitempty"`
}

func (m *ReceiptUnfreeze) Reset()                    { *m = ReceiptUnfreeze{} }
func (m *ReceiptUnfreeze) String() string            { return proto.CompactTextString(m) }
func (*ReceiptUnfreeze) ProtoMessage()               {}
func (*ReceiptUnfreeze) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ReceiptUnfreeze) GetUnfreezeID() string {
	if m != nil {
		return m.UnfreezeID
	}
	return ""
}

func (m *ReceiptUnfreeze) GetInitiator() string {
	if m != nil {
		return m.Initiator
	}
	return ""
}

func (m *ReceiptUnfreeze) GetBeneficiary() string {
	if m != nil {
		return m.Beneficiary
	}
	return ""
}

func (m *ReceiptUnfreeze) GetTokenName() string {
	if m != nil {
		return m.TokenName
	}
	return ""
}

// query
type QueryUnfreezeWithdraw struct {
	UnfreezeID string `protobuf:"bytes,1,opt,name=unfreezeID" json:"unfreezeID,omitempty"`
}

func (m *QueryUnfreezeWithdraw) Reset()                    { *m = QueryUnfreezeWithdraw{} }
func (m *QueryUnfreezeWithdraw) String() string            { return proto.CompactTextString(m) }
func (*QueryUnfreezeWithdraw) ProtoMessage()               {}
func (*QueryUnfreezeWithdraw) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *QueryUnfreezeWithdraw) GetUnfreezeID() string {
	if m != nil {
		return m.UnfreezeID
	}
	return ""
}

type ReplyQueryUnfreezeWithdraw struct {
	UnfreezeID      string `protobuf:"bytes,1,opt,name=unfreezeID" json:"unfreezeID,omitempty"`
	AvailableAmount int64  `protobuf:"varint,2,opt,name=availableAmount" json:"availableAmount,omitempty"`
}

func (m *ReplyQueryUnfreezeWithdraw) Reset()                    { *m = ReplyQueryUnfreezeWithdraw{} }
func (m *ReplyQueryUnfreezeWithdraw) String() string            { return proto.CompactTextString(m) }
func (*ReplyQueryUnfreezeWithdraw) ProtoMessage()               {}
func (*ReplyQueryUnfreezeWithdraw) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ReplyQueryUnfreezeWithdraw) GetUnfreezeID() string {
	if m != nil {
		return m.UnfreezeID
	}
	return ""
}

func (m *ReplyQueryUnfreezeWithdraw) GetAvailableAmount() int64 {
	if m != nil {
		return m.AvailableAmount
	}
	return 0
}

func init() {
	proto.RegisterType((*Unfreeze)(nil), "types.Unfreeze")
	proto.RegisterType((*UnfreezeAction)(nil), "types.UnfreezeAction")
	proto.RegisterType((*UnfreezeCreate)(nil), "types.UnfreezeCreate")
	proto.RegisterType((*UnfreezeWithdraw)(nil), "types.UnfreezeWithdraw")
	proto.RegisterType((*UnfreezeTerminate)(nil), "types.UnfreezeTerminate")
	proto.RegisterType((*ReceiptUnfreeze)(nil), "types.ReceiptUnfreeze")
	proto.RegisterType((*QueryUnfreezeWithdraw)(nil), "types.QueryUnfreezeWithdraw")
	proto.RegisterType((*ReplyQueryUnfreezeWithdraw)(nil), "types.ReplyQueryUnfreezeWithdraw")
}

func init() { proto.RegisterFile("unfreeze.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x6d, 0x92, 0x26, 0x6d, 0xa7, 0xa2, 0x0b, 0x16, 0x0b, 0x16, 0x42, 0xa8, 0x8a, 0x38, 0xf4,
	0x54, 0xa4, 0xae, 0x10, 0x5c, 0x97, 0xe5, 0x50, 0x2e, 0x48, 0x58, 0x8b, 0x38, 0xbb, 0xdd, 0x29,
	0x58, 0x24, 0x4e, 0xe4, 0x4e, 0x77, 0x15, 0xfe, 0x04, 0x3f, 0x8c, 0x0b, 0x07, 0x7e, 0x10, 0x8a,
	0xdd, 0xb4, 0x69, 0xba, 0x6a, 0xd5, 0x3d, 0xfa, 0xcd, 0x7b, 0xfe, 0x78, 0x6f, 0xc6, 0x30, 0x58,
	0xe9, 0x85, 0x41, 0xfc, 0x85, 0xe3, 0xdc, 0x64, 0x94, 0xb1, 0x90, 0x8a, 0x1c, 0x97, 0xf1, 0x5f,
	0x1f, 0xba, 0x5f, 0xd7, 0x15, 0xf6, 0x0a, 0xa0, 0x62, 0x7d, 0xfa, 0xc8, 0xbd, 0xa1, 0x37, 0xea,
	0x89, 0x1a, 0xc2, 0x5e, 0x42, 0x6f, 0x49, 0xd2, 0xd0, 0xb5, 0x4a, 0x91, 0xfb, 0x43, 0x6f, 0x14,
	0x88, 0x2d, 0x50, 0x56, 0x29, 0xfb, 0x89, 0xfa, 0xb3, 0x4c, 0x91, 0x07, 0x56, 0xbc, 0x05, 0xca,
	0xbd, 0x29, 0x23, 0x99, 0x5c, 0x65, 0x2b, 0x4d, 0xbc, 0x6d, 0xc5, 0x35, 0xa4, 0x54, 0x2b, 0xad,
	0x48, 0x49, 0xca, 0x0c, 0x0f, 0x9d, 0x7a, 0x03, 0xb0, 0x21, 0xf4, 0x67, 0xa8, 0x71, 0xa1, 0xe6,
	0x4a, 0x9a, 0x82, 0x47, 0xb6, 0x5e, 0x87, 0xd8, 0x33, 0x88, 0x72, 0x34, 0x2a, 0xbb, 0xe1, 0x1d,
	0xbb, 0xf7, 0x7a, 0xc5, 0x9e, 0x42, 0x98, 0xa2, 0xd4, 0x4b, 0xde, 0x1d, 0x7a, 0xa3, 0x50, 0xb8,
	0x45, 0xc9, 0x96, 0xa9, 0xbd, 0x49, 0xcf, 0xb1, 0xdd, 0x8a, 0xbd, 0x86, 0x47, 0x77, 0x8a, 0x7e,
	0xdc, 0x18, 0x79, 0x57, 0xbe, 0x69, 0xc9, 0xc1, 0xaa, 0x76, 0xc1, 0xf2, 0xae, 0x06, 0x53, 0xa9,
	0xb4, 0xd2, 0xdf, 0x79, 0xdf, 0xf9, 0xb0, 0x01, 0xe2, 0x3f, 0x1e, 0x0c, 0x2a, 0x4b, 0x2f, 0xe7,
	0xa4, 0x32, 0xcd, 0xde, 0x40, 0x34, 0x37, 0x28, 0x09, 0xad, 0xa9, 0xfd, 0xc9, 0xf9, 0xd8, 0xba,
	0x3f, 0xae, 0x68, 0x57, 0xb6, 0x38, 0x6d, 0x89, 0x35, 0x8d, 0xbd, 0x85, 0x6e, 0x75, 0xa4, 0x35,
	0xba, 0x3f, 0x79, 0xde, 0x90, 0x7c, 0x5b, 0x97, 0xa7, 0x2d, 0xb1, 0xa1, 0xb2, 0xf7, 0xd0, 0x23,
	0x34, 0xa9, 0xd2, 0xe5, 0x51, 0x81, 0xd5, 0xf1, 0x86, 0xee, 0xba, 0xaa, 0x4f, 0x5b, 0x62, 0x4b,
	0x66, 0x03, 0xf0, 0xa9, 0xb0, 0xb1, 0x84, 0xc2, 0xa7, 0xe2, 0x43, 0x07, 0xc2, 0x5b, 0x99, 0xac,
	0x30, 0xfe, 0x57, 0x7b, 0x8d, 0xbb, 0xe6, 0x6e, 0x1b, 0x78, 0x07, 0xdb, 0xc0, 0x3f, 0xdc, 0x06,
	0xc1, 0x5e, 0x1b, 0x34, 0x82, 0x6e, 0x1f, 0x0a, 0x3a, 0xbc, 0x3f, 0xe8, 0xe8, 0xfe, 0xa0, 0x3b,
	0xf5, 0xa0, 0xe3, 0x09, 0x3c, 0x6e, 0x3a, 0x79, 0xac, 0xfd, 0xe3, 0x0b, 0x78, 0xb2, 0xe7, 0xe2,
	0x51, 0xd1, 0x6f, 0x0f, 0xce, 0x04, 0xce, 0x51, 0xe5, 0x74, 0xca, 0x9c, 0x6d, 0x67, 0xc1, 0x3f,
	0x32, 0x0b, 0xc1, 0xbe, 0x45, 0x3b, 0x11, 0xb4, 0x1b, 0x11, 0xc4, 0xef, 0xe0, 0xfc, 0xcb, 0x0a,
	0x4d, 0x71, 0xf2, 0xfb, 0x17, 0xf0, 0x42, 0x60, 0x9e, 0x14, 0x0f, 0x52, 0xb3, 0x11, 0x9c, 0xc9,
	0x5b, 0xa9, 0x12, 0x39, 0x4b, 0xf0, 0xd2, 0x45, 0xe2, 0xbe, 0x90, 0x26, 0x3c, 0x8b, 0xec, 0x0f,
	0x75, 0xf1, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x2c, 0x93, 0xab, 0xb3, 0x04, 0x00, 0x00,
}
