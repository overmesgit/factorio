#!/bin/bash

TAG="gcr.io/factorio2022/mine:$(date +%Y%m%d%H%M%S)"
docker build -f infra/Dockerfile_mine -t "$TAG" .
docker push "$TAG"
for name in $(kubectl get deployments -o name | grep -o -P 'r\d+c\d+'); do
  kubectl patch deployment $name -p '
{"spec":
  {"template":{"spec":{"containers":[
    {"image": "'$TAG'", "name": "'$name'"}
    ]
    }}
  }}'
done
