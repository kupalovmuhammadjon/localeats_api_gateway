// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: dish.proto

package dish

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

type ReqCreateDish struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KitchenId   string   `protobuf:"bytes,1,opt,name=kitchen_id,json=kitchenId,proto3" json:"kitchen_id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price       float32  `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Category    string   `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	Ingredients []string `protobuf:"bytes,5,rep,name=ingredients,proto3" json:"ingredients,omitempty"`
	Description string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Available   bool     `protobuf:"varint,7,opt,name=available,proto3" json:"available,omitempty"`
}

func (x *ReqCreateDish) Reset() {
	*x = ReqCreateDish{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dish_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqCreateDish) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqCreateDish) ProtoMessage() {}

func (x *ReqCreateDish) ProtoReflect() protoreflect.Message {
	mi := &file_dish_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqCreateDish.ProtoReflect.Descriptor instead.
func (*ReqCreateDish) Descriptor() ([]byte, []int) {
	return file_dish_proto_rawDescGZIP(), []int{0}
}

func (x *ReqCreateDish) GetKitchenId() string {
	if x != nil {
		return x.KitchenId
	}
	return ""
}

func (x *ReqCreateDish) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReqCreateDish) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ReqCreateDish) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *ReqCreateDish) GetIngredients() []string {
	if x != nil {
		return x.Ingredients
	}
	return nil
}

func (x *ReqCreateDish) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ReqCreateDish) GetAvailable() bool {
	if x != nil {
		return x.Available
	}
	return false
}

type DishInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	KitchenId     string   `protobuf:"bytes,2,opt,name=kitchen_id,json=kitchenId,proto3" json:"kitchen_id,omitempty"`
	KitchenName   string   `protobuf:"bytes,3,opt,name=kitchen_name,json=kitchenName,proto3" json:"kitchen_name,omitempty"`
	Name          string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Price         float32  `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty"`
	Category      string   `protobuf:"bytes,6,opt,name=category,proto3" json:"category,omitempty"`
	Ingredients   []string `protobuf:"bytes,7,rep,name=ingredients,proto3" json:"ingredients,omitempty"`
	Description   string   `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	Available     bool     `protobuf:"varint,9,opt,name=available,proto3" json:"available,omitempty"`
	Allergens     []string `protobuf:"bytes,10,rep,name=allergens,proto3" json:"allergens,omitempty"`
	NutritionInfo string   `protobuf:"bytes,11,opt,name=nutrition_info,json=nutritionInfo,proto3" json:"nutrition_info,omitempty"`
	DietaryInfo   []string `protobuf:"bytes,12,rep,name=dietary_info,json=dietaryInfo,proto3" json:"dietary_info,omitempty"`
	CreatedAt     string   `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string   `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *DishInfo) Reset() {
	*x = DishInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dish_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DishInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DishInfo) ProtoMessage() {}

func (x *DishInfo) ProtoReflect() protoreflect.Message {
	mi := &file_dish_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DishInfo.ProtoReflect.Descriptor instead.
func (*DishInfo) Descriptor() ([]byte, []int) {
	return file_dish_proto_rawDescGZIP(), []int{1}
}

func (x *DishInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DishInfo) GetKitchenId() string {
	if x != nil {
		return x.KitchenId
	}
	return ""
}

func (x *DishInfo) GetKitchenName() string {
	if x != nil {
		return x.KitchenName
	}
	return ""
}

func (x *DishInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DishInfo) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *DishInfo) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *DishInfo) GetIngredients() []string {
	if x != nil {
		return x.Ingredients
	}
	return nil
}

func (x *DishInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *DishInfo) GetAvailable() bool {
	if x != nil {
		return x.Available
	}
	return false
}

func (x *DishInfo) GetAllergens() []string {
	if x != nil {
		return x.Allergens
	}
	return nil
}

func (x *DishInfo) GetNutritionInfo() string {
	if x != nil {
		return x.NutritionInfo
	}
	return ""
}

func (x *DishInfo) GetDietaryInfo() []string {
	if x != nil {
		return x.DietaryInfo
	}
	return nil
}

func (x *DishInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *DishInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type DishShortInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	KitchenId   string  `protobuf:"bytes,2,opt,name=kitchen_id,json=kitchenId,proto3" json:"kitchen_id,omitempty"`
	KitchenName string  `protobuf:"bytes,3,opt,name=kitchen_name,json=kitchenName,proto3" json:"kitchen_name,omitempty"`
	Price       float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Category    string  `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
	Available   bool    `protobuf:"varint,6,opt,name=available,proto3" json:"available,omitempty"`
}

