package model

import (
	"encoding/base64"

	"github.com/Olyxz16/sherpa/crypto"
)

type File struct {
    owner	    *User
    source      AuthSource
    reponame    string
    filename    string
    b64content  string
    b64nonce    string
}

func NewFile(owner *User, source AuthSource, reponame, filename, b64content, b64nonce string) *File {
	return &File {
		owner: owner,
		source: source,
		reponame: reponame,
		filename: filename,
		b64content: b64content,
		b64nonce: b64nonce,
	}
}

func (f *File) GetOwner() *User {
	return f.owner
}

func (f *File) GetSource() AuthSource {
	return f.source
}

func (f *File) GetReponame() string {
	return f.reponame
}

func (f *File) GetFilename() string {
	return f.filename
}

func (f *File) GetB64Content() string {
	return f.b64content
}

func (f *File) GetB64Nonce() string {
	return f.b64nonce
}

func (f *File) Encrypt(content string) error {
    filekey, err := base64.StdEncoding.DecodeString(f.owner.GetB64Filekey())
    if err != nil {
        return err
    }
    
	b64content, b64nonce, err := crypto.EncryptFile(filekey, content)
    if err != nil {
        return err
    }

	f.b64content = b64content
	f.b64nonce = b64nonce
	return nil
}

func (f *File) Decrypt() (string, error) {
    filekey, err := base64.StdEncoding.DecodeString(f.owner.GetB64Filekey())
    if err != nil {
        return "", err
    }

    content, err := crypto.DecryptFile(filekey, f.b64nonce, f.b64content)
    if err != nil {
        return "", err
    }
    
	return content, nil
}
