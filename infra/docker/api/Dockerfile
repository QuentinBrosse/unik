FROM golang:1.11-stretch

# Install DEP
RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64         && \
    chmod +x /usr/local/bin/dep

WORKDIR /go/src/github.com/quentinbrosse/scwgame

RUN go get github.com/pilu/fresh

COPY ./infra/docker/api/entrypoint.sh .

ARG app_env
ENV APP_ENV $app_env

COPY ./scripts/utils.sh ./scripts/build.sh ./scripts/configure.sh ./scripts/

# Install dep vendors
COPY ./Gopkg.lock ./Gopkg.toml ./
RUN ./scripts/configure.sh --dep --verbose

COPY ./infra/data/ /data/
COPY ./protobufs/ ./protobufs/
COPY ./pkg/api/ ./pkg/api/
COPY ./pkg/cmd/api/ ./pkg/cmd/api/
COPY ./pkg/std/ ./pkg/std/

RUN ./scripts/build.sh api

ENTRYPOINT ["./entrypoint.sh"]
