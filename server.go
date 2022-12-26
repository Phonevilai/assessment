package main

import (
	"fmt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Errorf("please consider environment variables: %s\n", err)
	}
}

func main() {
	fmt.Println("Please use server.go for main file")
	fmt.Println("start at port:", os.Getenv("PORT"))
}
