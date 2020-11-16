// Code generated by protoc-gen-go. DO NOT EDIT.
// source: jvm.proto

package types

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

//合约对象信息
type JVMContractObject struct {
	Addr                 string           `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Data                 *JVMContractData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *JVMContractObject) Reset()         { *m = JVMContractObject{} }
func (m *JVMContractObject) String() string { return proto.CompactTextString(m) }
func (*JVMContractObject) ProtoMessage()    {}
func (*JVMContractObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea1047aed7729a0b, []int{0}
}

func (m *JVMContractObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JVMContractObject.Unmarshal(m, b)
}
func (m *JVMContractObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JVMContractObject.Marshal(b, m, deterministic)
}
func (m *JVMContractObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JVMContractObject.Merge(m, src)
}
func (m *JVMContractObject) XXX_Size() int {
	return xxx_messageInfo_JVMContractObject.Size(m)
}
func (m *JVMContractObject) XXX_DiscardUnknown() {
	xxx_messageInfo_JVMContractObject.DiscardUnknown(m)
}

var xxx_messageInfo_JVMContractObject proto.InternalMessageInfo

func (m *JVMContractObject) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *JVMContractObject) GetData() *JVMContractData {
	if m != nil {
		return m.Data
	}
	return nil
}

// 存放合约固定数据
type JVMContractData struct {
	Creator              string   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Addr                 string   `protobuf:"bytes,4,opt,name=addr,proto3" json:"addr,omitempty"`
	Code                 []byte   `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
	CodeHash             []byte   `protobuf:"bytes,6,opt,name=codeHash,proto3" json:"codeHash,omitempty"`
	Abi                  []byte   `protobuf:"bytes,7,opt,name=abi,proto3" json:"abi,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JVMContractData) Reset()         { *m = JVMContractData{} }
func (m *JVMContractData) String() string { return proto.CompactTextString(m) }
func (*JVMContractData) ProtoMessage()    {}
func (*JVMContractData) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea1047aed7729a0b, []int{1}
}

func (m *JVMContractData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JVMContractData.Unmarshal(m, b)
}
func (m *JVMContractData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JVMContractData.Marshal(b, m, deterministic)
}
func (m *JVMContractData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JVMContractData.Merge(m, src)
}
func (m *JVMContractData) XXX_Size() int {
	return xxx_messageInfo_JVMContractData.Size(m)
}
func (m *JVMContractData) XXX_DiscardUnknown() {
	xxx_messageInfo_JVMContractData.DiscardUnknown(m)
}

var xxx_messageInfo_JVMContractData proto.InternalMessageInfo

func (m *JVMContractData) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *JVMContractData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *JVMContractData) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *JVMContractData) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *JVMContractData) GetCodeHash() []byte {
	if m != nil {
		return m.CodeHash
	}
	return nil
}

func (m *JVMContractData) GetAbi() []byte {
	if m != nil {
		return m.Abi
	}
	return nil
}

type LogJVMContractData struct {
	Creator              string   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Addr                 string   `protobuf:"bytes,3,opt,name=addr,proto3" json:"addr,omitempty"`
	CodeHash             string   `protobuf:"bytes,4,opt,name=codeHash,proto3" json:"codeHash,omitempty"`
	AbiHash              string   `protobuf:"bytes,5,opt,name=abiHash,proto3" json:"abiHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogJVMContractData) Reset()         { *m = LogJVMContractData{} }
func (m *LogJVMContractData) String() string { return proto.CompactTextString(m) }
func (*LogJVMContractData) ProtoMessage()    {}
func (*LogJVMContractData) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea1047aed7729a0b, []int{2}
}

func (m *LogJVMContractData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogJVMContractData.Unmarshal(m, b)
}
func (m *LogJVMContractData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogJVMContractData.Marshal(b, m, deterministic)
}
func (m *LogJVMContractData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogJVMContractData.Merge(m, src)
}
func (m *LogJVMContractData) XXX_Size() int {
	return xxx_messageInfo_LogJVMContractData.Size(m)
}
func (m *LogJVMContractData) XXX_DiscardUnknown() {
	xxx_messageInfo_LogJVMContractData.DiscardUnknown(m)
}

var xxx_messageInfo_LogJVMContractData proto.InternalMessageInfo

