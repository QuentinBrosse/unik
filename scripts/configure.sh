#!/usr/bin/env bash
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
WEBAPP_DIR="$ROOT_DIR/webapp"
PB_DIR="$ROOT_DIR/protobufs"
WEBAPP_PB_DIR="$WEBAPP_DIR/protobufs"

source ${SCRIPT_DIR}/utils.sh

##
# Log message only if verbose mode is enabled
##
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
      OPT_PROTO=1
  esac
  shift
done

##
# Ensure dependence
##
if [ ! -z ${OPT_DEP} ] ; then
  echoerr "Installing all dep vendors"
  pushd ${ROOT_DIR} > /dev/null
    dep ensure --vendor-only
  popd > /dev/null
fi

##
# Compiling protobufs in:
# - Go (for the server)
# - Js (for the client)
# - Js (for the grpc-web)
##
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
