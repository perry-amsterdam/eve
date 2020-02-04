// Code generated by protoc-gen-go. DO NOT EDIT.
// source: devconfig.proto

package config

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

// This is the response to a GET /api/v1/edgeDevice/config
// The EdgeDevConfig message carries all of the device's configuration from
// the controller to the device.
// The device will request these messages either periodically or as a result
// of some TBD notification.
// The message is assumed to be protected by a TLS session bound to the
// device certificate.
type EdgeDevConfig struct {
	Id          *UUIDandVersion      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Apps        []*AppInstanceConfig `protobuf:"bytes,4,rep,name=apps,proto3" json:"apps,omitempty"`
	Networks    []*NetworkConfig     `protobuf:"bytes,5,rep,name=networks,proto3" json:"networks,omitempty"`
	Datastores  []*DatastoreConfig   `protobuf:"bytes,6,rep,name=datastores,proto3" json:"datastores,omitempty"`
	LispInfo    *DeviceLispDetails   `protobuf:"bytes,7,opt,name=lispInfo,proto3" json:"lispInfo,omitempty"`
	Base        []*BaseOSConfig      `protobuf:"bytes,8,rep,name=base,proto3" json:"base,omitempty"`
	Reboot      *DeviceOpsCmd        `protobuf:"bytes,9,opt,name=reboot,proto3" json:"reboot,omitempty"`
	Backup      *DeviceOpsCmd        `protobuf:"bytes,10,opt,name=backup,proto3" json:"backup,omitempty"`
	ConfigItems []*ConfigItem        `protobuf:"bytes,11,rep,name=configItems,proto3" json:"configItems,omitempty"`
	// systemAdapterList - List of DeviceNetworkAdapters. Only Network
	//  adapters ( Ex: eth0, wlan1 etc ) have a corresponding SystemAdapter.
	// non-Network adapters do not have systemadapters.
	SystemAdapterList []*SystemAdapter `protobuf:"bytes,12,rep,name=systemAdapterList,proto3" json:"systemAdapterList,omitempty"`
	// deviceIoList - List of Physical Adapters. Includes both Network
	//  Adapters and Non-Network Adapters ( USB / Com etc )
	DeviceIoList []*PhysicalIO `protobuf:"bytes,13,rep,name=deviceIoList,proto3" json:"deviceIoList,omitempty"`
	// Override dmidecode info if set
	Manufacturer     string                   `protobuf:"bytes,14,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	ProductName      string                   `protobuf:"bytes,15,opt,name=productName,proto3" json:"productName,omitempty"`
	NetworkInstances []*NetworkInstanceConfig `protobuf:"bytes,16,rep,name=networkInstances,proto3" json:"networkInstances,omitempty"`
	// Information saved by device to make it easier to find in the controller
	Enterprise           string   `protobuf:"bytes,17,opt,name=enterprise,proto3" json:"enterprise,omitempty"`
	Name                 string   `protobuf:"bytes,18,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EdgeDevConfig) Reset()         { *m = EdgeDevConfig{} }
func (m *EdgeDevConfig) String() string { return proto.CompactTextString(m) }
func (*EdgeDevConfig) ProtoMessage()    {}
func (*EdgeDevConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc17241cd6d97458, []int{0}
}

func (m *EdgeDevConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EdgeDevConfig.Unmarshal(m, b)
}
func (m *EdgeDevConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EdgeDevConfig.Marshal(b, m, deterministic)
}
func (m *EdgeDevConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EdgeDevConfig.Merge(m, src)
}
func (m *EdgeDevConfig) XXX_Size() int {
	return xxx_messageInfo_EdgeDevConfig.Size(m)
}
func (m *EdgeDevConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_EdgeDevConfig.DiscardUnknown(m)
}

var xxx_messageInfo_EdgeDevConfig proto.InternalMessageInfo

func (m *EdgeDevConfig) GetId() *UUIDandVersion {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *EdgeDevConfig) GetApps() []*AppInstanceConfig {
	if m != nil {
		return m.Apps
	}
	return nil
}

func (m *EdgeDevConfig) GetNetworks() []*NetworkConfig {
	if m != nil {
		return m.Networks
	}
	return nil
}

func (m *EdgeDevConfig) GetDatastores() []*DatastoreConfig {
	if m != nil {
		return m.Datastores
	}
	return nil
}

func (m *EdgeDevConfig) GetLispInfo() *DeviceLispDetails {
	if m != nil {
		return m.LispInfo
	}
	return nil
}

func (m *EdgeDevConfig) GetBase() []*BaseOSConfig {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *EdgeDevConfig) GetReboot() *DeviceOpsCmd {
	if m != nil {
		return m.Reboot
	}
	return nil
}

func (m *EdgeDevConfig) GetBackup() *DeviceOpsCmd {
	if m != nil {
		return m.Backup
	}
	return nil
}

func (m *EdgeDevConfig) GetConfigItems() []*ConfigItem {
	if m != nil {
		return m.ConfigItems
	}
	return nil
}

func (m *EdgeDevConfig) GetSystemAdapterList() []*SystemAdapter {
	if m != nil {
		return m.SystemAdapterList
	}
	return nil
}

func (m *EdgeDevConfig) GetDeviceIoList() []*PhysicalIO {
	if m != nil {
		return m.DeviceIoList
	}
	return nil
}

func (m *EdgeDevConfig) GetManufacturer() string {
	if m != nil {
		return m.Manufacturer
	}
	return ""
}

func (m *EdgeDevConfig) GetProductName() string {
	if m != nil {
		return m.ProductName
	}
	return ""
}

func (m *EdgeDevConfig) GetNetworkInstances() []*NetworkInstanceConfig {
	if m != nil {
		return m.NetworkInstances
	}
	return nil
}

func (m *EdgeDevConfig) GetEnterprise() string {
	if m != nil {
		return m.Enterprise
	}
	return ""
}

func (m *EdgeDevConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ConfigRequest struct {
	ConfigHash           string   `protobuf:"bytes,1,opt,name=configHash,proto3" json:"configHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigRequest) Reset()         { *m = ConfigRequest{} }
func (m *ConfigRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest) ProtoMessage()    {}
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc17241cd6d97458, []int{1}
}

func (m *ConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest.Unmarshal(m, b)
}
func (m *ConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest.Marshal(b, m, deterministic)
}
func (m *ConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest.Merge(m, src)
}
func (m *ConfigRequest) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest.Size(m)
}
func (m *ConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest proto.InternalMessageInfo

func (m *ConfigRequest) GetConfigHash() string {
	if m != nil {
		return m.ConfigHash
	}
	return ""
}

type ConfigResponse struct {
	Config               *EdgeDevConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	ConfigHash           string         `protobuf:"bytes,2,opt,name=configHash,proto3" json:"configHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ConfigResponse) Reset()         { *m = ConfigResponse{} }
func (m *ConfigResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigResponse) ProtoMessage()    {}
func (*ConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc17241cd6d97458, []int{2}
}

func (m *ConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigResponse.Unmarshal(m, b)
}
func (m *ConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigResponse.Marshal(b, m, deterministic)
}
func (m *ConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigResponse.Merge(m, src)
}
func (m *ConfigResponse) XXX_Size() int {
	return xxx_messageInfo_ConfigResponse.Size(m)
}
func (m *ConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigResponse proto.InternalMessageInfo

func (m *ConfigResponse) GetConfig() *EdgeDevConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *ConfigResponse) GetConfigHash() string {
	if m != nil {
		return m.ConfigHash
	}
	return ""
}

func init() {
	proto.RegisterType((*EdgeDevConfig)(nil), "EdgeDevConfig")
	proto.RegisterType((*ConfigRequest)(nil), "ConfigRequest")
	proto.RegisterType((*ConfigResponse)(nil), "ConfigResponse")
}

func init() { proto.RegisterFile("devconfig.proto", fileDescriptor_fc17241cd6d97458) }

var fileDescriptor_fc17241cd6d97458 = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0x51, 0x4f, 0xdb, 0x3e,
	0x14, 0xc5, 0x05, 0xff, 0xd2, 0x3f, 0xdc, 0xd2, 0x02, 0x7e, 0x98, 0xac, 0x3d, 0x8c, 0xae, 0xda,
	0x10, 0x9a, 0x84, 0x33, 0xb1, 0xd7, 0xbd, 0x00, 0x9d, 0xb6, 0x4a, 0x08, 0xa6, 0x20, 0xa6, 0x69,
	0x6f, 0x6e, 0x7c, 0xdb, 0x5a, 0x34, 0xb6, 0x97, 0xeb, 0x74, 0x62, 0x5f, 0x6b, 0x5f, 0x70, 0x8a,
	0xe3, 0xa0, 0xa6, 0xec, 0x2d, 0x3e, 0xe7, 0xe7, 0x73, 0x9d, 0xeb, 0x6b, 0x38, 0x50, 0xb8, 0xca,
	0xac, 0x99, 0xe9, 0xb9, 0x70, 0x85, 0xf5, 0xf6, 0x65, 0x2d, 0xe4, 0xb9, 0x35, 0x8d, 0x20, 0x9d,
	0x6b, 0x11, 0x6c, 0x2a, 0x09, 0x2d, 0xb5, 0x77, 0x19, 0xf4, 0x2d, 0xa1, 0x4f, 0xde, 0x16, 0x72,
	0x8e, 0xcd, 0xd2, 0xa0, 0xd7, 0x86, 0x7c, 0x5c, 0x42, 0x8e, 0xb4, 0x88, 0xdf, 0x03, 0x85, 0xab,
	0xdc, 0x2a, 0x5c, 0xd6, 0xeb, 0xd1, 0x9f, 0x1d, 0xe8, 0x7f, 0x52, 0x73, 0x1c, 0xe3, 0xea, 0x2a,
	0x24, 0xb2, 0x63, 0xd8, 0xd6, 0x8a, 0x6f, 0x0d, 0xb7, 0x4e, 0x7b, 0xe7, 0x07, 0xe2, 0xfe, 0x7e,
	0x32, 0x96, 0x46, 0x7d, 0xc3, 0x82, 0xb4, 0x35, 0xe9, 0xb6, 0x56, 0xec, 0x04, 0x3a, 0xd2, 0x39,
	0xe2, 0x9d, 0xe1, 0x7f, 0xa7, 0xbd, 0x73, 0x26, 0x2e, 0x9c, 0x9b, 0x18, 0xf2, 0xd2, 0x64, 0x58,
	0x47, 0xa4, 0xc1, 0x67, 0xef, 0x60, 0xd7, 0xa0, 0xff, 0x65, 0x8b, 0x07, 0xe2, 0x3b, 0x81, 0x1d,
	0x88, 0x9b, 0x5a, 0x88, 0xdc, 0x93, 0xcf, 0xde, 0x03, 0x28, 0xe9, 0x65, 0xf5, 0x1b, 0x48, 0xbc,
	0x1b, 0xe8, 0x43, 0x31, 0x6e, 0xa4, 0xc8, 0xaf, 0x31, 0x4c, 0xc0, 0xee, 0x52, 0x93, 0x9b, 0x98,
	0x99, 0xe5, 0xff, 0x87, 0xc3, 0x32, 0x31, 0xc6, 0x95, 0xce, 0xf0, 0x5a, 0x93, 0x1b, 0xa3, 0x97,
	0x7a, 0x49, 0xe9, 0x13, 0xc3, 0x5e, 0x43, 0xa7, 0xea, 0x24, 0xdf, 0x0d, 0xd9, 0x7d, 0x71, 0x29,
	0x09, 0x6f, 0xef, 0x9a, 0x03, 0x57, 0x16, 0x7b, 0x0b, 0xdd, 0x02, 0xa7, 0xd6, 0x7a, 0xbe, 0x17,
	0x02, 0xfb, 0x31, 0xf0, 0xd6, 0xd1, 0x55, 0xae, 0xd2, 0x68, 0x56, 0xd8, 0x54, 0x66, 0x0f, 0xa5,
	0xe3, 0xf0, 0x4f, 0xac, 0x36, 0xd9, 0x19, 0xf4, 0xea, 0x3b, 0x9a, 0x78, 0xcc, 0x89, 0xf7, 0x42,
	0xdd, 0x9e, 0xb8, 0x7a, 0xd2, 0xd2, 0x75, 0x9f, 0x7d, 0x84, 0x23, 0x7a, 0x24, 0x8f, 0xf9, 0x85,
	0x92, 0xce, 0x63, 0x71, 0xad, 0xc9, 0xf3, 0xfd, 0xd8, 0xb6, 0xbb, 0x75, 0x27, 0x7d, 0x0e, 0xb2,
	0x04, 0xf6, 0x55, 0x38, 0xc4, 0xc4, 0x86, 0x8d, 0xfd, 0x58, 0xed, 0xeb, 0xe2, 0x91, 0x74, 0x26,
	0x97, 0x93, 0xdb, 0xb4, 0x05, 0xb0, 0x11, 0xec, 0xe7, 0xd2, 0x94, 0x33, 0x99, 0xf9, 0xb2, 0xc0,
	0x82, 0x0f, 0x86, 0x5b, 0xa7, 0x7b, 0x69, 0x4b, 0x63, 0x43, 0xe8, 0xb9, 0xc2, 0xaa, 0x32, 0xf3,
	0x37, 0x32, 0x47, 0x7e, 0x10, 0x90, 0x75, 0x89, 0x5d, 0xc2, 0x61, 0xbc, 0xc2, 0x66, 0x02, 0x88,
	0x1f, 0x86, 0xd2, 0x2f, 0x9a, 0xab, 0xde, 0x18, 0x8d, 0x67, 0x3c, 0x7b, 0x05, 0x80, 0xc6, 0x63,
	0xe1, 0x0a, 0x4d, 0xc8, 0x8f, 0x42, 0x91, 0x35, 0x85, 0x31, 0xe8, 0x98, 0xaa, 0x3c, 0x0b, 0x4e,
	0xf8, 0x1e, 0x25, 0xd0, 0x8f, 0x79, 0xf8, 0xb3, 0x44, 0xf2, 0x55, 0x48, 0xdd, 0xcc, 0x2f, 0x92,
	0x16, 0x61, 0x78, 0xf7, 0xd2, 0x35, 0x65, 0xf4, 0x1d, 0x06, 0xcd, 0x06, 0x72, 0xd6, 0x10, 0xb2,
	0x13, 0xe8, 0xd6, 0x7e, 0x1c, 0xf5, 0x81, 0x68, 0x3d, 0x83, 0x34, 0xba, 0x1b, 0xc9, 0xdb, 0x9b,
	0xc9, 0x97, 0x9f, 0xe1, 0x38, 0xb3, 0xb9, 0xf8, 0x8d, 0x0a, 0x95, 0x14, 0xd9, 0xd2, 0x96, 0x4a,
	0x94, 0x84, 0x45, 0xd5, 0xeb, 0xfa, 0x8d, 0xfd, 0x78, 0x33, 0xd7, 0x7e, 0x51, 0x4e, 0x45, 0x66,
	0xf3, 0x64, 0x39, 0x3b, 0x43, 0x35, 0xc7, 0x04, 0x57, 0x98, 0x48, 0xa7, 0x93, 0xb9, 0x4d, 0xea,
	0xb0, 0x69, 0x37, 0xc0, 0x1f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x54, 0xe9, 0xc1, 0x24,
	0x04, 0x00, 0x00,
}