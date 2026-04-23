package filmsrepository

import (
	"reflect"

	"github.com/doug-martin/goqu/v9"
	"github.com/rms-diego/book-my-session/internal/model"
	filmsdto "github.com/rms-diego/book-my-session/internal/modules/films/dto"
)

type filmsRepository struct {
	db *goqu.Database
}

type FilmsRepository interface {
	Create(payload filmsdto.CreateFilmRequest) error
}

func NewFilmsRepository(db *goqu.Database) FilmsRepository {
	return &filmsRepository{db}
}

func (r *filmsRepository) Create(payload filmsdto.CreateFilmRequest) error {
	var film model.Film
	t := reflect.TypeOf(payload)
	v := reflect.ValueOf(payload)

	cols := []any{}
	vals := goqu.Vals{}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fv := v.Field(i)

		if fv.IsZero() {
			continue
		}

		col := f.Tag.Get("db")
		if col == "" {
			continue
		}

		cols = append(cols, col)
		vals = append(vals, fv.Interface())
	}

	_, err := r.db.Insert(model.FILMS_TABLE).
		Cols(cols...).
		Vals(vals).
		Returning("*").
		Executor().
		ScanStruct(&film)

	if err != nil {
		return err
	}

	return nil
}
