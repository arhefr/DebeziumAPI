package v1

import (
	"context"
	"debez/pkg/contract"
)

func (h *Handler) CreateUser(ctx context.Context, in *contract.CreateUserRequest) (*contract.User, error) {
	return nil, nil
}

func (h *Handler) GetUser(ctx context.Context, in *contract.UserID) (*contract.User, error) {
	return nil, nil
}

func (h *Handler) GetUsers(ctx context.Context, in *contract.GetUsersRequest) (*contract.Users, error) {
	return nil, nil
}

func (h *Handler) UpdateUser(ctx context.Context, in *contract.UpdateUserRequest) (*contract.Void, error) {
	return nil, nil
}

func (h *Handler) DeleteUser(ctx context.Context, in *contract.UserID) (*contract.Void, error) {
	return nil, nil
}
