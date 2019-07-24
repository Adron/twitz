create table if not exists twitz.twitterAccounts
(
    id uuid,
    username text,
    email text,
    primary key (username, email)
);

