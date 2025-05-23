// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: game/joker/v1alpha1/joker.proto

package jokerv1alpha1

import (
	v1alpha1 "github.com/popberry-ai/popberry-proto/game/v1alpha1"
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

type GameRoundStatus int32

const (
	GameRoundStatus_GAME_ROUND_STATUS_UNSPECIFIED GameRoundStatus = 0
	GameRoundStatus_GAME_ROUND_STATUS_SUCCESS     GameRoundStatus = 1
	GameRoundStatus_GAME_ROUND_STATUS_CANCELLED   GameRoundStatus = 2
	GameRoundStatus_GAME_ROUND_STATUS_SETTLED     GameRoundStatus = 3
)

// Enum value maps for GameRoundStatus.
var (
	GameRoundStatus_name = map[int32]string{
		0: "GAME_ROUND_STATUS_UNSPECIFIED",
		1: "GAME_ROUND_STATUS_SUCCESS",
		2: "GAME_ROUND_STATUS_CANCELLED",
		3: "GAME_ROUND_STATUS_SETTLED",
	}
	GameRoundStatus_value = map[string]int32{
		"GAME_ROUND_STATUS_UNSPECIFIED": 0,
		"GAME_ROUND_STATUS_SUCCESS":     1,
		"GAME_ROUND_STATUS_CANCELLED":   2,
		"GAME_ROUND_STATUS_SETTLED":     3,
	}
)

func (x GameRoundStatus) Enum() *GameRoundStatus {
	p := new(GameRoundStatus)
	*p = x
	return p
}

func (x GameRoundStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameRoundStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_game_joker_v1alpha1_joker_proto_enumTypes[0].Descriptor()
}

func (GameRoundStatus) Type() protoreflect.EnumType {
	return &file_game_joker_v1alpha1_joker_proto_enumTypes[0]
}

func (x GameRoundStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameRoundStatus.Descriptor instead.
func (GameRoundStatus) EnumDescriptor() ([]byte, []int) {
	return file_game_joker_v1alpha1_joker_proto_rawDescGZIP(), []int{0}
}

type GameLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username    string  `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	GameCode    string  `protobuf:"bytes,2,opt,name=game_code,json=gameCode,proto3" json:"game_code,omitempty"`
	IsMobile    bool    `protobuf:"varint,3,opt,name=is_mobile,json=isMobile,proto3" json:"is_mobile,omitempty"`
	Token       *string `protobuf:"bytes,4,opt,name=token,proto3,oneof" json:"token,omitempty"`
	RedirectUrl *string `protobuf:"bytes,5,opt,name=redirect_url,json=redirectUrl,proto3,oneof" json:"redirect_url,omitempty"`
	Ip          *string `protobuf:"bytes,6,opt,name=ip,proto3,oneof" json:"ip,omitempty"`
	UserAgent   *string `protobuf:"bytes,7,opt,name=user_agent,json=userAgent,proto3,oneof" json:"user_agent,omitempty"`
}

func (x *GameLoginRequest) Reset() {
	*x = GameLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_joker_v1alpha1_joker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameLoginRequest) ProtoMessage() {}

func (x *GameLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_joker_v1alpha1_joker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameLoginRequest.ProtoReflect.Descriptor instead.
func (*GameLoginRequest) Descriptor() ([]byte, []int) {
	return file_game_joker_v1alpha1_joker_proto_rawDescGZIP(), []int{0}
}

func (x *GameLoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GameLoginRequest) GetGameCode() string {
	if x != nil {
		return x.GameCode
	}
	return ""
}

func (x *GameLoginRequest) GetIsMobile() bool {
	if x != nil {
		return x.IsMobile
	}
	return false
}

func (x *GameLoginRequest) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

func (x *GameLoginRequest) GetRedirectUrl() string {
	if x != nil && x.RedirectUrl != nil {
		return *x.RedirectUrl
	}
	return ""
}

func (x *GameLoginRequest) GetIp() string {
	if x != nil && x.Ip != nil {
		return *x.Ip
	}
	return ""
}

func (x *GameLoginRequest) GetUserAgent() string {
	if x != nil && x.UserAgent != nil {
		return *x.UserAgent
	}
	return ""
}

type GameLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess bool   `protobuf:"varint,1,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	Token     string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Url       string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GameLoginResponse) Reset() {
	*x = GameLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_joker_v1alpha1_joker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameLoginResponse) ProtoMessage() {}

func (x *GameLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_joker_v1alpha1_joker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameLoginResponse.ProtoReflect.Descriptor instead.
func (*GameLoginResponse) Descriptor() ([]byte, []int) {
	return file_game_joker_v1alpha1_joker_proto_rawDescGZIP(), []int{1}
}

func (x *GameLoginResponse) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *GameLoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GameLoginResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type GetGameListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize  int32   `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken string  `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	Query     *string `protobuf:"bytes,3,opt,name=query,proto3,oneof" json:"query,omitempty"`
}

func (x *GetGameListRequest) Reset() {
	*x = GetGameListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_joker_v1alpha1_joker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGameListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGameListRequest) ProtoMessage() {}

func (x *GetGameListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_joker_v1alpha1_joker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGameListRequest.ProtoReflect.Descriptor instead.
func (*GetGameListRequest) Descriptor() ([]byte, []int) {
	return file_game_joker_v1alpha1_joker_proto_rawDescGZIP(), []int{2}
}

func (x *GetGameListRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetGameListRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *GetGameListRequest) GetQuery() string {
	if x != nil && x.Query != nil {
		return *x.Query
	}
	return ""
}

type GetGameListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Games         []*v1alpha1.Game `protobuf:"bytes,1,rep,name=games,proto3" json:"games,omitempty"`
	NextPageToken string           `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	TotalSize     int32            `protobuf:"varint,3,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
}

func (x *GetGameListResponse) Reset() {
	*x = GetGameListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_joker_v1alpha1_joker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGameListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGameListResponse) ProtoMessage() {}

func (x *GetGameListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_joker_v1alpha1_joker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGameListResponse.ProtoReflect.Descriptor instead.
func (*GetGameListResponse) Descriptor() ([]byte, []int) {
	return file_game_joker_v1alpha1_joker_proto_rawDescGZIP(), []int{3}
}

func (x *GetGameListResponse) GetGames() []*v1alpha1.Game {
	if x != nil {
		return x.Games
	}
	return nil
}

func (x *GetGameListResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *GetGameListResponse) GetTotalSize() int32 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

type GetGameRoundStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoundId string `protobuf:"bytes,1,opt,name=round_id,json=roundId,proto3" json:"round_id,omitempty"`
}

func (x *GetGameRoundStatusRequest) Reset() {
	*x = GetGameRoundStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_joker_v1alpha1_joker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGameRoundStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGameRoundStatusRequest) ProtoMessage() {}

func (x *GetGameRoundStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_joker_v1alpha1_joker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGameRoundStatusRequest.ProtoReflect.Descriptor instead.
func (*GetGameRoundStatusRequest) Descriptor() ([]byte, []int) {
	return file_game_joker_v1alpha1_joker_proto_rawDescGZIP(), []int{4}
}

func (x *GetGameRoundStatusRequest) GetRoundId() string {
	if x != nil {
		return x.RoundId
	}
	return ""
}

type GetGameRoundStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status GameRoundStatus `protobuf:"varint,1,opt,name=status,proto3,enum=game.joker.v1alpha1.GameRoundStatus" json:"status,omitempty"`
	Amount float64         `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *GetGameRoundStatusResponse) Reset() {
	*x = GetGameRoundStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_joker_v1alpha1_joker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGameRoundStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGameRoundStatusResponse) ProtoMessage() {}

func (x *GetGameRoundStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_joker_v1alpha1_joker_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGameRoundStatusResponse.ProtoReflect.Descriptor instead.
func (*GetGameRoundStatusResponse) Descriptor() ([]byte, []int) {
	return file_game_joker_v1alpha1_joker_proto_rawDescGZIP(), []int{5}
}

func (x *GetGameRoundStatusResponse) GetStatus() GameRoundStatus {
	if x != nil {
		return x.Status
	}
	return GameRoundStatus_GAME_ROUND_STATUS_UNSPECIFIED
}

func (x *GetGameRoundStatusResponse) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type GetGameDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoundId string `protobuf:"bytes,1,opt,name=round_id,json=roundId,proto3" json:"round_id,omitempty"`
}

func (x *GetGameDetailRequest) Reset() {
	*x = GetGameDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_joker_v1alpha1_joker_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGameDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGameDetailRequest) ProtoMessage() {}

func (x *GetGameDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_joker_v1alpha1_joker_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGameDetailRequest.ProtoReflect.Descriptor instead.
func (*GetGameDetailRequest) Descriptor() ([]byte, []int) {
	return file_game_joker_v1alpha1_joker_proto_rawDescGZIP(), []int{6}
}

func (x *GetGameDetailRequest) GetRoundId() string {
	if x != nil {
		return x.RoundId
	}
	return ""
}

type GetGameDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess bool    `protobuf:"varint,1,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	Url       *string `protobuf:"bytes,2,opt,name=url,proto3,oneof" json:"url,omitempty"`
}

func (x *GetGameDetailResponse) Reset() {
	*x = GetGameDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_joker_v1alpha1_joker_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGameDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGameDetailResponse) ProtoMessage() {}

func (x *GetGameDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_joker_v1alpha1_joker_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGameDetailResponse.ProtoReflect.Descriptor instead.
func (*GetGameDetailResponse) Descriptor() ([]byte, []int) {
	return file_game_joker_v1alpha1_joker_proto_rawDescGZIP(), []int{7}
}

func (x *GetGameDetailResponse) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *GetGameDetailResponse) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

type GameSignoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *GameSignoutRequest) Reset() {
	*x = GameSignoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_joker_v1alpha1_joker_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameSignoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameSignoutRequest) ProtoMessage() {}

