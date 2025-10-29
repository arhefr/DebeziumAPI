package v1

import (
	"context"
	"debez/internal/models"
)

type Handler struct {
	service UserService
}

func NewHandler(service UserService) *Handler {
	return &Handler{service: service}
}

type UserService interface {
	GetUsers(ctx context.Context, offset, limit int) ([]*models.User, error)
	GetUser(ctx context.Context, userID *models.UserID) (*models.User, error)
	SaveUser(ctx context.Context, userCU *models.UserCreate) (*models.User, error)
	DeleteUser(ctx context.Context, userID *models.UserID) error
	UpdateUser(ctx context.Context, userCU *models.UserUpdate) error
}
