// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: proto/user.proto

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ROLE int32

const (
	ROLE_Admin     ROLE = 0
	ROLE_Editor    ROLE = 1
	ROLE_Moderator ROLE = 2
	ROLE_Seller    ROLE = 3
	ROLE_Customer  ROLE = 4
)

// Enum value maps for ROLE.
var (
	ROLE_name = map[int32]string{
		0: "Admin",
		1: "Editor",
		2: "Moderator",
		3: "Seller",
		4: "Customer",
	}
	ROLE_value = map[string]int32{
		"Admin":     0,
		"Editor":    1,
		"Moderator": 2,
		"Seller":    3,
		"Customer":  4,
	}
)

func (x ROLE) Enum() *ROLE {
	p := new(ROLE)
	*p = x
	return p
}

func (x ROLE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ROLE) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_user_proto_enumTypes[0].Descriptor()
}

func (ROLE) Type() protoreflect.EnumType {
	return &file_proto_user_proto_enumTypes[0]
}

func (x ROLE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ROLE.Descriptor instead.
func (ROLE) EnumDescriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{0}
}

type SignInUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNo  string `protobuf:"bytes,2,opt,name=phoneNo,proto3" json:"phoneNo,omitempty"`
	RoleType ROLE   `protobuf:"varint,4,opt,name=roleType,proto3,enum=proto.ROLE" json:"roleType,omitempty"`
}

func (x *SignInUser) Reset() {
	*x = SignInUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInUser) ProtoMessage() {}

func (x *SignInUser) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInUser.ProtoReflect.Descriptor instead.
func (*SignInUser) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{0}
}

func (x *SignInUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignInUser) GetPhoneNo() string {
	if x != nil {
		return x.PhoneNo
	}
	return ""
}

func (x *SignInUser) GetRoleType() ROLE {
	if x != nil {
		return x.RoleType
	}
	return ROLE_Admin
}

type UserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserId) Reset() {
	*x = UserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserId) ProtoMessage() {}

