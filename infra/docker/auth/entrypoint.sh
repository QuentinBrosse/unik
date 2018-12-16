#!/usr/bin/env bash
set -e

echo "Run in $APP_ENV env"

if [ "$APP_ENV" = "production" ]; then
  api $@
fi

if [ "$1" = 'api' ]; then
  cd ./pkg/cmd/api && fresh
fi

exec $@
