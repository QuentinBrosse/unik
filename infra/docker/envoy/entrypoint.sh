#!/bin/sh
set -e

if [ $# -eq 0 ]
then
    cat /etc/envoy/envoy.yaml
    /usr/local/bin/envoy --v2-config-only --mode validate -c /etc/envoy/envoy.yaml || exit 1
    exec /usr/local/bin/envoy --v2-config-only -c /etc/envoy/envoy.yaml
fi

exec "$@"
