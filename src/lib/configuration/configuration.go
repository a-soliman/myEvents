package configuration

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/a-soliman/projects/myEvents/src/lib/persistence/dblayer"
)

// DBTypeDefault the default type for db
var DBTypeDefault = dblayer.DBTYPE("mongodb")

// DBConnectionDefault default
var DBConnectionDefault = "mongodb://127.0.0.1"

// RestfulEPDefault default
var RestfulEPDefault = "localhost:8181"

// ServiceConfig struct
type ServiceConfig struct {
	Databasetype    dblayer.DBTYPE `json:"databasetype"`
	DBConnection    string         `json:"dbconnection"`
	RestfulEndpoint string         `json:"restfulapi_endpoint"`
}

// ExtractConfiguration extract the configuration file
func ExtractConfiguration(filename string) (ServiceConfig, error) {
	conf := ServiceConfig{
		DBTypeDefault,
		DBConnectionDefault,
		RestfulEPDefault,
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Configuration file not found. Continuing with default values.")
		return conf, err
	}

	err = json.NewDecoder(file).Decode(&conf)
	return conf, err
}
