package services

import (
	"unibot/internal/db"

	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID            int64
	ChatID        int64
	Username      string
	LastMessageID int64
	Login         string
	Password      string
}

func mapToDTO(u db.User) User {
	return User{
		ChatID:        u.ChatID.Int64,
		LastMessageID: u.LastMessageID.Int64,
		Login:         u.Login.String,
		Username:      u.Username.String,
		Password:      u.Password.String,
	}
}
func mapToModel(u User) db.CreateUserParams {
	return db.CreateUserParams{
		Login:         pgtype.Text{String: u.Login, Valid: true},
		Password:      pgtype.Text{String: u.Password, Valid: true},
		ChatID:        pgtype.Int8{Int64: u.ChatID, Valid: true},
		LastMessageID: pgtype.Int8{Int64: u.LastMessageID, Valid: true},
		Username:      pgtype.Text{String: u.Username, Valid: true},
	}
}
