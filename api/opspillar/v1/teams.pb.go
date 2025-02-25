// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: opspillar/v1/teams.proto

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
type Team struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Code        string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	LeaderId    uint32 `protobuf:"varint,5,opt,name=leader_id,json=leaderId,proto3" json:"leader_id,omitempty"`
}

func (x *Team) Reset() {
	*x = Team{}
	mi := &file_opspillar_v1_teams_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Team) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Team) ProtoMessage() {}

func (x *Team) ProtoReflect() protoreflect.Message {
	mi := &file_opspillar_v1_teams_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Team.ProtoReflect.Descriptor instead.
func (*Team) Descriptor() ([]byte, []int) {
	return file_opspillar_v1_teams_proto_rawDescGZIP(), []int{0}
}

func (x *Team) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Team) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Team) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Team) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Team) GetLeaderId() uint32 {
	if x != nil {
		return x.LeaderId
	}
	return 0
}

type TeamReadable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Code        string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Leader      string `protobuf:"bytes,5,opt,name=leader,proto3" json:"leader,omitempty"`
}

func (x *TeamReadable) Reset() {
	*x = TeamReadable{}
	mi := &file_opspillar_v1_teams_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TeamReadable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamReadable) ProtoMessage() {}

func (x *TeamReadable) ProtoReflect() protoreflect.Message {
	mi := &file_opspillar_v1_teams_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamReadable.ProtoReflect.Descriptor instead.
func (*TeamReadable) Descriptor() ([]byte, []int) {
	return file_opspillar_v1_teams_proto_rawDescGZIP(), []int{1}
}

func (x *TeamReadable) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TeamReadable) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TeamReadable) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *TeamReadable) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TeamReadable) GetLeader() string {
	if x != nil {
		return x.Leader
	}
	return ""
}

type CreateTeamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Teams []*Team `protobuf:"bytes,1,rep,name=teams,proto3" json:"teams,omitempty"`
}

func (x *CreateTeamsRequest) Reset() {
	*x = CreateTeamsRequest{}
	mi := &file_opspillar_v1_teams_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTeamsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTeamsRequest) ProtoMessage() {}

func (x *CreateTeamsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_opspillar_v1_teams_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTeamsRequest.ProtoReflect.Descriptor instead.
func (*CreateTeamsRequest) Descriptor() ([]byte, []int) {
	return file_opspillar_v1_teams_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTeamsRequest) GetTeams() []*Team {
	if x != nil {
		return x.Teams
	}
	return nil
}

type CreateTeamsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code    int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Action  string `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *CreateTeamsReply) Reset() {
	*x = CreateTeamsReply{}
	mi := &file_opspillar_v1_teams_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTeamsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTeamsReply) ProtoMessage() {}

func (x *CreateTeamsReply) ProtoReflect() protoreflect.Message {
	mi := &file_opspillar_v1_teams_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTeamsReply.ProtoReflect.Descriptor instead.
func (*CreateTeamsReply) Descriptor() ([]byte, []int) {
	return file_opspillar_v1_teams_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTeamsReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateTeamsReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CreateTeamsReply) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type UpdateTeamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Teams []*Team `protobuf:"bytes,1,rep,name=teams,proto3" json:"teams,omitempty"`
}

func (x *UpdateTeamsRequest) Reset() {
	*x = UpdateTeamsRequest{}
	mi := &file_opspillar_v1_teams_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTeamsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTeamsRequest) ProtoMessage() {}

func (x *UpdateTeamsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_opspillar_v1_teams_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTeamsRequest.ProtoReflect.Descriptor instead.
func (*UpdateTeamsRequest) Descriptor() ([]byte, []int) {
	return file_opspillar_v1_teams_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateTeamsRequest) GetTeams() []*Team {
	if x != nil {
		return x.Teams
	}
	return nil
}

type UpdateTeamsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code    int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Action  string `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *UpdateTeamsReply) Reset() {
	*x = UpdateTeamsReply{}
	mi := &file_opspillar_v1_teams_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTeamsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTeamsReply) ProtoMessage() {}

