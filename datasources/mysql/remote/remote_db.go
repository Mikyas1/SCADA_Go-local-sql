package remote

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

var (
	connectionURL [100]string
	servers       []Server
)

func init() {
	conf := configure("config.yml")
	setup(conf.Databases)
}

func configure(fileName string) Config {
	var config Config
	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("File not found: %v", err))
	}
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		panic(fmt.Sprintf("Read yaml: %v", err))
	}
	return config
}

// Setup database
func setup(databases []Server) {
	//"root:filling@tcp(192.168.100.47:3306)/ccc"
	for i := 0; i < len(databases); i++ {
		connectionURL[i] = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", databases[i].User, databases[i].Password, databases[i].IP, databases[i].Port, databases[i].Name)
	}
	servers = databases
}

// Open database connection
func Open(index int) (*sql.DB, *error) {
	db, err := sql.Open("mysql", connectionURL[index])
	if err != nil {
		color.Red(fmt.Sprintf("SQL ERROR: Couldn't connect to remote database with index %v", index))
		return nil, &err
	}
	if err = db.Ping(); err != nil {
		color.Red(fmt.Sprintf("SQL ERROR: %s. For branch with index %v", err, index))
		return nil, &err
	}
	color.Green(fmt.Sprintf("Successfully connected to remoteDB with index of %v", index))
	return db, nil
}

func Use23(index int) bool {
	return servers[index].Use23
}

func TotalBranches() int {
	return len(servers)
}