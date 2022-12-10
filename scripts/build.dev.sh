#!/bin/bash

# IF YOU UPDATE ANY LDFLAGS HERE, ALSO DO IT IN THE OTHER BUILD FILES IN THIS DIR

set -euo pipefail

PACKAGE="github.com/formulate-dev/cli"
VERSION="$(git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null | sed 's/^.//')"
BUILD_TIMESTAMP=$(date '+%Y-%m-%dT%H:%M:%S')

LDFLAGS=(
  "-X '${PACKAGE}/config.version=${VERSION}'"
  "-X '${PACKAGE}/config.buildTime=${BUILD_TIMESTAMP}'"
  "-X '${PACKAGE}/config.environment=dev'"
)

go build -ldflags="${LDFLAGS[*]}" -o dist/formulate.dev