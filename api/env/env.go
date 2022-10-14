package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Env(env_var string) string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading env file")
	}

	return os.Getenv(env_var)
}
