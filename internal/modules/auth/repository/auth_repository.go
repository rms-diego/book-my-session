package authrepository

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/rms-diego/book-my-session/internal/model"
	authdto "github.com/rms-diego/book-my-session/internal/modules/auth/dto"
)

type authRepository struct {
	db *goqu.Database
}

type AuthRepository interface {
	Create(ctx context.Context, data authdto.SignUpRequest) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}

func NewAuthRepository(db *goqu.Database) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) Create(ctx context.Context, data authdto.SignUpRequest) (*model.User, error) {
	var user model.User

	_, err := r.db.Insert(model.USERS_TABLE).
		Cols("name", "email", "password", "role").
		Vals(goqu.Vals{data.Name, data.Email, data.Password, data.Role}).
		Returning("*").
		Executor().
		ScanStructContext(ctx, &user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *authRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var data model.User

	found, err := r.db.From(model.USERS_TABLE).
		Where(goqu.Ex{"email": email}).
		ScanStructContext(ctx, &data)

	if err != nil {
		return nil, err
	}

	if !found {
		return nil, nil
	}

	return &data, nil
}
