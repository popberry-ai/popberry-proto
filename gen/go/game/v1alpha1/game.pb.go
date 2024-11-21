// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: game/v1alpha1/game.proto

package gamev1alpha1

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

// Represents the status of a game.
type GameStatus int32

const (
	GameStatus_GAME_STATUS_UNSPECIFIED GameStatus = 0 // Default value. This value is unused.
	GameStatus_GAME_STATUS_ACTIVE      GameStatus = 1 // The game is currently active and available for play.
	GameStatus_GAME_STATUS_MAINTENANCE GameStatus = 2 // The game is under maintenance and temporarily unavailable.
	GameStatus_GAME_STATUS_CLOSED      GameStatus = 3 // The game is permanently closed and unavailable.
)

// Enum value maps for GameStatus.
var (
	GameStatus_name = map[int32]string{
		0: "GAME_STATUS_UNSPECIFIED",
		1: "GAME_STATUS_ACTIVE",
		2: "GAME_STATUS_MAINTENANCE",
		3: "GAME_STATUS_CLOSED",
	}
	GameStatus_value = map[string]int32{
		"GAME_STATUS_UNSPECIFIED": 0,
		"GAME_STATUS_ACTIVE":      1,
		"GAME_STATUS_MAINTENANCE": 2,
		"GAME_STATUS_CLOSED":      3,
	}
)

func (x GameStatus) Enum() *GameStatus {
	p := new(GameStatus)
	*p = x
	return p
}

func (x GameStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_game_v1alpha1_game_proto_enumTypes[0].Descriptor()
}

func (GameStatus) Type() protoreflect.EnumType {
	return &file_game_v1alpha1_game_proto_enumTypes[0]
}

func (x GameStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameStatus.Descriptor instead.
func (GameStatus) EnumDescriptor() ([]byte, []int) {
	return file_game_v1alpha1_game_proto_rawDescGZIP(), []int{0}
}

// Represents the category of a game.
type GameCategory int32

const (
	GameCategory_GAME_CATEGORY_UNSPECIFIED GameCategory = 0 // Default value. This value is unused.
	GameCategory_GAME_CATEGORY_SLOT        GameCategory = 1 // Slot games.
	GameCategory_GAME_CATEGORY_LIVE        GameCategory = 2 // Live games.
	GameCategory_GAME_CATEGORY_SPORT       GameCategory = 3 // Sports games.
	GameCategory_GAME_CATEGORY_FISH        GameCategory = 4 // Fish games.
	GameCategory_GAME_CATEGORY_LOTTERY     GameCategory = 5 // Lottery games.
	GameCategory_GAME_CATEGORY_POKER       GameCategory = 6 // Poker games.
	GameCategory_GAME_CATEGORY_OTHER       GameCategory = 7 // Other types of games.
)

// Enum value maps for GameCategory.
var (
	GameCategory_name = map[int32]string{
		0: "GAME_CATEGORY_UNSPECIFIED",
		1: "GAME_CATEGORY_SLOT",
		2: "GAME_CATEGORY_LIVE",
		3: "GAME_CATEGORY_SPORT",
		4: "GAME_CATEGORY_FISH",
		5: "GAME_CATEGORY_LOTTERY",
		6: "GAME_CATEGORY_POKER",
		7: "GAME_CATEGORY_OTHER",
	}
	GameCategory_value = map[string]int32{
		"GAME_CATEGORY_UNSPECIFIED": 0,
		"GAME_CATEGORY_SLOT":        1,
		"GAME_CATEGORY_LIVE":        2,
		"GAME_CATEGORY_SPORT":       3,
		"GAME_CATEGORY_FISH":        4,
		"GAME_CATEGORY_LOTTERY":     5,
		"GAME_CATEGORY_POKER":       6,
		"GAME_CATEGORY_OTHER":       7,
	}
)

func (x GameCategory) Enum() *GameCategory {
	p := new(GameCategory)
	*p = x
	return p
}

func (x GameCategory) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameCategory) Descriptor() protoreflect.EnumDescriptor {
	return file_game_v1alpha1_game_proto_enumTypes[1].Descriptor()
}

func (GameCategory) Type() protoreflect.EnumType {
	return &file_game_v1alpha1_game_proto_enumTypes[1]
}

func (x GameCategory) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameCategory.Descriptor instead.
func (GameCategory) EnumDescriptor() ([]byte, []int) {
	return file_game_v1alpha1_game_proto_rawDescGZIP(), []int{1}
}

