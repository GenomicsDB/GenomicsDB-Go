package main

import (
  "log"

  gdb "github.com/GenomicsDB/GenomicsDB-Go/bindings"
)

func main() {
  log.Println("Got GenomicsDB version: ", gdb.GetVersion())
}
