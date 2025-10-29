package service

import (
	"context"
	"debez/internal/models"

	"errors"
)

const (
	defaultLimit  = 100
	defaultOffset = 1
)

var (
	ErrMissingRequiredFields = errors.New("error JSON missing required field or fields")
	ErrBadFieldValue         = errors.New("error JSON bad value of field or fields")
)

type UserRepository interface {
	Select(ctx context.Context, offset, limit int) ([]*models.User, error)
	SelectByID(ctx context.Context, userID *models.UserID) (*models.User, error)
	Insert(ctx context.Context, userCU *models.UserCreate) (*models.User, error)
	Update(ctx context.Context, userCU *models.UserUpdate) error
	Delete(ctx context.Context, userID *models.UserID) error
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {

	return &UserService{
		repo: repo,
	}
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

	user, err := s.repo.Select(ctx, offset, limit)
	return user, err
}

func (s *UserService) SaveUser(ctx context.Context, userCU *models.UserCreate) (*models.User, error) {

	if userCU.Email == "" || userCU.Name == "" || len(userCU.Role) == 0 {
		return nil, ErrMissingRequiredFields
	}

	return s.repo.Insert(ctx, userCU)
}

func (s *UserService) DeleteUser(ctx context.Context, userID *models.UserID) error {

	return s.repo.Delete(ctx, userID)
}

func (s *UserService) UpdateUser(ctx context.Context, userCU *models.UserUpdate) error {

	if userCU.Email == "" && userCU.Name == "" && len(userCU.Role) == 0 && userCU.LastName == "" {
		return ErrMissingRequiredFields
	}

	return s.repo.Update(ctx, userCU)
}
