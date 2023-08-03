/**
 *
 * The MIT License
 *
 * Copyright (c) 2023 dātma, inc™
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 *
 */

package bindings

// #cgo pkg-config: genomicsdb
// #cgo CXXFLAGS: -std=c++14
// #include <genomicsdb_go.h>
// #include <stdlib.h>
import "C"

import (
	"log"

	"github.com/GenomicsDB/GenomicsDB-Go/protobuf"

	"google.golang.org/protobuf/proto"
)

func GetVersion() string {
	return C.GoString(C.version())
}

func ptr[T any](v T) *T {
	return &v
}

func Query() {
	// Construct protobuf
	export_config := protobuf.ExportConfiguration{
		Workspace: ptr("/Users/nalini/genomicsdb_demo/ws"),
		Array:     &protobuf.ExportConfiguration_ArrayName{ArrayName: "allcontigs$1$3095677412"},
		VidMappingInfo: &protobuf.ExportConfiguration_VidMappingFile{
			VidMappingFile: "/Users/nalini/genomicsdb_demo/ws/vidmap.json"},
		CallsetMappingInfo: &protobuf.ExportConfiguration_CallsetMappingFile{
			CallsetMappingFile: "/Users/nalini/genomicsdb_demo/ws/callset.json",
		},
		QueryContigIntervals: []*protobuf.ContigInterval{
			{Contig: ptr("17"), Begin: ptr(int64(7571719)), End: ptr(int64(7590868))},
		},
		QueryRowRanges: []*protobuf.RowRangeList{
			{RangeList: []*protobuf.RowRange{{Low: ptr(int64(0)), High: ptr(int64(200000))}}},
		},
		Attributes:                       []string{"REF", "ALT", "GT"},
		QueryFilter:                      ptr("REF==\"A\" && ALT|=\"T\" && GT&=\"1/1\""),
		BypassIntersectingIntervalsPhase: ptr(true),
		EnableSharedPosixfsOptimizations: ptr(true),
	}

	data, err := proto.Marshal(&export_config)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	len := C.size_t(C.int(len(data)))
	p := C.malloc(len)
	defer C.free(p)
	cBuf := (*[1 << 30]byte)(p)
	copy(cBuf[:], data)

	genomicsdb_handle := C.connect(p, len)
	C.query_variant_calls(genomicsdb_handle)
	C.disconnect(genomicsdb_handle)
}
