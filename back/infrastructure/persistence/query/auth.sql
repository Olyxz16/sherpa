-- name: PersistAuth :exec
INSERT INTO Auth (
    uid, userId, source, access_token, expires_in, refresh_token, rt_expires_in
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
ON CONFLICT DO NOTHING;

-- name: FindAuthById :one
SELECT * FROM Auth a
JOIN UserData u on u.uid = a.userId
WHERE a.uid = $1
LIMIT 1;

-- name: FindAuthByUserId :one
SELECT * FROM Auth
WHERE userId = $1 AND source = $2
LIMIT 1;

