package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Dbconn *sql.DB
)

func InitDatabase() {
	var err error
	Dbconn, err = sql.Open("mysql", "root:example@tcp(localhost:3306)/bookDb")
	if err != nil {
		panic("Failed to connect to a Database")
	}
	fmt.Println("database connection successfully opened.")
}
