#!/usr/bin/env bash

IP="172.22.0.5"

echo 'Working against '${IP}'. Starting Apache Cassandra Schema Migrations.'

echo "migrate -source file://cassandra/ -database cassandra://$IP/twitz down"
migrate -source file://cassandra/ -database cassandra://$IP/twitz down
