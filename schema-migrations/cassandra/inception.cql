CREATE KEYSPACE IF NOT EXISTS twitz
    WITH REPLICATION = {
    'class':'SimpleStrategy',
    'replication_factor' : 1};