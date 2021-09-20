package env

import (
	"github.com/joho/godotenv"
	"os"
)

var (
	DbUserName string
	DbPassword string
	DbHost     string
	DbName     string
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Can't load .env file")
	}

	DbUserName = os.Getenv("DBUSERNAME")
	DbPassword = os.Getenv("DBPASSWORD")
	DbHost = os.Getenv("DBHOST")
	DbName = os.Getenv("DBNAME")
}
