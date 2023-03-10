#!/usr/bin/env bash

set -eo pipefail

DOCKERNAME_x86_64_alpine="starnamed:alpine"
DOCKERNAME_x86_64_buster="starnamed:buster"

docker build -t $DOCKERNAME_x86_64_alpine .

mkdir -p ./release

docker run --rm -tti --entrypoint="cp" -v $(pwd)/release:/release $DOCKERNAME_x86_64_alpine /usr/bin/starnamed /release/starnamed.linux.amd64

