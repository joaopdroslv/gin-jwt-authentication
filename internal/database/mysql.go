package database

import "database/sql"

func NewMysql(dsn string) (*sql.DB, error) {

	return sql.Open("mysql", dsn)
}
