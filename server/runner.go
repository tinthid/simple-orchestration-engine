package server

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
)

func (s *Server) taskRunner(data []byte) []byte {
	mapData := make(map[string]interface{})
	jsonString := string(data)
	err := json.Unmarshal(data, &mapData)
	if err != nil {
		return []byte(`error`)
	}

	/*

		"inputParams": {
			"first_name":
			"last_name":
			"email_address":
		}

	*/
	if _, nameCheck := mapData["name"]; nameCheck {
		result := gjson.Get(jsonString, "data.first_name")
		fmt.Println(result)

		// create request payload
	}
	return []byte(`error`)
	//if mapData["name"]
}
