package database

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var connection *sqlx.DB

func init() {
	var err error
	dbUser, dbPassword, dbName :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB")
	config := Config{
		User:     dbUser,
		Password: dbPassword,
		SSLMode:  "disable",
		DBName:   dbName,
	}
	fmt.Println(config.ConnectionString())
	connection, err = sqlx.Connect("postgres", config.ConnectionString())
	if err != nil {
		panic(err)
	}
}

func Connection() *sqlx.DB {
	return connection
}
