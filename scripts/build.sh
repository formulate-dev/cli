#!/bin/bash

# IF YOU UPDATE ANY LDFLAGS HERE, ALSO DO IT IN THE OTHER BUILD FILES IN THIS DIR
#
# This is the only file that builds the frontend too, on purpose.

set -euo pipefail

PACKAGE="github.com/formulate-dev/cli"
VERSION="$(git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null | sed 's/^.//')"
BUILD_TIMESTAMP=$(date '+%Y-%m-%dT%H:%M:%S')

LDFLAGS=(
  "-X '${PACKAGE}/config.version=${VERSION}'"
  "-X '${PACKAGE}/config.buildTime=${BUILD_TIMESTAMP}'"
  "-X '${PACKAGE}/config.environment=production'"
)

STATIC_LDFLAGS=(
  '-linkmode external -extldflags "-static"'
)

if [ "$(uname -s)" == "Linux" ]; then
( CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=musl-gcc go build -tags="embedfrontend" -ldflags="${LDFLAGS[*]} ${STATIC_LDFLAGS[*]}" -o dist/formulate.amd64.bin ) &
( CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CC=~/cross/aarch64-linux-musl-cross/bin/aarch64-linux-musl-gcc go build -tags="embedfrontend" -ldflags="${LDFLAGS[*]} ${STATIC_LDFLAGS[*]}" -o dist/formulate.arm64.bin ) &
( CGO_ENABLED=1 GOOS=linux GOARCH=arm CC=~/cross/arm-linux-musleabihf-cross/bin/arm-linux-musleabihf-gcc   go build -tags="embedfrontend" -ldflags="${LDFLAGS[*]} ${STATIC_LDFLAGS[*]}" -o dist/formulate.arm.bin ) &
wait -n
wait -n
wait -n
fi

# NOTE: These don't appear to be dynamic binaries, and I'm not sure what the implications are. 
#       This may need multiple version per major macOS release, for example.
if [ "$(uname -s)" == "Darwin" ]; then
CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build -tags="embedfrontend" -ldflags="${LDFLAGS[*]}" -o dist/formulate-macos-arm64.bin
CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -tags="embedfrontend" -ldflags="${LDFLAGS[*]}" -o dist/formulate-macos-amd64.bin
fi

chmod +x dist/*

# Compress
pushd dist
ls *.bin | xargs -I{} -P2 tar czvf {}.tar.gz {}
popd