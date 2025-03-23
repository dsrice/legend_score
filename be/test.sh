#!/bin/bash

go test -v -coverprofile=cover_file.out  ./...
go tool cover -html=cover_file.out -o cover_file.html