// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: tours_service.proto

package tours

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

type EquipmentIds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EquipmentIds []string `protobuf:"bytes,2,rep,name=equipmentIds,proto3" json:"equipmentIds,omitempty"`
}

func (x *EquipmentIds) Reset() {
	*x = EquipmentIds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tours_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EquipmentIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EquipmentIds) ProtoMessage() {}

func (x *EquipmentIds) ProtoReflect() protoreflect.Message {
	mi := &file_tours_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EquipmentIds.ProtoReflect.Descriptor instead.
func (*EquipmentIds) Descriptor() ([]byte, []int) {
	return file_tours_service_proto_rawDescGZIP(), []int{0}
}

func (x *EquipmentIds) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EquipmentIds) GetEquipmentIds() []string {
	if x != nil {
		return x.EquipmentIds
	}
	return nil
}

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tours_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_tours_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_tours_service_proto_rawDescGZIP(), []int{1}
}

func (x *Page) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Page) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type EquipmentId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *EquipmentId) Reset() {
	*x = EquipmentId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tours_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EquipmentId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EquipmentId) ProtoMessage() {}

func (x *EquipmentId) ProtoReflect() protoreflect.Message {
	mi := &file_tours_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EquipmentId.ProtoReflect.Descriptor instead.
func (*EquipmentId) Descriptor() ([]byte, []int) {
	return file_tours_service_proto_rawDescGZIP(), []int{2}
}

func (x *EquipmentId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateEquipmentId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Equipment *EquipmentResponse `protobuf:"bytes,1,opt,name=equipment,proto3" json:"equipment,omitempty"`
	Id        string             `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateEquipmentId) Reset() {
	*x = UpdateEquipmentId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tours_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEquipmentId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEquipmentId) ProtoMessage() {}

func (x *UpdateEquipmentId) ProtoReflect() protoreflect.Message {
	mi := &file_tours_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEquipmentId.ProtoReflect.Descriptor instead.
func (*UpdateEquipmentId) Descriptor() ([]byte, []int) {
	return file_tours_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateEquipmentId) GetEquipment() *EquipmentResponse {
	if x != nil {
		return x.Equipment
	}
	return nil
}

func (x *UpdateEquipmentId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EquipmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          *string `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description *string `protobuf:"bytes,3,opt,name=description,proto3,oneof" json:"description,omitempty"`
}

func (x *EquipmentResponse) Reset() {
	*x = EquipmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tours_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EquipmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EquipmentResponse) ProtoMessage() {}

func (x *EquipmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tours_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EquipmentResponse.ProtoReflect.Descriptor instead.
func (*EquipmentResponse) Descriptor() ([]byte, []int) {
	return file_tours_service_proto_rawDescGZIP(), []int{4}
}

func (x *EquipmentResponse) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *EquipmentResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EquipmentResponse) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

type EquipmentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EquipmentResponse []*EquipmentResponse `protobuf:"bytes,1,rep,name=EquipmentResponse,proto3" json:"EquipmentResponse,omitempty"`
}

func (x *EquipmentsResponse) Reset() {
	*x = EquipmentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tours_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EquipmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EquipmentsResponse) ProtoMessage() {}

func (x *EquipmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tours_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EquipmentsResponse.ProtoReflect.Descriptor instead.
func (*EquipmentsResponse) Descriptor() ([]byte, []int) {
	return file_tours_service_proto_rawDescGZIP(), []int{5}
}

func (x *EquipmentsResponse) GetEquipmentResponse() []*EquipmentResponse {
	if x != nil {
		return x.EquipmentResponse
	}
	return nil
}

type PagedEquipmentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results     []*EquipmentResponse `protobuf:"bytes,1,rep,name=Results,proto3" json:"Results,omitempty"`
	TotalCounts int32                `protobuf:"varint,2,opt,name=TotalCounts,proto3" json:"TotalCounts,omitempty"`
}

func (x *PagedEquipmentsResponse) Reset() {
	*x = PagedEquipmentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tours_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagedEquipmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagedEquipmentsResponse) ProtoMessage() {}

