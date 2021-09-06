#!/usr/bin/env sh

docker run -it --rm --privileged \
  --env ACCESS_TOKEN=$ACCESS_TOKEN \
  --env licensed=demo \
  -v $PWD:/go/src/khorevaa/odin \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -w /go/src/khorevaa/odin \
  goreleaser/goreleaser:latest release --snapshot --skip-publish --rm-dist