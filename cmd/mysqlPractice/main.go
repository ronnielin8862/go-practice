package main

import (
	"database/sql"
	"fmt"

	//"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:my-password@/test?charset=utf8")
	checkErr(err, db)

	//插入資料
	stmt, err := db.Prepare("INSERT INTO EMP (EMPNO , ENAME, JOB ,deptno ) VALUE (?,?,?,?)")
	checkErr(err, db)

	res, err := stmt.Exec("1111", "測試人員1", "Engineer", 20)
	checkErr(err, db)

	id, err := res.LastInsertId()
	checkErr(err, db)

	fmt.Println(id)

	affect, err := res.RowsAffected()
	checkErr(err, db)

	fmt.Println(affect)

	defer func() {
		db.Close()
		fmt.Print("關閉36")
	}()
}

func checkErr(err error, db *sql.DB) {
	if err != nil {
		panic(err)
	}

	defer func() {
		db.Close()
		fmt.Print("關閉47")
	}()
}
