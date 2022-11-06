package sql

import (
	db "github.com/Defake/day-assistant/data_access/db"
	// sql "database/sql"
	"log"
	"strconv"

	jsoni "github.com/Defake/day-assistant/data_access/json"
	meta "github.com/Defake/day-assistant/domain/meta"
)

func UpsertRecord(table string, id uint64, body jsoni.JsonSerializable) error {
	jsonString, err := body.ToJsonString()
	if err != nil {
		return err
	}

	doesHaveId := id != 0
	var idParam string
	var idValue string
	if doesHaveId {
		idParam = "id, "
		idValue = strconv.FormatUint(id, 10) + ", "
	}
	
	jsonBody := "'" + jsonString + "'"
	query := "INSERT INTO " + table +
		"(" + idParam + "created_at, updated_at, body) " +
		"VALUES (" + idValue + "NOW(), NOW(), " + jsonBody + ") " +
		"ON CONFLICT (id) DO UPDATE SET " +
		"updated_at = NOW(), " +
		"body = " + jsonBody + ";";
	
	_, err = db.Connection.Exec(query)

	return err
}

func ReadRecords(table string) ([]meta.ObjectWithMeta, error) {
	rows, err := db.Connection.Query("SELECT " +
		" body, " +
		" jsonb_build_object('id', id) " +
		"FROM " + table + ";")
	if err != nil {
		return nil, err
	}
	
	defer rows.Close()

	var results []meta.ObjectWithMeta
	
	for rows.Next() {
		var metaDataJson string
		var bodyJson string
		if err := rows.Scan(&bodyJson, &metaDataJson); err != nil {
			log.Fatal(err)
		}

		metaData, err := meta.FromJsonString(metaDataJson)
		if err != nil {
			return nil, err
		}

		obj := meta.ObjectWithMeta {Meta: *metaData, JsonBody: bodyJson}
		results = append(results, obj)
	}

	return results, nil
}

