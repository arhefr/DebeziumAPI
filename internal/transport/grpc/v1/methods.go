package v1

import (
	"context"
	pb "debez/pkg/contract/proto"
)

func (h *Handler) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.User, error) {
	return nil, nil
}

func (h *Handler) GetUser(ctx context.Context, in *pb.UserID) (*pb.User, error) {
	return nil, nil
}

func (h *Handler) GetUsers(ctx context.Context, in *pb.GetUsersRequest) (*pb.Users, error) {
	return nil, nil
}

func (h *Handler) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.Void, error) {
	return nil, nil
}

func (h *Handler) DeleteUser(ctx context.Context, in *pb.UserID) (*pb.Void, error) {
	return nil, nil
}
