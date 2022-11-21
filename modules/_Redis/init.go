package _Redis

import (
	"github.com/go-redis/redis"
)

var Redis *redis.Client

var Redis_IP string

func RedisInit() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     Redis_IP,
		Password: "", //默认空密码
		DB:       0,  //使用默认数据库
	})
	// defer client.Close() //最后关闭

	_, err := Redis.Ping().Result()
	if err != nil {
		panic("Redis Connect Error "+Redis_IP)
	}

	// fmt.Println("Connected result: ", pong)
}