func (x *GameSignoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_joker_v1alpha1_joker_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameSignoutRequest.ProtoReflect.Descriptor instead.
func (*GameSignoutRequest) Descriptor() ([]byte, []int) {
	return file_game_joker_v1alpha1_joker_proto_rawDescGZIP(), []int{8}
}

func (x *GameSignoutRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type GameSignoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess bool `protobuf:"varint,1,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
}

func (x *GameSignoutResponse) Reset() {
	*x = GameSignoutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_joker_v1alpha1_joker_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameSignoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameSignoutResponse) ProtoMessage() {}

func (x *GameSignoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_joker_v1alpha1_joker_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameSignoutResponse.ProtoReflect.Descriptor instead.
func (*GameSignoutResponse) Descriptor() ([]byte, []int) {
	return file_game_joker_v1alpha1_joker_proto_rawDescGZIP(), []int{9}
}

func (x *GameSignoutResponse) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

var File_game_joker_v1alpha1_joker_proto protoreflect.FileDescriptor

var file_game_joker_v1alpha1_joker_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x6a, 0x6f, 0x6b, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6a, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x6a, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x18, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x95, 0x02, 0x0a, 0x10, 0x47, 0x61, 0x6d, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x69, 0x73, 0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x69, 0x73, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b,
	0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x13,
	0x0a, 0x02, 0x69, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x02, 0x69, 0x70,
	0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75,
	0x72, 0x6c, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x70, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x22, 0x5a, 0x0a, 0x11, 0x47, 0x61, 0x6d, 0x65,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x22, 0x75, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x88, 0x01,
	0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x87, 0x01, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x05, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x26,
	0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x36, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65,
	0x52, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x64, 0x22, 0x72, 0x0a,
	0x1a, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x2e, 0x6a, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x31, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x49, 0x64, 0x22, 0x55, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x15, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x75, 0x72, 0x6c, 0x22, 0x30, 0x0a, 0x12, 0x47,
	0x61, 0x6d, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x34, 0x0a,
	0x13, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x2a, 0x93, 0x01, 0x0a, 0x0f, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x6f, 0x75, 0x6e,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x1d, 0x47, 0x41, 0x4d, 0x45, 0x5f,
	0x52, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x47, 0x41,
	0x4d, 0x45, 0x5f, 0x52, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x47, 0x41, 0x4d,
	0x45, 0x5f, 0x52, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43,
	0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x1d, 0x0a, 0x19, 0x47, 0x41,
	0x4d, 0x45, 0x5f, 0x52, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x53, 0x45, 0x54, 0x54, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x32, 0x97, 0x04, 0x0a, 0x0c, 0x4a, 0x6f,
	0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x09, 0x47, 0x61,
	0x6d, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x25, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x6a,
	0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x61,
	0x6d, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x6a, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x47,
	0x61, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x6a,
	0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x47, 0x61, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x28, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x6a, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x77, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x2e, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x6a, 0x6f, 0x6b, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65,
	0x52, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x6a, 0x6f, 0x6b, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65,
	0x52, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x29, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x6a, 0x6f,
	0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x47, 0x61, 0x6d, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2a, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x6a, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x62, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x6f, 0x75, 0x74, 0x12, 0x27,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x6a, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x6f, 0x75, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x6a,
	0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x61,
	0x6d, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0xea, 0x01, 0x0a, 0x25, 0x61, 0x69, 0x2e, 0x70, 0x6f, 0x70, 0x62, 0x65,
	0x72, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x6a,
	0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0a, 0x4a,
	0x6f, 0x6b, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x47, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x70, 0x62, 0x65, 0x72, 0x72, 0x79,
	0x2d, 0x61, 0x69, 0x2f, 0x70, 0x6f, 0x70, 0x62, 0x65, 0x72, 0x72, 0x79, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x6a, 0x6f, 0x6b, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x6a, 0x6f, 0x6b, 0x65, 0x72, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x47, 0x4a, 0x58, 0xaa, 0x02, 0x13, 0x47, 0x61, 0x6d,
	0x65, 0x2e, 0x4a, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xca, 0x02, 0x13, 0x47, 0x61, 0x6d, 0x65, 0x5c, 0x4a, 0x6f, 0x6b, 0x65, 0x72, 0x5c, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x1f, 0x47, 0x61, 0x6d, 0x65, 0x5c, 0x4a, 0x6f,
	0x6b, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x47, 0x61, 0x6d, 0x65, 0x3a,
	0x3a, 0x4a, 0x6f, 0x6b, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_joker_v1alpha1_joker_proto_rawDescOnce sync.Once
	file_game_joker_v1alpha1_joker_proto_rawDescData = file_game_joker_v1alpha1_joker_proto_rawDesc
)

func file_game_joker_v1alpha1_joker_proto_rawDescGZIP() []byte {
	file_game_joker_v1alpha1_joker_proto_rawDescOnce.Do(func() {
		file_game_joker_v1alpha1_joker_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_joker_v1alpha1_joker_proto_rawDescData)
	})
	return file_game_joker_v1alpha1_joker_proto_rawDescData
}

var file_game_joker_v1alpha1_joker_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_game_joker_v1alpha1_joker_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_game_joker_v1alpha1_joker_proto_goTypes = []any{
	(GameRoundStatus)(0),               // 0: game.joker.v1alpha1.GameRoundStatus
	(*GameLoginRequest)(nil),           // 1: game.joker.v1alpha1.GameLoginRequest
	(*GameLoginResponse)(nil),          // 2: game.joker.v1alpha1.GameLoginResponse
	(*GetGameListRequest)(nil),         // 3: game.joker.v1alpha1.GetGameListRequest
	(*GetGameListResponse)(nil),        // 4: game.joker.v1alpha1.GetGameListResponse
	(*GetGameRoundStatusRequest)(nil),  // 5: game.joker.v1alpha1.GetGameRoundStatusRequest
	(*GetGameRoundStatusResponse)(nil), // 6: game.joker.v1alpha1.GetGameRoundStatusResponse
	(*GetGameDetailRequest)(nil),       // 7: game.joker.v1alpha1.GetGameDetailRequest
	(*GetGameDetailResponse)(nil),      // 8: game.joker.v1alpha1.GetGameDetailResponse
	(*GameSignoutRequest)(nil),         // 9: game.joker.v1alpha1.GameSignoutRequest
	(*GameSignoutResponse)(nil),        // 10: game.joker.v1alpha1.GameSignoutResponse
	(*v1alpha1.Game)(nil),              // 11: game.v1alpha1.Game
}
var file_game_joker_v1alpha1_joker_proto_depIdxs = []int32{
	11, // 0: game.joker.v1alpha1.GetGameListResponse.games:type_name -> game.v1alpha1.Game
	0,  // 1: game.joker.v1alpha1.GetGameRoundStatusResponse.status:type_name -> game.joker.v1alpha1.GameRoundStatus
	1,  // 2: game.joker.v1alpha1.JokerService.GameLogin:input_type -> game.joker.v1alpha1.GameLoginRequest
	3,  // 3: game.joker.v1alpha1.JokerService.GetGameList:input_type -> game.joker.v1alpha1.GetGameListRequest
	5,  // 4: game.joker.v1alpha1.JokerService.GetGameRoundStatus:input_type -> game.joker.v1alpha1.GetGameRoundStatusRequest
	7,  // 5: game.joker.v1alpha1.JokerService.GetGameDetail:input_type -> game.joker.v1alpha1.GetGameDetailRequest
	9,  // 6: game.joker.v1alpha1.JokerService.GameSignout:input_type -> game.joker.v1alpha1.GameSignoutRequest
	2,  // 7: game.joker.v1alpha1.JokerService.GameLogin:output_type -> game.joker.v1alpha1.GameLoginResponse
	4,  // 8: game.joker.v1alpha1.JokerService.GetGameList:output_type -> game.joker.v1alpha1.GetGameListResponse
	6,  // 9: game.joker.v1alpha1.JokerService.GetGameRoundStatus:output_type -> game.joker.v1alpha1.GetGameRoundStatusResponse
	8,  // 10: game.joker.v1alpha1.JokerService.GetGameDetail:output_type -> game.joker.v1alpha1.GetGameDetailResponse
	10, // 11: game.joker.v1alpha1.JokerService.GameSignout:output_type -> game.joker.v1alpha1.GameSignoutResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_game_joker_v1alpha1_joker_proto_init() }
func file_game_joker_v1alpha1_joker_proto_init() {
	if File_game_joker_v1alpha1_joker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_joker_v1alpha1_joker_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GameLoginRequest); i {
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
		file_game_joker_v1alpha1_joker_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GameLoginResponse); i {
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
		file_game_joker_v1alpha1_joker_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetGameListRequest); i {
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
		file_game_joker_v1alpha1_joker_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetGameListResponse); i {
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
		file_game_joker_v1alpha1_joker_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetGameRoundStatusRequest); i {
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
		file_game_joker_v1alpha1_joker_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetGameRoundStatusResponse); i {
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
		file_game_joker_v1alpha1_joker_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetGameDetailRequest); i {
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
		file_game_joker_v1alpha1_joker_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetGameDetailResponse); i {
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
		file_game_joker_v1alpha1_joker_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GameSignoutRequest); i {
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
		file_game_joker_v1alpha1_joker_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*GameSignoutResponse); i {
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
	file_game_joker_v1alpha1_joker_proto_msgTypes[0].OneofWrappers = []any{}
	file_game_joker_v1alpha1_joker_proto_msgTypes[2].OneofWrappers = []any{}
	file_game_joker_v1alpha1_joker_proto_msgTypes[7].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_game_joker_v1alpha1_joker_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_game_joker_v1alpha1_joker_proto_goTypes,
		DependencyIndexes: file_game_joker_v1alpha1_joker_proto_depIdxs,
		EnumInfos:         file_game_joker_v1alpha1_joker_proto_enumTypes,
		MessageInfos:      file_game_joker_v1alpha1_joker_proto_msgTypes,
	}.Build()
	File_game_joker_v1alpha1_joker_proto = out.File
	file_game_joker_v1alpha1_joker_proto_rawDesc = nil
	file_game_joker_v1alpha1_joker_proto_goTypes = nil
	file_game_joker_v1alpha1_joker_proto_depIdxs = nil
}
