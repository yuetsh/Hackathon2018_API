#!/usr/bin/env bash

docker-compose down
docker images -q | xargs docker rmi
docker-compose up -d --build