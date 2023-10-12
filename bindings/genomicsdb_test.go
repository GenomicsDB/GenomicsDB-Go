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

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/GenomicsDB/GenomicsDB-Go/bindings/protobuf"
)

func TestGenomicsDBVersion(t *testing.T) {
	if !strings.HasPrefix(GetVersion(), "1.5") {
		t.Fatal("GenomicsDB Version is not 1.5")
	}
}

func createTestQueryConfig(workspace string, array string) GenomicsDBQueryConfig {
	return GenomicsDBQueryConfig{
		Workspace: workspace,
		Array:     array,
	}
}

func TestQueryNonExistentWorkspace(t *testing.T) {
	succeeded, errMsg, _ := GenomicsDBQuery(createTestQueryConfig("non-existent-ws", "non-existent-array"))
	if succeeded {
		t.Fatal("TestQueryNonExistentWorkspace should not succeed")
	} else if len(errMsg) == 0 {
		t.Fatal("TestQueryNonExistentWorkspace should return error message")
	}
	if len(errMsg) == 0 {
		t.Fatal("No error message from failed Query")
	}
}

func TestQueryWorkspaceWithInvalidURL(t *testing.T) {
	succeeded, errMsg, _ := GenomicsDBQuery(createTestQueryConfig("azb://my Container/myPath?endpoint=foo.blob.core.windows.net", "non-existent-array"))
	if succeeded {
		t.Fatal("TestQueryNonExistentWorkspace should not succeed")
	} else if len(errMsg) == 0 {
		t.Fatal("TestQueryNonExistentWorkspace should return error message")
	}
	if len(errMsg) == 0 {
		t.Fatal("No error message from failed Query")
	}
}

func TestQueryWorkspaceWithNonExistentWorkspaceURL(t *testing.T) {
	succeeded, errMsg, _ := GenomicsDBQuery(createTestQueryConfig("azb://myContainer/non-existent-workspace?endpoint=foo.blob.core.windows.net", "non-existent-array"))
	if succeeded {
		t.Fatal("TestQueryNonExistentWorkspace should not succeed")
	} else if len(errMsg) == 0 {
		t.Fatal("TestQueryNonExistentWorkspace should return error message")
	}
	if len(errMsg) == 0 {
		t.Fatal("No error message from failed Query")
	}
}

func TestQuery(t *testing.T) {
	config := createTestQueryConfig("test-ws", "allcontigs$1$3101976562")
	config.ContigIntervals = []*protobuf.ContigInterval{
		{Contig: ptr("1"), Begin: ptr(int64(1)), End: ptr(int64(20000))}}
	config.RowRanges = []*protobuf.RowRangeList{
		{RangeList: []*protobuf.RowRange{{Low: ptr(int64(0)), High: ptr(int64(2))}}}}
	config.Attributes = []string{"GT", "DP"}
	config.Filter = "REF == \"G\" && resolve(GT,REF,ALT) &= \"T/T\" && ALT |= \"T\""
	succeeded, errMsg, df := GenomicsDBQuery(config)

	if !succeeded {
		t.Fatal("TestQuery failed: " + errMsg)
	}

	fmt.Println(df)
}

func createTestQueryConfigWithVidMappingAndCallsetMappingFiles(workspace string, array string) GenomicsDBQueryConfig {
	return GenomicsDBQueryConfig{
		Workspace:          workspace,
		Array:              array,
		VidMappingFile:     filepath.Join(workspace, "vidmap.json"),
		CallsetMappingFile: filepath.Join(workspace, "callset.json"),
	}
}

func TestQueryWithVidMappingAndCallsetMappingFiles(t *testing.T) {
	config := createTestQueryConfigWithVidMappingAndCallsetMappingFiles("test-ws", "allcontigs$1$3101976562")
	config.ContigIntervals = []*protobuf.ContigInterval{
		{Contig: ptr("1"), Begin: ptr(int64(1)), End: ptr(int64(20000))}}
	config.RowRanges = []*protobuf.RowRangeList{
		{RangeList: []*protobuf.RowRange{{Low: ptr(int64(0)), High: ptr(int64(2))}}}}
	config.Attributes = []string{"GT", "DP"}
	config.Filter = "REF == \"G\" && resolve(GT, REF, ALT) &= \"T/T\" && ALT |= \"T\""
	succeeded, errMsg, df := GenomicsDBQuery(config)

	if !succeeded {
		t.Fatal("TestQuery failed:" + errMsg)
	}

	fmt.Println(df)
}

func TestGenomicsDBDemoData(t *testing.T) {
	genomicsDBDemoWS := os.Getenv("GENOMICSDB_DEMO_WS")
	if len(genomicsDBDemoWS) > 0 {
		config := createTestQueryConfig(genomicsDBDemoWS, "allcontigs$1$3095677412")
		config.ContigIntervals = []*protobuf.ContigInterval{
			{Contig: ptr("17"), Begin: ptr(int64(7571719)), End: ptr(int64(7590868))}}
		config.RowRanges = []*protobuf.RowRangeList{
			{RangeList: []*protobuf.RowRange{{Low: ptr(int64(0)), High: ptr(int64(200000))}}}}
		config.Attributes = []string{"REF", "ALT", "GT"}
		config.Filter = "REF==\"A\" && ALT|=\"T\" && resolve(GT,REF,ALT)&=\"T/T\""
		succeeded, errMsg, df := GenomicsDBQuery(config)
		if !succeeded {
			t.Fatal("TestQuery failed: " + errMsg)
		}
		fmt.Println(df)
	} else {
		t.Log("Skipping TestGenomicsDBDemoData. Set env GENOMICSDB_DEMO_WS to run this test")
	}
}
