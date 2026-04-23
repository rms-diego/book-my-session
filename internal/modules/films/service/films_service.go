package filmsservice

import (
	filmsdto "github.com/rms-diego/book-my-session/internal/modules/films/dto"
	filmsrepository "github.com/rms-diego/book-my-session/internal/modules/films/repository"
)

type filmsService struct {
	repository filmsrepository.FilmsRepository
}

type FilmsService interface {
	CreateFilm(payload filmsdto.CreateFilmRequest) error
}

func NewFilmsService(repository filmsrepository.FilmsRepository) FilmsService {
	return &filmsService{repository}
}

func (s *filmsService) CreateFilm(payload filmsdto.CreateFilmRequest) error {
	if err := s.repository.Create(payload); err != nil {
		return err
	}

	return nil
}
