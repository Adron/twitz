#!/usr/bin/env bash

IP="172.22.0.5"

./cassierunner -a $IP -q "create keyspace if not exists twitz with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };"

echo 'Working against '${IP}'. Starting Apache Cassandra Schema Migrations.'

echo "migrate -source file://migrations/ -database cassandra://$IP/twitz up"
migrate -source file://migrations/ -database cassandra://$IP/twitz up
