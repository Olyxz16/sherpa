CREATE TABLE IF NOT EXISTS "User" (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL,
    masterkey TEXT DEFAULT '',
    b64salt TEXT DEFAULT '',
    b64filekey TEXT DEFAULT ''
);

CREATE TABLE IF NOT EXISTS "Auth" (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES "User" (id),
    source VARCHAR(255),
    access_token VARCHAR(255),
    expires_in FLOAT,
    refresh_token VARCHAR(255),
    rt_expires_in FLOAT
);

CREATE TABLE IF NOT EXISTS "File" (
    owner_id INT REFERENCES "User" (id),
    source TEXT NOT NULL,
    reponame TEXT NOT NULL,
    filename TEXT NOT NULL,
    b64content TEXT,
    b64nonce TEXT,
    PRIMARY KEY (owner_id, source, reponame, filename)
);
