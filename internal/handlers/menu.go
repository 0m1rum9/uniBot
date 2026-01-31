package handlers

import (
	"context"
	"fmt"
	"time"
	"unibot/internal/services"
	"unibot/pkg/adapters/unix"

	"gopkg.in/telebot.v4"
)

type menu struct {
	u *services.UserService
}


func (m menu) Handle(c telebot.Context) error {
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

	buttons := &telebot.ReplyMarkup{}

	modules, err := unix.GetModules(token)
	if err != nil {
		return err
	}
	for _, module := range modules {
		buttons.InlineKeyboard = append(buttons.InlineKeyboard, []telebot.InlineButton{
			telebot.InlineButton{
				Unique: "module",
				Data:   fmt.Sprintf("%d", module.Id),
				Text:   module.Title,
			},
		})
	}

	return c.Send("Pick module:", buttons)

}

func NewMenu(u *services.UserService) menu {
	return menu{u: u}
}
