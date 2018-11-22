#!/usr/bin/env bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"

source ${SCRIPT_DIR}/utils.sh

export CGO_ENABLED=0

if [ $# -ne "1" ]
then
  echoerr "Usage: $0 package-name"
  exit 1
fi

echoerr "Building package $1"
go install -ldflags '-extldflags "-static"' "github.com/quentinbrosse/unik/pkg/cmd/$1"
