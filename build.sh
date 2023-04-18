#/bin/bash

export GOOS=js
export GOARCH=wasm
mkdir -p build
cp assets/* build
go build -o build/asciifyme.wasm
