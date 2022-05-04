#!/bin/bash

docker build -f infra/Dockerfile_mine -t 35.243.65.153:8080/overmes/mine:dev .
docker push 35.243.65.153:8080/overmes/mine:dev