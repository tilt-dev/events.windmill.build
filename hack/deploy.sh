#!/bin/bash
# Deploys events.windmill.build

cd $(dirname $(dirname $0))

set -ex

./hack/build.sh

TODAY=$(date +"%Y-%m-%d")
SECONDS=$(date +"%s")
TAG="$TODAY-$SECONDS"
docker build -t "gcr.io/windmill-prod/events-windmill-build:$TAG" .
docker push "gcr.io/windmill-prod/events-windmill-build:$TAG"

helm upgrade --install events-windmill-build ./chart \
     --set "imageName=gcr.io/windmill-prod/events-windmill-build:$TAG"
