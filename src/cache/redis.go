/*
 * @Author: Matheus Leal
 * @Date: 2022-07-01 22:54:39
 * @Last Modified by: Matheus Leal
 * @Last Modified time: 2022-07-02 11:57:48
 */
package cache

import (
	log "github.com/sirupsen/logrus"

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
		log.Error(err)
		return nil
	}
	return redisClient
}

func GetData(address string, password string, key string) string {
	client := cacheConn(address, password)
	// we can call get with a `Key`.
	val, err := client.Get(key).Result()
	if err != nil {
		log.Error("[CACHE] ", err)
		return ""
	}
	log.Debug("[CACHE] ", val)
	return val
}

func SetData(address string, password string, key string, value string) error {
	client := cacheConn(address, password)
	// we can call set with a `Key` and a `Value`.
	err := client.Set(key, value, 0).Err()
	// if there has been an error setting the value
	// handle the error
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func AppendData(address string, password string, key string, value string) error {
	client := cacheConn(address, password)
	// we can call set with a `Key` and a `Value`.
	err := client.Append(key, value).Err()
	// if there has been an error setting the value
	// handle the error
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
