package repository

import (
	"context"
	"testing"
	"time"

	"github.com/Olyxz16/sherpa/config"
	"github.com/Olyxz16/sherpa/domain/model"
	"github.com/Olyxz16/sherpa/infrastructure/persistence"
)

func TestCreateAndPersistAuthAndFindById(t *testing.T) {
    config := config.NewDatabaseConfig()
    persistence.New(config) 
    ctx := context.Background()

    authId := 45132

    userRepo := NewUserRepository()
    authRepo := NewAuthRepository()

    user := model.CreateUser("test_user")
    auth := model.NewAuth(authId, user, model.Github, "ghu_adyhbsdqpscuvbubhezb", "ghr_fiuafnjkscoliu", time.Now().Add(2*time.Minute).Nanosecond(), time.Now().Add(5*24*time.Hour).Nanosecond())

    err := userRepo.Persist(user, ctx)
    if err != nil {
        t.Error("Error persisting user", err)
    }
    err = authRepo.Persist(auth, ctx)
    if err != nil {
        t.Error("Error persisting auth", err)
    }

    expectedAuth, err := authRepo.Find(authId, ctx)
    if err != nil {
        t.Error("Error fetching auth", err)
    }
    expectedAuthId := expectedAuth.GetAuthID()
    if expectedAuthId != authId {
        t.Error("Auth ids don't match", expectedAuthId, authId)
    }
}

func TestCreateAndPersistAuthAndFindByUser(t *testing.T) {
    config := config.NewDatabaseConfig()
    persistence.New(config) 
    ctx := context.Background()

    authId := 78962
    source := model.Github

    userRepo := NewUserRepository()
    authRepo := NewAuthRepository()

    user := model.CreateUser("test_user2")
    auth := model.NewAuth(authId, user, source, "ghu_adyhjkli_oauahezb", "ghr_dahuzizizu,", time.Now().Add(2*time.Minute).Nanosecond(), time.Now().Add(5*24*time.Hour).Nanosecond())

    err := userRepo.Persist(user, ctx)
    if err != nil {
        t.Error("Error persisting user", err)
    }
    err = authRepo.Persist(auth, ctx)
    if err != nil {
        t.Error("Error persisting auth", err)
    }

    expectedAuth, err := authRepo.FindByUser(user, source, ctx)
    if err != nil {
        t.Error("Error fetching auth", err)
    }
    expectedAuthId := expectedAuth.GetAuthID()
    if expectedAuthId != authId {
        t.Error("Auth ids don't match", expectedAuthId, authId)
    }
}

func TestCreateUserAndAuthMatches(t *testing.T) {
    config := config.NewDatabaseConfig()
    persistence.New(config) 
    ctx := context.Background()

    authId := 98765123
    source := model.Github

    userRepo := NewUserRepository()
    authRepo := NewAuthRepository()

    user := model.CreateUser("test_user3")
    auth := model.NewAuth(authId, user, source, "ghu_azejokpogejhn", "ghr_tjaowpz_mfjqs", time.Now().Add(2*time.Minute).Nanosecond(), time.Now().Add(5*24*time.Hour).Nanosecond())

    err := userRepo.Persist(user, ctx)
    if err != nil {
        t.Error("Error persisting user", err)
    }
    err = authRepo.Persist(auth, ctx)
    if err != nil {
        t.Error("Error persisting auth", err)
    }

    expectedAuth, err := authRepo.Find(authId, ctx)
    if err != nil {
        t.Error("Error fetching auth", err)
    }
    expectedAuthId := expectedAuth.GetAuthID()
    if expectedAuthId != authId {
        t.Error("Auth ids don't match", expectedAuthId, authId)
    }
    doUsersMatch := user.GetID() == expectedAuth.GetUser().GetID()
    if !doUsersMatch {
        t.Error("Users don't match", user, expectedAuth.GetUser())
    }
}

func TestCreateAuthWithConcurrentUser(t *testing.T) {
    config := config.NewDatabaseConfig()
    persistence.New(config) 
    ctx := context.Background()

    authId := 98765123
    source := model.Github

    userRepo := NewUserRepository()
    authRepo := NewAuthRepository()

    user1 := model.CreateUser("test_user_64532")
    user2 := model.CreateUser("test_user_463287")
    auth1 := model.NewAuth(authId, user1, source, "ghu_azejokpogejhn", "ghr_tjaowpz_mfjqs", time.Now().Add(2*time.Minute).Nanosecond(), time.Now().Add(5*24*time.Hour).Nanosecond())
    auth2 := model.NewAuth(authId, user2, source, "ghu_azejokpogejhn", "ghr_tjaowpz_mfjqs", time.Now().Add(2*time.Minute).Nanosecond(), time.Now().Add(5*24*time.Hour).Nanosecond())


    err := userRepo.Persist(user1, ctx)
    if err != nil {
        t.Error("Error persisting user", err)
    }
    err = userRepo.Persist(user2, ctx)
    if err != nil {
        t.Error("Error persisting user", err)
    }
    
    err = authRepo.Persist(auth1, ctx)
    if err != nil {
        t.Error("Error persisting auth", err)
    }
    err = authRepo.Persist(auth2, ctx)
    if err == nil {
        t.Error("Should return an error")
    }

}
