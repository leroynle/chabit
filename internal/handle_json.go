package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Tasks struct {
	TaskID    int64
	Title     string
	Duration  string
	WeekStart string
	Goal      int64
	Complete  int64
}

func GetDataFromJsonTasks() []Tasks {
	byteValue, err := ioutil.ReadFile("data/tasks.json")
	if err != nil {
		log.Fatal(err)
	}
	var tasks []Tasks
	json.Unmarshal([]byte(byteValue), &tasks)
	return tasks
}

func WriteDataToJsonTasks(tasks []Tasks) {
	file, _ := json.MarshalIndent(tasks, "", " ")
	err := ioutil.WriteFile("data/tasks.json", file, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func GetDataFromJsonUtilites() map[string]interface{} {
	byteValue, err := ioutil.ReadFile("data/utilites.json")
	if err != nil {
		log.Fatal(err)
	}
	var result map[string]interface{}
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func WriteDataToJsonUtilites(val map[string]interface{}) {
	byteValue, err := json.MarshalIndent(val, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("data/utilites.json", byteValue, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
