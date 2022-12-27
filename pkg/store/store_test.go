package store

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestConnectToDB(t *testing.T) {

	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	t.Run("Test Connection", func(t *testing.T) {
		if err := godotenv.Load("../../dev.env"); err != nil {
			log.Errorf("please consider environment variables: %s\n", err)
		}
		NewDb(os.Getenv("DATABASE_URL"))
	})
}
