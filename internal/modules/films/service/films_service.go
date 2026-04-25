package filmsservice

import (
	"net/http"

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
	f, err := s.repository.FindById(id)
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
	f, err := s.repository.FindById(id)
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
