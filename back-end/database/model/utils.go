package model

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func isTableCreated(db *sqlx.DB, tableName string) bool {
	_, table_check := db.Query(fmt.Sprintf("select * from %s;", tableName))
	return table_check == nil
}
