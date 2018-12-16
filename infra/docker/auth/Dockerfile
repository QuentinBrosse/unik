FROM golang:1.11-stretch

ENV PATH=$PATH:/usr/local/protoc/bin

# Install Dependencies
RUN apt-get update && apt-get install unzip

# Install DEP
RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64         && \
    chmod +x /usr/local/bin/dep

# Install Protoc
RUN wget https://github.com/google/protobuf/releases/download/v3.6.0/protoc-3.6.0-linux-x86_64.zip                  && \
    unzip protoc-3.6.0-linux-x86_64.zip -d /usr/local/protoc                                                        && \
    rm -f protoc-3.6.0-linux-x86_64.zip

# Install protoc-gen-go
RUN git clone -b v1.2.0 --single-branch https://github.com/golang/protobuf $GOPATH/src/github.com/golang/protobuf   && \
    go install github.com/golang/protobuf/protoc-gen-go

# Install protoc-gen-grpc-web
RUN wget https://github.com/grpc/grpc-web/releases/download/1.0.3/protoc-gen-grpc-web-1.0.3-linux-x86_64               \
    -O /usr/local/bin/protoc-gen-grpc-web                                                                           && \
    chmod u+x /usr/local/bin/protoc-gen-grpc-web

WORKDIR /go/src/github.com/quentinbrosse/scwgame

RUN go get github.com/pilu/fresh

COPY ./infra/docker/auth/entrypoint.sh .

ARG app_env
ENV APP_ENV $app_env

COPY ./scripts/utils.sh ./scripts/build.sh ./scripts/configure.sh ./scripts/

COPY ./infra/data/ /data/

# Install dep vendors
COPY ./Gopkg.lock ./Gopkg.toml ./
RUN ./scripts/configure.sh --dep --verbose

# Configure protobufs
COPY ./protobufs/auth ./protobufs/auth
RUN ./scripts/configure.sh --proto --verbose

COPY ./pkg/auth/ ./pkg/auth/
COPY ./pkg/cmd/auth/ ./pkg/cmd/auth/
COPY ./pkg/std/ ./pkg/std/

RUN ./scripts/build.sh auth

ENTRYPOINT ["./entrypoint.sh"]
