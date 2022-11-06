package meta

import (
  json "encoding/json"
)

type MetaData struct {
	Id uint64 `json:"id"`
}

func FromJsonString(body string) (*MetaData, error) {
	var data MetaData
  err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}


type ObjectWithMeta struct {
	Meta MetaData
	JsonBody string
}

type WithMeta interface {
	SetId(uint64)
}

