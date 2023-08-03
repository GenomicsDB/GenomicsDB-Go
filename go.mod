module github.com/datma-health/GenomicsDB-Go

go 1.20

replace github.com/datma-health/GenomicsDB-Go/install-genomicsdb => ./install-genomicsdb

require (
	github.com/datma-health/GenomicsDB-Go/install-genomicsdb v0.0.0
	google.golang.org/protobuf v1.31.0 // indirect
)
