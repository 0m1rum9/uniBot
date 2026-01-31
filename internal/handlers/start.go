package handlers

import (
	"unibot/internal/db"
	"unibot/internal/services"

	"github.com/redis/go-redis/v9"
	"gopkg.in/telebot.v4"
)

type start struct {
	r  *redis.Client
	db *db.Queries
	u  *services.UserService
}

func (s start) Handle(c telebot.Context) error {

	return c.Send("/register <login> <password> to continue")
}

func NewStart(u *services.UserService) start {
	return start{
		u: u,
	}
}
