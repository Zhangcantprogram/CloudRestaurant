package tool

import (
	"context"
	"fmt"
	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type RedisStore struct {
	client *redis.Client
}

var RediStore RedisStore

// 初始化redis
func InitRedisStore() *RedisStore {
	redisConfig := GetConfig().Redis

	//对变量client进行初始化
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr + ":" + redisConfig.Port,
		Password: redisConfig.Password,
		DB:       redisConfig.Db,
	})

	RediStore := RedisStore{client: client}
	base64Captcha.SetCustomStore(&RediStore)

	fmt.Println("redis--------->", RediStore)

	return &RediStore

}

// set方法
func (rs *RedisStore) Set(id string, value string) {
	ctx := context.Background()
	err := rs.client.Set(ctx, id, value, time.Minute*10).Err()
	if err != nil {
		log.Println(err)
	}
}

// get方法
func (rs *RedisStore) Get(id string, clear bool) string {
	ctx := context.Background()
	val, err := rs.client.Get(ctx, id).Result()
	if err != nil {
		log.Println(err)
		return ""
	}
	if clear {
		err := rs.client.Del(ctx, id).Err()
		if err != nil {
			log.Println(err)
			return ""
		}
	}
	return val
}
