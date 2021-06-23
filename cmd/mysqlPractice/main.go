package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	// 设置日志格式为json格式
	log.SetFormatter(&logrus.TextFormatter{})

	// 设置将日志输出到标准输出（默认的输出为stderr,标准错误）
	// 日志消息输出可以是任意的io.writer类型
	log.SetOutput(os.Stdout)

	// 设置日志级别为warn以上
	log.SetLevel(logrus.DebugLevel)
}

func main() {

	log.Info("開始測試insert DB")

	EMPNO, ENAME, JOB, DEPTNO := 111, "第1個", "NO1", 1

	log.WithFields(logrus.Fields{"傳入內容 EMPNO ": EMPNO,
		"ENAME ":  ENAME,
		"JOB ":    JOB,
		"DEPTNO ": DEPTNO,
	}).Debug("不錯喔")

	log.Debug("測試不使用WithFields")

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
	fmt.Print("正常關閉")
}

func checkErr(err error, db *sql.DB) {
	if err != nil {

		defer func() {
			db.Close()
			fmt.Print("異常關閉")
		}()

		fmt.Print(err.Error())
		panic(err)
	}

}
