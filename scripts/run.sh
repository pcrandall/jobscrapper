#! /bin/bash

#embed config.yml
go-bindata -o config.go config

go run .
