package main

import (
	"fmt"
	"github.com/Phonevilai/assessment/pkg/config"
	"os"
)

func init() {
	config.GetEnv("dev.env")
}

func main() {
	fmt.Println("Please use server.go for main file")
	fmt.Println("start at port:", os.Getenv("PORT"))
}
