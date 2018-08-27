#!/bin/sh

rm -rf release/*

PACKAGE=main.go
RELEASE_PATH=release/autoexit

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $RELEASE_PATH-darwin $PACKAGE
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $RELEASE_PATH-linux $PACKAGE
CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -o $RELEASE_PATH-freebsd $PACKAGE
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $RELEASE_PATH-windows $PACKAGE

ls -al ./release
