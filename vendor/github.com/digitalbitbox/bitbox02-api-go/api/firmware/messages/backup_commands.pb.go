// Copyright 2019 Shift Cryptosecurity AG
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

// This file is named backup_commands to avoid conflicting header files with top-most backup.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: backup_commands.proto

package messages

import (
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

type CheckBackupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Silent bool `protobuf:"varint,1,opt,name=silent,proto3" json:"silent,omitempty"`
}

func (x *CheckBackupRequest) Reset() {
	*x = CheckBackupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backup_commands_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckBackupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckBackupRequest) ProtoMessage() {}

func (x *CheckBackupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backup_commands_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckBackupRequest.ProtoReflect.Descriptor instead.
func (*CheckBackupRequest) Descriptor() ([]byte, []int) {
	return file_backup_commands_proto_rawDescGZIP(), []int{0}
}

func (x *CheckBackupRequest) GetSilent() bool {
	if x != nil {
		return x.Silent
	}
	return false
}

type CheckBackupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CheckBackupResponse) Reset() {
	*x = CheckBackupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backup_commands_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckBackupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckBackupResponse) ProtoMessage() {}

func (x *CheckBackupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backup_commands_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckBackupResponse.ProtoReflect.Descriptor instead.
func (*CheckBackupResponse) Descriptor() ([]byte, []int) {
	return file_backup_commands_proto_rawDescGZIP(), []int{1}
}

func (x *CheckBackupResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Timestamp must be in UTC
type CreateBackupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp      uint32 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	TimezoneOffset int32  `protobuf:"varint,2,opt,name=timezone_offset,json=timezoneOffset,proto3" json:"timezone_offset,omitempty"`
}

func (x *CreateBackupRequest) Reset() {
	*x = CreateBackupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backup_commands_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBackupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBackupRequest) ProtoMessage() {}

func (x *CreateBackupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backup_commands_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBackupRequest.ProtoReflect.Descriptor instead.
func (*CreateBackupRequest) Descriptor() ([]byte, []int) {
	return file_backup_commands_proto_rawDescGZIP(), []int{2}
}

func (x *CreateBackupRequest) GetTimestamp() uint32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *CreateBackupRequest) GetTimezoneOffset() int32 {
	if x != nil {
		return x.TimezoneOffset
	}
	return 0
}

type ListBackupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListBackupsRequest) Reset() {
	*x = ListBackupsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backup_commands_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBackupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBackupsRequest) ProtoMessage() {}

func (x *ListBackupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backup_commands_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBackupsRequest.ProtoReflect.Descriptor instead.
func (*ListBackupsRequest) Descriptor() ([]byte, []int) {
	return file_backup_commands_proto_rawDescGZIP(), []int{3}
}

type BackupInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp uint32 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// uint32 timezone_offset = 3;
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *BackupInfo) Reset() {
	*x = BackupInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backup_commands_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupInfo) ProtoMessage() {}

func (x *BackupInfo) ProtoReflect() protoreflect.Message {
	mi := &file_backup_commands_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupInfo.ProtoReflect.Descriptor instead.
func (*BackupInfo) Descriptor() ([]byte, []int) {
	return file_backup_commands_proto_rawDescGZIP(), []int{4}
}

func (x *BackupInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BackupInfo) GetTimestamp() uint32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *BackupInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListBackupsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info []*BackupInfo `protobuf:"bytes,1,rep,name=info,proto3" json:"info,omitempty"`
}

func (x *ListBackupsResponse) Reset() {
	*x = ListBackupsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backup_commands_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBackupsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBackupsResponse) ProtoMessage() {}

func (x *ListBackupsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backup_commands_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBackupsResponse.ProtoReflect.Descriptor instead.
func (*ListBackupsResponse) Descriptor() ([]byte, []int) {
	return file_backup_commands_proto_rawDescGZIP(), []int{5}
}

func (x *ListBackupsResponse) GetInfo() []*BackupInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type RestoreBackupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp      uint32 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	TimezoneOffset int32  `protobuf:"varint,3,opt,name=timezone_offset,json=timezoneOffset,proto3" json:"timezone_offset,omitempty"`
}

