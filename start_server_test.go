package main

import (
	"fmt"
	"github.com/ronnielin8862/go-practice/cmd/db/postgres/psqlConn"
	"github.com/ronnielin8862/go-practice/globle"
	"testing"
)

var (
	id   int
	name string
	age  int
)

func TestStartServer(t *testing.T) {
	globle.StartServer()
	rows, err := psqlConn.GetPsql().Query("SELECT * FROM users")
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
}
