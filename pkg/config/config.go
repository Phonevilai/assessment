package config

import (
	"fmt"
	"github.com/joho/godotenv"
)

func GetEnv(fileName string) {
	if err := godotenv.Load(fileName); err != nil {
		fmt.Printf("please consider environment variables: %s\n", err)
	}
}
