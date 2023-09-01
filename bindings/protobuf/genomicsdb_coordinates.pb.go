//*
// The MIT License (MIT)
// Copyright (c) 2016-2018 Intel Corporation
// Copyright (c) dātma, inc™
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.7
// source: genomicsdb_coordinates.proto

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

type ContigPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contig   *string `protobuf:"bytes,1,req,name=contig" json:"contig,omitempty"`
	Position *int64  `protobuf:"varint,2,req,name=position" json:"position,omitempty"`
}

func (x *ContigPosition) Reset() {
	*x = ContigPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genomicsdb_coordinates_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContigPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContigPosition) ProtoMessage() {}

func (x *ContigPosition) ProtoReflect() protoreflect.Message {
	mi := &file_genomicsdb_coordinates_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContigPosition.ProtoReflect.Descriptor instead.
func (*ContigPosition) Descriptor() ([]byte, []int) {
	return file_genomicsdb_coordinates_proto_rawDescGZIP(), []int{0}
}

func (x *ContigPosition) GetContig() string {
	if x != nil && x.Contig != nil {
		return *x.Contig
	}
	return ""
}

func (x *ContigPosition) GetPosition() int64 {
	if x != nil && x.Position != nil {
		return *x.Position
	}
	return 0
}

type GenomicsDBColumn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Column:
	//
	//	*GenomicsDBColumn_TiledbColumn
	//	*GenomicsDBColumn_ContigPosition
	Column isGenomicsDBColumn_Column `protobuf_oneof:"column"`
}

func (x *GenomicsDBColumn) Reset() {
	*x = GenomicsDBColumn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genomicsdb_coordinates_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenomicsDBColumn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenomicsDBColumn) ProtoMessage() {}

func (x *GenomicsDBColumn) ProtoReflect() protoreflect.Message {
	mi := &file_genomicsdb_coordinates_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenomicsDBColumn.ProtoReflect.Descriptor instead.
func (*GenomicsDBColumn) Descriptor() ([]byte, []int) {
	return file_genomicsdb_coordinates_proto_rawDescGZIP(), []int{1}
}

func (m *GenomicsDBColumn) GetColumn() isGenomicsDBColumn_Column {
	if m != nil {
		return m.Column
	}
	return nil
}

func (x *GenomicsDBColumn) GetTiledbColumn() int64 {
	if x, ok := x.GetColumn().(*GenomicsDBColumn_TiledbColumn); ok {
		return x.TiledbColumn
	}
	return 0
}

func (x *GenomicsDBColumn) GetContigPosition() *ContigPosition {
	if x, ok := x.GetColumn().(*GenomicsDBColumn_ContigPosition); ok {
		return x.ContigPosition
	}
	return nil
}

type isGenomicsDBColumn_Column interface {
	isGenomicsDBColumn_Column()
}

type GenomicsDBColumn_TiledbColumn struct {
	TiledbColumn int64 `protobuf:"varint,1,opt,name=tiledb_column,json=tiledbColumn,oneof"`
}

type GenomicsDBColumn_ContigPosition struct {
	ContigPosition *ContigPosition `protobuf:"bytes,2,opt,name=contig_position,json=contigPosition,oneof"`
}

func (*GenomicsDBColumn_TiledbColumn) isGenomicsDBColumn_Column() {}

func (*GenomicsDBColumn_ContigPosition) isGenomicsDBColumn_Column() {}

type TileDBColumnInterval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Begin *int64 `protobuf:"varint,1,req,name=begin" json:"begin,omitempty"`
	End   *int64 `protobuf:"varint,2,req,name=end" json:"end,omitempty"`
}

func (x *TileDBColumnInterval) Reset() {
	*x = TileDBColumnInterval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genomicsdb_coordinates_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TileDBColumnInterval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TileDBColumnInterval) ProtoMessage() {}

