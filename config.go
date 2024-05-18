package main

import (
	"database/sql"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var rdb *redis.Client

func Connnect() error {
	var err error
	db, err = sql.Open("mysql", "root:admin123@/mydb")
	if err != nil {
		return err
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return nil
}
