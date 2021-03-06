// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/cluster/redis/redis_cluster.proto

package envoy_config_cluster_redis

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type RedisClusterConfig struct {
	ClusterRefreshRate           *duration.Duration    `protobuf:"bytes,1,opt,name=cluster_refresh_rate,json=clusterRefreshRate,proto3" json:"cluster_refresh_rate,omitempty"`
	ClusterRefreshTimeout        *duration.Duration    `protobuf:"bytes,2,opt,name=cluster_refresh_timeout,json=clusterRefreshTimeout,proto3" json:"cluster_refresh_timeout,omitempty"`
	RedirectRefreshInterval      *duration.Duration    `protobuf:"bytes,3,opt,name=redirect_refresh_interval,json=redirectRefreshInterval,proto3" json:"redirect_refresh_interval,omitempty"`
	RedirectRefreshThreshold     *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=redirect_refresh_threshold,json=redirectRefreshThreshold,proto3" json:"redirect_refresh_threshold,omitempty"`
	FailureRefreshThreshold      uint32                `protobuf:"varint,5,opt,name=failure_refresh_threshold,json=failureRefreshThreshold,proto3" json:"failure_refresh_threshold,omitempty"`
	HostDegradedRefreshThreshold uint32                `protobuf:"varint,6,opt,name=host_degraded_refresh_threshold,json=hostDegradedRefreshThreshold,proto3" json:"host_degraded_refresh_threshold,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}              `json:"-"`
	XXX_unrecognized             []byte                `json:"-"`
	XXX_sizecache                int32                 `json:"-"`
}

func (m *RedisClusterConfig) Reset()         { *m = RedisClusterConfig{} }
func (m *RedisClusterConfig) String() string { return proto.CompactTextString(m) }
func (*RedisClusterConfig) ProtoMessage()    {}
func (*RedisClusterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d6593a6ec218c02, []int{0}
}

func (m *RedisClusterConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisClusterConfig.Unmarshal(m, b)
}
func (m *RedisClusterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisClusterConfig.Marshal(b, m, deterministic)
}
func (m *RedisClusterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisClusterConfig.Merge(m, src)
}
func (m *RedisClusterConfig) XXX_Size() int {
	return xxx_messageInfo_RedisClusterConfig.Size(m)
}
func (m *RedisClusterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisClusterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RedisClusterConfig proto.InternalMessageInfo

func (m *RedisClusterConfig) GetClusterRefreshRate() *duration.Duration {
	if m != nil {
		return m.ClusterRefreshRate
	}
	return nil
}

func (m *RedisClusterConfig) GetClusterRefreshTimeout() *duration.Duration {
	if m != nil {
		return m.ClusterRefreshTimeout
	}
	return nil
}

func (m *RedisClusterConfig) GetRedirectRefreshInterval() *duration.Duration {
	if m != nil {
		return m.RedirectRefreshInterval
	}
	return nil
}

func (m *RedisClusterConfig) GetRedirectRefreshThreshold() *wrappers.UInt32Value {
	if m != nil {
		return m.RedirectRefreshThreshold
	}
	return nil
}

func (m *RedisClusterConfig) GetFailureRefreshThreshold() uint32 {
	if m != nil {
		return m.FailureRefreshThreshold
	}
	return 0
}

func (m *RedisClusterConfig) GetHostDegradedRefreshThreshold() uint32 {
	if m != nil {
		return m.HostDegradedRefreshThreshold
	}
	return 0
}

func init() {
	proto.RegisterType((*RedisClusterConfig)(nil), "envoy.config.cluster.redis.RedisClusterConfig")
}

func init() {
	proto.RegisterFile("envoy/config/cluster/redis/redis_cluster.proto", fileDescriptor_6d6593a6ec218c02)
}

var fileDescriptor_6d6593a6ec218c02 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x8e, 0xd3, 0x30,
	0x14, 0x45, 0x49, 0x87, 0xa9, 0x90, 0x11, 0x12, 0x58, 0xa0, 0xa4, 0xd5, 0x00, 0x23, 0x56, 0x15,
	0x0b, 0x5b, 0x9a, 0xd9, 0xb1, 0xcc, 0x14, 0xa1, 0xee, 0xaa, 0xa8, 0x65, 0x01, 0x8b, 0xc8, 0xad,
	0x5f, 0x52, 0x4b, 0x21, 0x8e, 0x9c, 0xe7, 0x42, 0x77, 0x7c, 0x0f, 0x9f, 0xc0, 0x17, 0xb0, 0xe5,
	0x77, 0x58, 0x20, 0x14, 0xdb, 0x41, 0xd0, 0x14, 0x31, 0x9b, 0xc4, 0xf2, 0xbd, 0xf7, 0xe8, 0x3e,
	0xeb, 0x11, 0x06, 0xf5, 0x5e, 0x1f, 0xf8, 0x56, 0xd7, 0x85, 0x2a, 0xf9, 0xb6, 0xb2, 0x2d, 0x82,
	0xe1, 0x06, 0xa4, 0x6a, 0xfd, 0x37, 0x0f, 0x77, 0xac, 0x31, 0x1a, 0x35, 0x9d, 0x3a, 0x3f, 0xf3,
	0x7e, 0xd6, 0x6b, 0xce, 0x39, 0x7d, 0x56, 0x6a, 0x5d, 0x56, 0xc0, 0x9d, 0x73, 0x63, 0x0b, 0x2e,
	0xad, 0x11, 0xa8, 0x74, 0xed, 0xb3, 0x43, 0xfd, 0xa3, 0x11, 0x4d, 0x03, 0xa6, 0x0d, 0xfa, 0x53,
	0x2b, 0x1b, 0xc1, 0x45, 0x5d, 0x6b, 0x74, 0xb1, 0x96, 0xb7, 0x28, 0xd0, 0xf6, 0x72, 0xbc, 0x17,
	0x95, 0x92, 0x02, 0x81, 0xf7, 0x07, 0x2f, 0xbc, 0xf8, 0x79, 0x46, 0x68, 0xd6, 0x35, 0xb8, 0xf1,
	0x75, 0x6e, 0x5c, 0x39, 0xba, 0x26, 0x8f, 0x43, 0xbf, 0xdc, 0x40, 0x61, 0xa0, 0xdd, 0xe5, 0x46,
	0x20, 0x24, 0xd1, 0x65, 0x34, 0xbb, 0x7f, 0x35, 0x61, 0xbe, 0x0d, 0xeb, 0xdb, 0xb0, 0x79, 0x68,
	0x9b, 0xde, 0xfb, 0x91, 0x9e, 0x7f, 0x89, 0x46, 0x2f, 0xef, 0x64, 0x34, 0x00, 0x32, 0x9f, 0xcf,
	0x04, 0x02, 0x7d, 0x4f, 0xe2, 0x63, 0x2c, 0xaa, 0x0f, 0xa0, 0x2d, 0x26, 0xa3, 0xdb, 0x93, 0x9f,
	0xfc, 0x4d, 0x5e, 0x79, 0x02, 0x5d, 0x93, 0x49, 0xf7, 0x96, 0x06, 0xb6, 0xf8, 0x9b, 0xae, 0x6a,
	0x04, 0xb3, 0x17, 0x55, 0x72, 0xf6, 0x1f, 0x7c, 0x16, 0xf7, 0xd9, 0x40, 0x5d, 0x84, 0x24, 0x7d,
	0x47, 0xa6, 0x03, 0x2c, 0xee, 0xba, 0x9f, 0xae, 0x64, 0x72, 0xd7, 0x71, 0x2f, 0x06, 0xdc, 0xf5,
	0xa2, 0xc6, 0xeb, 0xab, 0xb7, 0xa2, 0xb2, 0x90, 0x25, 0x47, 0xe8, 0x55, 0x9f, 0xa6, 0xaf, 0xc8,
	0xa4, 0x10, 0xaa, 0xb2, 0x06, 0x4e, 0xa0, 0xcf, 0x2f, 0xa3, 0xd9, 0x83, 0x2c, 0x0e, 0x86, 0x41,
	0xf6, 0x35, 0x79, 0xbe, 0xd3, 0x2d, 0xe6, 0x12, 0x4a, 0x23, 0x24, 0xc8, 0x13, 0x84, 0xb1, 0x23,
	0x5c, 0x74, 0xb6, 0x79, 0x70, 0x1d, 0x63, 0xd2, 0x37, 0x5f, 0x3f, 0x7f, 0xfb, 0x3e, 0x1e, 0x3d,
	0x8c, 0xc8, 0x4c, 0x69, 0xbf, 0xd1, 0x8d, 0xd1, 0x9f, 0x0e, 0xec, 0xdf, 0xcb, 0x9a, 0x3e, 0xfa,
	0x73, 0x63, 0x96, 0xdd, 0xc8, 0xcb, 0x68, 0x33, 0x76, 0xb3, 0x5f, 0xff, 0x0a, 0x00, 0x00, 0xff,
	0xff, 0x97, 0x72, 0xe4, 0x91, 0x16, 0x03, 0x00, 0x00,
}
