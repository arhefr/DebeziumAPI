package v1

import (
	"context"
	"debez/internal/models"
	"debez/pkg/contract"
)

type Handler struct {
	us UserService
	contract.UnimplementedUserServiceServer
}

type UserService interface {
	GetUsers(ctx context.Context, offset, limit int) ([]models.User, error)
	GetUser(ctx context.Context, userID *models.UserID) (*models.User, error)
	SaveUser(ctx context.Context, userCU *models.UserCU) (*models.User, error)
	DeleteUser(ctx context.Context, userID *models.UserID) error
	UpdateUser(ctx context.Context, userCU *models.UserCU) error
}

func NewHandler(us UserService) *Handler {
	return &Handler{us: us}
}
