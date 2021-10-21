#!/bin/sh

mkdir "Releases"

# 【darwin/amd64】
echo "start build darwin/amd64 ..."
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build  -o ./Releases/AKP-darwin-amd64 main.go

# 【linux/amd64】
echo "start build linux/amd64 ..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o ./Releases/AKP-linux-amd64 main.go

echo "Congratulations,all build success!"