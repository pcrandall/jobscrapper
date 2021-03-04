#! /bin/bash

#embed config.yml
go-bindata -o config.go config

GOOS=windows GOARCH=386 go build .
