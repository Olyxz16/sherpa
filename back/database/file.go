package database

import (
	"errors"
	"net/http"

	"github.com/Olyxz16/go-vue-template/database/utils"
)

type FileData struct {
    ownerId     int
    source      string
    repoName    string
    fileName    string
    content     string
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

func SaveFile(cookie *http.Cookie, repoName, fileName, content string) error {
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
        content             TEXT        DEFAULT ''
    )`
    _, err = db.Exec(q)
    if err != nil {
        panic(err)
    }
}
