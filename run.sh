#!/bin/sh

# node process kill
killall node

# docker run
docker-compose up -d

cd assets
yarn install
yarn run dev &
cd ../

dep ensure

go run main.go