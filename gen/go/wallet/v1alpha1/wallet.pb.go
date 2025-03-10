// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: wallet/v1alpha1/wallet.proto

package walletv1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for updating wallet balance
type UpdateBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Requests      []*RequestResponse `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests,omitempty"`                                  // List of request and response pairs
	TenantId      string             `protobuf:"bytes,2,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`                  // tenant id
	BetAmount     float64            `protobuf:"fixed64,3,opt,name=bet_amount,json=betAmount,proto3" json:"bet_amount,omitempty"`             // Bet amount
	WinAmount     float64            `protobuf:"fixed64,4,opt,name=win_amount,json=winAmount,proto3" json:"win_amount,omitempty"`             // Win amount
	WinloseAmount float64            `protobuf:"fixed64,5,opt,name=winlose_amount,json=winloseAmount,proto3" json:"winlose_amount,omitempty"` // Win/Loss amount
	Status        string             `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`                                      // Status of the transaction
	UserId        string             `protobuf:"bytes,7,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                        // User ID (or username)
	RoundId       string             `protobuf:"bytes,8,opt,name=round_id,json=roundId,proto3" json:"round_id,omitempty"`                     // Round ID
	TxnId         string             `protobuf:"bytes,9,opt,name=txn_id,json=txnId,proto3" json:"txn_id,omitempty"`                           // Transaction ID
	ParentId      string             `protobuf:"bytes,10,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`                 // Parent transaction ID (optional)
	GameId        int32              `protobuf:"varint,11,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`                      // Game ID
	GameCode      string             `protobuf:"bytes,12,opt,name=game_code,json=gameCode,proto3" json:"game_code,omitempty"`                 // Game code
	ProviderId    int32              `protobuf:"varint,13,opt,name=provider_id,json=providerId,proto3" json:"provider_id,omitempty"`          // Provider ID
	ProviderCode  string             `protobuf:"bytes,14,opt,name=provider_code,json=providerCode,proto3" json:"provider_code,omitempty"`     // Provider code
	Turnover      float64            `protobuf:"fixed64,15,opt,name=turnover,proto3" json:"turnover,omitempty"`
}

func (x *UpdateBalanceRequest) Reset() {
	*x = UpdateBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_v1alpha1_wallet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBalanceRequest) ProtoMessage() {}

func (x *UpdateBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_v1alpha1_wallet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBalanceRequest.ProtoReflect.Descriptor instead.
func (*UpdateBalanceRequest) Descriptor() ([]byte, []int) {
	return file_wallet_v1alpha1_wallet_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateBalanceRequest) GetRequests() []*RequestResponse {
	if x != nil {
		return x.Requests
	}
	return nil
}

func (x *UpdateBalanceRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *UpdateBalanceRequest) GetBetAmount() float64 {
	if x != nil {
		return x.BetAmount
	}
	return 0
}

func (x *UpdateBalanceRequest) GetWinAmount() float64 {
	if x != nil {
		return x.WinAmount
	}
	return 0
}

func (x *UpdateBalanceRequest) GetWinloseAmount() float64 {
	if x != nil {
		return x.WinloseAmount
	}
	return 0
}

func (x *UpdateBalanceRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UpdateBalanceRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateBalanceRequest) GetRoundId() string {
	if x != nil {
		return x.RoundId
	}
	return ""
}

func (x *UpdateBalanceRequest) GetTxnId() string {
	if x != nil {
		return x.TxnId
	}
	return ""
}

func (x *UpdateBalanceRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *UpdateBalanceRequest) GetGameId() int32 {
	if x != nil {
		return x.GameId
	}
	return 0
}

func (x *UpdateBalanceRequest) GetGameCode() string {
	if x != nil {
		return x.GameCode
	}
	return ""
}

func (x *UpdateBalanceRequest) GetProviderId() int32 {
	if x != nil {
		return x.ProviderId
	}
	return 0
}

func (x *UpdateBalanceRequest) GetProviderCode() string {
	if x != nil {
		return x.ProviderCode
	}
	return ""
}

func (x *UpdateBalanceRequest) GetTurnover() float64 {
	if x != nil {
		return x.Turnover
	}
	return 0
}

// Response message for updating wallet balance
type UpdateBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success       bool    `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`                                 // Indicates if the operation was successful
	TransactionId string  `protobuf:"bytes,2,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"` // MongoDB ObjectId as a string
	Before        float64 `protobuf:"fixed64,3,opt,name=before,proto3" json:"before,omitempty"`                                  // Balance before the transaction
	Amount        float64 `protobuf:"fixed64,4,opt,name=amount,proto3" json:"amount,omitempty"`                                  // Transaction amount
	After         float64 `protobuf:"fixed64,5,opt,name=after,proto3" json:"after,omitempty"`                                    // Balance after the transaction
	Message       string  `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`                                  // Optional message for additional information
}

func (x *UpdateBalanceResponse) Reset() {
	*x = UpdateBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_v1alpha1_wallet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBalanceResponse) ProtoMessage() {}

