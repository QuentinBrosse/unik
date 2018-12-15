# scwgame

## Start

1. [Install golang](https://golang.org/doc/install)
1. Clone this repo in your `GOPATH` and `cd` in the repo
1. Add CA to MacOS Keychain: `open ./infra/data/certifs/local/ca.pem`
1. Run the infra: `cd infra/environments/local && docker-compose build && docker-compose up -d`
1. In another terminal, run local webserver: `cd webapp/ && yarn start`


## Stop

1. Stop the infra: `cd infra/environments/local && docker-compose down -v`
2. Stop the webapp
