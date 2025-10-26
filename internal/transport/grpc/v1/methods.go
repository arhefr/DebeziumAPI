package v1

import (
	"context"
	"debez/internal/models"
	"debez/pkg/contract"
)

func (h *Handler) CreateUser(ctx context.Context, in *contract.CreateUserRequest) (*contract.User, error) {

	user, err := h.us.SaveUser(ctx, &models.UserCU{Name: in.Name, LastName: in.LastName, Email: in.Email, Role: in.Role})
	if err != nil {
		return nil, err
	}

	return &contract.User{Id: user.ID, Name: user.Name, LastName: user.LastName, Email: user.Email, Role: user.Role}, nil
}
