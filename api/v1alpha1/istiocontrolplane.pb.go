// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1alpha1/istiocontrolplane.proto

package v1alpha1

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	v1alpha1 "istio.io/api/mesh/v1alpha1"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// IstioControlPlane defines an Istio control plane
//
// <!-- crd generation tags
// +cue-gen:IstioControlPlane:groupName:servicemesh.cisco.com
// +cue-gen:IstioControlPlane:version:v1alpha1
// +cue-gen:IstioControlPlane:storageVersion
// +cue-gen:IstioControlPlane:annotations:helm.sh/resource-policy=keep
// +cue-gen:IstioControlPlane:subresource:status
// +cue-gen:IstioControlPlane:scope:Namespaced
// +cue-gen:IstioControlPlane:resource:shortNames=icp,istiocp
// +cue-gen:IstioControlPlane:printerColumn:name=Age,type=date,JSONPath=.metadata.creationTimestamp,description="CreationTimestamp is a timestamp
// representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations.
// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
// Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
// +cue-gen:IstioControlPlane:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +genclient
// +k8s:deepcopy-gen=true
// -->
type IstioControlPlaneSpec struct {
	// Contains the intended version for the Istio control plane.
	// +kubebuilder:validation:Pattern=^1.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version"`
	// Logging configurations.
	Logging *LoggingConfiguration `protobuf:"bytes,2,opt,name=logging,proto3" json:"logging"`
	// Defines mesh-wide settings for the Istio control plane.
	MeshConfig           *v1alpha1.MeshConfig `protobuf:"bytes,3,opt,name=mesh_config,json=meshConfig,proto3" json:"mesh_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *IstioControlPlaneSpec) Reset()         { *m = IstioControlPlaneSpec{} }
func (m *IstioControlPlaneSpec) String() string { return proto.CompactTextString(m) }
func (*IstioControlPlaneSpec) ProtoMessage()    {}
func (*IstioControlPlaneSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6817de833805cb8b, []int{0}
}
func (m *IstioControlPlaneSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IstioControlPlaneSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IstioControlPlaneSpec.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IstioControlPlaneSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IstioControlPlaneSpec.Merge(m, src)
}
func (m *IstioControlPlaneSpec) XXX_Size() int {
	return m.Size()
}
func (m *IstioControlPlaneSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_IstioControlPlaneSpec.DiscardUnknown(m)
}

var xxx_messageInfo_IstioControlPlaneSpec proto.InternalMessageInfo

func (m *IstioControlPlaneSpec) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *IstioControlPlaneSpec) GetLogging() *LoggingConfiguration {
	if m != nil {
		return m.Logging
	}
	return nil
}

func (m *IstioControlPlaneSpec) GetMeshConfig() *v1alpha1.MeshConfig {
	if m != nil {
		return m.MeshConfig
	}
	return nil
}

// Comma-separated minimum per-scope logging level of messages to output, in the form of <scope>:<level>,<scope>:<level>
// The control plane has different scopes depending on component, but can configure default log level across all components
// If empty, default scope and level will be used as configured in code
type LoggingConfiguration struct {
	// +kubebuilder:validation:Pattern=`^([a-zA-Z]+:[a-zA-Z]+,?)+$`
	Level                string   `protobuf:"bytes,1,opt,name=level,proto3" json:"level"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoggingConfiguration) Reset()         { *m = LoggingConfiguration{} }
func (m *LoggingConfiguration) String() string { return proto.CompactTextString(m) }
func (*LoggingConfiguration) ProtoMessage()    {}
func (*LoggingConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_6817de833805cb8b, []int{1}
}
func (m *LoggingConfiguration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LoggingConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LoggingConfiguration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LoggingConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoggingConfiguration.Merge(m, src)
}
func (m *LoggingConfiguration) XXX_Size() int {
	return m.Size()
}
func (m *LoggingConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_LoggingConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_LoggingConfiguration proto.InternalMessageInfo

func (m *LoggingConfiguration) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

// <!-- go code generation tags
// +genclient
// +k8s:deepcopy-gen=true
// -->
type IstioControlPlaneStatus struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IstioControlPlaneStatus) Reset()         { *m = IstioControlPlaneStatus{} }
func (m *IstioControlPlaneStatus) String() string { return proto.CompactTextString(m) }
func (*IstioControlPlaneStatus) ProtoMessage()    {}
func (*IstioControlPlaneStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_6817de833805cb8b, []int{2}
}
func (m *IstioControlPlaneStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IstioControlPlaneStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IstioControlPlaneStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IstioControlPlaneStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IstioControlPlaneStatus.Merge(m, src)
}
func (m *IstioControlPlaneStatus) XXX_Size() int {
	return m.Size()
}
func (m *IstioControlPlaneStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_IstioControlPlaneStatus.DiscardUnknown(m)
}

