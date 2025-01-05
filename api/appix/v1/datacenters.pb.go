// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: api/appix/v1/datacenters.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// gratos::model
type Datacenter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Datacenter) Reset() {
	*x = Datacenter{}
	mi := &file_api_appix_v1_datacenters_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Datacenter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Datacenter) ProtoMessage() {}

func (x *Datacenter) ProtoReflect() protoreflect.Message {
	mi := &file_api_appix_v1_datacenters_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Datacenter.ProtoReflect.Descriptor instead.
func (*Datacenter) Descriptor() ([]byte, []int) {
	return file_api_appix_v1_datacenters_proto_rawDescGZIP(), []int{0}
}

func (x *Datacenter) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Datacenter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Datacenter) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateDatacentersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datacenters []*Datacenter `protobuf:"bytes,1,rep,name=datacenters,proto3" json:"datacenters,omitempty"`
}

func (x *CreateDatacentersRequest) Reset() {
	*x = CreateDatacentersRequest{}
	mi := &file_api_appix_v1_datacenters_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateDatacentersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDatacentersRequest) ProtoMessage() {}

func (x *CreateDatacentersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_appix_v1_datacenters_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDatacentersRequest.ProtoReflect.Descriptor instead.
func (*CreateDatacentersRequest) Descriptor() ([]byte, []int) {
	return file_api_appix_v1_datacenters_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDatacentersRequest) GetDatacenters() []*Datacenter {
	if x != nil {
		return x.Datacenters
	}
	return nil
}

type CreateDatacentersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code    int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Action  string `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *CreateDatacentersReply) Reset() {
	*x = CreateDatacentersReply{}
	mi := &file_api_appix_v1_datacenters_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateDatacentersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDatacentersReply) ProtoMessage() {}

func (x *CreateDatacentersReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_appix_v1_datacenters_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDatacentersReply.ProtoReflect.Descriptor instead.
func (*CreateDatacentersReply) Descriptor() ([]byte, []int) {
	return file_api_appix_v1_datacenters_proto_rawDescGZIP(), []int{2}
}

func (x *CreateDatacentersReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateDatacentersReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CreateDatacentersReply) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type UpdateDatacentersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datacenters []*Datacenter `protobuf:"bytes,1,rep,name=datacenters,proto3" json:"datacenters,omitempty"`
}

func (x *UpdateDatacentersRequest) Reset() {
	*x = UpdateDatacentersRequest{}
	mi := &file_api_appix_v1_datacenters_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateDatacentersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDatacentersRequest) ProtoMessage() {}

func (x *UpdateDatacentersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_appix_v1_datacenters_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDatacentersRequest.ProtoReflect.Descriptor instead.
func (*UpdateDatacentersRequest) Descriptor() ([]byte, []int) {
	return file_api_appix_v1_datacenters_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateDatacentersRequest) GetDatacenters() []*Datacenter {
	if x != nil {
		return x.Datacenters
	}
	return nil
}

type UpdateDatacentersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code    int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Action  string `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *UpdateDatacentersReply) Reset() {
	*x = UpdateDatacentersReply{}
	mi := &file_api_appix_v1_datacenters_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateDatacentersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDatacentersReply) ProtoMessage() {}

func (x *UpdateDatacentersReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_appix_v1_datacenters_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDatacentersReply.ProtoReflect.Descriptor instead.
func (*UpdateDatacentersReply) Descriptor() ([]byte, []int) {
	return file_api_appix_v1_datacenters_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateDatacentersReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UpdateDatacentersReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *UpdateDatacentersReply) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type DeleteDatacentersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []uint32 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *DeleteDatacentersRequest) Reset() {
	*x = DeleteDatacentersRequest{}
	mi := &file_api_appix_v1_datacenters_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteDatacentersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDatacentersRequest) ProtoMessage() {}

