// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: buildings.proto

package buildings

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

type SendMailToCourierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic       string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Text        string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	FileUrl     string `protobuf:"bytes,3,opt,name=file_url,json=fileUrl,proto3" json:"file_url,omitempty"`
	CountryCode string `protobuf:"bytes,4,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	City        string `protobuf:"bytes,5,opt,name=city,proto3" json:"city,omitempty"`
	Limit       int32  `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset      int32  `protobuf:"varint,7,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *SendMailToCourierRequest) Reset() {
	*x = SendMailToCourierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buildings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMailToCourierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMailToCourierRequest) ProtoMessage() {}

func (x *SendMailToCourierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buildings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMailToCourierRequest.ProtoReflect.Descriptor instead.
func (*SendMailToCourierRequest) Descriptor() ([]byte, []int) {
	return file_buildings_proto_rawDescGZIP(), []int{0}
}

func (x *SendMailToCourierRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *SendMailToCourierRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *SendMailToCourierRequest) GetFileUrl() string {
	if x != nil {
		return x.FileUrl
	}
	return ""
}

func (x *SendMailToCourierRequest) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *SendMailToCourierRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *SendMailToCourierRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *SendMailToCourierRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetMailToCourierResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MailCouriers []*GetMailToCourierResponse_MailCourier `protobuf:"bytes,1,rep,name=mailCouriers,proto3" json:"mailCouriers,omitempty"`
}

func (x *GetMailToCourierResponse) Reset() {
	*x = GetMailToCourierResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buildings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMailToCourierResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMailToCourierResponse) ProtoMessage() {}

func (x *GetMailToCourierResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buildings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMailToCourierResponse.ProtoReflect.Descriptor instead.
func (*GetMailToCourierResponse) Descriptor() ([]byte, []int) {
	return file_buildings_proto_rawDescGZIP(), []int{1}
}

func (x *GetMailToCourierResponse) GetMailCouriers() []*GetMailToCourierResponse_MailCourier {
	if x != nil {
		return x.MailCouriers
	}
	return nil
}

type GetMailToCourierResponse_MailCourier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Topic     string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Text      string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	UserCount int32  `protobuf:"varint,4,opt,name=user_count,json=userCount,proto3" json:"user_count,omitempty"`
	Dt        string `protobuf:"bytes,5,opt,name=dt,proto3" json:"dt,omitempty"`
}

func (x *GetMailToCourierResponse_MailCourier) Reset() {
	*x = GetMailToCourierResponse_MailCourier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buildings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMailToCourierResponse_MailCourier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMailToCourierResponse_MailCourier) ProtoMessage() {}

func (x *GetMailToCourierResponse_MailCourier) ProtoReflect() protoreflect.Message {
	mi := &file_buildings_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMailToCourierResponse_MailCourier.ProtoReflect.Descriptor instead.
func (*GetMailToCourierResponse_MailCourier) Descriptor() ([]byte, []int) {
	return file_buildings_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetMailToCourierResponse_MailCourier) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetMailToCourierResponse_MailCourier) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *GetMailToCourierResponse_MailCourier) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *GetMailToCourierResponse_MailCourier) GetUserCount() int32 {
	if x != nil {
		return x.UserCount
	}
	return 0
}

func (x *GetMailToCourierResponse_MailCourier) GetDt() string {
	if x != nil {
		return x.Dt
	}
	return ""
}

var File_buildings_proto protoreflect.FileDescriptor

var file_buildings_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x01, 0x0a, 0x18, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61,
	0x69, 0x6c, 0x54, 0x6f, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x66, 0x69, 0x6c, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0xe3, 0x01, 0x0a,
	0x18, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6c, 0x54, 0x6f, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x6d, 0x61, 0x69,
	0x6c, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6c, 0x54,
	0x6f, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x52, 0x0c, 0x6d, 0x61,
	0x69, 0x6c, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x73, 0x1a, 0x76, 0x0a, 0x0b, 0x4d, 0x61,
	0x69, 0x6c, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x64, 0x74, 0x32, 0xa6, 0x01, 0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x4e, 0x0a, 0x11,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x54, 0x6f, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65,
	0x72, 0x12, 0x1f, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61,
	0x69, 0x6c, 0x54, 0x6f, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6c, 0x54, 0x6f, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6c, 0x54, 0x6f, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_buildings_proto_rawDescOnce sync.Once
	file_buildings_proto_rawDescData = file_buildings_proto_rawDesc
)

func file_buildings_proto_rawDescGZIP() []byte {
	file_buildings_proto_rawDescOnce.Do(func() {
		file_buildings_proto_rawDescData = protoimpl.X.CompressGZIP(file_buildings_proto_rawDescData)
	})
	return file_buildings_proto_rawDescData
}

var file_buildings_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_buildings_proto_goTypes = []interface{}{
	(*SendMailToCourierRequest)(nil),             // 0: admin.SendMailToCourierRequest
	(*GetMailToCourierResponse)(nil),             // 1: admin.GetMailToCourierResponse
	(*GetMailToCourierResponse_MailCourier)(nil), // 2: admin.GetMailToCourierResponse.MailCourier
	(*emptypb.Empty)(nil),                        // 3: google.protobuf.Empty
}
var file_buildings_proto_depIdxs = []int32{
	2, // 0: admin.GetMailToCourierResponse.mailCouriers:type_name -> admin.GetMailToCourierResponse.MailCourier
	0, // 1: admin.Admin.SendMailToCourier:input_type -> admin.SendMailToCourierRequest
	3, // 2: admin.Admin.GetMailToCourier:input_type -> google.protobuf.Empty
	3, // 3: admin.Admin.SendMailToCourier:output_type -> google.protobuf.Empty
	1, // 4: admin.Admin.GetMailToCourier:output_type -> admin.GetMailToCourierResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_buildings_proto_init() }
func file_buildings_proto_init() {
	if File_buildings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buildings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMailToCourierRequest); i {
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
		file_buildings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMailToCourierResponse); i {
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
		file_buildings_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMailToCourierResponse_MailCourier); i {
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
			RawDescriptor: file_buildings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_buildings_proto_goTypes,
		DependencyIndexes: file_buildings_proto_depIdxs,
		MessageInfos:      file_buildings_proto_msgTypes,
	}.Build()
	File_buildings_proto = out.File
	file_buildings_proto_rawDesc = nil
	file_buildings_proto_goTypes = nil
	file_buildings_proto_depIdxs = nil
}