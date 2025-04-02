-- name: PersistUser :exec
INSERT INTO UserData (
    uid, username, masterkey, b64salt, b64filekey
) VALUES (
    $1, $2, $3, $4, $5
)
ON CONFLICT (uid)
DO UPDATE SET
    username = EXCLUDED.username,
    masterkey = EXCLUDED.masterkey,
    b64salt = EXCLUDED.b64salt,
    b64filekey = EXCLUDED.b64filekey;

-- name: FindUser :one
SELECT * FROM UserData
WHERE uid = $1 LIMIT 1;

-- name: UpdateMasterkey :exec
UPDATE UserData
SET masterkey=$2,
b64salt=$3,
b64filekey=$4
WHERE uid=$1;
