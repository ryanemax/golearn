#!/bin/sh
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build server.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build client.go
mv client.exe ./bin/win64
mv server.exe ./bin/win64
go build server.go
go build client.go
mv client ./bin/win64
mv server ./bin/win64
