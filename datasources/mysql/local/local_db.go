package local

import (
	"database/sql"
	"fmt"
	"github.com/Mikyas1/SCADA_Go-local-sql/app/env"

	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
)

// Open database connection
func Open() (*sql.DB, *error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		env.DbUserName,
		env.DbPassword,
		env.DbHost,
		env.DbName)

	Client, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		color.Red(fmt.Sprintf("SQL ERROR: %s", err))
		return nil, &err
	}
	if err = Client.Ping(); err != nil {
		color.Red(fmt.Sprintf("SQL ERROR: %s", err))
		return nil, &err
	}

	color.Green("Successfully Connected to Local database.")
	return Client, nil
}
