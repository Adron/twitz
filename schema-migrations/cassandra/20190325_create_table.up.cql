create table if not exists twitz.twitterAccounts
(
    id uuid,
    username text,
    primary key (id, username)
);

