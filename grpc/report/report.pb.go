// Copyright 2021 ICT@MOPH.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.8
// source: src/report.proto

package grpc

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ReportGetLastServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cid string `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
}

func (x *ReportGetLastServiceRequest) Reset() {
	*x = ReportGetLastServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_report_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportGetLastServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportGetLastServiceRequest) ProtoMessage() {}

func (x *ReportGetLastServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_report_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportGetLastServiceRequest.ProtoReflect.Descriptor instead.
func (*ReportGetLastServiceRequest) Descriptor() ([]byte, []int) {
	return file_src_report_proto_rawDescGZIP(), []int{0}
}

func (x *ReportGetLastServiceRequest) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

type ReportGetLastServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*ReportGetLastServiceResponse_Results `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *ReportGetLastServiceResponse) Reset() {
	*x = ReportGetLastServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_report_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportGetLastServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportGetLastServiceResponse) ProtoMessage() {}

func (x *ReportGetLastServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_report_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportGetLastServiceResponse.ProtoReflect.Descriptor instead.
func (*ReportGetLastServiceResponse) Descriptor() ([]byte, []int) {
	return file_src_report_proto_rawDescGZIP(), []int{1}
}

func (x *ReportGetLastServiceResponse) GetResults() []*ReportGetLastServiceResponse_Results {
	if x != nil {
		return x.Results
	}
	return nil
}

type ReportGetLastServiceResponse_Results struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hospname     string `protobuf:"bytes,1,opt,name=hospname,proto3" json:"hospname,omitempty"`
	Hospcode     string `protobuf:"bytes,2,opt,name=hospcode,proto3" json:"hospcode,omitempty"`
	Hn           string `protobuf:"bytes,3,opt,name=hn,proto3" json:"hn,omitempty"`
	Vn           string `protobuf:"bytes,4,opt,name=vn,proto3" json:"vn,omitempty"`
	An           string `protobuf:"bytes,5,opt,name=an,proto3" json:"an,omitempty"`
	VisitDate    string `protobuf:"bytes,6,opt,name=visit_date,json=visitDate,proto3" json:"visit_date,omitempty"`
	VisitTime    string `protobuf:"bytes,7,opt,name=visit_time,json=visitTime,proto3" json:"visit_time,omitempty"`
	Pttype       string `protobuf:"bytes,8,opt,name=pttype,proto3" json:"pttype,omitempty"`
	PttypeNo     string `protobuf:"bytes,9,opt,name=pttype_no,json=pttypeNo,proto3" json:"pttype_no,omitempty"`
	PttypeName   string `protobuf:"bytes,10,opt,name=pttype_name,json=pttypeName,proto3" json:"pttype_name,omitempty"`
	Diagcode     string `protobuf:"bytes,11,opt,name=diagcode,proto3" json:"diagcode,omitempty"`
	Diagname     string `protobuf:"bytes,12,opt,name=diagname,proto3" json:"diagname,omitempty"`
	ProviderName string `protobuf:"bytes,13,opt,name=provider_name,json=providerName,proto3" json:"provider_name,omitempty"`
	HisProvider  string `protobuf:"bytes,14,opt,name=his_provider,json=hisProvider,proto3" json:"his_provider,omitempty"`
	Txid         string `protobuf:"bytes,15,opt,name=txid,proto3" json:"txid,omitempty"`
}

func (x *ReportGetLastServiceResponse_Results) Reset() {
	*x = ReportGetLastServiceResponse_Results{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_report_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportGetLastServiceResponse_Results) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportGetLastServiceResponse_Results) ProtoMessage() {}

