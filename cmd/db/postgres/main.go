package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1qaz2wsx"
	dbname   = "chelsea"
)

var db *sql.DB

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	newDb, err := sql.Open("postgres", psqlInfo)

	err = newDb.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	db = newDb
}

var (
	id   int
	name string
	age  int
)

func main() {

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			return
		}
		fmt.Println(id, name, age)
	}

	db.Close()
}
