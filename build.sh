#!/bin/bash

export GOPATH=$GOPATH:`pwd`:`pwd`/vendor
cd src/github.com/tinhtruong/aur-pkg-status
go build