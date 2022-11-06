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
	id         uint64
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

func SaveTask() {
	// TODO

	t := &Task{id: 1234, Name: "Test task", IsProject: true}
	err := sql.UpsertRecord("tasks", t.id, t)
	log.Printf("Upsert error: %s", err)
	
	// db.Connection.Query("")
	// email := "admin@test.com"
	// rows, err := db.Connection.Query("SELECT body->>'name' FROM users WHERE body->>'email' = $1", email)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()

	// var (
	// 	name string
	// )

	// for rows.Next() {
	// 	err := rows.Scan(&name)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println(name)
	// }

	// err = rows.Err()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
