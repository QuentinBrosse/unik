# scwgame

## Run

0. [Install golang](https://golang.org/doc/install)
0. Clone this repo in your `GOPATH` and `cd` in the repo
0. Add CA to MacOS Keychain app by opening `./infra/data/certifs/local/ca.pem`
0. Run the server: `cd infra/environments/local && docker-compose up -d`
0. In another terminal, run local webserver: `cd webapp/ && yarn start`