func (x *ReportGetLastServiceResponse_Results) ProtoReflect() protoreflect.Message {
	mi := &file_src_report_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportGetLastServiceResponse_Results.ProtoReflect.Descriptor instead.
func (*ReportGetLastServiceResponse_Results) Descriptor() ([]byte, []int) {
	return file_src_report_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ReportGetLastServiceResponse_Results) GetHospname() string {
	if x != nil {
		return x.Hospname
	}
	return ""
}

func (x *ReportGetLastServiceResponse_Results) GetHospcode() string {
	if x != nil {
		return x.Hospcode
	}
	return ""
}

func (x *ReportGetLastServiceResponse_Results) GetHn() string {
	if x != nil {
		return x.Hn
	}
	return ""
}

func (x *ReportGetLastServiceResponse_Results) GetVn() string {
	if x != nil {
		return x.Vn
	}
	return ""
}

func (x *ReportGetLastServiceResponse_Results) GetAn() string {
	if x != nil {
		return x.An
	}
	return ""
}

func (x *ReportGetLastServiceResponse_Results) GetVisitDate() string {
	if x != nil {
		return x.VisitDate
	}
	return ""
}

func (x *ReportGetLastServiceResponse_Results) GetVisitTime() string {
	if x != nil {
		return x.VisitTime
	}
	return ""
}

func (x *ReportGetLastServiceResponse_Results) GetPttype() string {
	if x != nil {
		return x.Pttype
	}
	return ""
}

func (x *ReportGetLastServiceResponse_Results) GetPttypeNo() string {
	if x != nil {
		return x.PttypeNo
	}
	return ""
}

func (x *ReportGetLastServiceResponse_Results) GetPttypeName() string {
	if x != nil {
		return x.PttypeName
	}
	return ""
}

func (x *ReportGetLastServiceResponse_Results) GetDiagcode() string {
	if x != nil {
		return x.Diagcode
	}
	return ""
}

func (x *ReportGetLastServiceResponse_Results) GetDiagname() string {
	if x != nil {
		return x.Diagname
	}
	return ""
}

func (x *ReportGetLastServiceResponse_Results) GetProviderName() string {
	if x != nil {
		return x.ProviderName
	}
	return ""
}

func (x *ReportGetLastServiceResponse_Results) GetHisProvider() string {
	if x != nil {
		return x.HisProvider
	}
	return ""
}

func (x *ReportGetLastServiceResponse_Results) GetTxid() string {
	if x != nil {
		return x.Txid
	}
	return ""
}

var File_src_report_proto protoreflect.FileDescriptor

var file_src_report_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x72, 0x63, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x2f, 0x0a, 0x1b, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69, 0x64, 0x22, 0x80, 0x04, 0x0a, 0x1c, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x1a, 0x99, 0x03, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x68, 0x6f, 0x73, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x68, 0x6f, 0x73, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x70,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x70,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x68, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x68, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x76, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x76, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x61, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x69, 0x73, 0x69, 0x74, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x69, 0x73, 0x69, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x74, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x74, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x74,
	0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x74, 0x74, 0x79, 0x70, 0x65, 0x4e, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x74, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x74,
	0x74, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x61, 0x67,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x61, 0x67,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x69, 0x73,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x78, 0x69, 0x64,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x69, 0x64, 0x32, 0x70, 0x0a, 0x0d,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a,
	0x14, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x21, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f,
	0x5a, 0x0d, 0x2e, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_report_proto_rawDescOnce sync.Once
	file_src_report_proto_rawDescData = file_src_report_proto_rawDesc
)

func file_src_report_proto_rawDescGZIP() []byte {
	file_src_report_proto_rawDescOnce.Do(func() {
		file_src_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_report_proto_rawDescData)
	})
	return file_src_report_proto_rawDescData
}

var file_src_report_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_src_report_proto_goTypes = []interface{}{
	(*ReportGetLastServiceRequest)(nil),          // 0: grpc.ReportGetLastServiceRequest
	(*ReportGetLastServiceResponse)(nil),         // 1: grpc.ReportGetLastServiceResponse
	(*ReportGetLastServiceResponse_Results)(nil), // 2: grpc.ReportGetLastServiceResponse.Results
}
var file_src_report_proto_depIdxs = []int32{
	2, // 0: grpc.ReportGetLastServiceResponse.results:type_name -> grpc.ReportGetLastServiceResponse.Results
	0, // 1: grpc.ReportService.ReportGetLastService:input_type -> grpc.ReportGetLastServiceRequest
	1, // 2: grpc.ReportService.ReportGetLastService:output_type -> grpc.ReportGetLastServiceResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_src_report_proto_init() }
func file_src_report_proto_init() {
	if File_src_report_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_report_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportGetLastServiceRequest); i {
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
		file_src_report_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportGetLastServiceResponse); i {
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
		file_src_report_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportGetLastServiceResponse_Results); i {
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
			RawDescriptor: file_src_report_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_report_proto_goTypes,
		DependencyIndexes: file_src_report_proto_depIdxs,
		MessageInfos:      file_src_report_proto_msgTypes,
	}.Build()
	File_src_report_proto = out.File
	file_src_report_proto_rawDesc = nil
	file_src_report_proto_goTypes = nil
	file_src_report_proto_depIdxs = nil
}