func (m *LogJVMContractData) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *LogJVMContractData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogJVMContractData) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *LogJVMContractData) GetCodeHash() string {
	if m != nil {
		return m.CodeHash
	}
	return ""
}

func (m *LogJVMContractData) GetAbiHash() string {
	if m != nil {
		return m.AbiHash
	}
	return ""
}

type JVMContractAction struct {
	// Types that are valid to be assigned to Value:
	//	*JVMContractAction_CreateJvmContract
	//	*JVMContractAction_CallJvmContract
	//	*JVMContractAction_UpdateJvmContract
	Value                isJVMContractAction_Value `protobuf_oneof:"value"`
	Ty                   int32                     `protobuf:"varint,4,opt,name=ty,proto3" json:"ty,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *JVMContractAction) Reset()         { *m = JVMContractAction{} }
func (m *JVMContractAction) String() string { return proto.CompactTextString(m) }
func (*JVMContractAction) ProtoMessage()    {}
func (*JVMContractAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea1047aed7729a0b, []int{3}
}

func (m *JVMContractAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JVMContractAction.Unmarshal(m, b)
}
func (m *JVMContractAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JVMContractAction.Marshal(b, m, deterministic)
}
func (m *JVMContractAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JVMContractAction.Merge(m, src)
}
func (m *JVMContractAction) XXX_Size() int {
	return xxx_messageInfo_JVMContractAction.Size(m)
}
func (m *JVMContractAction) XXX_DiscardUnknown() {
	xxx_messageInfo_JVMContractAction.DiscardUnknown(m)
}

var xxx_messageInfo_JVMContractAction proto.InternalMessageInfo

type isJVMContractAction_Value interface {
	isJVMContractAction_Value()
}

type JVMContractAction_CreateJvmContract struct {
	CreateJvmContract *CreateJvmContract `protobuf:"bytes,1,opt,name=createJvmContract,proto3,oneof"`
}

type JVMContractAction_CallJvmContract struct {
	CallJvmContract *CallJvmContract `protobuf:"bytes,2,opt,name=callJvmContract,proto3,oneof"`
}

type JVMContractAction_UpdateJvmContract struct {
	UpdateJvmContract *UpdateJvmContract `protobuf:"bytes,3,opt,name=updateJvmContract,proto3,oneof"`
}

func (*JVMContractAction_CreateJvmContract) isJVMContractAction_Value() {}

func (*JVMContractAction_CallJvmContract) isJVMContractAction_Value() {}

func (*JVMContractAction_UpdateJvmContract) isJVMContractAction_Value() {}

func (m *JVMContractAction) GetValue() isJVMContractAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *JVMContractAction) GetCreateJvmContract() *CreateJvmContract {
	if x, ok := m.GetValue().(*JVMContractAction_CreateJvmContract); ok {
		return x.CreateJvmContract
	}
	return nil
}

func (m *JVMContractAction) GetCallJvmContract() *CallJvmContract {
	if x, ok := m.GetValue().(*JVMContractAction_CallJvmContract); ok {
		return x.CallJvmContract
	}
	return nil
}

func (m *JVMContractAction) GetUpdateJvmContract() *UpdateJvmContract {
	if x, ok := m.GetValue().(*JVMContractAction_UpdateJvmContract); ok {
		return x.UpdateJvmContract
	}
	return nil
}

func (m *JVMContractAction) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*JVMContractAction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*JVMContractAction_CreateJvmContract)(nil),
		(*JVMContractAction_CallJvmContract)(nil),
		(*JVMContractAction_UpdateJvmContract)(nil),
	}
}

// 创建JVM合约
type CreateJvmContract struct {
	// 用户自定义Jvm合约名字，必须是user.Jvm.xxx的风格，且xxx由a-zA-Z0-9组成的4-16字符长度组成
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// 合约字节码
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateJvmContract) Reset()         { *m = CreateJvmContract{} }
func (m *CreateJvmContract) String() string { return proto.CompactTextString(m) }
func (*CreateJvmContract) ProtoMessage()    {}
func (*CreateJvmContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea1047aed7729a0b, []int{4}
}

func (m *CreateJvmContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateJvmContract.Unmarshal(m, b)
}
func (m *CreateJvmContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateJvmContract.Marshal(b, m, deterministic)
}
func (m *CreateJvmContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateJvmContract.Merge(m, src)
}
func (m *CreateJvmContract) XXX_Size() int {
	return xxx_messageInfo_CreateJvmContract.Size(m)
}
func (m *CreateJvmContract) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateJvmContract.DiscardUnknown(m)
}

var xxx_messageInfo_CreateJvmContract proto.InternalMessageInfo

func (m *CreateJvmContract) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateJvmContract) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

// 调用Jvm合约
type CallJvmContract struct {
	//合约名称
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	//执行参数
	ActionData           []string `protobuf:"bytes,2,rep,name=actionData,proto3" json:"actionData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallJvmContract) Reset()         { *m = CallJvmContract{} }
func (m *CallJvmContract) String() string { return proto.CompactTextString(m) }
func (*CallJvmContract) ProtoMessage()    {}
func (*CallJvmContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea1047aed7729a0b, []int{5}
}

func (m *CallJvmContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallJvmContract.Unmarshal(m, b)
}
func (m *CallJvmContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallJvmContract.Marshal(b, m, deterministic)
}
func (m *CallJvmContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallJvmContract.Merge(m, src)
}
func (m *CallJvmContract) XXX_Size() int {
	return xxx_messageInfo_CallJvmContract.Size(m)
}
func (m *CallJvmContract) XXX_DiscardUnknown() {
	xxx_messageInfo_CallJvmContract.DiscardUnknown(m)
}

var xxx_messageInfo_CallJvmContract proto.InternalMessageInfo

func (m *CallJvmContract) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CallJvmContract) GetActionData() []string {
	if m != nil {
		return m.ActionData
	}
	return nil
}

