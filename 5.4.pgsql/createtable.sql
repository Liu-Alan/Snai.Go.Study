CREATE TABLE userinfo(
    uid serial NOT NULL,
    username VARCHAR(64) NOT NULL DEFAULT '',
    department VARCHAR(64) NOT NULL DEFAULT '',
    created timestamp NOT NULL DEFAULT '1970-01-01',
    CONSTRAINT pk_userinfo PRIMARY KEY (uid)
);