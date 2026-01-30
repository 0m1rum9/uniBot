// Package cache provides caching abstraction
package cache

import "github.com/redis/go-redis/v9"

type Cache interface {
  SetState()
}

type Redis struct {
  redis *redis.Client
}


func(r *Redis) SetState() {
  
}
