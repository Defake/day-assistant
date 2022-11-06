package tasks

import (
	// "fmt"
	sql "github.com/Defake/day-assistant/data_access/sql"
	"log"
	"encoding/json"
	// "time"
)

// https://gobyexample.com/structs

type Task struct {
	Id         *uint64 `json:"id,omitempty"`
  Name       string `json:"name"`
	IsProject  bool `json:"isProject"`
	// weekDays   uint8 `json:"weekDays"`
	// time       time.Time `json:"time"`
}

func (t *Task) ToJsonString() (string, error) {
	bytes, err := json.Marshal(t)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func FromJsonString(body string) *Task {
	var task Task
  err := json.Unmarshal([]byte(body), &task)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &task
}

func SaveTaskToDb(task *Task) error {
	id := task.Id
	task.Id = nil
	err := sql.UpsertRecord("tasks", *id, task)
	return err
}

func SaveTask() {
	taskId := uint64(1234)
	t := &Task{Id: &taskId, Name: "Test task", IsProject: true}
	err := SaveTaskToDb(t)
	if err != nil {
		log.Printf("Upsert error: %s", err)		
	}

	tasks, err := sql.ReadRecords("tasks")
	if err != nil {
		log.Fatal(err)
	}

	for _, taskJson := range tasks {
		tt := FromJsonString(taskJson)
		log.Printf("%v\n", tt)
	}
 
}
