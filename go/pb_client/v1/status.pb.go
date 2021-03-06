// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/status.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Status condition specification
type StatusCondition struct {
	// Status type
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Status state
	Status bool `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	// Status reason
	Reason string `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
	// Status message
	Message string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	// last update time
	LastUpdateTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=last_update_time,json=lastUpdateTime,proto3" json:"last_update_time,omitempty"`
	// last transition time
	LastTransitionTime   *timestamp.Timestamp `protobuf:"bytes,6,opt,name=last_transition_time,json=lastTransitionTime,proto3" json:"last_transition_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *StatusCondition) Reset()         { *m = StatusCondition{} }
func (m *StatusCondition) String() string { return proto.CompactTextString(m) }
func (*StatusCondition) ProtoMessage()    {}
func (*StatusCondition) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8dac81fd98b7df8, []int{0}
}

func (m *StatusCondition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusCondition.Unmarshal(m, b)
}
func (m *StatusCondition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusCondition.Marshal(b, m, deterministic)
}
func (m *StatusCondition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusCondition.Merge(m, src)
}
func (m *StatusCondition) XXX_Size() int {
	return xxx_messageInfo_StatusCondition.Size(m)
}
func (m *StatusCondition) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusCondition.DiscardUnknown(m)
}

var xxx_messageInfo_StatusCondition proto.InternalMessageInfo

func (m *StatusCondition) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *StatusCondition) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *StatusCondition) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *StatusCondition) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *StatusCondition) GetLastUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdateTime
	}
	return nil
}

func (m *StatusCondition) GetLastTransitionTime() *timestamp.Timestamp {
	if m != nil {
		return m.LastTransitionTime
	}
	return nil
}

// Status specification
type Status struct {
	// The uuid of the run
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// The current status
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	// The status conditions timeline
	StatusConditions     []*StatusCondition `protobuf:"bytes,3,rep,name=status_conditions,json=statusConditions,proto3" json:"status_conditions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8dac81fd98b7df8, []int{1}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Status) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Status) GetStatusConditions() []*StatusCondition {
	if m != nil {
		return m.StatusConditions
	}
	return nil
}

// Request data to create/update entity status
type EntityStatusBodyRequest struct {
	// Owner of the namespace
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// Project where the experiement will be assigned
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// Unique integer identifier of the entity
	Uuid string `protobuf:"bytes,3,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// Status to set
	Status               *Status  `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntityStatusBodyRequest) Reset()         { *m = EntityStatusBodyRequest{} }
func (m *EntityStatusBodyRequest) String() string { return proto.CompactTextString(m) }
func (*EntityStatusBodyRequest) ProtoMessage()    {}
func (*EntityStatusBodyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8dac81fd98b7df8, []int{2}
}

func (m *EntityStatusBodyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityStatusBodyRequest.Unmarshal(m, b)
}
func (m *EntityStatusBodyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityStatusBodyRequest.Marshal(b, m, deterministic)
}
func (m *EntityStatusBodyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityStatusBodyRequest.Merge(m, src)
}
func (m *EntityStatusBodyRequest) XXX_Size() int {
	return xxx_messageInfo_EntityStatusBodyRequest.Size(m)
}
func (m *EntityStatusBodyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityStatusBodyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EntityStatusBodyRequest proto.InternalMessageInfo

func (m *EntityStatusBodyRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *EntityStatusBodyRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *EntityStatusBodyRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *EntityStatusBodyRequest) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*StatusCondition)(nil), "v1.StatusCondition")
	proto.RegisterType((*Status)(nil), "v1.Status")
	proto.RegisterType((*EntityStatusBodyRequest)(nil), "v1.EntityStatusBodyRequest")
}

func init() { proto.RegisterFile("v1/status.proto", fileDescriptor_d8dac81fd98b7df8) }

