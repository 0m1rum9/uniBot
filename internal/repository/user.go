// Package repository provides service for user entity
package repository

import (
	"context"
	"unibot/internal/db"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/redis/go-redis/v9"
)

type UserRepository interface {
	IsRegistered(ctx context.Context, chatID int64) (bool, error)
	GetUser(ctx context.Context, chatID int64) (User, error)
	CreateUser(ctx context.Context, user User) error 

}

type user struct {
	db *db.Queries
	r  *redis.Client
}

func NewUserRepository(
	db *db.Queries,
	r *redis.Client,
) *user {
	return &user{db: db, r: r}
}

func (u *user) IsRegistered(ctx context.Context, chatID int64) (bool, error) {
	_, err := u.GetUser(ctx, chatID)
	if err == pgx.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, err
}

func (u *user) GetUser(ctx context.Context, chatID int64) (User, error) {
	user, err := u.db.GetUser(ctx, pgtype.Int8{Int64: chatID, Valid: true})
	if err != nil {
		return User{}, err
	}
	return mapToDTO(user), nil
}

func (u *user) CreateUser(ctx context.Context, user User) error {
	err := u.db.CreateUser(ctx, mapToModel(user))
	return err
}
