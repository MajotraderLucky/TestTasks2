package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

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

func addAuthorsAndBooks() {
	db, err := sql.Open("mysql", "myuser:mypassword@tcp(db:3306)/mydb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	initSQL :=
		`CREATE TABLE IF NOT EXISTS authors (
  id INT AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(255)
	);`

	_, err = db.Exec(initSQL)
	if err != nil {
		log.Fatal(err)
	}

	initSQL2 :=
		`CREATE TABLE IF NOT EXISTS books (
	id INT AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(255),
  author_id INT,
	FOREIGN KEY (author_id) REFERENCES authors(id)
	);`

	_, err = db.Exec(initSQL2)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.ReadFile("books.json")
	if err != nil {
		log.Fatal(err)
	}

	var authors []Author
	err = json.Unmarshal(file, &authors)
	if err != nil {
		log.Fatal(err)
	}

	for _, author := range authors {
		insertAuthorSQL := "INSERT INTO authors (name) VALUES (?)"
		result, err := db.Exec(insertAuthorSQL, author.Name)
		if err != nil {
			log.Fatal(err)
		}

		authorID, err := result.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}

		for _, book := range author.Books {
			insertBookSQL := "INSERT INTO books (title, author_id) VALUES (?,?)"
			_, err := db.Exec(insertBookSQL, book, authorID)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	log.Println("Data inserted successfully")
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
		addAuthorsAndBooks()
	}

	logger.LogLine()

	logger.CleanLog()
}
