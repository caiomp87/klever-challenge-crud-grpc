// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: crypto.proto

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_crypto_proto_rawDescGZIP(), []int{0}
}

type Crypto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Crypto) Reset() {
	*x = Crypto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Crypto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Crypto) ProtoMessage() {}

func (x *Crypto) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Crypto.ProtoReflect.Descriptor instead.
func (*Crypto) Descriptor() ([]byte, []int) {
	return file_crypto_proto_rawDescGZIP(), []int{1}
}

func (x *Crypto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CryptoResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CryptoResult) Reset() {
	*x = CryptoResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptoResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptoResult) ProtoMessage() {}

func (x *CryptoResult) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptoResult.ProtoReflect.Descriptor instead.
func (*CryptoResult) Descriptor() ([]byte, []int) {
	return file_crypto_proto_rawDescGZIP(), []int{2}
}

func (x *CryptoResult) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CryptoResult) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_crypto_proto protoreflect.FileDescriptor

var file_crypto_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x6b, 0x6c, 0x65, 0x76, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x22,
	0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1c, 0x0a, 0x06, 0x43, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x0c, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xa3, 0x01, 0x0a, 0x0d, 0x43,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0c,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x12, 0x17, 0x2e, 0x6b,
	0x6c, 0x65, 0x76, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x2e, 0x43,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x1a, 0x1d, 0x2e, 0x6b, 0x6c, 0x65, 0x76, 0x65, 0x72, 0x43, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x73, 0x12, 0x16, 0x2e, 0x6b, 0x6c, 0x65, 0x76, 0x65, 0x72, 0x43, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e,
	0x6b, 0x6c, 0x65, 0x76, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x2e,
	0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x30, 0x01,
	0x42, 0x13, 0x5a, 0x11, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crypto_proto_rawDescOnce sync.Once
	file_crypto_proto_rawDescData = file_crypto_proto_rawDesc
)

func file_crypto_proto_rawDescGZIP() []byte {
	file_crypto_proto_rawDescOnce.Do(func() {
		file_crypto_proto_rawDescData = protoimpl.X.CompressGZIP(file_crypto_proto_rawDescData)
	})
	return file_crypto_proto_rawDescData
}

var file_crypto_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_crypto_proto_goTypes = []interface{}{
	(*Empty)(nil),        // 0: kleverChallenge.Empty
	(*Crypto)(nil),       // 1: kleverChallenge.Crypto
	(*CryptoResult)(nil), // 2: kleverChallenge.CryptoResult
}
var file_crypto_proto_depIdxs = []int32{
	1, // 0: kleverChallenge.CryptoService.CreateCrypto:input_type -> kleverChallenge.Crypto
	0, // 1: kleverChallenge.CryptoService.ListCryptos:input_type -> kleverChallenge.Empty
	2, // 2: kleverChallenge.CryptoService.CreateCrypto:output_type -> kleverChallenge.CryptoResult
	2, // 3: kleverChallenge.CryptoService.ListCryptos:output_type -> kleverChallenge.CryptoResult
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_crypto_proto_init() }
func file_crypto_proto_init() {
	if File_crypto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crypto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_crypto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Crypto); i {
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
		file_crypto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoResult); i {
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
			RawDescriptor: file_crypto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_crypto_proto_goTypes,
		DependencyIndexes: file_crypto_proto_depIdxs,
		MessageInfos:      file_crypto_proto_msgTypes,
	}.Build()
	File_crypto_proto = out.File
	file_crypto_proto_rawDesc = nil
	file_crypto_proto_goTypes = nil
	file_crypto_proto_depIdxs = nil
}
