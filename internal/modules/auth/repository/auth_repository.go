package authrepository

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/rms-diego/book-my-session/internal/model"
	authdto "github.com/rms-diego/book-my-session/internal/modules/auth/dto"
)

type authRepository struct {
	db *goqu.Database
}

type AuthRepository interface {
	Create(data authdto.SignUpRequest) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
}

func NewAuthRepository(db *goqu.Database) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) Create(data authdto.SignUpRequest) (*model.User, error) {
	var user model.User

	_, err := r.db.Insert(model.USERS_TABLE).
		Cols("name", "email", "password", "role").
		Vals(goqu.Vals{data.Name, data.Email, data.Password, data.Role}).
		Returning("*").
		Executor().
		ScanStruct(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *authRepository) GetByEmail(email string) (*model.User, error) {
	var data model.User

	found, err := r.db.From(model.USERS_TABLE).
		Where(goqu.Ex{"email": email}).
		ScanStruct(&data)

	if err != nil {
		return nil, err
	}

	if !found {
		return nil, nil
	}

	return &data, nil
}
