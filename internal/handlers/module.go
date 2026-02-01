package handlers

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"unibot/internal/services"
	"unibot/pkg/adapters/unix"

	"gopkg.in/telebot.v4"
)

type module struct {
	u *services.UserService
}

func (m module) Handle(c telebot.Context) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	user, err := m.u.GetUser(ctx, c.Chat().ID)
	if err != nil {
		return err
	}

	token, err := unix.Login(user.Login, user.Password)
	if err != nil {
		return err
	}

	moduleID, err := strconv.Atoi(c.Callback().Data)
	if err != nil {
		return err
	}
	courses, err := unix.GetCourses(token, moduleID)
	if err != nil {
		return err
	}
	// TODO pagination
	buttons := &telebot.ReplyMarkup{}
	row := []telebot.InlineButton{}
	for i, course := range courses {
		if course.IsPublic != true {
			continue
		}
		row = append(row, telebot.InlineButton{
			Unique: "course",
			Text:   course.Title,
			Data:   fmt.Sprintf("%d", course.Id),
		})

		if i%2 == 0 || i == len(courses)-1 {
			buttons.InlineKeyboard = append(buttons.InlineKeyboard, row)
			row = []telebot.InlineButton{}
		}
	}

	return c.Edit("Pick course:", buttons)
}

func NewModule(u *services.UserService) module {
	return module{
		u: u,
	}
}
