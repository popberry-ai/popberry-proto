#!/bin/sh

GO111MODULE=onbuf go install github.com/bufbuild/buf/cmd/buf@v1.45.0

go env -w GOPRIVATE=github.com/popberry-ai/*