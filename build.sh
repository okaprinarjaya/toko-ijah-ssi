#!/bin/bash
export GOPATH=$(pwd)
go get github.com/gorilla/mux github.com/mattn/go-sqlite3 github.com/speps/go-hashids
cd src/
go build -o ../toko-ijah
cd ..
export GOPATH=$HOME/go
