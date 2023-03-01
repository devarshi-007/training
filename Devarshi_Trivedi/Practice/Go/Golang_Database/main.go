package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// create a database object which can be used
	// to connect with database.
	db, err := sql.Open("mysql", "root:passwd@tcp(0.0.0.0:3306)/user")

	// handle error, if any.
	if err != nil {
		panic(err)
	}

	// database object has a method called Exec,
	// it executes a database query, but it does
	// not return any row as result.
	// Here we create a database table with a SQL query.
	_, err = db.Exec("CREATE TABLE user(id INT NOT NULL, name VARCHAR(20),PRIMARY KEY (ID));")

	// handle error
	if err != nil {
		panic(err)
	}

	fmt.Print("Successfully Created\n")

	// database object has a method Close,
	// which is used to free the resource.
	// Free the resource when the function
	// is returned.
	defer db.Close()
}
