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
	"fmt"
	"path/filepath"

	"github.com/GenomicsDB/GenomicsDB-Go/bindings/protobuf"
	"google.golang.org/protobuf/proto"
)

func GetVersion() string {
	return C.GoString(C.version())
}

func ptr[T any](v T) *T {
	return &v
}

type GenomicsDBQueryConfig struct {
	Workspace       string
	Array           string
	ContigIntervals []*protobuf.ContigInterval
	RowRanges       []*protobuf.RowRangeList
	Attributes      []string
	Filter          string
}

func Query(queryConfig GenomicsDBQueryConfig) (bool, string) {
	// Construct export config protobuf
	var exportConfig protobuf.ExportConfiguration

	exportConfig.Workspace = ptr(queryConfig.Workspace)
	exportConfig.Array = &protobuf.ExportConfiguration_ArrayName{ArrayName: queryConfig.Array}
	exportConfig.VidMappingInfo = &protobuf.ExportConfiguration_VidMappingFile{
		VidMappingFile: filepath.Join(queryConfig.Workspace, "vidmap.json")}
	exportConfig.CallsetMappingInfo = &protobuf.ExportConfiguration_CallsetMappingFile{
		CallsetMappingFile: filepath.Join(queryConfig.Workspace, "callset.json"),
	}
	exportConfig.BypassIntersectingIntervalsPhase = ptr(true)
	exportConfig.EnableSharedPosixfsOptimizations = ptr(true)

	if queryConfig.ContigIntervals != nil {
		exportConfig.QueryContigIntervals = queryConfig.ContigIntervals
	}

	if queryConfig.RowRanges != nil {
		exportConfig.QueryRowRanges = queryConfig.RowRanges
	}

	if len(queryConfig.Attributes) > 0 {
		exportConfig.Attributes = queryConfig.Attributes
	}

	if len(queryConfig.Filter) > 0 {
		exportConfig.QueryFilter = ptr(queryConfig.Filter)
	}

	data, err := proto.Marshal(&exportConfig)
	if err != nil {
		return false, fmt.Sprintln("marshaling error: ", err)
	}

	len := C.size_t(C.int(len(data)))
	p := C.malloc(len)
	defer C.free(p)
	cBuf := (*[1 << 30]byte)(p)
	copy(cBuf[:], data)

	var status C.status_t
	genomicsdb_handle := C.connect(p, len, &status)
	if status.succeeded == 0 || genomicsdb_handle == nil {
		return false, fmt.Sprintln("Could not connect to GenomicsDB : ", C.GoString(&status.error_message[0]))
	}
	C.query_variant_calls(genomicsdb_handle)
	C.disconnect(genomicsdb_handle)
	return true, ""
}
