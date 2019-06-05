// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: reservationservice/proto/reservation/reservation.proto

package reservation

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Message struct {
	Say string `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
}

func (m *Message) Reset()      { *m = Message{} }
func (*Message) ProtoMessage() {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9340195c3dda32e, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return m.Size()
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type Request struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *Request) Reset()      { *m = Request{} }
func (*Request) ProtoMessage() {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9340195c3dda32e, []int{1}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return m.Size()
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *Response) Reset()      { *m = Response{} }
func (*Response) ProtoMessage() {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9340195c3dda32e, []int{2}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return m.Size()
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type StreamingRequest struct {
	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (m *StreamingRequest) Reset()      { *m = StreamingRequest{} }
func (*StreamingRequest) ProtoMessage() {}
func (*StreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9340195c3dda32e, []int{3}
}
func (m *StreamingRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StreamingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StreamingRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StreamingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingRequest.Merge(m, src)
}
func (m *StreamingRequest) XXX_Size() int {
	return m.Size()
}
func (m *StreamingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingRequest proto.InternalMessageInfo

func (m *StreamingRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type StreamingResponse struct {
	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (m *StreamingResponse) Reset()      { *m = StreamingResponse{} }
func (*StreamingResponse) ProtoMessage() {}
func (*StreamingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9340195c3dda32e, []int{4}
}
func (m *StreamingResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StreamingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StreamingResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StreamingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingResponse.Merge(m, src)
}
func (m *StreamingResponse) XXX_Size() int {
	return m.Size()
}
func (m *StreamingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingResponse proto.InternalMessageInfo

func (m *StreamingResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Ping struct {
	Stroke int64 `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
}

func (m *Ping) Reset()      { *m = Ping{} }
func (*Ping) ProtoMessage() {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9340195c3dda32e, []int{5}
}
func (m *Ping) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return m.Size()
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

