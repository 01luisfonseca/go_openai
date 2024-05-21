package tools

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic(err)
	}
	log.Println("Loaded env")
}

// GetEnv gets the value of an environment variable
func GetEnv(key string) string {
	return os.Getenv(strings.ToUpper(key))
}