func (x *RestoreBackupRequest) Reset() {
	*x = RestoreBackupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backup_commands_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestoreBackupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestoreBackupRequest) ProtoMessage() {}

func (x *RestoreBackupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backup_commands_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestoreBackupRequest.ProtoReflect.Descriptor instead.
func (*RestoreBackupRequest) Descriptor() ([]byte, []int) {
	return file_backup_commands_proto_rawDescGZIP(), []int{6}
}

func (x *RestoreBackupRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RestoreBackupRequest) GetTimestamp() uint32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *RestoreBackupRequest) GetTimezoneOffset() int32 {
	if x != nil {
		return x.TimezoneOffset
	}
	return 0
}

var File_backup_commands_proto protoreflect.FileDescriptor

var file_backup_commands_proto_rawDesc = []byte{
	0x0a, 0x15, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x73, 0x68, 0x69, 0x66, 0x74, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x2e, 0x62, 0x69, 0x74, 0x62, 0x6f, 0x78, 0x30, 0x32, 0x22, 0x2c, 0x0a,
	0x12, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x6c, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x69, 0x6c, 0x65, 0x6e, 0x74, 0x22, 0x25, 0x0a, 0x13, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x5c, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x7a,
	0x6f, 0x6e, 0x65, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4e, 0x0a, 0x0a, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4b, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x68,
	0x69, 0x66, 0x74, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x62, 0x69, 0x74, 0x62, 0x6f, 0x78,
	0x30, 0x32, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x22, 0x6d, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x42, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x69, 0x6d,
	0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x4f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_backup_commands_proto_rawDescOnce sync.Once
	file_backup_commands_proto_rawDescData = file_backup_commands_proto_rawDesc
)

func file_backup_commands_proto_rawDescGZIP() []byte {
	file_backup_commands_proto_rawDescOnce.Do(func() {
		file_backup_commands_proto_rawDescData = protoimpl.X.CompressGZIP(file_backup_commands_proto_rawDescData)
	})
	return file_backup_commands_proto_rawDescData
}

var file_backup_commands_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_backup_commands_proto_goTypes = []interface{}{
	(*CheckBackupRequest)(nil),   // 0: shiftcrypto.bitbox02.CheckBackupRequest
	(*CheckBackupResponse)(nil),  // 1: shiftcrypto.bitbox02.CheckBackupResponse
	(*CreateBackupRequest)(nil),  // 2: shiftcrypto.bitbox02.CreateBackupRequest
	(*ListBackupsRequest)(nil),   // 3: shiftcrypto.bitbox02.ListBackupsRequest
	(*BackupInfo)(nil),           // 4: shiftcrypto.bitbox02.BackupInfo
	(*ListBackupsResponse)(nil),  // 5: shiftcrypto.bitbox02.ListBackupsResponse
	(*RestoreBackupRequest)(nil), // 6: shiftcrypto.bitbox02.RestoreBackupRequest
}
var file_backup_commands_proto_depIdxs = []int32{
	4, // 0: shiftcrypto.bitbox02.ListBackupsResponse.info:type_name -> shiftcrypto.bitbox02.BackupInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_backup_commands_proto_init() }
func file_backup_commands_proto_init() {
	if File_backup_commands_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_backup_commands_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckBackupRequest); i {
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
		file_backup_commands_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckBackupResponse); i {
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
		file_backup_commands_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBackupRequest); i {
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
		file_backup_commands_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBackupsRequest); i {
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
		file_backup_commands_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupInfo); i {
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
		file_backup_commands_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBackupsResponse); i {
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
		file_backup_commands_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestoreBackupRequest); i {
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
			RawDescriptor: file_backup_commands_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_backup_commands_proto_goTypes,
		DependencyIndexes: file_backup_commands_proto_depIdxs,
		MessageInfos:      file_backup_commands_proto_msgTypes,
	}.Build()
	File_backup_commands_proto = out.File
	file_backup_commands_proto_rawDesc = nil
	file_backup_commands_proto_goTypes = nil
	file_backup_commands_proto_depIdxs = nil
}
