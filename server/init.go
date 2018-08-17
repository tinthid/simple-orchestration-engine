package server

import (
	"github.com/tinthid/simple-orchestration-engine/rabbitmq"
	"github.com/tinthid/simple-orchestration-engine/redis"
)

type Server struct {
	RedisConn    *redis.Redis
	RabbitMqConn *rabbitmq.RabbitMQ
}

func CreateServer(redisConn *redis.Redis, rabbitMqConn *rabbitmq.RabbitMQ) (s *Server) {
	s = new(Server)
	s.RedisConn = redisConn
	s.RabbitMqConn = rabbitMqConn
	return
}
