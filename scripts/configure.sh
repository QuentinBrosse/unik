#!/usr/bin/env bash
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
WEBAPP_DIR="$ROOT_DIR/webapp"
PB_DIR="$ROOT_DIR/protobufs"
WEBAPP_PB_DIR="$WEBAPP_DIR/protobufs"

##
# Echo on stderr
##
function echoerr ()
{
  echo "$@" >&2
}

##
# Colorize output
##
function color() {
  case $1 in
    yellow) echoerr -e -n "\033[33m"   ;;
    green)  echoerr -e -n "\033[32m"   ;;
    red)    echoerr -e -n "\033[0;31m" ;;
  esac
  echoerr "$2"
  echoerr -e -n "\033[0m"
}

##
# Print Usage
##
function usage() {
  color yellow "Usage:"
  echoerr "  ${BASH_SOURCE[0]} [OPTIONS]"
  echoerr ""

  color yellow "Options:"

  color green "  -a, --all"
  echoerr -e "\tAll"

  color green "  -d, --dep"
  echoerr -e "\tInstall dep dependencies"

  color green "  -p, --proto"
  echoerr -e "\tCompile proto files"

  color green "  -h, --help"
  echoerr -e "\tDisplay this help"

  color green "  -v, --verbose"
  echoerr -e "\tEnable verbose mode"

  echoerr ""
  exit $1;
}

function log() {
  [ ! -z ${OPT_VERBOSE} ] && echo $@
}

OPT_DEP=""
OPT_PROTO=""
OPT_VERBOSE=""

##
# Parse arguments
##
while [[ $# > 0 ]]
do
  case "$1" in
    -d|--dep) OPT_DEP=1 ;;
    -p|--proto) OPT_PROTO=1 ;;
	  -h|--help) usage ;;
    -v|--verbose) OPT_VERBOSE="-v" ;;
    -a|--all)
      OPT_DEP=1
  esac
  shift
done

if [ ! -z ${OPT_DEP} ] ; then
  echoerr "Installing all dep vendors"
  pushd ${ROOT_DIR} > /dev/null
    dep ensure --vendor-only
  popd > /dev/null
fi

if [ ! -z ${OPT_PROTO} ] ; then
  echoerr "Compilling proto files"
  pushd ${ROOT_DIR} > /dev/null
    mkdir -p ${WEBAPP_PB_DIR}
    protoc \
      -I ${PB_DIR} \
      --go_out=plugins=grpc:${PB_DIR} \
      --js_out=import_style=commonjs:${WEBAPP_PB_DIR} \
      --grpc-web_out=import_style=commonjs,mode=grpcwebtext:${WEBAPP_PB_DIR} \
      ${PB_DIR}/api.proto
  popd > /dev/null
fi
