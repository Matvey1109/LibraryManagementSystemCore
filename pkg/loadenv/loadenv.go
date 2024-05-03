package loadenv

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadGlobalEnv() (string, string, string) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	typeOfStorage := os.Getenv("typeOfStorage")
	port := os.Getenv("port")
	pathToLogFile := os.Getenv("pathToLogFile")

	return typeOfStorage, port, pathToLogFile
}

func LoadMongoEnv() (string, string) {
	if err := godotenv.Load(".env_mongo"); err != nil {
		log.Fatal("error loading .env_mongo file")
	}

	MONGO_INITDB_ROOT_USERNAME := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	MONGO_INITDB_ROOT_PASSWORD := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")

	return MONGO_INITDB_ROOT_USERNAME, MONGO_INITDB_ROOT_PASSWORD
}
