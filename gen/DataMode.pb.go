// DO NOT EDIT!
// This code is auto generated by GO-XLSX-PROTOBUF
// VERSION 1.0
// go-protobuf v1.27.1

package gen

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

// Defined in table: data/define.xlsx:define
type EDataType int32

const (
	EDataType_UNKNOWN EDataType = 0
	EDataType_JSON    EDataType = 1 //json数据
	EDataType_XML     EDataType = 2 //xml数据
	EDataType_CSHARP  EDataType = 5
	EDataType_GOLANG  EDataType = 6
)

var (
	EDataType_name = map[int32]string{
		0: "UNKNOWN",
		1: "JSON",
		2: "XML",
		5: "CSHARP",
		6: "GOLANG",
	}

	EDataType_value = map[string]int32{
		"UNKNOWN": 0,
		"JSON":    1,
		"XML":     2,
		"CSHARP":  5,
		"GOLANG":  6,
	}
)

func (x EDataType) Enum() *EDataType {
	p := new(EDataType)
	*p = x
	return p
}

func (x EDataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EDataType) Descriptor() protoreflect.EnumDescriptor {
	return file_DataMode_proto_enumTypes[0].Descriptor()
}

func (EDataType) Type() protoreflect.EnumType {
	return &file_DataMode_proto_enumTypes[0]
}

func (x EDataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EDataType.Descriptor instead.
func (EDataType) EnumDescriptor() ([]byte, []int) {
	return file_DataMode_proto_rawDescGZIP(), []int{0}
}

// Defined in table: data/define.xlsx:define
type SInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string    `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"` //姓名
	Age      int32     `protobuf:"varint,2,opt,name=Age,proto3" json:"Age,omitempty"`  //年龄
	Sex      bool      `protobuf:"varint,3,opt,name=Sex,proto3" json:"Sex,omitempty"`  //性别
	Ids      []int32   `protobuf:"bytes,4,rep,name=Ids,proto3" json:"Ids,omitempty"`
	Props    []string  `protobuf:"bytes,5,rep,name=Props,proto3" json:"Props,omitempty"`
	DataType EDataType `protobuf:"varint,6,opt,name=DataType,proto3,enum=gen.EDataType" json:"DataType,omitempty"`
}

func (x *SInfo) Reset() {
	*x = SInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DataMode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SInfo) ProtoMessage() {}

func (x *SInfo) ProtoReflect() protoreflect.Message {
	mi := &file_DataMode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SInfo.ProtoReflect.Descriptor instead.
func (*SInfo) Descriptor() ([]byte, []int) {
	return file_DataMode_proto_rawDescGZIP(), []int{0}
}

func (x *SInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SInfo) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *SInfo) GetSex() bool {
	if x != nil {
		return x.Sex
	}
	return false
}

func (x *SInfo) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *SInfo) GetProps() []string {
	if x != nil {
		return x.Props
	}
	return nil
}

func (x *SInfo) GetDataType() EDataType {
	if x != nil {
		return x.DataType
	}
	return EDataType_UNKNOWN
}

// Defined in table: data/define.xlsx:define
type SInfo_ARRAY struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*SInfo `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (x *SInfo_ARRAY) Reset() {
	*x = SInfo_ARRAY{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DataMode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SInfo_ARRAY) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SInfo_ARRAY) ProtoMessage() {}

func (x *SInfo_ARRAY) ProtoReflect() protoreflect.Message {
	mi := &file_DataMode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SInfo_ARRAY.ProtoReflect.Descriptor instead.
func (*SInfo_ARRAY) Descriptor() ([]byte, []int) {
	return file_DataMode_proto_rawDescGZIP(), []int{1}
}

func (x *SInfo_ARRAY) GetItems() []*SInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

// Defined in table: data/model.xlsx:user
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`         //ID
	Name   string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`      //名称
	Age    uint32 `protobuf:"varint,3,opt,name=Age,proto3" json:"Age,omitempty"`       //年龄
	Sex    bool   `protobuf:"varint,4,opt,name=Sex,proto3" json:"Sex,omitempty"`       //性别
	Head   string `protobuf:"bytes,5,opt,name=Head,proto3" json:"Head,omitempty"`      //头像
	PClass uint32 `protobuf:"varint,6,opt,name=PClass,proto3" json:"PClass,omitempty"` //班级
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DataMode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_DataMode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_DataMode_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetAge() uint32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *User) GetSex() bool {
	if x != nil {
		return x.Sex
	}
	return false
}

func (x *User) GetHead() string {
	if x != nil {
		return x.Head
	}
	return ""
}

func (x *User) GetPClass() uint32 {
	if x != nil {
		return x.PClass
	}
	return 0
}

// Defined in table: data/model.xlsx:user
type User_ARRAY struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*User `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (x *User_ARRAY) Reset() {
	*x = User_ARRAY{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DataMode_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User_ARRAY) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User_ARRAY) ProtoMessage() {}

