package repository

import (
	"context"

	"github.com/Olyxz16/sherpa/domain/model"
)

type FileRepository interface {
	Find(owner *model.User, source model.AuthSource, reponame, filename string, ctx context.Context) (*model.File, error);		
	FindAll(owner *model.User, ctx context.Context) ([]*model.File, error);
	Persist(file *model.File, ctx context.Context) error;
}
