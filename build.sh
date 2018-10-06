#!/usr/bin/env bash

docker-compose down --rmi all
docker-compose up -d --build