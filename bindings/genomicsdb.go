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
	"unsafe"

	"github.com/GenomicsDB/GenomicsDB-Go/bindings/protobuf"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
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

func configure(queryConfig GenomicsDBQueryConfig) protobuf.ExportConfiguration {
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

	return exportConfig
}

var df dataframe.DataFrame

func constructDataFrame(genomicsdbQuery unsafe.Pointer) (bool, dataframe.DataFrame) {
	nGenomicFields := C.get_genomic_field_count(genomicsdbQuery)
	var genomicSeries = make([]series.Series, 4+nGenomicFields) // 4 for sample/chrom/pos/end

	nVariantCalls := C.uint64_t(C.get_count(genomicsdbQuery))
	fmt.Println("number of VariantCalls =", nVariantCalls)
	sampleNames := make([]string, nVariantCalls)
	chromosomes := make([]string, nVariantCalls)
	var i, j C.uint64_t
	var info C.info_t
	for i = 0; i < nVariantCalls; i++ {
		if C.get_sample_name(genomicsdbQuery, i, &info) == 1 {
			sampleNames[i] = C.GoString((*C.char)(unsafe.Pointer(info.ptr)))
		} else {
			return false, df
		}
		if C.get_chromosome(genomicsdbQuery, i, &info) == 1 {
			chromosomes[i] = C.GoString((*C.char)(unsafe.Pointer(info.ptr)))
		} else {
			return false, df
		}
	}
	positions := unsafe.Slice((*int)(unsafe.Pointer(C.get_positions(genomicsdbQuery))), nVariantCalls)
	end_positions := unsafe.Slice((*int)(unsafe.Pointer(C.get_end_positions(genomicsdbQuery))), nVariantCalls)

	genomicSeries[0] = series.New(sampleNames, series.String, "Sample")
	genomicSeries[1] = series.New(chromosomes, series.String, "CHROM")
	genomicSeries[2] = series.New(positions, series.Int, "POS")
	genomicSeries[3] = series.New(end_positions, series.Int, "END")

	for i = 0; i < nGenomicFields; i++ {
		if C.get_genomic_field_info(genomicsdbQuery, i, &info) == 1 {
			name := C.GoString(info.name)
			if info.kind == 0 {
				genomic_field := make([]string, nVariantCalls)
				for j = 0; j < nVariantCalls; j++ {
					string_field := C.get_genomic_string_field_at(info.ptr, j)
					if string_field != nil {
						genomic_field[j] = C.GoString((*C.char)(unsafe.Pointer(string_field)))
					} else {
						return false, df
					}
				}
				genomicSeries[i+4] = series.New(genomic_field, series.String, name)
			} else if info.kind == 1 {
				genomic_field := unsafe.Slice((*int)(unsafe.Pointer(info.ptr)), nVariantCalls)
				genomicSeries[i+4] = series.New(genomic_field, series.Int, name)
			} else if info.kind == 2 {
				genomic_field := unsafe.Slice((*float32)(unsafe.Pointer(info.ptr)), nVariantCalls)
				genomicSeries[i+4] = series.New(genomic_field, series.Float, name)
			} else {
				return false, df
			}
		} else {
			return false, df
		}
	}

	return true, dataframe.New(genomicSeries...)
}

func GenomicsDBQuery(queryConfig GenomicsDBQueryConfig) (bool, string, dataframe.DataFrame) {
	config := configure(queryConfig)
	data, err := proto.Marshal(&config)
	if err != nil {
		return false, fmt.Sprintln("marshaling error: ", err), df
	}

	len := C.size_t(C.int(len(data)))
	p := C.malloc(len)
	defer C.free(p)
	cBuf := (*[1 << 30]byte)(p)
	copy(cBuf[:], data)

	var status C.status_t
	genomicsdbHandle := C.connect(p, len, &status)
	if status.succeeded == 0 || genomicsdbHandle == nil {
		return false, fmt.Sprintln("Could not connect to GenomicsDB : ", C.GoString(&status.error_message[0])), df
	}

	genomicsdbQuery := C.query(genomicsdbHandle, &status)
	if status.succeeded == 0 || genomicsdbHandle == nil {
		C.disconnect(genomicsdbHandle)
		return false, fmt.Sprintln("Could not setup GenomicsDB query: ", C.GoString(&status.error_message[0])), df
	}

	succeed, genomicsdb_df := constructDataFrame(genomicsdbQuery)

	C.delete_query(genomicsdbQuery)
	C.disconnect(genomicsdbHandle)

	if !succeed {
		return false, "Exception occurred while constructing data frame", df
	} else {
		return true, "", genomicsdb_df
	}
}
