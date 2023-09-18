package main

import (
	"log"

	"github.com/MajotraderLucky/Utils/logger"
)

func main() {
	// Connect the logger package and make an entry in the log.
	logger := logger.Logger{}
	err := logger.CreateLogsDir()
	if err != nil {
		log.Println(err)
	}
	err = logger.OpenLogFile()
	if err != nil {
		log.Fatal(err)
	}
	logger.SetLogger()
	logger.LogLine()

	log.Println("Hello, Server gRPC!")
}
