package database

import (
	"encoding/base64"
	"math/rand"
	"reflect"
	"testing"

	"github.com/Olyxz16/sherpa/database/utils"
)

/*************************/
/* GetUserFromPlatformID */
/*************************/

func TestGetUserFromPlatformID(t *testing.T) {
    New()
    uid := 569991
    cookie, err := utils.GenerateUserCookie()
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
    if err != nil {
        t.Fatalf("GetUserFromPlatformId failed : %v", err)
    }

    if !reflect.DeepEqual(inputUserAuth, *actual) {
        t.Logf(`
            users not equal :\n
            expected : %v\n
            actual : %v\n`, inputUserAuth, actual)
        t.Fail()
    } 
}

func GetNonExistentUser(t *testing.T) {
    New()
    platformAuth := PlatformUserAuth {
        UserId: 5668988,
        PlatformId: 1222333,
        Source: "github.com",
        Access_token: "IJNYIHYHY",
        Refresh_token: "AZEEEEEE",
        Expires_in: 10000,
        Refresh_expires_in: 10000,
    }
    err := clean()
    if err != nil {
        t.Fatalf("Failed cleaning database : %v", err)
    }
    err = insertPlatform(platformAuth)
    if err != nil {
        t.Fatalf("Failed inserting platformAuth : %v", err)
    }

    actual, err := GetUserFromPlatformId(platformAuth)
    if err != nil {
        t.Fatalf("GetUserFromPlatformId failed: %v", err)
    }

    if actual != nil {
        t.Logf(`
            invalid result :\n
            expected: %v\n
            actual: %v\n`, nil, actual)
    }

}


/***************************/
/* GetUserOrCreateFromAuth */
/***************************/

func TestCreateUser(t *testing.T) {
    New()
    platformAuth := PlatformUserAuth {
        UserId: 5668988,
        PlatformId: 1222333,
        Source: "github.com",
        Access_token: "IJNYIHYHY",
        Refresh_token: "AZEEEEEE",
        Expires_in: 10000,
        Refresh_expires_in: 10000,
    }
    err := clean()
    if err != nil {
        t.Fatalf("Failed cleaning database : %v", err)
    }

    _, isNew, err := GetUserOrCreateFromAuth(platformAuth)

    if err != nil {
        t.Fatalf("Failed GetUserOrCreateFromAuth : %v", err)
    }
    if !isNew {
        t.Fatalf("isNew : actual %v | expected %v", isNew, true)
    }
}


/*************/
/*   Utils   */
/*************/

func mockUserAuth() (*UserAuth, error) {
    userId := rand.Int() % 100
    inputkey, err := utils.RandLetterString()
    if err != nil {
        return nil, err
    }
    cookie, err := utils.GenerateUserCookie()
    if err != nil {
        return nil, err
    }
    encodedMasterkey, b64salt, b64hash, err := utils.HashFromMasterkey(inputkey)
    hash, err := base64.StdEncoding.DecodeString(b64hash)
    if err != nil {
        return nil, err
    }
    _, _, b64filekey, err := utils.HashFromMasterkey(string(hash))
    if err != nil {
        return nil, err
    }
    user := &UserAuth{
        Uid: userId,
        Cookie: cookie,
        EncodedMasterkey: encodedMasterkey,
        Salt: b64salt,
        B64filekey: b64filekey,
    }
    return user, nil
}
func mockUserIdFromPlatform(platform PlatformUserAuth) (*UserAuth, error) {
    cookie, err := utils.GenerateUserCookie()
    if err != nil {
        return nil, err
    }
    user := &UserAuth{
        Uid: platform.UserId,
        Cookie: cookie,
    }
    return user, nil
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

    cookieStr, err := utils.MarshalCookie(user.Cookie)
    if err != nil {
        return err
    }
    _, err = db.Exec(q, user.Uid, cookieStr)
    return err
}
