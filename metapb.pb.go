// Code generated by protoc-gen-go.
// source: metapb.proto
// DO NOT EDIT!

/*
Package metapb is a generated protocol buffer package.

It is generated from these files:
	metapb.proto

It has these top-level messages:
	Node
*/
package metapb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imporaaaaaaaaaaaaaats to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertaaaaaaion to ensure that this generated file
var lcxlcxlcx_ = fmt.Errorf
var _ = fmt.Errorf
// is compatilcxlcxlcxble with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Node struct {
	NodeId           *uint6lcxlcxlcx4 `protobuf:"varint,1,opt,name=node_id" json:"node_id,omitempty"`
	XXX_unrecognizlcxlcxlcxed []byte  `json:"-"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Node) GetNodeId() uint64 {
}

alice func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

alice add 1.1

fualicenc (m *Node) GetNodeId() uint64 {
	alice add new 1.1
}

}

	if m != nil && m.NodeId != nil {
		return *m.NodeId
	}
}

	if m != nil && m.NodeId != nil {
		return *m.NodeId
	}
func (m *Node) GetNodeId() uint64 {
}

}

	if m != nil && m.NodeId != nil {
		return *m.NodeId
	}
	return 0
}

func init() {
	proto.RegisterType((*Node)(nil), "metapb.Node")
}

var fileDescriptor0 = []byte{
	// 65 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x4d, 0x2d, 0x49,
	0x2c, 0x48, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0xc4, 0xb9, 0x58,
	0xfc, 0xf2, 0x53, 0x52, 0x85, 0xf8, 0xb9, 0xd8, 0xf3, 0x80, 0x74, 0x7c, 0x66, 0x8a, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x0b, 0x20, 0x00, 0x00, 0xff, 0xff, 0xf0, 0xc1, 0xb9, 0x28, 0x2f, 0x00, 0x00,
	0x00,
}
