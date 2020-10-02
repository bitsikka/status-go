// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat_identity.proto

package protobuf

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// SourceType are the predefined types of image source allowed
type IdentityImage_SourceType int32

const (
	IdentityImage_UNKNOWN_SOURCE_TYPE IdentityImage_SourceType = 0
	// RAW_PAYLOAD image byte data
	IdentityImage_RAW_PAYLOAD IdentityImage_SourceType = 1
	// ENS_AVATAR uses the ENS record's resolver get-text-data.avatar data
	// The `payload` field will be ignored if ENS_AVATAR is selected
	// The application will read and parse the ENS avatar data as image payload data, URLs will be ignored
	// The parent `ChatMessageIdentity` must have a valid `ens_name` set
	IdentityImage_ENS_AVATAR IdentityImage_SourceType = 2
)

var IdentityImage_SourceType_name = map[int32]string{
	0: "UNKNOWN_SOURCE_TYPE",
	1: "RAW_PAYLOAD",
	2: "ENS_AVATAR",
}

var IdentityImage_SourceType_value = map[string]int32{
	"UNKNOWN_SOURCE_TYPE": 0,
	"RAW_PAYLOAD":         1,
	"ENS_AVATAR":          2,
}

func (x IdentityImage_SourceType) String() string {
	return proto.EnumName(IdentityImage_SourceType_name, int32(x))
}

func (IdentityImage_SourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7a652489000a5879, []int{1, 0}
}

// ChatIdentity represents the user defined identity associated with their public chat key
type ChatIdentity struct {
	// Lamport timestamp of the message
	Clock uint64 `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
	// ens_name is the valid ENS name associated with the chat key
	EnsName string `protobuf:"bytes,2,opt,name=ens_name,json=ensName,proto3" json:"ens_name,omitempty"`
	// images is a string indexed mapping of images associated with an identity
	Images               map[string]*IdentityImage `protobuf:"bytes,3,rep,name=images,proto3" json:"images,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ChatIdentity) Reset()         { *m = ChatIdentity{} }
func (m *ChatIdentity) String() string { return proto.CompactTextString(m) }
func (*ChatIdentity) ProtoMessage()    {}
func (*ChatIdentity) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a652489000a5879, []int{0}
}

func (m *ChatIdentity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatIdentity.Unmarshal(m, b)
}
func (m *ChatIdentity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatIdentity.Marshal(b, m, deterministic)
}
func (m *ChatIdentity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatIdentity.Merge(m, src)
}
func (m *ChatIdentity) XXX_Size() int {
	return xxx_messageInfo_ChatIdentity.Size(m)
}
func (m *ChatIdentity) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatIdentity.DiscardUnknown(m)
}

var xxx_messageInfo_ChatIdentity proto.InternalMessageInfo

func (m *ChatIdentity) GetClock() uint64 {
	if m != nil {
		return m.Clock
	}
	return 0
}

func (m *ChatIdentity) GetEnsName() string {
	if m != nil {
		return m.EnsName
	}
	return ""
}

func (m *ChatIdentity) GetImages() map[string]*IdentityImage {
	if m != nil {
		return m.Images
	}
	return nil
}

