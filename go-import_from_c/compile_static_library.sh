#!/bin/bash
gcc -c foo.c
ar -rcs libfoo.a foo.o
go build -ldflags "-linkmode external -extldflags -static" import.go

