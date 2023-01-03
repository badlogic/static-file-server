#!/bin/bash
GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" static-server.go
mv static-server static-server-macos
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" static-server.go
mv static-server static-server-linux
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" static-server.go
mv static-server.exe static-server-windows.exe