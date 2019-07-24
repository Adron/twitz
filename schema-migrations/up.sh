#!/usr/bin/env bash

IP="172.19.0.10"

echo 'Working against '${IP}'. Starting Apache Cassandra Schema Migrations.'

echo "migrate -source file://cassandra/ -database cassandra://$IP/twitz up"
migrate -source file://cassandra/ -database cassandra://$IP/twitz up
