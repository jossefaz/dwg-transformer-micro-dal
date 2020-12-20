package config

import (
	"github.com/jossefaz/dwg-transformer-micro-dal/model"
	"fmt"
	"os"

	"github.com/jossefaz/dwg-transformer-micro-utils/utils"
	"github.com/tkanos/gonfig"
)

var LocalConfig Configuration

type Configuration struct {
	Queue struct {
		Rabbitmq struct {
			ConnString string   `json:"ConnString"`
			QueueNames []string `json:"QueueNames"`
			Listennig  []string `json:"Listennig"`
			Result     utils.Result
		} `json:"Rabbitmq"`
	} `json:"Queue"`
	DB struct {
		Mysql struct {
			Schema map[string]model.Schema
		} `json:"Mysql"`
		Mssql struct {
			Schema map[string]model.Schema
		} `json:"Mysql"`
	} `json:"DB"`
}

var configEnv = map[string]string{
	"dev":  "config/config.dev.json",
	"prod": "config/config.prod.json",
}

func GetConfig(env string) {
	configuration := Configuration{}
	fmt.Println(os.Getwd())
	err := gonfig.GetConf(configEnv[env], &configuration)
	if err != nil {
		fmt.Println("Cannot read config file")
	}
	LocalConfig = configuration
	initReg()
}
