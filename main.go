package main

import (
	"fmt"

	"github.com/tinthid/simple-orchestration-engine/workflow"

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

	jsonString := `
		
			{
			"name": "create_order",
			"description": "create order",
			"version": "1",
			"routing_key": "create_order_tx",
			"tasks": [
			{
				"name": "customer_stock_tx",
				"type": "async",
				"tasks": [
				{
					"name": "upsert_customer",
					"request":{
					"topic": "customer.upsert",
					"input_params": {
						"customer_name": "Tinthid Jaikla"
					},
					"success_condition": "${create_order.response.success}"
					},
					"rollback": {
					"topic": "delete_customer",
					"input_params": {
						"customer_id": "${upsert_customer.response.id}"
					}
					}
				},
				{
					"name": "update_product_stock",
					"request": {
					"input_params": {
						"data": {
						"change": "${create_order.response.change}",
						"product_id": "${create_order.response.product_id}"
						}
					},
					"success_condition": "${update_product_stock.response.success}"
					},
					"rollback": {
					"topic": "stock.cancle_update",
					"input_params": {
						"transaction_id": "${update_product_stock.response.transaction_id}"
					}
					}
				}
				]
			}
			]
		}
	`

	wf, _ := workflow.CreateWorkflow()
	_ = wf.ParseJson(jsonString)
	s.Run()
	//fmt.Println(wf)

}
