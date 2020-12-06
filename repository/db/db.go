package db

import (
	"database/sql"
	"fmt"

	"wire_sample/env"
)

func OpenDB(e env.Conf) *sql.DB {
	dbPath := fmt.Sprintf("%s:%s@%s", e.DBSetting.USER, e.DBSetting.PASS, e.DBSetting.HOST)
	db, _ := sql.Open("mysql", dbPath)
	return db
}
