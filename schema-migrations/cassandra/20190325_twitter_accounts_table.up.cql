create keyspace if not exists twitz
    with replication = {
    'class':'SimpleStrategy',
    'replication_factor' : 1};

create table if not exists twitz.twitterAccounts
(
    id uuid,
    username text,
    email text,
    primary key (username, email)
);

