package model

import (
	"crypto/rand"
	"encoding/base64"
	"math"
	"math/big"

	"go.uber.org/zap"

	"github.com/Olyxz16/sherpa/crypto"
)

type User struct {
	uid				int
	username		string
	masterkey		string	
	b64salt			string
	b64filekey		string
}

func CreateUser(username string) *User {
	biguid, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt32))
	uid := int(biguid.Int64())
	return &User{
		uid: uid,
		username: username,
	}	
}

func NewUser(uid int, username, masterkey, b64salt, b64filekey string) *User {
	return &User {
		uid: uid,
		username: username,
		masterkey: masterkey,
		b64salt: b64salt,
		b64filekey: b64filekey,
	}
}

func (u *User) GetID() int {
	return u.uid;
}

func (u *User) GetUsername() string {
	return u.username
}

func (u *User) GetMasterkey() string {
	return u.masterkey
}

func (u *User) GetB64Salt() string {
	return u.b64salt
}

func (u *User) GetB64Filekey() string {
	return u.b64filekey
}

func (u *User) IsNew() bool {
	return u.masterkey == ""
}

func (u *User) SetUserMasterkey(plaintextMasterkey string) (error) {
    masterkey, b64salt, b64hash, err := crypto.HashFromMasterkey(plaintextMasterkey)
    if err != nil {
        zap.L().Error("SetUserMasterkey", zap.Error(err))
        return err
    }
    hash, err := base64.StdEncoding.DecodeString(b64hash)
    if err != nil {
        return err
    }
    _, _, b64filekey, err := crypto.HashFromMasterkey(string(hash))
    if err != nil {
        return err
    }
	
	u.masterkey = masterkey
	u.b64salt = b64salt
	u.b64filekey = b64filekey
    return nil
}
