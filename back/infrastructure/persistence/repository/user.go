package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/Olyxz16/sherpa/domain/model"
	"github.com/Olyxz16/sherpa/infrastructure/persistence"
	db "github.com/Olyxz16/sherpa/infrastructure/persistence/sqlc"
)

type UserRepository struct {
	q	   *db.Queries
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		q:		db.New(persistence.Conn()),
	};
}

func (r *UserRepository) Persist(user *model.User, ctx context.Context) error {
	params := db.PersistUserParams {
		Uid: int32(user.GetID()),
		Username: user.GetUsername(),
		Masterkey: pgtype.Text{String: user.GetMasterkey(), Valid: true},
		B64salt: pgtype.Text{String: user.GetB64Salt(), Valid: true},
		B64filekey: pgtype.Text{String: user.GetB64Filekey(), Valid: true},
	}
	err := r.q.PersistUser(ctx, params);
	return err;
}

func (r *UserRepository) FindFromID(uid int, ctx context.Context) (*model.User, error) {
	data, err := r.q.FindUser(ctx, int32(uid));
	if err != nil {
		return nil, err;
	}
	user := model.NewUser(
		int(data.Uid), 
		data.Username,
		data.Masterkey.String, 
		data.B64salt.String, 
		data.B64filekey.String,
	)
	return user, nil;
}

func (r *UserRepository) FindFromPlatformID(auth *model.Auth, ctx context.Context) (*model.User, error) {
	uid := auth.GetUser().GetID();
	data, err := r.q.FindUser(ctx, int32(uid));
	if err != nil {
		return nil, err
	}
	user := model.NewUser(
		int(data.Uid),
		data.Username,
		data.Masterkey.String,
		data.B64salt.String,
		data.B64filekey.String,
	)
	return user, nil;
}

func (r *UserRepository) UpdateMasterKey(user *model.User, ctx context.Context) error {
	params := db.UpdateMasterkeyParams{
		Uid: int32(user.GetID()),
		Masterkey: pgtype.Text{String: user.GetMasterkey(), Valid: true},
		B64salt: pgtype.Text{String: user.GetB64Salt(), Valid: true},
		B64filekey: pgtype.Text{String: user.GetB64Filekey(), Valid: true},
	};
	error := r.q.UpdateMasterkey(ctx, params);
	return error;
}