type Pong struct {
	Stroke int64 `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
}

func (m *Pong) Reset()      { *m = Pong{} }
func (*Pong) ProtoMessage() {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9340195c3dda32e, []int{6}
}
func (m *Pong) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return m.Size()
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "reservation.Message")
	proto.RegisterType((*Request)(nil), "reservation.Request")
	proto.RegisterType((*Response)(nil), "reservation.Response")
	proto.RegisterType((*StreamingRequest)(nil), "reservation.StreamingRequest")
	proto.RegisterType((*StreamingResponse)(nil), "reservation.StreamingResponse")
	proto.RegisterType((*Ping)(nil), "reservation.Ping")
	proto.RegisterType((*Pong)(nil), "reservation.Pong")
}

func init() {
	proto.RegisterFile("reservationservice/proto/reservation/reservation.proto", fileDescriptor_f9340195c3dda32e)
}

var fileDescriptor_f9340195c3dda32e = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x4e, 0xc2, 0x40,
	0x14, 0xc6, 0xe7, 0x05, 0x04, 0x7c, 0x6e, 0x60, 0x82, 0xc6, 0xa0, 0xbc, 0x98, 0x59, 0xe1, 0x06,
	0x88, 0x7f, 0x38, 0x80, 0xc6, 0x95, 0x31, 0x31, 0xf5, 0x04, 0x23, 0x99, 0x34, 0x8d, 0xb4, 0x83,
	0x9d, 0x62, 0x74, 0xe7, 0x11, 0x3c, 0x86, 0x17, 0x31, 0x71, 0xc9, 0x92, 0xa5, 0x0c, 0x1b, 0x97,
	0x1c, 0xc1, 0x74, 0x28, 0x49, 0x69, 0x64, 0xf7, 0xcd, 0xf7, 0xfb, 0xe6, 0xf5, 0x7b, 0xcd, 0xe0,
	0x20, 0x56, 0x46, 0xc5, 0x2f, 0x32, 0x09, 0x74, 0x94, 0x8a, 0x60, 0xa8, 0x7a, 0xe3, 0x58, 0x27,
	0xba, 0x97, 0x03, 0x79, 0xdd, 0x75, 0x94, 0xef, 0xe5, 0x2c, 0x71, 0x84, 0xd5, 0x3b, 0x65, 0x8c,
	0xf4, 0x15, 0xaf, 0x63, 0xc9, 0xc8, 0xb7, 0x43, 0x38, 0x81, 0xce, 0xae, 0x97, 0x4a, 0xd1, 0xc6,
	0xaa, 0xa7, 0x9e, 0x27, 0xca, 0x24, 0x9c, 0x63, 0x39, 0x92, 0xa1, 0xca, 0xa8, 0xd3, 0xe2, 0x18,
	0x6b, 0x9e, 0x32, 0xe3, 0xf4, 0xeb, 0xe9, 0xe5, 0xd0, 0xf8, 0xeb, 0xcb, 0xa1, 0xf1, 0x45, 0x07,
	0xeb, 0x0f, 0x49, 0xac, 0x64, 0x18, 0x44, 0xfe, 0x7a, 0x4a, 0x13, 0x77, 0x86, 0x7a, 0x12, 0x25,
	0x2e, 0x57, 0xf2, 0x56, 0x07, 0x71, 0x8a, 0x8d, 0x5c, 0x32, 0x1b, 0xf8, 0x7f, 0x94, 0xb0, 0x7c,
	0x1f, 0x44, 0x3e, 0x3f, 0xc0, 0x8a, 0x49, 0x62, 0xfd, 0xa4, 0x32, 0x9c, 0x9d, 0x1c, 0xd7, 0xdb,
	0xf9, 0xd9, 0x17, 0x60, 0xf5, 0xe6, 0x55, 0x86, 0xe3, 0x91, 0xe2, 0x97, 0x58, 0xbe, 0x96, 0xa3,
	0x11, 0x6f, 0x76, 0xf3, 0xff, 0x28, 0xab, 0xda, 0xda, 0x2f, 0xb8, 0xab, 0x5a, 0x82, 0xf1, 0x5b,
	0xac, 0xac, 0xda, 0xf2, 0xf6, 0x46, 0xa4, 0xb8, 0x6c, 0x8b, 0xb6, 0xe1, 0xf5, 0xa8, 0x3e, 0xf0,
	0x01, 0xd6, 0xd2, 0x7d, 0x5c, 0xe7, 0xc6, 0x46, 0x3e, 0xb5, 0x5b, 0x05, 0x4b, 0x47, 0xbe, 0x60,
	0x1d, 0xe8, 0xc3, 0xd5, 0xc5, 0x74, 0x4e, 0x6c, 0x36, 0x27, 0xb6, 0x9c, 0x13, 0xbc, 0x5b, 0x82,
	0x4f, 0x4b, 0xf0, 0x6d, 0x09, 0xa6, 0x96, 0xe0, 0xc7, 0x12, 0xfc, 0x5a, 0x62, 0x4b, 0x4b, 0xf0,
	0xb1, 0x20, 0x36, 0x5d, 0x10, 0x9b, 0x2d, 0x88, 0x3d, 0x56, 0xdc, 0x03, 0x38, 0xff, 0x0b, 0x00,
	0x00, 0xff, 0xff, 0xb5, 0xe1, 0x99, 0x3a, 0x3a, 0x02, 0x00, 0x00,
}

func (this *Message) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Message)
	if !ok {
		that2, ok := that.(Message)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Say != that1.Say {
		return false
	}
	return true
}
func (this *Request) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Request)
	if !ok {
		that2, ok := that.(Request)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	return true
}
func (this *Response) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Response)
	if !ok {
		that2, ok := that.(Response)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Msg != that1.Msg {
		return false
	}
	return true
}
func (this *StreamingRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StreamingRequest)
	if !ok {
		that2, ok := that.(StreamingRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Count != that1.Count {
		return false
	}
	return true
}
func (this *StreamingResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StreamingResponse)
	if !ok {
		that2, ok := that.(StreamingResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Count != that1.Count {
		return false
	}
	return true
}
func (this *Ping) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Ping)
	if !ok {
		that2, ok := that.(Ping)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Stroke != that1.Stroke {
		return false
	}
	return true
}
func (this *Pong) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Pong)
	if !ok {
		that2, ok := that.(Pong)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Stroke != that1.Stroke {
		return false
	}
	return true
}
func (this *Message) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&reservation.Message{")
	s = append(s, "Say: "+fmt.Sprintf("%#v", this.Say)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Request) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&reservation.Request{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Response) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&reservation.Response{")
	s = append(s, "Msg: "+fmt.Sprintf("%#v", this.Msg)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *StreamingRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&reservation.StreamingRequest{")
	s = append(s, "Count: "+fmt.Sprintf("%#v", this.Count)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *StreamingResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&reservation.StreamingResponse{")
	s = append(s, "Count: "+fmt.Sprintf("%#v", this.Count)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Ping) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&reservation.Ping{")
	s = append(s, "Stroke: "+fmt.Sprintf("%#v", this.Stroke)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Pong) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&reservation.Pong{")
	s = append(s, "Stroke: "+fmt.Sprintf("%#v", this.Stroke)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringReservation(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Say) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintReservation(dAtA, i, uint64(len(m.Say)))
		i += copy(dAtA[i:], m.Say)
	}
	return i, nil
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintReservation(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}

func (m *Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Response) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Msg) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintReservation(dAtA, i, uint64(len(m.Msg)))
		i += copy(dAtA[i:], m.Msg)
	}
	return i, nil
}

func (m *StreamingRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamingRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Count != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintReservation(dAtA, i, uint64(m.Count))
	}
	return i, nil
}

func (m *StreamingResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamingResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Count != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintReservation(dAtA, i, uint64(m.Count))
	}
	return i, nil
}

func (m *Ping) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ping) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Stroke != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintReservation(dAtA, i, uint64(m.Stroke))
	}
	return i, nil
}

func (m *Pong) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pong) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Stroke != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintReservation(dAtA, i, uint64(m.Stroke))
	}
	return i, nil
}

func encodeVarintReservation(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Say)
	if l > 0 {
		n += 1 + l + sovReservation(uint64(l))
	}
	return n
}

func (m *Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovReservation(uint64(l))
	}
	return n
}

func (m *Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovReservation(uint64(l))
	}
	return n
}

func (m *StreamingRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Count != 0 {
		n += 1 + sovReservation(uint64(m.Count))
	}
	return n
}

func (m *StreamingResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Count != 0 {
		n += 1 + sovReservation(uint64(m.Count))
	}
	return n
}

func (m *Ping) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Stroke != 0 {
		n += 1 + sovReservation(uint64(m.Stroke))
	}
	return n
}

func (m *Pong) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Stroke != 0 {
		n += 1 + sovReservation(uint64(m.Stroke))
	}
	return n
}

func sovReservation(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozReservation(x uint64) (n int) {
	return sovReservation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Message) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Message{`,
		`Say:` + fmt.Sprintf("%v", this.Say) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Request) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Request{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Response) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Response{`,
		`Msg:` + fmt.Sprintf("%v", this.Msg) + `,`,
		`}`,
	}, "")
	return s
}
func (this *StreamingRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StreamingRequest{`,
		`Count:` + fmt.Sprintf("%v", this.Count) + `,`,
		`}`,
	}, "")
	return s
}
func (this *StreamingResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StreamingResponse{`,
		`Count:` + fmt.Sprintf("%v", this.Count) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Ping) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Ping{`,
		`Stroke:` + fmt.Sprintf("%v", this.Stroke) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Pong) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Pong{`,
		`Stroke:` + fmt.Sprintf("%v", this.Stroke) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringReservation(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReservation
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Say", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReservation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReservation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReservation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Say = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReservation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReservation
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthReservation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReservation
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReservation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReservation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReservation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReservation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReservation
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthReservation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReservation
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReservation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReservation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReservation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReservation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReservation
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthReservation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StreamingRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReservation
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StreamingRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StreamingRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReservation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipReservation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReservation
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthReservation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StreamingResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReservation
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StreamingResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StreamingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReservation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipReservation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReservation
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthReservation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Ping) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReservation
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Ping: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ping: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stroke", wireType)
			}
			m.Stroke = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReservation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Stroke |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipReservation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReservation
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthReservation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Pong) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReservation
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Pong: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pong: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stroke", wireType)
			}
			m.Stroke = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReservation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Stroke |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipReservation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReservation
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthReservation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipReservation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowReservation
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowReservation
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowReservation
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthReservation
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthReservation
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowReservation
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipReservation(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthReservation
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthReservation = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowReservation   = fmt.Errorf("proto: integer overflow")
)
