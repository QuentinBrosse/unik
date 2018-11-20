#!/usr/bin/env bash
set -e

echo "Run in $APP_ENV env"

echo $#


if [ "$APP_ENV" = "production" ]; then
  server $@;
fi

if [ "$1" = 'server' ]; then
  cd ./pkg/cmd/server && fresh
fi

exec $@
