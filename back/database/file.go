package database

import (
	"encoding/base64"
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
func FetchFileContent(cookie *http.Cookie, source, repoName, fileName string) (string, error) {
    db := dbInstance.db
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
    content, err := utils.DecryptFile(filekey, b64nonce, b64content)
    return content, nil
}

func SaveFile(cookie *http.Cookie, source, repoName, fileName, content string) error {
    db := dbInstance.db

    user, err := getUserFromCookie(cookie)
    if err != nil {
        return err
    }
    
    filekey, err := base64.StdEncoding.DecodeString(user.B64filekey)
    b64content, b64nonce, err := utils.EncryptFile(filekey, content)
    if err != nil {
        return err
    }

    q := `INSERT INTO FileData
        (ownerId, source, repoName, fileName, b64content, b64nonce)
        VALUES ($1, $2, $3, $4, $5, $6)`

    tx, err := db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    _, err = tx.Exec(q, user.Uid, source, repoName, fileName, b64content, b64nonce)
    if err != nil {
        return err
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