func (x *TileDBColumnInterval) ProtoReflect() protoreflect.Message {
	mi := &file_genomicsdb_coordinates_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TileDBColumnInterval.ProtoReflect.Descriptor instead.
func (*TileDBColumnInterval) Descriptor() ([]byte, []int) {
	return file_genomicsdb_coordinates_proto_rawDescGZIP(), []int{2}
}

func (x *TileDBColumnInterval) GetBegin() int64 {
	if x != nil && x.Begin != nil {
		return *x.Begin
	}
	return 0
}

func (x *TileDBColumnInterval) GetEnd() int64 {
	if x != nil && x.End != nil {
		return *x.End
	}
	return 0
}

type ContigInterval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contig *string `protobuf:"bytes,1,req,name=contig" json:"contig,omitempty"`
	Begin  *int64  `protobuf:"varint,2,opt,name=begin" json:"begin,omitempty"`
	End    *int64  `protobuf:"varint,3,opt,name=end" json:"end,omitempty"`
}

func (x *ContigInterval) Reset() {
	*x = ContigInterval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genomicsdb_coordinates_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContigInterval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContigInterval) ProtoMessage() {}

func (x *ContigInterval) ProtoReflect() protoreflect.Message {
	mi := &file_genomicsdb_coordinates_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContigInterval.ProtoReflect.Descriptor instead.
func (*ContigInterval) Descriptor() ([]byte, []int) {
	return file_genomicsdb_coordinates_proto_rawDescGZIP(), []int{3}
}

func (x *ContigInterval) GetContig() string {
	if x != nil && x.Contig != nil {
		return *x.Contig
	}
	return ""
}

func (x *ContigInterval) GetBegin() int64 {
	if x != nil && x.Begin != nil {
		return *x.Begin
	}
	return 0
}

func (x *ContigInterval) GetEnd() int64 {
	if x != nil && x.End != nil {
		return *x.End
	}
	return 0
}

type GenomicsDBColumnInterval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Interval:
	//
	//	*GenomicsDBColumnInterval_TiledbColumnInterval
	//	*GenomicsDBColumnInterval_ContigInterval
	Interval isGenomicsDBColumnInterval_Interval `protobuf_oneof:"interval"`
}

func (x *GenomicsDBColumnInterval) Reset() {
	*x = GenomicsDBColumnInterval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genomicsdb_coordinates_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenomicsDBColumnInterval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenomicsDBColumnInterval) ProtoMessage() {}

