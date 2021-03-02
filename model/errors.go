package model

import (
	"encoding/json"
	"fmt"
	"github.com/jossefaz/dwg-transformer-micro-dal/log"
	tables "github.com/jossefaz/dwg-transformer-micro-data-struct"
	globalUtils "github.com/jossefaz/dwg-transformer-micro-utils/utils"
	"reflect"
)

func ErrorsRetrieve(db *CDb, keyval map[string]interface{}) ([]byte, error) {
	atts := []tables.CAD_check_errors{}
	db.Where(keyval).Find(&atts)
	b, _ := json.Marshal(atts)
	return b, nil
}

func Lut_Error_Retrieve(db *CDb, keyval map[string]interface{}) map[string]interface{} {
	sql, err := globalUtils.GetEnv("LOOKUP_ERRORS_SQL")
	if sql == "" {
		fmt.Printf("Environment variable %s must be set in order to update the error code correctly this service work", "LOOKUP_ERRORS_SQL")
	}
	if err != nil {
		log.Logger.Log.Error("Cannot retrieve LOOKUP_ERRORS_SQL from environment variables")
		return nil
	}
	atts := tables.LUT_cad_errors{}
	copyKeyval := make(map[string]interface{})
	for errorName, errorval := range keyval {
		testval := parsInt(errorval)
		if testval == 0 {
			db.Where(sql, errorName).Find(&atts)
			log.Logger.Log.Info(atts)
			copyKeyval[errorName] = atts.CombinedKey
		}
	}
	return copyKeyval
}

func parsInt(val interface{}) int {
	var testval int
	if reflect.TypeOf(val).Kind() == reflect.Float64 {
		testval = int(val.(float64))
	}
	if reflect.TypeOf(val).Kind() == reflect.Int {
		testval = val.(int)
	}
	return testval
}

func checkIfExist(db *CDb, id int, errorCode int) bool {
	atts := tables.CAD_check_errors{}
	row, err := db.Where(&tables.CAD_check_errors{Check_Status_Id: id, Error_Code: errorCode}).First(&atts).Rows()
	if err != nil {
		return false
	}
	if row.Next() {
		return true
	}
	return false

}

func ErrorsCreate(db *CDb, FkId map[string]interface{}, keyval map[string]interface{}) ([]byte, error) {
	keyval = Lut_Error_Retrieve(db, keyval)
	for _, errorCode := range keyval {
		checkId := parsInt(FkId["check_status_id"])
		errVal := parsInt(errorCode)
		if !checkIfExist(db, checkId, errVal) {
			atts := tables.CAD_check_errors{}
			atts.Check_Status_Id = checkId
			atts.Error_Code = errVal
			_, err := Create(atts, db)
			HandleDBErrors([]error{err})
		}
	}
	return []byte{}, nil
}
