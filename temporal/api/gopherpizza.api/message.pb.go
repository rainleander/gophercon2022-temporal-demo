// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: proto/pizza/v1/message.proto

package gopherpizza_api

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

type Crust int32

const (
	Crust_CRUST_NORMAL      Crust = 0
	Crust_CRUST_THIN        Crust = 1
	Crust_CRUST_GARLIC      Crust = 2
	Crust_CRUST_STUFFED     Crust = 3
	Crust_CRUST_GLUTEN_FREE Crust = 4
	Crust_CRUST_PRETZEL     Crust = 5
)

// Enum value maps for Crust.
var (
	Crust_name = map[int32]string{
		0: "CRUST_NORMAL",
		1: "CRUST_THIN",
		2: "CRUST_GARLIC",
		3: "CRUST_STUFFED",
		4: "CRUST_GLUTEN_FREE",
		5: "CRUST_PRETZEL",
	}
	Crust_value = map[string]int32{
		"CRUST_NORMAL":      0,
		"CRUST_THIN":        1,
		"CRUST_GARLIC":      2,
		"CRUST_STUFFED":     3,
		"CRUST_GLUTEN_FREE": 4,
		"CRUST_PRETZEL":     5,
	}
)

func (x Crust) Enum() *Crust {
	p := new(Crust)
	*p = x
	return p
}

func (x Crust) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Crust) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_pizza_v1_message_proto_enumTypes[0].Descriptor()
}

func (Crust) Type() protoreflect.EnumType {
	return &file_proto_pizza_v1_message_proto_enumTypes[0]
}

func (x Crust) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Crust.Descriptor instead.
func (Crust) EnumDescriptor() ([]byte, []int) {
	return file_proto_pizza_v1_message_proto_rawDescGZIP(), []int{0}
}

type Cheese int32

const (
	Cheese_CHEESE_MOZZARELLA Cheese = 0
	Cheese_CHEESE_CHEDDAR    Cheese = 1
	Cheese_CHEESE_NONE       Cheese = 2
	Cheese_CHEESE_ALL        Cheese = 3
)

// Enum value maps for Cheese.
var (
	Cheese_name = map[int32]string{
		0: "CHEESE_MOZZARELLA",
		1: "CHEESE_CHEDDAR",
		2: "CHEESE_NONE",
		3: "CHEESE_ALL",
	}
	Cheese_value = map[string]int32{
		"CHEESE_MOZZARELLA": 0,
		"CHEESE_CHEDDAR":    1,
		"CHEESE_NONE":       2,
		"CHEESE_ALL":        3,
	}
)

func (x Cheese) Enum() *Cheese {
	p := new(Cheese)
	*p = x
	return p
}

func (x Cheese) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Cheese) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_pizza_v1_message_proto_enumTypes[1].Descriptor()
}

func (Cheese) Type() protoreflect.EnumType {
	return &file_proto_pizza_v1_message_proto_enumTypes[1]
}

func (x Cheese) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Cheese.Descriptor instead.
func (Cheese) EnumDescriptor() ([]byte, []int) {
	return file_proto_pizza_v1_message_proto_rawDescGZIP(), []int{1}
}

type OrderStatus int32

const (
	OrderStatus_ORDER_RECEIVED         OrderStatus = 0
	OrderStatus_ORDER_PREPARING        OrderStatus = 1
	OrderStatus_ORDER_BAKING           OrderStatus = 2
	OrderStatus_ORDER_PENDING_PICKUP   OrderStatus = 3
	OrderStatus_ORDER_OUT_FOR_DELIVERY OrderStatus = 4
	OrderStatus_ORDER_DELIVERED        OrderStatus = 5
)

// Enum value maps for OrderStatus.
var (
	OrderStatus_name = map[int32]string{
		0: "ORDER_RECEIVED",
		1: "ORDER_PREPARING",
		2: "ORDER_BAKING",
		3: "ORDER_PENDING_PICKUP",
		4: "ORDER_OUT_FOR_DELIVERY",
		5: "ORDER_DELIVERED",
	}
	OrderStatus_value = map[string]int32{
		"ORDER_RECEIVED":         0,
		"ORDER_PREPARING":        1,
		"ORDER_BAKING":           2,
		"ORDER_PENDING_PICKUP":   3,
		"ORDER_OUT_FOR_DELIVERY": 4,
		"ORDER_DELIVERED":        5,
	}
)

func (x OrderStatus) Enum() *OrderStatus {
	p := new(OrderStatus)
	*p = x
	return p
}

func (x OrderStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_pizza_v1_message_proto_enumTypes[2].Descriptor()
}

func (OrderStatus) Type() protoreflect.EnumType {
	return &file_proto_pizza_v1_message_proto_enumTypes[2]
}

func (x OrderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderStatus.Descriptor instead.
func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_pizza_v1_message_proto_rawDescGZIP(), []int{2}
}

type PizzaOrderInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Crust    Crust    `protobuf:"varint,2,opt,name=crust,proto3,enum=gopherpizza.pizza.api.v1.Crust" json:"crust,omitempty"`
	Cheese   Cheese   `protobuf:"varint,3,opt,name=cheese,proto3,enum=gopherpizza.pizza.api.v1.Cheese" json:"cheese,omitempty"`
	Toppings []string `protobuf:"bytes,4,rep,name=toppings,proto3" json:"toppings,omitempty"`
}

func (x *PizzaOrderInfo) Reset() {
	*x = PizzaOrderInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pizza_v1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PizzaOrderInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PizzaOrderInfo) ProtoMessage() {}

func (x *PizzaOrderInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pizza_v1_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PizzaOrderInfo.ProtoReflect.Descriptor instead.
func (*PizzaOrderInfo) Descriptor() ([]byte, []int) {
	return file_proto_pizza_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *PizzaOrderInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PizzaOrderInfo) GetCrust() Crust {
	if x != nil {
		return x.Crust
	}
	return Crust_CRUST_NORMAL
}

func (x *PizzaOrderInfo) GetCheese() Cheese {
	if x != nil {
		return x.Cheese
	}
	return Cheese_CHEESE_MOZZARELLA
}

func (x *PizzaOrderInfo) GetToppings() []string {
	if x != nil {
		return x.Toppings
	}
	return nil
}

type PizzaOrderStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status OrderStatus     `protobuf:"varint,1,opt,name=status,proto3,enum=gopherpizza.pizza.api.v1.OrderStatus" json:"status,omitempty"`
	Order  *PizzaOrderInfo `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
	RunId  string          `protobuf:"bytes,3,opt,name=runId,proto3" json:"runId,omitempty"`
}

func (x *PizzaOrderStatus) Reset() {
	*x = PizzaOrderStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pizza_v1_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PizzaOrderStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PizzaOrderStatus) ProtoMessage() {}

func (x *PizzaOrderStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pizza_v1_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PizzaOrderStatus.ProtoReflect.Descriptor instead.
func (*PizzaOrderStatus) Descriptor() ([]byte, []int) {
	return file_proto_pizza_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *PizzaOrderStatus) GetStatus() OrderStatus {
	if x != nil {
		return x.Status
	}
	return OrderStatus_ORDER_RECEIVED
}

func (x *PizzaOrderStatus) GetOrder() *PizzaOrderInfo {
	if x != nil {
		return x.Order
	}
	return nil
}

func (x *PizzaOrderStatus) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

var File_proto_pizza_v1_message_proto protoreflect.FileDescriptor

var file_proto_pizza_v1_message_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x69, 0x7a, 0x7a, 0x61, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18,
	0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x70, 0x69, 0x7a, 0x7a, 0x61, 0x2e, 0x70, 0x69, 0x7a, 0x7a,
	0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x22, 0xad, 0x01, 0x0a, 0x0e, 0x50, 0x69, 0x7a,
	0x7a, 0x61, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x05, 0x63,
	0x72, 0x75, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x70,
	0x68, 0x65, 0x72, 0x70, 0x69, 0x7a, 0x7a, 0x61, 0x2e, 0x70, 0x69, 0x7a, 0x7a, 0x61, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x75, 0x73, 0x74, 0x52, 0x05, 0x63, 0x72, 0x75,
	0x73, 0x74, 0x12, 0x38, 0x0a, 0x06, 0x63, 0x68, 0x65, 0x65, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x70, 0x69, 0x7a, 0x7a, 0x61,
	0x2e, 0x70, 0x69, 0x7a, 0x7a, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68,
	0x65, 0x65, 0x73, 0x65, 0x52, 0x06, 0x63, 0x68, 0x65, 0x65, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x74, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x22, 0xa7, 0x01, 0x0a, 0x10, 0x50, 0x69, 0x7a,
	0x7a, 0x61, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3d, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e,
	0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x70, 0x69, 0x7a, 0x7a, 0x61, 0x2e, 0x70, 0x69, 0x7a, 0x7a,
	0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3e, 0x0a, 0x05,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f,
	0x70, 0x68, 0x65, 0x72, 0x70, 0x69, 0x7a, 0x7a, 0x61, 0x2e, 0x70, 0x69, 0x7a, 0x7a, 0x61, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x7a, 0x7a, 0x61, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x72, 0x75, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x75, 0x6e,
	0x49, 0x64, 0x2a, 0x78, 0x0a, 0x05, 0x43, 0x72, 0x75, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x0c, 0x43,
	0x52, 0x55, 0x53, 0x54, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x0e, 0x0a,
	0x0a, 0x43, 0x52, 0x55, 0x53, 0x54, 0x5f, 0x54, 0x48, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x10, 0x0a,
	0x0c, 0x43, 0x52, 0x55, 0x53, 0x54, 0x5f, 0x47, 0x41, 0x52, 0x4c, 0x49, 0x43, 0x10, 0x02, 0x12,
	0x11, 0x0a, 0x0d, 0x43, 0x52, 0x55, 0x53, 0x54, 0x5f, 0x53, 0x54, 0x55, 0x46, 0x46, 0x45, 0x44,
	0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x52, 0x55, 0x53, 0x54, 0x5f, 0x47, 0x4c, 0x55, 0x54,
	0x45, 0x4e, 0x5f, 0x46, 0x52, 0x45, 0x45, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x52, 0x55,
	0x53, 0x54, 0x5f, 0x50, 0x52, 0x45, 0x54, 0x5a, 0x45, 0x4c, 0x10, 0x05, 0x2a, 0x54, 0x0a, 0x06,
	0x43, 0x68, 0x65, 0x65, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x48, 0x45, 0x45, 0x53, 0x45,
	0x5f, 0x4d, 0x4f, 0x5a, 0x5a, 0x41, 0x52, 0x45, 0x4c, 0x4c, 0x41, 0x10, 0x00, 0x12, 0x12, 0x0a,
	0x0e, 0x43, 0x48, 0x45, 0x45, 0x53, 0x45, 0x5f, 0x43, 0x48, 0x45, 0x44, 0x44, 0x41, 0x52, 0x10,
	0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x48, 0x45, 0x45, 0x53, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x48, 0x45, 0x45, 0x53, 0x45, 0x5f, 0x41, 0x4c, 0x4c,
	0x10, 0x03, 0x2a, 0x93, 0x01, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x43, 0x45,
	0x49, 0x56, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f,
	0x50, 0x52, 0x45, 0x50, 0x41, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x4f,
	0x52, 0x44, 0x45, 0x52, 0x5f, 0x42, 0x41, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x18, 0x0a,
	0x14, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x50,
	0x49, 0x43, 0x4b, 0x55, 0x50, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x4f, 0x52, 0x44, 0x45, 0x52,
	0x5f, 0x4f, 0x55, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52,
	0x59, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x44, 0x45, 0x4c,
	0x49, 0x56, 0x45, 0x52, 0x45, 0x44, 0x10, 0x05, 0x42, 0x11, 0x5a, 0x0f, 0x67, 0x6f, 0x70, 0x68,
	0x65, 0x72, 0x70, 0x69, 0x7a, 0x7a, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_pizza_v1_message_proto_rawDescOnce sync.Once
	file_proto_pizza_v1_message_proto_rawDescData = file_proto_pizza_v1_message_proto_rawDesc
)

func file_proto_pizza_v1_message_proto_rawDescGZIP() []byte {
	file_proto_pizza_v1_message_proto_rawDescOnce.Do(func() {
		file_proto_pizza_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_pizza_v1_message_proto_rawDescData)
	})
	return file_proto_pizza_v1_message_proto_rawDescData
}

var file_proto_pizza_v1_message_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_proto_pizza_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_pizza_v1_message_proto_goTypes = []interface{}{
	(Crust)(0),               // 0: gopherpizza.pizza.api.v1.Crust
	(Cheese)(0),              // 1: gopherpizza.pizza.api.v1.Cheese
	(OrderStatus)(0),         // 2: gopherpizza.pizza.api.v1.OrderStatus
	(*PizzaOrderInfo)(nil),   // 3: gopherpizza.pizza.api.v1.PizzaOrderInfo
	(*PizzaOrderStatus)(nil), // 4: gopherpizza.pizza.api.v1.PizzaOrderStatus
}
var file_proto_pizza_v1_message_proto_depIdxs = []int32{
	0, // 0: gopherpizza.pizza.api.v1.PizzaOrderInfo.crust:type_name -> gopherpizza.pizza.api.v1.Crust
	1, // 1: gopherpizza.pizza.api.v1.PizzaOrderInfo.cheese:type_name -> gopherpizza.pizza.api.v1.Cheese
	2, // 2: gopherpizza.pizza.api.v1.PizzaOrderStatus.status:type_name -> gopherpizza.pizza.api.v1.OrderStatus
	3, // 3: gopherpizza.pizza.api.v1.PizzaOrderStatus.order:type_name -> gopherpizza.pizza.api.v1.PizzaOrderInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_pizza_v1_message_proto_init() }
func file_proto_pizza_v1_message_proto_init() {
	if File_proto_pizza_v1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_pizza_v1_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PizzaOrderInfo); i {
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
		file_proto_pizza_v1_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PizzaOrderStatus); i {
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
			RawDescriptor: file_proto_pizza_v1_message_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_pizza_v1_message_proto_goTypes,
		DependencyIndexes: file_proto_pizza_v1_message_proto_depIdxs,
		EnumInfos:         file_proto_pizza_v1_message_proto_enumTypes,
		MessageInfos:      file_proto_pizza_v1_message_proto_msgTypes,
	}.Build()
	File_proto_pizza_v1_message_proto = out.File
	file_proto_pizza_v1_message_proto_rawDesc = nil
	file_proto_pizza_v1_message_proto_goTypes = nil
	file_proto_pizza_v1_message_proto_depIdxs = nil
}
