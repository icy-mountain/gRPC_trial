// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: arithQuestioner/arith_questioner.proto

package __

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

type QuestionMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *QuestionMessage) Reset() {
	*x = QuestionMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arithQuestioner_arith_questioner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionMessage) ProtoMessage() {}

func (x *QuestionMessage) ProtoReflect() protoreflect.Message {
	mi := &file_arithQuestioner_arith_questioner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionMessage.ProtoReflect.Descriptor instead.
func (*QuestionMessage) Descriptor() ([]byte, []int) {
	return file_arithQuestioner_arith_questioner_proto_rawDescGZIP(), []int{0}
}

func (x *QuestionMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_arithQuestioner_arith_questioner_proto protoreflect.FileDescriptor

var file_arithQuestioner_arith_questioner_proto_rawDesc = []byte{
	0x0a, 0x26, 0x61, 0x72, 0x69, 0x74, 0x68, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x65,
	0x72, 0x2f, 0x61, 0x72, 0x69, 0x74, 0x68, 0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x72, 0x69, 0x74, 0x68, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x22, 0x2b, 0x0a, 0x0f, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x6b, 0x0a, 0x0f, 0x41, 0x72, 0x69, 0x74, 0x68, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x12, 0x58, 0x0a, 0x0c, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x74, 0x12, 0x20, 0x2e, 0x61, 0x72, 0x69, 0x74,
	0x68, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x20, 0x2e, 0x61, 0x72,
	0x69, 0x74, 0x68, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x2e, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x28,
	0x01, 0x30, 0x01, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_arithQuestioner_arith_questioner_proto_rawDescOnce sync.Once
	file_arithQuestioner_arith_questioner_proto_rawDescData = file_arithQuestioner_arith_questioner_proto_rawDesc
)

func file_arithQuestioner_arith_questioner_proto_rawDescGZIP() []byte {
	file_arithQuestioner_arith_questioner_proto_rawDescOnce.Do(func() {
		file_arithQuestioner_arith_questioner_proto_rawDescData = protoimpl.X.CompressGZIP(file_arithQuestioner_arith_questioner_proto_rawDescData)
	})
	return file_arithQuestioner_arith_questioner_proto_rawDescData
}

var file_arithQuestioner_arith_questioner_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_arithQuestioner_arith_questioner_proto_goTypes = []interface{}{
	(*QuestionMessage)(nil), // 0: arithquestioner.QuestionMessage
}
var file_arithQuestioner_arith_questioner_proto_depIdxs = []int32{
	0, // 0: arithquestioner.ArithQuestioner.QuestionChat:input_type -> arithquestioner.QuestionMessage
	0, // 1: arithquestioner.ArithQuestioner.QuestionChat:output_type -> arithquestioner.QuestionMessage
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_arithQuestioner_arith_questioner_proto_init() }
func file_arithQuestioner_arith_questioner_proto_init() {
	if File_arithQuestioner_arith_questioner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_arithQuestioner_arith_questioner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestionMessage); i {
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
			RawDescriptor: file_arithQuestioner_arith_questioner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_arithQuestioner_arith_questioner_proto_goTypes,
		DependencyIndexes: file_arithQuestioner_arith_questioner_proto_depIdxs,
		MessageInfos:      file_arithQuestioner_arith_questioner_proto_msgTypes,
	}.Build()
	File_arithQuestioner_arith_questioner_proto = out.File
	file_arithQuestioner_arith_questioner_proto_rawDesc = nil
	file_arithQuestioner_arith_questioner_proto_goTypes = nil
	file_arithQuestioner_arith_questioner_proto_depIdxs = nil
}
