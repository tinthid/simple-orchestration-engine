package workflow

type Workflow struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Version     string  `json:"version"`
	RoutingKey  string  `json:"routing_key"`
	Tasks       []*Task `json:"tasks"`
}

type RabbitMQPart struct {
}

type AsyncTask struct {
	Name string
	Type string
}

type SimpleTask struct {
	Name string
	Type string
}

type Task interface {
	//TaskName() string
}

/*
type Task interface {
	Name     string  `json:"send_email"`
	Type     string  `json:"type"`
	Tasks    *[]Task `json:"tasks"`
	Request  string
	Rollback string
}
*/
/*
type TaskData struct {
	Topic        string                 `json:"request_topic"`
	InputParams  map[string]interface{} `json:"input_params"`
	SuccessValue string                 `json:"success_value"`
}*/

func CreateWorkflow() (w *Workflow, err error) {
	w = new(Workflow)
	err = nil
	return
}

/*

{
  "name": "create_user",
  "description": "create user",
  "version": 1,
  "routing_key": "create_user_tx",
  "tasks": [
    {
      "name": "send_email",
      "request_topic": "email.verify_user",
      "type": "simple",
      "input_params": {
        	"email_address": "${create_user.data.email}",
        	"type": "register_success",
        	"data": {
        		"first_name" : "${create_user.data.first_name}",
        		"last_name" : "${create_user.data.last_name}"
        	}
      },
      "success_value": "${send_email.success}"
    }
  ]
}

*/

/*

{
  {
  "name": "create_order",
  "description": "create order",
  "version": 1,
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

*/
