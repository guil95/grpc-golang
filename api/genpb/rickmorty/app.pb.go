// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: api/app.proto

package rickmorty

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CharactersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total      int32        `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Characters []*Character `protobuf:"bytes,2,rep,name=characters,proto3" json:"characters,omitempty"`
}

func (x *CharactersResponse) Reset() {
	*x = CharactersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CharactersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CharactersResponse) ProtoMessage() {}

func (x *CharactersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CharactersResponse.ProtoReflect.Descriptor instead.
func (*CharactersResponse) Descriptor() ([]byte, []int) {
	return file_api_app_proto_rawDescGZIP(), []int{0}
}

func (x *CharactersResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *CharactersResponse) GetCharacters() []*Character {
	if x != nil {
		return x.Characters
	}
	return nil
}

type Character struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Species string `protobuf:"bytes,2,opt,name=species,proto3" json:"species,omitempty"`
}

func (x *Character) Reset() {
	*x = Character{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Character) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Character) ProtoMessage() {}

func (x *Character) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Character.ProtoReflect.Descriptor instead.
func (*Character) Descriptor() ([]byte, []int) {
	return file_api_app_proto_rawDescGZIP(), []int{1}
}

func (x *Character) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Character) GetSpecies() string {
	if x != nil {
		return x.Species
	}
	return ""
}

type PageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *PageRequest) Reset() {
	*x = PageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageRequest) ProtoMessage() {}

func (x *PageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageRequest.ProtoReflect.Descriptor instead.
func (*PageRequest) Descriptor() ([]byte, []int) {
	return file_api_app_proto_rawDescGZIP(), []int{2}
}

func (x *PageRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

var File_api_app_proto protoreflect.FileDescriptor

var file_api_app_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x72, 0x69, 0x63, 0x6b, 0x6d, 0x6f, 0x72, 0x74, 0x79, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a, 0x12, 0x43, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x34, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x69, 0x63, 0x6b, 0x6d, 0x6f,
	0x72, 0x74, 0x79, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x0a, 0x63,
	0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x22, 0x39, 0x0a, 0x09, 0x43, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x65,
	0x63, 0x69, 0x65, 0x73, 0x22, 0x21, 0x0a, 0x0b, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x32, 0x8b, 0x02, 0x0a, 0x10, 0x52, 0x69, 0x63, 0x6b,
	0x4d, 0x6f, 0x72, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x11,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x72, 0x69, 0x63, 0x6b,
	0x6d, 0x6f, 0x72, 0x74, 0x79, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x6c, 0x6c, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x1d, 0x2e, 0x72, 0x69, 0x63, 0x6b, 0x6d, 0x6f, 0x72, 0x74, 0x79, 0x2e, 0x43, 0x68,
	0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x30, 0x01, 0x12, 0x51, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x50, 0x61, 0x67, 0x65,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x16, 0x2e, 0x72,
	0x69, 0x63, 0x6b, 0x6d, 0x6f, 0x72, 0x74, 0x79, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x72, 0x69, 0x63, 0x6b, 0x6d, 0x6f, 0x72, 0x74, 0x79,
	0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x28, 0x01, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x72, 0x69, 0x63, 0x6b, 0x6d,
	0x6f, 0x72, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_app_proto_rawDescOnce sync.Once
	file_api_app_proto_rawDescData = file_api_app_proto_rawDesc
)

func file_api_app_proto_rawDescGZIP() []byte {
	file_api_app_proto_rawDescOnce.Do(func() {
		file_api_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_app_proto_rawDescData)
	})
	return file_api_app_proto_rawDescData
}

var file_api_app_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_app_proto_goTypes = []interface{}{
	(*CharactersResponse)(nil), // 0: rickmorty.CharactersResponse
	(*Character)(nil),          // 1: rickmorty.Character
	(*PageRequest)(nil),        // 2: rickmorty.PageRequest
	(*emptypb.Empty)(nil),      // 3: google.protobuf.Empty
}
var file_api_app_proto_depIdxs = []int32{
	1, // 0: rickmorty.CharactersResponse.characters:type_name -> rickmorty.Character
	3, // 1: rickmorty.RickMortyService.ListAllCharacters:input_type -> google.protobuf.Empty
	3, // 2: rickmorty.RickMortyService.ListAllCharactersServerStream:input_type -> google.protobuf.Empty
	2, // 3: rickmorty.RickMortyService.ListByPageClientStream:input_type -> rickmorty.PageRequest
	0, // 4: rickmorty.RickMortyService.ListAllCharacters:output_type -> rickmorty.CharactersResponse
	0, // 5: rickmorty.RickMortyService.ListAllCharactersServerStream:output_type -> rickmorty.CharactersResponse
	0, // 6: rickmorty.RickMortyService.ListByPageClientStream:output_type -> rickmorty.CharactersResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_app_proto_init() }
func file_api_app_proto_init() {
	if File_api_app_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CharactersResponse); i {
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
		file_api_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Character); i {
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
		file_api_app_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageRequest); i {
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
			RawDescriptor: file_api_app_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_app_proto_goTypes,
		DependencyIndexes: file_api_app_proto_depIdxs,
		MessageInfos:      file_api_app_proto_msgTypes,
	}.Build()
	File_api_app_proto = out.File
	file_api_app_proto_rawDesc = nil
	file_api_app_proto_goTypes = nil
	file_api_app_proto_depIdxs = nil
}
