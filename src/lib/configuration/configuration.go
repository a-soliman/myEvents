package configuration

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/a-soliman/projects/myEvents/src/lib/persistence/dblayer"
)

// Default values
var (
	DBTypeDefault            = dblayer.DBTYPE("mongodb")
	DBConnectionDefault      = "mongodb://127.0.0.1"
	RestfulEPDefault         = "localhost:8181"
	RestfulTLSEPDefault      = "localhost:9191"
	AMQPMessageBrokerDefault = "amqp://guest:guest@localhost:5672"
)

// ServiceConfig struct
type ServiceConfig struct {
	Databasetype       dblayer.DBTYPE `json:"databasetype"`
	DBConnection       string         `json:"dbconnection"`
	RestfulEndpoint    string         `json:"restfulapi_endpoint"`
	RestfulTLSEndpoint string         `json:"restfulapi_tlsendpoint"`
	AMQPMessageBroker  string         `json:"amqp_message_broker"`
}

// ExtractConfiguration extract the configuration file
func ExtractConfiguration(filename string) (ServiceConfig, error) {
	conf := ServiceConfig{
		DBTypeDefault,
		DBConnectionDefault,
		RestfulEPDefault,
		RestfulTLSEPDefault,
		AMQPMessageBrokerDefault,
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Configuration file not found. Continuing with default values.")
		return conf, err
	}

	err = json.NewDecoder(file).Decode(&conf)
	if broker := os.Getenv("AMQP_URL"); broker != "" {
		conf.AMQPMessageBroker = broker
	}
	return conf, err
}
