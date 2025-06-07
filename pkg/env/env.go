package env

import (
	"github.com/lpernett/godotenv"
	"log"
)

func NewEnv(filepath string) error {
	if err := godotenv.Load(filepath); err != nil {
		log.Println("Error while loading .env", err)
		return err
	}
	log.Println("Successfully loaded .env")
	return nil
}
