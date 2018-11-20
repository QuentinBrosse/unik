FROM envoyproxy/envoy:latest

COPY ./infra/docker/envoy/envoy.yaml /etc/envoy/envoy.yaml

CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml
