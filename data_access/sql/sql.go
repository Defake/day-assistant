package sql

import (
	db "github.com/Defake/day-assistant/data_access/db"
	jsoni "github.com/Defake/day-assistant/data_access/json"
	"strconv"
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

