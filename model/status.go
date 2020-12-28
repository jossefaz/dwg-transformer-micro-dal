package model

import (
	"fmt"
	tables "github.com/jossefaz/dwg-transformer-micro-data-struct"
	"github.com/jossefaz/dwg-transformer-micro-dal/log"
)

func StatusUpdate(db *CDb, where map[string]interface{}, update map[string]interface{}) ([]byte, error) {
	_, err := db.Model(tables.Cad_check_status{}).Where(where).Updates(update).Rows()
	if err != nil {
		log.Logger.Log.Error("cannot update rows ",err)
	}
	return []byte(fmt.Sprintf(string(db.RowsAffected), " rows were updated")), nil
}
