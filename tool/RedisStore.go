package tool

import (
	"context"
	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type RedisStore struct {
	redisClient *redis.Client
}

var Redis *redis.Client

// 初始化redis
func InitRedisStore() *RedisStore {
	redisConfig := GetConfig().Redis

	//对全局变量global.Redis进行初始化
	Redis = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr + ":" + redisConfig.Port,
		Password: redisConfig.Password,
		DB:       redisConfig.Db,
	})

	customeStore := &RedisStore{Redis}
	base64Captcha.SetCustomStore(customeStore)

	return customeStore

}

// set方法
func (rs *RedisStore) Set(id string, value string) {
	ctx := context.Background()
	err := rs.redisClient.Set(ctx, id, value, time.Minute*10).Err()
	if err != nil {
		log.Println(err)
	}
}

// get方法
func (rs *RedisStore) Get(id string, clear bool) string {
	ctx := context.Background()
	val, err := rs.redisClient.Get(ctx, id).Result()
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	if clear {
		err := rs.redisClient.Del(ctx, id).Err()
		if err != nil {
			log.Println(err.Error())
			return ""
		}
	}
	return val
}
