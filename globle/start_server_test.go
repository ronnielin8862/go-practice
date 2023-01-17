package globle

import (
	"fmt"
	"github.com/ronnielin8862/go-practice/cmd/db/postgres/psql"
	"testing"
)

var (
	id   int
	name string
	age  int
)

func TestStartServer(t *testing.T) {
	StartServer()
	rows, err := psql.Get().Query("SELECT * FROM users")
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

// 這裡的測試結果，fx在啟動server時就把所有服務先啟動了，並不是調用到才會實例化
// 這樣在unitest階段並沒有傳說中的優勢... 反而自己手動選擇要啟動的服務，更加靈活
func TestStartServerDI(t *testing.T) {
	fmt.Println("========= before start server")
	StartServer2()
	fmt.Println("========= invoke psql")

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
