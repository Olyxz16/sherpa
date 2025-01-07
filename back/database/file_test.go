package database

import (
	"encoding/base64"
	"testing"

	"github.com/Olyxz16/go-vue-template/database/utils"
)

func TestFetchUserFileData(t *testing.T) {
    New()
    user, err := mockUserAuth()
    if err != nil {
        t.Fatalf("Error mocking userauth : %v", err)
    }
    expectedContent, fileData, err := mockFileDataFromUser(*user)
    if err != nil {
        t.Fatalf("Error mocking file data : %v", err)
    }
    
    err = clean()
    if err != nil {
        t.Fatalf("Error cleaning database : %v", err)
    }
    err = insertUser(*user)
    if err != nil {
        t.Fatalf("Error inserting user : %v", err)
    }
    err = insertFileData(*fileData)
    if err != nil {
        t.Fatalf("Error inserting fileData : %v", err)
    }
    
    actualContent, err := FetchFileContent(user.Cookie, fileData.Source, fileData.RepoName, fileData.FileName)
    if err != nil {
        t.Fatalf("Error in FetchFile : %v", err)
    }

    if expectedContent != actualContent {
        t.Fatalf(`Error: file content should be equal
                    expected : %s
                    actual : %s`, expectedContent, actualContent)
    }
    
}


func mockFileDataFromUser(userAuth UserAuth) (string, *FileData, error) {
    content, err := utils.RandLetterString()
    if err != nil {
        return "", nil, err
    }
    
    filekey, err := base64.StdEncoding.DecodeString(userAuth.B64filekey)
    b64content, b64nonce, err := utils.EncryptFile(filekey, content)
    if err != nil {
        return "", nil, err
    }
    fileData := &FileData{
        OwnerId: userAuth.Uid,
        Source: "github.com",
        RepoName: "TestRepository",
        FileName: ".env",
        B64Content: b64content,
        B64Nonce: b64nonce,
    }
    return "", fileData, nil
}

func insertFileData(file FileData) (error) {
    db := dbInstance.db
    q := `INSERT INTO FileData
        (ownerId, source, reponame, filename, b64content, b64nonce)
        VALUES ($1, $2, $3, $4, $5, $6)`

    _, err := db.Exec(q, file.OwnerId, file.Source, file.RepoName, file.FileName, file.B64Content, file.B64Nonce)
    return err
}
