package database

import (
    "testing"
)

func TestFetchUserUnencryptedFileData(t *testing.T) {
    New()
    user, err := mockUserAuth()
    if err != nil {
        t.Fatalf("Error mocking userauth : %v", err)
    }
    fileData := mockFileDataFromUserId(user.Uid)
    
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
    
    content, err := FetchFile(user.Cookie, fileData.repoName, fileData.fileName)  
    if err != nil {
        t.Fatalf("Error in FetchFile : %v", err)
    }

    if content != fileData.content {
        t.Fatalf(`Error: file content should be equal
                    expected : %s
                    actual : %s`, fileData.content, content)
    }
    
}


func mockFileDataFromUserId(userId int) (*FileData) {
    return &FileData{
        ownerId: userId,
        source: "github.com",
        repoName: "TestRepository",
        fileName: ".env",
        content: "DEBUG=true\nPORT=80",
    }
}

func insertFileData(file FileData) (error) {
    db := dbInstance.db
    q := `INSERT INTO FileData
        (ownerId, source, reponame, filename, content)
        VALUES ($1, $2, $3, $4, $5)`

    _, err := db.Exec(q, file.ownerId, file.source, file.repoName, file.fileName, file.content)
    return err
}