func (x *UserId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserId.ProtoReflect.Descriptor instead.
func (*UserId) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserId) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SignInAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsValid  bool   `protobuf:"varint,1,opt,name=isValid,proto3" json:"isValid,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SignInAck) Reset() {
	*x = SignInAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInAck) ProtoMessage() {}

func (x *SignInAck) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInAck.ProtoReflect.Descriptor instead.
func (*SignInAck) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{2}
}

func (x *SignInAck) GetIsValid() bool {
	if x != nil {
		return x.IsValid
	}
	return false
}

func (x *SignInAck) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	PhoneNo         string   `protobuf:"bytes,2,opt,name=phoneNo,proto3" json:"phoneNo,omitempty"`
	Email           string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password        string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	HomeAdress      *Address `protobuf:"bytes,5,opt,name=homeAdress,proto3" json:"homeAdress,omitempty"`
	DeliveryAddress *Address `protobuf:"bytes,6,opt,name=deliveryAddress,proto3" json:"deliveryAddress,omitempty"`
	RoleType        ROLE     `protobuf:"varint,7,opt,name=roleType,proto3,enum=proto.ROLE" json:"roleType,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[3]
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
	return file_proto_user_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetPhoneNo() string {
	if x != nil {
		return x.PhoneNo
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetHomeAdress() *Address {
	if x != nil {
		return x.HomeAdress
	}
	return nil
}

func (x *User) GetDeliveryAddress() *Address {
	if x != nil {
		return x.DeliveryAddress
	}
	return nil
}

func (x *User) GetRoleType() ROLE {
	if x != nil {
		return x.RoleType
	}
	return ROLE_Admin
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNo     string `protobuf:"bytes,1,opt,name=phoneNo,proto3" json:"phoneNo,omitempty"`
	AddressLine string `protobuf:"bytes,2,opt,name=addressLine,proto3" json:"addressLine,omitempty"`
	City        string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	PinCode     string `protobuf:"bytes,4,opt,name=pinCode,proto3" json:"pinCode,omitempty"`
	Landmark    string `protobuf:"bytes,5,opt,name=landmark,proto3" json:"landmark,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{4}
}

func (x *Address) GetPhoneNo() string {
	if x != nil {
		return x.PhoneNo
	}
	return ""
}

func (x *Address) GetAddressLine() string {
	if x != nil {
		return x.AddressLine
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetPinCode() string {
	if x != nil {
		return x.PinCode
	}
	return ""
}

func (x *Address) GetLandmark() string {
	if x != nil {
		return x.Landmark
	}
	return ""
}

var File_proto_user_proto protoreflect.FileDescriptor

var file_proto_user_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x0a, 0x53, 0x69, 0x67,
	0x6e, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x6f, 0x12, 0x27, 0x0a, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x4f, 0x4c, 0x45, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x18, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x41, 0x0a, 0x09, 0x53, 0x69,
	0x67, 0x6e, 0x49, 0x6e, 0x41, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xf9, 0x01,
	0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x2e, 0x0a, 0x0a, 0x68, 0x6f, 0x6d, 0x65, 0x41, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0a, 0x68, 0x6f, 0x6d, 0x65,
	0x41, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x38, 0x0a, 0x0f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x0f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x27, 0x0a, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x4f, 0x4c, 0x45, 0x52,
	0x08, 0x72, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x07, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x6f, 0x12,
	0x20, 0x0a, 0x0b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x6e, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x6e,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x69, 0x6e, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x64, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x64, 0x6d, 0x61, 0x72, 0x6b, 0x2a, 0x46, 0x0a, 0x04, 0x52,
	0x4f, 0x4c, 0x45, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x6f,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x65, 0x6c,
	0x6c, 0x65, 0x72, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x10, 0x04, 0x32, 0x6c, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x2a, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x41, 0x63, 0x6b, 0x22,
	0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_user_proto_rawDescOnce sync.Once
	file_proto_user_proto_rawDescData = file_proto_user_proto_rawDesc
)

func file_proto_user_proto_rawDescGZIP() []byte {
	file_proto_user_proto_rawDescOnce.Do(func() {
		file_proto_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_user_proto_rawDescData)
	})
	return file_proto_user_proto_rawDescData
}

var file_proto_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_user_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_user_proto_goTypes = []interface{}{
	(ROLE)(0),          // 0: proto.ROLE
	(*SignInUser)(nil), // 1: proto.SignInUser
	(*UserId)(nil),     // 2: proto.UserId
	(*SignInAck)(nil),  // 3: proto.SignInAck
	(*User)(nil),       // 4: proto.User
	(*Address)(nil),    // 5: proto.Address
}
var file_proto_user_proto_depIdxs = []int32{
	0, // 0: proto.SignInUser.roleType:type_name -> proto.ROLE
	5, // 1: proto.User.homeAdress:type_name -> proto.Address
	5, // 2: proto.User.deliveryAddress:type_name -> proto.Address
	0, // 3: proto.User.roleType:type_name -> proto.ROLE
	4, // 4: proto.AuthUser.CreateUser:input_type -> proto.User
	1, // 5: proto.AuthUser.GetUserByID:input_type -> proto.SignInUser
	2, // 6: proto.AuthUser.CreateUser:output_type -> proto.UserId
	3, // 7: proto.AuthUser.GetUserByID:output_type -> proto.SignInAck
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_user_proto_init() }
func file_proto_user_proto_init() {
	if File_proto_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInUser); i {
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
		file_proto_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserId); i {
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
		file_proto_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInAck); i {
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
		file_proto_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
			RawDescriptor: file_proto_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_user_proto_goTypes,
		DependencyIndexes: file_proto_user_proto_depIdxs,
		EnumInfos:         file_proto_user_proto_enumTypes,
		MessageInfos:      file_proto_user_proto_msgTypes,
	}.Build()
	File_proto_user_proto = out.File
	file_proto_user_proto_rawDesc = nil
	file_proto_user_proto_goTypes = nil
	file_proto_user_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthUserClient is the client API for AuthUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthUserClient interface {
	CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserId, error)
	GetUserByID(ctx context.Context, in *SignInUser, opts ...grpc.CallOption) (*SignInAck, error)
}

type authUserClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthUserClient(cc grpc.ClientConnInterface) AuthUserClient {
	return &authUserClient{cc}
}

func (c *authUserClient) CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserId, error) {
	out := new(UserId)
	err := c.cc.Invoke(ctx, "/proto.AuthUser/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authUserClient) GetUserByID(ctx context.Context, in *SignInUser, opts ...grpc.CallOption) (*SignInAck, error) {
	out := new(SignInAck)
	err := c.cc.Invoke(ctx, "/proto.AuthUser/GetUserByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthUserServer is the server API for AuthUser service.
type AuthUserServer interface {
	CreateUser(context.Context, *User) (*UserId, error)
	GetUserByID(context.Context, *SignInUser) (*SignInAck, error)
}

// UnimplementedAuthUserServer can be embedded to have forward compatible implementations.
type UnimplementedAuthUserServer struct {
}

func (*UnimplementedAuthUserServer) CreateUser(context.Context, *User) (*UserId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedAuthUserServer) GetUserByID(context.Context, *SignInUser) (*SignInAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByID not implemented")
}

func RegisterAuthUserServer(s *grpc.Server, srv AuthUserServer) {
	s.RegisterService(&_AuthUser_serviceDesc, srv)
}

func _AuthUser_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthUserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthUser/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthUserServer).CreateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthUser_GetUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthUserServer).GetUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthUser/GetUserByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthUserServer).GetUserByID(ctx, req.(*SignInUser))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthUser_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AuthUser",
	HandlerType: (*AuthUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _AuthUser_CreateUser_Handler,
		},
		{
			MethodName: "GetUserByID",
			Handler:    _AuthUser_GetUserByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user.proto",
}
