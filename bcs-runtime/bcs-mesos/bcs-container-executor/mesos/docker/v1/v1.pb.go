

// Code generated by protoc-gen-go.
// source: mesos/docker/v1.proto
// DO NOT EDIT!

/*
Package docker_spec_v1 is a generated protocol buffer package.

It is generated from these files:

	mesos/docker/v1.proto

It has these top-level messages:

	Label
	ImageManifest
*/
package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// *
// The Label of a docker image, used for adding custom metadata.
// https://docs.docker.com/engine/userguide/labels-custom-metadata
type Label struct {
	Key              *string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value            *string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Label) Reset()                    { *m = Label{} }
func (m *Label) String() string            { return proto.CompactTextString(m) }
func (*Label) ProtoMessage()               {}
func (*Label) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Label) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Label) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

// *
// Protobuf for the Docker v1 image manifest JSON schema:
// https://github.com/docker/docker/blob/master/image/spec/v1.md
type ImageManifest struct {
	// Following docker code to define and order protobuf fields.
	// https://github.com/docker/docker/blob/master/image/image.go
	Id     *string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Parent *string `protobuf:"bytes,2,opt,name=parent" json:"parent,omitempty"`
	// This field is used to comment user added comment. It is not
	// covered in docker v1 doc, but it is included in docker code.
	Comment *string `protobuf:"bytes,3,opt,name=comment" json:"comment,omitempty"`
	Created *string `protobuf:"bytes,4,opt,name=created" json:"created,omitempty"`
	// Container is the id of the container used to commit. It is
	// not covered in docker v1 doc, but included in docker code.
	Container        *string               `protobuf:"bytes,5,opt,name=container" json:"container,omitempty"`
	ContainerConfig  *ImageManifest_Config `protobuf:"bytes,6,opt,name=container_config,json=containerConfig" json:"container_config,omitempty"`
	DockerVersion    *string               `protobuf:"bytes,7,opt,name=docker_version,json=dockerVersion" json:"docker_version,omitempty"`
	Author           *string               `protobuf:"bytes,8,opt,name=author" json:"author,omitempty"`
	Config           *ImageManifest_Config `protobuf:"bytes,9,opt,name=config" json:"config,omitempty"`
	Architecture     *string               `protobuf:"bytes,10,opt,name=architecture" json:"architecture,omitempty"`
	Os               *string               `protobuf:"bytes,11,opt,name=os" json:"os,omitempty"`
	Size             *uint32               `protobuf:"varint,12,opt,name=Size" json:"Size,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *ImageManifest) Reset()                    { *m = ImageManifest{} }
func (m *ImageManifest) String() string            { return proto.CompactTextString(m) }
func (*ImageManifest) ProtoMessage()               {}
func (*ImageManifest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ImageManifest) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *ImageManifest) GetParent() string {
	if m != nil && m.Parent != nil {
		return *m.Parent
	}
	return ""
}

func (m *ImageManifest) GetComment() string {
	if m != nil && m.Comment != nil {
		return *m.Comment
	}
	return ""
}

func (m *ImageManifest) GetCreated() string {
	if m != nil && m.Created != nil {
		return *m.Created
	}
	return ""
}

func (m *ImageManifest) GetContainer() string {
	if m != nil && m.Container != nil {
		return *m.Container
	}
	return ""
}

func (m *ImageManifest) GetContainerConfig() *ImageManifest_Config {
	if m != nil {
		return m.ContainerConfig
	}
	return nil
}

func (m *ImageManifest) GetDockerVersion() string {
	if m != nil && m.DockerVersion != nil {
		return *m.DockerVersion
	}
	return ""
}

func (m *ImageManifest) GetAuthor() string {
	if m != nil && m.Author != nil {
		return *m.Author
	}
	return ""
}

func (m *ImageManifest) GetConfig() *ImageManifest_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *ImageManifest) GetArchitecture() string {
	if m != nil && m.Architecture != nil {
		return *m.Architecture
	}
	return ""
}

func (m *ImageManifest) GetOs() string {
	if m != nil && m.Os != nil {
		return *m.Os
	}
	return ""
}

func (m *ImageManifest) GetSize() uint32 {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return 0
}

// TODO(gilbert): Add other config fields.
// Currently necessary rumtime config fields only. Please see:
// https://github.com/docker/docker/blob/master/runconfig/config.go
type ImageManifest_Config struct {
	Hostname   *string  `protobuf:"bytes,1,opt,name=Hostname" json:"Hostname,omitempty"`
	Entrypoint []string `protobuf:"bytes,2,rep,name=Entrypoint" json:"Entrypoint,omitempty"`
	Env        []string `protobuf:"bytes,3,rep,name=Env" json:"Env,omitempty"`
	User       *string  `protobuf:"bytes,4,opt,name=User" json:"User,omitempty"`
	Cmd        []string `protobuf:"bytes,5,rep,name=Cmd" json:"Cmd,omitempty"`
	WorkingDir *string  `protobuf:"bytes,6,opt,name=WorkingDir" json:"WorkingDir,omitempty"`
	// Name of the image as it was passed by the operator.
	Image *string `protobuf:"bytes,8,opt,name=Image" json:"Image,omitempty"`
	// NOTE: We cannot use 'Labels' here because otherwise,
	// json->protobuf parsing will fail. 'labels' is manually
	// set during parsing.
	Labels           []*Label `protobuf:"bytes,9,rep,name=labels" json:"labels,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ImageManifest_Config) Reset()                    { *m = ImageManifest_Config{} }
func (m *ImageManifest_Config) String() string            { return proto.CompactTextString(m) }
func (*ImageManifest_Config) ProtoMessage()               {}
func (*ImageManifest_Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *ImageManifest_Config) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

func (m *ImageManifest_Config) GetEntrypoint() []string {
	if m != nil {
		return m.Entrypoint
	}
	return nil
}

func (m *ImageManifest_Config) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *ImageManifest_Config) GetUser() string {
	if m != nil && m.User != nil {
		return *m.User
	}
	return ""
}

