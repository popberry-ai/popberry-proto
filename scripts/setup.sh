#!/bin/sh

GO111MODULE=on GOBIN=/usr/local/bin go install github.com/bufbuild/buf/cmd/buf@latest

go env -w GOPRIVATE=github.com/popberry-ai/*