package main

import (
	"database/sql"
	"github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./db/db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("insert into logs(idLog, cot) values(?,?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec("123", "teste")
	if err != nil {
		panic(err)
	}
	

}
