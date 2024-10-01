// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: extension.proto

package hiddifyrpc

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

type ExtensionResponseType int32

const (
	ExtensionResponseType_NOTHING     ExtensionResponseType = 0
	ExtensionResponseType_UPDATE_UI   ExtensionResponseType = 1
	ExtensionResponseType_SHOW_DIALOG ExtensionResponseType = 2
	ExtensionResponseType_END         ExtensionResponseType = 3
)

// Enum value maps for ExtensionResponseType.
var (
	ExtensionResponseType_name = map[int32]string{
		0: "NOTHING",
		1: "UPDATE_UI",
		2: "SHOW_DIALOG",
		3: "END",
	}
	ExtensionResponseType_value = map[string]int32{
		"NOTHING":     0,
		"UPDATE_UI":   1,
		"SHOW_DIALOG": 2,
		"END":         3,
	}
)

func (x ExtensionResponseType) Enum() *ExtensionResponseType {
	p := new(ExtensionResponseType)
	*p = x
	return p
}

func (x ExtensionResponseType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExtensionResponseType) Descriptor() protoreflect.EnumDescriptor {
	return file_extension_proto_enumTypes[0].Descriptor()
}

func (ExtensionResponseType) Type() protoreflect.EnumType {
	return &file_extension_proto_enumTypes[0]
}

func (x ExtensionResponseType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExtensionResponseType.Descriptor instead.
func (ExtensionResponseType) EnumDescriptor() ([]byte, []int) {
	return file_extension_proto_rawDescGZIP(), []int{0}
}

type ExtensionActionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExtensionId string       `protobuf:"bytes,1,opt,name=extension_id,json=extensionId,proto3" json:"extension_id,omitempty"`
	Code        ResponseCode `protobuf:"varint,2,opt,name=code,proto3,enum=hiddifyrpc.ResponseCode" json:"code,omitempty"`
	Message     string       `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ExtensionActionResult) Reset() {
	*x = ExtensionActionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extension_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtensionActionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtensionActionResult) ProtoMessage() {}

func (x *ExtensionActionResult) ProtoReflect() protoreflect.Message {
	mi := &file_extension_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtensionActionResult.ProtoReflect.Descriptor instead.
func (*ExtensionActionResult) Descriptor() ([]byte, []int) {
	return file_extension_proto_rawDescGZIP(), []int{0}
}

func (x *ExtensionActionResult) GetExtensionId() string {
	if x != nil {
		return x.ExtensionId
	}
	return ""
}

func (x *ExtensionActionResult) GetCode() ResponseCode {
	if x != nil {
		return x.Code
	}
	return ResponseCode_OK
}

func (x *ExtensionActionResult) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ExtensionList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Extensions []*Extension `protobuf:"bytes,1,rep,name=extensions,proto3" json:"extensions,omitempty"`
}

func (x *ExtensionList) Reset() {
	*x = ExtensionList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extension_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtensionList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtensionList) ProtoMessage() {}

func (x *ExtensionList) ProtoReflect() protoreflect.Message {
	mi := &file_extension_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtensionList.ProtoReflect.Descriptor instead.
func (*ExtensionList) Descriptor() ([]byte, []int) {
	return file_extension_proto_rawDescGZIP(), []int{1}
}

func (x *ExtensionList) GetExtensions() []*Extension {
	if x != nil {
		return x.Extensions
	}
	return nil
}

type EditExtensionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExtensionId string `protobuf:"bytes,1,opt,name=extension_id,json=extensionId,proto3" json:"extension_id,omitempty"`
	Enable      bool   `protobuf:"varint,2,opt,name=enable,proto3" json:"enable,omitempty"`
}

func (x *EditExtensionRequest) Reset() {
	*x = EditExtensionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extension_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditExtensionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditExtensionRequest) ProtoMessage() {}

func (x *EditExtensionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_extension_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditExtensionRequest.ProtoReflect.Descriptor instead.
func (*EditExtensionRequest) Descriptor() ([]byte, []int) {
	return file_extension_proto_rawDescGZIP(), []int{2}
}

func (x *EditExtensionRequest) GetExtensionId() string {
	if x != nil {
		return x.ExtensionId
	}
	return ""
}

func (x *EditExtensionRequest) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

type Extension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Enable      bool   `protobuf:"varint,4,opt,name=enable,proto3" json:"enable,omitempty"`
}

func (x *Extension) Reset() {
	*x = Extension{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extension_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Extension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Extension) ProtoMessage() {}

func (x *Extension) ProtoReflect() protoreflect.Message {
	mi := &file_extension_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Extension.ProtoReflect.Descriptor instead.
func (*Extension) Descriptor() ([]byte, []int) {
	return file_extension_proto_rawDescGZIP(), []int{3}
}

func (x *Extension) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Extension) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Extension) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Extension) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

type ExtensionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExtensionId string            `protobuf:"bytes,1,opt,name=extension_id,json=extensionId,proto3" json:"extension_id,omitempty"`
	Data        map[string]string `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ExtensionRequest) Reset() {
	*x = ExtensionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extension_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtensionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtensionRequest) ProtoMessage() {}

