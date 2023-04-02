// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: rpc_get_exercises_by_initial.proto

package pb

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

type GetExercisesByInitialRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetExercisesByInitialRequest) Reset() {
	*x = GetExercisesByInitialRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_get_exercises_by_initial_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExercisesByInitialRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExercisesByInitialRequest) ProtoMessage() {}

func (x *GetExercisesByInitialRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_get_exercises_by_initial_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExercisesByInitialRequest.ProtoReflect.Descriptor instead.
func (*GetExercisesByInitialRequest) Descriptor() ([]byte, []int) {
	return file_rpc_get_exercises_by_initial_proto_rawDescGZIP(), []int{0}
}

type GetExercisesByInitialResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exercises map[string]*Exercise `protobuf:"bytes,1,rep,name=exercises,proto3" json:"exercises,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetExercisesByInitialResponse) Reset() {
	*x = GetExercisesByInitialResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_get_exercises_by_initial_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExercisesByInitialResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExercisesByInitialResponse) ProtoMessage() {}

func (x *GetExercisesByInitialResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_get_exercises_by_initial_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExercisesByInitialResponse.ProtoReflect.Descriptor instead.
func (*GetExercisesByInitialResponse) Descriptor() ([]byte, []int) {
	return file_rpc_get_exercises_by_initial_proto_rawDescGZIP(), []int{1}
}

func (x *GetExercisesByInitialResponse) GetExercises() map[string]*Exercise {
	if x != nil {
		return x.Exercises
	}
	return nil
}

var File_rpc_get_exercises_by_initial_proto protoreflect.FileDescriptor

var file_rpc_get_exercises_by_initial_proto_rawDesc = []byte{
	0x0a, 0x22, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69,
	0x73, 0x65, 0x73, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x45,
	0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x42, 0x79, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xbb, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x74,
	0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x42, 0x79, 0x49, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x09, 0x65, 0x78,
	0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x42,
	0x79, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x09, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x1a, 0x4a, 0x0a, 0x0e, 0x45, 0x78,
	0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x22,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x70, 0x62, 0x2e, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x61, 0x64, 0x7a, 0x65, 0x79, 0x6e, 0x61, 0x6c, 0x2f,
	0x67, 0x79, 0x6d, 0x62, 0x72, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_get_exercises_by_initial_proto_rawDescOnce sync.Once
	file_rpc_get_exercises_by_initial_proto_rawDescData = file_rpc_get_exercises_by_initial_proto_rawDesc
)

func file_rpc_get_exercises_by_initial_proto_rawDescGZIP() []byte {
	file_rpc_get_exercises_by_initial_proto_rawDescOnce.Do(func() {
		file_rpc_get_exercises_by_initial_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_get_exercises_by_initial_proto_rawDescData)
	})
	return file_rpc_get_exercises_by_initial_proto_rawDescData
}

var file_rpc_get_exercises_by_initial_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rpc_get_exercises_by_initial_proto_goTypes = []interface{}{
	(*GetExercisesByInitialRequest)(nil),  // 0: pb.GetExercisesByInitialRequest
	(*GetExercisesByInitialResponse)(nil), // 1: pb.GetExercisesByInitialResponse
	nil,                                   // 2: pb.GetExercisesByInitialResponse.ExercisesEntry
	(*Exercise)(nil),                      // 3: pb.Exercise
}
var file_rpc_get_exercises_by_initial_proto_depIdxs = []int32{
	2, // 0: pb.GetExercisesByInitialResponse.exercises:type_name -> pb.GetExercisesByInitialResponse.ExercisesEntry
	3, // 1: pb.GetExercisesByInitialResponse.ExercisesEntry.value:type_name -> pb.Exercise
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpc_get_exercises_by_initial_proto_init() }
func file_rpc_get_exercises_by_initial_proto_init() {
	if File_rpc_get_exercises_by_initial_proto != nil {
		return
	}
	file_exercise_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_get_exercises_by_initial_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExercisesByInitialRequest); i {
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
		file_rpc_get_exercises_by_initial_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExercisesByInitialResponse); i {
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
			RawDescriptor: file_rpc_get_exercises_by_initial_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_get_exercises_by_initial_proto_goTypes,
		DependencyIndexes: file_rpc_get_exercises_by_initial_proto_depIdxs,
		MessageInfos:      file_rpc_get_exercises_by_initial_proto_msgTypes,
	}.Build()
	File_rpc_get_exercises_by_initial_proto = out.File
	file_rpc_get_exercises_by_initial_proto_rawDesc = nil
	file_rpc_get_exercises_by_initial_proto_goTypes = nil
	file_rpc_get_exercises_by_initial_proto_depIdxs = nil
}
