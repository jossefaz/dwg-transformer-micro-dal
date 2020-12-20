package main

import (
	"github.com/jossefaz/dwg-transformer-micro-dal/config"
	"github.com/jossefaz/dwg-transformer-micro-dal/log"
	"github.com/jossefaz/dwg-transformer-micro-dal/utils"
	"fmt"
	"os"

	"github.com/jossefaz/dwg-transformer-micro-utils/queue"
	globalUtils "github.com/jossefaz/dwg-transformer-micro-utils/utils"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	environment, err := globalUtils.GetEnv("DEV_PROD")
	if environment == "" {
		fmt.Printf("Environment variable %s must be set in order to make this service work", "DEV_PROD")
		os.Exit(1)
	}
	utils.HandleError(err, "Error while getting env variable", err != nil)
	config.GetConfig(environment)
	log.GetLogger(environment)
}

func main() {

	queueConf := config.LocalConfig.Queue.Rabbitmq
	rmqConn, err := queue.NewRabbit(queueConf.ConnString, queueConf.QueueNames)
	utils.HandleError(err, "Error Occured when RabbitMQ Init", err != nil)
	defer rmqConn.Conn.Close()
	defer rmqConn.ChanL.Close()
	rmqConn.OpenListening(queueConf.Listennig, utils.MessageReceiver)

}
