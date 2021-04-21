package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var MysqlDB *sqlx.DB

func init() {
	var err error

	MysqlDB, err = sqlx.Connect("mysql", DATA_SOURCE)

	if err != nil {
		fmt.Println(err)
		return
	}
}
