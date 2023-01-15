package globle

import (
	"github.com/ronnielin8862/go-practice/cmd/db/postgres/psqlConn"
	"github.com/ronnielin8862/go-practice/config"
)

func StartServer() {
	//go func() {
	//	log.Println(http.ListenAndServe("localhost:8211", nil))
	//}()

	config.LoadGlobalConfig()

	psqlConn.InitPsql()

}
