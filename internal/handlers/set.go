// Package handlers provide
package handlers

import (
	"unibot/internal/db"
	"unibot/internal/middlewares"
	"unibot/internal/services"

	"github.com/redis/go-redis/v9"
	"gopkg.in/telebot.v4"
	"gopkg.in/telebot.v4/middleware"
)

func SetHandlers(b *telebot.Bot,
	db *db.Queries,
	r *redis.Client,
	u *services.UserService) {

	// Global handler for registered users
	b.Use(middlewares.NewRegisterMiddleware(
		NewMenu(u).Handle, u,
	).Handle, middleware.AutoRespond())

	b.Handle(
		telebot.OnText,
		NewStart(u).Handle,
	)

	b.Handle("/register", NewRegister(u).Handle)
	b.Handle(&telebot.InlineButton{Unique: "module"}, NewModule(u).Handle)
	b.Handle(&telebot.InlineButton{Unique: "course"}, NewCourse(u).Handle)
	b.Handle(&telebot.InlineButton{Unique: "topic"}, NewTopic(u).Handle)
	b.Handle(&telebot.InlineButton{Unique: "course"}, NewCourse(u).Handle)

}
