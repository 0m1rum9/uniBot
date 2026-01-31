// Package services provides service for user entity
package services

import (
	"context"
	"unibot/internal/db"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/redis/go-redis/v9"
)

type UserService struct {
	db *db.Queries
	r  *redis.Client
}

func CreateUserService(
	db *db.Queries,
	r *redis.Client,
) *UserService {
	return &UserService{db: db, r: r}
}

func (u *UserService) IsRegistered(ctx context.Context, chatID int64) (bool, error) {
	_, err := u.GetUser(ctx, chatID)
	if err == pgx.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, err
}

func (u *UserService) GetUser(ctx context.Context, chatID int64) (User, error) {
	user, err := u.db.GetUser(ctx, pgtype.Int8{Int64: chatID, Valid: true})
	if err != nil {
		return User{}, err
	}
	return mapToDTO(user), nil
}

func (u *UserService) CreateUser(ctx context.Context, user User) error {
	err := u.db.CreateUser(ctx, mapToModel(user))
	return err
}
