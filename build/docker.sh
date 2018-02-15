#!/bin/bash
wd=/go/src/github.com/fabiolb/fabio

docker run --rm -it -e CGO_ENABLED=0 -v $(pwd):$wd -w $wd golang:latest go build -a -ldflags '-s' . 

