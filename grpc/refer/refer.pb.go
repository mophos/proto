// Copyright 2021 ICT@MOPH.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: src/refer.proto

package grpc

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

type ReferOutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hospcode string `protobuf:"bytes,1,opt,name=hospcode,proto3" json:"hospcode,omitempty"`
	Date     string `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *ReferOutRequest) Reset() {
	*x = ReferOutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_refer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReferOutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReferOutRequest) ProtoMessage() {}

func (x *ReferOutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_refer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReferOutRequest.ProtoReflect.Descriptor instead.
func (*ReferOutRequest) Descriptor() ([]byte, []int) {
	return file_src_refer_proto_rawDescGZIP(), []int{0}
}

func (x *ReferOutRequest) GetHospcode() string {
	if x != nil {
		return x.Hospcode
	}
	return ""
}

func (x *ReferOutRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type ReferVisitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hospcode string `protobuf:"bytes,1,opt,name=hospcode,proto3" json:"hospcode,omitempty"`
	Date     string `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *ReferVisitRequest) Reset() {
	*x = ReferVisitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_refer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReferVisitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReferVisitRequest) ProtoMessage() {}

func (x *ReferVisitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_refer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReferVisitRequest.ProtoReflect.Descriptor instead.
func (*ReferVisitRequest) Descriptor() ([]byte, []int) {
	return file_src_refer_proto_rawDescGZIP(), []int{1}
}

func (x *ReferVisitRequest) GetHospcode() string {
	if x != nil {
		return x.Hospcode
	}
	return ""
}

func (x *ReferVisitRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type ReferOutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*ReferOutResponse_Results `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *ReferOutResponse) Reset() {
	*x = ReferOutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_refer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReferOutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReferOutResponse) ProtoMessage() {}

func (x *ReferOutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_refer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReferOutResponse.ProtoReflect.Descriptor instead.
func (*ReferOutResponse) Descriptor() ([]byte, []int) {
	return file_src_refer_proto_rawDescGZIP(), []int{2}
}

func (x *ReferOutResponse) GetResults() []*ReferOutResponse_Results {
	if x != nil {
		return x.Results
	}
	return nil
}

type ReferVisitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*ReferVisitResponse_Results `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *ReferVisitResponse) Reset() {
	*x = ReferVisitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_refer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReferVisitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReferVisitResponse) ProtoMessage() {}

func (x *ReferVisitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_refer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReferVisitResponse.ProtoReflect.Descriptor instead.
func (*ReferVisitResponse) Descriptor() ([]byte, []int) {
	return file_src_refer_proto_rawDescGZIP(), []int{3}
}

func (x *ReferVisitResponse) GetResults() []*ReferVisitResponse_Results {
	if x != nil {
		return x.Results
	}
	return nil
}

type ReferOutResponse_Results struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hospcode       string `protobuf:"bytes,1,opt,name=hospcode,proto3" json:"hospcode,omitempty"`
	Cid            string `protobuf:"bytes,2,opt,name=cid,proto3" json:"cid,omitempty"`
	Hn             string `protobuf:"bytes,3,opt,name=hn,proto3" json:"hn,omitempty"`
	Fname          string `protobuf:"bytes,4,opt,name=fname,proto3" json:"fname,omitempty"`
	Lname          string `protobuf:"bytes,5,opt,name=lname,proto3" json:"lname,omitempty"`
	Birthday       string `protobuf:"bytes,6,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Sex            string `protobuf:"bytes,7,opt,name=sex,proto3" json:"sex,omitempty"`
	Vn             string `protobuf:"bytes,8,opt,name=vn,proto3" json:"vn,omitempty"`
	ReferDate      string `protobuf:"bytes,9,opt,name=refer_date,json=referDate,proto3" json:"refer_date,omitempty"`
	ReferTime      string `protobuf:"bytes,10,opt,name=refer_time,json=referTime,proto3" json:"refer_time,omitempty"`
	ReferNumber    string `protobuf:"bytes,11,opt,name=refer_number,json=referNumber,proto3" json:"refer_number,omitempty"`
	ReferPoint     string `protobuf:"bytes,12,opt,name=refer_point,json=referPoint,proto3" json:"refer_point,omitempty"`
	Pttype         string `protobuf:"bytes,13,opt,name=pttype,proto3" json:"pttype,omitempty"`
	PttypeName     string `protobuf:"bytes,14,opt,name=pttype_name,json=pttypeName,proto3" json:"pttype_name,omitempty"`
	Treatment      string `protobuf:"bytes,15,opt,name=treatment,proto3" json:"treatment,omitempty"`
	Diagcode       string `protobuf:"bytes,16,opt,name=diagcode,proto3" json:"diagcode,omitempty"`
	Diagname       string `protobuf:"bytes,17,opt,name=diagname,proto3" json:"diagname,omitempty"`
	PreDiagnosis   string `protobuf:"bytes,18,opt,name=pre_diagnosis,json=preDiagnosis,proto3" json:"pre_diagnosis,omitempty"`
	Pmh            string `protobuf:"bytes,19,opt,name=pmh,proto3" json:"pmh,omitempty"`
	Hpi            string `protobuf:"bytes,20,opt,name=hpi,proto3" json:"hpi,omitempty"`
	ReferHospcode  string `protobuf:"bytes,22,opt,name=refer_hospcode,json=referHospcode,proto3" json:"refer_hospcode,omitempty"`
	ReferHospname  string `protobuf:"bytes,23,opt,name=refer_hospname,json=referHospname,proto3" json:"refer_hospname,omitempty"`
	ExpireDate     string `protobuf:"bytes,24,opt,name=expire_date,json=expireDate,proto3" json:"expire_date,omitempty"`
	ReferTypeId    string `protobuf:"bytes,25,opt,name=refer_type_id,json=referTypeId,proto3" json:"refer_type_id,omitempty"`
	ReferTypeName  string `protobuf:"bytes,26,opt,name=refer_type_name,json=referTypeName,proto3" json:"refer_type_name,omitempty"`
	ReferCauseId   string `protobuf:"bytes,27,opt,name=refer_cause_id,json=referCauseId,proto3" json:"refer_cause_id,omitempty"`
	ReferCauseName string `protobuf:"bytes,28,opt,name=refer_cause_name,json=referCauseName,proto3" json:"refer_cause_name,omitempty"`
	RfrcsId        string `protobuf:"bytes,29,opt,name=rfrcs_id,json=rfrcsId,proto3" json:"rfrcs_id,omitempty"`
	RfrcsName      string `protobuf:"bytes,30,opt,name=rfrcs_name,json=rfrcsName,proto3" json:"rfrcs_name,omitempty"`
}

func (x *ReferOutResponse_Results) Reset() {
	*x = ReferOutResponse_Results{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_refer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReferOutResponse_Results) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReferOutResponse_Results) ProtoMessage() {}

func (x *ReferOutResponse_Results) ProtoReflect() protoreflect.Message {
	mi := &file_src_refer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReferOutResponse_Results.ProtoReflect.Descriptor instead.
func (*ReferOutResponse_Results) Descriptor() ([]byte, []int) {
	return file_src_refer_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ReferOutResponse_Results) GetHospcode() string {
	if x != nil {
		return x.Hospcode
	}
	return ""
}

func (x *ReferOutResponse_Results) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

func (x *ReferOutResponse_Results) GetHn() string {
	if x != nil {
		return x.Hn
	}
	return ""
}

func (x *ReferOutResponse_Results) GetFname() string {
	if x != nil {
		return x.Fname
	}
	return ""
}

func (x *ReferOutResponse_Results) GetLname() string {
	if x != nil {
		return x.Lname
	}
	return ""
}

func (x *ReferOutResponse_Results) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *ReferOutResponse_Results) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *ReferOutResponse_Results) GetVn() string {
	if x != nil {
		return x.Vn
	}
	return ""
}

func (x *ReferOutResponse_Results) GetReferDate() string {
	if x != nil {
		return x.ReferDate
	}
	return ""
}

func (x *ReferOutResponse_Results) GetReferTime() string {
	if x != nil {
		return x.ReferTime
	}
	return ""
}

func (x *ReferOutResponse_Results) GetReferNumber() string {
	if x != nil {
		return x.ReferNumber
	}
	return ""
}

func (x *ReferOutResponse_Results) GetReferPoint() string {
	if x != nil {
		return x.ReferPoint
	}
	return ""
}

func (x *ReferOutResponse_Results) GetPttype() string {
	if x != nil {
		return x.Pttype
	}
	return ""
}

func (x *ReferOutResponse_Results) GetPttypeName() string {
	if x != nil {
		return x.PttypeName
	}
	return ""
}

func (x *ReferOutResponse_Results) GetTreatment() string {
	if x != nil {
		return x.Treatment
	}
	return ""
}

func (x *ReferOutResponse_Results) GetDiagcode() string {
	if x != nil {
		return x.Diagcode
	}
	return ""
}

func (x *ReferOutResponse_Results) GetDiagname() string {
	if x != nil {
		return x.Diagname
	}
	return ""
}

func (x *ReferOutResponse_Results) GetPreDiagnosis() string {
	if x != nil {
		return x.PreDiagnosis
	}
	return ""
}

func (x *ReferOutResponse_Results) GetPmh() string {
	if x != nil {
		return x.Pmh
	}
	return ""
}

func (x *ReferOutResponse_Results) GetHpi() string {
	if x != nil {
		return x.Hpi
	}
	return ""
}

func (x *ReferOutResponse_Results) GetReferHospcode() string {
	if x != nil {
		return x.ReferHospcode
	}
	return ""
}

func (x *ReferOutResponse_Results) GetReferHospname() string {
	if x != nil {
		return x.ReferHospname
	}
	return ""
}

func (x *ReferOutResponse_Results) GetExpireDate() string {
	if x != nil {
		return x.ExpireDate
	}
	return ""
}

func (x *ReferOutResponse_Results) GetReferTypeId() string {
	if x != nil {
		return x.ReferTypeId
	}
	return ""
}

func (x *ReferOutResponse_Results) GetReferTypeName() string {
	if x != nil {
		return x.ReferTypeName
	}
	return ""
}

func (x *ReferOutResponse_Results) GetReferCauseId() string {
	if x != nil {
		return x.ReferCauseId
	}
	return ""
}

func (x *ReferOutResponse_Results) GetReferCauseName() string {
	if x != nil {
		return x.ReferCauseName
	}
	return ""
}

func (x *ReferOutResponse_Results) GetRfrcsId() string {
	if x != nil {
		return x.RfrcsId
	}
	return ""
}

func (x *ReferOutResponse_Results) GetRfrcsName() string {
	if x != nil {
		return x.RfrcsName
	}
	return ""
}

type ReferVisitResponse_Results struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hospcode       string `protobuf:"bytes,1,opt,name=hospcode,proto3" json:"hospcode,omitempty"`
	VisitDate      string `protobuf:"bytes,2,opt,name=visit_date,json=visitDate,proto3" json:"visit_date,omitempty"`
	VisitTime      string `protobuf:"bytes,3,opt,name=visit_time,json=visitTime,proto3" json:"visit_time,omitempty"`
	Cid            string `protobuf:"bytes,4,opt,name=cid,proto3" json:"cid,omitempty"`
	Hn             string `protobuf:"bytes,5,opt,name=hn,proto3" json:"hn,omitempty"`
	Vn             string `protobuf:"bytes,6,opt,name=vn,proto3" json:"vn,omitempty"`
	An             string `protobuf:"bytes,7,opt,name=an,proto3" json:"an,omitempty"`
	Fname          string `protobuf:"bytes,8,opt,name=fname,proto3" json:"fname,omitempty"`
	Lname          string `protobuf:"bytes,9,opt,name=lname,proto3" json:"lname,omitempty"`
	Birthday       string `protobuf:"bytes,10,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Sex            string `protobuf:"bytes,11,opt,name=sex,proto3" json:"sex,omitempty"`
	Pttype         string `protobuf:"bytes,12,opt,name=pttype,proto3" json:"pttype,omitempty"`
	PttypeNo       string `protobuf:"bytes,13,opt,name=pttype_no,json=pttypeNo,proto3" json:"pttype_no,omitempty"`
	PttypeName     string `protobuf:"bytes,14,opt,name=pttype_name,json=pttypeName,proto3" json:"pttype_name,omitempty"`
	HospmainCode   string `protobuf:"bytes,15,opt,name=hospmain_code,json=hospmainCode,proto3" json:"hospmain_code,omitempty"`
	HospmainName   string `protobuf:"bytes,16,opt,name=hospmain_name,json=hospmainName,proto3" json:"hospmain_name,omitempty"`
	HospsubCode    string `protobuf:"bytes,17,opt,name=hospsub_code,json=hospsubCode,proto3" json:"hospsub_code,omitempty"`
	HospsubName    string `protobuf:"bytes,18,opt,name=hospsub_name,json=hospsubName,proto3" json:"hospsub_name,omitempty"`
	MainDepartment string `protobuf:"bytes,19,opt,name=main_department,json=mainDepartment,proto3" json:"main_department,omitempty"`
	SubDepartment  string `protobuf:"bytes,20,opt,name=sub_department,json=subDepartment,proto3" json:"sub_department,omitempty"`
	Diagcode       string `protobuf:"bytes,21,opt,name=diagcode,proto3" json:"diagcode,omitempty"`
	Diagname       string `protobuf:"bytes,22,opt,name=diagname,proto3" json:"diagname,omitempty"`
}

func (x *ReferVisitResponse_Results) Reset() {
	*x = ReferVisitResponse_Results{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_refer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReferVisitResponse_Results) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReferVisitResponse_Results) ProtoMessage() {}

func (x *ReferVisitResponse_Results) ProtoReflect() protoreflect.Message {
	mi := &file_src_refer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReferVisitResponse_Results.ProtoReflect.Descriptor instead.
func (*ReferVisitResponse_Results) Descriptor() ([]byte, []int) {
	return file_src_refer_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ReferVisitResponse_Results) GetHospcode() string {
	if x != nil {
		return x.Hospcode
	}
	return ""
}

func (x *ReferVisitResponse_Results) GetVisitDate() string {
	if x != nil {
		return x.VisitDate
	}
	return ""
}

func (x *ReferVisitResponse_Results) GetVisitTime() string {
	if x != nil {
		return x.VisitTime
	}
	return ""
}

func (x *ReferVisitResponse_Results) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

func (x *ReferVisitResponse_Results) GetHn() string {
	if x != nil {
		return x.Hn
	}
	return ""
}

func (x *ReferVisitResponse_Results) GetVn() string {
	if x != nil {
		return x.Vn
	}
	return ""
}

func (x *ReferVisitResponse_Results) GetAn() string {
	if x != nil {
		return x.An
	}
	return ""
}

func (x *ReferVisitResponse_Results) GetFname() string {
	if x != nil {
		return x.Fname
	}
	return ""
}

func (x *ReferVisitResponse_Results) GetLname() string {
	if x != nil {
		return x.Lname
	}
	return ""
}

func (x *ReferVisitResponse_Results) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *ReferVisitResponse_Results) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *ReferVisitResponse_Results) GetPttype() string {
	if x != nil {
		return x.Pttype
	}
	return ""
}

