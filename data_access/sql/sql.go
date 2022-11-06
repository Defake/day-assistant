package sql

import (
	db "github.com/Defake/day-assistant/data_access/db"
	// sql "database/sql"
	jsoni "github.com/Defake/day-assistant/data_access/json"
	"strconv"
	"log"
)

func UpsertRecord(table string, id uint64, body jsoni.JsonSerializable) error {
	jsonString, err := body.ToJsonString()
	if err != nil {
		return err
	}
	
	jsonBody := "'" + jsonString + "'"
	query := "INSERT INTO " + table +
		"(id, created_at, updated_at, body) " +
		"VALUES (" + strconv.FormatUint(id, 10) + ", NOW(), NOW(), " + jsonBody + ") " +
		"ON CONFLICT (id) DO UPDATE SET " +
		"updated_at = NOW(), " +
		"body = " + jsonBody + ";";
	
	_, err = db.Connection.Exec(query)

	return err
}

func ReadRecords(table string) ([]string, error) {
	rows, err := db.Connection.Query("SELECT body || jsonb_build_object('id', id) FROM " + table + ";")
	if err != nil {
		log.Fatal(err)
	}
	
	defer rows.Close()

	var results []string
	
	for rows.Next() {
		var body string
		if err := rows.Scan(&body); err != nil {
			log.Fatal(err)
		}
		results = append(results, body)
	}

	return results, nil
}

