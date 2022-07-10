// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: proto/quotes.proto

package proto

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

type QuoteData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quote  string `protobuf:"bytes,1,opt,name=quote,proto3" json:"quote,omitempty"`
	Author string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *QuoteData) Reset() {
	*x = QuoteData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quotes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuoteData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuoteData) ProtoMessage() {}

func (x *QuoteData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quotes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuoteData.ProtoReflect.Descriptor instead.
func (*QuoteData) Descriptor() ([]byte, []int) {
	return file_proto_quotes_proto_rawDescGZIP(), []int{0}
}

func (x *QuoteData) GetQuote() string {
	if x != nil {
		return x.Quote
	}
	return ""
}

func (x *QuoteData) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type AddQuoteReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quote *QuoteData `protobuf:"bytes,1,opt,name=quote,proto3" json:"quote,omitempty"`
}

func (x *AddQuoteReply) Reset() {
	*x = AddQuoteReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quotes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddQuoteReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddQuoteReply) ProtoMessage() {}

func (x *AddQuoteReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quotes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddQuoteReply.ProtoReflect.Descriptor instead.
func (*AddQuoteReply) Descriptor() ([]byte, []int) {
	return file_proto_quotes_proto_rawDescGZIP(), []int{1}
}

func (x *AddQuoteReply) GetQuote() *QuoteData {
	if x != nil {
		return x.Quote
	}
	return nil
}

type ListQuotesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quotes []*QuoteData `protobuf:"bytes,1,rep,name=quotes,proto3" json:"quotes,omitempty"`
}

func (x *ListQuotesReply) Reset() {
	*x = ListQuotesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quotes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQuotesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQuotesReply) ProtoMessage() {}

func (x *ListQuotesReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quotes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQuotesReply.ProtoReflect.Descriptor instead.
func (*ListQuotesReply) Descriptor() ([]byte, []int) {
	return file_proto_quotes_proto_rawDescGZIP(), []int{2}
}

func (x *ListQuotesReply) GetQuotes() []*QuoteData {
	if x != nil {
		return x.Quotes
	}
	return nil
}

type ListQuotesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListQuotesRequest) Reset() {
	*x = ListQuotesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quotes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQuotesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQuotesRequest) ProtoMessage() {}

func (x *ListQuotesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quotes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQuotesRequest.ProtoReflect.Descriptor instead.
func (*ListQuotesRequest) Descriptor() ([]byte, []int) {
	return file_proto_quotes_proto_rawDescGZIP(), []int{3}
}

var File_proto_quotes_proto protoreflect.FileDescriptor

var file_proto_quotes_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x39, 0x0a, 0x09,
	0x51, 0x75, 0x6f, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x6f,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x38, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x51, 0x75,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73,
	0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x22, 0x3c, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x29, 0x0a, 0x06, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x51, 0x75,
	0x6f, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x22,
	0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x32, 0x84, 0x01, 0x0a, 0x06, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x12,
	0x42, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x19, 0x2e,
	0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65,
	0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12,
	0x11, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x1a, 0x15, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_quotes_proto_rawDescOnce sync.Once
	file_proto_quotes_proto_rawDescData = file_proto_quotes_proto_rawDesc
)

func file_proto_quotes_proto_rawDescGZIP() []byte {
	file_proto_quotes_proto_rawDescOnce.Do(func() {
		file_proto_quotes_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_quotes_proto_rawDescData)
	})
	return file_proto_quotes_proto_rawDescData
}

var file_proto_quotes_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_quotes_proto_goTypes = []interface{}{
	(*QuoteData)(nil),         // 0: quotes.QuoteData
	(*AddQuoteReply)(nil),     // 1: quotes.AddQuoteReply
	(*ListQuotesReply)(nil),   // 2: quotes.ListQuotesReply
	(*ListQuotesRequest)(nil), // 3: quotes.ListQuotesRequest
}
var file_proto_quotes_proto_depIdxs = []int32{
	0, // 0: quotes.AddQuoteReply.quote:type_name -> quotes.QuoteData
	0, // 1: quotes.ListQuotesReply.quotes:type_name -> quotes.QuoteData
	3, // 2: quotes.Quotes.ListQuotes:input_type -> quotes.ListQuotesRequest
	0, // 3: quotes.Quotes.AddQuote:input_type -> quotes.QuoteData
	2, // 4: quotes.Quotes.ListQuotes:output_type -> quotes.ListQuotesReply
	1, // 5: quotes.Quotes.AddQuote:output_type -> quotes.AddQuoteReply
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_quotes_proto_init() }
func file_proto_quotes_proto_init() {
	if File_proto_quotes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_quotes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuoteData); i {
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
		file_proto_quotes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddQuoteReply); i {
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
		file_proto_quotes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListQuotesReply); i {
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
		file_proto_quotes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListQuotesRequest); i {
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
			RawDescriptor: file_proto_quotes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_quotes_proto_goTypes,
		DependencyIndexes: file_proto_quotes_proto_depIdxs,
		MessageInfos:      file_proto_quotes_proto_msgTypes,
	}.Build()
	File_proto_quotes_proto = out.File
	file_proto_quotes_proto_rawDesc = nil
	file_proto_quotes_proto_goTypes = nil
	file_proto_quotes_proto_depIdxs = nil
}
