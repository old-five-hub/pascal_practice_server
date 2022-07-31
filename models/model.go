package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func InitDb() {
	var err error
	Db, err = sql.Open("mysql", "root:yR.3Ad*P9A3VaZiFNDmst9mN@tcp(127.0.0.1:3306)/pascal_practice")
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)
}
