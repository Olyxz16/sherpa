package database

import (
	"reflect"
	"regexp"
	"testing"
)


func TestGenerateCookie(t *testing.T) {
    n := 100
    reg, err := regexp.Compile("[a-zA-Z]*")
    if err != nil {
        t.Log("Failed compiling regexp")
        t.Fail()
    }
    for i := 0 ; i < n ; i++ {
        cookie, err := generateUserCookie()
        if err != nil {
            t.Log("Failed generating cookie")
            t.Fail()
            continue
        }
        val := cookie.Value
        l := len(val)
        if l != cookie_size {
            t.Log("Failed: different size")
            t.Fail()
            continue
        }
        matches := reg.MatchString(val) 
        if !matches {
            t.Log("Failed: values don't match")
            t.Fail()
            continue
        }
    }
}


func TestMarshalCookie(t *testing.T) {
    n := 100
    for i := 0 ; i < n ; i++ {
        cookie, err := generateUserCookie() 
        if err != nil {
            t.Log("Failed generating cookie")
            t.Fail()
            continue
        }
        cookieStr, err := marshalCookie(cookie) 
        if err != nil {
            t.Logf("Failed marshalling cookie : %v\n", err)
            t.Fail()
            continue
        }
        cookieCopy, err := unmarshalCookie(cookieStr)
        if err != nil {
            t.Logf("Failed unmarshalling cookie : %v\n", err)
            t.Fail()
            continue
        }
        if !reflect.DeepEqual(cookie, cookieCopy) {
            t.Log("Failed: Cookies not equal")
            t.Fail()
            continue
        }
    }
}