var fileDescriptor_d8dac81fd98b7df8 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x3d, 0x4f, 0xf3, 0x30,
	0x14, 0x85, 0x95, 0xa6, 0xcd, 0xfb, 0xd6, 0x95, 0x68, 0x31, 0x15, 0x98, 0x2e, 0x44, 0x99, 0x32,
	0xa5, 0x6a, 0xf9, 0x03, 0x88, 0x8f, 0x8d, 0x29, 0x94, 0xb9, 0x72, 0x1b, 0x53, 0x05, 0xb5, 0x76,
	0xc8, 0xbd, 0x09, 0xca, 0x88, 0xf8, 0xe3, 0xc8, 0x76, 0x1c, 0xa0, 0x0b, 0x9b, 0xcf, 0xfd, 0xf2,
	0x73, 0x0e, 0x19, 0xd7, 0x8b, 0x39, 0x20, 0xc7, 0x0a, 0x92, 0xa2, 0x54, 0xa8, 0x68, 0xaf, 0x5e,
	0xcc, 0x2e, 0x77, 0x4a, 0xed, 0xf6, 0x62, 0x6e, 0x2a, 0x9b, 0xea, 0x65, 0xce, 0x65, 0x63, 0xdb,
	0xb3, 0xab, 0xe3, 0x16, 0xe6, 0x07, 0x01, 0xc8, 0x0f, 0x85, 0x1d, 0x88, 0x3e, 0x7b, 0x64, 0xfc,
	0x64, 0x0e, 0xde, 0x29, 0x99, 0xe5, 0x98, 0x2b, 0x49, 0x29, 0xe9, 0x63, 0x53, 0x08, 0xe6, 0x85,
	0x5e, 0x3c, 0x4c, 0xcd, 0x9b, 0x9e, 0x93, 0xc0, 0xfe, 0xcb, 0x7a, 0xa1, 0x17, 0xff, 0x4f, 0x5b,
	0xa5, 0xeb, 0xa5, 0xe0, 0xa0, 0x24, 0xf3, 0xcd, 0x74, 0xab, 0x28, 0x23, 0xff, 0x0e, 0x02, 0x80,
	0xef, 0x04, 0xeb, 0x9b, 0x86, 0x93, 0xf4, 0x9e, 0x4c, 0xf6, 0x1c, 0x70, 0x5d, 0x15, 0x19, 0x47,
	0xb1, 0xd6, 0x40, 0x6c, 0x10, 0x7a, 0xf1, 0x68, 0x39, 0x4b, 0x2c, 0x6d, 0xe2, 0x68, 0x93, 0x95,
	0xa3, 0x4d, 0x4f, 0xf4, 0xce, 0xb3, 0x59, 0xd1, 0x45, 0xfa, 0x48, 0xa6, 0xe6, 0x0a, 0x96, 0x5c,
	0x82, 0xc1, 0xb6, 0x97, 0x82, 0x3f, 0x2f, 0x51, 0xbd, 0xb7, 0xea, 0xd6, 0x74, 0x23, 0xaa, 0x49,
	0x60, 0x43, 0xd0, 0xde, 0xab, 0x2a, 0xcf, 0x9c, 0x77, 0xfd, 0x3e, 0xf2, 0x3e, 0xec, 0xbc, 0xdf,
	0x90, 0x53, 0xfb, 0x5a, 0x6f, 0x5d, 0x76, 0xc0, 0xfc, 0xd0, 0x8f, 0x47, 0xcb, 0xb3, 0xa4, 0x5e,
	0x24, 0x47, 0xb9, 0xa6, 0x13, 0xf8, 0x5d, 0x80, 0xe8, 0xc3, 0x23, 0x17, 0x0f, 0x12, 0x73, 0x6c,
	0xec, 0xec, 0xad, 0xca, 0x9a, 0x54, 0xbc, 0x55, 0x02, 0x90, 0x4e, 0xc9, 0x40, 0xbd, 0x4b, 0x51,
	0xb6, 0x28, 0x56, 0xe8, 0x5c, 0x8b, 0x52, 0xbd, 0x8a, 0x2d, 0xb6, 0x30, 0x4e, 0x76, 0xe4, 0xfe,
	0x0f, 0xf2, 0xa8, 0x23, 0xef, 0x9b, 0x5c, 0xc8, 0x37, 0x96, 0x73, 0xb1, 0x09, 0x4c, 0x46, 0xd7,
	0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x10, 0x61, 0x27, 0x43, 0x5b, 0x02, 0x00, 0x00,
}
