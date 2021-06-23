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
	// 設置日誌格式為json格式, or  test格式
	log.SetFormatter(&logrus.TextFormatter{})

	// 设置将日志输出到标准输出（默认的输出为stderr,标准错误）
	// 日志消息输出可以是任意的io.writer类型
	log.SetOutput(os.Stdout)

	// 設置級別
	log.SetLevel(logrus.DebugLevel)
}

func main() {

	f, err := os.OpenFile("logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("file open error : %v", err)
	}

	defer f.Close()
	log.SetOutput(f)

	log.Info("開始測試insert DB")

	EMPNO, ENAME, JOB, DEPTNO := 111, "第1個", "NO1", 1

	log.WithFields(logrus.Fields{"傳入內容 EMPNO ": EMPNO,
		"ENAME ":  ENAME,
		"JOB ":    JOB,
		"DEPTNO ": DEPTNO,
	}).Debug("不錯喔")

	//也可以這樣直接使用，但是遇到需要呈現某些值的時候，推薦withfield
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
