package service

import (
	"time"

	"gopkg.in/redis.v5"
)

func CacheSet(key string, val interface{}, timeOut time.Duration) error {
	return redisClient.Set(key, val, timeOut).Err()
}

func CacheGet(key string) (string, error) {
	return redisClient.Get(key).Result()
}

func CacheDel(key string) {
	redisClient.Del(key)
}

func CacheExist(key string) bool {
	return redisClient.Exists(key).Val()
}

func GetRedisClient() *redis.ClusterClient {
	return redisClient
}

func SessionSet(key string, field string, val interface{}) error {
	return redisClient.HSet(key, field, val).Err()
}

func SessionGetAll(key string) (map[string]string, error) {
	return redisClient.HGetAll(key).Result()
}

func SessionGet(key string, field string) *redis.StringCmd {
	return redisClient.HGet(key, field)
}

func SessionIsFieldExist(key string, field string) bool {
	return redisClient.HExists(key, field).Val()
}

func SessonFieldDelete(key string, field string) error {
	return redisClient.HDel(key, field).Err()
}
