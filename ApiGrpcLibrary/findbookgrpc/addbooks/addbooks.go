package main

import (
	"fmt"
	"log"

	"github.com/MajotraderLucky/Utils/logger"
)

type Author struct {
	Name  string   `json:"name"`
	Books []string `json:"books"`
}

type Book struct {
	Title string `json:"title"`
}

func main() {
	logger := logger.Logger{}
	err := logger.CreateLogsDir()
	if err != nil {
		fmt.Println(err)
	}
	err = logger.OpenLogFile()
	if err != nil {
		log.Fatal(err)
	}
	logger.SetLogger()
	logger.LogLine()

	log.Println("Start adding books and authors to the database")
	logger.LogLine()
}
