CREATE TABLE IF NOT EXISTS UserData (
    uid         SERIAL  PRIMARY KEY,
    username    TEXT    NOT NULL,
    masterkey   TEXT    DEFAULT '',
    b64Salt     TEXT    DEFAULT '',
    b64Filekey  TEXT    DEFAULT ''
);

CREATE TABLE IF NOT EXISTS Auth (
    uid             SERIAL,
    userId          INT     REFERENCES UserData(uid),
    source          VARCHAR(255),
    access_token    VARCHAR(255),
    expires_in      FLOAT,
    refresh_token   VARCHAR(255),
    rt_expires_in   FLOAT,
    PRIMARY KEY (uid, userId)
);

CREATE TABLE IF NOT EXISTS File (
    ownerId     INT     REFERENCES UserData(uid),
    source      TEXT    NOT NULL,
    reponame    TEXT    NOT NULL,
    filename    TEXT    NOT NULL,
    b64content  TEXT,
    b64nonce    TEXT,
    PRIMARY KEY(ownerId, source, repoName, fileName)
);
