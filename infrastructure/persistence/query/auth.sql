-- name: PersistAuth :exec
INSERT INTO "Auth" (
    id, user_id, source, access_token, expires_in, refresh_token, rt_expires_in
    ) VALUES (
    $1, $2, $3, $4, $5, $6, $7
    )
    ON CONFLICT (user_id, source)
    DO UPDATE SET
        access_token = EXCLUDED.access_token,
        expires_in = EXCLUDED.expires_in,
        refresh_token = EXCLUDED.refresh_token,
        rt_expires_in = EXCLUDED.rt_expires_in
;


-- name: FindAuthById :one
SELECT * FROM "Auth" a
JOIN "User" u on u.id = a.user_id
WHERE a.id = $1
LIMIT 1;

-- name: FindAuthByUserId :one
SELECT * FROM "Auth"
WHERE user_id = $1 AND source = $2
LIMIT 1;

