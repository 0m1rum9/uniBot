package handlers

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"unibot/internal/repository"
	"unibot/pkg/adapters/unix"

	"gopkg.in/telebot.v4"
)

type course struct {
	u repository.UserRepository
}

func (cr course) Handle(c telebot.Context) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	user, err := cr.u.GetUser(ctx, c.Chat().ID)
	if err != nil {
		return err
	}

	token, err := unix.Login(user.Login, user.Password)
	if err != nil {
		return err
	}

	courseID, err := strconv.Atoi(c.Callback().Data)
	if err != nil {
		return err
	}

	course, err := unix.GetCourse(token, courseID)
	if err != nil {
		return err
	}
	buttons := &telebot.ReplyMarkup{}
	row := []telebot.InlineButton{}

	for i, topic := range course.Topics {
		if !topic.IsPublic {
			continue 
		}
		text := fmt.Sprintf("%d. ", i + 1) + topic.Title
		if topic.IsPass {
			text += "✅"
		} else {
			text += "❌"
		}
		row = append(row, telebot.InlineButton{
			Data:   fmt.Sprintf("%d:%d", topic.Id, course.Id),
			Unique: "topic",
			Text:   text,
		})

		buttons.InlineKeyboard = append(buttons.InlineKeyboard, row)
		row = []telebot.InlineButton{}
	}
	return c.Edit("Watch", buttons)

}

func NewCourse(u repository.UserRepository) course {
	return course{
		u: u,
	}
}
