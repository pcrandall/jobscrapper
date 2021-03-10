#! /bin/bash
#embed config.yml
packr2
GOOS=windows GOARCH=386 go build .
