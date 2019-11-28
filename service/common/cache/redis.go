package cache

import (
	"github.com/go-redis/redis"
	"time"
)

func NewRedisClient(db ...int) *redis.Client {
	dbCode := 0
	if len(db) > 0 {
		dbCode = db[0]
	}

	client := redis.NewClient(&redis.Options{
		Addr:         "192.168.10.33:6379",
		PoolSize:     1000,
		ReadTimeout:  time.Millisecond * time.Duration(100),
		WriteTimeout: time.Millisecond * time.Duration(100),
		IdleTimeout:  time.Second * time.Duration(60),
		DB:           dbCode,
	})

	return client
}
