package main

import (
	"database/sql"
	"fmt"

	//"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	EMPNO, ENAME, JOB, DEPTNO := 555, "第5個", "NO5", 5
	// var EMPNO int16 = 333
	// var ENAME string = "第3個"
	// var JOB string = "NO3"
	// var DEPTNO int16 = 3

	InsertEMP(&EMPNO, ENAME, JOB, DEPTNO)

}

func InsertEMP(EMPNO *int, ENAME string, JOB string, DEPTNO int) {
	db, err := sql.Open("mysql", "root:my-password@/test")
	checkErr(err, db)

	//插入資料
	stmt, err := db.Prepare("INSERT INTO EMP (EMPNO , ENAME, JOB ,DEPTNO ) VALUE (?,?,?,?)")
	if err != nil {
		checkErr(err, db)
	}

	res, err := stmt.Exec(*EMPNO, ENAME, JOB, DEPTNO)
	checkErr(err, db)

	id, err := res.LastInsertId()
	checkErr(err, db)

	fmt.Println(id)

	affect, err := res.RowsAffected()
	checkErr(err, db)

	fmt.Println(affect)

	db.Close()
	fmt.Print("關閉36")
}

func checkErr(err error, db *sql.DB) {
	if err != nil {

		defer func() {
			db.Close()
			fmt.Print("關閉47")
		}()

		fmt.Print(err.Error())
		panic(err)
	}

}
