#!/usr/bin/env bash

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
