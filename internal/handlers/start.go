package handlers

import (
	"unibot/internal/services"

	"gopkg.in/telebot.v4"
)

type start struct {
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
