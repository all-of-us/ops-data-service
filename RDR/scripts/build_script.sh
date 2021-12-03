#!/usr/bin/env bash
for GOOS in darwin linux windows; do
    for GOARCH in amd64; do
        if [ $GOOS = "windows" ]; then
            GOARCH+='.exe'
        fi
        go build -v -o ops-data-$GOOS-$GOARCH
    done
done
