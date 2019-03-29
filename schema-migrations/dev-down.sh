#!/usr/bin/env bash

echo "Removing Apache Cassandra Node."
docker stop cassandra-twitz
docker rm cassandra-twitz

echo "Removing Docker dev network."
docker network rm devtwitz

echo "Listing Docker networks."
docker network ls
