//*
// The MIT License (MIT)
// Copyright (c) 2016-2017 Intel Corporation
// Copyright (c) 2023 dātma, inc™
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

//*
// Note: Variable names are kept aligned with
// GenomicsDB JSON configuration.
//
// To build, use protocol buffer version >3.2.0 and the following command from base directory:
// $ protoc -Isrc/resources/ --java_out=src/main/java/ src/resources/genomicsdb_vid_mapping.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.7
// source: genomicsdb_vid_mapping.proto

package protobuf

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

type FieldLengthDescriptorComponentPB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to LengthDescriptor:
	//
	//	*FieldLengthDescriptorComponentPB_VariableLengthDescriptor
	//	*FieldLengthDescriptorComponentPB_FixedLength
	LengthDescriptor isFieldLengthDescriptorComponentPB_LengthDescriptor `protobuf_oneof:"length_descriptor"`
}

func (x *FieldLengthDescriptorComponentPB) Reset() {
	*x = FieldLengthDescriptorComponentPB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genomicsdb_vid_mapping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldLengthDescriptorComponentPB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldLengthDescriptorComponentPB) ProtoMessage() {}

func (x *FieldLengthDescriptorComponentPB) ProtoReflect() protoreflect.Message {
	mi := &file_genomicsdb_vid_mapping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldLengthDescriptorComponentPB.ProtoReflect.Descriptor instead.
func (*FieldLengthDescriptorComponentPB) Descriptor() ([]byte, []int) {
	return file_genomicsdb_vid_mapping_proto_rawDescGZIP(), []int{0}
}

func (m *FieldLengthDescriptorComponentPB) GetLengthDescriptor() isFieldLengthDescriptorComponentPB_LengthDescriptor {
	if m != nil {
		return m.LengthDescriptor
	}
	return nil
}

func (x *FieldLengthDescriptorComponentPB) GetVariableLengthDescriptor() string {
	if x, ok := x.GetLengthDescriptor().(*FieldLengthDescriptorComponentPB_VariableLengthDescriptor); ok {
		return x.VariableLengthDescriptor
	}
	return ""
}

func (x *FieldLengthDescriptorComponentPB) GetFixedLength() int32 {
	if x, ok := x.GetLengthDescriptor().(*FieldLengthDescriptorComponentPB_FixedLength); ok {
		return x.FixedLength
	}
	return 0
}

type isFieldLengthDescriptorComponentPB_LengthDescriptor interface {
	isFieldLengthDescriptorComponentPB_LengthDescriptor()
}

type FieldLengthDescriptorComponentPB_VariableLengthDescriptor struct {
	VariableLengthDescriptor string `protobuf:"bytes,1,opt,name=variable_length_descriptor,json=variableLengthDescriptor,oneof"`
}

type FieldLengthDescriptorComponentPB_FixedLength struct {
	FixedLength int32 `protobuf:"varint,2,opt,name=fixed_length,json=fixedLength,oneof"`
}

func (*FieldLengthDescriptorComponentPB_VariableLengthDescriptor) isFieldLengthDescriptorComponentPB_LengthDescriptor() {
}

func (*FieldLengthDescriptorComponentPB_FixedLength) isFieldLengthDescriptorComponentPB_LengthDescriptor() {
}

type GenomicsDBFieldInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                     *string                             `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Type                     []string                            `protobuf:"bytes,2,rep,name=type" json:"type,omitempty"`
	VcfFieldClass            []string                            `protobuf:"bytes,3,rep,name=vcf_field_class,json=vcfFieldClass" json:"vcf_field_class,omitempty"`
	VcfType                  *string                             `protobuf:"bytes,4,opt,name=vcf_type,json=vcfType" json:"vcf_type,omitempty"`
	Length                   []*FieldLengthDescriptorComponentPB `protobuf:"bytes,5,rep,name=length" json:"length,omitempty"`
	VcfDelimiter             []string                            `protobuf:"bytes,6,rep,name=vcf_delimiter,json=vcfDelimiter" json:"vcf_delimiter,omitempty"`
	VCFFieldCombineOperation *string                             `protobuf:"bytes,7,opt,name=VCF_field_combine_operation,json=VCFFieldCombineOperation" json:"VCF_field_combine_operation,omitempty"`
	// useful when multiple fields of different types/length with the same
	// name (FILTER, FORMAT, INFO)  are defined in the VCF header
	VcfName                       *string `protobuf:"bytes,8,opt,name=vcf_name,json=vcfName" json:"vcf_name,omitempty"`
	DisableRemapMissingWithNonRef *bool   `protobuf:"varint,9,opt,name=disable_remap_missing_with_non_ref,json=disableRemapMissingWithNonRef,def=0" json:"disable_remap_missing_with_non_ref,omitempty"`
}

// Default values for GenomicsDBFieldInfo fields.
const (
	Default_GenomicsDBFieldInfo_DisableRemapMissingWithNonRef = bool(false)
)

func (x *GenomicsDBFieldInfo) Reset() {
	*x = GenomicsDBFieldInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genomicsdb_vid_mapping_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenomicsDBFieldInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenomicsDBFieldInfo) ProtoMessage() {}

func (x *GenomicsDBFieldInfo) ProtoReflect() protoreflect.Message {
	mi := &file_genomicsdb_vid_mapping_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenomicsDBFieldInfo.ProtoReflect.Descriptor instead.
func (*GenomicsDBFieldInfo) Descriptor() ([]byte, []int) {
	return file_genomicsdb_vid_mapping_proto_rawDescGZIP(), []int{1}
}

func (x *GenomicsDBFieldInfo) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *GenomicsDBFieldInfo) GetType() []string {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *GenomicsDBFieldInfo) GetVcfFieldClass() []string {
	if x != nil {
		return x.VcfFieldClass
	}
	return nil
}

func (x *GenomicsDBFieldInfo) GetVcfType() string {
	if x != nil && x.VcfType != nil {
		return *x.VcfType
	}
	return ""
}

func (x *GenomicsDBFieldInfo) GetLength() []*FieldLengthDescriptorComponentPB {
	if x != nil {
		return x.Length
	}
	return nil
}

func (x *GenomicsDBFieldInfo) GetVcfDelimiter() []string {
	if x != nil {
		return x.VcfDelimiter
	}
	return nil
}

func (x *GenomicsDBFieldInfo) GetVCFFieldCombineOperation() string {
	if x != nil && x.VCFFieldCombineOperation != nil {
		return *x.VCFFieldCombineOperation
	}
	return ""
}

func (x *GenomicsDBFieldInfo) GetVcfName() string {
	if x != nil && x.VcfName != nil {
		return *x.VcfName
	}
	return ""
}

func (x *GenomicsDBFieldInfo) GetDisableRemapMissingWithNonRef() bool {
	if x != nil && x.DisableRemapMissingWithNonRef != nil {
		return *x.DisableRemapMissingWithNonRef
	}
	return Default_GenomicsDBFieldInfo_DisableRemapMissingWithNonRef
}

type Chromosome struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name               *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Length             *int64  `protobuf:"varint,2,req,name=length" json:"length,omitempty"`
	TiledbColumnOffset *int64  `protobuf:"varint,3,req,name=tiledb_column_offset,json=tiledbColumnOffset" json:"tiledb_column_offset,omitempty"`
}

func (x *Chromosome) Reset() {
	*x = Chromosome{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genomicsdb_vid_mapping_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chromosome) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chromosome) ProtoMessage() {}

func (x *Chromosome) ProtoReflect() protoreflect.Message {
	mi := &file_genomicsdb_vid_mapping_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chromosome.ProtoReflect.Descriptor instead.
func (*Chromosome) Descriptor() ([]byte, []int) {
	return file_genomicsdb_vid_mapping_proto_rawDescGZIP(), []int{2}
}

func (x *Chromosome) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Chromosome) GetLength() int64 {
	if x != nil && x.Length != nil {
		return *x.Length
	}
	return 0
}

func (x *Chromosome) GetTiledbColumnOffset() int64 {
	if x != nil && x.TiledbColumnOffset != nil {
		return *x.TiledbColumnOffset
	}
	return 0
}

type VidMappingPB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields  []*GenomicsDBFieldInfo `protobuf:"bytes,1,rep,name=fields" json:"fields,omitempty"`
	Contigs []*Chromosome          `protobuf:"bytes,2,rep,name=contigs" json:"contigs,omitempty"`
}

func (x *VidMappingPB) Reset() {
	*x = VidMappingPB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genomicsdb_vid_mapping_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VidMappingPB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VidMappingPB) ProtoMessage() {}

func (x *VidMappingPB) ProtoReflect() protoreflect.Message {
	mi := &file_genomicsdb_vid_mapping_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VidMappingPB.ProtoReflect.Descriptor instead.
func (*VidMappingPB) Descriptor() ([]byte, []int) {
	return file_genomicsdb_vid_mapping_proto_rawDescGZIP(), []int{3}
}

func (x *VidMappingPB) GetFields() []*GenomicsDBFieldInfo {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *VidMappingPB) GetContigs() []*Chromosome {
	if x != nil {
		return x.Contigs
	}
	return nil
}

var File_genomicsdb_vid_mapping_proto protoreflect.FileDescriptor

var file_genomicsdb_vid_mapping_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x67, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x64, 0x62, 0x5f, 0x76, 0x69, 0x64,
	0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c,
	0x01, 0x0a, 0x20, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x50, 0x42, 0x12, 0x3e, 0x0a, 0x1a, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x18, 0x76, 0x61, 0x72, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x12, 0x23, 0x0a, 0x0c, 0x66, 0x69, 0x78, 0x65, 0x64, 0x5f, 0x6c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0b, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x42, 0x13, 0x0a, 0x11, 0x6c, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x22, 0x8c, 0x03,
	0x0a, 0x13, 0x47, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x44, 0x42, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a,
	0x0f, 0x76, 0x63, 0x66, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x76, 0x63, 0x66, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x63, 0x66, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x63, 0x66, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x39, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x50, 0x42, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x23, 0x0a, 0x0d, 0x76,
	0x63, 0x66, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0c, 0x76, 0x63, 0x66, 0x44, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72,
	0x12, 0x3d, 0x0a, 0x1b, 0x56, 0x43, 0x46, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x63, 0x6f,
	0x6d, 0x62, 0x69, 0x6e, 0x65, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x56, 0x43, 0x46, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43,
	0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x19, 0x0a, 0x08, 0x76, 0x63, 0x66, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x63, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x50, 0x0a, 0x22, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x6d, 0x61, 0x70, 0x5f, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6e, 0x67, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x6e, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x1d, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x6d, 0x61, 0x70, 0x4d, 0x69, 0x73, 0x73, 0x69,
	0x6e, 0x67, 0x57, 0x69, 0x74, 0x68, 0x4e, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x22, 0x6a, 0x0a, 0x0a,
	0x43, 0x68, 0x72, 0x6f, 0x6d, 0x6f, 0x73, 0x6f, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x02, 0x28, 0x03, 0x52, 0x06,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x30, 0x0a, 0x14, 0x74, 0x69, 0x6c, 0x65, 0x64, 0x62,
	0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03,
	0x20, 0x02, 0x28, 0x03, 0x52, 0x12, 0x74, 0x69, 0x6c, 0x65, 0x64, 0x62, 0x43, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x63, 0x0a, 0x0c, 0x56, 0x69, 0x64, 0x4d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x50, 0x42, 0x12, 0x2c, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x47, 0x65, 0x6e, 0x6f, 0x6d,
	0x69, 0x63, 0x73, 0x44, 0x42, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x25, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x67,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x43, 0x68, 0x72, 0x6f, 0x6d, 0x6f,
	0x73, 0x6f, 0x6d, 0x65, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x67, 0x73, 0x42, 0x38, 0x0a,
	0x14, 0x6f, 0x72, 0x67, 0x2e, 0x67, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x64, 0x62, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x15, 0x47, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x44,
	0x42, 0x56, 0x69, 0x64, 0x4d, 0x61, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x09, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
}

var (
	file_genomicsdb_vid_mapping_proto_rawDescOnce sync.Once
	file_genomicsdb_vid_mapping_proto_rawDescData = file_genomicsdb_vid_mapping_proto_rawDesc
)

func file_genomicsdb_vid_mapping_proto_rawDescGZIP() []byte {
	file_genomicsdb_vid_mapping_proto_rawDescOnce.Do(func() {
		file_genomicsdb_vid_mapping_proto_rawDescData = protoimpl.X.CompressGZIP(file_genomicsdb_vid_mapping_proto_rawDescData)
	})
	return file_genomicsdb_vid_mapping_proto_rawDescData
}

var file_genomicsdb_vid_mapping_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_genomicsdb_vid_mapping_proto_goTypes = []interface{}{
	(*FieldLengthDescriptorComponentPB)(nil), // 0: FieldLengthDescriptorComponentPB
	(*GenomicsDBFieldInfo)(nil),              // 1: GenomicsDBFieldInfo
	(*Chromosome)(nil),                       // 2: Chromosome
	(*VidMappingPB)(nil),                     // 3: VidMappingPB
}
var file_genomicsdb_vid_mapping_proto_depIdxs = []int32{
	0, // 0: GenomicsDBFieldInfo.length:type_name -> FieldLengthDescriptorComponentPB
	1, // 1: VidMappingPB.fields:type_name -> GenomicsDBFieldInfo
	2, // 2: VidMappingPB.contigs:type_name -> Chromosome
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_genomicsdb_vid_mapping_proto_init() }
func file_genomicsdb_vid_mapping_proto_init() {
	if File_genomicsdb_vid_mapping_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_genomicsdb_vid_mapping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldLengthDescriptorComponentPB); i {
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
		file_genomicsdb_vid_mapping_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenomicsDBFieldInfo); i {
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
		file_genomicsdb_vid_mapping_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chromosome); i {
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
		file_genomicsdb_vid_mapping_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VidMappingPB); i {
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
	file_genomicsdb_vid_mapping_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*FieldLengthDescriptorComponentPB_VariableLengthDescriptor)(nil),
		(*FieldLengthDescriptorComponentPB_FixedLength)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_genomicsdb_vid_mapping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_genomicsdb_vid_mapping_proto_goTypes,
		DependencyIndexes: file_genomicsdb_vid_mapping_proto_depIdxs,
		MessageInfos:      file_genomicsdb_vid_mapping_proto_msgTypes,
	}.Build()
	File_genomicsdb_vid_mapping_proto = out.File
	file_genomicsdb_vid_mapping_proto_rawDesc = nil
	file_genomicsdb_vid_mapping_proto_goTypes = nil
	file_genomicsdb_vid_mapping_proto_depIdxs = nil
}