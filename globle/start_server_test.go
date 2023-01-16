package globle

import (
	"fmt"
	"github.com/ronnielin8862/go-practice/cmd/db/postgres/psqlConn"
	"testing"
)

var (
	id   int
	name string
	age  int
)

func TestStartServer(t *testing.T) {
	StartServer()
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

func TestStartServerDI(t *testing.T) {
	StartServer2()

	rows, err := psql2.Query("SELECT * FROM users")
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