func (m *ImageManifest_Config) GetCmd() []string {
	if m != nil {
		return m.Cmd
	}
	return nil
}

func (m *ImageManifest_Config) GetWorkingDir() string {
	if m != nil && m.WorkingDir != nil {
		return *m.WorkingDir
	}
	return ""
}

func (m *ImageManifest_Config) GetImage() string {
	if m != nil && m.Image != nil {
		return *m.Image
	}
	return ""
}

func (m *ImageManifest_Config) GetLabels() []*Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*Label)(nil), "docker.spec.v1.Label")
	proto.RegisterType((*ImageManifest)(nil), "docker.spec.v1.ImageManifest")
	proto.RegisterType((*ImageManifest_Config)(nil), "docker.spec.v1.ImageManifest.Config")
}

func init() { proto.RegisterFile("mesos/docker/v1.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x52, 0xcf, 0xab, 0xd3, 0x40,
	0x10, 0x26, 0xcd, 0x4b, 0xde, 0xcb, 0xf4, 0xf5, 0xf9, 0x58, 0xac, 0x2c, 0x45, 0x24, 0x14, 0x85,
	0x5c, 0x4c, 0x69, 0xcf, 0xde, 0x6a, 0x41, 0x41, 0x11, 0x22, 0xea, 0xb1, 0xac, 0xc9, 0xb4, 0x5d,
	0xda, 0xec, 0x86, 0xdd, 0x6d, 0xa0, 0x5e, 0xfd, 0x4b, 0xfd, 0x4f, 0x64, 0x7f, 0x58, 0x5b, 0x6f,
	0xde, 0xe6, 0xfb, 0xbe, 0x99, 0xdd, 0x99, 0x6f, 0x06, 0xc6, 0x2d, 0x6a, 0xa9, 0x67, 0x8d, 0xac,
	0xf7, 0xa8, 0x66, 0xfd, 0xbc, 0xec, 0x94, 0x34, 0x92, 0x3c, 0x78, 0xa2, 0xd4, 0x1d, 0xd6, 0x65,
	0x3f, 0x9f, 0xce, 0x20, 0xf9, 0xc0, 0xbe, 0xe3, 0x81, 0x3c, 0x42, 0xbc, 0xc7, 0x13, 0x8d, 0xf2,
	0xa8, 0xc8, 0x2a, 0x1b, 0x92, 0xa7, 0x90, 0xf4, 0xec, 0x70, 0x44, 0x3a, 0x70, 0x9c, 0x07, 0xd3,
	0x9f, 0x09, 0x8c, 0xde, 0xb7, 0x6c, 0x8b, 0x1f, 0x99, 0xe0, 0x1b, 0xd4, 0x86, 0x3c, 0xc0, 0x80,
	0x37, 0xa1, 0x70, 0xc0, 0x1b, 0xf2, 0x0c, 0xd2, 0x8e, 0x29, 0x14, 0x26, 0x14, 0x06, 0x44, 0x28,
	0xdc, 0xd6, 0xb2, 0x6d, 0xad, 0x10, 0x3b, 0xe1, 0x0f, 0x74, 0x8a, 0x42, 0x66, 0xb0, 0xa1, 0x37,
	0x41, 0xf1, 0x90, 0x3c, 0x87, 0xac, 0x96, 0xc2, 0x30, 0x2e, 0x50, 0xd1, 0xc4, 0x69, 0x7f, 0x09,
	0xf2, 0x09, 0x1e, 0xcf, 0x60, 0x5d, 0x4b, 0xb1, 0xe1, 0x5b, 0x9a, 0xe6, 0x51, 0x31, 0x5c, 0xbc,
	0x2c, 0xaf, 0xe7, 0x2c, 0xaf, 0x5a, 0x2e, 0x97, 0x2e, 0xb7, 0x7a, 0x72, 0xae, 0xf6, 0x04, 0x79,
	0x05, 0xc1, 0x9f, 0x75, 0x8f, 0x4a, 0x73, 0x29, 0xe8, 0xad, 0xfb, 0x73, 0xe4, 0xd9, 0xaf, 0x9e,
	0xb4, 0x13, 0xb2, 0xa3, 0xd9, 0x49, 0x45, 0xef, 0xfc, 0x84, 0x1e, 0x91, 0x37, 0x90, 0x86, 0x2e,
	0xb2, 0xff, 0xe8, 0x22, 0xd4, 0x90, 0x29, 0xdc, 0x33, 0x55, 0xef, 0xb8, 0xc1, 0xda, 0x1c, 0x15,
	0x52, 0x70, 0x6f, 0x5f, 0x71, 0xd6, 0x6b, 0xa9, 0xe9, 0xd0, 0x7b, 0x2d, 0x35, 0x21, 0x70, 0xf3,
	0x99, 0xff, 0x40, 0x7a, 0x9f, 0x47, 0xc5, 0xa8, 0x72, 0xf1, 0xe4, 0x57, 0x04, 0x69, 0x98, 0x67,
	0x02, 0x77, 0xef, 0xa4, 0x36, 0x82, 0xb5, 0x18, 0x16, 0x74, 0xc6, 0xe4, 0x05, 0xc0, 0x4a, 0x18,
	0x75, 0xea, 0x24, 0x77, 0xab, 0x8a, 0x8b, 0xac, 0xba, 0x60, 0xec, 0x41, 0xac, 0x44, 0x4f, 0x63,
	0x27, 0xd8, 0xd0, 0x7e, 0xf6, 0x45, 0xa3, 0x0a, 0x3b, 0x72, 0xb1, 0xcd, 0x5a, 0xb6, 0x0d, 0x4d,
	0x7c, 0xd6, 0xb2, 0x6d, 0xec, 0xbb, 0xdf, 0xa4, 0xda, 0x73, 0xb1, 0x7d, 0xcb, 0x95, 0x5b, 0x47,
	0x56, 0x5d, 0x30, 0xf6, 0xac, 0x9c, 0x0d, 0xc1, 0x3b, 0x0f, 0xc8, 0x6b, 0x48, 0x0f, 0xf6, 0x0e,
	0x35, 0xcd, 0xf2, 0xb8, 0x18, 0x2e, 0xc6, 0xff, 0x5a, 0xe7, 0xae, 0xb4, 0x0a, 0x49, 0xbf, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x60, 0xb7, 0xc5, 0xb6, 0xde, 0x02, 0x00, 0x00,
}
