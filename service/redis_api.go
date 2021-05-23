package service

import (
	"time"

	"gopkg.in/redis.v5"
)

func ClearCache() {
	redisClient.FlushAll()
}

func CacheSet(key string, val interface{}, timeOut time.Duration) error {
	if redisClient != nil {
		return redisClient.Set(key, val, timeOut).Err()
	}
	return redisClientCluster.Set(key, val, timeOut).Err()
}

func CacheGet(key string) (string, error) {
	if redisClient != nil {
		return redisClient.Get(key).Result()
	}
	return redisClientCluster.Get(key).Result()
}

func CacheDel(key string) {
	if redisClient != nil {
		redisClient.Del(key)
		return
	}
	redisClientCluster.Del(key)
}

func CacheExist(key string) bool {
	if redisClient != nil {
		return redisClient.Exists(key).Val()
	}
	return redisClientCluster.Exists(key).Val()
}

func GetredisClientCluster() *redis.ClusterClient {
	return redisClientCluster
}

func SessionSet(key string, field string, val interface{}) error {
	if redisClient != nil {
		return redisClient.HSet(key, field, val).Err()
	}
	return redisClientCluster.HSet(key, field, val).Err()
}

func SessionGetAll(key string) (map[string]string, error) {
	if redisClient != nil {
		return redisClient.HGetAll(key).Result()
	}
	return redisClientCluster.HGetAll(key).Result()
}

func SessionGet(key string, field string) *redis.StringCmd {
	if redisClient != nil {
		return redisClient.HGet(key, field)
	}
	return redisClientCluster.HGet(key, field)
}

func SessionIsFieldExist(key string, field string) bool {
	if redisClient != nil {
		return redisClient.HExists(key, field).Val()
	}
	return redisClientCluster.HExists(key, field).Val()
}

func SessonFieldDelete(key string, field string) error {
	if redisClient != nil {
		return redisClient.HDel(key, field).Err()
	}
	return redisClientCluster.HDel(key, field).Err()
}
