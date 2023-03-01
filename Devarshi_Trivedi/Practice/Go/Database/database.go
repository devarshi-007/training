package main

import (
	"database/sql"
	"log"
)

var Dbconn *sql.DB

func SetupDatabase() {
	var err error
	Dbconn, err := sql.Open("mysql", "root:example@tcp(localhost:3306)/userDb")
	if err != nil {
		log.Fatal(err)
	}
	Dbconn.SetMaxOpenConns(3)
}
