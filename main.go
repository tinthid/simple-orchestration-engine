package main

import (
	"fmt"

	"github.com/tinthid/simple-orchestration-engine/rabbitmq"
	"github.com/tinthid/simple-orchestration-engine/redis"
	"github.com/tinthid/simple-orchestration-engine/server"
)

func main() {
	var err error
	var redisServer *redis.Redis
	var rabbitmqServer *rabbitmq.RabbitMQ

	rabbitmqServer, err = rabbitmq.CreateRabbitMQ()
	if err != nil {
		fmt.Println("Create RabbitMQ Server error, messsage: ", err)
		return
	}

	redisServer, err = redis.CreateRedis()
	if err != nil {
		fmt.Println("Create Redis Server error, messsage: ", err)
		return
	}

	s := server.CreateServer(redisServer, rabbitmqServer)
	fmt.Println(s)
}
