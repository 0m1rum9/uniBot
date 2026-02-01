// Package services provides bla
package services

import (
	"log"

	"gopkg.in/telebot.v4"
)

type NotificationService struct {
  b *telebot.Bot
}


func(n *NotificationService) Update(msg string, editable *telebot.Message) {
  
  _, err := n.b.Edit(editable, msg)
  if err != nil {
    log.Print(err)
  }
}

func(n *NotificationService) Notificate(msg string, chatID int64) (*telebot.Message, error) {
  m, err := n.b.Send(&telebot.Chat{ID: chatID}, msg)
  if err != nil {
    return nil, err
  }
  return m, nil
}

func NewNotificationService(b *telebot.Bot) *NotificationService{
  return &NotificationService {
    b: b,
  } 
}
