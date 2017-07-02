package wechat

import (
	"github.com/go-redis/redis"
)

type Wechat struct {
	AppID     string
	AppSecret string
}

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
}