func (x *PagedEquipmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tours_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagedEquipmentsResponse.ProtoReflect.Descriptor instead.
func (*PagedEquipmentsResponse) Descriptor() ([]byte, []int) {
	return file_tours_service_proto_rawDescGZIP(), []int{6}
}

func (x *PagedEquipmentsResponse) GetResults() []*EquipmentResponse {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *PagedEquipmentsResponse) GetTotalCounts() int32 {
	if x != nil {
		return x.TotalCounts
	}
	return 0
}

var File_tours_service_proto protoreflect.FileDescriptor

var file_tours_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x74, 0x6f, 0x75, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x0c, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x71, 0x75,
	0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x22, 0x36, 0x0a, 0x04, 0x50, 0x61, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x22, 0x1d, 0x0a, 0x0b, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x55, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x09, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x09, 0x65, 0x71,
	0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x7a, 0x0a, 0x11, 0x45, 0x71, 0x75, 0x69, 0x70,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x13, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03,
	0x5f, 0x69, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x56, 0x0a, 0x12, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x11, 0x45, 0x71, 0x75,
	0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x11, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x69, 0x0a, 0x17, 0x50,
	0x61, 0x67, 0x65, 0x64, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x32, 0xe0, 0x02, 0x0a, 0x04, 0x54, 0x6f, 0x75, 0x72, 0x12,
	0x3d, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x45,
	0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0d, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x1a, 0x13, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x05, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x1a, 0x18, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x64,
	0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x45, 0x71, 0x75, 0x69, 0x70,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0c, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x1a, 0x12, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x2e, 0x45, 0x71,
	0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a,
	0x12, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45,
	0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x1a, 0x12, 0x2e, 0x45,
	0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x35, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x71, 0x75, 0x69,
	0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0c, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x1a, 0x12, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x75, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tours_service_proto_rawDescOnce sync.Once
	file_tours_service_proto_rawDescData = file_tours_service_proto_rawDesc
)

func file_tours_service_proto_rawDescGZIP() []byte {
	file_tours_service_proto_rawDescOnce.Do(func() {
		file_tours_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_tours_service_proto_rawDescData)
	})
	return file_tours_service_proto_rawDescData
}

var file_tours_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_tours_service_proto_goTypes = []interface{}{
	(*EquipmentIds)(nil),            // 0: EquipmentIds
	(*Page)(nil),                    // 1: Page
	(*EquipmentId)(nil),             // 2: EquipmentId
	(*UpdateEquipmentId)(nil),       // 3: UpdateEquipmentId
	(*EquipmentResponse)(nil),       // 4: EquipmentResponse
	(*EquipmentsResponse)(nil),      // 5: EquipmentsResponse
	(*PagedEquipmentsResponse)(nil), // 6: PagedEquipmentsResponse
}
var file_tours_service_proto_depIdxs = []int32{
	4, // 0: UpdateEquipmentId.equipment:type_name -> EquipmentResponse
	4, // 1: EquipmentsResponse.EquipmentResponse:type_name -> EquipmentResponse
	4, // 2: PagedEquipmentsResponse.Results:type_name -> EquipmentResponse
	0, // 3: Tour.GetAvailableEquipment:input_type -> EquipmentIds
	1, // 4: Tour.GetAllEquipment:input_type -> Page
	2, // 5: Tour.GetEquipment:input_type -> EquipmentId
	4, // 6: Tour.CreateEquipment:input_type -> EquipmentResponse
	3, // 7: Tour.UpdateEquipment:input_type -> UpdateEquipmentId
	2, // 8: Tour.DeleteEquipment:input_type -> EquipmentId
	5, // 9: Tour.GetAvailableEquipment:output_type -> EquipmentsResponse
	6, // 10: Tour.GetAllEquipment:output_type -> PagedEquipmentsResponse
	4, // 11: Tour.GetEquipment:output_type -> EquipmentResponse
	4, // 12: Tour.CreateEquipment:output_type -> EquipmentResponse
	4, // 13: Tour.UpdateEquipment:output_type -> EquipmentResponse
	4, // 14: Tour.DeleteEquipment:output_type -> EquipmentResponse
	9, // [9:15] is the sub-list for method output_type
	3, // [3:9] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tours_service_proto_init() }
func file_tours_service_proto_init() {
	if File_tours_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tours_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EquipmentIds); i {
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
		file_tours_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
		file_tours_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EquipmentId); i {
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
		file_tours_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEquipmentId); i {
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
		file_tours_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EquipmentResponse); i {
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
		file_tours_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EquipmentsResponse); i {
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
		file_tours_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PagedEquipmentsResponse); i {
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
	file_tours_service_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tours_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tours_service_proto_goTypes,
		DependencyIndexes: file_tours_service_proto_depIdxs,
		MessageInfos:      file_tours_service_proto_msgTypes,
	}.Build()
	File_tours_service_proto = out.File
	file_tours_service_proto_rawDesc = nil
	file_tours_service_proto_goTypes = nil
	file_tours_service_proto_depIdxs = nil
}