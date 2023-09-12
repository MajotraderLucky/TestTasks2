package dbaccess

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
)

type Database struct {
	db *sql.DB
}

type Author struct {
	Name  string   `json:"name"`
	Books []string `json:"books"`
}

type Book struct {
	Title string `json:"title"`
}

func (d *Database) Connect() error {
	db, err := sql.Open("mysql", "myuser:mypassword@tcp(db:3306)/mydb")
	if err != nil {
		return err
	}
	d.db = db
	return nil
}

func (d *Database) Ping() error {
	err := d.db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) TakeTablesNames() error {
	// Get a list database tables
	rows, err := d.db.Query("SHOW TABLES")
	if err != nil {
		return err
	}
	defer rows.Close()

	// Read table names and write them to the log
	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			return err
		}
		log.Println("There are tables in the database:", tableName)
	}

	if err = rows.Err(); err != nil {
		return err
	}
	return nil
}

func (d *Database) CheckAuthors() bool {
	// Check the available data in the authors table
	query := "SELECT COUNT(*) FROM authors"
	var count int
	err := d.db.QueryRow(query).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		return false
	} else {
		return true
	}
}

func (d *Database) CheckBooks() bool {
	// Check the available data in the books table
	query := "SELECT COUNT(*) FROM books"
	var count int
	err := d.db.QueryRow(query).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		return false
	} else {
		return true
	}
}

func (d *Database) ReadTableAuthors() error {
	// Output header for the start function
	log.Println("-----------Starting the function ReadTableAuthors-----------")
	// Execute the SELECT * FROM authors request
	query := "SELECT * FROM authors"
	rows, err := d.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	// The output of the authors in the log
	log.Println("The authors in the database:")
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			return err
		}
		log.Printf("ID: %d, Name: %s\n", id, name)
	}
	if err = rows.Err(); err != nil {
		return err
	}
	// Output footer for the finished function
	log.Println("-----------Finished the function ReadTableAuthors-----------")
	return nil
}

func (d *Database) ReadTableBooks() error {
	// Output header for the start function
	log.Println("-----------Starting the function ReadTableBooks-----------")
	// Execute the SELECT * FROM books request
	query := "SELECT * FROM books"
	rows, err := d.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	// The output of the books in the log
	log.Println("The books in the database:")
	for rows.Next() {
		var id int
		var title string
		var author int
		err := rows.Scan(&id, &title, &author)
		if err != nil {
			return err
		}
		log.Printf("ID: %d, Title: %s, Author: %d\n", id, title, author)
	}
	if err = rows.Err(); err != nil {
		return err
	}
	// Output footer for the finished function
	log.Println("-----------Finished the function ReadTableBooks-----------")
	return nil
}

func (d *Database) AddAuthorsAndBooks() error {
	// Output header for the start function
	log.Println("-----------Starting the function addAuthorsAndBooks-----------")
	initSQL :=
		`CREATE TABLE IF NOT EXISTS authors (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255)
	);`
	_, err := d.db.Exec(initSQL)
	if err != nil {
		return err
	}

	initSQL2 :=
		`CREATE TABLE IF NOT EXISTS books (
			id INT AUTO_INCREMENT PRIMARY KEY,
			title VARCHAR(255),
			author_id INT,
			FOREIGN KEY (author_id) REFERENCES authors(id)
			);`

	_, err = d.db.Exec(initSQL2)
	if err != nil {
		return err
	}

	file, err := os.ReadFile("books.json")
	if err != nil {
		return err
	}
	var authors []Author
	err = json.Unmarshal(file, &authors)
	if err != nil {
		return err
	}

	for _, author := range authors {
		insertAuthorSQL := "INSERT INTO authors (name) VALUES (?)"
		result, err := d.db.Exec(insertAuthorSQL, author.Name)
		if err != nil {
			return err
		}
		authorID, err := result.LastInsertId()
		if err != nil {
			return err
		}
		for _, book := range author.Books {
			insertBookSQL := "INSERT INTO books (title, author_id) VALUES (?,?)"
			_, err := d.db.Exec(insertBookSQL, book, authorID)
			if err != nil {
				return err
			}
		}
	}
	log.Println("Data inserted successfully")
	// Output footer for the finished function
	log.Println("-----------Finished the function addAuthorsAndBooks-----------")
	return nil
}
