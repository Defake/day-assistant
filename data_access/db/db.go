package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.11"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "habits_assistant"
)

func ConnectDb() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	email := "admin@test.com"
	rows, err := db.Query("SELECT body->>'name' FROM users WHERE body->>'email' = $1", email)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var (
		name string
	)

	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
