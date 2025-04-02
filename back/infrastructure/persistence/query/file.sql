-- name: FindFile :one
SELECT * FROM File
WHERE ownerId = $1 
AND source = $2
AND reponame = $3
AND filename = $4
LIMIT 1;

-- name: FindAllFiles :many
SELECT * FROM File
WHERE ownerId = $1;

-- name: PersistFile :exec
INSERT INTO File (
    ownerId, source, reponame, filename, b64content, b64nonce    
) VALUES (
    $1, $2, $3, $4, $5, $6
);
