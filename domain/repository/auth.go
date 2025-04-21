package repository

import (
	"context"
	
	"github.com/Olyxz16/sherpa/domain/model"
)


type AuthRepository interface {
	Persist(auth *model.Auth, ctx context.Context) error;
	Find(id int, ctx context.Context) (*model.Auth, error);
	FindByUser(user *model.User, source model.AuthSource, ctx context.Context) (*model.Auth, error);
}

