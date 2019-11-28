package lock

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"sync"
	"time"
	"xhyl-micro/service/common/cache"
)

/*
	使用Redis实现分布式锁
*/

var (
	redisClient *redis.Client
	mutex       sync.Mutex
)

func init() {
	redisClient = cache.NewRedisClient()
}

//存入方法名作为key，生成一个uuid作为value，并设置过期时间
func SetRedisLock(method string) {
	for {
		mutex.Lock()
		lock, err := redisClient.Get(key(method)).Result()
		if err == nil && lock != "" {
			mutex.Unlock()
			continue
		} else {
			redisClient.Set(key(method), uuid.New(), time.Second*3)
			mutex.Unlock()
			break
		}
	}
}

//释放分布式锁
func ReleaseRedisLock(method string) {
	redisClient.Del(key(method))
}

func key(key string) string {
	return fmt.Sprintf("distri_lock:%s", key)
}
