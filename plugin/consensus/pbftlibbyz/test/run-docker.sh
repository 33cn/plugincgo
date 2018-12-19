#!/bin/sh

docker network inspect pbft-net || docker network create --subnet=172.172.0.0/16 pbft-net

docker run --net pbft-net --ip 172.172.0.2 --name replica-1 -dit --rm pbftlibbyz -index=1 -configPath=".."
docker run --net pbft-net --ip 172.172.0.3 --name replica-2 -dit --rm pbftlibbyz -index=2 -configPath=".."
docker run --net pbft-net --ip 172.172.0.4 --name replica-3 -dit --rm pbftlibbyz -index=3 -configPath=".."
docker run --net pbft-net --ip 172.172.0.5 --name replica-4 -dit --rm pbftlibbyz -index=4 -configPath=".."
docker run --net pbft-net --ip 172.172.0.6 --name client -dit --rm pbftlibbyz -index=5 -configPath=".."

docker logs -f -t client