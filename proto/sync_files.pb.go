// Code generated by protoc-gen-go.
// source: sync_files.proto
// DO NOT EDIT!

/*
Package sync_proto is a generated protocol buffer package.

It is generated from these files:
	sync_files.proto

It has these top-level messages:
	SendFileInfo
	SyncFileReq
	SyncFileData
*/
package sync_proto

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

type SYNC int32

const (
	SYNC_Msg_SendFileInfo SYNC = 1
	SYNC_Msg_SyncFileReq  SYNC = 2
	SYNC_Msg_SyncFileData SYNC = 3
)

var SYNC_name = map[int32]string{
	1: "Msg_SendFileInfo",
	2: "Msg_SyncFileReq",
	3: "Msg_SyncFileData",
}
var SYNC_value = map[string]int32{
	"Msg_SendFileInfo": 1,
	"Msg_SyncFileReq":  2,
	"Msg_SyncFileData": 3,
}

func (x SYNC) Enum() *SYNC {
	p := new(SYNC)
	*p = x
	return p
}
func (x SYNC) String() string {
	return proto.EnumName(SYNC_name, int32(x))
}
func (x *SYNC) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SYNC_value, data, "SYNC")
	if err != nil {
		return err
	}
	*x = SYNC(value)
	return nil
}
func (SYNC) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// client -> Server Type : 0001
type SendFileInfo struct {
	FileName         *string `protobuf:"bytes,1,req,name=file_name" json:"file_name,omitempty"`
	HashCode         *uint64 `protobuf:"varint,2,req,name=hash_code" json:"hash_code,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SendFileInfo) Reset()                    { *m = SendFileInfo{} }
func (m *SendFileInfo) String() string            { return proto.CompactTextString(m) }
func (*SendFileInfo) ProtoMessage()               {}
func (*SendFileInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SendFileInfo) GetFileName() string {
	if m != nil && m.FileName != nil {
		return *m.FileName
	}
	return ""
}

func (m *SendFileInfo) GetHashCode() uint64 {
	if m != nil && m.HashCode != nil {
		return *m.HashCode
	}
	return 0
}

// server -> client Type : 8002
type SyncFileReq struct {
	FileName         *string `protobuf:"bytes,1,req,name=file_name" json:"file_name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SyncFileReq) Reset()                    { *m = SyncFileReq{} }
func (m *SyncFileReq) String() string            { return proto.CompactTextString(m) }
func (*SyncFileReq) ProtoMessage()               {}
func (*SyncFileReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SyncFileReq) GetFileName() string {
	if m != nil && m.FileName != nil {
		return *m.FileName
	}
	return ""
}

// client -> client Type : 0002
type SyncFileData struct {
	FileName         *string `protobuf:"bytes,1,req,name=file_name" json:"file_name,omitempty"`
	FileMode         *uint32 `protobuf:"varint,2,opt,name=file_mode" json:"file_mode,omitempty"`
	TotalPacks       *uint32 `protobuf:"varint,3,req,name=total_packs" json:"total_packs,omitempty"`
	CurrentPacks     *uint32 `protobuf:"varint,4,req,name=current_packs" json:"current_packs,omitempty"`
	Data             []byte  `protobuf:"bytes,5,req,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SyncFileData) Reset()                    { *m = SyncFileData{} }
func (m *SyncFileData) String() string            { return proto.CompactTextString(m) }
func (*SyncFileData) ProtoMessage()               {}
func (*SyncFileData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SyncFileData) GetFileName() string {
	if m != nil && m.FileName != nil {
		return *m.FileName
	}
	return ""
}

func (m *SyncFileData) GetFileMode() uint32 {
	if m != nil && m.FileMode != nil {
		return *m.FileMode
	}
	return 0
}

func (m *SyncFileData) GetTotalPacks() uint32 {
	if m != nil && m.TotalPacks != nil {
		return *m.TotalPacks
	}
	return 0
}

func (m *SyncFileData) GetCurrentPacks() uint32 {
	if m != nil && m.CurrentPacks != nil {
		return *m.CurrentPacks
	}
	return 0
}

func (m *SyncFileData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*SendFileInfo)(nil), "sync_proto.SendFileInfo")
	proto.RegisterType((*SyncFileReq)(nil), "sync_proto.SyncFileReq")
	proto.RegisterType((*SyncFileData)(nil), "sync_proto.SyncFileData")
	proto.RegisterEnum("sync_proto.SYNC", SYNC_name, SYNC_value)
}

func init() { proto.RegisterFile("sync_files.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0xbd, 0x4e, 0xc4, 0x30,
	0x10, 0x84, 0x15, 0x9f, 0x29, 0x6e, 0xcf, 0x11, 0xc6, 0x07, 0x92, 0x4b, 0x2b, 0x95, 0x45, 0x41,
	0xc5, 0x1b, 0x80, 0x40, 0x14, 0x50, 0x90, 0x8a, 0xca, 0x5a, 0x39, 0x3e, 0x12, 0x91, 0xd8, 0x21,
	0x36, 0x45, 0xde, 0x1e, 0xc5, 0xfc, 0xa5, 0xe0, 0xba, 0xd5, 0xa7, 0x6f, 0x46, 0xb3, 0xc0, 0xe3,
	0xec, 0xad, 0x39, 0x74, 0xbd, 0x8b, 0x57, 0xe3, 0x14, 0x52, 0x10, 0x90, 0x49, 0xbe, 0xab, 0x6b,
	0x60, 0xb5, 0xf3, 0xcd, 0x5d, 0xd7, 0xbb, 0x07, 0x7f, 0x08, 0xe2, 0x0c, 0xb6, 0x8b, 0x6a, 0x3c,
	0x0e, 0x4e, 0x16, 0x8a, 0xe8, 0xed, 0x82, 0x5a, 0x8c, 0xad, 0xb1, 0xa1, 0x71, 0x92, 0x28, 0xa2,
	0x69, 0xa5, 0x60, 0x57, 0xcf, 0xde, 0x2e, 0xa9, 0x67, 0xf7, 0xfe, 0x4f, 0xa8, 0xf2, 0xc0, 0x7e,
	0x8c, 0x5b, 0x4c, 0x78, 0xa4, 0x37, 0xa3, 0xe1, 0xab, 0xb7, 0xd0, 0xa5, 0xd8, 0xc3, 0x2e, 0x85,
	0x84, 0xbd, 0x19, 0xd1, 0xbe, 0x45, 0xb9, 0x51, 0x44, 0x97, 0xe2, 0x02, 0x4a, 0xfb, 0x31, 0x4d,
	0xce, 0xa7, 0x6f, 0x4c, 0x33, 0x66, 0x40, 0x1b, 0x4c, 0x28, 0x4f, 0x14, 0xd1, 0xec, 0xf2, 0x1e,
	0x68, 0xfd, 0xf2, 0x74, 0x23, 0xce, 0x81, 0x3f, 0xc6, 0x57, 0xb3, 0xfe, 0x89, 0x17, 0x62, 0x0f,
	0xa7, 0x99, 0xfe, 0x6d, 0xe6, 0xe4, 0x57, 0x5d, 0xcd, 0xe4, 0x9b, 0xcf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xe4, 0x7f, 0x6e, 0xfd, 0x2f, 0x01, 0x00, 0x00,
}
