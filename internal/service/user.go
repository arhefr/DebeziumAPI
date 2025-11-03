package service

import (
	"context"
	"debez/internal/models"

	"errors"
)

const (
	defaultLimit  = 10
	defaultOffset = 1
)

var (
	ErrMissingRequiredFields = errors.New("error JSON missing required field or fields")
	ErrBadFieldValue         = errors.New("error JSON bad value of field or fields")
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {

	return &UserService{
		repo: repo,
	}
}

type UserRepository interface {
	Select(ctx context.Context, offset, limit int) ([]*models.User, error)
	SelectByID(ctx context.Context, userID *models.UserID) (*models.User, error)
	Insert(ctx context.Context, userC *models.UserCreate) (*models.User, error)
	Update(ctx context.Context, userU *models.UserUpdate) error
	Delete(ctx context.Context, userID *models.UserID) error
}

func (s *UserService) GetUser(ctx context.Context, userID *models.UserID) (*models.User, error) {

	users, err := s.repo.SelectByID(ctx, userID)
	return users, err
}

func (s *UserService) GetUsers(ctx context.Context, offset, limit int) ([]*models.User, error) {

	if offset == 0 {
		offset = defaultOffset
	}

	if limit == 0 {
		limit = defaultLimit
	}

	users, err := s.repo.Select(ctx, offset, limit)
	return users, err
}

func (s *UserService) SaveUser(ctx context.Context, userC *models.UserCreate) (*models.User, error) {

	if userC.Email == "" || userC.Name == "" || userC.LastName == "" {
		return nil, ErrMissingRequiredFields
	}

	return s.repo.Insert(ctx, userC)
}

func (s *UserService) DeleteUser(ctx context.Context, userID *models.UserID) error {

	return s.repo.Delete(ctx, userID)
}

func (s *UserService) UpdateUser(ctx context.Context, userU *models.UserUpdate) error {

	return s.repo.Update(ctx, userU)
}