// Represents the provider of a game.
type GameProvider int32

const (
	GameProvider_GAME_PROVIDER_UNSPECIFIED GameProvider = 0 // Default value. This value is unused.
	GameProvider_GAME_PROVIDER_JOKER       GameProvider = 1 // Joker game provider.
)

// Enum value maps for GameProvider.
var (
	GameProvider_name = map[int32]string{
		0: "GAME_PROVIDER_UNSPECIFIED",
		1: "GAME_PROVIDER_JOKER",
	}
	GameProvider_value = map[string]int32{
		"GAME_PROVIDER_UNSPECIFIED": 0,
		"GAME_PROVIDER_JOKER":       1,
	}
)

func (x GameProvider) Enum() *GameProvider {
	p := new(GameProvider)
	*p = x
	return p
}

func (x GameProvider) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameProvider) Descriptor() protoreflect.EnumDescriptor {
	return file_game_v1alpha1_game_proto_enumTypes[2].Descriptor()
}

func (GameProvider) Type() protoreflect.EnumType {
	return &file_game_v1alpha1_game_proto_enumTypes[2]
}

func (x GameProvider) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameProvider.Descriptor instead.
func (GameProvider) EnumDescriptor() ([]byte, []int) {
	return file_game_v1alpha1_game_proto_rawDescGZIP(), []int{2}
}

// Represents a game with various attributes.
type Game struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                              // Unique identifier for the game.
	Name     string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                          // Name of the game.
	Provider GameProvider `protobuf:"varint,3,opt,name=provider,proto3,enum=game.v1alpha1.GameProvider" json:"provider,omitempty"` // Provider of the game.
	ImageUrl string       `protobuf:"bytes,4,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`                  // URL to the game's image.
	Category GameCategory `protobuf:"varint,5,opt,name=category,proto3,enum=game.v1alpha1.GameCategory" json:"category,omitempty"` // Category of the game.
	Tags     []string     `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`                                          // Tags associated with the game.
	Status   GameStatus   `protobuf:"varint,7,opt,name=status,proto3,enum=game.v1alpha1.GameStatus" json:"status,omitempty"`       // Current status of the game.
	WinRate  *int32       `protobuf:"varint,8,opt,name=win_rate,json=winRate,proto3,oneof" json:"win_rate,omitempty"`              // Optional win rate of the game.
}

func (x *Game) Reset() {
	*x = Game{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_v1alpha1_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Game) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Game) ProtoMessage() {}

func (x *Game) ProtoReflect() protoreflect.Message {
	mi := &file_game_v1alpha1_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Game.ProtoReflect.Descriptor instead.
func (*Game) Descriptor() ([]byte, []int) {
	return file_game_v1alpha1_game_proto_rawDescGZIP(), []int{0}
}

func (x *Game) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Game) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Game) GetProvider() GameProvider {
	if x != nil {
		return x.Provider
	}
	return GameProvider_GAME_PROVIDER_UNSPECIFIED
}

func (x *Game) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Game) GetCategory() GameCategory {
	if x != nil {
		return x.Category
	}
	return GameCategory_GAME_CATEGORY_UNSPECIFIED
}

func (x *Game) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Game) GetStatus() GameStatus {
	if x != nil {
		return x.Status
	}
	return GameStatus_GAME_STATUS_UNSPECIFIED
}

func (x *Game) GetWinRate() int32 {
	if x != nil && x.WinRate != nil {
		return *x.WinRate
	}
	return 0
}

var File_game_v1alpha1_game_proto protoreflect.FileDescriptor

