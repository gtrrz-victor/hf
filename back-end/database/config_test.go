package database

import "testing"

func TestConnectionString(t *testing.T) {
	config := Config{
		User:     "user",
		Password: "pass",
		SSLMode:  "disable",
		DBName:   "dbname",
	}
	connectionStringExpected := "user=" + config.User + " password=" + config.Password + " dbname=" + config.DBName + " sslmode=" + config.SSLMode
	connectionStringActual := config.ConnectionString()
	if connectionStringActual != connectionStringExpected {
		t.Errorf("config.ConnectionString() = %s; want %s", connectionStringActual, connectionStringExpected)
	}

}
