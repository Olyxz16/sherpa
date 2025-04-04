package repository

import (
	"context"
	"testing"
	"time"

	"github.com/Olyxz16/sherpa/config"
	"github.com/Olyxz16/sherpa/domain/model"
	"github.com/Olyxz16/sherpa/infrastructure/persistence"
)

func TestCreateAndPersistAuth(t *testing.T) {
    config := config.NewDatabaseConfig()
    persistence.New(config) 
    ctx := context.Background()

    authId := 45132

    userRepo := NewUserRepository()
    authRepo := NewAuthRepository()

    user := model.CreateUser("test_user")
    auth := model.NewAuth(authId, user, model.Github, "ghu_adyhbsdqpscuvbubhezb", "ghr_fiuafnjkscoliu", time.Now().Add(2*time.Minute).Nanosecond(), time.Now().Add(5*24*time.Hour).Nanosecond())

    userRepo.Persist(user, ctx)
    authRepo.Persist(auth, ctx)

    expectedAuth, err := authRepo.Find(45132, ctx)
    if err != nil {
        t.Error("Error fetching auth", err)
    }
    expectedAuthId := expectedAuth.GetAuthID()
    if expectedAuthId != authId {
        t.Error("Auth ids don't match", expectedAuthId, authId)
    }
    
}
