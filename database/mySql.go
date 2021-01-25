package database

import (
	"github.com/jmoiron/sqlx"
)

var MysqlDB *sqlx.DB

func init() {
	var  err error
	MysqlDB, err = sqlx.Connect("mysql", "")
	if err != nil {
		return
	}

}

