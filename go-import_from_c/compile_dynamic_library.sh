#!/bin/bash
gcc -shared -o libfoo.so foo.c
#Put shared library in /usr/lib/
go build import.go

