package repository

import (
	"context"

	"github.com/Olyxz16/sherpa/domain/model"
)

type UserRepository interface {
	FindFromID(id int, ctx context.Context) (*model.User, error);	
	FindFromPlatformID(auth *model.Auth, ctx context.Context) (*model.User, error);
	Create(user *model.User, ctx context.Context) (*model.User, error);
	UpdateMasterKey(user *model.User, ctx context.Context) error;
}
