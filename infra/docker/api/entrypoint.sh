#!/usr/bin/env bash
set -e

echo "Run in $APP_ENV env"

if [ "$APP_ENV" = "production" ]; then
  auth $@
fi

if [ "$1" = 'auth' ]; then
  cd ./pkg/cmd/auth && fresh
fi

exec $@
