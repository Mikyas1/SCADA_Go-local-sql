package local

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	Client   *sql.DB
	username = "root"
	password = "Engineering"
	host     = "localhost"
	schema   = "SCADA_MASTER"
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		schema,
	)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("Local database successfully Connected.")

}