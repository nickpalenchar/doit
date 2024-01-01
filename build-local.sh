#!/bin/bash

if ! command -v go &> /dev/null; then
    echo "Go is not installed. Please install Go and try again."
    exit 1
fi

VERSION="$(cat ../version.txt)"

echo "Building version $VERSION"
go build -ldflags "-X main.Version=${VERSION}" -o doit

if [ $? -ne 0 ]; then
    echo "Build failed. Please fix any errors and try again."
    exit 1
fi

sudo mv doit /usr/local/bin/doit

if [ $? -eq 0 ]; then
    echo "doit binary has been successfully installed to /usr/local/bin/doit"
else
    echo "Failed to move doit binary. Please check permissions and try again."
fi