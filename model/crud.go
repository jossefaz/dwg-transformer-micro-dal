package model

import (
	"encoding/json"
	"fmt"
)

func Retrieve(atts interface{}, db *CDb, keyval map[string]interface{}) ([]byte, error) {
	db.Where(keyval).Find(atts)
	b, _ := json.Marshal(atts)
	return b, nil
}

func Create(atts interface{}, db *CDb) ([]byte, error) {
	db.Create(atts)
	return []byte(fmt.Sprintf(string(db.RowsAffected), " rows were updated")), nil
}