func (x *UpdateBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_v1alpha1_wallet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBalanceResponse.ProtoReflect.Descriptor instead.
func (*UpdateBalanceResponse) Descriptor() ([]byte, []int) {
	return file_wallet_v1alpha1_wallet_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateBalanceResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UpdateBalanceResponse) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *UpdateBalanceResponse) GetBefore() float64 {
	if x != nil {
		return x.Before
	}
	return 0
}

func (x *UpdateBalanceResponse) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *UpdateBalanceResponse) GetAfter() float64 {
	if x != nil {
		return x.After
	}
	return 0
}

func (x *UpdateBalanceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Request message for identifying user existence with various methods
type IdentifyUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`          // User ID (optional)
	UserToken string `protobuf:"bytes,2,opt,name=user_token,json=userToken,proto3" json:"user_token,omitempty"` // User Token (optional)
	Username  string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`                    // Username (optional)
}

func (x *IdentifyUserRequest) Reset() {
	*x = IdentifyUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_v1alpha1_wallet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentifyUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentifyUserRequest) ProtoMessage() {}

func (x *IdentifyUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_v1alpha1_wallet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentifyUserRequest.ProtoReflect.Descriptor instead.
func (*IdentifyUserRequest) Descriptor() ([]byte, []int) {
	return file_wallet_v1alpha1_wallet_proto_rawDescGZIP(), []int{2}
}

func (x *IdentifyUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *IdentifyUserRequest) GetUserToken() string {
	if x != nil {
		return x.UserToken
	}
	return ""
}

func (x *IdentifyUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

// Response message for identifying user existence
type IdentifyUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists  bool   `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`  // Indicates if the user exists
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // Optional message for additional information
}

func (x *IdentifyUserResponse) Reset() {
	*x = IdentifyUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_v1alpha1_wallet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentifyUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentifyUserResponse) ProtoMessage() {}

func (x *IdentifyUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_v1alpha1_wallet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentifyUserResponse.ProtoReflect.Descriptor instead.
func (*IdentifyUserResponse) Descriptor() ([]byte, []int) {
	return file_wallet_v1alpha1_wallet_proto_rawDescGZIP(), []int{3}
}

func (x *IdentifyUserResponse) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

func (x *IdentifyUserResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Request message for updating the transaction response after another service has processed the request
type UpdateTransactionResponseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxnId    string `protobuf:"bytes,1,opt,name=txn_id,json=txnId,proto3" json:"txn_id,omitempty"` // Transaction ID to identify which record to update
	Response string `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`        // Updated response data (JSON or any other format)
}

func (x *UpdateTransactionResponseRequest) Reset() {
	*x = UpdateTransactionResponseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_v1alpha1_wallet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTransactionResponseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTransactionResponseRequest) ProtoMessage() {}

func (x *UpdateTransactionResponseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_v1alpha1_wallet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTransactionResponseRequest.ProtoReflect.Descriptor instead.
func (*UpdateTransactionResponseRequest) Descriptor() ([]byte, []int) {
	return file_wallet_v1alpha1_wallet_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateTransactionResponseRequest) GetTxnId() string {
	if x != nil {
		return x.TxnId
	}
	return ""
}

func (x *UpdateTransactionResponseRequest) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

// Response message for updating the transaction response
type UpdateTransactionResponseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // Indicates if the update was successful
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`  // Optional message for additional information
}

func (x *UpdateTransactionResponseResponse) Reset() {
	*x = UpdateTransactionResponseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_v1alpha1_wallet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTransactionResponseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTransactionResponseResponse) ProtoMessage() {}

func (x *UpdateTransactionResponseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_v1alpha1_wallet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTransactionResponseResponse.ProtoReflect.Descriptor instead.
func (*UpdateTransactionResponseResponse) Descriptor() ([]byte, []int) {
	return file_wallet_v1alpha1_wallet_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateTransactionResponseResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UpdateTransactionResponseResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Message for request and response pairs
type RequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request  string `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`   // Request data
	Response string `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"` // Response data (JSON or any other format)
}

func (x *RequestResponse) Reset() {
	*x = RequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_v1alpha1_wallet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestResponse) ProtoMessage() {}

func (x *RequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_v1alpha1_wallet_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestResponse.ProtoReflect.Descriptor instead.
func (*RequestResponse) Descriptor() ([]byte, []int) {
	return file_wallet_v1alpha1_wallet_proto_rawDescGZIP(), []int{6}
}

func (x *RequestResponse) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

func (x *RequestResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_wallet_v1alpha1_wallet_proto protoreflect.FileDescriptor

var file_wallet_v1alpha1_wallet_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22,
	0xee, 0x03, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x65, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x62, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x69, 0x6e, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x77, 0x69, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x25, 0x0a, 0x0e, 0x77, 0x69, 0x6e, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x77, 0x69, 0x6e, 0x6c, 0x6f,
	0x73, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x78, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x78, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x75, 0x72, 0x6e, 0x6f, 0x76, 0x65, 0x72,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x74, 0x75, 0x72, 0x6e, 0x6f, 0x76, 0x65, 0x72,
	0x22, 0xb8, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62,
	0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x62, 0x65, 0x66,
	0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x66, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x69, 0x0a, 0x13, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x14, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x55, 0x0a, 0x20, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x78, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x78, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x57, 0x0a, 0x21, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x47, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xd1, 0x02, 0x0a, 0x0d, 0x57, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x0d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x25, 0x2e, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x0c, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x19, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xd3, 0x01,
	0x0a, 0x21, 0x61, 0x69, 0x2e, 0x70, 0x6f, 0x70, 0x62, 0x65, 0x72, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x42, 0x0b, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x6f, 0x70, 0x62, 0x65, 0x72, 0x72, 0x79, 0x2d, 0x61, 0x69, 0x2f, 0x70, 0x6f, 0x70, 0x62, 0x65,
	0x72, 0x72, 0x79, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x57, 0x58, 0x58, 0xaa, 0x02,
	0x0f, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xca, 0x02, 0x0f, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xe2, 0x02, 0x1b, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x10, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wallet_v1alpha1_wallet_proto_rawDescOnce sync.Once
	file_wallet_v1alpha1_wallet_proto_rawDescData = file_wallet_v1alpha1_wallet_proto_rawDesc
)

func file_wallet_v1alpha1_wallet_proto_rawDescGZIP() []byte {
	file_wallet_v1alpha1_wallet_proto_rawDescOnce.Do(func() {
		file_wallet_v1alpha1_wallet_proto_rawDescData = protoimpl.X.CompressGZIP(file_wallet_v1alpha1_wallet_proto_rawDescData)
	})
	return file_wallet_v1alpha1_wallet_proto_rawDescData
}

var file_wallet_v1alpha1_wallet_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_wallet_v1alpha1_wallet_proto_goTypes = []any{
	(*UpdateBalanceRequest)(nil),              // 0: wallet.v1alpha1.UpdateBalanceRequest
	(*UpdateBalanceResponse)(nil),             // 1: wallet.v1alpha1.UpdateBalanceResponse
	(*IdentifyUserRequest)(nil),               // 2: wallet.v1alpha1.IdentifyUserRequest
	(*IdentifyUserResponse)(nil),              // 3: wallet.v1alpha1.IdentifyUserResponse
	(*UpdateTransactionResponseRequest)(nil),  // 4: wallet.v1alpha1.UpdateTransactionResponseRequest
	(*UpdateTransactionResponseResponse)(nil), // 5: wallet.v1alpha1.UpdateTransactionResponseResponse
	(*RequestResponse)(nil),                   // 6: wallet.v1alpha1.RequestResponse
}
var file_wallet_v1alpha1_wallet_proto_depIdxs = []int32{
	6, // 0: wallet.v1alpha1.UpdateBalanceRequest.requests:type_name -> wallet.v1alpha1.RequestResponse
	0, // 1: wallet.v1alpha1.WalletService.UpdateBalance:input_type -> wallet.v1alpha1.UpdateBalanceRequest
	2, // 2: wallet.v1alpha1.WalletService.IdentifyUser:input_type -> wallet.v1alpha1.IdentifyUserRequest
	4, // 3: wallet.v1alpha1.WalletService.UpdateTransactionResponse:input_type -> wallet.v1alpha1.UpdateTransactionResponseRequest
	1, // 4: wallet.v1alpha1.WalletService.UpdateBalance:output_type -> wallet.v1alpha1.UpdateBalanceResponse
	3, // 5: wallet.v1alpha1.WalletService.IdentifyUser:output_type -> wallet.v1alpha1.IdentifyUserResponse
	5, // 6: wallet.v1alpha1.WalletService.UpdateTransactionResponse:output_type -> wallet.v1alpha1.UpdateTransactionResponseResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_wallet_v1alpha1_wallet_proto_init() }
func file_wallet_v1alpha1_wallet_proto_init() {
	if File_wallet_v1alpha1_wallet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wallet_v1alpha1_wallet_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateBalanceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_v1alpha1_wallet_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateBalanceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_v1alpha1_wallet_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*IdentifyUserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_v1alpha1_wallet_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*IdentifyUserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_v1alpha1_wallet_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateTransactionResponseRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_v1alpha1_wallet_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateTransactionResponseResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_v1alpha1_wallet_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*RequestResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wallet_v1alpha1_wallet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wallet_v1alpha1_wallet_proto_goTypes,
		DependencyIndexes: file_wallet_v1alpha1_wallet_proto_depIdxs,
		MessageInfos:      file_wallet_v1alpha1_wallet_proto_msgTypes,
	}.Build()
	File_wallet_v1alpha1_wallet_proto = out.File
	file_wallet_v1alpha1_wallet_proto_rawDesc = nil
	file_wallet_v1alpha1_wallet_proto_goTypes = nil
	file_wallet_v1alpha1_wallet_proto_depIdxs = nil
}
