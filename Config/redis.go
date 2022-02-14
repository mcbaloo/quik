package Config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

func RedisConnection() redis.Client {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("eror occured while loading environment varaible")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	return *client
}
func SetRedisData(client *redis.Client, key string, data interface{}) error {
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return client.Set(key, value, 0).Err()
}
func GetRedisData(c *redis.Client, key string, value interface{}) (string, error) {
	data, err := c.Get(key).Result()
	if err != nil {
		return "", err
	}
	return redis.Nil.Error(), json.Unmarshal([]byte(data), value)
}
func RemoveRedisData(client *redis.Client, key string) error {
	return client.Del(key).Err()
}