// 更新Jvm合约
type UpdateJvmContract struct {
	// 用户需要更新的Jvm合约
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// 合约字节码
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateJvmContract) Reset()         { *m = UpdateJvmContract{} }
func (m *UpdateJvmContract) String() string { return proto.CompactTextString(m) }
func (*UpdateJvmContract) ProtoMessage()    {}
func (*UpdateJvmContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea1047aed7729a0b, []int{6}
}

func (m *UpdateJvmContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateJvmContract.Unmarshal(m, b)
}
func (m *UpdateJvmContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateJvmContract.Marshal(b, m, deterministic)
}
func (m *UpdateJvmContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateJvmContract.Merge(m, src)
}
func (m *UpdateJvmContract) XXX_Size() int {
	return xxx_messageInfo_UpdateJvmContract.Size(m)
}
func (m *UpdateJvmContract) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateJvmContract.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateJvmContract proto.InternalMessageInfo

func (m *UpdateJvmContract) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateJvmContract) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

// 存放本地数据库的数据
type ReceiptLocalData struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	CurValue             []byte   `protobuf:"bytes,2,opt,name=curValue,proto3" json:"curValue,omitempty"`
	PreValue             []byte   `protobuf:"bytes,3,opt,name=preValue,proto3" json:"preValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiptLocalData) Reset()         { *m = ReceiptLocalData{} }
func (m *ReceiptLocalData) String() string { return proto.CompactTextString(m) }
func (*ReceiptLocalData) ProtoMessage()    {}
func (*ReceiptLocalData) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea1047aed7729a0b, []int{7}
}

func (m *ReceiptLocalData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiptLocalData.Unmarshal(m, b)
}
func (m *ReceiptLocalData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiptLocalData.Marshal(b, m, deterministic)
}
func (m *ReceiptLocalData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiptLocalData.Merge(m, src)
}
func (m *ReceiptLocalData) XXX_Size() int {
	return xxx_messageInfo_ReceiptLocalData.Size(m)
}
func (m *ReceiptLocalData) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiptLocalData.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiptLocalData proto.InternalMessageInfo

func (m *ReceiptLocalData) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *ReceiptLocalData) GetCurValue() []byte {
	if m != nil {
		return m.CurValue
	}
	return nil
}

func (m *ReceiptLocalData) GetPreValue() []byte {
	if m != nil {
		return m.PreValue
	}
	return nil
}

// 合约创建/调用日志
type ReceiptJVMContract struct {
	Caller               string   `protobuf:"bytes,1,opt,name=caller,proto3" json:"caller,omitempty"`
	ContractName         string   `protobuf:"bytes,2,opt,name=contractName,proto3" json:"contractName,omitempty"`
	ContractAddr         string   `protobuf:"bytes,3,opt,name=contractAddr,proto3" json:"contractAddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiptJVMContract) Reset()         { *m = ReceiptJVMContract{} }
