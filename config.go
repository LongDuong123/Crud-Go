package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connnect() error {
	var err error
	db, err = sql.Open("mysql", "root:admin123@/mydb")
	if err != nil {
		return err
	}
	return nil
}
