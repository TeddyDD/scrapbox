#!/bin/sh

set -eufx

mkdir -p bin
CGO_ENABLED=0 go build -ldflags="-s -w" -o bin github.com/teddydd/scrapbox/cmd/...