var file_game_v1alpha1_game_proto_rawDesc = []byte{
	0x0a, 0x18, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x67, 0x61, 0x6d, 0x65,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0xad, 0x02, 0x0a, 0x04, 0x47, 0x61,
	0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x37, 0x0a, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47,
	0x61, 0x6d, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x67, 0x61, 0x6d, 0x65,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x08,
	0x77, 0x69, 0x6e, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00,
	0x52, 0x07, 0x77, 0x69, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x77, 0x69, 0x6e, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x2a, 0x76, 0x0a, 0x0a, 0x47, 0x61, 0x6d,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x17, 0x47, 0x41, 0x4d, 0x45, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17,
	0x47, 0x41, 0x4d, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4d, 0x41, 0x49, 0x4e,
	0x54, 0x45, 0x4e, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x47, 0x41, 0x4d,
	0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10,
	0x03, 0x2a, 0xdb, 0x01, 0x0a, 0x0c, 0x47, 0x61, 0x6d, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x1d, 0x0a, 0x19, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47,
	0x4f, 0x52, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x16, 0x0a, 0x12, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f,
	0x52, 0x59, 0x5f, 0x53, 0x4c, 0x4f, 0x54, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x47, 0x41, 0x4d,
	0x45, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x4c, 0x49, 0x56, 0x45, 0x10,
	0x02, 0x12, 0x17, 0x0a, 0x13, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f,
	0x52, 0x59, 0x5f, 0x53, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x47, 0x41,
	0x4d, 0x45, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x46, 0x49, 0x53, 0x48,
	0x10, 0x04, 0x12, 0x19, 0x0a, 0x15, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47,
	0x4f, 0x52, 0x59, 0x5f, 0x4c, 0x4f, 0x54, 0x54, 0x45, 0x52, 0x59, 0x10, 0x05, 0x12, 0x17, 0x0a,
	0x13, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x50,
	0x4f, 0x4b, 0x45, 0x52, 0x10, 0x06, 0x12, 0x17, 0x0a, 0x13, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x43,
	0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x07, 0x2a,
	0x46, 0x0a, 0x0c, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12,
	0x1d, 0x0a, 0x19, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17,
	0x0a, 0x13, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f,
	0x4a, 0x4f, 0x4b, 0x45, 0x52, 0x10, 0x01, 0x42, 0xbc, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x09, 0x47,
	0x61, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x70, 0x62, 0x65, 0x72, 0x72, 0x79, 0x2d,
	0x61, 0x69, 0x2f, 0x70, 0x6f, 0x70, 0x62, 0x65, 0x72, 0x72, 0x79, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x67, 0x61, 0x6d, 0x65, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x47, 0x58, 0x58, 0xaa, 0x02, 0x0d, 0x47, 0x61, 0x6d, 0x65,
	0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x0d, 0x47, 0x61, 0x6d, 0x65,
	0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x19, 0x47, 0x61, 0x6d, 0x65,
	0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x47, 0x61, 0x6d, 0x65, 0x3a, 0x3a, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_v1alpha1_game_proto_rawDescOnce sync.Once
	file_game_v1alpha1_game_proto_rawDescData = file_game_v1alpha1_game_proto_rawDesc
)

func file_game_v1alpha1_game_proto_rawDescGZIP() []byte {
	file_game_v1alpha1_game_proto_rawDescOnce.Do(func() {
		file_game_v1alpha1_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_v1alpha1_game_proto_rawDescData)
	})
	return file_game_v1alpha1_game_proto_rawDescData
}

var file_game_v1alpha1_game_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_game_v1alpha1_game_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_game_v1alpha1_game_proto_goTypes = []any{
	(GameStatus)(0),   // 0: game.v1alpha1.GameStatus
	(GameCategory)(0), // 1: game.v1alpha1.GameCategory
	(GameProvider)(0), // 2: game.v1alpha1.GameProvider
	(*Game)(nil),      // 3: game.v1alpha1.Game
}
var file_game_v1alpha1_game_proto_depIdxs = []int32{
	2, // 0: game.v1alpha1.Game.provider:type_name -> game.v1alpha1.GameProvider
	1, // 1: game.v1alpha1.Game.category:type_name -> game.v1alpha1.GameCategory
	0, // 2: game.v1alpha1.Game.status:type_name -> game.v1alpha1.GameStatus
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_game_v1alpha1_game_proto_init() }
func file_game_v1alpha1_game_proto_init() {
	if File_game_v1alpha1_game_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_v1alpha1_game_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Game); i {
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
	file_game_v1alpha1_game_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_game_v1alpha1_game_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_game_v1alpha1_game_proto_goTypes,
		DependencyIndexes: file_game_v1alpha1_game_proto_depIdxs,
		EnumInfos:         file_game_v1alpha1_game_proto_enumTypes,
		MessageInfos:      file_game_v1alpha1_game_proto_msgTypes,
	}.Build()
	File_game_v1alpha1_game_proto = out.File
	file_game_v1alpha1_game_proto_rawDesc = nil
	file_game_v1alpha1_game_proto_goTypes = nil
	file_game_v1alpha1_game_proto_depIdxs = nil
}
