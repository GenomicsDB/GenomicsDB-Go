package main

import (
	"fmt"
	"log"
	"os"

	gdb "github.com/GenomicsDB/GenomicsDB-Go/bindings"
	"github.com/GenomicsDB/GenomicsDB-Go/bindings/protobuf"
)

func main() {
	log.Println("Got GenomicsDB version: ", gdb.GetVersion())

	// small test
	config := gdb.GenomicsDBQueryConfig{
		Workspace: os.Getenv("OMICSDS_WORKSPACE"),
		Array:     os.Getenv("OMICSDS_ARRAY"),
	}
	contig, begin, end, row1, row2 := "1", int64(1), int64(20000), int64(0), int64(2)

	config.ContigIntervals = []*protobuf.ContigInterval{
		{Contig: &contig, Begin: &begin, End: &end}}
	config.RowRanges = []*protobuf.RowRangeList{
		{RangeList: []*protobuf.RowRange{{Low: &row1, High: &row2}}}}
	config.Attributes = []string{"GT", "ALT", "REF"}
	succeeded, errmsg, df := gdb.GenomicsDBQuery(config)

	if !succeeded {
		log.Fatalf("TestQuery failed %v", errmsg)
	}

	fmt.Println(df)
}
