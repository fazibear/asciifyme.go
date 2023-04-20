#/bin/bash

export GOOS=js
export GOARCH=wasm
mkdir -p build
cp index.html build/
if which -s tinygo
then
  cp "$(tinygo env GOROOT)/misc/wasm/wasm_exec.js" build/
  tinygo build -o build/asciifyme.wasm -target wasm
else
  cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" build/
  go build -o build/asciifyme.wasm
fi
