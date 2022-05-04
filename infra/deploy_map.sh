#!/bin/bash

docker build -f infra/Dockerfile_map -t 35.243.65.153:8080/overmes/map:dev .
docker push 35.243.65.153:8080/overmes/map:dev