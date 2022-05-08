#!/bin/bash
set -e

TAG="gcr.io/factorio2022/mine:$(date +%Y%m%d%H%M%S)"
docker build -f infra/Dockerfile_mine -t "$TAG" -t gcr.io/factorio2022/mine:latest .
docker push "$TAG"
docker push gcr.io/factorio2022/mine:latest
for name in $(kubectl get deployments -o name | grep -o -P 'r\d+c\d+'); do
  kubectl patch deployment $name -p '
{"spec":
  {"template":{"spec":{"containers":[
    {"image": "'$TAG'", "name": "'$name'"}
    ]
    }}
  }}'
done