var xxx_messageInfo_IstioControlPlaneStatus proto.InternalMessageInfo

func init() {
	proto.RegisterType((*IstioControlPlaneSpec)(nil), "istio_operator.v2.api.v1alpha1.IstioControlPlaneSpec")
	proto.RegisterType((*LoggingConfiguration)(nil), "istio_operator.v2.api.v1alpha1.LoggingConfiguration")
	proto.RegisterType((*IstioControlPlaneStatus)(nil), "istio_operator.v2.api.v1alpha1.IstioControlPlaneStatus")
}

func init() {
	proto.RegisterFile("api/v1alpha1/istiocontrolplane.proto", fileDescriptor_6817de833805cb8b)
}

var fileDescriptor_6817de833805cb8b = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x3d, 0x4f, 0xeb, 0x30,
	0x14, 0x95, 0xfb, 0xf4, 0x5e, 0x55, 0x77, 0x78, 0x52, 0x54, 0x44, 0xe9, 0x90, 0x54, 0x15, 0x43,
	0x19, 0x70, 0xd4, 0x02, 0x62, 0x85, 0x76, 0x42, 0x02, 0x09, 0x85, 0x0d, 0x86, 0xca, 0x49, 0x5d,
	0xc7, 0x92, 0x9b, 0x6b, 0x39, 0x4e, 0x06, 0x7e, 0x21, 0x23, 0xbf, 0xa0, 0x82, 0x8c, 0xfd, 0x15,
	0xc8, 0x71, 0x43, 0x91, 0x40, 0x4c, 0x3e, 0xba, 0x1f, 0xe7, 0xdc, 0x73, 0xaf, 0xf1, 0x31, 0x55,
	0x22, 0x2c, 0x27, 0x54, 0xaa, 0x94, 0x4e, 0x42, 0x91, 0x1b, 0x01, 0x09, 0x64, 0x46, 0x83, 0x54,
	0x92, 0x66, 0x8c, 0x28, 0x0d, 0x06, 0x3c, 0xbf, 0x4e, 0x2c, 0x40, 0x31, 0x4d, 0x0d, 0x68, 0x52,
	0x4e, 0x09, 0x55, 0x82, 0x34, 0x7d, 0x83, 0xc1, 0x9a, 0xe5, 0xe9, 0x9e, 0x26, 0x81, 0x6c, 0x25,
	0xb8, 0xeb, 0x1d, 0xf4, 0x38, 0x70, 0xa8, 0x61, 0x68, 0xd1, 0x2e, 0x1a, 0x70, 0x00, 0x2e, 0x59,
	0x68, 0xe5, 0x57, 0x82, 0xc9, 0xe5, 0x22, 0x66, 0x29, 0x2d, 0x05, 0x68, 0x57, 0x30, 0x7a, 0x47,
	0xf8, 0xe0, 0xc6, 0xaa, 0xce, 0xdd, 0x38, 0xf7, 0x76, 0x9c, 0x07, 0xc5, 0x12, 0xef, 0x04, 0xb7,
	0x4b, 0xa6, 0x73, 0x01, 0x59, 0x1f, 0x0d, 0xd1, 0xb8, 0x33, 0xfb, 0x5f, 0x5d, 0xa3, 0xd6, 0x76,
	0x13, 0x34, 0xe1, 0xa8, 0x01, 0xde, 0x13, 0x6e, 0x4b, 0xe0, 0x5c, 0x64, 0xbc, 0xdf, 0x1a, 0xa2,
	0x71, 0x77, 0x7a, 0x4e, 0x7e, 0x77, 0x42, 0x6e, 0x5d, 0xf9, 0xbc, 0x76, 0x50, 0x68, 0x6a, 0x04,
	0x64, 0xb3, 0xae, 0x25, 0xdf, 0x11, 0x45, 0x0d, 0xf0, 0xae, 0x70, 0xd7, 0xda, 0x5e, 0x38, 0xb7,
	0xfd, 0x3f, 0xb5, 0x40, 0xe0, 0x04, 0x88, 0xcd, 0xec, 0x59, 0xef, 0x58, 0x9e, 0x3a, 0xca, 0x08,
	0xaf, 0x3f, 0xf1, 0xe8, 0x12, 0xf7, 0x7e, 0xd2, 0xf3, 0x02, 0xfc, 0x57, 0xb2, 0x92, 0xc9, 0x9d,
	0xbf, 0xce, 0x76, 0x13, 0xb8, 0x40, 0xe4, 0x9e, 0xd1, 0x11, 0x3e, 0xfc, 0xbe, 0x1b, 0x43, 0x4d,
	0x91, 0xcf, 0xe6, 0x2f, 0x95, 0x8f, 0x5e, 0x2b, 0x1f, 0xbd, 0x55, 0x3e, 0x7a, 0xbc, 0xe0, 0xc2,
	0xa4, 0x45, 0x4c, 0x12, 0x58, 0x87, 0x31, 0xcd, 0x9e, 0xa9, 0x48, 0x24, 0x14, 0x4b, 0x77, 0xe8,
	0xd3, 0x66, 0x0b, 0x61, 0x39, 0x0d, 0xbf, 0xfe, 0x83, 0xf8, 0x5f, 0x7d, 0x83, 0xb3, 0x8f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x97, 0xa5, 0x39, 0xa5, 0x1e, 0x02, 0x00, 0x00,
}

