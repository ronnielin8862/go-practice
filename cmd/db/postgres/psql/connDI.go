package psql

import (
	"database/sql"
	"fmt"
	"github.com/ronnielin8862/go-practice/config"

	_ "github.com/lib/pq"
)

//var db2 *sql.DB

func InitPsql2(psqlConfig *config.GlobalConfig2) *sql.DB {

	fmt.Println("==== InitPsql2 ==== ", psqlConfig)

	psql := psqlConfig.Psql2

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

	return newDb
}
