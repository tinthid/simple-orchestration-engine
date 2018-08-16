package redis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-redis/redis"
)

type Redis struct {
	RedisClient *redis.Client
}

func CreateRedis() (*Redis, error) {
	redisObject := new(Redis)
	jsonFile, err := os.Open("config/redis.json")

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	redisJson := make(map[string]interface{})

	json.Unmarshal(byteValue, &redisJson)

	endpoint := redisJson["redis_endpoint"].(string)
	password := redisJson["password"].(string)
	// database := redisJson["database_name"].(string)

	redisObject.RedisClient = redis.NewClient(&redis.Options{
		Addr:     endpoint,
		Password: password,
		DB:       0,
	})

	pong, err := redisObject.RedisClient.Ping().Result()
	if err != nil {
		fmt.Println(pong, err)
	}

	return redisObject, err
}
