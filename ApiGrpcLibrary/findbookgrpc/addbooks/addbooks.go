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

func readTableAuthors() {
	db, err := sql.Open("mysql", "myuser:mypassword@tcp(db:3306)/mydb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//  Execute the SELECT * FROM authors request
	query := "SELECT * FROM authors"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Check the available data in the table
	if !rows.Next() {
		log.Println("There are no authors in the database")
		return
	}

	// The output of the authors in the log
	log.Println("The authors in the database:")
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID: %d, Name: %s\n", id, name)
	}
}

func checkAuthors() bool {
	db, err := sql.Open("mysql", "myuser:mypassword@tcp(db:3306)/mydb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check the available data in the table
	query := "SELECT COUNT(*) FROM authors"
	var count int
	err = db.QueryRow(query).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		return false
	} else {
		return true
	}
}

func readTableBooks() {
	db, err := sql.Open("mysql", "myuser:mypassword@tcp(db:3306)/mydb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Execute the SELECT * FROM books
	query := "SELECT * FROM books"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Check the available data in the table
	if !rows.Next() {
		log.Println("There are no books in the database")
		return
	}

	// Output of the book titles in the log
	log.Println("List of the books in the database:")
	for rows.Next() {
		var id int
		var title string
		var author string
		err := rows.Scan(&id, &title, &author)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID: %d, Title: %s, Author: %s\n", id, title, author)
	}
}

func checkBooks() bool {
	db, err := sql.Open("mysql", "myuser:mypassword@tcp(db:3306)/mydb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check the available data in the table
	query := "SELECT COUNT(*) FROM books"
	var count int
	err = db.QueryRow(query).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		return false
	} else {
		return true
	}
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

	readTableAuthors()

	readTableBooks()

	if !checkAuthors() && !checkBooks() {
		log.Println("The base is empty")
		addAuthorsAndBooks()
	}

	readTableBooks()

	logger.CleanLog()
}
