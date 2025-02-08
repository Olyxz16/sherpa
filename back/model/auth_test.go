package model

import (
	"math/rand"
	"testing"

    "github.com/Olyxz16/sherpa/utils"
)


func TestAuthenticateNewUser(t *testing.T) {
    New()
    inputPlatform := mockPlatformUser()

    
    if err := clean() ; err != nil {
        t.Fatalf("Failed cleaning database : %v", err)
    }

    actual, isNew, err := AuthenticateUser(inputPlatform)
    if err != nil {
        t.Fatalf("Error during authentication : %v", err)
    }

    if !isNew {
        t.Fatal(`Error : isNew wrong value
                    expected : true
                    actual : false`)
    }
    userId, exists, err := userIdFromPlatformId(inputPlatform)
    if err != nil {
        t.Fatal("Failed fetching userId")
    }
    if !exists {
        t.Fatal(`Error : exists wrong value
                    expected : true
                    actual : false`)
    }
    if userId != actual.Uid {
        t.Fatalf(`Error : user ids don't match
                    expected : %v,
                    actual : %v`, actual.Uid, userId)
    }
}

func TestAuthenticateExistingUser(t *testing.T) {
    New()
    inputPlatform := mockPlatformUser()
    inputUser, err := mockUserIdFromPlatform(inputPlatform)
    if err != nil {
        t.Fatal("Error creating mock Userid")
    }

    if err := clean() ; err != nil {
        t.Fatalf("Failed cleaning database : %v", err)
    }
    if err := insertUser(*inputUser) ; err != nil {
        t.Fatalf("Failed inserting userAuth : %v", err) 
    }
    if err := insertPlatform(inputPlatform) ; err != nil {
        t.Fatalf("Failed inserting platformAuth : %v", err) 
    }

    actual, isNew, err := AuthenticateUser(inputPlatform)
    if err != nil {
        t.Fatalf("Error during authentication : %v", err)
    }

    if isNew {
        t.Fatal(`Error : isNew wrong value
                    expected : false
                    actual : true`)
    }
    userId, exists, err := userIdFromPlatformId(inputPlatform)
    if err != nil {
        t.Fatalf("Failed fetching userId : %v", err)
    }
    if !exists {
        t.Fatal(`Error : exists wrong value
                    expected : true
                    actual : false`)
    }
    if inputPlatform.UserId != actual.Uid {
        t.Fatalf(`Error : mock user id doesn't match created user id
                    expected : %v,
                    actual: %v`, inputPlatform.UserId, actual.Uid)
    }
    if userId != actual.Uid {
        t.Fatalf(`Error : user ids don't match
                    expected : %v,
                    actual : %v`, actual.Uid, userId)
    }
}



func userIdFromPlatformId(platform PlatformUserAuth) (int, bool, error) {
    db := instance.db
    q := `SELECT userId FROM PlatformUserAuth
        WHERE platformId=$1`
    row := db.QueryRow(q, platform.PlatformId)
    var userId int
    err := row.Scan(&userId) 
    if err != nil {
        return -1, false, err
    }
    return userId, true, nil
}


func mockPlatformUser() PlatformUserAuth {
    acces_token, err := utils.RandLetterString()
    if err != nil {
        panic("Failed generating string")
    }
    refresh_token, err := utils.RandLetterString()
    if err != nil {
        panic("Failed generating string")
    }
    return PlatformUserAuth{
        UserId: int(rand.Int31()),
        PlatformId: int(rand.Int31()),
        Source: "github.com",
        Access_token: acces_token,
        Refresh_token: refresh_token,
        Expires_in: 10000,
        Refresh_expires_in: 10000,
    }
}