func (m *ReceiptJVMContract) String() string { return proto.CompactTextString(m) }
func (*ReceiptJVMContract) ProtoMessage()    {}
func (*ReceiptJVMContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea1047aed7729a0b, []int{8}
}

func (m *ReceiptJVMContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiptJVMContract.Unmarshal(m, b)
}
func (m *ReceiptJVMContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiptJVMContract.Marshal(b, m, deterministic)
}
func (m *ReceiptJVMContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiptJVMContract.Merge(m, src)
}
func (m *ReceiptJVMContract) XXX_Size() int {
	return xxx_messageInfo_ReceiptJVMContract.Size(m)
}
func (m *ReceiptJVMContract) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiptJVMContract.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiptJVMContract proto.InternalMessageInfo

func (m *ReceiptJVMContract) GetCaller() string {
	if m != nil {
		return m.Caller
	}
	return ""
}

func (m *ReceiptJVMContract) GetContractName() string {
	if m != nil {
		return m.ContractName
	}
	return ""
}

func (m *ReceiptJVMContract) GetContractAddr() string {
	if m != nil {
		return m.ContractAddr
	}
	return ""
}

// 用于保存JVM只能合约中的状态数据变更
type JVMStateChangeItem struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	PreValue             []byte   `protobuf:"bytes,2,opt,name=preValue,proto3" json:"preValue,omitempty"`
	CurrentValue         []byte   `protobuf:"bytes,3,opt,name=currentValue,proto3" json:"currentValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JVMStateChangeItem) Reset()         { *m = JVMStateChangeItem{} }
func (m *JVMStateChangeItem) String() string { return proto.CompactTextString(m) }
func (*JVMStateChangeItem) ProtoMessage()    {}
func (*JVMStateChangeItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea1047aed7729a0b, []int{9}
}

func (m *JVMStateChangeItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JVMStateChangeItem.Unmarshal(m, b)
}
func (m *JVMStateChangeItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JVMStateChangeItem.Marshal(b, m, deterministic)
}
func (m *JVMStateChangeItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JVMStateChangeItem.Merge(m, src)
}
func (m *JVMStateChangeItem) XXX_Size() int {
	return xxx_messageInfo_JVMStateChangeItem.Size(m)
}
func (m *JVMStateChangeItem) XXX_DiscardUnknown() {
	xxx_messageInfo_JVMStateChangeItem.DiscardUnknown(m)
}

var xxx_messageInfo_JVMStateChangeItem proto.InternalMessageInfo

func (m *JVMStateChangeItem) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *JVMStateChangeItem) GetPreValue() []byte {
	if m != nil {
		return m.PreValue
	}
	return nil
}

func (m *JVMStateChangeItem) GetCurrentValue() []byte {
	if m != nil {
		return m.CurrentValue
	}
	return nil
}

type CheckJVMContractNameReq struct {
	JvmContractName      string   `protobuf:"bytes,1,opt,name=JvmContractName,proto3" json:"JvmContractName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckJVMContractNameReq) Reset()         { *m = CheckJVMContractNameReq{} }
func (m *CheckJVMContractNameReq) String() string { return proto.CompactTextString(m) }
func (*CheckJVMContractNameReq) ProtoMessage()    {}
func (*CheckJVMContractNameReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea1047aed7729a0b, []int{10}
}

func (m *CheckJVMContractNameReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckJVMContractNameReq.Unmarshal(m, b)
}
func (m *CheckJVMContractNameReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckJVMContractNameReq.Marshal(b, m, deterministic)
}
func (m *CheckJVMContractNameReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckJVMContractNameReq.Merge(m, src)
}
func (m *CheckJVMContractNameReq) XXX_Size() int {
	return xxx_messageInfo_CheckJVMContractNameReq.Size(m)
}
func (m *CheckJVMContractNameReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckJVMContractNameReq.DiscardUnknown(m)
}

var xxx_messageInfo_CheckJVMContractNameReq proto.InternalMessageInfo

func (m *CheckJVMContractNameReq) GetJvmContractName() string {
	if m != nil {
		return m.JvmContractName
	}
	return ""
}

