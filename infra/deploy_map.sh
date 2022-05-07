#!/bin/bash

TAG="gcr.io/factorio2022/map:$(date +%Y%m%d%H%M%S)"
docker build -f infra/Dockerfile_map -t "$TAG" .
docker push "$TAG"
kubectl patch deployment map-dep -p '
{"spec":
  {"template":{"spec":{"containers":[
    {"image": "'"$TAG"'", "name": "map"}
    ]
    }}
  }}'
