// Package handlers provide
package handlers

import (
	 "unibot/internal/services"
	"unibot/internal/middlewares"
	"unibot/internal/repository"

	"gopkg.in/telebot.v4"
	"gopkg.in/telebot.v4/middleware"
)

func SetHandlers(b *telebot.Bot,
	u repository.UserRepository,
	n *services.NotificationService,
	) {

	// Global handler for registered users
	b.Use(middlewares.NewRegisterMiddleware(
		NewMenu(u).Handle, u,
	).Handle, middleware.AutoRespond())

	b.Handle(
		telebot.OnText,
		NewStart(u).Handle,
	)
	// TODO create service layer
	b.Handle("/register", NewRegister(u).Handle)

	b.Handle(&telebot.InlineButton{Unique: "module"}, NewModule(u).Handle)
	b.Handle(&telebot.InlineButton{Unique: "course"}, NewCourse(u).Handle)
	b.Handle(&telebot.InlineButton{Unique: "course"}, NewCourse(u).Handle)

	b.Handle(&telebot.InlineButton{Unique: "topic"}, NewTopic(u, n).Handle)
}