func (x *GenomicsDBColumnInterval) ProtoReflect() protoreflect.Message {
	mi := &file_genomicsdb_coordinates_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenomicsDBColumnInterval.ProtoReflect.Descriptor instead.
func (*GenomicsDBColumnInterval) Descriptor() ([]byte, []int) {
	return file_genomicsdb_coordinates_proto_rawDescGZIP(), []int{4}
}

func (m *GenomicsDBColumnInterval) GetInterval() isGenomicsDBColumnInterval_Interval {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (x *GenomicsDBColumnInterval) GetTiledbColumnInterval() *TileDBColumnInterval {
	if x, ok := x.GetInterval().(*GenomicsDBColumnInterval_TiledbColumnInterval); ok {
		return x.TiledbColumnInterval
	}
	return nil
}

func (x *GenomicsDBColumnInterval) GetContigInterval() *ContigInterval {
	if x, ok := x.GetInterval().(*GenomicsDBColumnInterval_ContigInterval); ok {
		return x.ContigInterval
	}
	return nil
}

type isGenomicsDBColumnInterval_Interval interface {
	isGenomicsDBColumnInterval_Interval()
}

type GenomicsDBColumnInterval_TiledbColumnInterval struct {
	TiledbColumnInterval *TileDBColumnInterval `protobuf:"bytes,1,opt,name=tiledb_column_interval,json=tiledbColumnInterval,oneof"`
}

type GenomicsDBColumnInterval_ContigInterval struct {
	ContigInterval *ContigInterval `protobuf:"bytes,2,opt,name=contig_interval,json=contigInterval,oneof"`
}

func (*GenomicsDBColumnInterval_TiledbColumnInterval) isGenomicsDBColumnInterval_Interval() {}

func (*GenomicsDBColumnInterval_ContigInterval) isGenomicsDBColumnInterval_Interval() {}

type GenomicsDBColumnOrInterval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ColumnOrInterval:
	//
	//	*GenomicsDBColumnOrInterval_Column
	//	*GenomicsDBColumnOrInterval_ColumnInterval
	ColumnOrInterval isGenomicsDBColumnOrInterval_ColumnOrInterval `protobuf_oneof:"column_or_interval"`
}

func (x *GenomicsDBColumnOrInterval) Reset() {
	*x = GenomicsDBColumnOrInterval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genomicsdb_coordinates_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenomicsDBColumnOrInterval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenomicsDBColumnOrInterval) ProtoMessage() {}

func (x *GenomicsDBColumnOrInterval) ProtoReflect() protoreflect.Message {
	mi := &file_genomicsdb_coordinates_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenomicsDBColumnOrInterval.ProtoReflect.Descriptor instead.
func (*GenomicsDBColumnOrInterval) Descriptor() ([]byte, []int) {
	return file_genomicsdb_coordinates_proto_rawDescGZIP(), []int{5}
}

func (m *GenomicsDBColumnOrInterval) GetColumnOrInterval() isGenomicsDBColumnOrInterval_ColumnOrInterval {
	if m != nil {
		return m.ColumnOrInterval
	}
	return nil
}

func (x *GenomicsDBColumnOrInterval) GetColumn() *GenomicsDBColumn {
	if x, ok := x.GetColumnOrInterval().(*GenomicsDBColumnOrInterval_Column); ok {
		return x.Column
	}
	return nil
}

func (x *GenomicsDBColumnOrInterval) GetColumnInterval() *GenomicsDBColumnInterval {
	if x, ok := x.GetColumnOrInterval().(*GenomicsDBColumnOrInterval_ColumnInterval); ok {
		return x.ColumnInterval
	}
	return nil
}

type isGenomicsDBColumnOrInterval_ColumnOrInterval interface {
	isGenomicsDBColumnOrInterval_ColumnOrInterval()
}

type GenomicsDBColumnOrInterval_Column struct {
	Column *GenomicsDBColumn `protobuf:"bytes,1,opt,name=column,oneof"`
}

type GenomicsDBColumnOrInterval_ColumnInterval struct {
	ColumnInterval *GenomicsDBColumnInterval `protobuf:"bytes,2,opt,name=column_interval,json=columnInterval,oneof"`
}

func (*GenomicsDBColumnOrInterval_Column) isGenomicsDBColumnOrInterval_ColumnOrInterval() {}

func (*GenomicsDBColumnOrInterval_ColumnInterval) isGenomicsDBColumnOrInterval_ColumnOrInterval() {}

var File_genomicsdb_coordinates_proto protoreflect.FileDescriptor

var file_genomicsdb_coordinates_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x67, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x64, 0x62, 0x5f, 0x63, 0x6f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44,
	0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x67, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x67, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x02, 0x28, 0x03, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7f, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x73,
	0x44, 0x42, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x25, 0x0a, 0x0d, 0x74, 0x69, 0x6c, 0x65,
	0x64, 0x62, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48,
	0x00, 0x52, 0x0c, 0x74, 0x69, 0x6c, 0x65, 0x64, 0x62, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12,
	0x3a, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x67, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x69,
	0x67, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0e, 0x63, 0x6f, 0x6e,
	0x74, 0x69, 0x67, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0x0a, 0x06, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x22, 0x3e, 0x0a, 0x14, 0x54, 0x69, 0x6c, 0x65, 0x44, 0x42, 0x43,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x02, 0x28, 0x03, 0x52, 0x05, 0x62, 0x65,
	0x67, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x03,
	0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x50, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x67, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x74, 0x69,
	0x67, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x67, 0x12,
	0x14, 0x0a, 0x05, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x62, 0x65, 0x67, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0xb1, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x6e, 0x6f,
	0x6d, 0x69, 0x63, 0x73, 0x44, 0x42, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x12, 0x4d, 0x0a, 0x16, 0x74, 0x69, 0x6c, 0x65, 0x64, 0x62, 0x5f, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x54, 0x69, 0x6c, 0x65, 0x44, 0x42, 0x43, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x14, 0x74,
	0x69, 0x6c, 0x65, 0x64, 0x62, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x12, 0x3a, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x67, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x48, 0x00, 0x52,
	0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x42,
	0x0a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x22, 0xa5, 0x01, 0x0a, 0x1a,
	0x47, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x44, 0x42, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x4f, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x2b, 0x0a, 0x06, 0x63, 0x6f,
	0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x47, 0x65, 0x6e,
	0x6f, 0x6d, 0x69, 0x63, 0x73, 0x44, 0x42, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x48, 0x00, 0x52,
	0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x44, 0x0a, 0x0f, 0x63, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x47, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x44, 0x42, 0x43, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x0e, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x42, 0x14, 0x0a,
	0x12, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x42, 0x2e, 0x0a, 0x14, 0x6f, 0x72, 0x67, 0x2e, 0x67, 0x65, 0x6e, 0x6f, 0x6d,
	0x69, 0x63, 0x73, 0x64, 0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x0b, 0x43, 0x6f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x5a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f,
}

var (
	file_genomicsdb_coordinates_proto_rawDescOnce sync.Once
	file_genomicsdb_coordinates_proto_rawDescData = file_genomicsdb_coordinates_proto_rawDesc
)

func file_genomicsdb_coordinates_proto_rawDescGZIP() []byte {
	file_genomicsdb_coordinates_proto_rawDescOnce.Do(func() {
		file_genomicsdb_coordinates_proto_rawDescData = protoimpl.X.CompressGZIP(file_genomicsdb_coordinates_proto_rawDescData)
	})
	return file_genomicsdb_coordinates_proto_rawDescData
}

var file_genomicsdb_coordinates_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_genomicsdb_coordinates_proto_goTypes = []interface{}{
	(*ContigPosition)(nil),             // 0: ContigPosition
	(*GenomicsDBColumn)(nil),           // 1: GenomicsDBColumn
	(*TileDBColumnInterval)(nil),       // 2: TileDBColumnInterval
	(*ContigInterval)(nil),             // 3: ContigInterval
	(*GenomicsDBColumnInterval)(nil),   // 4: GenomicsDBColumnInterval
	(*GenomicsDBColumnOrInterval)(nil), // 5: GenomicsDBColumnOrInterval
}
var file_genomicsdb_coordinates_proto_depIdxs = []int32{
	0, // 0: GenomicsDBColumn.contig_position:type_name -> ContigPosition
	2, // 1: GenomicsDBColumnInterval.tiledb_column_interval:type_name -> TileDBColumnInterval
	3, // 2: GenomicsDBColumnInterval.contig_interval:type_name -> ContigInterval
	1, // 3: GenomicsDBColumnOrInterval.column:type_name -> GenomicsDBColumn
	4, // 4: GenomicsDBColumnOrInterval.column_interval:type_name -> GenomicsDBColumnInterval
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_genomicsdb_coordinates_proto_init() }
func file_genomicsdb_coordinates_proto_init() {
	if File_genomicsdb_coordinates_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_genomicsdb_coordinates_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContigPosition); i {
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
		file_genomicsdb_coordinates_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenomicsDBColumn); i {
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
		file_genomicsdb_coordinates_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TileDBColumnInterval); i {
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
		file_genomicsdb_coordinates_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContigInterval); i {
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
		file_genomicsdb_coordinates_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenomicsDBColumnInterval); i {
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
		file_genomicsdb_coordinates_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenomicsDBColumnOrInterval); i {
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
	file_genomicsdb_coordinates_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*GenomicsDBColumn_TiledbColumn)(nil),
		(*GenomicsDBColumn_ContigPosition)(nil),
	}
	file_genomicsdb_coordinates_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*GenomicsDBColumnInterval_TiledbColumnInterval)(nil),
		(*GenomicsDBColumnInterval_ContigInterval)(nil),
	}
	file_genomicsdb_coordinates_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*GenomicsDBColumnOrInterval_Column)(nil),
		(*GenomicsDBColumnOrInterval_ColumnInterval)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_genomicsdb_coordinates_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_genomicsdb_coordinates_proto_goTypes,
		DependencyIndexes: file_genomicsdb_coordinates_proto_depIdxs,
		MessageInfos:      file_genomicsdb_coordinates_proto_msgTypes,
	}.Build()
	File_genomicsdb_coordinates_proto = out.File
	file_genomicsdb_coordinates_proto_rawDesc = nil
	file_genomicsdb_coordinates_proto_goTypes = nil
	file_genomicsdb_coordinates_proto_depIdxs = nil
}
