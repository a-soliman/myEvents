package dblayer

import (
	"github.com/a-soliman/projects/myEvents/src/lib/persistence"
	"github.com/a-soliman/projects/myEvents/src/lib/persistence/mongolayer"
)

// DBTYPE a string based type represents the db type
type DBTYPE string

const (
	// MONGODB mongodb
	MONGODB DBTYPE = "mongodb"
	// DYNAMODB dynamodb
	DYNAMODB DBTYPE = "dynamodb"
)

// NewPersistenceLayer returns dbHandler
func NewPersistenceLayer(options DBTYPE, connection string) (persistence.DatabaseHandler, error) {
	switch options {
	case MONGODB:
		return mongolayer.NewMongoDBLayer(connection)
	}
	return nil, nil
}
