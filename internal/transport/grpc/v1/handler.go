package v1

import (
	"context"
	"debez/internal/models"
	pb "debez/pkg/contract/proto"
)

type Handler struct {
	us UserService
	pb.UnimplementedUserServiceServer
}

type UserService interface {
	GetUsers(ctx context.Context, offset, limit int) ([]*models.User, error)
	GetUser(ctx context.Context, userID *models.UserID) (*models.User, error)
	SaveUser(ctx context.Context, userCU *models.UserCreate) (*models.User, error)
	DeleteUser(ctx context.Context, userID *models.UserID) error
	UpdateUser(ctx context.Context, userCU *models.UserUpdate) error
}

func NewHandler(us UserService) *Handler {
	return &Handler{us: us}
}
