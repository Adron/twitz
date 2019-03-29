CREATE KEYSPACE IF NOT EXISTS twitzers
    WITH REPLICATION = {
    'class':'SimpleStrategy',
    'replication_factor' : 1};