package database

import (
	"errors"
	"net/http"

	"github.com/Olyxz16/go-vue-template/database/utils"
)

type FileData struct {
    OwnerId     int
    Source      string
    RepoName    string
    FileName    string
    B64Content  string
    B64Nonce    string
}

// Error handling for missing cookie, repo or file
func FetchFile(cookie *http.Cookie, repoName, fileName string) (string, error) {
    db := dbInstance.db
    q := `SELECT content FROM FileData
            INNER JOIN (
                SELECT uid FROM UserAuth
                WHERE cookie=$1
            ) 
            ON ownerId = uid
            WHERE reponame=$2
            AND filename=$3`

    cookieStr, err := utils.MarshalCookie(cookie)
    if err != nil {
        return "", err
    }
    rows, err := db.Query(q, cookieStr, repoName, fileName)
    if err != nil {
        /*var pqerr *pq.Error
        var ok bool
        if pqerr, ok = err.(*pq.Error) ; ok {
        }*/
        return "", err
    }
    defer rows.Close()

    var content string
    if !rows.Next() {
        return "", errors.New("Missing data")
    }
    err = rows.Scan(&content)
    if err != nil {
        return "", err
    }

    return content, nil
}

func SaveFile(cookie *http.Cookie, source, repoName, fileName, content string) error {
    db := dbInstance.db

    user, err := getUserFromCookie(cookie)
    if err != nil {
        return err
    }

    encryptedContent, nonce, err := utils.EncryptFile(user.EncodedMasterkey, content)
    if err != nil {
        return err
    }

    q := `INSERT INTO FileData
        (ownerId, source, repoName, fileName, encodedContent, nonce)
        VALUES ($1, $2, $3, $4, $5, $6)`

    tx, err := db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    rows, err := tx.Query(q, user.Uid, source, repoName, fileName, encryptedContent, nonce)
    if err != nil {
        return err
    }
    defer rows.Close()

    var uid int
    if rows.Next() {
        if err := rows.Scan(&uid) ; err != nil {
            return err
        }
    }

    err = tx.Commit()
    if err != nil {
        return err
    }
    return nil
}


func init() {
    migrateFileData()
}

func migrateFileData() {
    New()
    db := dbInstance.db

    migrated, err := isUserMigrated()
    if err != nil {
        panic(err)
    }

    if !migrated {
        migrateUserAuth()
    }

    q := `CREATE TABLE IF NOT EXISTS FileData (
        ownerId             INT         REFERENCES UserAuth(uid),
        repoName            TEXT        DEFAULT '',
        source              TEXT        DEFAULT '',
        filename            TEXT        DEFAULT '.env',
        b64Content          TEXT        DEFAULT '',
        b64Nonce            TEXT        DEFAULT ''
    )`
    _, err = db.Exec(q)
    if err != nil {
        panic(err)
    }
}