func (x *UpdateTeamsReply) ProtoReflect() protoreflect.Message {
	mi := &file_opspillar_v1_teams_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTeamsReply.ProtoReflect.Descriptor instead.
func (*UpdateTeamsReply) Descriptor() ([]byte, []int) {
	return file_opspillar_v1_teams_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateTeamsReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UpdateTeamsReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *UpdateTeamsReply) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type DeleteTeamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []uint32 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *DeleteTeamsRequest) Reset() {
	*x = DeleteTeamsRequest{}
	mi := &file_opspillar_v1_teams_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTeamsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTeamsRequest) ProtoMessage() {}

func (x *DeleteTeamsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_opspillar_v1_teams_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTeamsRequest.ProtoReflect.Descriptor instead.
func (*DeleteTeamsRequest) Descriptor() ([]byte, []int) {
	return file_opspillar_v1_teams_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteTeamsRequest) GetIds() []uint32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type DeleteTeamsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code    int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Action  string `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *DeleteTeamsReply) Reset() {
	*x = DeleteTeamsReply{}
	mi := &file_opspillar_v1_teams_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTeamsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTeamsReply) ProtoMessage() {}

func (x *DeleteTeamsReply) ProtoReflect() protoreflect.Message {
	mi := &file_opspillar_v1_teams_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTeamsReply.ProtoReflect.Descriptor instead.
func (*DeleteTeamsReply) Descriptor() ([]byte, []int) {
	return file_opspillar_v1_teams_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteTeamsReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DeleteTeamsReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DeleteTeamsReply) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type GetTeamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTeamsRequest) Reset() {
	*x = GetTeamsRequest{}
	mi := &file_opspillar_v1_teams_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTeamsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTeamsRequest) ProtoMessage() {}

func (x *GetTeamsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_opspillar_v1_teams_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTeamsRequest.ProtoReflect.Descriptor instead.
func (*GetTeamsRequest) Descriptor() ([]byte, []int) {
	return file_opspillar_v1_teams_proto_rawDescGZIP(), []int{8}
}

func (x *GetTeamsRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetTeamsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code    int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Action  string `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	Team    *Team  `protobuf:"bytes,4,opt,name=team,proto3" json:"team,omitempty"`
}

func (x *GetTeamsReply) Reset() {
	*x = GetTeamsReply{}
	mi := &file_opspillar_v1_teams_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTeamsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTeamsReply) ProtoMessage() {}

func (x *GetTeamsReply) ProtoReflect() protoreflect.Message {
	mi := &file_opspillar_v1_teams_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTeamsReply.ProtoReflect.Descriptor instead.
func (*GetTeamsReply) Descriptor() ([]byte, []int) {
	return file_opspillar_v1_teams_proto_rawDescGZIP(), []int{9}
}

func (x *GetTeamsReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetTeamsReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetTeamsReply) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *GetTeamsReply) GetTeam() *Team {
	if x != nil {
		return x.Team
	}
	return nil
}

type ListTeamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page      uint32   `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  uint32   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Names     []string `protobuf:"bytes,3,rep,name=names,proto3" json:"names,omitempty"`
	Codes     []string `protobuf:"bytes,4,rep,name=codes,proto3" json:"codes,omitempty"`
	LeadersId []uint32 `protobuf:"varint,5,rep,packed,name=leaders_id,json=leadersId,proto3" json:"leaders_id,omitempty"`
	Ids       []uint32 `protobuf:"varint,6,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *ListTeamsRequest) Reset() {
	*x = ListTeamsRequest{}
	mi := &file_opspillar_v1_teams_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTeamsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTeamsRequest) ProtoMessage() {}

