package bootstrap

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Print(err)
	}
}
