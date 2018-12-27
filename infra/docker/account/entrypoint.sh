#!/usr/bin/env bash
set -e

echo "Run in $APP_ENV env"

if [ "$APP_ENV" = "production" ]; then
  account $@
fi

if [ "$1" = 'account' ]; then
  cd ./pkg
  reflex -sr '\.go$' -- go run ./cmd/account/
fi

exec $@