func (x *ListTeamsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_opspillar_v1_teams_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTeamsRequest.ProtoReflect.Descriptor instead.
func (*ListTeamsRequest) Descriptor() ([]byte, []int) {
	return file_opspillar_v1_teams_proto_rawDescGZIP(), []int{10}
}

func (x *ListTeamsRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListTeamsRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListTeamsRequest) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

func (x *ListTeamsRequest) GetCodes() []string {
	if x != nil {
		return x.Codes
	}
	return nil
}

func (x *ListTeamsRequest) GetLeadersId() []uint32 {
	if x != nil {
		return x.LeadersId
	}
	return nil
}

func (x *ListTeamsRequest) GetIds() []uint32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type ListTeamsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string  `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code    int32   `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Action  string  `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	Teams   []*Team `protobuf:"bytes,4,rep,name=teams,proto3" json:"teams,omitempty"`
}

func (x *ListTeamsReply) Reset() {
	*x = ListTeamsReply{}
	mi := &file_opspillar_v1_teams_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTeamsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTeamsReply) ProtoMessage() {}

func (x *ListTeamsReply) ProtoReflect() protoreflect.Message {
	mi := &file_opspillar_v1_teams_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTeamsReply.ProtoReflect.Descriptor instead.
func (*ListTeamsReply) Descriptor() ([]byte, []int) {
	return file_opspillar_v1_teams_proto_rawDescGZIP(), []int{11}
}

func (x *ListTeamsReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ListTeamsReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ListTeamsReply) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *ListTeamsReply) GetTeams() []*Team {
	if x != nil {
		return x.Teams
	}
	return nil
}

var File_opspillar_v1_teams_proto protoreflect.FileDescriptor

var file_opspillar_v1_teams_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6f, 0x70, 0x73, 0x70, 0x69, 0x6c, 0x6c, 0x61, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x65, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61, 0x70, 0x69, 0x2e,
	0x6f, 0x70, 0x73, 0x70, 0x69, 0x6c, 0x6c, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x04, 0x54, 0x65,
	0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x80, 0x01, 0x0a, 0x0c, 0x54, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x42, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x73, 0x70, 0x69, 0x6c, 0x6c, 0x61,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x05, 0x74, 0x65, 0x61, 0x6d, 0x73,
	0x22, 0x58, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x42, 0x0a, 0x12, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2c, 0x0a, 0x05, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x73, 0x70, 0x69, 0x6c, 0x6c, 0x61, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x05, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x22, 0x58,
	0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x26, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x22, 0x58, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x81, 0x01,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x73, 0x70, 0x69, 0x6c,
	0x6c, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x04, 0x74, 0x65, 0x61,
	0x6d, 0x22, 0xa0, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f,
	0x64, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x09, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x03, 0x69, 0x64, 0x73, 0x22, 0x84, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x61,
	0x6d, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a,
	0x05, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6f, 0x70, 0x73, 0x70, 0x69, 0x6c, 0x6c, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x65, 0x61, 0x6d, 0x52, 0x05, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x32, 0xd3, 0x04, 0x0a, 0x05,
	0x54, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x78, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x65, 0x61, 0x6d, 0x73, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x73, 0x70, 0x69,
	0x6c, 0x6c, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6f, 0x70, 0x73, 0x70, 0x69, 0x6c, 0x6c, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x78, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x24,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x73, 0x70, 0x69, 0x6c, 0x6c, 0x61, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x73, 0x70, 0x69,
	0x6c, 0x6c, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x61, 0x6d, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19,
	0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x61,
	0x6d, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x78, 0x0a, 0x0b, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f,
	0x70, 0x73, 0x70, 0x69, 0x6c, 0x6c, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x73, 0x70, 0x69, 0x6c, 0x6c, 0x61, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x6a, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x12,
	0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x73, 0x70, 0x69, 0x6c, 0x6c, 0x61, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x73, 0x70, 0x69, 0x6c, 0x6c,
	0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12,
	0x70, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x22, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6f, 0x70, 0x73, 0x70, 0x69, 0x6c, 0x6c, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x73, 0x70, 0x69, 0x6c, 0x6c, 0x61, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x6c, 0x69, 0x73,
	0x74, 0x42, 0x33, 0x0a, 0x10, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x73, 0x70, 0x69, 0x6c, 0x6c,
	0x61, 0x72, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x1d, 0x6f, 0x70, 0x73, 0x70, 0x69, 0x6c, 0x6c,
	0x61, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70, 0x73, 0x70, 0x69, 0x6c, 0x6c, 0x61, 0x72,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_opspillar_v1_teams_proto_rawDescOnce sync.Once
	file_opspillar_v1_teams_proto_rawDescData = file_opspillar_v1_teams_proto_rawDesc
)

func file_opspillar_v1_teams_proto_rawDescGZIP() []byte {
	file_opspillar_v1_teams_proto_rawDescOnce.Do(func() {
		file_opspillar_v1_teams_proto_rawDescData = protoimpl.X.CompressGZIP(file_opspillar_v1_teams_proto_rawDescData)
	})
	return file_opspillar_v1_teams_proto_rawDescData
}

var file_opspillar_v1_teams_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_opspillar_v1_teams_proto_goTypes = []any{
	(*Team)(nil),               // 0: api.opspillar.v1.Team
	(*TeamReadable)(nil),       // 1: api.opspillar.v1.TeamReadable
	(*CreateTeamsRequest)(nil), // 2: api.opspillar.v1.CreateTeamsRequest
	(*CreateTeamsReply)(nil),   // 3: api.opspillar.v1.CreateTeamsReply
	(*UpdateTeamsRequest)(nil), // 4: api.opspillar.v1.UpdateTeamsRequest
	(*UpdateTeamsReply)(nil),   // 5: api.opspillar.v1.UpdateTeamsReply
	(*DeleteTeamsRequest)(nil), // 6: api.opspillar.v1.DeleteTeamsRequest
	(*DeleteTeamsReply)(nil),   // 7: api.opspillar.v1.DeleteTeamsReply
	(*GetTeamsRequest)(nil),    // 8: api.opspillar.v1.GetTeamsRequest
	(*GetTeamsReply)(nil),      // 9: api.opspillar.v1.GetTeamsReply
	(*ListTeamsRequest)(nil),   // 10: api.opspillar.v1.ListTeamsRequest
	(*ListTeamsReply)(nil),     // 11: api.opspillar.v1.ListTeamsReply
}
var file_opspillar_v1_teams_proto_depIdxs = []int32{
	0,  // 0: api.opspillar.v1.CreateTeamsRequest.teams:type_name -> api.opspillar.v1.Team
	0,  // 1: api.opspillar.v1.UpdateTeamsRequest.teams:type_name -> api.opspillar.v1.Team
	0,  // 2: api.opspillar.v1.GetTeamsReply.team:type_name -> api.opspillar.v1.Team
	0,  // 3: api.opspillar.v1.ListTeamsReply.teams:type_name -> api.opspillar.v1.Team
	2,  // 4: api.opspillar.v1.Teams.CreateTeams:input_type -> api.opspillar.v1.CreateTeamsRequest
	4,  // 5: api.opspillar.v1.Teams.UpdateTeams:input_type -> api.opspillar.v1.UpdateTeamsRequest
	6,  // 6: api.opspillar.v1.Teams.DeleteTeams:input_type -> api.opspillar.v1.DeleteTeamsRequest
	8,  // 7: api.opspillar.v1.Teams.GetTeams:input_type -> api.opspillar.v1.GetTeamsRequest
	10, // 8: api.opspillar.v1.Teams.ListTeams:input_type -> api.opspillar.v1.ListTeamsRequest
	3,  // 9: api.opspillar.v1.Teams.CreateTeams:output_type -> api.opspillar.v1.CreateTeamsReply
	5,  // 10: api.opspillar.v1.Teams.UpdateTeams:output_type -> api.opspillar.v1.UpdateTeamsReply
	7,  // 11: api.opspillar.v1.Teams.DeleteTeams:output_type -> api.opspillar.v1.DeleteTeamsReply
	9,  // 12: api.opspillar.v1.Teams.GetTeams:output_type -> api.opspillar.v1.GetTeamsReply
	11, // 13: api.opspillar.v1.Teams.ListTeams:output_type -> api.opspillar.v1.ListTeamsReply
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_opspillar_v1_teams_proto_init() }
func file_opspillar_v1_teams_proto_init() {
	if File_opspillar_v1_teams_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_opspillar_v1_teams_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_opspillar_v1_teams_proto_goTypes,
		DependencyIndexes: file_opspillar_v1_teams_proto_depIdxs,
		MessageInfos:      file_opspillar_v1_teams_proto_msgTypes,
	}.Build()
	File_opspillar_v1_teams_proto = out.File
	file_opspillar_v1_teams_proto_rawDesc = nil
	file_opspillar_v1_teams_proto_goTypes = nil
	file_opspillar_v1_teams_proto_depIdxs = nil
}
