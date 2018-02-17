// Code generated by protoc-gen-go. DO NOT EDIT.
// source: team.proto

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	team.proto

It has these top-level messages:
	CreateRequest
	GetRequest
	GetResponse
	IsAvailableRequest
	Address
*/
package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "appscode.com/api/dtypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateRequest struct {
	Name               string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	DisplayName        string   `protobuf:"bytes,2,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	Email              string   `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	UserName           string   `protobuf:"bytes,4,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	Password           string   `protobuf:"bytes,5,opt,name=password" json:"password,omitempty"`
	InviteEmails       []string `protobuf:"bytes,6,rep,name=invite_emails,json=inviteEmails" json:"invite_emails,omitempty"`
	Subscription       string   `protobuf:"bytes,7,opt,name=subscription" json:"subscription,omitempty"`
	InitialUnits       int64    `protobuf:"varint,8,opt,name=initial_units,json=initialUnits" json:"initial_units,omitempty"`
	PaymentMethodNonce string   `protobuf:"bytes,9,opt,name=payment_method_nonce,json=paymentMethodNonce" json:"payment_method_nonce,omitempty"`
	BillingAddress     *Address `protobuf:"bytes,10,opt,name=billing_address,json=billingAddress" json:"billing_address,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *CreateRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *CreateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateRequest) GetInviteEmails() []string {
	if m != nil {
		return m.InviteEmails
	}
	return nil
}

func (m *CreateRequest) GetSubscription() string {
	if m != nil {
		return m.Subscription
	}
	return ""
}

func (m *CreateRequest) GetInitialUnits() int64 {
	if m != nil {
		return m.InitialUnits
	}
	return 0
}

func (m *CreateRequest) GetPaymentMethodNonce() string {
	if m != nil {
		return m.PaymentMethodNonce
	}
	return ""
}

func (m *CreateRequest) GetBillingAddress() *Address {
	if m != nil {
		return m.BillingAddress
	}
	return nil
}

type GetRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetResponse struct {
	Phid string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetResponse) GetPhid() string {
	if m != nil {
		return m.Phid
	}
	return ""
}

type IsAvailableRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *IsAvailableRequest) Reset()                    { *m = IsAvailableRequest{} }
func (m *IsAvailableRequest) String() string            { return proto.CompactTextString(m) }
func (*IsAvailableRequest) ProtoMessage()               {}
func (*IsAvailableRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *IsAvailableRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Example
// result = Braintree::Address.create(
//   :first_name          => 'Jenna',
//   :last_name           => 'Smith',
//   :company             => 'Braintree',
//   :street_address      => '1 E Main St',
//   :extended_address    => 'Suite 403',
//   :locality            => 'Chicago',
//   :region              => 'Illinois',
//   :postal_code         => '60622',
//   :country_code_numeric => '840'
// )
type Address struct {
	FirstName       string `protobuf:"bytes,1,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName        string `protobuf:"bytes,2,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	Company         string `protobuf:"bytes,3,opt,name=company" json:"company,omitempty"`
	StreetAddress   string `protobuf:"bytes,4,opt,name=street_address,json=streetAddress" json:"street_address,omitempty"`
	ExtendedAddress string `protobuf:"bytes,5,opt,name=extended_address,json=extendedAddress" json:"extended_address,omitempty"`
	Locality        string `protobuf:"bytes,6,opt,name=locality" json:"locality,omitempty"`
	Region          string `protobuf:"bytes,7,opt,name=region" json:"region,omitempty"`
	PostalCode      string `protobuf:"bytes,8,opt,name=postal_code,json=postalCode" json:"postal_code,omitempty"`
	// Ref https://developers.braintreepayments.com/reference/general/countries/ruby
	CountryCodeNumeric string `protobuf:"bytes,9,opt,name=country_code_numeric,json=countryCodeNumeric" json:"country_code_numeric,omitempty"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Address) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Address) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Address) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *Address) GetStreetAddress() string {
	if m != nil {
		return m.StreetAddress
	}
	return ""
}

func (m *Address) GetExtendedAddress() string {
	if m != nil {
		return m.ExtendedAddress
	}
	return ""
}

func (m *Address) GetLocality() string {
	if m != nil {
		return m.Locality
	}
	return ""
}

func (m *Address) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *Address) GetPostalCode() string {
	if m != nil {
		return m.PostalCode
	}
	return ""
}

func (m *Address) GetCountryCodeNumeric() string {
	if m != nil {
		return m.CountryCodeNumeric
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "appscode.namespace.v1beta1.CreateRequest")
	proto.RegisterType((*GetRequest)(nil), "appscode.namespace.v1beta1.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "appscode.namespace.v1beta1.GetResponse")
	proto.RegisterType((*IsAvailableRequest)(nil), "appscode.namespace.v1beta1.IsAvailableRequest")
	proto.RegisterType((*Address)(nil), "appscode.namespace.v1beta1.Address")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Teams service

type TeamsClient interface {
	// Creates a new namespace, if name is valid and no namespace with same name exists.
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	// Gets a namespace, if exists. This can be used to check existence of a namespace.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// Check if a namespace name is available meaning name is valid and no namespace with same name exists.
	IsAvailable(ctx context.Context, in *IsAvailableRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
}

type teamsClient struct {
	cc *grpc.ClientConn
}

func NewTeamsClient(cc *grpc.ClientConn) TeamsClient {
	return &teamsClient{cc}
}

func (c *teamsClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.namespace.v1beta1.Teams/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/appscode.namespace.v1beta1.Teams/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) IsAvailable(ctx context.Context, in *IsAvailableRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.namespace.v1beta1.Teams/IsAvailable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Teams service

type TeamsServer interface {
	// Creates a new namespace, if name is valid and no namespace with same name exists.
	Create(context.Context, *CreateRequest) (*appscode_dtypes.VoidResponse, error)
	// Gets a namespace, if exists. This can be used to check existence of a namespace.
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// Check if a namespace name is available meaning name is valid and no namespace with same name exists.
	IsAvailable(context.Context, *IsAvailableRequest) (*appscode_dtypes.VoidResponse, error)
}

func RegisterTeamsServer(s *grpc.Server, srv TeamsServer) {
	s.RegisterService(&_Teams_serviceDesc, srv)
}

func _Teams_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.namespace.v1beta1.Teams/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.namespace.v1beta1.Teams/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_IsAvailable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAvailableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).IsAvailable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.namespace.v1beta1.Teams/IsAvailable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).IsAvailable(ctx, req.(*IsAvailableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Teams_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.namespace.v1beta1.Teams",
	HandlerType: (*TeamsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Teams_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Teams_Get_Handler,
		},
		{
			MethodName: "IsAvailable",
			Handler:    _Teams_IsAvailable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "team.proto",
}

func init() { proto.RegisterFile("team.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 703 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6a, 0xdb, 0x48,
	0x14, 0x46, 0x76, 0x62, 0x5b, 0xc7, 0xf9, 0x59, 0x86, 0xb0, 0x08, 0x67, 0xb3, 0xeb, 0x28, 0xec,
	0xae, 0xb3, 0xcb, 0xda, 0x49, 0x76, 0xaf, 0xc2, 0x42, 0x9b, 0x84, 0x12, 0x0a, 0x4d, 0x08, 0xa6,
	0xed, 0x45, 0x6f, 0xc4, 0x58, 0x3a, 0x75, 0xa6, 0x48, 0x33, 0x53, 0xcd, 0x38, 0xad, 0x29, 0xbd,
	0x49, 0x5f, 0xa0, 0xd0, 0x9b, 0xbe, 0x43, 0xaf, 0x0b, 0x7d, 0x8f, 0x3e, 0x42, 0xfb, 0x20, 0x65,
	0x66, 0x24, 0xdb, 0x21, 0xc4, 0x69, 0x6f, 0x8c, 0xce, 0xf7, 0x7d, 0x73, 0xce, 0x67, 0x7d, 0x67,
	0x04, 0xa0, 0x91, 0x66, 0x5d, 0x99, 0x0b, 0x2d, 0x48, 0x8b, 0x4a, 0xa9, 0x62, 0x91, 0x60, 0x97,
	0xd3, 0x0c, 0x95, 0xa4, 0x31, 0x76, 0x2f, 0x76, 0x07, 0xa8, 0xe9, 0x6e, 0xeb, 0x97, 0xa1, 0x10,
	0xc3, 0x14, 0x7b, 0x54, 0xb2, 0x1e, 0xe5, 0x5c, 0x68, 0xaa, 0x99, 0xe0, 0xca, 0x9d, 0x6c, 0xfd,
	0x5a, 0x9e, 0xbc, 0x81, 0xdf, 0x9a, 0x74, 0x8e, 0x45, 0x66, 0x35, 0x89, 0x1e, 0x4b, 0x54, 0x3d,
	0xfb, 0xeb, 0x44, 0xe1, 0x9b, 0x2a, 0x2c, 0x1f, 0xe5, 0x48, 0x35, 0xf6, 0xf1, 0xf9, 0x08, 0x95,
	0x26, 0x04, 0x16, 0x8c, 0x93, 0xc0, 0x6b, 0x7b, 0x1d, 0xbf, 0x6f, 0x9f, 0xc9, 0x26, 0x2c, 0x25,
	0x4c, 0xc9, 0x94, 0x8e, 0x23, 0xcb, 0x55, 0x2c, 0xd7, 0x2c, 0xb0, 0x53, 0x23, 0x59, 0x83, 0x45,
	0xcc, 0x28, 0x4b, 0x83, 0xaa, 0xe5, 0x5c, 0x41, 0xd6, 0xc1, 0x1f, 0x29, 0xcc, 0xdd, 0xa9, 0x05,
	0xcb, 0x34, 0x0c, 0x60, 0x8f, 0xb4, 0xa0, 0x21, 0xa9, 0x52, 0x2f, 0x44, 0x9e, 0x04, 0x8b, 0x8e,
	0x2b, 0x6b, 0xb2, 0x05, 0xcb, 0x8c, 0x5f, 0x30, 0x8d, 0x91, 0x6d, 0xa4, 0x82, 0x5a, 0xbb, 0xda,
	0xf1, 0xfb, 0x4b, 0x0e, 0xbc, 0x67, 0x31, 0x12, 0xc2, 0x92, 0x1a, 0x0d, 0x54, 0x9c, 0x33, 0x69,
	0xfe, 0x78, 0x50, 0xb7, 0x4d, 0xae, 0x60, 0xae, 0x11, 0xd3, 0x8c, 0xa6, 0xd1, 0x88, 0x33, 0xad,
	0x82, 0x46, 0xdb, 0xeb, 0x54, 0x4d, 0x23, 0x0b, 0x3e, 0x32, 0x18, 0xd9, 0x81, 0x35, 0x49, 0xc7,
	0x19, 0x72, 0x1d, 0x65, 0xa8, 0xcf, 0x45, 0x12, 0x71, 0xc1, 0x63, 0x0c, 0x7c, 0xdb, 0x90, 0x14,
	0xdc, 0x89, 0xa5, 0x4e, 0x0d, 0x43, 0x1e, 0xc0, 0xea, 0x80, 0xa5, 0x29, 0xe3, 0xc3, 0x88, 0x26,
	0x49, 0x8e, 0x4a, 0x05, 0xd0, 0xf6, 0x3a, 0xcd, 0xbd, 0xad, 0xee, 0xcd, 0x81, 0x76, 0x0f, 0x9c,
	0xb4, 0xbf, 0x52, 0x9c, 0x2d, 0xea, 0xb0, 0x0d, 0x70, 0x8c, 0x7a, 0x4e, 0x02, 0xe1, 0x26, 0x34,
	0xad, 0x42, 0x49, 0xc1, 0x15, 0x1a, 0x89, 0x3c, 0x67, 0x49, 0x29, 0x31, 0xcf, 0x61, 0x07, 0xc8,
	0x7d, 0x75, 0x70, 0x41, 0x59, 0x4a, 0x07, 0xe9, 0xbc, 0x38, 0xc3, 0x4f, 0x15, 0xa8, 0x17, 0xa3,
	0xc9, 0x06, 0xc0, 0x53, 0x96, 0x2b, 0x1d, 0xcd, 0xa8, 0x7c, 0x8b, 0xd8, 0x8c, 0xd6, 0xc1, 0x4f,
	0x69, 0xc9, 0xba, 0xd8, 0x1b, 0x06, 0xb0, 0x64, 0x00, 0xf5, 0x58, 0x64, 0x92, 0xf2, 0x71, 0x91,
	0x7a, 0x59, 0x92, 0xdf, 0x61, 0x45, 0xe9, 0x1c, 0x51, 0x4f, 0xde, 0x8e, 0x0b, 0x7f, 0xd9, 0xa1,
	0xe5, 0xf0, 0x6d, 0xf8, 0x09, 0x5f, 0x6a, 0xe4, 0x09, 0x26, 0x13, 0xa1, 0xdb, 0x84, 0xd5, 0x12,
	0x2f, 0xa5, 0x2d, 0x68, 0xa4, 0x22, 0xa6, 0x29, 0xd3, 0xe3, 0xa0, 0x56, 0xf8, 0x28, 0x6a, 0xf2,
	0x33, 0xd4, 0x72, 0x1c, 0x4e, 0x37, 0xa0, 0xa8, 0xc8, 0x6f, 0xd0, 0x94, 0x42, 0x69, 0x9a, 0x46,
	0x26, 0x0f, 0x9b, 0xbc, 0xdf, 0x07, 0x07, 0x1d, 0x89, 0x04, 0x4d, 0xee, 0xb1, 0x18, 0x71, 0x9d,
	0x8f, 0xad, 0x22, 0xe2, 0xa3, 0x0c, 0x73, 0x16, 0x97, 0xb9, 0x17, 0x9c, 0x91, 0x9e, 0x3a, 0x66,
	0xef, 0x4b, 0x15, 0x16, 0x1f, 0x22, 0xcd, 0x14, 0x79, 0xeb, 0x41, 0xcd, 0xdd, 0x1c, 0xb2, 0x3d,
	0x2f, 0xf3, 0x2b, 0xb7, 0xab, 0xb5, 0x31, 0x95, 0xba, 0xdb, 0xd8, 0x7d, 0x2c, 0x58, 0x52, 0xe6,
	0x1a, 0xfe, 0x7f, 0xf9, 0x31, 0xa8, 0x34, 0xbc, 0xcb, 0xcf, 0x5f, 0xdf, 0x55, 0x76, 0xc2, 0xbf,
	0x7b, 0xd1, 0x95, 0x3b, 0x3e, 0x69, 0xde, 0x2b, 0x9a, 0xf7, 0xcc, 0x87, 0x44, 0xf5, 0x9e, 0x29,
	0xc1, 0xf7, 0xbd, 0xbf, 0xc8, 0x7b, 0x0f, 0xaa, 0xc7, 0xa8, 0xc9, 0x1f, 0xf3, 0xfc, 0x4c, 0x17,
	0xad, 0xf5, 0xe7, 0xad, 0xba, 0xc2, 0xd6, 0xdd, 0x19, 0x5b, 0xff, 0x91, 0xbd, 0xef, 0xb4, 0xf5,
	0xca, 0x10, 0xaf, 0xad, 0x3b, 0xf2, 0xc1, 0x83, 0xe6, 0xcc, 0x76, 0x92, 0xee, 0xbc, 0xd1, 0xd7,
	0xd7, 0xf8, 0xb6, 0xf7, 0x76, 0x32, 0x63, 0xf0, 0x80, 0xdc, 0xf9, 0x31, 0x83, 0x4c, 0xfd, 0x43,
	0xcb, 0x79, 0xd6, 0xed, 0xe1, 0x3e, 0x84, 0xb1, 0xc8, 0xa6, 0x23, 0xa9, 0x64, 0xd7, 0x6d, 0x1e,
	0xfa, 0x66, 0x11, 0xce, 0xcc, 0x67, 0xf4, 0xcc, 0x7b, 0x52, 0x2f, 0xd0, 0x41, 0xcd, 0x7e, 0x58,
	0xff, 0xfd, 0x16, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x38, 0x86, 0xd2, 0xe5, 0x05, 0x00, 0x00,
}