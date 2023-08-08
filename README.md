# GenomicsDB-Go
Experimental golang bindings to the native [GenomicsDB](https://github.com/GenomicsDB/GenomicsDB) library. Only queries are supported for now. For importing vcf files into GenomicsDB, use the command line tools - `vcf2genomicsdb` or `gatk GenomicsDBImport`.

## For users
Installation : Only Linux and MacOS are supported
### Step 1 - Prerequisites
This step is required to install native genomicsdb if it is not available.
- system prerequisites for native genomicsdb are minimal, see [scripts](https://github.com/GenomicsDB/GenomicsDB/tree/master/scripts/prereqs/system) to install them if needed
- _go get github.com/golang/protobuf/protoc-gen-go_
- verify that $GOPATH is included in the $PATH environment variable
- _go install github.com/GenomicsDB/GenomicsDB-Go/install-genomicsdb@latest_
- $GOPATH/bin/install-genomicsdb
  - By default, the genomicsdb shared library is installed in `/usr/local` and may require `sudo` access. Set environment variable `GENOMICSDB_INSTALL_DIR` to any custom location before invoking `InstallGenomicsDB`.
  - Generates a helper `genomicdb.env` file that exports `PKG_CONFIG_PATH` for building and `DYLD/LD_LIBRARY_PATH` for usage. This file can be sourced if the native GenomicsDB was installed in a custom location.
### Step 2 - Building
- If a custom location was used for installing GenomicsDB in Step 1, set environment variable `PKG_CONFIG_PATH` to `$GENOMICSDB_INSTALL_DIR/lib/pkgconfig/genomicsdb.pc`
- _go get -u github.com/GenomicsDB/GenomicsDB-Go/bindings_ from a `go module`
  
### Example Usage
If a custom location was used for installing GenomicsDB in Step 1, set environment variable `DYLD_LIBRARY_PATH` for MacOS or `LD_LIBRARY_PATH` for Linux to include `$GENOMICSDB_INSTALL_DIR/lib`.
  
```bash
#!/bin/bash
mkdir example-go
pushd example-go
go mod init example/example-go
go get -u github.com/GenomicsDB/GenomicsDB-Go/bindings
cat > main.go << EOF
package main

import (
  "log"

  gdb "github.com/GenomicsDB/GenomicsDB-Go/bindings"
)

func main() {
  log.Println("Got GenomicsDB version: ", gdb.GetVersion())
}
EOF
git build
git run .
popd
```


  

## For release
`GenomicsDB-Go/bindings` and `GenomicsDB-Go/install_genomicsdb` are separate modules, so both bindings and install_genomicsdb folders have to be git tagged separately for `go install` and `go get` to function. For example, create `bindings/v0.0.1` and `install_genomicsdb/v0.0.1` git tags to get the steps for users outlined above to work. Note: If only one of the modules is updated, bump the version only associated with that module.

