package psqlConn

import (
	"database/sql"
	"fmt"
	"github.com/ronnielin8862/go-practice/config"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitPsql() {

	psql := config.Config.Psql

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		psql.Host, psql.Port, psql.User, psql.Password, psql.DBName)

	newDb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Failed to connect to psql : ", err)
		panic(err)
	}

	err = newDb.Ping()

	fmt.Println("Successfully connected!")

	db = newDb
}

func GetPsql() *sql.DB {
	return db
}
