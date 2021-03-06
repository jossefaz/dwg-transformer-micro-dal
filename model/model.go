package model

import (
	"errors"
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	tables "github.com/jossefaz/dwg-transformer-micro-data-struct"
	globalUtils "github.com/jossefaz/dwg-transformer-micro-utils/utils"
)

type CDb struct {
	*gorm.DB
}

type Schema struct {
	ConnString string
	Name       string
	Dialect    string
}

type Timestamp time.Time

func HandleDBErrors(errs []error) error {
	if len(errs) > 0 {
		var b1 strings.Builder
		for _, err := range errs {
			b1.WriteString(fmt.Sprintln(err))
		}
		return errors.New(b1.String())
	}
	return nil
}

func ConnectToDb(dialect string, connString string) (*CDb, error) {
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		return &CDb{}, err
	}
	dup := CDb{db}
	return &dup, nil
}

func (db *CDb) RetrieveRow(dbQ *globalUtils.DbQuery) ([]byte, error) {

	switch dbQ.Table {
	case "CAD_check_status":
		status := []tables.Cad_check_status{}
		res, err := Retrieve(&status, db, dbQ.ORMKeyVal)
		if err != nil {
			return nil, err
		}
		return res, nil
	case "CAD_check_errors":
		errors := []tables.CAD_check_errors{}
		res, err := Retrieve(&errors, db, dbQ.ORMKeyVal)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	return []byte{}, errors.New("any tables allowed correspond to the requested table name")
}

func (db *CDb) UpdateRow(dbQ *globalUtils.DbQuery) ([]byte, error) {
	switch dbQ.Table {
	case "CAD_check_status":
		res, err := StatusUpdate(db, dbQ.Id, dbQ.ORMKeyVal)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	return []byte{}, errors.New("any tables allowed correspond to the requested table name")
}

func (db *CDb) CreateRow(dbQ *globalUtils.DbQuery) ([]byte, error) {
	switch dbQ.Table {
	case "CAD_check_errors":
		res, err := ErrorsCreate(db, dbQ.Id, dbQ.ORMKeyVal)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	return []byte{}, errors.New("any tables allowed correspond to the requested table name")
}
