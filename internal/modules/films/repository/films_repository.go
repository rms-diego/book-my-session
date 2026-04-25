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
	Update(id string, payload filmsdto.UpdateFilmRequest) error
	Delete(id string) error
	GetAll() (*[]model.Film, error)
	GetById(id string) (*model.Film, error)
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

func (r *filmsRepository) Update(id string, payload filmsdto.UpdateFilmRequest) error {
	updateData := make(map[string]any)
	t := reflect.TypeOf(payload)
	v := reflect.ValueOf(payload)

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

		if col == "deleted" && fv.Elem().Bool() == false {
			updateData["deleted_at"] = nil
		}

		if col == "deleted" && fv.Elem().Bool() {
			updateData["deleted_at"] = goqu.L("NOW()")
		}

		updateData[col] = fv.Interface()
	}

	if len(updateData) == 0 {
		return nil
	}

	updateData["updated_at"] = goqu.L("NOW()")

	_, err := r.db.Update(model.FILMS_TABLE).
		Set(updateData).
		Where(goqu.Ex{"id": id}).
		Executor().
		Exec()

	return err
}

func (r *filmsRepository) Delete(id string) error {
	_, err := r.db.Update(model.FILMS_TABLE).
		Set(goqu.Record{"deleted": true, "deleted_at": goqu.L("NOW()")}).
		Where(goqu.Ex{"id": id}).
		Executor().
		Exec()

	return err
}

func (r *filmsRepository) GetById(id string) (*model.Film, error) {
	var film model.Film

	found, err := r.db.From(model.FILMS_TABLE).
		Where(goqu.Ex{"id": id}).
		Select("*").
		ScanStruct(&film)

	if err != nil {
		return nil, err
	}

	if !found {
		return nil, nil
	}

	return &film, nil
}

func (r *filmsRepository) GetAll() (*[]model.Film, error) {
	var films []model.Film

	err := r.db.From(model.FILMS_TABLE).
		Where(goqu.Ex{"deleted": false, "deleted_at": nil}).
		Select("*").
		ScanStructs(&films)

	if err != nil {
		return nil, err
	}

	return &films, nil
}
