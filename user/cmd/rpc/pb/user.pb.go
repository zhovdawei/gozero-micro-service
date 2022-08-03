// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: pb/user.proto

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

// 查询User
type QueryUserByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *QueryUserByIdReq) Reset() {
	*x = QueryUserByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryUserByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryUserByIdReq) ProtoMessage() {}

func (x *QueryUserByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryUserByIdReq.ProtoReflect.Descriptor instead.
func (*QueryUserByIdReq) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{0}
}

func (x *QueryUserByIdReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// 查询User响应结果
type UserResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       int64  `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	IdCard       string `protobuf:"bytes,3,opt,name=IdCard,proto3" json:"IdCard,omitempty"`
	Email        string `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	Phone        string `protobuf:"bytes,5,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Password     string `protobuf:"bytes,6,opt,name=Password,proto3" json:"Password,omitempty"`
	Gender       int32  `protobuf:"varint,7,opt,name=Gender,proto3" json:"Gender,omitempty"`
	Birthday     int64  `protobuf:"varint,8,opt,name=Birthday,proto3" json:"Birthday,omitempty"`
	ProvinceCode int32  `protobuf:"varint,9,opt,name=ProvinceCode,proto3" json:"ProvinceCode,omitempty"`
	Province     string `protobuf:"bytes,10,opt,name=Province,proto3" json:"Province,omitempty"`
	CityCode     int32  `protobuf:"varint,11,opt,name=CityCode,proto3" json:"CityCode,omitempty"`
	City         string `protobuf:"bytes,12,opt,name=City,proto3" json:"City,omitempty"`
	AreaCode     int32  `protobuf:"varint,13,opt,name=AreaCode,proto3" json:"AreaCode,omitempty"`
	Area         string `protobuf:"bytes,14,opt,name=Area,proto3" json:"Area,omitempty"`
	CreateAt     int64  `protobuf:"varint,15,opt,name=CreateAt,proto3" json:"CreateAt,omitempty"`
	UpdateAt     int64  `protobuf:"varint,16,opt,name=UpdateAt,proto3" json:"UpdateAt,omitempty"`
}

func (x *UserResp) Reset() {
	*x = UserResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResp) ProtoMessage() {}

func (x *UserResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResp.ProtoReflect.Descriptor instead.
func (*UserResp) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserResp) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserResp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserResp) GetIdCard() string {
	if x != nil {
		return x.IdCard
	}
	return ""
}

func (x *UserResp) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserResp) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserResp) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserResp) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *UserResp) GetBirthday() int64 {
	if x != nil {
		return x.Birthday
	}
	return 0
}

func (x *UserResp) GetProvinceCode() int32 {
	if x != nil {
		return x.ProvinceCode
	}
	return 0
}

func (x *UserResp) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *UserResp) GetCityCode() int32 {
	if x != nil {
		return x.CityCode
	}
	return 0
}

func (x *UserResp) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *UserResp) GetAreaCode() int32 {
	if x != nil {
		return x.AreaCode
	}
	return 0
}

func (x *UserResp) GetArea() string {
	if x != nil {
		return x.Area
	}
	return ""
}

func (x *UserResp) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *UserResp) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

// 查询UserPost
type QueryUserPostByUserIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *QueryUserPostByUserIdReq) Reset() {
	*x = QueryUserPostByUserIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryUserPostByUserIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryUserPostByUserIdReq) ProtoMessage() {}

func (x *QueryUserPostByUserIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryUserPostByUserIdReq.ProtoReflect.Descriptor instead.
func (*QueryUserPostByUserIdReq) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{2}
}

func (x *QueryUserPostByUserIdReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// 查询UserPost数组
type UserPostArrayResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserPosts []*UserPostResp `protobuf:"bytes,1,rep,name=userPosts,proto3" json:"userPosts,omitempty"`
}

func (x *UserPostArrayResp) Reset() {
	*x = UserPostArrayResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPostArrayResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPostArrayResp) ProtoMessage() {}

func (x *UserPostArrayResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPostArrayResp.ProtoReflect.Descriptor instead.
func (*UserPostArrayResp) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserPostArrayResp) GetUserPosts() []*UserPostResp {
	if x != nil {
		return x.UserPosts
	}
	return nil
}

// 查询单条UserPost数据
type QueryUserPostReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId int64 `protobuf:"varint,1,opt,name=PostId,proto3" json:"PostId,omitempty"`
}

func (x *QueryUserPostReq) Reset() {
	*x = QueryUserPostReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryUserPostReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryUserPostReq) ProtoMessage() {}

func (x *QueryUserPostReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryUserPostReq.ProtoReflect.Descriptor instead.
func (*QueryUserPostReq) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{4}
}

func (x *QueryUserPostReq) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

// UserPost模型
type UserPostResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId       int64  `protobuf:"varint,1,opt,name=PostId,proto3" json:"PostId,omitempty"`
	UserId       int64  `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	UserName     string `protobuf:"bytes,3,opt,name=UserName,proto3" json:"UserName,omitempty"`
	UserPhone    string `protobuf:"bytes,4,opt,name=UserPhone,proto3" json:"UserPhone,omitempty"`
	ProvinceCode int32  `protobuf:"varint,5,opt,name=ProvinceCode,proto3" json:"ProvinceCode,omitempty"`
	Province     string `protobuf:"bytes,6,opt,name=Province,proto3" json:"Province,omitempty"`
	CityCode     int32  `protobuf:"varint,7,opt,name=CityCode,proto3" json:"CityCode,omitempty"`
	City         string `protobuf:"bytes,8,opt,name=City,proto3" json:"City,omitempty"`
	AreaCode     int32  `protobuf:"varint,9,opt,name=AreaCode,proto3" json:"AreaCode,omitempty"`
	Area         string `protobuf:"bytes,10,opt,name=Area,proto3" json:"Area,omitempty"`
	Address      string `protobuf:"bytes,11,opt,name=Address,proto3" json:"Address,omitempty"`
	Status       int32  `protobuf:"varint,12,opt,name=Status,proto3" json:"Status,omitempty"`
	CreateAt     int64  `protobuf:"varint,13,opt,name=CreateAt,proto3" json:"CreateAt,omitempty"`
	UpdateAt     int64  `protobuf:"varint,14,opt,name=UpdateAt,proto3" json:"UpdateAt,omitempty"`
}

