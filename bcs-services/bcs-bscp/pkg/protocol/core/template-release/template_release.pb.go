// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: template_release.proto

package pbtr

import (
	base "bscp.io/pkg/protocol/core/base"
	config_item "bscp.io/pkg/protocol/core/config-item"
	content "bscp.io/pkg/protocol/core/content"
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

// Template source resource reference: pkg/dal/table/template_release.go
type TemplateRelease struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32                     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Spec       *TemplateReleaseSpec       `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Attachment *TemplateReleaseAttachment `protobuf:"bytes,3,opt,name=attachment,proto3" json:"attachment,omitempty"`
	Revision   *base.CreatedRevision      `protobuf:"bytes,4,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *TemplateRelease) Reset() {
	*x = TemplateRelease{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_release_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateRelease) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateRelease) ProtoMessage() {}

func (x *TemplateRelease) ProtoReflect() protoreflect.Message {
	mi := &file_template_release_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateRelease.ProtoReflect.Descriptor instead.
func (*TemplateRelease) Descriptor() ([]byte, []int) {
	return file_template_release_proto_rawDescGZIP(), []int{0}
}

func (x *TemplateRelease) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TemplateRelease) GetSpec() *TemplateReleaseSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *TemplateRelease) GetAttachment() *TemplateReleaseAttachment {
	if x != nil {
		return x.Attachment
	}
	return nil
}

func (x *TemplateRelease) GetRevision() *base.CreatedRevision {
	if x != nil {
		return x.Revision
	}
	return nil
}

// TemplateReleaseSpec source resource reference: pkg/dal/table/template_release.go
type TemplateReleaseSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReleaseName string                      `protobuf:"bytes,1,opt,name=release_name,json=releaseName,proto3" json:"release_name,omitempty"`
	ReleaseMemo string                      `protobuf:"bytes,2,opt,name=release_memo,json=releaseMemo,proto3" json:"release_memo,omitempty"`
	Name        string                      `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Path        string                      `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	FileType    string                      `protobuf:"bytes,5,opt,name=file_type,json=fileType,proto3" json:"file_type,omitempty"` // file_type is enum type
	FileMode    string                      `protobuf:"bytes,6,opt,name=file_mode,json=fileMode,proto3" json:"file_mode,omitempty"` // file_mode is enum type
	Permission  *config_item.FilePermission `protobuf:"bytes,7,opt,name=permission,proto3" json:"permission,omitempty"`
	ContentSpec *content.ContentSpec        `protobuf:"bytes,8,opt,name=content_spec,json=contentSpec,proto3" json:"content_spec,omitempty"`
}

func (x *TemplateReleaseSpec) Reset() {
	*x = TemplateReleaseSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_release_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateReleaseSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateReleaseSpec) ProtoMessage() {}

func (x *TemplateReleaseSpec) ProtoReflect() protoreflect.Message {
	mi := &file_template_release_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateReleaseSpec.ProtoReflect.Descriptor instead.
func (*TemplateReleaseSpec) Descriptor() ([]byte, []int) {
	return file_template_release_proto_rawDescGZIP(), []int{1}
}

func (x *TemplateReleaseSpec) GetReleaseName() string {
	if x != nil {
		return x.ReleaseName
	}
	return ""
}

func (x *TemplateReleaseSpec) GetReleaseMemo() string {
	if x != nil {
		return x.ReleaseMemo
	}
	return ""
}

func (x *TemplateReleaseSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TemplateReleaseSpec) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *TemplateReleaseSpec) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

func (x *TemplateReleaseSpec) GetFileMode() string {
	if x != nil {
		return x.FileMode
	}
	return ""
}

func (x *TemplateReleaseSpec) GetPermission() *config_item.FilePermission {
	if x != nil {
		return x.Permission
	}
	return nil
}

func (x *TemplateReleaseSpec) GetContentSpec() *content.ContentSpec {
	if x != nil {
		return x.ContentSpec
	}
	return nil
}