type CheckJVMAddrResp struct {
	ExistAlready         bool     `protobuf:"varint,1,opt,name=existAlready,proto3" json:"existAlready,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckJVMAddrResp) Reset()         { *m = CheckJVMAddrResp{} }
func (m *CheckJVMAddrResp) String() string { return proto.CompactTextString(m) }
func (*CheckJVMAddrResp) ProtoMessage()    {}
func (*CheckJVMAddrResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea1047aed7729a0b, []int{11}
}

func (m *CheckJVMAddrResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckJVMAddrResp.Unmarshal(m, b)
}
func (m *CheckJVMAddrResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckJVMAddrResp.Marshal(b, m, deterministic)
}
func (m *CheckJVMAddrResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckJVMAddrResp.Merge(m, src)
}
func (m *CheckJVMAddrResp) XXX_Size() int {
	return xxx_messageInfo_CheckJVMAddrResp.Size(m)
}
func (m *CheckJVMAddrResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckJVMAddrResp.DiscardUnknown(m)
}

var xxx_messageInfo_CheckJVMAddrResp proto.InternalMessageInfo

func (m *CheckJVMAddrResp) GetExistAlready() bool {
	if m != nil {
		return m.ExistAlready
	}
	return false
}

type JVMQueryReq struct {
	Contract             string   `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
	Para                 []string `protobuf:"bytes,2,rep,name=para,proto3" json:"para,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JVMQueryReq) Reset()         { *m = JVMQueryReq{} }
func (m *JVMQueryReq) String() string { return proto.CompactTextString(m) }
func (*JVMQueryReq) ProtoMessage()    {}
func (*JVMQueryReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea1047aed7729a0b, []int{12}
}

func (m *JVMQueryReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JVMQueryReq.Unmarshal(m, b)
}
func (m *JVMQueryReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JVMQueryReq.Marshal(b, m, deterministic)
}
func (m *JVMQueryReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JVMQueryReq.Merge(m, src)
}
func (m *JVMQueryReq) XXX_Size() int {
	return xxx_messageInfo_JVMQueryReq.Size(m)
}
func (m *JVMQueryReq) XXX_DiscardUnknown() {
	xxx_messageInfo_JVMQueryReq.DiscardUnknown(m)
}

var xxx_messageInfo_JVMQueryReq proto.InternalMessageInfo

func (m *JVMQueryReq) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *JVMQueryReq) GetPara() []string {
	if m != nil {
		return m.Para
	}
	return nil
}

type JVMQueryResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Result               []string `protobuf:"bytes,2,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JVMQueryResponse) Reset()         { *m = JVMQueryResponse{} }
func (m *JVMQueryResponse) String() string { return proto.CompactTextString(m) }
func (*JVMQueryResponse) ProtoMessage()    {}
func (*JVMQueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea1047aed7729a0b, []int{13}
}

func (m *JVMQueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JVMQueryResponse.Unmarshal(m, b)
}
func (m *JVMQueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JVMQueryResponse.Marshal(b, m, deterministic)
}
func (m *JVMQueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JVMQueryResponse.Merge(m, src)
}
func (m *JVMQueryResponse) XXX_Size() int {
	return xxx_messageInfo_JVMQueryResponse.Size(m)
}
func (m *JVMQueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JVMQueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JVMQueryResponse proto.InternalMessageInfo

func (m *JVMQueryResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *JVMQueryResponse) GetResult() []string {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*JVMContractObject)(nil), "types.JVMContractObject")
	proto.RegisterType((*JVMContractData)(nil), "types.JVMContractData")
	proto.RegisterType((*LogJVMContractData)(nil), "types.LogJVMContractData")
	proto.RegisterType((*JVMContractAction)(nil), "types.JVMContractAction")
	proto.RegisterType((*CreateJvmContract)(nil), "types.CreateJvmContract")
	proto.RegisterType((*CallJvmContract)(nil), "types.CallJvmContract")
	proto.RegisterType((*UpdateJvmContract)(nil), "types.UpdateJvmContract")
	proto.RegisterType((*ReceiptLocalData)(nil), "types.ReceiptLocalData")
	proto.RegisterType((*ReceiptJVMContract)(nil), "types.ReceiptJVMContract")
	proto.RegisterType((*JVMStateChangeItem)(nil), "types.JVMStateChangeItem")
	proto.RegisterType((*CheckJVMContractNameReq)(nil), "types.CheckJVMContractNameReq")
	proto.RegisterType((*CheckJVMAddrResp)(nil), "types.CheckJVMAddrResp")
	proto.RegisterType((*JVMQueryReq)(nil), "types.JVMQueryReq")
	proto.RegisterType((*JVMQueryResponse)(nil), "types.JVMQueryResponse")
}

