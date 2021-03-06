// Code generated by protoc-gen-go. DO NOT EDIT.
// source: RequestProtoQuote.proto

/*
Package RequestProtoQuote is a generated protocol buffer package.

It is generated from these files:
	RequestProtoQuote.proto

It has these top-level messages:
	RequestProtoQuote
	ProtoResponse
*/
package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RequestProtoQuote struct {
	ID              string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	AccountID       string `protobuf:"bytes,2,opt,name=AccountID" json:"AccountID,omitempty"`
	ProductNumber   string `protobuf:"bytes,3,opt,name=ProductNumber" json:"ProductNumber,omitempty"`
	ProductCategory string `protobuf:"bytes,4,opt,name=ProductCategory" json:"ProductCategory,omitempty"`
	Quantity        string `protobuf:"bytes,5,opt,name=Quantity" json:"Quantity,omitempty"`
}

func (m *RequestProtoQuote) Reset()                    { *m = RequestProtoQuote{} }
func (m *RequestProtoQuote) String() string            { return proto.CompactTextString(m) }
func (*RequestProtoQuote) ProtoMessage()               {}
func (*RequestProtoQuote) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RequestProtoQuote) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *RequestProtoQuote) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *RequestProtoQuote) GetProductNumber() string {
	if m != nil {
		return m.ProductNumber
	}
	return ""
}

func (m *RequestProtoQuote) GetProductCategory() string {
	if m != nil {
		return m.ProductCategory
	}
	return ""
}

func (m *RequestProtoQuote) GetQuantity() string {
	if m != nil {
		return m.Quantity
	}
	return ""
}

type ProtoResponse struct {
	Message               string `protobuf:"bytes,1,opt,name=Message" json:"Message,omitempty"`
	UnitPrice             string `protobuf:"bytes,2,opt,name=UnitPrice" json:"UnitPrice,omitempty"`
	PriceValidationPeriod string `protobuf:"bytes,3,opt,name=PriceValidationPeriod" json:"PriceValidationPeriod,omitempty"`
}

func (m *ProtoResponse) Reset()                    { *m = ProtoResponse{} }
func (m *ProtoResponse) String() string            { return proto.CompactTextString(m) }
func (*ProtoResponse) ProtoMessage()               {}
func (*ProtoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ProtoResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ProtoResponse) GetUnitPrice() string {
	if m != nil {
		return m.UnitPrice
	}
	return ""
}

func (m *ProtoResponse) GetPriceValidationPeriod() string {
	if m != nil {
		return m.PriceValidationPeriod
	}
	return ""
}

func init() {
	proto.RegisterType((*RequestProtoQuote)(nil), "RequestProtoQuote")
	proto.RegisterType((*ProtoResponse)(nil), "ProtoResponse")
}

func init() { proto.RegisterFile("RequestProtoQuote.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbd, 0x4a, 0xc6, 0x30,
	0x14, 0x86, 0x69, 0xfd, 0xfd, 0x0e, 0xa8, 0x18, 0x10, 0x83, 0x38, 0xc8, 0x87, 0x43, 0x27, 0x17,
	0xbd, 0x01, 0xb1, 0x4b, 0x07, 0x25, 0x2d, 0xe8, 0x9e, 0xa6, 0x87, 0x12, 0xd0, 0x9c, 0x9a, 0x9c,
	0x0c, 0x1d, 0xbc, 0x24, 0xef, 0x51, 0x1a, 0xda, 0x8a, 0x3f, 0xdb, 0x79, 0x9f, 0xf7, 0x85, 0x3c,
	0x04, 0xce, 0x1b, 0x7c, 0x8f, 0x18, 0x58, 0x79, 0x62, 0xaa, 0x23, 0x31, 0xde, 0x0c, 0xd3, 0xb9,
	0xfd, 0xcc, 0xe0, 0xf4, 0x4f, 0x27, 0x8e, 0x21, 0xaf, 0x4a, 0x99, 0x5d, 0x65, 0xc5, 0xa6, 0xc9,
	0xab, 0x52, 0x5c, 0xc2, 0xe6, 0xde, 0x18, 0x8a, 0x8e, 0xab, 0x52, 0xe6, 0x09, 0x7f, 0x03, 0x71,
	0x0d, 0x47, 0xca, 0x53, 0x17, 0x0d, 0x3f, 0xc5, 0xb7, 0x16, 0xbd, 0xdc, 0x49, 0x8b, 0x9f, 0x50,
	0x14, 0x70, 0x32, 0x83, 0x07, 0xcd, 0xd8, 0x93, 0x1f, 0xe5, 0x6e, 0xda, 0xfd, 0xc6, 0xe2, 0x02,
	0x0e, 0xeb, 0xa8, 0x1d, 0x5b, 0x1e, 0xe5, 0x5e, 0x9a, 0xac, 0x79, 0xfb, 0x91, 0xde, 0x62, 0x6a,
	0x30, 0x0c, 0xe4, 0x02, 0x0a, 0x09, 0x07, 0x8f, 0x18, 0x82, 0xee, 0x71, 0xf6, 0x5d, 0xe2, 0x24,
	0xfd, 0xec, 0x2c, 0x2b, 0x6f, 0x0d, 0x2e, 0xd2, 0x2b, 0x10, 0x77, 0x70, 0x96, 0x8e, 0x17, 0xfd,
	0x6a, 0x3b, 0xcd, 0x96, 0x9c, 0x42, 0x6f, 0xa9, 0x9b, 0xe5, 0xff, 0x2f, 0xdb, 0xfd, 0xf4, 0x6b,
	0xb7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x69, 0x49, 0x11, 0xbb, 0x50, 0x01, 0x00, 0x00,
}
