package model

import (
	"testing"

	"github.com/jossefaz/dwg-transformer-micro-utils/test"
	globalUtils "github.com/jossefaz/dwg-transformer-micro-utils/utils"
)

var datbaseQuery = globalUtils.DbQuery{
	DbType: "mysql",
	Id: map[string]interface{}{
		"Id": 5,
	},
	Table: "cad_check_status",
	ORMKeyVal: map[string]interface{}{
		"status_code": 0,
	},
}

var datbaseErrorQuery = globalUtils.DbQuery{
	DbType: "mysql",
	Schema: "dwg_transformer",
	CrudT:  "create",
	Table:  "cad_check_errors",
	Id: map[string]interface{}{
		"check_status_id": 6,
	},
	ORMKeyVal: map[string]interface{}{
		"BorderExist": 1,
		"InsideJer":   0,
	},
}

func TestConnectToDb(t *testing.T) {
	_, err := ConnectToDb("mysql", "root:Dev123456!@(localhost)/dwg_transformer?charset=utf8&parseTime=True&loc=Local")
	test.Ok(t, err)
}

func TestConnectToDbMSServer(t *testing.T) {
	_, err := ConnectToDb("mssql", "root:Dev123456!@(localhost)/dwg_transformer?charset=utf8&parseTime=True&loc=Local")
	test.Ok(t, err)
}

func TestCDb_Retrieve(t *testing.T) {
	cdb, err := ConnectToDb("mysql", "root:Dev123456!@(localhost)/dwg_transformer?charset=utf8&parseTime=True&loc=Local")
	test.Ok(t, err)
	_, err1 := cdb.RetrieveRow(&datbaseQuery)
	test.Ok(t, err1)
}

func TestErrorsCreate(t *testing.T) {
	cdb, err := ConnectToDb("mysql", "root:Dev123456!@(localhost)/dwg_transformer?charset=utf8&parseTime=True&loc=Local")
	test.Ok(t, err)
	_, err1 := cdb.CreateRow(&datbaseErrorQuery)
	test.Ok(t, err1)
}

func TestCDb_Update(t *testing.T) {
	cdb, err := ConnectToDb("mysql", "root:Dev123456!@(localhost)/dwg_transformer?charset=utf8&parseTime=True&loc=Local")
	test.Ok(t, err)
	_, err1 := cdb.UpdateRow(&datbaseQuery)
	test.Ok(t, err1)
}
