#!/bin/bash
mkdir -p bin
go build -o bin/tippecanoe-concat-array src/main.go

# Copy additional files to the output directory
# cp -r some_directory/* output_directory/