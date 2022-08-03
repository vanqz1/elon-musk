package main

import (
	"log"
	"os"
	"tweets/api"
)

func main() {
	service := api.NewService()
	err := service.Start()
	if err != nil {
		log.Fatalf("[Fatal] %s", err)
		os.Exit(1)
	}
}
