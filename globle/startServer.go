package globle

import (
	"github.com/ronnielin8862/go-practice/cmd/db/postgres/psqlConn"
	"github.com/ronnielin8862/go-practice/config"
	"log"
)

func StartServer() {
	//go func() {
	//	log.Println(http.ListenAndServe("localhost:8211", nil))
	//}()

	cfg, err := config.LoadGlobalConfig("/Users/ronnie/Library/Mobile Documents/com~apple~CloudDocs/Documents/coding/code/go-practice/config.toml")
	if err != nil {
		log.Fatal("init config err :", err)
		return
	}

	psqlConn.InitPsql(cfg)

}
