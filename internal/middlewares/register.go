// Package middlewares provides telebot.middleware for handlers
package middlewares

import (
	"context"
	"time"
	"unibot/internal/repository"

	"gopkg.in/telebot.v4"
)

type register struct {
	registeredHandler telebot.HandlerFunc
	u                 repository.UserRepository
}

func NewRegisterMiddleware(registeredHandler telebot.HandlerFunc, u repository.UserRepository) register {
	return register{registeredHandler: registeredHandler, u: u}
}

func (r register) Handle(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		if c.Callback() != nil {
			return next(c)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		isRegistered, err := r.u.IsRegistered(ctx, c.Chat().ID)

		if err != nil {
			return err
		}
		if isRegistered {
			return r.registeredHandler(c)
		}

		return next(c)
	}
}
