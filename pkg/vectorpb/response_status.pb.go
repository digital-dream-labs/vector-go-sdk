// Copyright (c) 2018 Anki, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License in the file LICENSE.txt or at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Response status

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.4
// source: response_status.proto

package vectorpb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ResponseStatus_StatusCode int32

const (
	ResponseStatus_UNKNOWN ResponseStatus_StatusCode = 0
	// The message has completed as expected.
	ResponseStatus_RESPONSE_RECEIVED ResponseStatus_StatusCode = 1
	// The message has been sent to the robot.
	ResponseStatus_REQUEST_PROCESSING ResponseStatus_StatusCode = 2
	// The message has been handled successfully at the interface level.
	ResponseStatus_OK ResponseStatus_StatusCode = 3
	// The user was not authorizied.
	ResponseStatus_FORBIDDEN ResponseStatus_StatusCode = 100
	// The requested attribute was not found.
	ResponseStatus_NOT_FOUND ResponseStatus_StatusCode = 101
	// Currently updating values from another call.
	ResponseStatus_ERROR_UPDATE_IN_PROGRESS ResponseStatus_StatusCode = 102
)

// Enum value maps for ResponseStatus_StatusCode.
var (
	ResponseStatus_StatusCode_name = map[int32]string{
		0:   "UNKNOWN",
		1:   "RESPONSE_RECEIVED",
		2:   "REQUEST_PROCESSING",
		3:   "OK",
		100: "FORBIDDEN",
		101: "NOT_FOUND",
		102: "ERROR_UPDATE_IN_PROGRESS",
	}
	ResponseStatus_StatusCode_value = map[string]int32{
		"UNKNOWN":                  0,
		"RESPONSE_RECEIVED":        1,
		"REQUEST_PROCESSING":       2,
		"OK":                       3,
		"FORBIDDEN":                100,
		"NOT_FOUND":                101,
		"ERROR_UPDATE_IN_PROGRESS": 102,
	}
)

func (x ResponseStatus_StatusCode) Enum() *ResponseStatus_StatusCode {
	p := new(ResponseStatus_StatusCode)
	*p = x
	return p
}

func (x ResponseStatus_StatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResponseStatus_StatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_response_status_proto_enumTypes[0].Descriptor()
}

func (ResponseStatus_StatusCode) Type() protoreflect.EnumType {
	return &file_response_status_proto_enumTypes[0]
}

func (x ResponseStatus_StatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResponseStatus_StatusCode.Descriptor instead.
func (ResponseStatus_StatusCode) EnumDescriptor() ([]byte, []int) {
	return file_response_status_proto_rawDescGZIP(), []int{0, 0}
}

// A shared response message sent back as part of most requests.
// This will indicate the generic state of the request.
type ResponseStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The generic status code to give high-level insight into the progress of a given message.
	Code ResponseStatus_StatusCode `protobuf:"varint,1,opt,name=code,proto3,enum=Anki.Vector.external_interface.ResponseStatus_StatusCode" json:"code,omitempty"`
}

func (x *ResponseStatus) Reset() {
	*x = ResponseStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_response_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseStatus) ProtoMessage() {}

func (x *ResponseStatus) ProtoReflect() protoreflect.Message {
	mi := &file_response_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseStatus.ProtoReflect.Descriptor instead.
func (*ResponseStatus) Descriptor() ([]byte, []int) {
	return file_response_status_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseStatus) GetCode() ResponseStatus_StatusCode {
	if x != nil {
		return x.Code
	}
	return ResponseStatus_UNKNOWN
}

var File_response_status_proto protoreflect.FileDescriptor

var file_response_status_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x41, 0x6e, 0x6b, 0x69, 0x2e, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x22, 0xee, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4d, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x39, 0x2e, 0x41, 0x6e, 0x6b, 0x69, 0x2e,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x8c, 0x01, 0x0a, 0x0a, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53,
	0x45, 0x5f, 0x52, 0x45, 0x43, 0x45, 0x49, 0x56, 0x45, 0x44, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12,
	0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x49,
	0x4e, 0x47, 0x10, 0x02, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09,
	0x46, 0x4f, 0x52, 0x42, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x10, 0x64, 0x12, 0x0d, 0x0a, 0x09, 0x4e,
	0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x52,
	0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x66, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x2d, 0x64,
	0x72, 0x65, 0x61, 0x6d, 0x2d, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x2d, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x76, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_response_status_proto_rawDescOnce sync.Once
	file_response_status_proto_rawDescData = file_response_status_proto_rawDesc
)

func file_response_status_proto_rawDescGZIP() []byte {
	file_response_status_proto_rawDescOnce.Do(func() {
		file_response_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_response_status_proto_rawDescData)
	})
	return file_response_status_proto_rawDescData
}

var file_response_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_response_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_response_status_proto_goTypes = []interface{}{
	(ResponseStatus_StatusCode)(0), // 0: Anki.Vector.external_interface.ResponseStatus.StatusCode
	(*ResponseStatus)(nil),         // 1: Anki.Vector.external_interface.ResponseStatus
}
var file_response_status_proto_depIdxs = []int32{
	0, // 0: Anki.Vector.external_interface.ResponseStatus.code:type_name -> Anki.Vector.external_interface.ResponseStatus.StatusCode
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_response_status_proto_init() }
func file_response_status_proto_init() {
	if File_response_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_response_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_response_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_response_status_proto_goTypes,
		DependencyIndexes: file_response_status_proto_depIdxs,
		EnumInfos:         file_response_status_proto_enumTypes,
		MessageInfos:      file_response_status_proto_msgTypes,
	}.Build()
	File_response_status_proto = out.File
	file_response_status_proto_rawDesc = nil
	file_response_status_proto_goTypes = nil
	file_response_status_proto_depIdxs = nil
}
