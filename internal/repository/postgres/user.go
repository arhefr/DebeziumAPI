package repository

import (
	"context"
	"debez/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (r *UserRepository) Select(ctx context.Context, offset, limit int) ([]models.User, error) {
	return []models.User{}, nil
}

func (r *UserRepository) SelectByID(ctx context.Context, userID *models.UserID) (*models.User, error) {
	return &models.User{}, nil
}

func (r *UserRepository) Insert(ctx context.Context, user *models.UserCU) (*models.User, error) {
	return &models.User{}, nil
}

func (r *UserRepository) Update(ctx context.Context, user *models.UserCU) error {
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, userID *models.UserID) error {
	return nil
}
