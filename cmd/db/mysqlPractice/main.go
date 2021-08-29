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
	// 設置日誌格式為test格式
	log.SetFormatter(&logrus.TextFormatter{})

	// 設置日誌標準輸出（默认的输出为stderr,标准错误）
	log.SetOutput(os.Stdout)

	// 設置級別
	log.SetLevel(logrus.DebugLevel)
}

type Emp struct {
	Empno  uint16
	Ename  string
	Job    string
	Deptno uint16
}

type EmpTest struct {
	Empno  uint16
	Ename  string
	Job    string
	Deptno uint16
}

func main() {

	f, err := os.OpenFile("logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("file open error : %v", err)
	}

	defer f.Close()
	log.SetOutput(f)

	log.Info("開始測試insert DB")

	// EMPNO, ENAME, JOB, DEPTNO := 111, "第1個", "NO1", 1
	var emp Emp

	empData := emp.DataProcessing()

	log.WithFields(logrus.Fields{"傳入內容 EMPNO ": empData.Empno,
		"ENAME ":  empData.Ename,
		"JOB ":    empData.Job,
		"DEPTNO ": empData.Deptno,
	}).Debug("不錯喔")

	//也可以這樣直接使用，但是遇到需要呈現某些值的時候，推薦withfield
	log.Debug("測試不使用WithFields")

	InsertEMP(&empData.Empno, empData.Ename, empData.Job, empData.Deptno)

}

func (emp *Emp) DataProcessing() (empData *Emp) {

	emp.Empno = 111
	emp.Ename = "第1個"
	emp.Job = "No1"
	emp.Deptno = 111

	test := new(EmpTest)

	test.Deptno = 111

	// empTest := new(empTest)

	return emp

}

func InsertEMP(EMPNO *uint16, ENAME string, JOB string, DEPTNO uint16) {
	db, err := sql.Open("mysql", "root:my-password@/test") //user:password@/dbname
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
