package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type sqlEnv struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func GetSQLEnv() sqlEnv {
	// load env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("failed to load env")
	}

	return sqlEnv{
		Host:     os.Getenv("host"),
		Port:     os.Getenv("dbPort"),
		User:     os.Getenv("user"),
		Password: os.Getenv("password"),
		DBName:   os.Getenv("dbName"),
	}
}

func GetServerEnv() string {
	// load env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("failed to load env")
	}

	return os.Getenv("GO_PORT")
}

func GoDotEnv(key string) string {
	env := make(chan string, 1)

	if os.Getenv("GO_ENV") != "production" {
		godotenv.Load(".env")
		env <- os.Getenv(key)
	} else {
		env <- os.Getenv(key)
	}

	return <-env
}
