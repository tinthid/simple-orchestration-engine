package rabbitmq

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/streadway/amqp"
)

var RabbitConn *amqp.Connection

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func RabbitMQInit() {
	jsonFile, err := os.Open("config/rabbitmq.json")
	
	if err != nil {
		fmt.Println(err) 	
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	rabbitmqJson := make(map[string]interface{})

	json.Unmarshal(byteValue, &rabbitmqJson)

	amqpURL := rabbitmqJson["rabbitmq_server"].(string)

	RabbitConn, err = amqp.Dial(amqpURL)
	failOnError(err, "Failed to connect to RabbitMQ")
}
