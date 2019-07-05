alter table twitz.twitteraccounts
    add geoEnabled boolean;

alter table twitz.twitteraccounts
    add profileImageUri text;

alter table twitz.twitteraccounts
    add profileBackgroundImageUri text;

alter table twitz.twitteraccounts
    add statusesCount int;