// ProfileImage represents data associated with a user's profile image
type IdentityImage struct {
	// payload is a context based payload for the profile image data,
	// context is determined by the `source_type`
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	// source_type signals the image payload source
	SourceType IdentityImage_SourceType `protobuf:"varint,2,opt,name=source_type,json=sourceType,proto3,enum=protobuf.IdentityImage_SourceType" json:"source_type,omitempty"`
	// image_type signals the image type and method of parsing the payload
	ImageType            ImageType `protobuf:"varint,3,opt,name=image_type,json=imageType,proto3,enum=protobuf.ImageType" json:"image_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *IdentityImage) Reset()         { *m = IdentityImage{} }
func (m *IdentityImage) String() string { return proto.CompactTextString(m) }
func (*IdentityImage) ProtoMessage()    {}
func (*IdentityImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a652489000a5879, []int{1}
}

func (m *IdentityImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentityImage.Unmarshal(m, b)
}
func (m *IdentityImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentityImage.Marshal(b, m, deterministic)
}
func (m *IdentityImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentityImage.Merge(m, src)
}
func (m *IdentityImage) XXX_Size() int {
	return xxx_messageInfo_IdentityImage.Size(m)
}
func (m *IdentityImage) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentityImage.DiscardUnknown(m)
}

var xxx_messageInfo_IdentityImage proto.InternalMessageInfo

func (m *IdentityImage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *IdentityImage) GetSourceType() IdentityImage_SourceType {
	if m != nil {
		return m.SourceType
	}
	return IdentityImage_UNKNOWN_SOURCE_TYPE
}

func (m *IdentityImage) GetImageType() ImageType {
	if m != nil {
		return m.ImageType
	}
	return ImageType_UNKNOWN_IMAGE_TYPE
}

func init() {
	proto.RegisterEnum("protobuf.IdentityImage_SourceType", IdentityImage_SourceType_name, IdentityImage_SourceType_value)
	proto.RegisterType((*ChatIdentity)(nil), "protobuf.ChatIdentity")
	proto.RegisterMapType((map[string]*IdentityImage)(nil), "protobuf.ChatIdentity.ImagesEntry")
	proto.RegisterType((*IdentityImage)(nil), "protobuf.IdentityImage")
}

func init() {
	proto.RegisterFile("chat_identity.proto", fileDescriptor_7a652489000a5879)
}

var fileDescriptor_7a652489000a5879 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xcb, 0x6e, 0xf2, 0x30,
	0x10, 0x85, 0x7f, 0x93, 0x9f, 0xdb, 0x84, 0xd2, 0xc8, 0x54, 0x22, 0x65, 0x85, 0xb2, 0x62, 0xd3,
	0x2c, 0xe8, 0xa6, 0x62, 0x17, 0xd1, 0x54, 0x42, 0xad, 0x02, 0x32, 0x50, 0xc4, 0xca, 0x32, 0xc1,
	0x2d, 0x11, 0xe4, 0x22, 0xe2, 0x54, 0xca, 0x53, 0xf6, 0x3d, 0xfa, 0x14, 0x15, 0x36, 0x29, 0x61,
	0xe5, 0xe3, 0x99, 0x73, 0xbe, 0xa3, 0x81, 0x8e, 0xbf, 0x63, 0x82, 0x06, 0x5b, 0x1e, 0x89, 0x40,
	0xe4, 0x76, 0x72, 0x8c, 0x45, 0x8c, 0x1b, 0xf2, 0xd9, 0x64, 0x1f, 0x3d, 0x9d, 0x47, 0x59, 0x98,
	0xaa, 0xb1, 0xf5, 0x8d, 0xa0, 0x35, 0xde, 0x31, 0x31, 0x39, 0xbb, 0xf1, 0x1d, 0x54, 0xfd, 0x43,
	0xec, 0xef, 0x4d, 0xd4, 0x47, 0x83, 0xff, 0x44, 0x7d, 0xf0, 0x3d, 0x34, 0x78, 0x94, 0xd2, 0x88,
	0x85, 0xdc, 0xac, 0xf4, 0xd1, 0xa0, 0x49, 0xea, 0x3c, 0x4a, 0x3d, 0x16, 0x72, 0x3c, 0x82, 0x5a,
	0x10, 0xb2, 0x4f, 0x9e, 0x9a, 0x5a, 0x5f, 0x1b, 0xe8, 0x43, 0xcb, 0x2e, 0x9a, 0xec, 0x32, 0xd8,
	0x9e, 0x48, 0x93, 0x1b, 0x89, 0x63, 0x4e, 0xce, 0x89, 0x1e, 0x01, 0xbd, 0x34, 0xc6, 0x06, 0x68,
	0x7b, 0x9e, 0xcb, 0xe6, 0x26, 0x39, 0x49, 0xfc, 0x00, 0xd5, 0x2f, 0x76, 0xc8, 0x54, 0xa9, 0x3e,
	0xec, 0x5e, 0xd8, 0x05, 0x57, 0xe6, 0x89, 0x72, 0x8d, 0x2a, 0x4f, 0xc8, 0xfa, 0x41, 0x70, 0x73,
	0xb5, 0xc4, 0x26, 0xd4, 0x13, 0x96, 0x1f, 0x62, 0xb6, 0x95, 0xe8, 0x16, 0x29, 0xbe, 0x78, 0x0c,
	0x7a, 0x1a, 0x67, 0x47, 0x9f, 0x53, 0x91, 0x27, 0xaa, 0xa4, 0x5d, 0x3e, 0xe0, 0x8a, 0x63, 0xcf,
	0xa5, 0x75, 0x91, 0x27, 0x9c, 0x40, 0xfa, 0xa7, 0xf1, 0x10, 0x40, 0x9e, 0xa3, 0x18, 0x9a, 0x64,
	0x74, 0x4a, 0x8c, 0xd3, 0x4e, 0x86, 0x9a, 0x41, 0x21, 0xad, 0x17, 0x80, 0x0b, 0x0d, 0x77, 0xa1,
	0xb3, 0xf4, 0x5e, 0xbd, 0xe9, 0xca, 0xa3, 0xf3, 0xe9, 0x92, 0x8c, 0x5d, 0xba, 0x58, 0xcf, 0x5c,
	0xe3, 0x1f, 0xbe, 0x05, 0x9d, 0x38, 0x2b, 0x3a, 0x73, 0xd6, 0x6f, 0x53, 0xe7, 0xd9, 0x40, 0xb8,
	0x0d, 0xe0, 0x7a, 0x73, 0xea, 0xbc, 0x3b, 0x0b, 0x87, 0x18, 0x95, 0x4d, 0x4d, 0xd6, 0x3c, 0xfe,
	0x06, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x30, 0xc2, 0x9f, 0xf3, 0x01, 0x00, 0x00,
}
