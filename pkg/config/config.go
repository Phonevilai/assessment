package config

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func getEnv(fileName string) {
	if err := godotenv.Load(fileName); err != nil {
		log.Errorf("please consider environment variables: %s\n", err)
	}
}
