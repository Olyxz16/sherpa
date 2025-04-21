package repository

import (
	"context"
	"reflect"
	"testing"

	"github.com/Olyxz16/sherpa/config"
	"github.com/Olyxz16/sherpa/domain/model"
	"github.com/Olyxz16/sherpa/infrastructure/persistence"
)

func TestCreateAndPersistNewUser(t *testing.T) {
    config := config.NewDatabaseConfig()
    persistence.New(config) 
    ctx := context.Background()
    
    userRepo := NewUserRepository()
    user := model.CreateUser("test_user_96432")
    
    err := userRepo.Persist(user, ctx)
    if err != nil {
        t.Error("Error persisting user", err)
    }

    expectedUser, err := userRepo.FindFromID(user.GetID(), ctx)
    if err != nil {
        t.Error("Error finding user", err)
    }

    if user.GetUsername() != expectedUser.GetUsername() {
        t.Error("Users don't match", user, expectedUser)
    }
}

func TestCreateAndPersistExistingUser(t *testing.T) {
    config := config.NewDatabaseConfig()
    persistence.New(config) 
    ctx := context.Background()
    
    userRepo := NewUserRepository()
    user := model.CreateUser("test_user_413354")
    user.SetUserMasterkey("eazezee")
    
    err := userRepo.Persist(user, ctx)
    if err != nil {
        t.Error("Error persisting user", err)
    }

    expectedUser, err := userRepo.FindFromID(user.GetID(), ctx)
    if err != nil {
        t.Error("Error finding user", err)
    }

    doUsersMatch := reflect.DeepEqual(user, expectedUser)
    if !doUsersMatch {
        t.Error("Users don't match", user, expectedUser)
    }
}
