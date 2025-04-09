package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

//env reader

func LoadEnv() error {

	file, fileErr := os.Open(".env")
	if os.IsNotExist(fileErr) {
		return errors.New(".env file not found")
	}

	err := godotenv.Load(file.Name())
	if err != nil {
		return errors.New("error reading .env file")
	}

	log.Print(".env charged successfully")

	return nil
}

func GetMYQSLDBName() string {
	return os.Getenv("MYSQL_DB_NAME")
}
