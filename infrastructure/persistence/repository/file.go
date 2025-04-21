package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/Olyxz16/sherpa/domain/model"
	"github.com/Olyxz16/sherpa/infrastructure/persistence"
	db "github.com/Olyxz16/sherpa/infrastructure/persistence/sqlc"
)

type FileRepository struct {
	q	*db.Queries
}

func NewFileRepository() *FileRepository {
	return &FileRepository{
		q: db.New(persistence.Conn()),
	};
}

func (r *FileRepository) Persist(file *model.File, ctx context.Context) error {
	params := db.PersistFileParams {
		OwnerID: int32(file.GetOwner().GetID()),
		Source: string(file.GetSource()),
		Reponame: file.GetReponame(),
		Filename: file.GetFilename(),
		B64content: pgtype.Text{String: file.GetB64Content(), Valid: true},
		B64nonce: pgtype.Text{String: file.GetB64Nonce(), Valid: true},
	}
	err := r.q.PersistFile(ctx, params)
	return err
}

func (r *FileRepository) Find(owner *model.User, source model.AuthSource, reponame, filename string, ctx context.Context) (*model.File, error) {
	params := db.FindFileParams {
		OwnerID: int32(owner.GetID()),
		Source: string(source),
		Reponame: reponame,
		Filename: filename,
	}
	data, err := r.q.FindFile(ctx, params)
	if err != nil {
		return nil, err
	}

	file := model.NewFile(
		owner,
		source,
		reponame,
		filename,
		data.B64content.String,
		data.B64nonce.String,
	)
	return file, err
}

func (r *FileRepository) FindAll(owner *model.User, ctx context.Context) ([]*model.File, error) {
	data, err := r.q.FindAllFiles(ctx, int32(owner.GetID()))
	if err != nil {
		return []*model.File{}, err
	}
	
	length := len(data)
	files := make([]*model.File, length)
	for _, v := range(data) {
		file := model.NewFile(
			owner,
			model.AuthSource(v.Source),
			v.Reponame,
			v.Filename,
			v.B64content.String,
			v.B64nonce.String,
		)
		files = append(files, file)
	}

	return files, nil
}