func (m *IstioControlPlaneSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IstioControlPlaneSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IstioControlPlaneSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.MeshConfig != nil {
		{
			size, err := m.MeshConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIstiocontrolplane(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Logging != nil {
		{
			size, err := m.Logging.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIstiocontrolplane(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintIstiocontrolplane(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LoggingConfiguration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoggingConfiguration) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LoggingConfiguration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Level) > 0 {
		i -= len(m.Level)
		copy(dAtA[i:], m.Level)
		i = encodeVarintIstiocontrolplane(dAtA, i, uint64(len(m.Level)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IstioControlPlaneStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IstioControlPlaneStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IstioControlPlaneStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintIstiocontrolplane(dAtA []byte, offset int, v uint64) int {
	offset -= sovIstiocontrolplane(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IstioControlPlaneSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovIstiocontrolplane(uint64(l))
	}
	if m.Logging != nil {
		l = m.Logging.Size()
		n += 1 + l + sovIstiocontrolplane(uint64(l))
	}
	if m.MeshConfig != nil {
		l = m.MeshConfig.Size()
		n += 1 + l + sovIstiocontrolplane(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LoggingConfiguration) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Level)
	if l > 0 {
		n += 1 + l + sovIstiocontrolplane(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *IstioControlPlaneStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovIstiocontrolplane(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIstiocontrolplane(x uint64) (n int) {
	return sovIstiocontrolplane(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IstioControlPlaneSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIstiocontrolplane
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
			return fmt.Errorf("proto: IstioControlPlaneSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IstioControlPlaneSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIstiocontrolplane
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
				return ErrInvalidLengthIstiocontrolplane
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIstiocontrolplane
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Logging", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIstiocontrolplane
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIstiocontrolplane
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIstiocontrolplane
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Logging == nil {
				m.Logging = &LoggingConfiguration{}
			}
			if err := m.Logging.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MeshConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIstiocontrolplane
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIstiocontrolplane
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIstiocontrolplane
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MeshConfig == nil {
				m.MeshConfig = &v1alpha1.MeshConfig{}
			}
			if err := m.MeshConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIstiocontrolplane(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIstiocontrolplane
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LoggingConfiguration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIstiocontrolplane
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
			return fmt.Errorf("proto: LoggingConfiguration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoggingConfiguration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Level", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIstiocontrolplane
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
				return ErrInvalidLengthIstiocontrolplane
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIstiocontrolplane
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Level = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIstiocontrolplane(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIstiocontrolplane
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IstioControlPlaneStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIstiocontrolplane
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
			return fmt.Errorf("proto: IstioControlPlaneStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IstioControlPlaneStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipIstiocontrolplane(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIstiocontrolplane
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipIstiocontrolplane(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIstiocontrolplane
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
					return 0, ErrIntOverflowIstiocontrolplane
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowIstiocontrolplane
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
				return 0, ErrInvalidLengthIstiocontrolplane
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIstiocontrolplane
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIstiocontrolplane
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIstiocontrolplane        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIstiocontrolplane          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIstiocontrolplane = fmt.Errorf("proto: unexpected end of group")
)
