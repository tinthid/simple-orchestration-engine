package server

import (
	"fmt"

	"github.com/tinthid/simple-orchestration-engine/rabbitmq"
	"github.com/tinthid/simple-orchestration-engine/redis"
	"github.com/tinthid/simple-orchestration-engine/workflow"
)

type Server struct {
	RedisConn    *redis.Redis
	RabbitMQConn *rabbitmq.RabbitMQ
	Workflows    map[string]*workflow.Workflow
}

func CreateServer(redisConn *redis.Redis, rabbitMQConn *rabbitmq.RabbitMQ) (s *Server) {
	s = new(Server)
	s.Workflows = make(map[string]*workflow.Workflow)
	s.Workflows["create_order_tx"] = &workflow.Workflow{}
	s.Workflows["create_user_tx"] = &workflow.Workflow{}
	s.RedisConn = redisConn
	s.RabbitMQConn = rabbitMQConn
	return
}

func (s *Server) Run() {

	for workflowName, workflowValue := range s.Workflows {
		fmt.Printf("key[%s] value[%s]\n", workflowName, workflowValue)
		// Create RabbitMQ Server per Workflow here
	}

	s.RabbitMQConn.ServerRPC("fillgoods_oe", "workflow.run", "fillgoods_oe", s.taskRunner)
	data := []byte(`{
		"name": "create_user_tx",
		"data": {
		  "first_name": "Tinthid",
		  "last_name": "Jaikla",
		  "email_address": "tinthid.456@gmail.com"
		}
	  }`)
	fmt.Println(1)
	s.taskRunner(data)

	//s.RabbitMQConn.ServerRPC("")

}
