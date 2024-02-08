#!/bin/bash

export CGO_ENABLED=0

export VERSION="1.2.3"
export NAME="MyApp"
export GIT_COMMIT=$(git rev-parse HEAD)

echo
echo CGO_ENABLED=$CGO_ENABLED
echo VERSION=$VERSION
echo NAME=$NAME
echo GIT_COMMIT=$GIT_COMMIT
echo