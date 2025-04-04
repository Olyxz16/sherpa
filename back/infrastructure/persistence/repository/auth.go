package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/Olyxz16/sherpa/domain/model"
	"github.com/Olyxz16/sherpa/infrastructure/persistence"
	db "github.com/Olyxz16/sherpa/infrastructure/persistence/sqlc"
)

type AuthRepository struct {
	q	*db.Queries
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{
		q: db.New(persistence.Conn()),
	};
}

func (r *AuthRepository) Persist(auth *model.Auth, ctx context.Context) error {
	params := db.PersistAuthParams {
		ID: int32(auth.GetAuthID()),		
        UserID: pgtype.Int4{Int32: int32(auth.GetUser().GetID()), Valid: true},
		Source: pgtype.Text{String: string(auth.GetPlatormSource()), Valid: true},
		AccessToken: pgtype.Text{String: auth.GetAccessToken(), Valid: true},
		RefreshToken: pgtype.Text{String: auth.GetRefreshToken(), Valid: true},
		ExpiresIn: pgtype.Float8{Float64: float64(auth.GetExpiresIn()), Valid: true},
		RtExpiresIn: pgtype.Float8{Float64: float64(auth.GetRefreshTokenExpireIn()), Valid: true},
	}
	err := r.q.PersistAuth(ctx, params)
	return err
}

func (r *AuthRepository) Find(id int, ctx context.Context) (*model.Auth, error) {
	data, err := r.q.FindAuthById(ctx, int32(id))
	if err != nil {
		return nil, err
	}
	user := model.NewUser(	
		int(data.ID_2),
		data.Username,
		data.Masterkey.String,
		data.B64salt.String,
		data.B64filekey.String,
	)
	auth := model.NewAuth(
		int(data.ID),
		user,
		model.AuthSource(data.Source.String),
		data.AccessToken.String,
		data.RefreshToken.String,
		int(data.ExpiresIn.Float64),
		int(data.RtExpiresIn.Float64),
	)
	return auth, nil;
}

func (r *AuthRepository) FindByUser(user *model.User, source model.AuthSource, ctx context.Context) (*model.Auth, error) {
	params := db.FindAuthByUserIdParams {
		UserID: pgtype.Int4{Int32: int32(user.GetID()), Valid: true},	
		Source: pgtype.Text{String: string(source), Valid: true},
	}
	data, err := r.q.FindAuthByUserId(ctx, params)
	if err != nil {
		return nil, err
	}
	auth := model.NewAuth(
		int(data.ID),
		user,
		model.AuthSource(data.Source.String),
		data.AccessToken.String,
		data.RefreshToken.String,
		int(data.ExpiresIn.Float64),
		int(data.RtExpiresIn.Float64),
	)
	return auth, nil;
}

