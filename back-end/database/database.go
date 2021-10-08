package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var connection *sqlx.DB

func init() {
	var err error
	config := Config{
		User:     "postgres",
		Password: "test123",
		SSLMode:  "disable",
		DBName:   "postgres",
	}
	connection, err = sqlx.Connect("postgres", config.ConnectionString())
	if err != nil {
		panic(err)
	}
}

func Connection() *sqlx.DB {
	return connection
}