func (x *UserPostResp) Reset() {
	*x = UserPostResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPostResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPostResp) ProtoMessage() {}

func (x *UserPostResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPostResp.ProtoReflect.Descriptor instead.
func (*UserPostResp) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{5}
}

func (x *UserPostResp) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *UserPostResp) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserPostResp) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserPostResp) GetUserPhone() string {
	if x != nil {
		return x.UserPhone
	}
	return ""
}

func (x *UserPostResp) GetProvinceCode() int32 {
	if x != nil {
		return x.ProvinceCode
	}
	return 0
}

func (x *UserPostResp) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *UserPostResp) GetCityCode() int32 {
	if x != nil {
		return x.CityCode
	}
	return 0
}

func (x *UserPostResp) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *UserPostResp) GetAreaCode() int32 {
	if x != nil {
		return x.AreaCode
	}
	return 0
}

func (x *UserPostResp) GetArea() string {
	if x != nil {
		return x.Area
	}
	return ""
}

func (x *UserPostResp) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UserPostResp) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UserPostResp) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *UserPostResp) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

var File_pb_user_proto protoreflect.FileDescriptor

var file_pb_user_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0x2a, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0xa2, 0x03, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x64, 0x43, 0x61,
	0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x49, 0x64, 0x43, 0x61, 0x72, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x22, 0x0a, 0x0c,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x43, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x43, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08,
	0x41, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x41, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x72, 0x65, 0x61,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x41, 0x72, 0x65, 0x61, 0x12, 0x1a, 0x0a, 0x08,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x22, 0x32, 0x0a, 0x18, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x6f, 0x73, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x2a, 0x0a,
	0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0x82, 0x03, 0x0a, 0x0c, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f,
	0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x50, 0x6f, 0x73, 0x74,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x43, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x43, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x41, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x41, 0x72, 0x65, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x41, 0x72, 0x65, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x32, 0xbb,
	0x01, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x49, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x1c,
	0x2e, 0x70, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73,
	0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x70,
	0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x37, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x6f, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_user_proto_rawDescOnce sync.Once
	file_pb_user_proto_rawDescData = file_pb_user_proto_rawDesc
)

func file_pb_user_proto_rawDescGZIP() []byte {
	file_pb_user_proto_rawDescOnce.Do(func() {
		file_pb_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_user_proto_rawDescData)
	})
	return file_pb_user_proto_rawDescData
}

var file_pb_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pb_user_proto_goTypes = []interface{}{
	(*QueryUserByIdReq)(nil),         // 0: pb.QueryUserByIdReq
	(*UserResp)(nil),                 // 1: pb.UserResp
	(*QueryUserPostByUserIdReq)(nil), // 2: pb.QueryUserPostByUserIdReq
	(*UserPostArrayResp)(nil),        // 3: pb.UserPostArrayResp
	(*QueryUserPostReq)(nil),         // 4: pb.QueryUserPostReq
	(*UserPostResp)(nil),             // 5: pb.UserPostResp
}
var file_pb_user_proto_depIdxs = []int32{
	5, // 0: pb.UserPostArrayResp.userPosts:type_name -> pb.UserPostResp
	0, // 1: pb.user.QueryUser:input_type -> pb.QueryUserByIdReq
	2, // 2: pb.user.QueryUserPostArray:input_type -> pb.QueryUserPostByUserIdReq
	4, // 3: pb.user.QueryUserPost:input_type -> pb.QueryUserPostReq
	1, // 4: pb.user.QueryUser:output_type -> pb.UserResp
	3, // 5: pb.user.QueryUserPostArray:output_type -> pb.UserPostArrayResp
	5, // 6: pb.user.QueryUserPost:output_type -> pb.UserPostResp
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_user_proto_init() }
func file_pb_user_proto_init() {
	if File_pb_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryUserByIdReq); i {
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
		file_pb_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResp); i {
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
		file_pb_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryUserPostByUserIdReq); i {
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
		file_pb_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPostArrayResp); i {
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
		file_pb_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryUserPostReq); i {
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
		file_pb_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPostResp); i {
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
			RawDescriptor: file_pb_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_user_proto_goTypes,
		DependencyIndexes: file_pb_user_proto_depIdxs,
		MessageInfos:      file_pb_user_proto_msgTypes,
	}.Build()
	File_pb_user_proto = out.File
	file_pb_user_proto_rawDesc = nil
	file_pb_user_proto_goTypes = nil
	file_pb_user_proto_depIdxs = nil
}