func init() {
	proto.RegisterFile("jvm.proto", fileDescriptor_ea1047aed7729a0b)
}

var fileDescriptor_ea1047aed7729a0b = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x3d, 0x8f, 0xda, 0x40,
	0x10, 0x3d, 0x63, 0x38, 0x8e, 0x01, 0x05, 0xd8, 0xe2, 0x62, 0xa5, 0x88, 0xd0, 0x56, 0x28, 0x05,
	0xc5, 0x45, 0x4a, 0x13, 0xa5, 0xe0, 0x4c, 0x24, 0x82, 0x8e, 0x8b, 0xb2, 0xa7, 0x50, 0xa5, 0x59,
	0xd6, 0x93, 0x83, 0x3b, 0x63, 0x3b, 0xeb, 0x35, 0x0a, 0x7d, 0xca, 0xb4, 0xf9, 0xbf, 0xd1, 0xae,
	0x6d, 0x58, 0x9b, 0x26, 0x45, 0x2a, 0xcf, 0xc7, 0xce, 0xcc, 0x9b, 0x99, 0x37, 0x86, 0xce, 0xd3,
	0x7e, 0x37, 0x49, 0x64, 0xac, 0x62, 0xd2, 0x52, 0x87, 0x04, 0x53, 0xfa, 0x00, 0xc3, 0xc5, 0x6a,
	0xe9, 0xc7, 0x91, 0x92, 0x5c, 0xa8, 0xcf, 0xeb, 0x27, 0x14, 0x8a, 0x10, 0x68, 0xf2, 0x20, 0x90,
	0x9e, 0x33, 0x72, 0xc6, 0x1d, 0x66, 0x64, 0xf2, 0x06, 0x9a, 0x01, 0x57, 0xdc, 0x6b, 0x8c, 0x9c,
	0x71, 0xf7, 0xe6, 0x7a, 0x62, 0xc2, 0x27, 0x56, 0xec, 0x8c, 0x2b, 0xce, 0xcc, 0x1b, 0xfa, 0xc7,
	0x81, 0x7e, 0xcd, 0x43, 0x3c, 0x68, 0x0b, 0x89, 0x5c, 0xc5, 0x65, 0xda, 0x52, 0xd5, 0xd5, 0x22,
	0xbe, 0x43, 0x93, 0xb9, 0xc3, 0x8c, 0x7c, 0x44, 0xd0, 0xb4, 0x10, 0x10, 0x68, 0x8a, 0x38, 0x40,
	0xaf, 0x35, 0x72, 0xc6, 0x3d, 0x66, 0x64, 0xf2, 0x0a, 0xae, 0xf4, 0x77, 0xce, 0xd3, 0x8d, 0x77,
	0x69, 0xec, 0x47, 0x9d, 0x0c, 0xc0, 0xe5, 0xeb, 0xad, 0xd7, 0x36, 0x66, 0x2d, 0xd2, 0xdf, 0x0e,
	0x90, 0xbb, 0xf8, 0xf1, 0xff, 0x40, 0x73, 0x2d, 0x68, 0x36, 0x8c, 0x1c, 0xf2, 0x09, 0x86, 0x07,
	0x6d, 0xbe, 0xde, 0x1a, 0x57, 0x2b, 0xcf, 0x5e, 0xa8, 0xf4, 0x57, 0xa3, 0x32, 0xfc, 0xa9, 0x50,
	0xdb, 0x38, 0x22, 0x73, 0x18, 0x9a, 0xf2, 0xb8, 0xd8, 0xef, 0x4a, 0x97, 0xc1, 0xd5, 0xbd, 0xf1,
	0x8a, 0xa9, 0xfb, 0x75, 0xff, 0xfc, 0x82, 0x9d, 0x07, 0x91, 0x5b, 0xe8, 0x0b, 0x1e, 0x86, 0x76,
	0x9e, 0xea, 0xf6, 0xfc, 0xaa, 0x77, 0x7e, 0xc1, 0xea, 0x01, 0x1a, 0x4d, 0x96, 0x04, 0x35, 0x34,
	0x6e, 0x05, 0xcd, 0xd7, 0xba, 0x5f, 0xa3, 0x39, 0x0b, 0x22, 0x2f, 0xa0, 0xa1, 0x0e, 0x66, 0x3a,
	0x2d, 0xd6, 0x50, 0x87, 0xdb, 0x36, 0xb4, 0xf6, 0x3c, 0xcc, 0x90, 0xbe, 0x87, 0xe1, 0x59, 0x43,
	0xc7, 0xc9, 0x3b, 0xd5, 0xc9, 0x1b, 0x02, 0x14, 0xdb, 0xd0, 0x32, 0xfd, 0x08, 0xfd, 0x5a, 0x17,
	0xfa, 0xd9, 0xbd, 0x15, 0xaa, 0x65, 0xf2, 0x1a, 0x80, 0x9b, 0xf1, 0xce, 0x72, 0x0e, 0xbb, 0xe3,
	0x0e, 0xb3, 0x2c, 0x1a, 0xc3, 0x59, 0x1b, 0xff, 0x8c, 0xe1, 0x1b, 0x0c, 0x18, 0x0a, 0xdc, 0x26,
	0xea, 0x2e, 0x16, 0x3c, 0x34, 0x9c, 0x1a, 0x80, 0xfb, 0x8c, 0x07, 0x13, 0xda, 0x63, 0x5a, 0x34,
	0x1c, 0xc9, 0xe4, 0x4a, 0xb7, 0x6c, 0xa2, 0x35, 0x55, 0x0b, 0x5d, 0xfb, 0x12, 0x89, 0xb9, 0xcf,
	0xcd, 0x7d, 0xa5, 0x4e, 0x15, 0x90, 0x22, 0xbb, 0xc5, 0x15, 0x72, 0x0d, 0x97, 0x7a, 0x55, 0x58,
	0x52, 0xb6, 0xd0, 0x08, 0x85, 0x9e, 0x28, 0xde, 0xdc, 0x9f, 0x98, 0x5b, 0xb1, 0xd9, 0x6f, 0xa6,
	0x27, 0x26, 0x57, 0x6c, 0xf4, 0x3b, 0x90, 0xc5, 0x6a, 0xf9, 0xa0, 0xb8, 0x42, 0x7f, 0xc3, 0xa3,
	0x47, 0xfc, 0xa4, 0x70, 0x67, 0x77, 0xd5, 0x39, 0x76, 0x75, 0x44, 0xde, 0xa8, 0x22, 0x37, 0x75,
	0x32, 0x29, 0x31, 0x52, 0x76, 0x67, 0x15, 0x1b, 0xf5, 0xe1, 0xa5, 0xbf, 0x41, 0xf1, 0x6c, 0xf5,
	0xa6, 0x31, 0x32, 0xfc, 0x41, 0xc6, 0xd0, 0xb7, 0xb6, 0x61, 0xad, 0xb4, 0x6e, 0xa6, 0xef, 0x60,
	0x50, 0x26, 0xd1, 0xe0, 0x19, 0xa6, 0x89, 0x2e, 0x8e, 0x3f, 0xb7, 0xa9, 0x9a, 0x86, 0x12, 0x79,
	0x90, 0x63, 0xbe, 0x62, 0x15, 0x1b, 0xfd, 0x00, 0xdd, 0xc5, 0x6a, 0xf9, 0x25, 0x43, 0x79, 0xd0,
	0x05, 0xcd, 0x15, 0x5b, 0x07, 0x67, 0xae, 0xf8, 0xc4, 0x85, 0x84, 0xcb, 0x92, 0x3a, 0x46, 0xa6,
	0x33, 0x18, 0x9c, 0xc2, 0xd3, 0x24, 0x8e, 0x52, 0xd4, 0xd7, 0x9e, 0x66, 0x42, 0x60, 0x9a, 0x16,
	0x15, 0x4b, 0x55, 0x6f, 0x4c, 0x62, 0x9a, 0x85, 0xaa, 0xc8, 0x51, 0x68, 0xeb, 0x4b, 0xf3, 0x3f,
	0x7e, 0xfb, 0x37, 0x00, 0x00, 0xff, 0xff, 0x59, 0x7e, 0xf4, 0xad, 0x9c, 0x05, 0x00, 0x00,
}
