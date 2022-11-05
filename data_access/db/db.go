package db

import (
	"database/sql"
	"fmt"
	"github.com/Defake/day-assistant/data_access/migrations"
	// _ "log"

	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "habits_assistant"
)

var Connection *sql.DB = nil

func ConnectDb() {
	fmt.Print("Connecting to the database... ")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Failed.")
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Failed.")
		panic(err)
	}

	fmt.Println("Connected.")

	fmt.Print("Running migrations... ")
	err = migrations.RunMigrations(db)
	if err != nil {
		fmt.Println("Failed.")
		panic(err)
	}
	fmt.Println("Migrated.")

	Connection = db
}