func (x *ReferVisitResponse_Results) GetPttypeNo() string {
	if x != nil {
		return x.PttypeNo
	}
	return ""
}

func (x *ReferVisitResponse_Results) GetPttypeName() string {
	if x != nil {
		return x.PttypeName
	}
	return ""
}

func (x *ReferVisitResponse_Results) GetHospmainCode() string {
	if x != nil {
		return x.HospmainCode
	}
	return ""
}

func (x *ReferVisitResponse_Results) GetHospmainName() string {
	if x != nil {
		return x.HospmainName
	}
	return ""
}

func (x *ReferVisitResponse_Results) GetHospsubCode() string {
	if x != nil {
		return x.HospsubCode
	}
	return ""
}

func (x *ReferVisitResponse_Results) GetHospsubName() string {
	if x != nil {
		return x.HospsubName
	}
	return ""
}

func (x *ReferVisitResponse_Results) GetMainDepartment() string {
	if x != nil {
		return x.MainDepartment
	}
	return ""
}

func (x *ReferVisitResponse_Results) GetSubDepartment() string {
	if x != nil {
		return x.SubDepartment
	}
	return ""
}

func (x *ReferVisitResponse_Results) GetDiagcode() string {
	if x != nil {
		return x.Diagcode
	}
	return ""
}

