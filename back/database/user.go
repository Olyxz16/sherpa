package database

import (
	"fmt"
    "net/http"

	"github.com/Olyxz16/go-vue-template/logging"
)


type UserAuth struct {
    Uid             int
    Cookie          *http.Cookie
}


func GetUserFromPlatformId(user PlatformUserAuth) (*UserAuth, error) {
    db := dbInstance.db
    q := `SELECT uid, cookie FROM UserAuth
            JOIN PlatformUserAuth ON uid = userId
            WHERE platformId=$1`

    rows, err := db.Query(q, user.PlatformId)
    if err != nil {
        logging.ErrLog(fmt.Sprintf("GetUserFromPlatformId : %v", err))
        return nil, err
    }
    defer rows.Close()

    var result *UserAuth
    if rows.Next() {
        var uid int
        var cookieStr string
        var cookie *http.Cookie
        if err := rows.Scan(&uid, &cookieStr) ; err != nil {
            return nil, err
        }
        if cookie, err = unmarshalCookie(cookieStr) ; err != nil {
            return nil, err    
        }
        result = &UserAuth{ Uid: uid, Cookie: cookie }
    }
    return result, nil
}

// TODO Handle cookie collision
func GetUserOrCreateFromAuth(platformUser PlatformUserAuth) (*UserAuth, bool, error) {
    db := dbInstance.db
    currUser, err := GetUserFromPlatformId(platformUser)
    if err != nil {
        logging.ErrLog(fmt.Sprintf("GetUserOrCreateFromAuth : %v", err))
        return nil, false, err
    }
    if currUser != nil {
        return currUser, false, nil
    }
    
    q := `INSERT INTO UserAuth
        (cookie)
        VALUES ($1)
        ON CONFLICT (cookie) DO NOTHING
        RETURNING (uid)`

    tx, err := db.Begin()
    if err != nil {
        return nil, false, err
    }
    defer tx.Rollback()

    cookie, err := generateUserCookie()
    if err != nil {
        return nil, false, err
    }
    cookieStr, err := marshalCookie(cookie)
    if err != nil {
        return nil, false, err
    }
    rows, err := tx.Query(q, cookieStr)
    if err != nil {
        logging.ErrLog(fmt.Sprintf("GetUserOrCreateFromAuth : %v", err))
        return nil, false, err
    }
    defer rows.Close()

    var uid int
    if rows.Next() {
        if err := rows.Scan(&uid) ; err != nil {
            return nil, false, err
        }
    }

    err = tx.Commit()
    user := &UserAuth{ Uid: uid, Cookie: cookie }
    return user, true, err
}


func init() {
    migrateUserAuth()
}

func migrateUserAuth() {
    New()
    db := dbInstance.db
    q := `CREATE TABLE IF NOT EXISTS UserAuth (
        uid             SERIAL PRIMARY KEY,
        cookie          TEXT UNIQUE
    )`
    _, err := db.Exec(q)
    if err != nil {
        panic(err)
    }
}
