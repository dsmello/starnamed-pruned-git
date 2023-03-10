#!/bin/sh

# set -eo pipefail


GO_MOD_FILE=$1/go.mod
COSMOS_VERSION=$(grep github.com/cosmos/cosmos-sdk $GO_MOD_FILE|cut -d " " -f 2 | head -n 1)
SWAGGER_URL=https://raw.githubusercontent.com/cosmos/cosmos-sdk/$COSMOS_VERSION/client/docs/swagger-ui/swagger.yaml

SWAGGER_DEST=$1/client/docs/swagger-ui/swagger.yaml

wget $SWAGGER_URL -q -O $SWAGGER_DEST 1> /dev/null