func (x *ReferVisitResponse_Results) GetDiagname() string {
	if x != nil {
		return x.Diagname
	}
	return ""
}

var File_src_refer_proto protoreflect.FileDescriptor

var file_src_refer_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x72, 0x63, 0x2f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x41, 0x0a, 0x0f, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x4f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f,
	0x73, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f,
	0x73, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x43, 0x0a, 0x11, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x56, 0x69, 0x73, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22,
	0x9f, 0x07, 0x0a, 0x10, 0x52, 0x65, 0x66, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0xd0,
	0x06, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f,
	0x73, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f,
	0x73, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x68, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x68, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73,
	0x65, 0x78, 0x12, 0x0e, 0x0a, 0x02, 0x76, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x76, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x66, 0x65, 0x72, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x66, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x74, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x74, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x74, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x74, 0x74, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x72, 0x65, 0x61, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x72, 0x65, 0x61, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x69, 0x61, 0x67, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x69, 0x61, 0x67, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x61, 0x67, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x61, 0x67, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x5f, 0x64, 0x69, 0x61, 0x67, 0x6e,
	0x6f, 0x73, 0x69, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x44,
	0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x69, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6d, 0x68, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x6d, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x68, 0x70,
	0x69, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x68, 0x70, 0x69, 0x12, 0x25, 0x0a, 0x0e,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x5f, 0x68, 0x6f, 0x73, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x16,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x66, 0x65, 0x72, 0x48, 0x6f, 0x73, 0x70, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x66, 0x65, 0x72, 0x5f, 0x68, 0x6f, 0x73,
	0x70, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x48, 0x6f, 0x73, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x19, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x66, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12,
	0x26, 0x0a, 0x0f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x66, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x5f, 0x63, 0x61, 0x75, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x66, 0x65, 0x72, 0x43, 0x61, 0x75, 0x73, 0x65, 0x49, 0x64, 0x12, 0x28, 0x0a,
	0x10, 0x72, 0x65, 0x66, 0x65, 0x72, 0x5f, 0x63, 0x61, 0x75, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x66, 0x65, 0x72, 0x43, 0x61,
	0x75, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x66, 0x72, 0x63, 0x73,
	0x5f, 0x69, 0x64, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x66, 0x72, 0x63, 0x73,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x66, 0x72, 0x63, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x66, 0x72, 0x63, 0x73, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0xc0, 0x05, 0x0a, 0x12, 0x52, 0x65, 0x66, 0x65, 0x72, 0x56, 0x69, 0x73, 0x69, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x56, 0x69, 0x73, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x1a, 0xed, 0x04, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x76, 0x69, 0x73, 0x69, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x76, 0x69, 0x73, 0x69, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76,
	0x69, 0x73, 0x69, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x76, 0x69, 0x73, 0x69, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x68, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x68, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x76, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x76, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x61, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x61, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x66, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74,
	0x68, 0x64, 0x61, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74,
	0x68, 0x64, 0x61, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x74, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x74, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x74, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x74, 0x74, 0x79, 0x70, 0x65, 0x4e, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x74, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x74, 0x74, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x68, 0x6f, 0x73, 0x70, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x68, 0x6f, 0x73, 0x70, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x68, 0x6f, 0x73, 0x70, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x68, 0x6f, 0x73, 0x70, 0x6d, 0x61,
	0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x6f, 0x73, 0x70, 0x73, 0x75,
	0x62, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x6f,
	0x73, 0x70, 0x73, 0x75, 0x62, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x6f, 0x73,
	0x70, 0x73, 0x75, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x68, 0x6f, 0x73, 0x70, 0x73, 0x75, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f,
	0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x61, 0x69, 0x6e, 0x44, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x75, 0x62, 0x5f, 0x64, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73,
	0x75, 0x62, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x69, 0x61, 0x67, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x69, 0x61, 0x67, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x61, 0x67,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x61, 0x67,
	0x6e, 0x61, 0x6d, 0x65, 0x32, 0x89, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x52, 0x65, 0x66, 0x65, 0x72, 0x4f, 0x75,
	0x74, 0x12, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x4f, 0x75,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3c, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x15, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x66, 0x65,
	0x72, 0x56, 0x69, 0x73, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x3b, 0x67, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_refer_proto_rawDescOnce sync.Once
	file_src_refer_proto_rawDescData = file_src_refer_proto_rawDesc
)

