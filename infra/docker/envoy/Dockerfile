FROM envoyproxy/envoy:latest

COPY ./infra/docker/envoy/entrypoint.sh .
COPY ./infra/data/ /etc/envoy/data/
COPY ./infra/docker/envoy/envoy.yaml /etc/envoy/envoy.yaml

ENTRYPOINT ["./entrypoint.sh"]
