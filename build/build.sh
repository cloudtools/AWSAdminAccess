#!/bin/bash

ARCHITECTURES=amd64
OPERATING_SYSTEMS="linux darwin"
for GOARCH in $ARCHITECTURES; do
    for GOOS in $OPERATING_SYSTEMS; do
        GOOS=$GOOS GOARCH=$GOARCH go build -o AWSAdminAccess-$GOOS-$GOARCH
        gzip -f AWSAdminAccess-$GOOS-$GOARCH
    done
done
