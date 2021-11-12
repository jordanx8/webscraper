// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: proto/playerstuff.proto

package webscraper

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

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rank     int32  `protobuf:"varint,1,opt,name=rank,proto3" json:"rank,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	School   string `protobuf:"bytes,3,opt,name=school,proto3" json:"school,omitempty"`
	Position string `protobuf:"bytes,4,opt,name=position,proto3" json:"position,omitempty"`
	Nextgame string `protobuf:"bytes,5,opt,name=nextgame,proto3" json:"nextgame,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playerstuff_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playerstuff_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_proto_playerstuff_proto_rawDescGZIP(), []int{0}
}

func (x *Player) GetRank() int32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *Player) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Player) GetSchool() string {
	if x != nil {
		return x.School
	}
	return ""
}

func (x *Player) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *Player) GetNextgame() string {
	if x != nil {
		return x.Nextgame
	}
	return ""
}

type PlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rank     int32  `protobuf:"varint,1,opt,name=rank,proto3" json:"rank,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	School   string `protobuf:"bytes,3,opt,name=school,proto3" json:"school,omitempty"`
	Position string `protobuf:"bytes,4,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *PlayerRequest) Reset() {
	*x = PlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playerstuff_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerRequest) ProtoMessage() {}

func (x *PlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playerstuff_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerRequest.ProtoReflect.Descriptor instead.
func (*PlayerRequest) Descriptor() ([]byte, []int) {
	return file_proto_playerstuff_proto_rawDescGZIP(), []int{1}
}

func (x *PlayerRequest) GetRank() int32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *PlayerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PlayerRequest) GetSchool() string {
	if x != nil {
		return x.School
	}
	return ""
}

func (x *PlayerRequest) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

type PlayerArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Players []*Player `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *PlayerArray) Reset() {
	*x = PlayerArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playerstuff_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerArray) ProtoMessage() {}

func (x *PlayerArray) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playerstuff_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerArray.ProtoReflect.Descriptor instead.
func (*PlayerArray) Descriptor() ([]byte, []int) {
	return file_proto_playerstuff_proto_rawDescGZIP(), []int{2}
}

func (x *PlayerArray) GetPlayers() []*Player {
	if x != nil {
		return x.Players
	}
	return nil
}

var File_proto_playerstuff_proto protoreflect.FileDescriptor

var file_proto_playerstuff_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x74,
	0x75, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x77, 0x65, 0x62, 0x73, 0x63,
	0x72, 0x61, 0x70, 0x65, 0x72, 0x22, 0x80, 0x01, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x72, 0x61, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x6f,
	0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x6e, 0x65, 0x78, 0x74, 0x67, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x65, 0x78, 0x74, 0x67, 0x61, 0x6d, 0x65, 0x22, 0x6b, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3b, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41,
	0x72, 0x72, 0x61, 0x79, 0x12, 0x2c, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x63, 0x72, 0x61, 0x70,
	0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x32, 0x57, 0x0a, 0x11, 0x77, 0x65, 0x62, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x19, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x63, 0x72, 0x61, 0x70,
	0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x41, 0x72, 0x72, 0x61, 0x79, 0x22, 0x00, 0x42, 0x20, 0x5a, 0x1e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x6f, 0x72, 0x64, 0x61, 0x6e,
	0x78, 0x38, 0x2f, 0x77, 0x65, 0x62, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_playerstuff_proto_rawDescOnce sync.Once
	file_proto_playerstuff_proto_rawDescData = file_proto_playerstuff_proto_rawDesc
)

func file_proto_playerstuff_proto_rawDescGZIP() []byte {
	file_proto_playerstuff_proto_rawDescOnce.Do(func() {
		file_proto_playerstuff_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_playerstuff_proto_rawDescData)
	})
	return file_proto_playerstuff_proto_rawDescData
}

var file_proto_playerstuff_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_playerstuff_proto_goTypes = []interface{}{
	(*Player)(nil),        // 0: webscraper.Player
	(*PlayerRequest)(nil), // 1: webscraper.PlayerRequest
	(*PlayerArray)(nil),   // 2: webscraper.PlayerArray
}
var file_proto_playerstuff_proto_depIdxs = []int32{
	0, // 0: webscraper.PlayerArray.players:type_name -> webscraper.Player
	1, // 1: webscraper.webscraperService.GetPlayers:input_type -> webscraper.PlayerRequest
	2, // 2: webscraper.webscraperService.GetPlayers:output_type -> webscraper.PlayerArray
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_playerstuff_proto_init() }
func file_proto_playerstuff_proto_init() {
	if File_proto_playerstuff_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_playerstuff_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
		file_proto_playerstuff_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerRequest); i {
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
		file_proto_playerstuff_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerArray); i {
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
			RawDescriptor: file_proto_playerstuff_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_playerstuff_proto_goTypes,
		DependencyIndexes: file_proto_playerstuff_proto_depIdxs,
		MessageInfos:      file_proto_playerstuff_proto_msgTypes,
	}.Build()
	File_proto_playerstuff_proto = out.File
	file_proto_playerstuff_proto_rawDesc = nil
	file_proto_playerstuff_proto_goTypes = nil
	file_proto_playerstuff_proto_depIdxs = nil
}