func (x *User_ARRAY) ProtoReflect() protoreflect.Message {
	mi := &file_DataMode_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User_ARRAY.ProtoReflect.Descriptor instead.
func (*User_ARRAY) Descriptor() ([]byte, []int) {
	return file_DataMode_proto_rawDescGZIP(), []int{3}
}

func (x *User_ARRAY) GetItems() []*User {
	if x != nil {
		return x.Items
	}
	return nil
}

// Defined in table: data/model.xlsx:class
type PClass struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    uint32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`                        //ID
	Name  string    `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`                     //名称
	Level uint32    `protobuf:"varint,3,opt,name=Level,proto3" json:"Level,omitempty"`                  //等级
	Users []uint32  `protobuf:"bytes,4,rep,name=Users,proto3" json:"Users,omitempty"`                   //人员
	Type  EDataType `protobuf:"varint,5,opt,name=Type,proto3,enum=gen.EDataType" json:"Type,omitempty"` //数据类型
}

func (x *PClass) Reset() {
	*x = PClass{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DataMode_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PClass) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PClass) ProtoMessage() {}

func (x *PClass) ProtoReflect() protoreflect.Message {
	mi := &file_DataMode_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PClass.ProtoReflect.Descriptor instead.
func (*PClass) Descriptor() ([]byte, []int) {
	return file_DataMode_proto_rawDescGZIP(), []int{4}
}

func (x *PClass) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *PClass) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PClass) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *PClass) GetUsers() []uint32 {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *PClass) GetType() EDataType {
	if x != nil {
		return x.Type
	}
	return EDataType_UNKNOWN
}

// Defined in table: data/model.xlsx:class
type PClass_ARRAY struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*PClass `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (x *PClass_ARRAY) Reset() {
	*x = PClass_ARRAY{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DataMode_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PClass_ARRAY) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PClass_ARRAY) ProtoMessage() {}

func (x *PClass_ARRAY) ProtoReflect() protoreflect.Message {
	mi := &file_DataMode_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PClass_ARRAY.ProtoReflect.Descriptor instead.
func (*PClass_ARRAY) Descriptor() ([]byte, []int) {
	return file_DataMode_proto_rawDescGZIP(), []int{5}
}

func (x *PClass_ARRAY) GetItems() []*PClass {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_DataMode_proto protoreflect.FileDescriptor

var file_DataMode_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x67, 0x65, 0x6e, 0x22, 0x93, 0x01, 0x0a, 0x05, 0x53, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x41, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x03, 0x53, 0x65, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x49, 0x64, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x49, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x6f,
	0x70, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x12,
	0x2a, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0e, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x45, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x22, 0x2f, 0x0a, 0x0b, 0x53,
	0x49, 0x6e, 0x66, 0x6f, 0x5f, 0x41, 0x52, 0x52, 0x41, 0x59, 0x12, 0x20, 0x0a, 0x05, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x65, 0x6e, 0x2e,
	0x53, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x7a, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x41, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x65,
	0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x53, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04,
	0x48, 0x65, 0x61, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x48, 0x65, 0x61, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x50, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x50, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x22, 0x2d, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72,
	0x5f, 0x41, 0x52, 0x52, 0x41, 0x59, 0x12, 0x1f, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x7c, 0x0a, 0x06, 0x50, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x05, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x12, 0x22, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0e, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x45, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0x31, 0x0a, 0x0c, 0x50, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x5f,
	0x41, 0x52, 0x52, 0x41, 0x59, 0x12, 0x21, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x50, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x2a, 0x43, 0x0a, 0x09, 0x45, 0x44, 0x61, 0x74,
	0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x53, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03,
	0x58, 0x4d, 0x4c, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x53, 0x48, 0x41, 0x52, 0x50, 0x10,
	0x05, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x4f, 0x4c, 0x41, 0x4e, 0x47, 0x10, 0x06, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DataMode_proto_rawDescOnce sync.Once
	file_DataMode_proto_rawDescData = file_DataMode_proto_rawDesc
)

func file_DataMode_proto_rawDescGZIP() []byte {
	file_DataMode_proto_rawDescOnce.Do(func() {
		file_DataMode_proto_rawDescData = protoimpl.X.CompressGZIP(file_DataMode_proto_rawDescData)
	})
	return file_DataMode_proto_rawDescData
}

var file_DataMode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_DataMode_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_DataMode_proto_goTypes = []interface{}{
	(EDataType)(0),       // 0: gen.EDataType
	(*SInfo)(nil),        // 1: gen.SInfo
	(*SInfo_ARRAY)(nil),  // 2: gen.SInfo_ARRAY
	(*User)(nil),         // 3: gen.User
	(*User_ARRAY)(nil),   // 4: gen.User_ARRAY
	(*PClass)(nil),       // 5: gen.PClass
	(*PClass_ARRAY)(nil), // 6: gen.PClass_ARRAY

}

var file_DataMode_proto_depIdxs = []int32{
	0, // 0: gen.SInfo:type_name -> gen.EDataType
	1, // 1: gen.SInfo_ARRAY:type_name -> gen.SInfo
	3, // 2: gen.User_ARRAY:type_name -> gen.User
	0, // 3: gen.PClass:type_name -> gen.EDataType
	5, // 4: gen.PClass_ARRAY:type_name -> gen.PClass
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_DataMode_proto_init() }
func file_DataMode_proto_init() {
	if File_DataMode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_DataMode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SInfo); i {
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

		file_DataMode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SInfo_ARRAY); i {
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

		file_DataMode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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

		file_DataMode_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User_ARRAY); i {
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

		file_DataMode_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PClass); i {
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

		file_DataMode_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PClass_ARRAY); i {
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
			RawDescriptor: file_DataMode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DataMode_proto_goTypes,
		DependencyIndexes: file_DataMode_proto_depIdxs,
		EnumInfos:         file_DataMode_proto_enumTypes,
		MessageInfos:      file_DataMode_proto_msgTypes,
	}.Build()
	File_DataMode_proto = out.File
	file_DataMode_proto_rawDesc = nil
	file_DataMode_proto_goTypes = nil
	file_DataMode_proto_depIdxs = nil
}