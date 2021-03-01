package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var MysqlDB *sqlx.DB

func init() {
	var err error

	//MysqlDB, err = sqlx.Connect("mysql", "b9b59bf12bb030:3e741152@tcp(eu-cdbr-west-03.cleardb.net:3306)/heroku_e9281e77bbf34e2")
	//local:
	MysqlDB, err = sqlx.Connect("mysql", "root:@tcp(localhost:3306)/sport_kld_db")

	if err != nil {
		fmt.Println(err)
		return
	}
}
