#!/usr/bin/env bash

IP="172.19.0.10"

echo 'Working against '${IP}'. Starting Apache Cassandra Schema Migrations.'
cqlsh -f 'cassandra/inception.cql'  ${IP}
cqlsh -f 'cassandra/20190325_create_table.up.cql' ${IP}
