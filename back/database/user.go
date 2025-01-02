package database

import (
	"fmt"
    "net/http"

    "github.com/Olyxz16/go-vue-template/database/utils"
	"github.com/Olyxz16/go-vue-template/logging"
)


type UserAuth struct {
    Uid                 int
    Cookie              *http.Cookie
    EncodedMasterkey    string
    Salt                string
}


func GetUserFromPlatformId(user PlatformUserAuth) (*UserAuth, error) {
    db := dbInstance.db
    q := `SELECT uid, cookie, encodedMasterkey FROM UserAuth
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
        var encodedMasterkey string
        if err := rows.Scan(&uid, &cookieStr, &encodedMasterkey) ; err != nil {
            return nil, err
        }
        if cookie, err = utils.UnmarshalCookie(cookieStr) ; err != nil {
            return nil, err    
        }
        result = &UserAuth{ Uid: uid, Cookie: cookie, EncodedMasterkey: encodedMasterkey }
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

    cookie, err := utils.GenerateUserCookie()
    if err != nil {
        return nil, false, err
    }
    cookieStr, err := utils.MarshalCookie(cookie)
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

func SetUserMasterkey(cookie *http.Cookie, masterkey string) (error) {
    db := dbInstance.db
    q := `UPDATE UserAuth
        SET encodedMasterkey=$1,
        salt=$2
        WHERE cookie=$3`

    tx, err := db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    cookieStr, err := utils.MarshalCookie(cookie)
    if err != nil {
        logging.ErrLog(fmt.Sprintf("SetUserMasterkey : %v", err))
        return err
    }
    encodedMasterkey, b64Salt, _, err := utils.HashFromMasterkey(masterkey)
    if err != nil {
        logging.ErrLog(fmt.Sprintf("SetUserMasterkey : %v", err))
        return err
    }
    rows, err := tx.Query(q, encodedMasterkey, b64Salt, cookieStr)
    if err != nil {
        logging.ErrLog(fmt.Sprintf("SetUserMasterkey : %v", err))
        return err
    }
    defer rows.Close()

    err = tx.Commit()
    return err
}


func init() {
    migrateUserAuth()
}

func migrateUserAuth() {
    New()
    db := dbInstance.db
    q := `CREATE TABLE IF NOT EXISTS UserAuth (
        uid                 SERIAL PRIMARY KEY,
        cookie              TEXT UNIQUE DEFAULT '',
        encodedMasterkey    TEXT        DEFAULT '',
        salt                TEXT        DEFAULT ''
    )`
    _, err := db.Exec(q)
    if err != nil {
        panic(err)
    }
}
