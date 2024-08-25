package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

const (
	maxOpenDbConn = 25
	maxIdleDBConn = 25
	maxDBLifetime = 5 * time.Minute
)

func newDB(driver, dsn string) (*sql.DB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetMaxIdleConns(maxIdleDBConn)
	db.SetConnMaxLifetime(maxDBLifetime)

	return db, nil
}
