package infra

import "database/sql"

func NewMysqlDatabase(dsn string) (*sql.DB, error) {

	return sql.Open("mysql", dsn)
}
