# GenomicsDB-Go
Experimental golang bindings to GenomicsDB

## For users
### Step 1 optional
This step is required to install native genomicsdb if it is not available. 
- go install github.com/GenomicsDB/GenomicsDB-Go/install-genomicsdb@latest
- $GOPATH/bin/InstallGenomicsDB
  - By default, the genomicsdb shared library is installed in `/usr/local` and requires `sudo` access. Set environment variable `GENOMICSDB_INSTALL_DIR` to any  location before invoking `InstallGenomicsDB`.
### Step 2
- Optional if a custom location was used for installing GenomicsDB in Step 1, set `PKG_CONFIG_PATH` to `$GENOMICSDB_INSTALL_DIR/lib/pkgconfig/genomicsdb.pc`
- go get -u github.com/GenomicsDB/GenomicsDB-Go/bindings

## For release
`GenomicsDB-Go/bindings` and `GenomicsDB-Go/install_genomicsdb` are separate modules, so both bindings and install_genomicsdb folders have to be tagged separately. For example, create `bindings/v0.0.1` and `install_genomicsdb/v0.0.1` to get the steps for users outlined above to work. Note: If only one of the modules is updated, bump the version only associated with that module.

