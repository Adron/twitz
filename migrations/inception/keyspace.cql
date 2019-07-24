create keyspace if not exists twitz
    with replication = {
    'class':'SimpleStrategy',
    'replication_factor' : 1};
