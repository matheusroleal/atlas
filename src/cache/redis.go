/*
 * @Author: Matheus Leal
 * @Date: 2022-07-01 22:54:39
 * @Last Modified by: Matheus Leal
 * @Last Modified time: 2022-07-03 14:26:13
 */
package cache

import (
	log "github.com/sirupsen/logrus"

	"github.com/go-redis/redis"
)

/**
 * Start a new connection with a Redis server.
 *
 * @param   address				string				Redis server address
 *					password			string	 			Redis server password
 * @return  							redis.Client	A Redis server client
 */
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

/**
 * Get a data from a Redis server.
 *
 * @param		address				string				Redis server address
 *					password			string	 			Redis server password
 *					key						string	 			Redis server key
 * @return								string				A Redis server key value
 */
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

/**
 * Insert a data from a Redis server.
 *
 * @param		address				string				Redis server address
 *					password			string	 			Redis server password
 *					key						string	 			Redis server key
 *					value					string	 			Redis server key value
 * @return								error					A Redis server error
 */
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

/**
 * Append a data for a existent key at a Redis server.
 *
 * @param		address				string				Redis server address
 *					password			string	 			Redis server password
 *					key						string	 			Redis server key
 *					value					string	 			Redis server key value
 * @return								error					A Redis server error
 */
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
