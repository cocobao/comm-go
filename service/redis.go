package service

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/redis.v5"
)

var (
	redisClientCluster *redis.ClusterClient
	redisClient        *redis.Client
)

func SetupRedis(addrs string) *redis.Client {
	redisClient = redis.NewClient(&redis.Options{
		Addr:         addrs + ":6379",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
	return redisClient
}

//链接redis集群
func SetupRedisCluster(addrs []string, pwd string) *redis.ClusterClient {
	redisClientCluster = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:              addrs,
		Password:           pwd,
		MaxRedirects:       16,
		ReadOnly:           true,
		RouteByLatency:     true,
		DialTimeout:        10000 * time.Millisecond,
		ReadTimeout:        30000 * time.Millisecond,
		WriteTimeout:       30000 * time.Millisecond,
		PoolSize:           10,
		PoolTimeout:        35000 * time.Millisecond,
		IdleTimeout:        600 * time.Second,
		IdleCheckFrequency: 60 * time.Second,
	})
	pong, err := redisClientCluster.Ping().Result()
	if err != nil {
		fmt.Println(err, "connect redis server failed!", pong)
		os.Exit(0)
	}
	fmt.Println("connect redis success!")
	return redisClientCluster
}
