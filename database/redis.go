package database

import (
	"github.com/go-redis/redis"
	"password-manager-backend/logger"
)

var Redis *redis.Client

func ConnectRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	_, err := Redis.Ping().Result()

	if err != nil {
		logger.LogWarning("Couldn't connect to Redis database")
	} else {
		logger.LogInfo("Connected to Redis database")
	}
}
