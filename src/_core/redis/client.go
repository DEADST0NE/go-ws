package redis

import (
	"context"
	"os"

	myContext "exex-chart/src/_core/context"

	log "github.com/sirupsen/logrus"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

var Client *redis.Client

var Nil = redis.Nil

func NewRedisClient(redisURL string) *redis.Client {
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Errorf("ERROR CONNECTING TO REDIS: %v", err)
		os.Exit(1)
	}

	client := redis.NewClient(opts)

	// Проверка соединения
	_, err = client.Ping(Ctx).Result()
	if err != nil {
		log.Errorf("ERROR CONNECTING TO REDIS: %v", err)
		os.Exit(1)
	}

	log.Info("CONNECTED TO REDIS")
	return client
}

func Init() {
	Client = NewRedisClient(myContext.Config.Redis.Host)

	if Client == nil {
		log.Error("FAILED TO CONNECT TO REDIS")
		os.Exit(1)
		return
	}
}
