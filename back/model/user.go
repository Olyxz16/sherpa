package model

import (
	"encoding/base64"
	"net/http"

    "go.uber.org/zap"

	"github.com/Olyxz16/sherpa/utils"
)


type UserAuth struct {
    Uid                 int
    Cookie              *http.Cookie
    EncodedMasterkey    string
    Salt                string
    B64filekey             string
}


func getUserFromCookie(cookie *http.Cookie) (*UserAuth, error) {
    db := instance.db
    
    q := `SELECT uid, encodedMasterkey, salt, b64filekey FROM UserAuth
            WHERE cookie=$1`

    cookieStr, err := utils.MarshalCookie(cookie)
    if err != nil {
        return nil, err
    }

    row := db.QueryRow(q, cookieStr)

    var userAuth UserAuth
    if err := row.Scan(&userAuth.Uid, &userAuth.EncodedMasterkey, &userAuth.Salt, &userAuth.B64filekey) ; err != nil {
        return nil, err
    }
    userAuth.Cookie = cookie

    return &userAuth, nil
}


func GetUserFromPlatformId(user PlatformUserAuth) (*UserAuth, error) {
    db := instance.db
    q := `SELECT uid, cookie, encodedMasterkey, b64filekey FROM UserAuth
            JOIN PlatformUserAuth ON uid = userId
            WHERE platformId=$1`

    row := db.QueryRow(q, user.PlatformId)

    var result UserAuth
    var cookieStr string
    if err := row.Scan(&result.Uid, &cookieStr, &result.EncodedMasterkey, &result.B64filekey) ; err != nil {
        return nil, nil
    }

    if cookie, err := utils.UnmarshalCookie(cookieStr) ; err == nil {
        result.Cookie = cookie
    } else {
        return nil, err
    }

    return &result, nil
}

// TODO Handle cookie collision
func GetUserOrCreateFromAuth(platformUser PlatformUserAuth) (*UserAuth, bool, error) {
    db := instance.db
    currUser, err := GetUserFromPlatformId(platformUser)
    if err != nil {
        zap.L().DPanic("GetUserOrCreateFromAuth", zap.Error(err))
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
    row := tx.QueryRow(q, cookieStr)

    var uid int
    if err := row.Scan(&uid) ; err != nil {
        return nil, false, err
    }

    if err := tx.Commit() ; err != nil {
        return nil, false, err 
    }
    user := &UserAuth{ Uid: uid, Cookie: cookie }
    return user, true, nil
}

func SetUserMasterkey(cookie *http.Cookie, masterkey string) (error) {
    db := instance.db
    q := `UPDATE UserAuth
        SET encodedMasterkey=$1,
        salt=$2,
        b64filekey=$3
        WHERE cookie=$4`

    tx, err := db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    cookieStr, err := utils.MarshalCookie(cookie)
    if err != nil {
        zap.L().Error("SetUserMasterkey", zap.Error(err))
        return err
    }
    encodedMasterkey, b64Salt, b64Hash, err := utils.HashFromMasterkey(masterkey)
    if err != nil {
        zap.L().Error("SetUserMasterkey", zap.Error(err))
        return err
    }
    hash, err := base64.StdEncoding.DecodeString(b64Hash)
    if err != nil {
        return err
    }
    _, _, b64Filekey, err := utils.HashFromMasterkey(string(hash))
    if err != nil {
        return err
    }

    if _, err := tx.Exec(q, encodedMasterkey, b64Salt, b64Filekey, cookieStr) ; err != nil {
        zap.L().Error("SetUserMasterkey", zap.Error(err))
        return err
    }

    if err := tx.Commit() ; err != nil {
        return err
    }
    return nil
}


func init() {
    migrateUserAuth()
}

func migrateUserAuth() {
    New()
    db := instance.db
    q := `CREATE TABLE IF NOT EXISTS UserAuth (
        uid                 SERIAL PRIMARY KEY,
        cookie              TEXT UNIQUE DEFAULT '',
        encodedMasterkey    TEXT        DEFAULT '',
        salt                TEXT        DEFAULT '',
        b64filekey          TEXT        DEFAULT ''
    )`
    _, err := db.Exec(q)
    if err != nil {
        panic(err)
    }
}

func isUserMigrated() (bool, error) {
    db := instance.db
    q := `SELECT EXISTS (
    SELECT FROM
    pg_tables
    WHERE
    schemaname = 'public' AND
    tablename  = 'userauth'
    );`
    row := db.QueryRow(q)

    var result bool
    if err := row.Scan(&result) ; err != nil {
        panic(err)
    }

    return result, nil
}
