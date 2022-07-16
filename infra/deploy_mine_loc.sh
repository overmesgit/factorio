#!/bin/bash
set -e

TAG="gcr.io/factorio2022/mine:$(find mine -type f -print0 | sort -z | xargs -0 sha1sum | sha1sum | head -c 10)"
echo "$TAG"
docker build -f infra/Dockerfile_mine -t "$TAG" -t gcr.io/factorio2022/mine:latest .
minikube image load --daemon "$TAG"
minikube image load --daemon gcr.io/factorio2022/mine:latest
for name in $(kubectl get deployments -o name | grep -o -P 'r\d+c\d+'); do
  kubectl patch deployment "$name" -p '
{"spec":
  {"template":{"spec":{"containers":[
    {"image": "'"$TAG"'", "name": "'"$name"'"}
    ]
    }}
  }}'
done
