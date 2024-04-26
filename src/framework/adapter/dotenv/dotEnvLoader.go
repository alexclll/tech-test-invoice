package dotenv

import (
	"log"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err.Error())
	}
}
