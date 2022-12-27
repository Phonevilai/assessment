package main

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	if err := godotenv.Load("dev.env"); err != nil {
		log.Errorf("please consider environment variables: %s\n", err)
	}
}

func main() {

}
