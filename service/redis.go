package service

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/redis.v5"
)

var redisClient *redis.ClusterClient

//链接redis集群
func SetupRedis(addrs []string, pwd string) *redis.ClusterClient {
	redisClient = redis.NewClusterClient(&redis.ClusterOptions{
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
	pong, err := redisClient.Ping().Result()
	if err != nil {
		fmt.Println(err, "connect redis server failed!", pong)
		os.Exit(0)
	}
	fmt.Println("connect redis success!")
	return redisClient
}
