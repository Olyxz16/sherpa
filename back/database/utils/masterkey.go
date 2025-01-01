package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type params struct {
    iterations  uint32
    memory      uint32
    thread      uint8
    saltLength  uint32
    keyLength   uint32
}

var (
    ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
    ErrIncompatibleVersion = errors.New("incompatible version of argon2")
    p = params {
        iterations: 3,
        memory: 64 * 1024,
        thread: 4,
        saltLength: 16,
        keyLength: 32,
    }
)

func HashFromMasterkey(masterkey string) (encodedHash, b64Salt, b64Hash string, err error) {
    salt, err := generateSalt()
    if err != nil {
        return "", "", "", err
    }
    hash := argon2.IDKey([]byte(masterkey), salt, p.iterations, p.memory, p.thread, p.keyLength)

    b64Salt = base64.RawStdEncoding.EncodeToString(salt)
    b64Hash = base64.RawStdEncoding.EncodeToString(hash)

    encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.memory, p.iterations, p.thread, b64Salt, b64Hash)
    
    return encodedHash, b64Salt, b64Hash, nil
}

func CompareHashAndPassword(masterkey, encodedHash string) (match bool, err error) {
    params, salt, hash, err := decodeHash(encodedHash)
    if err != nil {
        return false, err
    }

    otherHash := argon2.IDKey([]byte(masterkey), salt, params.iterations, params.memory, params.thread, params.keyLength)

    if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
        return true, nil
    }
    return false, nil
}


func decodeHash(encodedHash string) (par *params, salt, hash []byte, err error) {
    vals := strings.Split(encodedHash, "$")
    if len(vals) != 6 {
        return nil, nil, nil, ErrInvalidHash
    }

    var version int
    _, err = fmt.Sscanf(vals[2], "v=%d", &version)
    if err != nil {
        return nil, nil, nil, err
    }
    if version != argon2.Version {
        return nil, nil, nil, ErrIncompatibleVersion
    }

    par = &params{}
    _, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &par.memory, &par.iterations, &par.thread)
    if err != nil {
        return nil, nil, nil, err
    }

    salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
    if err != nil {
        return nil, nil, nil, err
    }
    par.saltLength = uint32(len(salt))

    hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
    if err != nil {
        return nil, nil, nil, err
    }
    par.keyLength = uint32(len(hash))

    return par, salt, hash, nil
}


func generateSalt() ([]byte, error) {
	b := make([]byte, p.saltLength)
	_, err := rand.Read(b)
    if err != nil {
        return nil, err
    }
    return b, nil
}