func file_src_refer_proto_rawDescGZIP() []byte {
	file_src_refer_proto_rawDescOnce.Do(func() {
		file_src_refer_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_refer_proto_rawDescData)
	})
	return file_src_refer_proto_rawDescData
}

var file_src_refer_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_src_refer_proto_goTypes = []interface{}{
	(*ReferOutRequest)(nil),            // 0: grpc.ReferOutRequest
	(*ReferVisitRequest)(nil),          // 1: grpc.ReferVisitRequest
	(*ReferOutResponse)(nil),           // 2: grpc.ReferOutResponse
	(*ReferVisitResponse)(nil),         // 3: grpc.ReferVisitResponse
	(*ReferOutResponse_Results)(nil),   // 4: grpc.ReferOutResponse.Results
	(*ReferVisitResponse_Results)(nil), // 5: grpc.ReferVisitResponse.Results
}
var file_src_refer_proto_depIdxs = []int32{
	4, // 0: grpc.ReferOutResponse.results:type_name -> grpc.ReferOutResponse.Results
	5, // 1: grpc.ReferVisitResponse.results:type_name -> grpc.ReferVisitResponse.Results
	0, // 2: grpc.ReferService.ReferOut:input_type -> grpc.ReferOutRequest
	0, // 3: grpc.ReferService.Service:input_type -> grpc.ReferOutRequest
	2, // 4: grpc.ReferService.ReferOut:output_type -> grpc.ReferOutResponse
	3, // 5: grpc.ReferService.Service:output_type -> grpc.ReferVisitResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_src_refer_proto_init() }
func file_src_refer_proto_init() {
	if File_src_refer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_refer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReferOutRequest); i {
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
		file_src_refer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReferVisitRequest); i {
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
		file_src_refer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReferOutResponse); i {
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
		file_src_refer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReferVisitResponse); i {
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
		file_src_refer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReferOutResponse_Results); i {
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
		file_src_refer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReferVisitResponse_Results); i {
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
			RawDescriptor: file_src_refer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_refer_proto_goTypes,
		DependencyIndexes: file_src_refer_proto_depIdxs,
		MessageInfos:      file_src_refer_proto_msgTypes,
	}.Build()
	File_src_refer_proto = out.File
	file_src_refer_proto_rawDesc = nil
	file_src_refer_proto_goTypes = nil
	file_src_refer_proto_depIdxs = nil
}