func (x *DeleteDatacentersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_appix_v1_datacenters_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDatacentersRequest.ProtoReflect.Descriptor instead.
func (*DeleteDatacentersRequest) Descriptor() ([]byte, []int) {
	return file_api_appix_v1_datacenters_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteDatacentersRequest) GetIds() []uint32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type DeleteDatacentersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code    int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Action  string `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *DeleteDatacentersReply) Reset() {
	*x = DeleteDatacentersReply{}
	mi := &file_api_appix_v1_datacenters_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteDatacentersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDatacentersReply) ProtoMessage() {}

func (x *DeleteDatacentersReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_appix_v1_datacenters_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDatacentersReply.ProtoReflect.Descriptor instead.
func (*DeleteDatacentersReply) Descriptor() ([]byte, []int) {
	return file_api_appix_v1_datacenters_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteDatacentersReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DeleteDatacentersReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DeleteDatacentersReply) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type GetDatacentersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetDatacentersRequest) Reset() {
	*x = GetDatacentersRequest{}
	mi := &file_api_appix_v1_datacenters_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDatacentersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDatacentersRequest) ProtoMessage() {}

func (x *GetDatacentersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_appix_v1_datacenters_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDatacentersRequest.ProtoReflect.Descriptor instead.
func (*GetDatacentersRequest) Descriptor() ([]byte, []int) {
	return file_api_appix_v1_datacenters_proto_rawDescGZIP(), []int{7}
}

func (x *GetDatacentersRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetDatacentersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message    string      `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code       int32       `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Action     string      `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	Datacenter *Datacenter `protobuf:"bytes,4,opt,name=datacenter,proto3" json:"datacenter,omitempty"`
}

func (x *GetDatacentersReply) Reset() {
	*x = GetDatacentersReply{}
	mi := &file_api_appix_v1_datacenters_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDatacentersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDatacentersReply) ProtoMessage() {}

func (x *GetDatacentersReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_appix_v1_datacenters_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDatacentersReply.ProtoReflect.Descriptor instead.
func (*GetDatacentersReply) Descriptor() ([]byte, []int) {
	return file_api_appix_v1_datacenters_proto_rawDescGZIP(), []int{8}
}

func (x *GetDatacentersReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetDatacentersReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetDatacentersReply) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *GetDatacentersReply) GetDatacenter() *Datacenter {
	if x != nil {
		return x.Datacenter
	}
	return nil
}

type ListDatacentersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     uint32   `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize uint32   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Ids      []uint32 `protobuf:"varint,3,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Names    []string `protobuf:"bytes,4,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *ListDatacentersRequest) Reset() {
	*x = ListDatacentersRequest{}
	mi := &file_api_appix_v1_datacenters_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListDatacentersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDatacentersRequest) ProtoMessage() {}

func (x *ListDatacentersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_appix_v1_datacenters_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDatacentersRequest.ProtoReflect.Descriptor instead.
func (*ListDatacentersRequest) Descriptor() ([]byte, []int) {
	return file_api_appix_v1_datacenters_proto_rawDescGZIP(), []int{9}
}

func (x *ListDatacentersRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListDatacentersRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListDatacentersRequest) GetIds() []uint32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *ListDatacentersRequest) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

type ListDatacentersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message     string        `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code        int32         `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Action      string        `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	Datacenters []*Datacenter `protobuf:"bytes,4,rep,name=datacenters,proto3" json:"datacenters,omitempty"`
}

func (x *ListDatacentersReply) Reset() {
	*x = ListDatacentersReply{}
	mi := &file_api_appix_v1_datacenters_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListDatacentersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDatacentersReply) ProtoMessage() {}

func (x *ListDatacentersReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_appix_v1_datacenters_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDatacentersReply.ProtoReflect.Descriptor instead.
func (*ListDatacentersReply) Descriptor() ([]byte, []int) {
	return file_api_appix_v1_datacenters_proto_rawDescGZIP(), []int{10}
}

func (x *ListDatacentersReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ListDatacentersReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ListDatacentersReply) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *ListDatacentersReply) GetDatacenters() []*Datacenter {
	if x != nil {
		return x.Datacenters
	}
	return nil
}

var File_api_appix_v1_datacenters_proto protoreflect.FileDescriptor

var file_api_appix_v1_datacenters_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x69, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x0a,
	0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x56, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0b,
	0x64, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x69, 0x78, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x0b, 0x64, 0x61, 0x74,
	0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x22, 0x5e, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x56, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x70, 0x70, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73,
	0x22, 0x5e, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x2c, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x5e,
	0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x27,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x95, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x70, 0x70, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x22,
	0x71, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x22, 0x98, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x3a, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70,
	0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x52, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x32, 0xad, 0x05,
	0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x12, 0x88, 0x01,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x73, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x69, 0x78, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x70, 0x70, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x12, 0x26,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70,
	0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x25, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x70, 0x70, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x69, 0x78, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a,
	0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x7a,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73,
	0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x69,
	0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a,
	0x12, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x80, 0x01, 0x0a, 0x0f, 0x4c,
	0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x12, 0x24,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x69, 0x78,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d,
	0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x27, 0x0a,
	0x0c, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a,
	0x15, 0x61, 0x70, 0x70, 0x69, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x69, 0x78,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_appix_v1_datacenters_proto_rawDescOnce sync.Once
	file_api_appix_v1_datacenters_proto_rawDescData = file_api_appix_v1_datacenters_proto_rawDesc
)

func file_api_appix_v1_datacenters_proto_rawDescGZIP() []byte {
	file_api_appix_v1_datacenters_proto_rawDescOnce.Do(func() {
		file_api_appix_v1_datacenters_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_appix_v1_datacenters_proto_rawDescData)
	})
	return file_api_appix_v1_datacenters_proto_rawDescData
}

var file_api_appix_v1_datacenters_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_appix_v1_datacenters_proto_goTypes = []any{
	(*Datacenter)(nil),               // 0: api.appix.v1.Datacenter
	(*CreateDatacentersRequest)(nil), // 1: api.appix.v1.CreateDatacentersRequest
	(*CreateDatacentersReply)(nil),   // 2: api.appix.v1.CreateDatacentersReply
	(*UpdateDatacentersRequest)(nil), // 3: api.appix.v1.UpdateDatacentersRequest
	(*UpdateDatacentersReply)(nil),   // 4: api.appix.v1.UpdateDatacentersReply
	(*DeleteDatacentersRequest)(nil), // 5: api.appix.v1.DeleteDatacentersRequest
	(*DeleteDatacentersReply)(nil),   // 6: api.appix.v1.DeleteDatacentersReply
	(*GetDatacentersRequest)(nil),    // 7: api.appix.v1.GetDatacentersRequest
	(*GetDatacentersReply)(nil),      // 8: api.appix.v1.GetDatacentersReply
	(*ListDatacentersRequest)(nil),   // 9: api.appix.v1.ListDatacentersRequest
	(*ListDatacentersReply)(nil),     // 10: api.appix.v1.ListDatacentersReply
}
var file_api_appix_v1_datacenters_proto_depIdxs = []int32{
	0,  // 0: api.appix.v1.CreateDatacentersRequest.datacenters:type_name -> api.appix.v1.Datacenter
	0,  // 1: api.appix.v1.UpdateDatacentersRequest.datacenters:type_name -> api.appix.v1.Datacenter
	0,  // 2: api.appix.v1.GetDatacentersReply.datacenter:type_name -> api.appix.v1.Datacenter
	0,  // 3: api.appix.v1.ListDatacentersReply.datacenters:type_name -> api.appix.v1.Datacenter
	1,  // 4: api.appix.v1.Datacenters.CreateDatacenters:input_type -> api.appix.v1.CreateDatacentersRequest
	3,  // 5: api.appix.v1.Datacenters.UpdateDatacenters:input_type -> api.appix.v1.UpdateDatacentersRequest
	5,  // 6: api.appix.v1.Datacenters.DeleteDatacenters:input_type -> api.appix.v1.DeleteDatacentersRequest
	7,  // 7: api.appix.v1.Datacenters.GetDatacenters:input_type -> api.appix.v1.GetDatacentersRequest
	9,  // 8: api.appix.v1.Datacenters.ListDatacenters:input_type -> api.appix.v1.ListDatacentersRequest
	2,  // 9: api.appix.v1.Datacenters.CreateDatacenters:output_type -> api.appix.v1.CreateDatacentersReply
	4,  // 10: api.appix.v1.Datacenters.UpdateDatacenters:output_type -> api.appix.v1.UpdateDatacentersReply
	6,  // 11: api.appix.v1.Datacenters.DeleteDatacenters:output_type -> api.appix.v1.DeleteDatacentersReply
	8,  // 12: api.appix.v1.Datacenters.GetDatacenters:output_type -> api.appix.v1.GetDatacentersReply
	10, // 13: api.appix.v1.Datacenters.ListDatacenters:output_type -> api.appix.v1.ListDatacentersReply
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_appix_v1_datacenters_proto_init() }
func file_api_appix_v1_datacenters_proto_init() {
	if File_api_appix_v1_datacenters_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_appix_v1_datacenters_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_appix_v1_datacenters_proto_goTypes,
		DependencyIndexes: file_api_appix_v1_datacenters_proto_depIdxs,
		MessageInfos:      file_api_appix_v1_datacenters_proto_msgTypes,
	}.Build()
	File_api_appix_v1_datacenters_proto = out.File
	file_api_appix_v1_datacenters_proto_rawDesc = nil
	file_api_appix_v1_datacenters_proto_goTypes = nil
	file_api_appix_v1_datacenters_proto_depIdxs = nil
}
