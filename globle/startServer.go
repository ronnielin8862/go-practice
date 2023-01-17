package globle

import (
	"github.com/ronnielin8862/go-practice/cmd/db/postgres/psql"
	"github.com/ronnielin8862/go-practice/config"
)

func StartServer() {
	//go func() {
	//	log.Println(http.ListenAndServe("localhost:8211", nil))
	//}()

	config.InitConfig()

	psql.Get()

}
