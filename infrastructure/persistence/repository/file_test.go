package repository

import (
	"context"
	"testing"

	"github.com/Olyxz16/sherpa/config"
	"github.com/Olyxz16/sherpa/domain/model"
	"github.com/Olyxz16/sherpa/infrastructure/persistence"
)

func TestCreateAndPersistNewAccountFile(t *testing.T) {
    config := config.NewDatabaseConfig()
    persistence.New(config) 
    ctx := context.Background()

    source := model.Github
    reponame := "test_repo_name_4651"
    filename := "test_file_name_9643"
    content := "test_content_5212"

    userRepo := NewUserRepository()
    fileRepo := NewFileRepository()

    user := model.CreateUser("test_user_3152")
    err := user.SetUserMasterkey("hdaiajksdqsdq")
    if err != nil {
        t.Error("Error setting user masterkey", err)
    }
    file := model.CreateFile(user, source, reponame, filename)
    err = file.Encrypt(content)
    if err != nil {
        t.Error("Error encrypting file", err)
    }

    err = userRepo.Persist(user, ctx)
    if err != nil {
        t.Error("Error persisting user", err)
    }
    err = fileRepo.Persist(file, ctx)
    if err != nil {
        t.Error("Error persisting file", err)
    }

    actualFile, err := fileRepo.Find(user, source, reponame, filename, ctx)
    if err != nil {
        t.Error("Error fetching file", err)
    }
    actualContent, err := actualFile.Decrypt()
    if err != nil {
        t.Error("Error decrypting file", err)
    }
    expectedContent, err := file.Decrypt()
    if err != nil {
        t.Error("Error decrypting file", err)
    }
    if actualContent != expectedContent {
        t.Error("File contents don't match")
    }
}
