package globle

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ronnielin8862/go-practice/cmd/db/postgres/psqlConn"
	"github.com/ronnielin8862/go-practice/config"
	"go.uber.org/fx"
)

var config2 *config.GlobalConfig2
var psql2 *sql.DB

func StartServer2() {
	//go func() {
	//	log.Println(http.ListenAndServe("localhost:8211", nil))
	//}()
	//var reader io.Reader
	app := fx.New(
		fx.Provide(config.LoadGlobalConfig2),
		fx.Populate(&config2),
		fx.Provide(psqlConn.InitPsql2),
		fx.Populate(&psql2),
	)

	if err := app.Start(context.Background()); err != nil {
		fmt.Println(err)
	}
	//defer app.Stop(context.Background())
	fmt.Println("config = ", config2)
	fmt.Println("psql = ", psql2)

}
