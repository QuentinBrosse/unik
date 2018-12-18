# scwgame

## Start

1. [Install golang](https://golang.org/doc/install)
1. Add `GOPATH/bin` to your `PATH`: `export PATH=$PATH:$GOPATH/bin`
1. Install [dep](https://golang.github.io/dep/): `brew install dep`
1. Install [GRPC](https://grpc.io): `go get -u google.golang.org/grpc`
1. Install [protobuf](https://developers.google.com/protocol-buffers/): `brew install protobuf`
1. Install [protoc-gen-go](https://github.com/golang/protobuf/tree/master/protoc-gen-go): `go get -u github.com/golang/protobuf/protoc-gen-go`
1. [Install protoc-gen-grpc-web](https://github.com/grpc/grpc-web#code-generator-plugin)
1. Clone this repo in your `GOPATH` and `cd` in the repo
1. Configure the project: `./scripts/configure.sh -a -v`
1. Add CA to MacOS Keychain: `open ./infra/data/certifs/local/ca.pem` and trust it:
   1. In Keychain, go into the Certificates section and locate the certificate you just added
   1. Double click on it, enter the trust section and under “When using this certificate” select “Always Trust”
1. Run the infra: `cd infra/environments/local && docker-compose build && docker-compose up -d`
1. In another terminal, run local webserver: `cd webapp/ && yarn install && yarn start`


## Stop

1. Stop the infra: `cd infra/environments/local && docker-compose down -v`
2. Stop the webapp
