# unik
The unique link

## Run

1. [Install golang](https://golang.org/doc/install)
2. Clone this repo in your `GOPATH` and `cd` in the repo
3. Configure: `./scripts/configure.sh -a -v`
4. Build server: `./scripts/build.sh server`
5. Run the server: `cd infra/environments/local && docker-compose up -d`
6. Run webpack process (in watch mode): `cd webapp/ && yarn watch`
7. In another terminal, run local webserver: `cd webapp/ && yarn serve`
