// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: file.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const findAllFiles = `-- name: FindAllFiles :many
SELECT owner_id, source, reponame, filename, b64content, b64nonce FROM "File"
WHERE owner_id = $1
`

func (q *Queries) FindAllFiles(ctx context.Context, ownerID int32) ([]File, error) {
	rows, err := q.db.Query(ctx, findAllFiles, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []File
	for rows.Next() {
		var i File
		if err := rows.Scan(
			&i.OwnerID,
			&i.Source,
			&i.Reponame,
			&i.Filename,
			&i.B64content,
			&i.B64nonce,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findFile = `-- name: FindFile :one
SELECT owner_id, source, reponame, filename, b64content, b64nonce FROM "File"
WHERE owner_id = $1 
AND source = $2
AND reponame = $3
AND filename = $4
LIMIT 1
`

type FindFileParams struct {
	OwnerID  int32
	Source   string
	Reponame string
	Filename string
}

func (q *Queries) FindFile(ctx context.Context, arg FindFileParams) (File, error) {
	row := q.db.QueryRow(ctx, findFile,
		arg.OwnerID,
		arg.Source,
		arg.Reponame,
		arg.Filename,
	)
	var i File
	err := row.Scan(
		&i.OwnerID,
		&i.Source,
		&i.Reponame,
		&i.Filename,
		&i.B64content,
		&i.B64nonce,
	)
	return i, err
}

const persistFile = `-- name: PersistFile :exec
INSERT INTO "File" (
    owner_id, source, reponame, filename, b64content, b64nonce    
) VALUES (
    $1, $2, $3, $4, $5, $6
)
`

type PersistFileParams struct {
	OwnerID    int32
	Source     string
	Reponame   string
	Filename   string
	B64content pgtype.Text
	B64nonce   pgtype.Text
}

func (q *Queries) PersistFile(ctx context.Context, arg PersistFileParams) error {
	_, err := q.db.Exec(ctx, persistFile,
		arg.OwnerID,
		arg.Source,
		arg.Reponame,
		arg.Filename,
		arg.B64content,
		arg.B64nonce,
	)
	return err
}