func (x *ExtensionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_extension_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtensionRequest.ProtoReflect.Descriptor instead.
func (*ExtensionRequest) Descriptor() ([]byte, []int) {
	return file_extension_proto_rawDescGZIP(), []int{4}
}

func (x *ExtensionRequest) GetExtensionId() string {
	if x != nil {
		return x.ExtensionId
	}
	return ""
}

func (x *ExtensionRequest) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

type ExtensionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        ExtensionResponseType `protobuf:"varint,1,opt,name=type,proto3,enum=hiddifyrpc.ExtensionResponseType" json:"type,omitempty"`
	ExtensionId string                `protobuf:"bytes,2,opt,name=extension_id,json=extensionId,proto3" json:"extension_id,omitempty"`
	JsonUi      string                `protobuf:"bytes,3,opt,name=json_ui,json=jsonUi,proto3" json:"json_ui,omitempty"`
}

func (x *ExtensionResponse) Reset() {
	*x = ExtensionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extension_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtensionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtensionResponse) ProtoMessage() {}

func (x *ExtensionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_extension_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtensionResponse.ProtoReflect.Descriptor instead.
func (*ExtensionResponse) Descriptor() ([]byte, []int) {
	return file_extension_proto_rawDescGZIP(), []int{5}
}

func (x *ExtensionResponse) GetType() ExtensionResponseType {
	if x != nil {
		return x.Type
	}
	return ExtensionResponseType_NOTHING
}

func (x *ExtensionResponse) GetExtensionId() string {
	if x != nil {
		return x.ExtensionId
	}
	return ""
}

func (x *ExtensionResponse) GetJsonUi() string {
	if x != nil {
		return x.JsonUi
	}
	return ""
}

var File_extension_proto protoreflect.FileDescriptor

var file_extension_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x68, 0x69, 0x64, 0x64, 0x69, 0x66, 0x79, 0x72, 0x70, 0x63, 0x1a, 0x0a, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x15, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x68, 0x69, 0x64, 0x64, 0x69, 0x66, 0x79, 0x72, 0x70,
	0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x46,
	0x0a, 0x0d, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x35, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x68, 0x69, 0x64, 0x64, 0x69, 0x66, 0x79, 0x72, 0x70, 0x63,
	0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x51, 0x0a, 0x14, 0x45, 0x64, 0x69, 0x74, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x6b, 0x0a, 0x09, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x22, 0xaa, 0x01, 0x0a, 0x10, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x3a,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x68,
	0x69, 0x64, 0x64, 0x69, 0x66, 0x79, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x86, 0x01, 0x0a, 0x11, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x68, 0x69, 0x64, 0x64, 0x69, 0x66,
	0x79, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x75, 0x69, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6a, 0x73, 0x6f, 0x6e, 0x55, 0x69, 0x2a, 0x4d, 0x0a, 0x15,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x4f, 0x54, 0x48, 0x49, 0x4e, 0x47,
	0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x49, 0x10,
	0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x48, 0x4f, 0x57, 0x5f, 0x44, 0x49, 0x41, 0x4c, 0x4f, 0x47,
	0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x4e, 0x44, 0x10, 0x03, 0x32, 0xb1, 0x04, 0x0a, 0x14,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x11, 0x2e, 0x68, 0x69, 0x64, 0x64, 0x69, 0x66, 0x79,
	0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x68, 0x69, 0x64, 0x64,
	0x69, 0x66, 0x79, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x12, 0x1c, 0x2e, 0x68, 0x69, 0x64, 0x64, 0x69, 0x66, 0x79, 0x72, 0x70, 0x63, 0x2e, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x68, 0x69, 0x64, 0x64, 0x69, 0x66, 0x79, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x12, 0x56, 0x0a, 0x0d, 0x45, 0x64, 0x69, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x68, 0x69, 0x64, 0x64, 0x69, 0x66, 0x79, 0x72, 0x70, 0x63,
	0x2e, 0x45, 0x64, 0x69, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x68, 0x69, 0x64, 0x64, 0x69, 0x66, 0x79, 0x72,
	0x70, 0x63, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0a, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x12, 0x1c, 0x2e, 0x68, 0x69, 0x64, 0x64, 0x69,
	0x66, 0x79, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x68, 0x69, 0x64, 0x64, 0x69, 0x66, 0x79,
	0x72, 0x70, 0x63, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x06, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x12, 0x1c, 0x2e, 0x68, 0x69, 0x64, 0x64, 0x69, 0x66, 0x79, 0x72,
	0x70, 0x63, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x68, 0x69, 0x64, 0x64, 0x69, 0x66, 0x79, 0x72, 0x70, 0x63,
	0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70,
	0x12, 0x1c, 0x2e, 0x68, 0x69, 0x64, 0x64, 0x69, 0x66, 0x79, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x68, 0x69, 0x64, 0x64, 0x69, 0x66, 0x79, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x55, 0x49, 0x12, 0x1c, 0x2e, 0x68,
	0x69, 0x64, 0x64, 0x69, 0x66, 0x79, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x68, 0x69, 0x64,
	0x64, 0x69, 0x66, 0x79, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42,
	0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x68, 0x69, 0x64, 0x64, 0x69, 0x66, 0x79, 0x72, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_extension_proto_rawDescOnce sync.Once
	file_extension_proto_rawDescData = file_extension_proto_rawDesc
)

