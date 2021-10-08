package database

import "fmt"

type Config struct {
	User,
	Password,
	SSLMode,
	DBName string
}

func (c Config) ConnectionString() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", c.User, c.Password, c.DBName, c.SSLMode)
}
