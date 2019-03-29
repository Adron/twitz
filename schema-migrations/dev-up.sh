#!/usr/bin/env bash

IP="172.19.0.10"
SUBNET="172.19.0.0/16"
GATEWAY="172.19.0.1"

echo "Creating Docker Network"
docker network create --subnet="$SUBNET" --gateway="$GATEWAY" --attachable=true devtwitz

echo "Listing Docker networks."
docker network ls

echo "Inspecting 'devtwitz'"
docker network inspect devtwitz

echo "Creating Apache Cassandra Node."
docker run --name cassandra-twitz --network devtwitz --ip="$IP"  -e CASSANDRA_BROADCAST_ADDRESS="$IP" -it -d cassandra:3.11.4
docker inspect cassandra-twitz
