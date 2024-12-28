package database

import (
	"reflect"
	"testing"
)


func TestGetUserFromPlatformID(t *testing.T) {
    New()
    uid := 569991
    cookie, err := generateUserCookie()
    if err != nil {
        t.Fatal("Failed generating cookie")
    }
    platformAuth := PlatformUserAuth {
        UserId: uid,
        PlatformId: 6666622,
        Source: "github.com",
        Access_token: "AAAAAAAAAAA",
        Refresh_token: "RESSSSSSSSS",
        Expires_in: 10000,
        Refresh_expires_in: 10000,
    }
    inputUserAuth := UserAuth{
        Uid: uid,
        Cookie: cookie,
    }
    err = clean()
    if err != nil {
        t.Fatalf("Failed cleaning database : %v", err)
    }
    err = insertUser(inputUserAuth)
    if err != nil {
        t.Fatalf("Failed inserting userAuth : %v", err)
    }
    err = insertPlatform(platformAuth)
    if err != nil {
        t.Fatalf("Failed inserting platformAuth : %v", err)
    }

    actual, err := GetUserFromPlatformId(platformAuth)

    if !reflect.DeepEqual(inputUserAuth, *actual) {
        t.Logf(`
            users not equal :\n
            expected : %v\n
            actual : %v\n`, inputUserAuth, actual)
        t.Fail()
    } 
}

func clean() error {
    db := dbInstance.db
    q := `TRUNCATE TABLE UserAuth CASCADE`
    _, err := db.Exec(q)
    if err != nil {
        return err
    }
    q = `TRUNCATE TABLE PlatformUserAuth CASCADE`
    _, err = db.Exec(q)
    return err
}
func insertPlatform(auth PlatformUserAuth) error {
    db := dbInstance.db
    q := `INSERT INTO PlatformUserAuth
        (userId, platformId, source, access_token, expires_in, refresh_token, rt_expires_in)
        VALUES ($1, $2, $3, $4, $5, $6, $7)`
    _, err := db.Exec(q, auth.UserId, auth.PlatformId, auth.Source, auth.Access_token, auth.Expires_in, auth.Refresh_token, auth.Refresh_expires_in)
    return err
}
func insertUser(user UserAuth) error {
    db := dbInstance.db
    q := `INSERT INTO UserAuth
        (uid, cookie)
        VALUES ($1, $2)`

    cookieStr, err := marshalCookie(user.Cookie)
    if err != nil {
        return err
    }
    _, err = db.Exec(q, user.Uid, cookieStr)
    return err
}
