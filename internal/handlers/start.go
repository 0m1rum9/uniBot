package handlers

import (
	"unibot/internal/db"

	"github.com/redis/go-redis/v9"
	"gopkg.in/telebot.v4"
)


type Start struct {
  r *redis.Client
  db *db.DBTX
}
func(s Start) handle(c telebot.Context) error {
  
  if s.r.Get()
  
  
  return c.Send("/register <login> <password> to continue")
}
