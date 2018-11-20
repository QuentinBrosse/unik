FROM golang:1.11-stretch

# Install DEP
RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64         && \
    chmod +x /usr/local/bin/dep

WORKDIR /go/src/github.com/quentinbrosse/unik

RUN go get github.com/pilu/fresh

COPY ./infra/docker/server/entrypoint.sh .

ARG app_env
ENV APP_ENV $app_env

COPY ./scripts/build.sh ./scripts/configure.sh ./scripts/


# Install dep vendors
COPY ./Gopkg.lock ./Gopkg.toml ./
RUN ./scripts/configure.sh --dep --verbose

COPY ./protobufs ./protobufs/
COPY ./pkg/server ./pkg/server/
COPY ./pkg/cmd/server ./pkg/cmd/server/

RUN ./scripts/build.sh server

ENTRYPOINT ["./entrypoint.sh"]
