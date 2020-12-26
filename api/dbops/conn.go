package dbops

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err error
)

func init() {
	//Open connection
	dbConn, err = sql.Open("mysql", "root:SYDsdtc1203@tcp(localhost:3306)/video_server?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}