package psql

import (
	"database/sql"
	"fmt"
	"github.com/ronnielin8862/go-practice/config"
	"sync"

	_ "github.com/lib/pq"
)

var db *sql.DB
var once sync.Once

func initPsql() {

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

func Get() *sql.DB {
	once.Do(initPsql)
	return db
}
