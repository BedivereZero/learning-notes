package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:0000000000000000@tcp(0.0.0.0:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("select name from users")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var name string
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		rows.Scan(&name)
		log.Println(name)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
