package database

import (
	"database/sql"
)

type DbConfig struct {
	Host     string
	Port     string
	UserName string
	Password string
}

func (config *DbConfig) GrtConnection(dbName string) *sql.DB {
	return nil

}
