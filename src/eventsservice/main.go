package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/a-soliman/projects/myEvents/src/eventsservice/rest"
	"github.com/a-soliman/projects/myEvents/src/lib/configuration"
	"github.com/a-soliman/projects/myEvents/src/lib/persistence/dblayer"
)

func main() {
	confPath := flag.String("conf", `./configuration/config.json`, "flag to set the path to the configuration json file")
	flag.Parse()
	// extract configuration
	config, _ := configuration.ExtractConfiguration(*confPath)
	fmt.Println("Connection to Database")
	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	// Start restful API
	httpErrChan, httpTLSErrChan := rest.ServeAPI(config.RestfulEndpoint, config.RestfulTLSEndpoint, dbhandler)
	select {
	case err := <-httpErrChan:
		log.Fatal("HTTP Error: ", err)
	case err := <-httpTLSErrChan:
		log.Fatal("HTTPS Error: ", err)
	}
}