// TemplateReleaseAttachment source resource reference: pkg/dal/table/template_release.go
type TemplateReleaseAttachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizId           uint32 `protobuf:"varint,1,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"`
	TemplateSpaceId uint32 `protobuf:"varint,2,opt,name=template_space_id,json=templateSpaceId,proto3" json:"template_space_id,omitempty"`
	TemplateId      uint32 `protobuf:"varint,3,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
}

func (x *TemplateReleaseAttachment) Reset() {
	*x = TemplateReleaseAttachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_release_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateReleaseAttachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateReleaseAttachment) ProtoMessage() {}

func (x *TemplateReleaseAttachment) ProtoReflect() protoreflect.Message {
	mi := &file_template_release_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateReleaseAttachment.ProtoReflect.Descriptor instead.
func (*TemplateReleaseAttachment) Descriptor() ([]byte, []int) {
	return file_template_release_proto_rawDescGZIP(), []int{2}
}

func (x *TemplateReleaseAttachment) GetBizId() uint32 {
	if x != nil {
		return x.BizId
	}
	return 0
}

func (x *TemplateReleaseAttachment) GetTemplateSpaceId() uint32 {
	if x != nil {
		return x.TemplateSpaceId
	}
	return 0
}

func (x *TemplateReleaseAttachment) GetTemplateId() uint32 {
	if x != nil {
		return x.TemplateId
	}
	return 0
}

var File_template_release_proto protoreflect.FileDescriptor

var file_template_release_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x62, 0x74, 0x72, 0x1a, 0x29,
	0x62, 0x73, 0x63, 0x70, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x37, 0x62, 0x73, 0x63, 0x70, 0x2e,
	0x69, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2d, 0x69, 0x74, 0x65, 0x6d,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2f, 0x62, 0x73, 0x63, 0x70, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x01, 0x0a, 0x0f, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2d, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x62, 0x74, 0x72, 0x2e, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x70, 0x65, 0x63,
	0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x3f, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x62, 0x74,
	0x72, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xae, 0x02, 0x0a,
	0x13, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x34, 0x0a, 0x0a,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x70, 0x62, 0x63, 0x69, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x70,
	0x65, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x22, 0x7f, 0x0a,
	0x19, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x69,
	0x7a, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49,
	0x64, 0x12, 0x2a, 0x0a, 0x11, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x42, 0x31,
	0x5a, 0x2f, 0x62, 0x73, 0x63, 0x70, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x2d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x3b, 0x70, 0x62, 0x74,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_template_release_proto_rawDescOnce sync.Once
	file_template_release_proto_rawDescData = file_template_release_proto_rawDesc
)

func file_template_release_proto_rawDescGZIP() []byte {
	file_template_release_proto_rawDescOnce.Do(func() {
		file_template_release_proto_rawDescData = protoimpl.X.CompressGZIP(file_template_release_proto_rawDescData)
	})
	return file_template_release_proto_rawDescData
}

var file_template_release_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_template_release_proto_goTypes = []interface{}{
	(*TemplateRelease)(nil),            // 0: pbtr.TemplateRelease
	(*TemplateReleaseSpec)(nil),        // 1: pbtr.TemplateReleaseSpec
	(*TemplateReleaseAttachment)(nil),  // 2: pbtr.TemplateReleaseAttachment
	(*base.CreatedRevision)(nil),       // 3: pbbase.CreatedRevision
	(*config_item.FilePermission)(nil), // 4: pbci.FilePermission
	(*content.ContentSpec)(nil),        // 5: pbcontent.ContentSpec
}
var file_template_release_proto_depIdxs = []int32{
	1, // 0: pbtr.TemplateRelease.spec:type_name -> pbtr.TemplateReleaseSpec
	2, // 1: pbtr.TemplateRelease.attachment:type_name -> pbtr.TemplateReleaseAttachment
	3, // 2: pbtr.TemplateRelease.revision:type_name -> pbbase.CreatedRevision
	4, // 3: pbtr.TemplateReleaseSpec.permission:type_name -> pbci.FilePermission
	5, // 4: pbtr.TemplateReleaseSpec.content_spec:type_name -> pbcontent.ContentSpec
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_template_release_proto_init() }
func file_template_release_proto_init() {
	if File_template_release_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_template_release_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateRelease); i {
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
		file_template_release_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateReleaseSpec); i {
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
		file_template_release_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateReleaseAttachment); i {
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
			RawDescriptor: file_template_release_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_template_release_proto_goTypes,
		DependencyIndexes: file_template_release_proto_depIdxs,
		MessageInfos:      file_template_release_proto_msgTypes,
	}.Build()
	File_template_release_proto = out.File
	file_template_release_proto_rawDesc = nil
	file_template_release_proto_goTypes = nil
	file_template_release_proto_depIdxs = nil
}
