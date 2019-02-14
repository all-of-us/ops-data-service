#!/usr/bin/env bash
for GOOS in darwin linux windows; do
    for GOARCH in 386 amd64; do
        go build -v -o ops-data-$GOOS-$GOARCH
    done
done
