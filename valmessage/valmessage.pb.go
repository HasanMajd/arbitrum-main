// Code generated by protoc-gen-go. DO NOT EDIT.
// source: valmessage.proto

package valmessage

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	protocol "github.com/offchainlabs/arb-util/protocol"
	value "github.com/offchainlabs/arb-util/value"
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

type Address struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34ccd35396e2606, []int{0}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type TokenTypeBuf struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenTypeBuf) Reset()         { *m = TokenTypeBuf{} }
func (m *TokenTypeBuf) String() string { return proto.CompactTextString(m) }
func (*TokenTypeBuf) ProtoMessage()    {}
func (*TokenTypeBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34ccd35396e2606, []int{1}
}

func (m *TokenTypeBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenTypeBuf.Unmarshal(m, b)
}
func (m *TokenTypeBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenTypeBuf.Marshal(b, m, deterministic)
}
func (m *TokenTypeBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenTypeBuf.Merge(m, src)
}
func (m *TokenTypeBuf) XXX_Size() int {
	return xxx_messageInfo_TokenTypeBuf.Size(m)
}
func (m *TokenTypeBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenTypeBuf.DiscardUnknown(m)
}

var xxx_messageInfo_TokenTypeBuf proto.InternalMessageInfo

func (m *TokenTypeBuf) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type VMConfiguration struct {
	GracePeriod           uint64               `protobuf:"varint,1,opt,name=grace_period,json=gracePeriod,proto3" json:"grace_period,omitempty"`
	EscrowRequired        *value.BigIntegerBuf `protobuf:"bytes,2,opt,name=escrow_required,json=escrowRequired,proto3" json:"escrow_required,omitempty"`
	EscrowCurrency        *Address             `protobuf:"bytes,3,opt,name=escrow_currency,json=escrowCurrency,proto3" json:"escrow_currency,omitempty"`
	AssertKeys            []*Address           `protobuf:"bytes,4,rep,name=assert_keys,json=assertKeys,proto3" json:"assert_keys,omitempty"`
	MaxExecutionStepCount uint32               `protobuf:"varint,5,opt,name=max_execution_step_count,json=maxExecutionStepCount,proto3" json:"max_execution_step_count,omitempty"`
	Owner                 *Address             `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}             `json:"-"`
	XXX_unrecognized      []byte               `json:"-"`
	XXX_sizecache         int32                `json:"-"`
}

func (m *VMConfiguration) Reset()         { *m = VMConfiguration{} }
func (m *VMConfiguration) String() string { return proto.CompactTextString(m) }
func (*VMConfiguration) ProtoMessage()    {}
func (*VMConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34ccd35396e2606, []int{2}
}

func (m *VMConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VMConfiguration.Unmarshal(m, b)
}
func (m *VMConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VMConfiguration.Marshal(b, m, deterministic)
}
func (m *VMConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VMConfiguration.Merge(m, src)
}
func (m *VMConfiguration) XXX_Size() int {
	return xxx_messageInfo_VMConfiguration.Size(m)
}
func (m *VMConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_VMConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_VMConfiguration proto.InternalMessageInfo

func (m *VMConfiguration) GetGracePeriod() uint64 {
	if m != nil {
		return m.GracePeriod
	}
	return 0
}

func (m *VMConfiguration) GetEscrowRequired() *value.BigIntegerBuf {
	if m != nil {
		return m.EscrowRequired
	}
	return nil
}

func (m *VMConfiguration) GetEscrowCurrency() *Address {
	if m != nil {
		return m.EscrowCurrency
	}
	return nil
}

func (m *VMConfiguration) GetAssertKeys() []*Address {
	if m != nil {
		return m.AssertKeys
	}
	return nil
}

func (m *VMConfiguration) GetMaxExecutionStepCount() uint32 {
	if m != nil {
		return m.MaxExecutionStepCount
	}
	return 0
}

func (m *VMConfiguration) GetOwner() *Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

type CreateVMValidatorRequest struct {
	Config               *VMConfiguration `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	VmId                 *value.HashBuf   `protobuf:"bytes,2,opt,name=vm_id,json=vmId,proto3" json:"vm_id,omitempty"`
	VmState              *value.HashBuf   `protobuf:"bytes,3,opt,name=vm_state,json=vmState,proto3" json:"vm_state,omitempty"`
	ChallengeManagerNum  uint32           `protobuf:"varint,4,opt,name=challengeManagerNum,proto3" json:"challengeManagerNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CreateVMValidatorRequest) Reset()         { *m = CreateVMValidatorRequest{} }
func (m *CreateVMValidatorRequest) String() string { return proto.CompactTextString(m) }
func (*CreateVMValidatorRequest) ProtoMessage()    {}
func (*CreateVMValidatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34ccd35396e2606, []int{3}
}

func (m *CreateVMValidatorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateVMValidatorRequest.Unmarshal(m, b)
}
func (m *CreateVMValidatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateVMValidatorRequest.Marshal(b, m, deterministic)
}
func (m *CreateVMValidatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateVMValidatorRequest.Merge(m, src)
}
func (m *CreateVMValidatorRequest) XXX_Size() int {
	return xxx_messageInfo_CreateVMValidatorRequest.Size(m)
}
func (m *CreateVMValidatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateVMValidatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateVMValidatorRequest proto.InternalMessageInfo

func (m *CreateVMValidatorRequest) GetConfig() *VMConfiguration {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *CreateVMValidatorRequest) GetVmId() *value.HashBuf {
	if m != nil {
		return m.VmId
	}
	return nil
}

func (m *CreateVMValidatorRequest) GetVmState() *value.HashBuf {
	if m != nil {
		return m.VmState
	}
	return nil
}

func (m *CreateVMValidatorRequest) GetChallengeManagerNum() uint32 {
	if m != nil {
		return m.ChallengeManagerNum
	}
	return 0
}

type UnanimousAssertionValidatorNotification struct {
	Accepted             bool     `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
	Signatures           [][]byte `protobuf:"bytes,2,rep,name=signatures,proto3" json:"signatures,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnanimousAssertionValidatorNotification) Reset() {
	*m = UnanimousAssertionValidatorNotification{}
}
func (m *UnanimousAssertionValidatorNotification) String() string { return proto.CompactTextString(m) }
func (*UnanimousAssertionValidatorNotification) ProtoMessage()    {}
func (*UnanimousAssertionValidatorNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34ccd35396e2606, []int{4}
}

func (m *UnanimousAssertionValidatorNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnanimousAssertionValidatorNotification.Unmarshal(m, b)
}
func (m *UnanimousAssertionValidatorNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnanimousAssertionValidatorNotification.Marshal(b, m, deterministic)
}
func (m *UnanimousAssertionValidatorNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnanimousAssertionValidatorNotification.Merge(m, src)
}
func (m *UnanimousAssertionValidatorNotification) XXX_Size() int {
	return xxx_messageInfo_UnanimousAssertionValidatorNotification.Size(m)
}
func (m *UnanimousAssertionValidatorNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_UnanimousAssertionValidatorNotification.DiscardUnknown(m)
}

var xxx_messageInfo_UnanimousAssertionValidatorNotification proto.InternalMessageInfo

func (m *UnanimousAssertionValidatorNotification) GetAccepted() bool {
	if m != nil {
		return m.Accepted
	}
	return false
}

func (m *UnanimousAssertionValidatorNotification) GetSignatures() [][]byte {
	if m != nil {
		return m.Signatures
	}
	return nil
}

type CreateVMFinalizedValidatorNotification struct {
	Approved             bool     `protobuf:"varint,1,opt,name=approved,proto3" json:"approved,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateVMFinalizedValidatorNotification) Reset() {
	*m = CreateVMFinalizedValidatorNotification{}
}
func (m *CreateVMFinalizedValidatorNotification) String() string { return proto.CompactTextString(m) }
func (*CreateVMFinalizedValidatorNotification) ProtoMessage()    {}
func (*CreateVMFinalizedValidatorNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34ccd35396e2606, []int{5}
}

func (m *CreateVMFinalizedValidatorNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateVMFinalizedValidatorNotification.Unmarshal(m, b)
}
func (m *CreateVMFinalizedValidatorNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateVMFinalizedValidatorNotification.Marshal(b, m, deterministic)
}
func (m *CreateVMFinalizedValidatorNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateVMFinalizedValidatorNotification.Merge(m, src)
}
func (m *CreateVMFinalizedValidatorNotification) XXX_Size() int {
	return xxx_messageInfo_CreateVMFinalizedValidatorNotification.Size(m)
}
func (m *CreateVMFinalizedValidatorNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateVMFinalizedValidatorNotification.DiscardUnknown(m)
}

var xxx_messageInfo_CreateVMFinalizedValidatorNotification proto.InternalMessageInfo

func (m *CreateVMFinalizedValidatorNotification) GetApproved() bool {
	if m != nil {
		return m.Approved
	}
	return false
}

type SignedMessage struct {
	Message              *protocol.MessageBuf `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Signature            []byte               `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SignedMessage) Reset()         { *m = SignedMessage{} }
func (m *SignedMessage) String() string { return proto.CompactTextString(m) }
func (*SignedMessage) ProtoMessage()    {}
func (*SignedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34ccd35396e2606, []int{6}
}

