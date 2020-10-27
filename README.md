# Dwg-transformer : DAL

This module is a part of the dwg-transformer-microservice tool. It's called the **DAL** because he is responsible for accessing the database (in our hybrid architecture for the moment there is only one DB because other microservices do not require any storing, they just need metadata to indicates which files they have to convert/deal with)


# Configuration

Only one environment variables must be set before running this program :

**$DEV_PROD** :  define the environment. Must be set to one of those values : "prod" / "dev", depending on the configuration files that you rely on.

Based on the **$DEV_PROD** variable, there is two configuration files. 

 1. config.prod.json (for production)
 2. config.dev.json (for dev)
 
 **Warning** : all configuration files must be set in the folder “config”. The relative path of the file would be then for example :

    ROOT_FOLDER/config/config.prod.json

 Here is an example of the configuration file. 

```
{  
  "Queue" : {  
    "Rabbitmq" : {  
      "ConnString" : "amqp://guest:guest@localhost:5672connection_attempts=5&retry_delay=5",  
	  "QueueNames" : [ "Dal_Req", "Dal_Res"],  
	  "Listennig" : ["Dal_Req"],  
	  "Result" : {  
	        "Success" : "Dal_Res",  
			"Fail" : "Dal_Res",  
			"From" : "DAL"  
	  }  
    }  
  },  
  "DB" :  {  
    "Mysql" : {  
      "Schema" : {  
        "dwg_transformer" : {  
          "ConnString" :  "root:password@(localhost:3306)/dwg_transformer?charset=utf8&parseTime=True&loc=Local",  
		"Name" : "dwg_transformer",  
		"Dialect" : "mysql"  
		 }  
      }  
    }  
  }  
}
```
 **Queue** : define which message broker will be used and how to connect to it. Here is an example with RabbitMQ connection string to localhost.
---> **ConnString** : the connection string to the queue
---> **QueueNames** : the queues that this service will open ( if not exists) an publish to them
---> **Listennig** : the queues that this service will subscribe and listen for messages

## Dependencies

This service, as the orchestrator, is dependent on at least 4 others services to run correctly, an do the job.

 1. **DAL** service : implement access to the dwg database.
 2. **transformer** service : responsible to convert the DWG files to DXF or others format
 3. **worker** service : will perform validity check on the converted file
 4. **Message Broker** service  : responsible for the communication between microservices


## Flow

The DAL workflow is very simple :

```mermaid
sequenceDiagram
Message Broker ->> DAL:  DAL request
DAL-->>DB: Send SQL request
DB-->> DAL: Result Buffer
DAL-->> Message Broker: DAL Response


