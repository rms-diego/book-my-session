package filmsservice

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"

	s3gateway "github.com/rms-diego/book-my-session/internal/gateway/s3-gateway"
	"github.com/rms-diego/book-my-session/internal/model"
	filmsdto "github.com/rms-diego/book-my-session/internal/modules/films/dto"
	filmsrepository "github.com/rms-diego/book-my-session/internal/modules/films/repository"
	"github.com/rms-diego/book-my-session/pkg/exception"
)

type filmsService struct {
	repository filmsrepository.FilmsRepository
	s3gateway  s3gateway.S3GatewayInterface
}

type FilmsService interface {
	Create(ctx context.Context, payload filmsdto.CreateFilmRequest) error
	Update(ctx context.Context, id string, payload filmsdto.UpdateFilmRequest) error
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context) (*[]model.Film, error)
	GetById(ctx context.Context, id string) (*model.Film, error)
	UploadThumbnail(ctx context.Context, id string, file *multipart.FileHeader) error
}

func NewFilmsService(repository filmsrepository.FilmsRepository, s3gateway s3gateway.S3GatewayInterface) FilmsService {
	return &filmsService{repository, s3gateway}
}

func (s *filmsService) Create(ctx context.Context, payload filmsdto.CreateFilmRequest) error {
	if err := s.repository.Create(ctx, payload); err != nil {
		return err
	}

	return nil
}

func (s *filmsService) Update(ctx context.Context, id string, payload filmsdto.UpdateFilmRequest) error {
	f, err := s.repository.GetById(ctx, id)
	if err != nil {
		return err
	}

	if f == nil {
		return exception.NewException("film not found", http.StatusNotFound)
	}

	if err := s.repository.Update(ctx, id, payload); err != nil {
		return err
	}

	return nil
}

func (s *filmsService) Delete(ctx context.Context, id string) error {
	f, err := s.repository.GetById(ctx, id)
	if err != nil {
		return err
	}

	if f == nil {
		return exception.NewException("film not found", http.StatusNotFound)
	}

	if err := s.repository.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (s *filmsService) GetAll(ctx context.Context) (*[]model.Film, error) {
	films, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	if films == nil || len(*films) == 0 {
		return &[]model.Film{}, nil
	}

	return films, nil
}

func (s *filmsService) GetById(ctx context.Context, id string) (*model.Film, error) {
	film, err := s.repository.GetById(ctx, id)
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

func (s *filmsService) UploadThumbnail(ctx context.Context, id string, data *multipart.FileHeader) error {
	film, err := s.repository.GetById(ctx, id)
	if err != nil {
		return err
	}

	if film == nil {
		return exception.NewException("film not found", http.StatusNotFound)
	}

	f, err := data.Open()
	if err != nil {
		return err
	}

	defer f.Close()

	filename := fmt.Sprintf("movies-thumbnail/%s.%s", id, data.Filename)
	url, err := s.s3gateway.Upload(ctx, f, filename)
	if err != nil {
		return err
	}

	if err := s.repository.Update(ctx, id, filmsdto.UpdateFilmRequest{Thumbnail: url}); err != nil {
		return err
	}

	return nil
}
