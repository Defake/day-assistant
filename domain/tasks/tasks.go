package tasks

import (
	sql "github.com/Defake/day-assistant/data_access/sql"
	// "log"
	log "github.com/Defake/day-assistant/util/logging"
	"encoding/json"
	// "time"
)

// https://gobyexample.com/structs

type Task struct {
	Id         uint64 `json:"id,omitempty"`
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
		log.Error.Println(err)
		return nil
	}

	return &task
}

func ReadTasks() []*Task {
	objects, err := sql.ReadRecords("tasks")
	if err != nil {
		log.Error.Println(err)
	}

	var result []*Task
	for _, object := range objects {
		body := FromJsonString(object.JsonBody)
		result = append(result, &Task{
			Id: object.Meta.Id,
			Name: body.Name,
			IsProject: body.IsProject,
		})
	}

	return result
}

func SaveTask(task *Task) error {
	body := Task{
		Name: task.Name,
		IsProject: task.IsProject}
	err := sql.UpsertRecord("tasks", task.Id, &body)
	return err
}

func SaveTaskExample() {
	t := &Task{Id: 1234, Name: "Test task", IsProject: true}
	err := SaveTask(t)
	if err != nil {
		log.Error.Printf("Upsert error: %s\n", err)		
	}

	tasks := ReadTasks()
	for _, task := range tasks {
		log.Info.Printf("%v\n", task)
	}
 
}
