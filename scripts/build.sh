#! /bin/bash

#embed config.yml
go-bindata -o config.go config

# GOOS=linux GOARCH=386 go build .
go build .
