package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB() *sql.DB {
	con, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/alura_store")
	if err != nil {
		panic(err.Error())
	}
	return con
}
