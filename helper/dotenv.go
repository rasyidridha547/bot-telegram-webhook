package helper

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// DotEnvVar get .env or configmap
func DotEnvVar(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println(fmt.Sprintf("error load %s env. Getting env from server", key))
	}

	return os.Getenv(key)
}
