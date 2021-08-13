// DO NOT EDIT!
// This code is auto generated by go-xlsx-exporter
// VERSION 1.0
// go-protobuf v1.27.1

package go_xlsx_exporter

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

// Defined in table: 
type Language struct { 
    state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

    ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`  
    Text string `protobuf:"bytes,2,opt,name=Text,proto3" json:"Text,omitempty"`  
}

func (x *Language) Reset() {
	*x = Language{}
	if protoimpl.UnsafeEnabled {
		mi := &file___inner_language___proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Language) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Language) ProtoMessage() {}

func (x *Language) ProtoReflect() protoreflect.Message {
	mi := &file___inner_language___proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Language.ProtoReflect.Descriptor instead.
func (*Language) Descriptor() ([]byte, []int) {
	return file___inner_language___proto_rawDescGZIP(), []int{0}
}

func (x *Language) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Language) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}


// Defined in table: 
type Language_ARRAY struct { 
    state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

    Items []*Language `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`  
}

func (x *Language_ARRAY) Reset() {
	*x = Language_ARRAY{}
	if protoimpl.UnsafeEnabled {
		mi := &file___inner_language___proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Language_ARRAY) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Language_ARRAY) ProtoMessage() {}

func (x *Language_ARRAY) ProtoReflect() protoreflect.Message {
	mi := &file___inner_language___proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Language_ARRAY.ProtoReflect.Descriptor instead.
func (*Language_ARRAY) Descriptor() ([]byte, []int) {
	return file___inner_language___proto_rawDescGZIP(), []int{1}
}
    
func (x *Language_ARRAY) GetItems() []*Language {
	if x != nil {
		return x.Items
	}
	return nil
}


var File___inner_language___proto protoreflect.FileDescriptor

var file___inner_language___proto_rawDesc = []byte{
0x0a,0x18,0x5f,0x5f,0x69,0x6e,0x6e,0x65,0x72,0x5f,0x6c,0x61,0x6e,0x67,0x75,0x61,
0x67,0x65,0x5f,0x5f,0x2e,0x70,0x72,0x6f,0x74,0x6f,0x12,0x10,0x67,0x6f,0x5f,0x78,
0x6c,0x73,0x78,0x5f,0x65,0x78,0x70,0x6f,0x72,0x74,0x65,0x72,0x22,0x2e,0x0a,0x08,
0x4c,0x61,0x6e,0x67,0x75,0x61,0x67,0x65,0x12,0x0e,0x0a,0x02,0x49,0x44,0x18,0x01,
0x20,0x01,0x28,0x09,0x52,0x02,0x49,0x44,0x12,0x12,0x0a,0x04,0x54,0x65,0x78,0x74,
0x18,0x02,0x20,0x01,0x28,0x09,0x52,0x04,0x54,0x65,0x78,0x74,0x22,0x42,0x0a,0x0e,
0x4c,0x61,0x6e,0x67,0x75,0x61,0x67,0x65,0x5f,0x41,0x52,0x52,0x41,0x59,0x12,0x30,
0x0a,0x05,0x49,0x74,0x65,0x6d,0x73,0x18,0x01,0x20,0x03,0x28,0x0b,0x32,0x1a,0x2e,
0x67,0x6f,0x5f,0x78,0x6c,0x73,0x78,0x5f,0x65,0x78,0x70,0x6f,0x72,0x74,0x65,0x72,
0x2e,0x4c,0x61,0x6e,0x67,0x75,0x61,0x67,0x65,0x52,0x05,0x49,0x74,0x65,0x6d,0x73,
0x62,0x06,0x70,0x72,0x6f,0x74,0x6f,0x33,
}

var (
	file___inner_language___proto_rawDescOnce sync.Once
	file___inner_language___proto_rawDescData = file___inner_language___proto_rawDesc
)

func file___inner_language___proto_rawDescGZIP() []byte {
	file___inner_language___proto_rawDescOnce.Do(func() {
		file___inner_language___proto_rawDescData = protoimpl.X.CompressGZIP(file___inner_language___proto_rawDescData)
	})
	return file___inner_language___proto_rawDescData
}

var file___inner_language___proto_enumTypes = make([]protoimpl.EnumInfo, 0)
var file___inner_language___proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file___inner_language___proto_goTypes = []interface{}{
	 (*Language)(nil),				// 0: go_xlsx_exporter.Language
	 (*Language_ARRAY)(nil),				// 1: go_xlsx_exporter.Language_ARRAY
	
}

var file___inner_language___proto_depIdxs = []int32{
	0, // 0: go_xlsx_exporter.Language_ARRAY:type_name -> go_xlsx_exporter.Language
1, // [1:1] is the sub-list for method output_type
1, // [1:1] is the sub-list for method input_type
1, // [1:1] is the sub-list for extension type_name
1, // [1:1] is the sub-list for extension extendee
0, // [0:1] is the sub-list for field type_name
}

func init() { file___inner_language___proto_init() }
func file___inner_language___proto_init() {
	if File___inner_language___proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file___inner_language___proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Language); i {
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
		
		file___inner_language___proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Language_ARRAY); i {
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
			RawDescriptor: file___inner_language___proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file___inner_language___proto_goTypes,
		DependencyIndexes: file___inner_language___proto_depIdxs,
		EnumInfos:         file___inner_language___proto_enumTypes,
		MessageInfos:      file___inner_language___proto_msgTypes,
	}.Build()
	File___inner_language___proto = out.File
	file___inner_language___proto_rawDesc = nil
	file___inner_language___proto_goTypes = nil
	file___inner_language___proto_depIdxs = nil
}