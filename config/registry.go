package config

import (
	"github.com/jossefaz/dwg-transformer-micro-dal/model"
)

var SchemaReg = map[string]map[string]model.Schema{}

func initReg() {
	SchemaReg["mysql"] = map[string]model.Schema{
		"dwg_transformer": LocalConfig.DB.Mysql.Schema["dwg_transformer"],
	}
	SchemaReg["mssql"] = map[string]model.Schema{
		"dbo": LocalConfig.DB.Mssql.Schema["dbo"],
	}

}
func GetDBConf(dbtype string, schema string) model.Schema {
	return SchemaReg[dbtype][schema]
}
