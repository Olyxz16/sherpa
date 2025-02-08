package model

import (
	"encoding/base64"
	"net/http"

	"github.com/Olyxz16/sherpa/utils"
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
func FetchFileContent(cookie *http.Cookie, source, repoName, fileName string) (string, error) {
    db := instance.db
    q := `SELECT b64content, b64nonce, b64filekey FROM FileData
            INNER JOIN (
                SELECT uid, b64filekey FROM UserAuth
                WHERE cookie=$1
            ) 
            ON ownerId = uid
            WHERE source=$2
            AND reponame=$3
            AND filename=$4`

    cookieStr, err := utils.MarshalCookie(cookie)
    if err != nil {
        return "", err
    }
    row := db.QueryRow(q, cookieStr, source, repoName, fileName)

    var b64content string
    var b64nonce string
    var b64filekey string
    err = row.Scan(&b64content, &b64nonce, &b64filekey)
    if err != nil {
        return "", err
    }
    
    filekey, err := base64.StdEncoding.DecodeString(b64filekey)
    if err != nil {
        return "", err
    }
    content, err := utils.DecryptFile(filekey, b64nonce, b64content)
    if err != nil {
        return "", err
    }
    return content, nil
}

func SaveFile(cookie *http.Cookie, source, repoName, fileName, content string) error {
    db := instance.db

    user, err := getUserFromCookie(cookie)
    if err != nil {
        return err
    }
    
    filekey, err := base64.StdEncoding.DecodeString(user.B64filekey)
    if err != nil {
        return err
    }
    b64content, b64nonce, err := utils.EncryptFile(filekey, content)
    if err != nil {
        return err
    }

    q := `INSERT INTO FileData
        (ownerId, source, repoName, fileName, b64content, b64nonce)
        VALUES ($1, $2, $3, $4, $5, $6)
        ON CONFLICT (ownerId, source, repoName, fileName)
        DO UPDATE SET
            b64content=EXCLUDED.b64content,
            b64nonce=EXCLUDED.b64nonce`
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    if _, err = tx.Exec(q, user.Uid, source, repoName, fileName, b64content, b64nonce) ; err != nil {
        return err
    }

    if err = tx.Commit() ; err != nil {
        return err
    }
    return nil
}


func init() {
    migrateFileData()
}

func migrateFileData() {
    New()
    db := instance.db

    migrated, err := isUserMigrated()
    if err != nil {
        panic(err)
    }

    if !migrated {
        migrateUserAuth()
    }

    q := `CREATE TABLE IF NOT EXISTS FileData (
        ownerId             INT         REFERENCES UserAuth(uid),
        source              TEXT        NOT NULL,
        repoName            TEXT        NOT NULL,
        filename            TEXT        NOT NULL,
        b64Content          TEXT,
        b64Nonce            TEXT,
        PRIMARY KEY(ownerId, source, repoName, fileName)
    )`
    if _, err = db.Exec(q) ; err != nil {
        panic(err)
    }
}
