// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: trading.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Stock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code            string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name            string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	TotalStockCount uint32 `protobuf:"varint,4,opt,name=total_stock_count,json=totalStockCount,proto3" json:"total_stock_count,omitempty"`
}

func (x *Stock) Reset() {
	*x = Stock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trading_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stock) ProtoMessage() {}

func (x *Stock) ProtoReflect() protoreflect.Message {
	mi := &file_trading_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stock.ProtoReflect.Descriptor instead.
func (*Stock) Descriptor() ([]byte, []int) {
	return file_trading_proto_rawDescGZIP(), []int{0}
}

func (x *Stock) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Stock) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Stock) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Stock) GetTotalStockCount() uint32 {
	if x != nil {
		return x.TotalStockCount
	}
	return 0
}

type StockListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StockList []*Stock `protobuf:"bytes,1,rep,name=stock_list,json=stockList,proto3" json:"stock_list,omitempty"`
}

func (x *StockListResp) Reset() {
	*x = StockListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trading_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockListResp) ProtoMessage() {}

func (x *StockListResp) ProtoReflect() protoreflect.Message {
	mi := &file_trading_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockListResp.ProtoReflect.Descriptor instead.
func (*StockListResp) Descriptor() ([]byte, []int) {
	return file_trading_proto_rawDescGZIP(), []int{1}
}

func (x *StockListResp) GetStockList() []*Stock {
	if x != nil {
		return x.StockList
	}
	return nil
}

var File_trading_proto protoreflect.FileDescriptor

var file_trading_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x3e, 0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x2d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e,
	0x67, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x4c, 0x69,
	0x73, 0x74, 0x32, 0x4b, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x40, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x2e,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42,
	0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x68,
	0x69, 0x6c, 0x62, 0x75, 0x74, 0x2f, 0x66, 0x69, 0x6e, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trading_proto_rawDescOnce sync.Once
	file_trading_proto_rawDescData = file_trading_proto_rawDesc
)

func file_trading_proto_rawDescGZIP() []byte {
	file_trading_proto_rawDescOnce.Do(func() {
		file_trading_proto_rawDescData = protoimpl.X.CompressGZIP(file_trading_proto_rawDescData)
	})
	return file_trading_proto_rawDescData
}

var file_trading_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_trading_proto_goTypes = []interface{}{
	(*Stock)(nil),         // 0: trading.Stock
	(*StockListResp)(nil), // 1: trading.StockListResp
	(*emptypb.Empty)(nil), // 2: google.protobuf.Empty
}
var file_trading_proto_depIdxs = []int32{
	0, // 0: trading.StockListResp.stock_list:type_name -> trading.Stock
	2, // 1: trading.Trading.GetStockList:input_type -> google.protobuf.Empty
	1, // 2: trading.Trading.GetStockList:output_type -> trading.StockListResp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_trading_proto_init() }
func file_trading_proto_init() {
	if File_trading_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trading_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stock); i {
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
		file_trading_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockListResp); i {
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
			RawDescriptor: file_trading_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_trading_proto_goTypes,
		DependencyIndexes: file_trading_proto_depIdxs,
		MessageInfos:      file_trading_proto_msgTypes,
	}.Build()
	File_trading_proto = out.File
	file_trading_proto_rawDesc = nil
	file_trading_proto_goTypes = nil
	file_trading_proto_depIdxs = nil
}
