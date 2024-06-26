// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.12.4
// source: pb/product_service/product.proto

package product_service

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type ProductSortType int32

const (
	ProductSortType_NAME_DESC       ProductSortType = 0
	ProductSortType_NAME_ASC        ProductSortType = 1
	ProductSortType_CREATED_AT_DESC ProductSortType = 2
	ProductSortType_CREATED_AT_ASC  ProductSortType = 3
)

// Enum value maps for ProductSortType.
var (
	ProductSortType_name = map[int32]string{
		0: "NAME_DESC",
		1: "NAME_ASC",
		2: "CREATED_AT_DESC",
		3: "CREATED_AT_ASC",
	}
	ProductSortType_value = map[string]int32{
		"NAME_DESC":       0,
		"NAME_ASC":        1,
		"CREATED_AT_DESC": 2,
		"CREATED_AT_ASC":  3,
	}
)

func (x ProductSortType) Enum() *ProductSortType {
	p := new(ProductSortType)
	*p = x
	return p
}

func (x ProductSortType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProductSortType) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_product_service_product_proto_enumTypes[0].Descriptor()
}

func (ProductSortType) Type() protoreflect.EnumType {
	return &file_pb_product_service_product_proto_enumTypes[0]
}

func (x ProductSortType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProductSortType.Descriptor instead.
func (ProductSortType) EnumDescriptor() ([]byte, []int) {
	return file_pb_product_service_product_proto_rawDescGZIP(), []int{0}
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name        string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Price       float64              `protobuf:"fixed64,3,opt,name=price,proto3" json:"price"`
	Stock       int64                `protobuf:"varint,4,opt,name=stock,proto3" json:"stock"`
	Description string               `protobuf:"bytes,5,opt,name=description,proto3" json:"description"`
	ImageUrl    string               `protobuf:"bytes,6,opt,name=image_url,json=imageUrl,proto3" json:"image_url"`
	CreatedAt   *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt   *timestamp.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	DeletedAt   *timestamp.Timestamp `protobuf:"bytes,9,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_product_service_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_pb_product_service_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_pb_product_service_product_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Product) GetStock() int64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Product) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Product) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Product) GetDeletedAt() *timestamp.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type Products struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products"`
}

func (x *Products) Reset() {
	*x = Products{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_product_service_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Products) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Products) ProtoMessage() {}

func (x *Products) ProtoReflect() protoreflect.Message {
	mi := &file_pb_product_service_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Products.ProtoReflect.Descriptor instead.
func (*Products) Descriptor() ([]byte, []int) {
	return file_pb_product_service_product_proto_rawDescGZIP(), []int{1}
}

func (x *Products) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

// ProductSearchRequest :nodoc:
type ProductSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size     int64           `protobuf:"varint,1,opt,name=size,proto3" json:"size"`
	Page     int64           `protobuf:"varint,2,opt,name=page,proto3" json:"page"`
	Query    string          `protobuf:"bytes,3,opt,name=query,proto3" json:"query"`
	Filter   *ProductFilter  `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter"`
	SortType ProductSortType `protobuf:"varint,5,opt,name=sort_type,json=sortType,proto3,enum=pb.product_service.ProductSortType" json:"sort_type"`
}

func (x *ProductSearchRequest) Reset() {
	*x = ProductSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_product_service_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductSearchRequest) ProtoMessage() {}

func (x *ProductSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_product_service_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductSearchRequest.ProtoReflect.Descriptor instead.
func (*ProductSearchRequest) Descriptor() ([]byte, []int) {
	return file_pb_product_service_product_proto_rawDescGZIP(), []int{2}
}

func (x *ProductSearchRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ProductSearchRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ProductSearchRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *ProductSearchRequest) GetFilter() *ProductFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *ProductSearchRequest) GetSortType() ProductSortType {
	if x != nil {
		return x.SortType
	}
	return ProductSortType_NAME_DESC
}

type ProductFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsDeleted bool `protobuf:"varint,1,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted"`
}

