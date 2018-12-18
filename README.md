# scwgame

## Start

1. [Install golang](https://golang.org/doc/install)
1. Install [dep](https://golang.github.io/dep/): `brew install dep`
1. [Install GRPC and protoc-gen-go](https://grpc.io/docs/quickstart/go.html#install-grpc): `brew install protobuf`
1. [Install protoc-gen-grpc-web](https://github.com/grpc/grpc-web#code-generator-plugin)
1. Clone this repo in your `GOPATH` and `cd` in the repo
1. Configure the project: `./scripts/configure.sh -a -v`
1. Add CA to MacOS Keychain: `open ./infra/data/certifs/local/ca.pem`
1. Run the infra: `cd infra/environments/local && docker-compose build && docker-compose up -d`
1. In another terminal, run local webserver: `cd webapp/ && yarn start`


## Stop

1. Stop the infra: `cd infra/environments/local && docker-compose down -v`
2. Stop the webapp