func (x *DishShortInfo) Reset() {
	*x = DishShortInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dish_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DishShortInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DishShortInfo) ProtoMessage() {}

func (x *DishShortInfo) ProtoReflect() protoreflect.Message {
	mi := &file_dish_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DishShortInfo.ProtoReflect.Descriptor instead.
func (*DishShortInfo) Descriptor() ([]byte, []int) {
	return file_dish_proto_rawDescGZIP(), []int{2}
}

func (x *DishShortInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DishShortInfo) GetKitchenId() string {
	if x != nil {
		return x.KitchenId
	}
	return ""
}

func (x *DishShortInfo) GetKitchenName() string {
	if x != nil {
		return x.KitchenName
	}
	return ""
}

func (x *DishShortInfo) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *DishShortInfo) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *DishShortInfo) GetAvailable() bool {
	if x != nil {
		return x.Available
	}
	return false
}

type Dishes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dishes []*DishShortInfo `protobuf:"bytes,1,rep,name=dishes,proto3" json:"dishes,omitempty"`
	Total  int64            `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Page   int32            `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit  int32            `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *Dishes) Reset() {
	*x = Dishes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dish_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dishes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dishes) ProtoMessage() {}

func (x *Dishes) ProtoReflect() protoreflect.Message {
	mi := &file_dish_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dishes.ProtoReflect.Descriptor instead.
func (*Dishes) Descriptor() ([]byte, []int) {
	return file_dish_proto_rawDescGZIP(), []int{3}
}

func (x *Dishes) GetDishes() []*DishShortInfo {
	if x != nil {
		return x.Dishes
	}
	return nil
}

func (x *Dishes) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Dishes) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Dishes) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ReqUpdateDish struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price       float32  `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Category    string   `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	Ingredients []string `protobuf:"bytes,5,rep,name=ingredients,proto3" json:"ingredients,omitempty"`
	Description string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Available   bool     `protobuf:"varint,7,opt,name=available,proto3" json:"available,omitempty"`
}

func (x *ReqUpdateDish) Reset() {
	*x = ReqUpdateDish{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dish_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqUpdateDish) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqUpdateDish) ProtoMessage() {}

func (x *ReqUpdateDish) ProtoReflect() protoreflect.Message {
	mi := &file_dish_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqUpdateDish.ProtoReflect.Descriptor instead.
func (*ReqUpdateDish) Descriptor() ([]byte, []int) {
	return file_dish_proto_rawDescGZIP(), []int{4}
}

func (x *ReqUpdateDish) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReqUpdateDish) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReqUpdateDish) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ReqUpdateDish) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *ReqUpdateDish) GetIngredients() []string {
	if x != nil {
		return x.Ingredients
	}
	return nil
}

func (x *ReqUpdateDish) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ReqUpdateDish) GetAvailable() bool {
	if x != nil {
		return x.Available
	}
	return false
}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dish_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_dish_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_dish_proto_rawDescGZIP(), []int{5}
}

func (x *Id) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dish_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_dish_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_dish_proto_rawDescGZIP(), []int{6}
}

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Page  int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit int32  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dish_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_dish_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_dish_proto_rawDescGZIP(), []int{7}
}

func (x *Pagination) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Pagination) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Pagination) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type NutritionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Allergens     []string `protobuf:"bytes,1,rep,name=allergens,proto3" json:"allergens,omitempty"`
	Calories      int32    `protobuf:"varint,2,opt,name=calories,proto3" json:"calories,omitempty"`
	Protein       int32    `protobuf:"varint,3,opt,name=protein,proto3" json:"protein,omitempty"`
	Carbohydrates int32    `protobuf:"varint,4,opt,name=carbohydrates,proto3" json:"carbohydrates,omitempty"`
	Fat           int32    `protobuf:"varint,5,opt,name=fat,proto3" json:"fat,omitempty"`
	DietaryInfo   []string `protobuf:"bytes,6,rep,name=dietary_info,json=dietaryInfo,proto3" json:"dietary_info,omitempty"`
}

func (x *NutritionInfo) Reset() {
	*x = NutritionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dish_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NutritionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NutritionInfo) ProtoMessage() {}

func (x *NutritionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_dish_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NutritionInfo.ProtoReflect.Descriptor instead.
func (*NutritionInfo) Descriptor() ([]byte, []int) {
	return file_dish_proto_rawDescGZIP(), []int{8}
}

func (x *NutritionInfo) GetAllergens() []string {
	if x != nil {
		return x.Allergens
	}
	return nil
}

func (x *NutritionInfo) GetCalories() int32 {
	if x != nil {
		return x.Calories
	}
	return 0
}

func (x *NutritionInfo) GetProtein() int32 {
	if x != nil {
		return x.Protein
	}
	return 0
}

func (x *NutritionInfo) GetCarbohydrates() int32 {
	if x != nil {
		return x.Carbohydrates
	}
	return 0
}

func (x *NutritionInfo) GetFat() int32 {
	if x != nil {
		return x.Fat
	}
	return 0
}

func (x *NutritionInfo) GetDietaryInfo() []string {
	if x != nil {
		return x.DietaryInfo
	}
	return nil
}

var File_dish_proto protoreflect.FileDescriptor

var file_dish_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x69, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x64, 0x69,
	0x73, 0x68, 0x22, 0xd6, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x69, 0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65,
	0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x67,
	0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b,
	0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x22, 0xaa, 0x03, 0x0a, 0x08,
	0x44, 0x69, 0x73, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6b, 0x69, 0x74, 0x63,
	0x68, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6b, 0x69,
	0x74, 0x63, 0x68, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6b, 0x69, 0x74, 0x63, 0x68,
	0x65, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6b,
	0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x73, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x73,
	0x12, 0x25, 0x0a, 0x0e, 0x6e, 0x75, 0x74, 0x72, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x75, 0x74, 0x72, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x65, 0x74, 0x61,
	0x72, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x65, 0x74, 0x61, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xb1, 0x01, 0x0a, 0x0d, 0x44, 0x69, 0x73,
	0x68, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6b, 0x69,
	0x74, 0x63, 0x68, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6b, 0x69, 0x74,
	0x63, 0x68, 0x65, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x75, 0x0a, 0x06,
	0x44, 0x69, 0x73, 0x68, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x64, 0x69, 0x73, 0x68, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x69, 0x73, 0x68, 0x2e, 0x44, 0x69,
	0x73, 0x68, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x64, 0x69, 0x73,
	0x68, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0xc7, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x69, 0x73, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x69,
	0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x14, 0x0a,
	0x02, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x46, 0x0a, 0x0a, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0xbe, 0x01, 0x0a, 0x0d, 0x4e, 0x75, 0x74, 0x72, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x67, 0x65,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x67,
	0x65, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x6c, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x61, 0x72,
	0x62, 0x6f, 0x68, 0x79, 0x64, 0x72, 0x61, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x63, 0x61, 0x72, 0x62, 0x6f, 0x68, 0x79, 0x64, 0x72, 0x61, 0x74, 0x65, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x66, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x66, 0x61,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x65, 0x74, 0x61, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x65, 0x74, 0x61, 0x72, 0x79,
	0x49, 0x6e, 0x66, 0x6f, 0x32, 0xca, 0x02, 0x0a, 0x04, 0x44, 0x69, 0x73, 0x68, 0x12, 0x31, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x68, 0x12, 0x13, 0x2e, 0x64, 0x69,
	0x73, 0x68, 0x2e, 0x52, 0x65, 0x71, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x68,
	0x1a, 0x0e, 0x2e, 0x64, 0x69, 0x73, 0x68, 0x2e, 0x44, 0x69, 0x73, 0x68, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x31, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x68, 0x12, 0x13,
	0x2e, 0x64, 0x69, 0x73, 0x68, 0x2e, 0x52, 0x65, 0x71, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44,
	0x69, 0x73, 0x68, 0x1a, 0x0e, 0x2e, 0x64, 0x69, 0x73, 0x68, 0x2e, 0x44, 0x69, 0x73, 0x68, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x2b, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x68, 0x65, 0x73,
	0x12, 0x10, 0x2e, 0x64, 0x69, 0x73, 0x68, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x0c, 0x2e, 0x64, 0x69, 0x73, 0x68, 0x2e, 0x44, 0x69, 0x73, 0x68, 0x65, 0x73,
	0x12, 0x27, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x68, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x08, 0x2e, 0x64, 0x69, 0x73, 0x68, 0x2e, 0x49, 0x64, 0x1a, 0x0e, 0x2e, 0x64, 0x69, 0x73, 0x68,
	0x2e, 0x44, 0x69, 0x73, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x22, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x44, 0x69, 0x73, 0x68, 0x12, 0x08, 0x2e, 0x64, 0x69, 0x73, 0x68, 0x2e, 0x49,
	0x64, 0x1a, 0x0a, 0x2e, 0x64, 0x69, 0x73, 0x68, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x26, 0x0a,
	0x0e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x68, 0x49, 0x64, 0x12,
	0x08, 0x2e, 0x64, 0x69, 0x73, 0x68, 0x2e, 0x49, 0x64, 0x1a, 0x0a, 0x2e, 0x64, 0x69, 0x73, 0x68,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x3a, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e,
	0x75, 0x74, 0x72, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x13, 0x2e, 0x64,
	0x69, 0x73, 0x68, 0x2e, 0x4e, 0x75, 0x74, 0x72, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x1a, 0x0e, 0x2e, 0x64, 0x69, 0x73, 0x68, 0x2e, 0x44, 0x69, 0x73, 0x68, 0x49, 0x6e, 0x66,
	0x6f, 0x42, 0x0f, 0x5a, 0x0d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x69,
	0x73, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dish_proto_rawDescOnce sync.Once
	file_dish_proto_rawDescData = file_dish_proto_rawDesc
)

func file_dish_proto_rawDescGZIP() []byte {
	file_dish_proto_rawDescOnce.Do(func() {
		file_dish_proto_rawDescData = protoimpl.X.CompressGZIP(file_dish_proto_rawDescData)
	})
	return file_dish_proto_rawDescData
}

var file_dish_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_dish_proto_goTypes = []interface{}{
	(*ReqCreateDish)(nil), // 0: dish.ReqCreateDish
	(*DishInfo)(nil),      // 1: dish.DishInfo
	(*DishShortInfo)(nil), // 2: dish.DishShortInfo
	(*Dishes)(nil),        // 3: dish.Dishes
	(*ReqUpdateDish)(nil), // 4: dish.ReqUpdateDish
	(*Id)(nil),            // 5: dish.Id
	(*Void)(nil),          // 6: dish.Void
	(*Pagination)(nil),    // 7: dish.Pagination
	(*NutritionInfo)(nil), // 8: dish.NutritionInfo
}
var file_dish_proto_depIdxs = []int32{
	2, // 0: dish.Dishes.dishes:type_name -> dish.DishShortInfo
	0, // 1: dish.Dish.CreateDish:input_type -> dish.ReqCreateDish
	4, // 2: dish.Dish.UpdateDish:input_type -> dish.ReqUpdateDish
	7, // 3: dish.Dish.GetDishes:input_type -> dish.Pagination
	5, // 4: dish.Dish.GetDishById:input_type -> dish.Id
	5, // 5: dish.Dish.DeleteDish:input_type -> dish.Id
	5, // 6: dish.Dish.ValidateDishId:input_type -> dish.Id
	8, // 7: dish.Dish.UpdateNutritionInfo:input_type -> dish.NutritionInfo
	1, // 8: dish.Dish.CreateDish:output_type -> dish.DishInfo
	1, // 9: dish.Dish.UpdateDish:output_type -> dish.DishInfo
	3, // 10: dish.Dish.GetDishes:output_type -> dish.Dishes
	1, // 11: dish.Dish.GetDishById:output_type -> dish.DishInfo
	6, // 12: dish.Dish.DeleteDish:output_type -> dish.Void
	6, // 13: dish.Dish.ValidateDishId:output_type -> dish.Void
	1, // 14: dish.Dish.UpdateNutritionInfo:output_type -> dish.DishInfo
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_dish_proto_init() }
func file_dish_proto_init() {
	if File_dish_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dish_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqCreateDish); i {
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
		file_dish_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DishInfo); i {
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
		file_dish_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DishShortInfo); i {
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
		file_dish_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dishes); i {
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
		file_dish_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqUpdateDish); i {
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
		file_dish_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_dish_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
		file_dish_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
		file_dish_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NutritionInfo); i {
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
			RawDescriptor: file_dish_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dish_proto_goTypes,
		DependencyIndexes: file_dish_proto_depIdxs,
		MessageInfos:      file_dish_proto_msgTypes,
	}.Build()
	File_dish_proto = out.File
	file_dish_proto_rawDesc = nil
	file_dish_proto_goTypes = nil
	file_dish_proto_depIdxs = nil
}
