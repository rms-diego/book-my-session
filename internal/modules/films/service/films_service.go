package filmsservice

import (
	"net/http"

	"github.com/rms-diego/book-my-session/internal/model"
	filmsdto "github.com/rms-diego/book-my-session/internal/modules/films/dto"
	filmsrepository "github.com/rms-diego/book-my-session/internal/modules/films/repository"
	"github.com/rms-diego/book-my-session/pkg/exception"
)

type filmsService struct {
	repository filmsrepository.FilmsRepository
}

type FilmsService interface {
	Create(payload filmsdto.CreateFilmRequest) error
	Update(id string, payload filmsdto.UpdateFilmRequest) error
	Delete(id string) error
	GetAll() (*[]model.Film, error)
	GetById(id string) (*model.Film, error)
}

func NewFilmsService(repository filmsrepository.FilmsRepository) FilmsService {
	return &filmsService{repository}
}

func (s *filmsService) Create(payload filmsdto.CreateFilmRequest) error {
	if err := s.repository.Create(payload); err != nil {
		return err
	}

	return nil
}

func (s *filmsService) Update(id string, payload filmsdto.UpdateFilmRequest) error {
	f, err := s.repository.GetById(id)
	if err != nil {
		return err
	}

	if f == nil {
		return exception.NewException("film not found", http.StatusNotFound)
	}

	if err := s.repository.Update(id, payload); err != nil {
		return err
	}

	return nil
}

func (s *filmsService) Delete(id string) error {
	f, err := s.repository.GetById(id)
	if err != nil {
		return err
	}

	if f == nil {
		return exception.NewException("film not found", http.StatusNotFound)
	}

	if err := s.repository.Delete(id); err != nil {
		return err
	}

	return nil
}

func (s *filmsService) GetAll() (*[]model.Film, error) {
	films, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	if films == nil || len(*films) == 0 {
		return &[]model.Film{}, nil
	}

	return films, nil
}

func (s *filmsService) GetById(id string) (*model.Film, error) {
	film, err := s.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	if film == nil {
		return nil, exception.NewException("film not found", http.StatusNotFound)
	}

	if film.Deleted && film.DeletedAt != nil {
		return nil, exception.NewException("film not found", http.StatusNotFound)
	}

	return film, nil
}
