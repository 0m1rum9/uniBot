// Package handlers provide
package handlers

import "gopkg.in/telebot.v4"


func SetHandlers(b *telebot.Bot) {
  b.Handle(telebot.OnText, Start{}.handle)
}
