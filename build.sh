#!/bin/bash

go fmt ./...
go test ./...
GOOS=windows GOARCH=386 go build -o reportgen_x86.exe