func (m *SignedMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedMessage.Unmarshal(m, b)
}
func (m *SignedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedMessage.Marshal(b, m, deterministic)
}
func (m *SignedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedMessage.Merge(m, src)
}
func (m *SignedMessage) XXX_Size() int {
	return xxx_messageInfo_SignedMessage.Size(m)
}
func (m *SignedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SignedMessage proto.InternalMessageInfo

func (m *SignedMessage) GetMessage() *protocol.MessageBuf {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *SignedMessage) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type UnanimousAssertionValidatorRequest struct {
	BeforeHash           *value.HashBuf          `protobuf:"bytes,1,opt,name=beforeHash,proto3" json:"beforeHash,omitempty"`
	BeforeInbox          *value.HashBuf          `protobuf:"bytes,2,opt,name=beforeInbox,proto3" json:"beforeInbox,omitempty"`
	SequenceNum          uint64                  `protobuf:"varint,3,opt,name=sequenceNum,proto3" json:"sequenceNum,omitempty"`
	TimeBounds           *protocol.TimeBoundsBuf `protobuf:"bytes,4,opt,name=timeBounds,proto3" json:"timeBounds,omitempty"`
	SignedMessages       []*SignedMessage        `protobuf:"bytes,5,rep,name=signedMessages,proto3" json:"signedMessages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *UnanimousAssertionValidatorRequest) Reset()         { *m = UnanimousAssertionValidatorRequest{} }
func (m *UnanimousAssertionValidatorRequest) String() string { return proto.CompactTextString(m) }
func (*UnanimousAssertionValidatorRequest) ProtoMessage()    {}
func (*UnanimousAssertionValidatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34ccd35396e2606, []int{7}
}

func (m *UnanimousAssertionValidatorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnanimousAssertionValidatorRequest.Unmarshal(m, b)
}
func (m *UnanimousAssertionValidatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnanimousAssertionValidatorRequest.Marshal(b, m, deterministic)
}
func (m *UnanimousAssertionValidatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnanimousAssertionValidatorRequest.Merge(m, src)
}
func (m *UnanimousAssertionValidatorRequest) XXX_Size() int {
	return xxx_messageInfo_UnanimousAssertionValidatorRequest.Size(m)
}
func (m *UnanimousAssertionValidatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnanimousAssertionValidatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnanimousAssertionValidatorRequest proto.InternalMessageInfo

func (m *UnanimousAssertionValidatorRequest) GetBeforeHash() *value.HashBuf {
	if m != nil {
		return m.BeforeHash
	}
	return nil
}

func (m *UnanimousAssertionValidatorRequest) GetBeforeInbox() *value.HashBuf {
	if m != nil {
		return m.BeforeInbox
	}
	return nil
}

func (m *UnanimousAssertionValidatorRequest) GetSequenceNum() uint64 {
	if m != nil {
		return m.SequenceNum
	}
	return 0
}

func (m *UnanimousAssertionValidatorRequest) GetTimeBounds() *protocol.TimeBoundsBuf {
	if m != nil {
		return m.TimeBounds
	}
	return nil
}

func (m *UnanimousAssertionValidatorRequest) GetSignedMessages() []*SignedMessage {
	if m != nil {
		return m.SignedMessages
	}
	return nil
}

type ValidatorRequest struct {
	RequestId *value.HashBuf `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Types that are valid to be assigned to Request:
	//	*ValidatorRequest_Unanimous
	//	*ValidatorRequest_UnanimousNotification
	//	*ValidatorRequest_Create
	//	*ValidatorRequest_CreateNotification
	Request              isValidatorRequest_Request `protobuf_oneof:"request"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ValidatorRequest) Reset()         { *m = ValidatorRequest{} }
func (m *ValidatorRequest) String() string { return proto.CompactTextString(m) }
func (*ValidatorRequest) ProtoMessage()    {}
func (*ValidatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34ccd35396e2606, []int{8}
}

func (m *ValidatorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidatorRequest.Unmarshal(m, b)
}
func (m *ValidatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidatorRequest.Marshal(b, m, deterministic)
}
func (m *ValidatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorRequest.Merge(m, src)
}
func (m *ValidatorRequest) XXX_Size() int {
	return xxx_messageInfo_ValidatorRequest.Size(m)
}
func (m *ValidatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorRequest proto.InternalMessageInfo

func (m *ValidatorRequest) GetRequestId() *value.HashBuf {
	if m != nil {
		return m.RequestId
	}
	return nil
}

type isValidatorRequest_Request interface {
	isValidatorRequest_Request()
}

type ValidatorRequest_Unanimous struct {
	Unanimous *UnanimousAssertionValidatorRequest `protobuf:"bytes,2,opt,name=unanimous,proto3,oneof"`
}

type ValidatorRequest_UnanimousNotification struct {
	UnanimousNotification *UnanimousAssertionValidatorNotification `protobuf:"bytes,3,opt,name=unanimousNotification,proto3,oneof"`
}

type ValidatorRequest_Create struct {
	Create *CreateVMValidatorRequest `protobuf:"bytes,4,opt,name=create,proto3,oneof"`
}

type ValidatorRequest_CreateNotification struct {
	CreateNotification *CreateVMFinalizedValidatorNotification `protobuf:"bytes,5,opt,name=createNotification,proto3,oneof"`
}

func (*ValidatorRequest_Unanimous) isValidatorRequest_Request() {}

func (*ValidatorRequest_UnanimousNotification) isValidatorRequest_Request() {}

func (*ValidatorRequest_Create) isValidatorRequest_Request() {}

func (*ValidatorRequest_CreateNotification) isValidatorRequest_Request() {}

func (m *ValidatorRequest) GetRequest() isValidatorRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *ValidatorRequest) GetUnanimous() *UnanimousAssertionValidatorRequest {
	if x, ok := m.GetRequest().(*ValidatorRequest_Unanimous); ok {
		return x.Unanimous
	}
	return nil
}

func (m *ValidatorRequest) GetUnanimousNotification() *UnanimousAssertionValidatorNotification {
	if x, ok := m.GetRequest().(*ValidatorRequest_UnanimousNotification); ok {
		return x.UnanimousNotification
	}
	return nil
}

func (m *ValidatorRequest) GetCreate() *CreateVMValidatorRequest {
	if x, ok := m.GetRequest().(*ValidatorRequest_Create); ok {
		return x.Create
	}
	return nil
}

func (m *ValidatorRequest) GetCreateNotification() *CreateVMFinalizedValidatorNotification {
	if x, ok := m.GetRequest().(*ValidatorRequest_CreateNotification); ok {
		return x.CreateNotification
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ValidatorRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ValidatorRequest_Unanimous)(nil),
		(*ValidatorRequest_UnanimousNotification)(nil),
		(*ValidatorRequest_Create)(nil),
		(*ValidatorRequest_CreateNotification)(nil),
	}
}

type CreateVMFollowerResponse struct {
	Accepted             bool     `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateVMFollowerResponse) Reset()         { *m = CreateVMFollowerResponse{} }
func (m *CreateVMFollowerResponse) String() string { return proto.CompactTextString(m) }
func (*CreateVMFollowerResponse) ProtoMessage()    {}
func (*CreateVMFollowerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34ccd35396e2606, []int{9}
}

func (m *CreateVMFollowerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateVMFollowerResponse.Unmarshal(m, b)
}
func (m *CreateVMFollowerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateVMFollowerResponse.Marshal(b, m, deterministic)
}
func (m *CreateVMFollowerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateVMFollowerResponse.Merge(m, src)
}
func (m *CreateVMFollowerResponse) XXX_Size() int {
	return xxx_messageInfo_CreateVMFollowerResponse.Size(m)
}
func (m *CreateVMFollowerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateVMFollowerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateVMFollowerResponse proto.InternalMessageInfo

func (m *CreateVMFollowerResponse) GetAccepted() bool {
	if m != nil {
		return m.Accepted
	}
	return false
}

func (m *CreateVMFollowerResponse) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type UnanimousAssertionFollowerResponse struct {
	Accepted             bool           `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
	Signature            []byte         `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	AssertionHash        *value.HashBuf `protobuf:"bytes,3,opt,name=assertion_hash,json=assertionHash,proto3" json:"assertion_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UnanimousAssertionFollowerResponse) Reset()         { *m = UnanimousAssertionFollowerResponse{} }
func (m *UnanimousAssertionFollowerResponse) String() string { return proto.CompactTextString(m) }
func (*UnanimousAssertionFollowerResponse) ProtoMessage()    {}
func (*UnanimousAssertionFollowerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34ccd35396e2606, []int{10}
}

func (m *UnanimousAssertionFollowerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnanimousAssertionFollowerResponse.Unmarshal(m, b)
}
func (m *UnanimousAssertionFollowerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnanimousAssertionFollowerResponse.Marshal(b, m, deterministic)
}
func (m *UnanimousAssertionFollowerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnanimousAssertionFollowerResponse.Merge(m, src)
}
func (m *UnanimousAssertionFollowerResponse) XXX_Size() int {
	return xxx_messageInfo_UnanimousAssertionFollowerResponse.Size(m)
}
func (m *UnanimousAssertionFollowerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UnanimousAssertionFollowerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UnanimousAssertionFollowerResponse proto.InternalMessageInfo

func (m *UnanimousAssertionFollowerResponse) GetAccepted() bool {
	if m != nil {
		return m.Accepted
	}
	return false
}

func (m *UnanimousAssertionFollowerResponse) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *UnanimousAssertionFollowerResponse) GetAssertionHash() *value.HashBuf {
	if m != nil {
		return m.AssertionHash
	}
	return nil
}

type FollowerResponse struct {
	RequestId *value.HashBuf `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Types that are valid to be assigned to Response:
	//	*FollowerResponse_Create
	//	*FollowerResponse_Unanimous
	Response             isFollowerResponse_Response `protobuf_oneof:"response"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *FollowerResponse) Reset()         { *m = FollowerResponse{} }
func (m *FollowerResponse) String() string { return proto.CompactTextString(m) }
func (*FollowerResponse) ProtoMessage()    {}
func (*FollowerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34ccd35396e2606, []int{11}
}

func (m *FollowerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FollowerResponse.Unmarshal(m, b)
}
func (m *FollowerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FollowerResponse.Marshal(b, m, deterministic)
}
func (m *FollowerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FollowerResponse.Merge(m, src)
}
func (m *FollowerResponse) XXX_Size() int {
	return xxx_messageInfo_FollowerResponse.Size(m)
}
func (m *FollowerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FollowerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FollowerResponse proto.InternalMessageInfo

func (m *FollowerResponse) GetRequestId() *value.HashBuf {
	if m != nil {
		return m.RequestId
	}
	return nil
}

type isFollowerResponse_Response interface {
	isFollowerResponse_Response()
}

type FollowerResponse_Create struct {
	Create *CreateVMFollowerResponse `protobuf:"bytes,2,opt,name=create,proto3,oneof"`
}

type FollowerResponse_Unanimous struct {
	Unanimous *UnanimousAssertionFollowerResponse `protobuf:"bytes,3,opt,name=unanimous,proto3,oneof"`
}

func (*FollowerResponse_Create) isFollowerResponse_Response() {}

func (*FollowerResponse_Unanimous) isFollowerResponse_Response() {}

func (m *FollowerResponse) GetResponse() isFollowerResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *FollowerResponse) GetCreate() *CreateVMFollowerResponse {
	if x, ok := m.GetResponse().(*FollowerResponse_Create); ok {
		return x.Create
	}
	return nil
}

func (m *FollowerResponse) GetUnanimous() *UnanimousAssertionFollowerResponse {
	if x, ok := m.GetResponse().(*FollowerResponse_Unanimous); ok {
		return x.Unanimous
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FollowerResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FollowerResponse_Create)(nil),
		(*FollowerResponse_Unanimous)(nil),
	}
}

func init() {
	proto.RegisterType((*Address)(nil), "valmessage.Address")
	proto.RegisterType((*TokenTypeBuf)(nil), "valmessage.TokenTypeBuf")
	proto.RegisterType((*VMConfiguration)(nil), "valmessage.VMConfiguration")
	proto.RegisterType((*CreateVMValidatorRequest)(nil), "valmessage.CreateVMValidatorRequest")
	proto.RegisterType((*UnanimousAssertionValidatorNotification)(nil), "valmessage.UnanimousAssertionValidatorNotification")
	proto.RegisterType((*CreateVMFinalizedValidatorNotification)(nil), "valmessage.CreateVMFinalizedValidatorNotification")
	proto.RegisterType((*SignedMessage)(nil), "valmessage.SignedMessage")
	proto.RegisterType((*UnanimousAssertionValidatorRequest)(nil), "valmessage.UnanimousAssertionValidatorRequest")
	proto.RegisterType((*ValidatorRequest)(nil), "valmessage.ValidatorRequest")
	proto.RegisterType((*CreateVMFollowerResponse)(nil), "valmessage.CreateVMFollowerResponse")
	proto.RegisterType((*UnanimousAssertionFollowerResponse)(nil), "valmessage.UnanimousAssertionFollowerResponse")
	proto.RegisterType((*FollowerResponse)(nil), "valmessage.FollowerResponse")
}

func init() { proto.RegisterFile("valmessage.proto", fileDescriptor_b34ccd35396e2606) }

var fileDescriptor_b34ccd35396e2606 = []byte{
	// 866 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc7, 0x9b, 0xa4, 0xe9, 0xc7, 0x49, 0xda, 0x2d, 0xb3, 0x5d, 0xad, 0x29, 0x08, 0x82, 0x59,
	0x41, 0xf7, 0x62, 0x93, 0x2a, 0x05, 0xed, 0x0d, 0x20, 0x35, 0x85, 0x55, 0x2a, 0xd4, 0x0a, 0xb9,
	0xa5, 0x17, 0x48, 0x28, 0x9a, 0xd8, 0x27, 0xce, 0xa8, 0xf6, 0x8c, 0x99, 0x19, 0xa7, 0x2d, 0xef,
	0xc1, 0x1b, 0x70, 0xcf, 0x83, 0x70, 0xc5, 0x03, 0xf0, 0x2e, 0xc8, 0xe3, 0x71, 0xec, 0x66, 0x93,
	0xb2, 0x48, 0xdc, 0x58, 0x9e, 0x33, 0xff, 0x73, 0xe6, 0xcc, 0xef, 0x9c, 0x33, 0xb0, 0x37, 0xa3,
	0x51, 0x8c, 0x4a, 0xd1, 0x10, 0xbb, 0x89, 0x14, 0x5a, 0x10, 0x28, 0x2d, 0x07, 0xcf, 0x8d, 0xc9,
	0x17, 0x51, 0xaf, 0xf8, 0xc9, 0x45, 0x07, 0xef, 0xcd, 0x68, 0x94, 0x62, 0xcf, 0x7c, 0x73, 0x93,
	0xfb, 0x31, 0x6c, 0x9e, 0x04, 0x81, 0x44, 0xa5, 0xc8, 0x3e, 0x34, 0xcd, 0x8e, 0x53, 0xeb, 0xd4,
	0x0e, 0xdb, 0x5e, 0xbe, 0x70, 0x5f, 0x40, 0xfb, 0x4a, 0xdc, 0x20, 0xbf, 0xba, 0x4f, 0x70, 0x90,
	0x4e, 0x56, 0xa8, 0xfe, 0xac, 0xc3, 0x93, 0xeb, 0xf3, 0x53, 0xc1, 0x27, 0x2c, 0x4c, 0x25, 0xd5,
	0x4c, 0x70, 0xf2, 0x09, 0xb4, 0x43, 0x49, 0x7d, 0x1c, 0x25, 0x28, 0x99, 0x08, 0x8c, 0xc3, 0xba,
	0xd7, 0x32, 0xb6, 0x1f, 0x8c, 0x89, 0x7c, 0x0d, 0x4f, 0x50, 0xf9, 0x52, 0xdc, 0x8e, 0x24, 0xfe,
	0x92, 0x32, 0x89, 0x81, 0x53, 0xef, 0xd4, 0x0e, 0x5b, 0xfd, 0xfd, 0x6e, 0x9e, 0xe4, 0x80, 0x85,
	0x67, 0x5c, 0x63, 0x88, 0x72, 0x90, 0x4e, 0xbc, 0xdd, 0x5c, 0xec, 0x59, 0x2d, 0xf9, 0x6a, 0xee,
	0xee, 0xa7, 0x52, 0x22, 0xf7, 0xef, 0x9d, 0x86, 0x71, 0x7f, 0xda, 0xad, 0x00, 0xb2, 0xf7, 0x2b,
	0xbc, 0x4f, 0xad, 0x94, 0x7c, 0x01, 0x2d, 0xaa, 0x14, 0x4a, 0x3d, 0xba, 0xc1, 0x7b, 0xe5, 0xac,
	0x77, 0x1a, 0xab, 0x3c, 0x21, 0xd7, 0x7d, 0x8f, 0xf7, 0x8a, 0xbc, 0x06, 0x27, 0xa6, 0x77, 0x23,
	0xbc, 0x43, 0x3f, 0xcd, 0xae, 0x39, 0x52, 0x1a, 0x93, 0x91, 0x2f, 0x52, 0xae, 0x9d, 0x66, 0xa7,
	0x76, 0xb8, 0xe3, 0x3d, 0x8b, 0xe9, 0xdd, 0x77, 0xc5, 0xf6, 0xa5, 0xc6, 0xe4, 0x34, 0xdb, 0x24,
	0x2f, 0xa1, 0x29, 0x6e, 0x39, 0x4a, 0x67, 0x63, 0x75, 0x8a, 0xb9, 0xc2, 0xfd, 0xab, 0x06, 0xce,
	0xa9, 0x44, 0xaa, 0xf1, 0xfa, 0xfc, 0x9a, 0x46, 0x2c, 0xa0, 0x5a, 0xc8, 0xec, 0xd6, 0xa8, 0x34,
	0x39, 0x86, 0x0d, 0xdf, 0x70, 0x36, 0x40, 0x5b, 0xfd, 0x0f, 0xaa, 0x81, 0x16, 0x6a, 0xe0, 0x59,
	0x29, 0xf9, 0x14, 0x9a, 0xb3, 0x78, 0xc4, 0x0a, 0xbc, 0xbb, 0x16, 0xef, 0x90, 0xaa, 0x69, 0x06,
	0x76, 0x7d, 0x16, 0x9f, 0x05, 0xe4, 0x25, 0x6c, 0xcd, 0xe2, 0x91, 0xd2, 0x54, 0xa3, 0xe5, 0xb8,
	0xa8, 0xdb, 0x9c, 0xc5, 0x97, 0xd9, 0x36, 0x39, 0x82, 0xa7, 0xfe, 0x94, 0x46, 0x11, 0xf2, 0x10,
	0xcf, 0x29, 0xa7, 0x21, 0xca, 0x8b, 0x34, 0x76, 0xd6, 0x0d, 0x80, 0x65, 0x5b, 0x2e, 0xc2, 0xe7,
	0x3f, 0x72, 0xca, 0x59, 0x2c, 0x52, 0x75, 0x62, 0x70, 0x32, 0xc1, 0xe7, 0x97, 0xbb, 0x10, 0x9a,
	0x4d, 0x98, 0x9f, 0x37, 0xce, 0x01, 0x6c, 0x51, 0xdf, 0xc7, 0x44, 0x63, 0xde, 0x34, 0x5b, 0xde,
	0x7c, 0x4d, 0x3e, 0x02, 0x50, 0x2c, 0xe4, 0x54, 0xa7, 0x12, 0x95, 0x53, 0xef, 0x34, 0x0e, 0xdb,
	0x5e, 0xc5, 0xe2, 0x7e, 0x0b, 0x9f, 0x15, 0xe4, 0xde, 0x30, 0x4e, 0x23, 0xf6, 0x2b, 0x06, 0xab,
	0x4f, 0x49, 0x12, 0x29, 0x66, 0x95, 0x53, 0xec, 0xda, 0xfd, 0x19, 0x76, 0x2e, 0x59, 0xc8, 0x31,
	0x38, 0xcf, 0xb9, 0x92, 0x2e, 0x6c, 0x5a, 0xc4, 0x96, 0xfa, 0x7e, 0x77, 0x3e, 0x5b, 0x56, 0x63,
	0xf8, 0x58, 0x11, 0xf9, 0x10, 0xb6, 0xe7, 0x49, 0x19, 0xe6, 0x6d, 0xaf, 0x34, 0xb8, 0xbf, 0xd7,
	0xc1, 0x7d, 0x04, 0x46, 0x51, 0xe9, 0x2e, 0xc0, 0x18, 0x27, 0x42, 0x62, 0x86, 0xdf, 0x9e, 0xbb,
	0x58, 0x91, 0x8a, 0x82, 0x1c, 0x41, 0x2b, 0x5f, 0x9d, 0xf1, 0xb1, 0xb8, 0x5b, 0x51, 0xea, 0xaa,
	0x84, 0x74, 0xa0, 0xa5, 0xb2, 0xc3, 0xb8, 0x8f, 0x59, 0xf9, 0x1a, 0xf9, 0x84, 0x56, 0x4c, 0xe4,
	0x35, 0x80, 0x66, 0x31, 0x0e, 0x44, 0xca, 0x03, 0x65, 0xea, 0xdb, 0xea, 0x3f, 0x2f, 0xef, 0x7e,
	0x35, 0xdf, 0x33, 0xc9, 0x94, 0x52, 0x72, 0x02, 0xbb, 0xaa, 0x8a, 0x50, 0x39, 0x4d, 0x33, 0x60,
	0xef, 0x57, 0xdb, 0xf5, 0x01, 0x64, 0x6f, 0xc1, 0xc1, 0xfd, 0xa3, 0x01, 0x7b, 0x6f, 0x41, 0x79,
	0x05, 0x20, 0xf3, 0xdf, 0xac, 0x9d, 0x97, 0x43, 0xd9, 0xb6, 0x8a, 0xb3, 0x80, 0x5c, 0xc0, 0x76,
	0x5a, 0x90, 0xb6, 0x44, 0xba, 0xd5, 0x0c, 0xfe, 0xbd, 0x0c, 0xc3, 0x35, 0xaf, 0x0c, 0x41, 0x6e,
	0xe0, 0xd9, 0x7c, 0x51, 0x6d, 0x27, 0x3b, 0x30, 0xc7, 0xef, 0x18, 0xbb, 0xea, 0x3a, 0x5c, 0xf3,
	0x96, 0xc7, 0x24, 0xdf, 0xc0, 0x86, 0x6f, 0x9a, 0xd9, 0x82, 0x7f, 0x51, 0x8d, 0xbe, 0xea, 0x81,
	0x18, 0xae, 0x79, 0xd6, 0x8b, 0x04, 0x40, 0xf2, 0xbf, 0x07, 0x99, 0x36, 0x4d, 0xac, 0xfe, 0xb2,
	0x58, 0x8f, 0x8f, 0xcc, 0x70, 0xcd, 0x5b, 0x12, 0x6f, 0xb0, 0x0d, 0x9b, 0x96, 0xb7, 0x7b, 0x55,
	0xbe, 0x5b, 0x6f, 0x44, 0x14, 0x89, 0x5b, 0x94, 0x1e, 0xaa, 0x44, 0x70, 0x85, 0x8f, 0x4e, 0xf5,
	0xe3, 0xe3, 0xf2, 0x5b, 0x6d, 0xd9, 0xb8, 0xfc, 0x7f, 0x07, 0x90, 0x2f, 0x61, 0x97, 0x16, 0x61,
	0x47, 0xd3, 0x6c, 0xd8, 0x96, 0x3f, 0x7f, 0x3b, 0x73, 0x55, 0x66, 0x71, 0xff, 0xae, 0xc1, 0xde,
	0x5b, 0x59, 0xfc, 0xc7, 0xfe, 0x2c, 0x4b, 0x5c, 0x5f, 0x5d, 0xe2, 0xc5, 0x43, 0x2a, 0x25, 0x7e,
	0xd0, 0xdf, 0x8d, 0x77, 0xe9, 0xef, 0x25, 0xc1, 0xca, 0x10, 0x03, 0x80, 0x2d, 0x69, 0x37, 0x06,
	0xfd, 0x9f, 0x8e, 0x42, 0xa6, 0xa7, 0xe9, 0xb8, 0xeb, 0x8b, 0xb8, 0x27, 0x26, 0x13, 0x7f, 0x4a,
	0x19, 0x8f, 0xe8, 0x58, 0xf5, 0xa8, 0x1c, 0xbf, 0x9a, 0x15, 0x3d, 0xd2, 0x2b, 0xcf, 0x1b, 0x6f,
	0x98, 0xa7, 0xe1, 0xf8, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xee, 0x72, 0xd8, 0x58, 0xa2, 0x08,
	0x00, 0x00,
}
