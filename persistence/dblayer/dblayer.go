package dblayer

import (
	"aws-api/persistence"
	"aws-api/persistence/mysql-layer"
)

type DBTYPE string

const (
	MYSQLDB  DBTYPE = "mysqldb"
	DYNAMODB DBTYPE = "dynamodb"
)

func NewPersistenceLayer(options DBTYPE, connection string) (persistence.DatabaseHandler, error) {

	switch options {
	case MYSQLDB:
		return mysql_layer.NewMySqlLayer(connection)
	}
	return nil, nil
}
