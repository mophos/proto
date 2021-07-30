// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MhealthServiceClient is the client API for MhealthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MhealthServiceClient interface {
	AppointmentDateserve(ctx context.Context, in *DateserveRequest, opts ...grpc.CallOption) (*AppointmentDateserveResponse, error)
}

type mhealthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMhealthServiceClient(cc grpc.ClientConnInterface) MhealthServiceClient {
	return &mhealthServiceClient{cc}
}

func (c *mhealthServiceClient) AppointmentDateserve(ctx context.Context, in *DateserveRequest, opts ...grpc.CallOption) (*AppointmentDateserveResponse, error) {
	out := new(AppointmentDateserveResponse)
	err := c.cc.Invoke(ctx, "/grpc.MhealthService/AppointmentDateserve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MhealthServiceServer is the server API for MhealthService service.
// All implementations must embed UnimplementedMhealthServiceServer
// for forward compatibility
type MhealthServiceServer interface {
	AppointmentDateserve(context.Context, *DateserveRequest) (*AppointmentDateserveResponse, error)
	mustEmbedUnimplementedMhealthServiceServer()
}

// UnimplementedMhealthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMhealthServiceServer struct {
}

func (UnimplementedMhealthServiceServer) AppointmentDateserve(context.Context, *DateserveRequest) (*AppointmentDateserveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppointmentDateserve not implemented")
}
func (UnimplementedMhealthServiceServer) mustEmbedUnimplementedMhealthServiceServer() {}

// UnsafeMhealthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MhealthServiceServer will
// result in compilation errors.
type UnsafeMhealthServiceServer interface {
	mustEmbedUnimplementedMhealthServiceServer()
}

func RegisterMhealthServiceServer(s grpc.ServiceRegistrar, srv MhealthServiceServer) {
	s.RegisterService(&MhealthService_ServiceDesc, srv)
}

func _MhealthService_AppointmentDateserve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DateserveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MhealthServiceServer).AppointmentDateserve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MhealthService/AppointmentDateserve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MhealthServiceServer).AppointmentDateserve(ctx, req.(*DateserveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MhealthService_ServiceDesc is the grpc.ServiceDesc for MhealthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MhealthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.MhealthService",
	HandlerType: (*MhealthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppointmentDateserve",
			Handler:    _MhealthService_AppointmentDateserve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/h4u.proto",
}

// MasterServiceClient is the client API for MasterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MasterServiceClient interface {
	DoctorList(ctx context.Context, in *HospcodeRequest, opts ...grpc.CallOption) (*ListDoctorResponse, error)
	ClinicList(ctx context.Context, in *HospcodeRequest, opts ...grpc.CallOption) (*ListClinicResponse, error)
	HisProviderList(ctx context.Context, in *HospcodeRequest, opts ...grpc.CallOption) (*HisProviderResponse, error)
	CountRecord(ctx context.Context, in *ProviderRequest, opts ...grpc.CallOption) (*CountResponse, error)
}

type masterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterServiceClient(cc grpc.ClientConnInterface) MasterServiceClient {
	return &masterServiceClient{cc}
}

func (c *masterServiceClient) DoctorList(ctx context.Context, in *HospcodeRequest, opts ...grpc.CallOption) (*ListDoctorResponse, error) {
	out := new(ListDoctorResponse)
	err := c.cc.Invoke(ctx, "/grpc.MasterService/DoctorList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) ClinicList(ctx context.Context, in *HospcodeRequest, opts ...grpc.CallOption) (*ListClinicResponse, error) {
	out := new(ListClinicResponse)
	err := c.cc.Invoke(ctx, "/grpc.MasterService/ClinicList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) HisProviderList(ctx context.Context, in *HospcodeRequest, opts ...grpc.CallOption) (*HisProviderResponse, error) {
	out := new(HisProviderResponse)
	err := c.cc.Invoke(ctx, "/grpc.MasterService/HisProviderList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) CountRecord(ctx context.Context, in *ProviderRequest, opts ...grpc.CallOption) (*CountResponse, error) {
	out := new(CountResponse)
	err := c.cc.Invoke(ctx, "/grpc.MasterService/CountRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterServiceServer is the server API for MasterService service.
// All implementations must embed UnimplementedMasterServiceServer
// for forward compatibility
type MasterServiceServer interface {
	DoctorList(context.Context, *HospcodeRequest) (*ListDoctorResponse, error)
	ClinicList(context.Context, *HospcodeRequest) (*ListClinicResponse, error)
	HisProviderList(context.Context, *HospcodeRequest) (*HisProviderResponse, error)
	CountRecord(context.Context, *ProviderRequest) (*CountResponse, error)
	mustEmbedUnimplementedMasterServiceServer()
}

// UnimplementedMasterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMasterServiceServer struct {
}

func (UnimplementedMasterServiceServer) DoctorList(context.Context, *HospcodeRequest) (*ListDoctorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoctorList not implemented")
}
func (UnimplementedMasterServiceServer) ClinicList(context.Context, *HospcodeRequest) (*ListClinicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClinicList not implemented")
}
func (UnimplementedMasterServiceServer) HisProviderList(context.Context, *HospcodeRequest) (*HisProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HisProviderList not implemented")
}
func (UnimplementedMasterServiceServer) CountRecord(context.Context, *ProviderRequest) (*CountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountRecord not implemented")
}
func (UnimplementedMasterServiceServer) mustEmbedUnimplementedMasterServiceServer() {}

// UnsafeMasterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MasterServiceServer will
// result in compilation errors.
type UnsafeMasterServiceServer interface {
	mustEmbedUnimplementedMasterServiceServer()
}

func RegisterMasterServiceServer(s grpc.ServiceRegistrar, srv MasterServiceServer) {
	s.RegisterService(&MasterService_ServiceDesc, srv)
}

func _MasterService_DoctorList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HospcodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).DoctorList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MasterService/DoctorList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).DoctorList(ctx, req.(*HospcodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_ClinicList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HospcodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).ClinicList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MasterService/ClinicList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).ClinicList(ctx, req.(*HospcodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_HisProviderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HospcodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).HisProviderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MasterService/HisProviderList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).HisProviderList(ctx, req.(*HospcodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_CountRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).CountRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MasterService/CountRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).CountRecord(ctx, req.(*ProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MasterService_ServiceDesc is the grpc.ServiceDesc for MasterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MasterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.MasterService",
	HandlerType: (*MasterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoctorList",
			Handler:    _MasterService_DoctorList_Handler,
		},
		{
			MethodName: "ClinicList",
			Handler:    _MasterService_ClinicList_Handler,
		},
		{
			MethodName: "HisProviderList",
			Handler:    _MasterService_HisProviderList_Handler,
		},
		{
			MethodName: "CountRecord",
			Handler:    _MasterService_CountRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/h4u.proto",
}

// AuthenServiceClient is the client API for AuthenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenServiceClient interface {
	Authen(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error)
}

type authenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenServiceClient(cc grpc.ClientConnInterface) AuthenServiceClient {
	return &authenServiceClient{cc}
}

func (c *authenServiceClient) Authen(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, "/grpc.AuthenService/Authen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenServiceServer is the server API for AuthenService service.
// All implementations must embed UnimplementedAuthenServiceServer
// for forward compatibility
type AuthenServiceServer interface {
	Authen(context.Context, *TokenRequest) (*TokenResponse, error)
	mustEmbedUnimplementedAuthenServiceServer()
}

// UnimplementedAuthenServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenServiceServer struct {
}

func (UnimplementedAuthenServiceServer) Authen(context.Context, *TokenRequest) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authen not implemented")
}
func (UnimplementedAuthenServiceServer) mustEmbedUnimplementedAuthenServiceServer() {}

// UnsafeAuthenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenServiceServer will
// result in compilation errors.
type UnsafeAuthenServiceServer interface {
	mustEmbedUnimplementedAuthenServiceServer()
}

func RegisterAuthenServiceServer(s grpc.ServiceRegistrar, srv AuthenServiceServer) {
	s.RegisterService(&AuthenService_ServiceDesc, srv)
}

func _AuthenService_Authen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenServiceServer).Authen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.AuthenService/Authen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenServiceServer).Authen(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthenService_ServiceDesc is the grpc.ServiceDesc for AuthenService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.AuthenService",
	HandlerType: (*AuthenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authen",
			Handler:    _AuthenService_Authen_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/h4u.proto",
}

// H4UServiceClient is the client API for H4UService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type H4UServiceClient interface {
	H4UPersonal(ctx context.Context, in *Request, opts ...grpc.CallOption) (*H4UPersonalResponse, error)
	H4UVisit(ctx context.Context, in *Request, opts ...grpc.CallOption) (*H4UVisitResponse, error)
	H4ULab(ctx context.Context, in *Request, opts ...grpc.CallOption) (*H4ULabResponse, error)
	H4UOrder(ctx context.Context, in *Request, opts ...grpc.CallOption) (*H4UOrderResponse, error)
	H4UDiagnosis(ctx context.Context, in *Request, opts ...grpc.CallOption) (*H4UDiagnosisResponse, error)
}

type h4UServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewH4UServiceClient(cc grpc.ClientConnInterface) H4UServiceClient {
	return &h4UServiceClient{cc}
}

func (c *h4UServiceClient) H4UPersonal(ctx context.Context, in *Request, opts ...grpc.CallOption) (*H4UPersonalResponse, error) {
	out := new(H4UPersonalResponse)
	err := c.cc.Invoke(ctx, "/grpc.H4UService/H4UPersonal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *h4UServiceClient) H4UVisit(ctx context.Context, in *Request, opts ...grpc.CallOption) (*H4UVisitResponse, error) {
	out := new(H4UVisitResponse)
	err := c.cc.Invoke(ctx, "/grpc.H4UService/H4UVisit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *h4UServiceClient) H4ULab(ctx context.Context, in *Request, opts ...grpc.CallOption) (*H4ULabResponse, error) {
	out := new(H4ULabResponse)
	err := c.cc.Invoke(ctx, "/grpc.H4UService/H4ULab", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *h4UServiceClient) H4UOrder(ctx context.Context, in *Request, opts ...grpc.CallOption) (*H4UOrderResponse, error) {
	out := new(H4UOrderResponse)
	err := c.cc.Invoke(ctx, "/grpc.H4UService/H4UOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *h4UServiceClient) H4UDiagnosis(ctx context.Context, in *Request, opts ...grpc.CallOption) (*H4UDiagnosisResponse, error) {
	out := new(H4UDiagnosisResponse)
	err := c.cc.Invoke(ctx, "/grpc.H4UService/H4UDiagnosis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// H4UServiceServer is the server API for H4UService service.
// All implementations must embed UnimplementedH4UServiceServer
// for forward compatibility
type H4UServiceServer interface {
	H4UPersonal(context.Context, *Request) (*H4UPersonalResponse, error)
	H4UVisit(context.Context, *Request) (*H4UVisitResponse, error)
	H4ULab(context.Context, *Request) (*H4ULabResponse, error)
	H4UOrder(context.Context, *Request) (*H4UOrderResponse, error)
	H4UDiagnosis(context.Context, *Request) (*H4UDiagnosisResponse, error)
	mustEmbedUnimplementedH4UServiceServer()
}

// UnimplementedH4UServiceServer must be embedded to have forward compatible implementations.
type UnimplementedH4UServiceServer struct {
}

func (UnimplementedH4UServiceServer) H4UPersonal(context.Context, *Request) (*H4UPersonalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method H4UPersonal not implemented")
}
func (UnimplementedH4UServiceServer) H4UVisit(context.Context, *Request) (*H4UVisitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method H4UVisit not implemented")
}
func (UnimplementedH4UServiceServer) H4ULab(context.Context, *Request) (*H4ULabResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method H4ULab not implemented")
}
func (UnimplementedH4UServiceServer) H4UOrder(context.Context, *Request) (*H4UOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method H4UOrder not implemented")
}
func (UnimplementedH4UServiceServer) H4UDiagnosis(context.Context, *Request) (*H4UDiagnosisResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method H4UDiagnosis not implemented")
}
func (UnimplementedH4UServiceServer) mustEmbedUnimplementedH4UServiceServer() {}

// UnsafeH4UServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to H4UServiceServer will
// result in compilation errors.
type UnsafeH4UServiceServer interface {
	mustEmbedUnimplementedH4UServiceServer()
}

func RegisterH4UServiceServer(s grpc.ServiceRegistrar, srv H4UServiceServer) {
	s.RegisterService(&H4UService_ServiceDesc, srv)
}

func _H4UService_H4UPersonal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(H4UServiceServer).H4UPersonal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.H4UService/H4UPersonal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(H4UServiceServer).H4UPersonal(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _H4UService_H4UVisit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(H4UServiceServer).H4UVisit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.H4UService/H4UVisit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(H4UServiceServer).H4UVisit(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _H4UService_H4ULab_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(H4UServiceServer).H4ULab(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.H4UService/H4ULab",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(H4UServiceServer).H4ULab(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _H4UService_H4UOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(H4UServiceServer).H4UOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.H4UService/H4UOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(H4UServiceServer).H4UOrder(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _H4UService_H4UDiagnosis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(H4UServiceServer).H4UDiagnosis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.H4UService/H4UDiagnosis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(H4UServiceServer).H4UDiagnosis(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// H4UService_ServiceDesc is the grpc.ServiceDesc for H4UService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var H4UService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.H4UService",
	HandlerType: (*H4UServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "H4UPersonal",
			Handler:    _H4UService_H4UPersonal_Handler,
		},
		{
			MethodName: "H4UVisit",
			Handler:    _H4UService_H4UVisit_Handler,
		},
		{
			MethodName: "H4ULab",
			Handler:    _H4UService_H4ULab_Handler,
		},
		{
			MethodName: "H4UOrder",
			Handler:    _H4UService_H4UOrder_Handler,
		},
		{
			MethodName: "H4UDiagnosis",
			Handler:    _H4UService_H4UDiagnosis_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/h4u.proto",
}
