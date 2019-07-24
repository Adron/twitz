alter table twitz.twitteraccounts
    add (name text,
    createAt text,
    description text,
    followersCount int,
    friendsCount int,
    following boolean,
    twitterID bigint,
    twitterIDstr text,
    listedCount int,
    location text);