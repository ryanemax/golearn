#!/bin/sh
rm ./bin/linux/*
rm ./bin/win64/*
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build server.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build client.go
mv client.exe ./bin/win64
mv server.exe ./bin/win64
go build server.go
go build client.go
mv client ./bin/linux
mv server ./bin/linux
