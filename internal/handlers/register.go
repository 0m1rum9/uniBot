package handlers

import (
	"context"
	"time"
	"unibot/internal/repository"
	"unibot/pkg/adapters/unix"

	"gopkg.in/telebot.v4"
)

type register struct {
	u repository.UserRepository
}

func (r register) Handle(c telebot.Context) error {

	if len(c.Args()) < 2 {
		return c.Send("Provide both login and password")
	}
	login := c.Args()[0]
	password := c.Args()[1]

	_, err := unix.Login(login, password)

	if err == unix.ErrInvalidCredentials {
		return c.Send("Invalid credentials, try again")
	}

	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	r.u.CreateUser(ctx, repository.User{
		Username:      c.Chat().Username,
		ChatID:        c.Chat().ID,
		Login:         login,
		Password:      password,
	})
	return c.Send("Registered!\nYou can now reach menu by typing anything")

}

func NewRegister(u repository.UserRepository) register {

	return register{
		u: u,
	}
}