func (x *ProductFilter) Reset() {
	*x = ProductFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_product_service_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductFilter) ProtoMessage() {}

func (x *ProductFilter) ProtoReflect() protoreflect.Message {
	mi := &file_pb_product_service_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductFilter.ProtoReflect.Descriptor instead.
func (*ProductFilter) Descriptor() ([]byte, []int) {
	return file_pb_product_service_product_proto_rawDescGZIP(), []int{3}
}

func (x *ProductFilter) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

var File_pb_product_service_product_proto protoreflect.FileDescriptor

var file_pb_product_service_product_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55,
	0x72, 0x6c, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x43, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12,
	0x37, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0xd1, 0x01, 0x0a, 0x14, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x39, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x09, 0x73, 0x6f,
	0x72, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e,
	0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x6f, 0x72, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x08, 0x73, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x2e, 0x0a, 0x0d,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x2a, 0x57, 0x0a, 0x0f,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0d, 0x0a, 0x09, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x10, 0x00, 0x12, 0x0c,
	0x0a, 0x08, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x41, 0x53, 0x43, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x41, 0x54, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x10,
	0x02, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x41, 0x54, 0x5f,
	0x41, 0x53, 0x43, 0x10, 0x03, 0x42, 0x14, 0x5a, 0x12, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pb_product_service_product_proto_rawDescOnce sync.Once
	file_pb_product_service_product_proto_rawDescData = file_pb_product_service_product_proto_rawDesc
)

func file_pb_product_service_product_proto_rawDescGZIP() []byte {
	file_pb_product_service_product_proto_rawDescOnce.Do(func() {
		file_pb_product_service_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_product_service_product_proto_rawDescData)
	})
	return file_pb_product_service_product_proto_rawDescData
}

var file_pb_product_service_product_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_product_service_product_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_product_service_product_proto_goTypes = []interface{}{
	(ProductSortType)(0),         // 0: pb.product_service.ProductSortType
	(*Product)(nil),              // 1: pb.product_service.Product
	(*Products)(nil),             // 2: pb.product_service.Products
	(*ProductSearchRequest)(nil), // 3: pb.product_service.ProductSearchRequest
	(*ProductFilter)(nil),        // 4: pb.product_service.ProductFilter
	(*timestamp.Timestamp)(nil),  // 5: google.protobuf.Timestamp
}
var file_pb_product_service_product_proto_depIdxs = []int32{
	5, // 0: pb.product_service.Product.created_at:type_name -> google.protobuf.Timestamp
	5, // 1: pb.product_service.Product.updated_at:type_name -> google.protobuf.Timestamp
	5, // 2: pb.product_service.Product.deleted_at:type_name -> google.protobuf.Timestamp
	1, // 3: pb.product_service.Products.products:type_name -> pb.product_service.Product
	4, // 4: pb.product_service.ProductSearchRequest.filter:type_name -> pb.product_service.ProductFilter
	0, // 5: pb.product_service.ProductSearchRequest.sort_type:type_name -> pb.product_service.ProductSortType
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pb_product_service_product_proto_init() }
func file_pb_product_service_product_proto_init() {
	if File_pb_product_service_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_product_service_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_pb_product_service_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Products); i {
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
		file_pb_product_service_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductSearchRequest); i {
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
		file_pb_product_service_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductFilter); i {
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
			RawDescriptor: file_pb_product_service_product_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_product_service_product_proto_goTypes,
		DependencyIndexes: file_pb_product_service_product_proto_depIdxs,
		EnumInfos:         file_pb_product_service_product_proto_enumTypes,
		MessageInfos:      file_pb_product_service_product_proto_msgTypes,
	}.Build()
	File_pb_product_service_product_proto = out.File
	file_pb_product_service_product_proto_rawDesc = nil
	file_pb_product_service_product_proto_goTypes = nil
	file_pb_product_service_product_proto_depIdxs = nil
}
