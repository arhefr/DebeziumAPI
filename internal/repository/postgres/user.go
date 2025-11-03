package repository

import (
	"context"
	"debez/internal/models"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *UserRepository) Select(ctx context.Context, offset, limit int) ([]*models.User, error) {

	query, args, err := r.builder.
		Select("id", "name", "last_name", "email", "role").
		From("users").
		Limit(uint64(offset)).
		Offset(uint64(offset)).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("failed build sql query: %v", err)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}

		err := rows.Scan(&user.ID, &user.Name, &user.LastName, &user.Email, &user.Role)
		if err != nil {
			return nil, fmt.Errorf("failed scan: %v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil && err != pgx.ErrNoRows {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) SelectByID(ctx context.Context, userID *models.UserID) (*models.User, error) {

	query, args, err := r.builder.
		Select("id", "name", "last_name", "email", "role").
		From("users").
		Where("id = ?", userID.ID).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("failed build sql query: %v", err)
	}

	user := &models.User{}
	err = r.db.QueryRow(ctx, query, args...).Scan(&user.ID, &user.Name, &user.LastName, &user.Email, &user.Role)
	if err != nil && err != pgx.ErrNoRows {
		return nil, fmt.Errorf("failed scan: %v", err)
	}
	return user, nil
}

func (r *UserRepository) Insert(ctx context.Context, userC *models.UserCreate) (*models.User, error) {

	query, args, err := r.builder.
		Insert("users").
		Columns("email", "name", "last_name").
		Values(userC.Email, userC.Name, userC.LastName).
		Suffix("RETURNING \"id\", \"name\", \"last_name\", \"email\", \"role\"").
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("failed build sql query: %v", err)
	}

	user := &models.User{}
	err = r.db.QueryRow(ctx, query, args...).Scan(&user.ID, &user.Name, &user.LastName, &user.Email, &user.Role)
	if err != nil {
		return nil, fmt.Errorf("failed scan: %v", err)
	}

	return user, nil
}

func (r *UserRepository) Update(ctx context.Context, userU *models.UserUpdate) error {

	query, args, err := r.builder.
		Update("users").
		Set("name", userU.Name).
		Set("last_name", userU.LastName).
		Set("email", userU.Email).
		Set("role", userU.Role).
		Where("id = ?", userU.ID).
		ToSql()

	if err != nil {
		return fmt.Errorf("failed build sql query: %v", err)
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed exec query: %v", err)
	}

	return nil
}

func (r *UserRepository) Delete(ctx context.Context, userID *models.UserID) error {

	query, args, err := r.builder.
		Delete("users").
		Where("id = ?", userID.ID).
		ToSql()

	if err != nil {
		return fmt.Errorf("failed build sql query: %v", err)
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed exec query: %v", err)
	}

	return nil
}
