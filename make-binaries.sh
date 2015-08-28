#!/usr/bin/env bash
# Generate binaries for [darwin, windows]x [386, amd64]
set -e
set -v
GOOS=darwin GOARCH=386 go build -o bin/hearthstone-json-i386-macosx
GOOS=darwin GOARCH=amd64 go build -o bin/hearthstone-json-x86_64-macosx
GOOS=windows GOARCH=386 go build -o bin/hearthstone-json-i386-win32.exe
GOOS=windows GOARCH=amd64 go build -o bin/hearthstone-json-x86_64-win32.exe