package database

import (
	"encoding/json"
	"net/http"
)


type UserAuth struct {
    Uid             int
    Cookie          *http.Cookie
}


func GetUserFromPlatformId(user PlatformUserAuth) (*UserAuth, error) {
    db := dbInstance.db
    q := `SELECT uid, cookie FROM UserAuth
            WHERE PlatformUserAuth.id=$1`

    rows, err := db.Query(q, user.PlatformId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var result *UserAuth
    if rows.Next() {
        var uid int
        var cookieStr []byte
        var cookie *http.Cookie
        if err := rows.Scan(&uid, cookieStr) ; err != nil {
            return nil, err
        }
        if err := json.Unmarshal(cookieStr, cookie) ; err != nil {
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
        return nil, false, err
    }
    if currUser != nil {
        return currUser, false, nil
    }
    
    cookie, err := generateUserCookie()
    if err != nil {
        return nil, false, err
    }
    q := `INSERT INTO UserAuth
        (cookie)
        VALUES ($1)
        RETURNING (uid)
        ON CONFLICT (cookie) DO NOTHING`

    tx, err := db.Begin()
    if err != nil {
        return nil, false, err
    }
    defer tx.Rollback()

    cookieStr, err := json.Marshal(cookie)
    rows, err := tx.Query(q, cookieStr)
    if err != nil {
        return nil, false, err
    }

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
        uid             SERIAL PRIMARY KEY
        cookie          BYTEA UNIQUE
    )`
    _, err := db.Exec(q)
    if err != nil {
        panic(err)
    }
}
