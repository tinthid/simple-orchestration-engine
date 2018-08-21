package workflow

import (
	"encoding/json"
)

func (w *Workflow) ParseJson(jsonString string) error {
	err := json.Unmarshal([]byte(jsonString), w)
	if err != nil {
		return err
	}
	var workflowTask []Task
	for _, t := range w.Tasks {
		curTask := (*t).(map[string]interface{})
		if curTask["type"] == "async" {
			newTask := AsyncTask{
				Name: curTask["type"].(string),
				Type: curTask["type"].(string),
			}
			workflowTask = append(workflowTask, &newTask)
		} else if curTask["type"] == "simple" {
			newTask := SimpleTask{
				Name: curTask["type"].(string),
				Type: curTask["type"].(string),
			}
			workflowTask = append(workflowTask, &newTask)
		}
	}
	return nil
}

/*
func (w Workflow) AddTask(task Task) error {
	return nil
}
*/
