package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type databaseMySQL struct {
	Mysql *sql.DB
}

func ConnnectMySql() (*databaseMySQL, error) {
	mySql, err := sql.Open("mysql", "root:admin123@tcp(mysql-container:3306)/mydb")

	if err != nil {
		return nil, err
	}
	if err = mySql.Ping(); err != nil {
		return nil, err
	}
	return &databaseMySQL{Mysql: mySql}, nil
}
