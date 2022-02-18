#!/bin/sh

go build -ldflags "-s -w" -o turtle src/cmd/turtle.go
