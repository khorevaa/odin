#!/usr/bin/env sh

go-bindata -pkg keys -nometadata -o ./license/keys/data.go -ignore keypair.pem -prefix license/cmd/keys ./license/cmd/keys
