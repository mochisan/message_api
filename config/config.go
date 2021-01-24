package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	// SQLConnectionName .
	SQLConnectionName string
)

type environment string

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(".env not found")
	}

	databaseName := os.Getenv("DATABASE_NAME")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databaseUserName := os.Getenv("DATABASE_USER_NAME")
	databaseHost := os.Getenv("DATABASE_HOST")
	databasePort := os.Getenv("DATABASE_PORT")

	SQLConnectionName = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4,utf8&parseTime=true", databaseUserName, databasePassword, databaseHost, databasePort, databaseName)
}
