// Code generated by protoc-gen-go.
// source: test.proto
// DO NOT EDIT!

package example

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type FOO int32

const (
	FOO_X FOO = 17
)

var FOO_name = map[int32]string{
	17: "X",
}
var FOO_value = map[string]int32{
	"X": 17,
}

func (x FOO) Enum() *FOO {
	p := new(FOO)
	*p = x
	return p
}
func (x FOO) String() string {
	return proto.EnumName(FOO_name, int32(x))
}
func (x *FOO) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FOO_value, data, "FOO")
	if err != nil {
		return err
	}
	*x = FOO(value)
	return nil
}
func (FOO) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type Test struct {
	Label            *string             `protobuf:"bytes,1,req,name=label" json:"label,omitempty"`
	Type             *int32              `protobuf:"varint,2,opt,name=type,def=77" json:"type,omitempty"`
	Reps             []int64             `protobuf:"varint,3,rep,name=reps" json:"reps,omitempty"`
	Optionalgroup    *Test_OptionalGroup `protobuf:"group,4,opt,name=OptionalGroup" json:"optionalgroup,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *Test) Reset()                    { *m = Test{} }
func (m *Test) String() string            { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()               {}
func (*Test) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

const Default_Test_Type int32 = 77

func (m *Test) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *Test) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_Test_Type
}

func (m *Test) GetReps() []int64 {
	if m != nil {
		return m.Reps
	}
	return nil
}

func (m *Test) GetOptionalgroup() *Test_OptionalGroup {
	if m != nil {
		return m.Optionalgroup
	}
	return nil
}

type Test_OptionalGroup struct {
	RequiredField    *string `protobuf:"bytes,5,req,name=RequiredField" json:"RequiredField,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Test_OptionalGroup) Reset()         { *m = Test_OptionalGroup{} }
func (m *Test_OptionalGroup) String() string { return proto.CompactTextString(m) }
func (*Test_OptionalGroup) ProtoMessage()    {}

func (m *Test_OptionalGroup) GetRequiredField() string {
	if m != nil && m.RequiredField != nil {
		return *m.RequiredField
	}
	return ""
}

func init() {
	proto.RegisterType((*Test)(nil), "example.Test")
	proto.RegisterType((*Test_OptionalGroup)(nil), "example.Test.OptionalGroup")
	proto.RegisterEnum("example.FOO", FOO_name, FOO_value)
}

var fileDescriptor2 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0x55,
	0x9a, 0xc8, 0xc8, 0xc5, 0x12, 0x02, 0x14, 0x17, 0xe2, 0xe5, 0x62, 0xcd, 0x49, 0x4c, 0x4a, 0xcd,
	0x91, 0x60, 0x54, 0x60, 0xd2, 0xe0, 0x14, 0x12, 0xe0, 0x62, 0x29, 0xa9, 0x2c, 0x48, 0x95, 0x60,
	0x52, 0x60, 0xd4, 0x60, 0xb5, 0x62, 0x32, 0x37, 0x17, 0xe2, 0xe1, 0x62, 0x29, 0x4a, 0x2d, 0x28,
	0x96, 0x60, 0x56, 0x60, 0xd6, 0x60, 0x16, 0x32, 0xe2, 0xe2, 0xcd, 0x2f, 0x28, 0xc9, 0xcc, 0xcf,
	0x4b, 0xcc, 0x49, 0x2f, 0xca, 0x2f, 0x2d, 0x90, 0x60, 0x01, 0x2a, 0xe4, 0x32, 0x92, 0xd6, 0x83,
	0x1a, 0xac, 0x07, 0x32, 0x54, 0xcf, 0x1f, 0xaa, 0xc4, 0x1d, 0xa4, 0x44, 0x4a, 0x8d, 0x8b, 0x17,
	0x45, 0x40, 0x48, 0x94, 0x8b, 0x37, 0x28, 0xb5, 0xb0, 0x34, 0xb3, 0x28, 0x35, 0xc5, 0x2d, 0x33,
	0x35, 0x27, 0x45, 0x82, 0x15, 0x64, 0xb7, 0x16, 0x0f, 0x17, 0xb3, 0x9b, 0xbf, 0xbf, 0x10, 0x2b,
	0x17, 0x63, 0x84, 0x80, 0x20, 0x20, 0x00, 0x00, 0xff, 0xff, 0x24, 0xb1, 0x18, 0x6a, 0xb7, 0x00,
	0x00, 0x00,
}
