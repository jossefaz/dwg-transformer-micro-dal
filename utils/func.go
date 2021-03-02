package utils

import (
	"github.com/jossefaz/dwg-transformer-micro-dal/config"
	"github.com/jossefaz/dwg-transformer-micro-dal/log"
	"github.com/jossefaz/dwg-transformer-micro-dal/model"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/streadway/amqp"
	"github.com/jossefaz/dwg-transformer-micro-utils/queue"
	globalUtils "github.com/jossefaz/dwg-transformer-micro-utils/utils"
)

func HandleError(err error, msg string, exit bool) {
	if err != nil {
		log.Logger.Log.Error(fmt.Sprintf("%s: %s", msg, err))
	}
	if exit {
		os.Exit(1)
	}
}

func MessageReceiver(m amqp.Delivery, rmq *queue.Rabbitmq) {
	dbQ, err := unpackMessage(m)
	HandleError(err, "cannot unpack message received in DAL to DB", err != nil)
	dbconf := config.GetDBConf(dbQ.DbType, dbQ.Schema)
	cdb, err := model.ConnectToDb(dbconf.Dialect, dbconf.ConnString)
	if cdb != nil {
		connection, err := cdb.DB.DB()
		HandleError(err, "cannot connect to DB", err != nil)
		defer func(){
			if err := connection.Close(); err !=nil {
				HandleError(err, "Cannot close DB connection", true)
			}
		}()
		res, err := dispatcher(cdb, dbQ)
		if err != nil {
			log.Logger.Log.Error(err)
		} else {
			message, err:= rmq.SendMessage(res, "Dal_Res", map[string]interface{}{
				"From": "DAL",
				"To":   "Dal_Res",
				"Type": dbQ.CrudT,
			})
			if err != nil {
				HandleError(err, "Cannot send message to the Queue", false)
			}
			log.Logger.Log.Info(fmt.Sprintf("%s", message))
		}
	} else {
		log.Logger.Log.Error(fmt.Sprintf("Cannot open connection with database check connection string or network error : %s", err.Error()))
	}

}

func dispatcher(db *model.CDb, dbQ *globalUtils.DbQuery) ([]byte, error) {
	switch dbQ.CrudT {
	case "retrieve":
		res, err := db.RetrieveRow(dbQ)
		if err != nil {
			return nil, err
		}
		return res, nil
	case "update":
		res, err := db.UpdateRow(dbQ)
		if err != nil {
			return nil, err
		}
		return res, nil
	case "create":
		res, err := db.CreateRow(dbQ)
		if err != nil {
			return nil, err
		}
		return res, nil

	default:
		return nil, errors.New("CRUD operation must be one of the following : retrieve, update | delete and create not supported yet")
	}

}

func unpackMessage(m amqp.Delivery) (*globalUtils.DbQuery, error) {
	dbQ := &globalUtils.DbQuery{}
	err := json.Unmarshal(m.Body, dbQ)
	if err := m.Ack(false); err != nil {
		log.Logger.Log.Error("Error acknowledging message : %s", err)
		return &globalUtils.DbQuery{}, err
	}
	HandleError(err, "Error decoding DB message", false)
	return dbQ, nil
}

//"mysql", "root:Dev123456!@(localhost)/dwg_transformer?charset=utf8&parseTime=True&loc=Local"