func file_extension_proto_rawDescGZIP() []byte {
	file_extension_proto_rawDescOnce.Do(func() {
		file_extension_proto_rawDescData = protoimpl.X.CompressGZIP(file_extension_proto_rawDescData)
	})
	return file_extension_proto_rawDescData
}

var file_extension_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_extension_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_extension_proto_goTypes = []any{
	(ExtensionResponseType)(0),    // 0: hiddifyrpc.ExtensionResponseType
	(*ExtensionActionResult)(nil), // 1: hiddifyrpc.ExtensionActionResult
	(*ExtensionList)(nil),         // 2: hiddifyrpc.ExtensionList
	(*EditExtensionRequest)(nil),  // 3: hiddifyrpc.EditExtensionRequest
	(*Extension)(nil),             // 4: hiddifyrpc.Extension
	(*ExtensionRequest)(nil),      // 5: hiddifyrpc.ExtensionRequest
	(*ExtensionResponse)(nil),     // 6: hiddifyrpc.ExtensionResponse
	nil,                           // 7: hiddifyrpc.ExtensionRequest.DataEntry
	(ResponseCode)(0),             // 8: hiddifyrpc.ResponseCode
	(*Empty)(nil),                 // 9: hiddifyrpc.Empty
}
var file_extension_proto_depIdxs = []int32{
	8,  // 0: hiddifyrpc.ExtensionActionResult.code:type_name -> hiddifyrpc.ResponseCode
	4,  // 1: hiddifyrpc.ExtensionList.extensions:type_name -> hiddifyrpc.Extension
	7,  // 2: hiddifyrpc.ExtensionRequest.data:type_name -> hiddifyrpc.ExtensionRequest.DataEntry
	0,  // 3: hiddifyrpc.ExtensionResponse.type:type_name -> hiddifyrpc.ExtensionResponseType
	9,  // 4: hiddifyrpc.ExtensionHostService.ListExtensions:input_type -> hiddifyrpc.Empty
	5,  // 5: hiddifyrpc.ExtensionHostService.Connect:input_type -> hiddifyrpc.ExtensionRequest
	3,  // 6: hiddifyrpc.ExtensionHostService.EditExtension:input_type -> hiddifyrpc.EditExtensionRequest
	5,  // 7: hiddifyrpc.ExtensionHostService.SubmitForm:input_type -> hiddifyrpc.ExtensionRequest
	5,  // 8: hiddifyrpc.ExtensionHostService.Cancel:input_type -> hiddifyrpc.ExtensionRequest
	5,  // 9: hiddifyrpc.ExtensionHostService.Stop:input_type -> hiddifyrpc.ExtensionRequest
	5,  // 10: hiddifyrpc.ExtensionHostService.GetUI:input_type -> hiddifyrpc.ExtensionRequest
	2,  // 11: hiddifyrpc.ExtensionHostService.ListExtensions:output_type -> hiddifyrpc.ExtensionList
	6,  // 12: hiddifyrpc.ExtensionHostService.Connect:output_type -> hiddifyrpc.ExtensionResponse
	1,  // 13: hiddifyrpc.ExtensionHostService.EditExtension:output_type -> hiddifyrpc.ExtensionActionResult
	1,  // 14: hiddifyrpc.ExtensionHostService.SubmitForm:output_type -> hiddifyrpc.ExtensionActionResult
	1,  // 15: hiddifyrpc.ExtensionHostService.Cancel:output_type -> hiddifyrpc.ExtensionActionResult
	1,  // 16: hiddifyrpc.ExtensionHostService.Stop:output_type -> hiddifyrpc.ExtensionActionResult
	1,  // 17: hiddifyrpc.ExtensionHostService.GetUI:output_type -> hiddifyrpc.ExtensionActionResult
	11, // [11:18] is the sub-list for method output_type
	4,  // [4:11] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_extension_proto_init() }
func file_extension_proto_init() {
	if File_extension_proto != nil {
		return
	}
	file_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_extension_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ExtensionActionResult); i {
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
		file_extension_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ExtensionList); i {
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
		file_extension_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*EditExtensionRequest); i {
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
		file_extension_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Extension); i {
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
		file_extension_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ExtensionRequest); i {
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
		file_extension_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ExtensionResponse); i {
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
			RawDescriptor: file_extension_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_extension_proto_goTypes,
		DependencyIndexes: file_extension_proto_depIdxs,
		EnumInfos:         file_extension_proto_enumTypes,
		MessageInfos:      file_extension_proto_msgTypes,
	}.Build()
	File_extension_proto = out.File
	file_extension_proto_rawDesc = nil
	file_extension_proto_goTypes = nil
	file_extension_proto_depIdxs = nil
}
