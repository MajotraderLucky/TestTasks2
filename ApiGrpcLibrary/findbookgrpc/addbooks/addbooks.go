package main

import (
	"fmt"
	"log"

	"github.com/MajotraderLucky/TestTasks2/Repo/dbaccess"
	"github.com/MajotraderLucky/Utils/logger"
	_ "github.com/go-sql-driver/mysql"
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

	// Database connection
	db := dbaccess.Database{}
	err = db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected successfully")

	// Database ping
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database pinged successfully")
	logger.LogLine()

	// Take tables from the database
	err = db.TakeTablesNames()
	if err != nil {
		log.Fatal(err)
	}
	logger.LogLine()

	// Check table authors
	if db.CheckAuthors() {
		// If the authors table is not empty, call the function ReadTableAuthors
		err = db.ReadTableAuthors()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// If the authors table is empty, we output a message
		log.Println("There are no authors in the database")
	}

	// Check table books
	if db.CheckBooks() {
		// If the books table is not empty, call the function ReadTableBooks
		err = db.ReadTableBooks()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// If the books table is empty, we output a message
		log.Println("There are no books in the database")
	}

	if !db.CheckAuthors() && !db.CheckBooks() {
		log.Println("The base is empty")
		err = db.AddAuthorsAndBooks()
		if err != nil {
			log.Fatal(err)
		}
	}

	logger.LogLine()

	logger.CleanLog()
}
