package config

import (
	     "database/sql"
       _ "github.com/lib/pq"
		"fmt"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres",
		"postgres://postgres:Yfehsp2203@localhost/postgres?sslmode=disable")

	if err != nil {
		panic("Can not connect to DB")
	}

	if err = Db.Ping(); err != nil {
		panic("Pingin error")
	}

	fmt.Println("Successfully connected")
}
