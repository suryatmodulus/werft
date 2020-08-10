// Code generated by protoc-gen-go. DO NOT EDIT.
// source: repo-plugin.proto

package common

import (
	context "context"
	fmt "fmt"
	v1 "github.com/csweichel/werft/pkg/api/v1"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type RepoHostRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RepoHostRequest) Reset()         { *m = RepoHostRequest{} }
func (m *RepoHostRequest) String() string { return proto.CompactTextString(m) }
func (*RepoHostRequest) ProtoMessage()    {}
func (*RepoHostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d8b2585439eaf91, []int{0}
}

func (m *RepoHostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepoHostRequest.Unmarshal(m, b)
}
func (m *RepoHostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepoHostRequest.Marshal(b, m, deterministic)
}
func (m *RepoHostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepoHostRequest.Merge(m, src)
}
func (m *RepoHostRequest) XXX_Size() int {
	return xxx_messageInfo_RepoHostRequest.Size(m)
}
func (m *RepoHostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RepoHostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RepoHostRequest proto.InternalMessageInfo

type RepoHostResponse struct {
	Host                 string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RepoHostResponse) Reset()         { *m = RepoHostResponse{} }
func (m *RepoHostResponse) String() string { return proto.CompactTextString(m) }
func (*RepoHostResponse) ProtoMessage()    {}
func (*RepoHostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d8b2585439eaf91, []int{1}
}

func (m *RepoHostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepoHostResponse.Unmarshal(m, b)
}
func (m *RepoHostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepoHostResponse.Marshal(b, m, deterministic)
}
func (m *RepoHostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepoHostResponse.Merge(m, src)
}
func (m *RepoHostResponse) XXX_Size() int {
	return xxx_messageInfo_RepoHostResponse.Size(m)
}
func (m *RepoHostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RepoHostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RepoHostResponse proto.InternalMessageInfo

func (m *RepoHostResponse) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

type ResolveRequest struct {
	Repository           *v1.Repository `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ResolveRequest) Reset()         { *m = ResolveRequest{} }
func (m *ResolveRequest) String() string { return proto.CompactTextString(m) }
func (*ResolveRequest) ProtoMessage()    {}
func (*ResolveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d8b2585439eaf91, []int{2}
}

func (m *ResolveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResolveRequest.Unmarshal(m, b)
}
func (m *ResolveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResolveRequest.Marshal(b, m, deterministic)
}
func (m *ResolveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResolveRequest.Merge(m, src)
}
func (m *ResolveRequest) XXX_Size() int {
	return xxx_messageInfo_ResolveRequest.Size(m)
}
func (m *ResolveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResolveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResolveRequest proto.InternalMessageInfo

func (m *ResolveRequest) GetRepository() *v1.Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

type ResolveResponse struct {
	Repository           *v1.Repository `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ResolveResponse) Reset()         { *m = ResolveResponse{} }
func (m *ResolveResponse) String() string { return proto.CompactTextString(m) }
func (*ResolveResponse) ProtoMessage()    {}
func (*ResolveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d8b2585439eaf91, []int{3}
}

func (m *ResolveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResolveResponse.Unmarshal(m, b)
}
func (m *ResolveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResolveResponse.Marshal(b, m, deterministic)
}
func (m *ResolveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResolveResponse.Merge(m, src)
}
func (m *ResolveResponse) XXX_Size() int {
	return xxx_messageInfo_ResolveResponse.Size(m)
}
func (m *ResolveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResolveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResolveResponse proto.InternalMessageInfo

func (m *ResolveResponse) GetRepository() *v1.Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

type ContentInitContainerRequest struct {
	Repository           *v1.Repository `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ContentInitContainerRequest) Reset()         { *m = ContentInitContainerRequest{} }
func (m *ContentInitContainerRequest) String() string { return proto.CompactTextString(m) }
func (*ContentInitContainerRequest) ProtoMessage()    {}
func (*ContentInitContainerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d8b2585439eaf91, []int{4}
}

func (m *ContentInitContainerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentInitContainerRequest.Unmarshal(m, b)
}
func (m *ContentInitContainerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentInitContainerRequest.Marshal(b, m, deterministic)
}
func (m *ContentInitContainerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentInitContainerRequest.Merge(m, src)
}
func (m *ContentInitContainerRequest) XXX_Size() int {
	return xxx_messageInfo_ContentInitContainerRequest.Size(m)
}
func (m *ContentInitContainerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentInitContainerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContentInitContainerRequest proto.InternalMessageInfo

func (m *ContentInitContainerRequest) GetRepository() *v1.Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

type ContentInitContainerResponse struct {
	Container            []byte   `protobuf:"bytes,1,opt,name=container,proto3" json:"container,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContentInitContainerResponse) Reset()         { *m = ContentInitContainerResponse{} }
func (m *ContentInitContainerResponse) String() string { return proto.CompactTextString(m) }
func (*ContentInitContainerResponse) ProtoMessage()    {}
func (*ContentInitContainerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d8b2585439eaf91, []int{5}
}

func (m *ContentInitContainerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentInitContainerResponse.Unmarshal(m, b)
}
func (m *ContentInitContainerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentInitContainerResponse.Marshal(b, m, deterministic)
}
func (m *ContentInitContainerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentInitContainerResponse.Merge(m, src)
}
func (m *ContentInitContainerResponse) XXX_Size() int {
	return xxx_messageInfo_ContentInitContainerResponse.Size(m)
}
func (m *ContentInitContainerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentInitContainerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContentInitContainerResponse proto.InternalMessageInfo

func (m *ContentInitContainerResponse) GetContainer() []byte {
	if m != nil {
		return m.Container
	}
	return nil
}

type DownloadRequest struct {
	Repository           *v1.Repository `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	Path                 string         `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DownloadRequest) Reset()         { *m = DownloadRequest{} }
func (m *DownloadRequest) String() string { return proto.CompactTextString(m) }
func (*DownloadRequest) ProtoMessage()    {}
func (*DownloadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d8b2585439eaf91, []int{6}
}

func (m *DownloadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadRequest.Unmarshal(m, b)
}
func (m *DownloadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadRequest.Marshal(b, m, deterministic)
}
func (m *DownloadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadRequest.Merge(m, src)
}
func (m *DownloadRequest) XXX_Size() int {
	return xxx_messageInfo_DownloadRequest.Size(m)
}
func (m *DownloadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadRequest proto.InternalMessageInfo

func (m *DownloadRequest) GetRepository() *v1.Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *DownloadRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type DownloadResponse struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadResponse) Reset()         { *m = DownloadResponse{} }
func (m *DownloadResponse) String() string { return proto.CompactTextString(m) }
func (*DownloadResponse) ProtoMessage()    {}
func (*DownloadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d8b2585439eaf91, []int{7}
}

func (m *DownloadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadResponse.Unmarshal(m, b)
}
func (m *DownloadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadResponse.Marshal(b, m, deterministic)
}
func (m *DownloadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadResponse.Merge(m, src)
}
func (m *DownloadResponse) XXX_Size() int {
	return xxx_messageInfo_DownloadResponse.Size(m)
}
func (m *DownloadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadResponse proto.InternalMessageInfo

func (m *DownloadResponse) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type ListFilesRequest struct {
	Repository           *v1.Repository `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	Path                 string         `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListFilesRequest) Reset()         { *m = ListFilesRequest{} }
func (m *ListFilesRequest) String() string { return proto.CompactTextString(m) }
func (*ListFilesRequest) ProtoMessage()    {}
func (*ListFilesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d8b2585439eaf91, []int{8}
}

func (m *ListFilesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFilesRequest.Unmarshal(m, b)
}
func (m *ListFilesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFilesRequest.Marshal(b, m, deterministic)
}
func (m *ListFilesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFilesRequest.Merge(m, src)
}
func (m *ListFilesRequest) XXX_Size() int {
	return xxx_messageInfo_ListFilesRequest.Size(m)
}
func (m *ListFilesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFilesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListFilesRequest proto.InternalMessageInfo

func (m *ListFilesRequest) GetRepository() *v1.Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *ListFilesRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type ListFilesReponse struct {
	Paths                []string `protobuf:"bytes,1,rep,name=paths,proto3" json:"paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListFilesReponse) Reset()         { *m = ListFilesReponse{} }
func (m *ListFilesReponse) String() string { return proto.CompactTextString(m) }
func (*ListFilesReponse) ProtoMessage()    {}
func (*ListFilesReponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d8b2585439eaf91, []int{9}
}

func (m *ListFilesReponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFilesReponse.Unmarshal(m, b)
}
func (m *ListFilesReponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFilesReponse.Marshal(b, m, deterministic)
}
func (m *ListFilesReponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFilesReponse.Merge(m, src)
}
func (m *ListFilesReponse) XXX_Size() int {
	return xxx_messageInfo_ListFilesReponse.Size(m)
}
func (m *ListFilesReponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFilesReponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListFilesReponse proto.InternalMessageInfo

func (m *ListFilesReponse) GetPaths() []string {
	if m != nil {
		return m.Paths
	}
	return nil
}

type GetRemoteAnnotationsRequest struct {
	Repository           *v1.Repository `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetRemoteAnnotationsRequest) Reset()         { *m = GetRemoteAnnotationsRequest{} }
func (m *GetRemoteAnnotationsRequest) String() string { return proto.CompactTextString(m) }
func (*GetRemoteAnnotationsRequest) ProtoMessage()    {}
func (*GetRemoteAnnotationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d8b2585439eaf91, []int{10}
}

func (m *GetRemoteAnnotationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRemoteAnnotationsRequest.Unmarshal(m, b)
}
func (m *GetRemoteAnnotationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRemoteAnnotationsRequest.Marshal(b, m, deterministic)
}
func (m *GetRemoteAnnotationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRemoteAnnotationsRequest.Merge(m, src)
}
func (m *GetRemoteAnnotationsRequest) XXX_Size() int {
	return xxx_messageInfo_GetRemoteAnnotationsRequest.Size(m)
}
func (m *GetRemoteAnnotationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRemoteAnnotationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRemoteAnnotationsRequest proto.InternalMessageInfo

func (m *GetRemoteAnnotationsRequest) GetRepository() *v1.Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

type GetRemoteAnnotationsResponse struct {
	Annotations          map[string]string `protobuf:"bytes,1,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetRemoteAnnotationsResponse) Reset()         { *m = GetRemoteAnnotationsResponse{} }
func (m *GetRemoteAnnotationsResponse) String() string { return proto.CompactTextString(m) }
func (*GetRemoteAnnotationsResponse) ProtoMessage()    {}
func (*GetRemoteAnnotationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d8b2585439eaf91, []int{11}
}

func (m *GetRemoteAnnotationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRemoteAnnotationsResponse.Unmarshal(m, b)
}
func (m *GetRemoteAnnotationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRemoteAnnotationsResponse.Marshal(b, m, deterministic)
}
func (m *GetRemoteAnnotationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRemoteAnnotationsResponse.Merge(m, src)
}
func (m *GetRemoteAnnotationsResponse) XXX_Size() int {
	return xxx_messageInfo_GetRemoteAnnotationsResponse.Size(m)
}
func (m *GetRemoteAnnotationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRemoteAnnotationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetRemoteAnnotationsResponse proto.InternalMessageInfo

func (m *GetRemoteAnnotationsResponse) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func init() {
	proto.RegisterType((*RepoHostRequest)(nil), "repoplugin.RepoHostRequest")
	proto.RegisterType((*RepoHostResponse)(nil), "repoplugin.RepoHostResponse")
	proto.RegisterType((*ResolveRequest)(nil), "repoplugin.ResolveRequest")
	proto.RegisterType((*ResolveResponse)(nil), "repoplugin.ResolveResponse")
	proto.RegisterType((*ContentInitContainerRequest)(nil), "repoplugin.ContentInitContainerRequest")
	proto.RegisterType((*ContentInitContainerResponse)(nil), "repoplugin.ContentInitContainerResponse")
	proto.RegisterType((*DownloadRequest)(nil), "repoplugin.DownloadRequest")
	proto.RegisterType((*DownloadResponse)(nil), "repoplugin.DownloadResponse")
	proto.RegisterType((*ListFilesRequest)(nil), "repoplugin.ListFilesRequest")
	proto.RegisterType((*ListFilesReponse)(nil), "repoplugin.ListFilesReponse")
	proto.RegisterType((*GetRemoteAnnotationsRequest)(nil), "repoplugin.GetRemoteAnnotationsRequest")
	proto.RegisterType((*GetRemoteAnnotationsResponse)(nil), "repoplugin.GetRemoteAnnotationsResponse")
	proto.RegisterMapType((map[string]string)(nil), "repoplugin.GetRemoteAnnotationsResponse.AnnotationsEntry")
}

func init() {
	proto.RegisterFile("repo-plugin.proto", fileDescriptor_0d8b2585439eaf91)
}

var fileDescriptor_0d8b2585439eaf91 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x6d, 0xb6, 0xb1, 0xb5, 0xb7, 0x68, 0xcb, 0xac, 0x3d, 0x54, 0x69, 0x1f, 0x26, 0x3f, 0x40,
	0x1f, 0x20, 0xd3, 0xca, 0x0b, 0x20, 0x84, 0x18, 0x0c, 0xca, 0x24, 0x90, 0x90, 0x25, 0x78, 0x80,
	0xa7, 0x50, 0x0c, 0x8b, 0x96, 0xfa, 0x86, 0xf8, 0xb6, 0x55, 0x7f, 0x05, 0x3f, 0x88, 0x3f, 0x87,
	0x1c, 0xbb, 0xcd, 0x87, 0x42, 0x11, 0xd5, 0xde, 0xec, 0x7b, 0xaf, 0x8f, 0x8f, 0x4f, 0xce, 0x09,
	0x1c, 0x67, 0x32, 0xc5, 0x87, 0x69, 0x32, 0xfb, 0x11, 0xab, 0x30, 0xcd, 0x90, 0x90, 0x81, 0x29,
	0xd9, 0x4a, 0xc0, 0xa2, 0x34, 0x3e, 0x9b, 0x9f, 0x9f, 0x2d, 0x64, 0xf6, 0x9d, 0x6c, 0x9f, 0x1f,
	0xc3, 0x91, 0x90, 0x29, 0xbe, 0x45, 0x4d, 0x42, 0xfe, 0x9c, 0x49, 0x4d, 0xfc, 0x1e, 0xf8, 0x45,
	0x49, 0xa7, 0xa8, 0xb4, 0x64, 0x0c, 0xf6, 0xae, 0x51, 0x53, 0xcf, 0x3b, 0xf5, 0x86, 0x1d, 0x91,
	0xaf, 0xf9, 0x0b, 0x38, 0x14, 0x52, 0x63, 0x32, 0x97, 0xee, 0x24, 0x0b, 0x21, 0xbf, 0x4e, 0xc7,
	0x84, 0xd9, 0x32, 0x9f, 0xed, 0x8e, 0x0e, 0xc3, 0xf9, 0x79, 0x28, 0xd6, 0x55, 0x51, 0x9a, 0xe0,
	0x17, 0xe6, 0x72, 0x87, 0xe0, 0x2e, 0xfa, 0x5f, 0x88, 0xf7, 0xd0, 0x7f, 0x85, 0x8a, 0xa4, 0xa2,
	0x2b, 0x15, 0x93, 0x59, 0x46, 0xb1, 0x92, 0xd9, 0xb6, 0x8c, 0x9e, 0xc1, 0xa0, 0x19, 0xce, 0xd1,
	0x1b, 0x40, 0x67, 0xb2, 0x2a, 0xe6, 0x70, 0x77, 0x45, 0x51, 0xe0, 0x1f, 0xe1, 0xe8, 0x12, 0x17,
	0x2a, 0xc1, 0xe8, 0xdb, 0x96, 0x04, 0x8c, 0xd0, 0x69, 0x44, 0xd7, 0xbd, 0x1d, 0x2b, 0xb4, 0x59,
	0xf3, 0x07, 0xe0, 0x17, 0xb0, 0x8e, 0x48, 0x0f, 0x0e, 0x26, 0x96, 0xa8, 0xa3, 0xb1, 0xda, 0xf2,
	0x4f, 0xe0, 0xbf, 0x8b, 0x35, 0xbd, 0x89, 0x13, 0xa9, 0x6f, 0x93, 0xc5, 0xb0, 0x82, 0x6b, 0x59,
	0x9c, 0xc0, 0x1d, 0xd3, 0xd3, 0x3d, 0xef, 0x74, 0x77, 0xd8, 0x11, 0x76, 0x63, 0xbe, 0xc9, 0x58,
	0x92, 0x90, 0x53, 0x24, 0x79, 0xa1, 0x14, 0x52, 0x44, 0x31, 0xaa, 0x6d, 0xc9, 0xf0, 0xdf, 0x1e,
	0x0c, 0x9a, 0xf1, 0x9c, 0x16, 0x5f, 0xa0, 0x1b, 0x15, 0xe5, 0x9c, 0x4b, 0x77, 0xf4, 0x24, 0x2c,
	0x9c, 0x1f, 0x6e, 0x3a, 0x1e, 0x96, 0x6a, 0xaf, 0x15, 0x65, 0x4b, 0x51, 0x46, 0x0b, 0x9e, 0x83,
	0x5f, 0x1f, 0x60, 0x3e, 0xec, 0xde, 0xc8, 0xa5, 0x0b, 0x83, 0x59, 0x1a, 0x21, 0xe6, 0x51, 0x32,
	0x93, 0x4e, 0x31, 0xbb, 0x79, 0xba, 0xf3, 0xd8, 0x1b, 0xfd, 0xda, 0xb3, 0x71, 0xb2, 0x8f, 0xf9,
	0x90, 0xf3, 0x61, 0x63, 0x68, 0xaf, 0x22, 0xc6, 0xfa, 0x65, 0xa2, 0xb5, 0x2c, 0x06, 0x83, 0xe6,
	0xa6, 0x65, 0xce, 0x5b, 0xec, 0x12, 0x0e, 0x5c, 0x82, 0x58, 0x50, 0x1d, 0x2d, 0x07, 0x33, 0xe8,
	0x37, 0xf6, 0xd6, 0x28, 0x37, 0x70, 0xd2, 0xe4, 0x7a, 0x76, 0xbf, 0x7c, 0x6c, 0x43, 0xcc, 0x82,
	0xe1, 0xbf, 0x07, 0xd7, 0x97, 0x8d, 0xa1, 0xbd, 0x72, 0x73, 0xf5, 0xed, 0xb5, 0xe8, 0x54, 0xdf,
	0x5e, 0x0f, 0x00, 0x6f, 0xb1, 0x2b, 0xe8, 0xac, 0x0d, 0xc9, 0x2a, 0xc3, 0x75, 0xff, 0x07, 0x7f,
	0xeb, 0x96, 0x04, 0x68, 0xb2, 0x48, 0x55, 0x80, 0x0d, 0x9e, 0xae, 0x0a, 0xb0, 0xc9, 0x6d, 0xbc,
	0xf5, 0xb2, 0xfd, 0x79, 0x7f, 0x82, 0xd3, 0x29, 0xaa, 0xaf, 0xfb, 0xf9, 0x3f, 0xf8, 0xd1, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x07, 0x33, 0xf1, 0xb8, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RepositoryPluginClient is the client API for RepositoryPlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RepositoryPluginClient interface {
	// RepoHost returns the host which this plugins integrates with
	RepoHost(ctx context.Context, in *RepoHostRequest, opts ...grpc.CallOption) (*RepoHostResponse, error)
	// Resolve resolves the repo's revision based on its ref(erence)
	Resolve(ctx context.Context, in *ResolveRequest, opts ...grpc.CallOption) (*ResolveResponse, error)
	// ContentInitContainer produces the init container YAML required to initialize
	// the build context from this repository in /workspace.
	ContentInitContainer(ctx context.Context, in *ContentInitContainerRequest, opts ...grpc.CallOption) (*ContentInitContainerResponse, error)
	// Download downloads a file from the repository.
	Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*DownloadResponse, error)
	// ListFiles lists all files in a directory.
	ListFiles(ctx context.Context, in *ListFilesRequest, opts ...grpc.CallOption) (*ListFilesReponse, error)
	// GetRemoteAnnotations extracts werft annotations form information associated
	// with a particular commit, e.g. the commit message, PRs or merge requests.
	// Implementors can expect the revision of the repo object to be set.
	GetRemoteAnnotations(ctx context.Context, in *GetRemoteAnnotationsRequest, opts ...grpc.CallOption) (*GetRemoteAnnotationsResponse, error)
}

type repositoryPluginClient struct {
	cc grpc.ClientConnInterface
}

func NewRepositoryPluginClient(cc grpc.ClientConnInterface) RepositoryPluginClient {
	return &repositoryPluginClient{cc}
}

func (c *repositoryPluginClient) RepoHost(ctx context.Context, in *RepoHostRequest, opts ...grpc.CallOption) (*RepoHostResponse, error) {
	out := new(RepoHostResponse)
	err := c.cc.Invoke(ctx, "/repoplugin.RepositoryPlugin/RepoHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryPluginClient) Resolve(ctx context.Context, in *ResolveRequest, opts ...grpc.CallOption) (*ResolveResponse, error) {
	out := new(ResolveResponse)
	err := c.cc.Invoke(ctx, "/repoplugin.RepositoryPlugin/Resolve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryPluginClient) ContentInitContainer(ctx context.Context, in *ContentInitContainerRequest, opts ...grpc.CallOption) (*ContentInitContainerResponse, error) {
	out := new(ContentInitContainerResponse)
	err := c.cc.Invoke(ctx, "/repoplugin.RepositoryPlugin/ContentInitContainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryPluginClient) Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*DownloadResponse, error) {
	out := new(DownloadResponse)
	err := c.cc.Invoke(ctx, "/repoplugin.RepositoryPlugin/Download", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryPluginClient) ListFiles(ctx context.Context, in *ListFilesRequest, opts ...grpc.CallOption) (*ListFilesReponse, error) {
	out := new(ListFilesReponse)
	err := c.cc.Invoke(ctx, "/repoplugin.RepositoryPlugin/ListFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryPluginClient) GetRemoteAnnotations(ctx context.Context, in *GetRemoteAnnotationsRequest, opts ...grpc.CallOption) (*GetRemoteAnnotationsResponse, error) {
	out := new(GetRemoteAnnotationsResponse)
	err := c.cc.Invoke(ctx, "/repoplugin.RepositoryPlugin/GetRemoteAnnotations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RepositoryPluginServer is the server API for RepositoryPlugin service.
type RepositoryPluginServer interface {
	// RepoHost returns the host which this plugins integrates with
	RepoHost(context.Context, *RepoHostRequest) (*RepoHostResponse, error)
	// Resolve resolves the repo's revision based on its ref(erence)
	Resolve(context.Context, *ResolveRequest) (*ResolveResponse, error)
	// ContentInitContainer produces the init container YAML required to initialize
	// the build context from this repository in /workspace.
	ContentInitContainer(context.Context, *ContentInitContainerRequest) (*ContentInitContainerResponse, error)
	// Download downloads a file from the repository.
	Download(context.Context, *DownloadRequest) (*DownloadResponse, error)
	// ListFiles lists all files in a directory.
	ListFiles(context.Context, *ListFilesRequest) (*ListFilesReponse, error)
	// GetRemoteAnnotations extracts werft annotations form information associated
	// with a particular commit, e.g. the commit message, PRs or merge requests.
	// Implementors can expect the revision of the repo object to be set.
	GetRemoteAnnotations(context.Context, *GetRemoteAnnotationsRequest) (*GetRemoteAnnotationsResponse, error)
}

// UnimplementedRepositoryPluginServer can be embedded to have forward compatible implementations.
type UnimplementedRepositoryPluginServer struct {
}

func (*UnimplementedRepositoryPluginServer) RepoHost(ctx context.Context, req *RepoHostRequest) (*RepoHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RepoHost not implemented")
}
func (*UnimplementedRepositoryPluginServer) Resolve(ctx context.Context, req *ResolveRequest) (*ResolveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Resolve not implemented")
}
func (*UnimplementedRepositoryPluginServer) ContentInitContainer(ctx context.Context, req *ContentInitContainerRequest) (*ContentInitContainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContentInitContainer not implemented")
}
func (*UnimplementedRepositoryPluginServer) Download(ctx context.Context, req *DownloadRequest) (*DownloadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Download not implemented")
}
func (*UnimplementedRepositoryPluginServer) ListFiles(ctx context.Context, req *ListFilesRequest) (*ListFilesReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFiles not implemented")
}
func (*UnimplementedRepositoryPluginServer) GetRemoteAnnotations(ctx context.Context, req *GetRemoteAnnotationsRequest) (*GetRemoteAnnotationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRemoteAnnotations not implemented")
}

func RegisterRepositoryPluginServer(s *grpc.Server, srv RepositoryPluginServer) {
	s.RegisterService(&_RepositoryPlugin_serviceDesc, srv)
}

func _RepositoryPlugin_RepoHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepoHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryPluginServer).RepoHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repoplugin.RepositoryPlugin/RepoHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryPluginServer).RepoHost(ctx, req.(*RepoHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryPlugin_Resolve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryPluginServer).Resolve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repoplugin.RepositoryPlugin/Resolve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryPluginServer).Resolve(ctx, req.(*ResolveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryPlugin_ContentInitContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContentInitContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryPluginServer).ContentInitContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repoplugin.RepositoryPlugin/ContentInitContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryPluginServer).ContentInitContainer(ctx, req.(*ContentInitContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryPlugin_Download_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryPluginServer).Download(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repoplugin.RepositoryPlugin/Download",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryPluginServer).Download(ctx, req.(*DownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryPlugin_ListFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryPluginServer).ListFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repoplugin.RepositoryPlugin/ListFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryPluginServer).ListFiles(ctx, req.(*ListFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryPlugin_GetRemoteAnnotations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRemoteAnnotationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryPluginServer).GetRemoteAnnotations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repoplugin.RepositoryPlugin/GetRemoteAnnotations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryPluginServer).GetRemoteAnnotations(ctx, req.(*GetRemoteAnnotationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RepositoryPlugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "repoplugin.RepositoryPlugin",
	HandlerType: (*RepositoryPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RepoHost",
			Handler:    _RepositoryPlugin_RepoHost_Handler,
		},
		{
			MethodName: "Resolve",
			Handler:    _RepositoryPlugin_Resolve_Handler,
		},
		{
			MethodName: "ContentInitContainer",
			Handler:    _RepositoryPlugin_ContentInitContainer_Handler,
		},
		{
			MethodName: "Download",
			Handler:    _RepositoryPlugin_Download_Handler,
		},
		{
			MethodName: "ListFiles",
			Handler:    _RepositoryPlugin_ListFiles_Handler,
		},
		{
			MethodName: "GetRemoteAnnotations",
			Handler:    _RepositoryPlugin_GetRemoteAnnotations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "repo-plugin.proto",
}
