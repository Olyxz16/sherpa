package database

import (
	"fmt"
	"net/http"
    
    "github.com/Olyxz16/go-vue-template/database/utils"
	"github.com/Olyxz16/go-vue-template/logging"
)

type PlatformUserAuth struct {
    UserId              int
    PlatformId          int
    Source              string
    Access_token        string
    Refresh_token       string
    Expires_in          int
    Refresh_expires_in  int
}


func AuthenticateUser(auth PlatformUserAuth) (*UserAuth, bool, error) {
    db := dbInstance.db

    user, isNew, err := GetUserOrCreateFromAuth(auth)
    auth.UserId = user.Uid
    if err != nil {
        return nil, false, err
    }
    if !isNew {
        return user, false, nil
    }

    q := `INSERT INTO PlatformUserAuth
    (userId, platformId, source, access_token, expires_in, refresh_token, rt_expires_in)
    VALUES ($1, $2, $3, $4, $5, $6, $7)
    ON CONFLICT (userId, platformId)
    DO UPDATE SET
        source=EXCLUDED.source,
        access_token=EXCLUDED.access_token,
        expires_in=EXCLUDED.expires_in,
        refresh_token=EXCLUDED.refresh_token,
        rt_expires_in=EXCLUDED.rt_expires_in`
    tx, err := db.Begin()
    if err != nil {
        return nil, false, err
    }
    defer tx.Rollback()

    _, err = tx.Exec(q,
        auth.UserId,
        auth.PlatformId,
        auth.Source,
        auth.Access_token, 
        auth.Expires_in, 
        auth.Refresh_token,
        auth.Refresh_expires_in)
    if err != nil {
        logging.ErrLog(fmt.Sprintf("AuthenticateUser : %v", err))
        return nil, false, err
    }

    err = tx.Commit()
    return user, true, err
}

// TODO rows protection when rows length 0
func TokenFromCookie(cookie *http.Cookie, source string) (string, error) {
    db := dbInstance.db
    q := `SELECT access_token FROM PlatformUserAuth
            JOIN UserAuth ON UserId = UserId
            WHERE cookie=$1 AND source=$2`

    cookieStr, err := utils.MarshalCookie(cookie)
    if err != nil {
        return "", err
    }
    rows, err := db.Query(q, cookieStr, source)
    if err != nil {
        return "", err
    }
    defer rows.Close()

    var access_token string
    rows.Next()
    err = rows.Scan(&access_token)
    if err != nil {
        return "", err
    }

    return access_token, nil
}




func init() {
    migrateGithubAuth()
}

func migrateGithubAuth() {
    New()
    db := dbInstance.db

    allowed, err := isUserMigrated()
    if err != nil {
        panic(err)
    }

    if !allowed {
        migrateUserAuth()
    }

    q := `CREATE TABLE IF NOT EXISTS PlatformUserAuth (
    userId          INT          REFERENCES UserAuth(uid),
    platformId      INT,
    source          VARCHAR(255),
    access_token    VARCHAR(255),
    expires_in      FLOAT,
    refresh_token   VARCHAR(255),
    rt_expires_in   FLOAT,
    PRIMARY KEY (userId, platformId)
    )`
    _, err = db.Exec(q)
    if err != nil {
        panic(err)
    }
}
func isUserMigrated() (bool, error) {
    db := dbInstance.db
    q := `SELECT EXISTS (
    SELECT FROM 
    pg_tables
    WHERE 
    schemaname = 'public' AND 
    tablename  = 'userauth'
    );`
    rows, err := db.Query(q)
    if err != nil {
        return false, err
    }
    defer rows.Close()

    var result bool
    rows.Next()
    rows.Scan(&result)

    return result, nil
}
