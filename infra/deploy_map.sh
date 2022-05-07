#!/bin/bash
set -e

TAG="gcr.io/factorio2022/map:$(date +%Y%m%d%H%M%S)"
docker build -f infra/Dockerfile_map -t "$TAG" -t gcr.io/factorio2022/map:latest .
docker push "$TAG"
docker push gcr.io/factorio2022/map:latest
kubectl patch deployment map-dep -p '
{"spec":
  {"template":{"spec":{"containers":[
    {"image": "'"$TAG"'", "name": "map"}
    ]
    }}
  }}'
