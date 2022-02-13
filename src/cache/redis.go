package cache

import (
	"log"

	"github.com/go-redis/redis"
)

func cacheConn(address string, password string) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		panic(err.Error())
	}
	return redisClient
}

func GetData(address string, password string, key string) string {
	client := cacheConn(address, password)
	// we can call get with a `Key`.
	val, err := client.Get(key).Result()
	if err != nil {
		log.Println("[CACHE] ", err)
		return ""
	}
	log.Println("[CACHE] ", val)
	return val
}

func SetData(address string, password string, key string, value string) {
	client := cacheConn(address, password)
	// we can call set with a `Key` and a `Value`.
	err := client.Set(key, value, 0).Err()
	// if there has been an error setting the value
	// handle the error
	if err != nil {
		log.Println("[CACHE] ", err)
	}
}

func AppendData(address string, password string, key string, value string) {
	client := cacheConn(address, password)
	// we can call set with a `Key` and a `Value`.
	err := client.Append(key, value).Err()
	// if there has been an error setting the value
	// handle the error
	if err != nil {
		log.Println("[CACHE] ", err)
	}
}
