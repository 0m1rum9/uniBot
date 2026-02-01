package handlers

import (
	"unibot/internal/repository"

	"gopkg.in/telebot.v4"
)

type start struct {
	u  repository.UserRepository
}

func (s start) Handle(c telebot.Context) error {

	return c.Send("/register <login> <password> to continue")
}

func NewStart(u repository.UserRepository) start {
	return start{
		u: u,
	}
}
