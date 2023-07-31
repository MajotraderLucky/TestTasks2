package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "myuser:mypassword@tcp(172.24.0.1:3306)/mydb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")

	// Создание таблицы authors
	createAuthorsTable :=
		`CREATE TABLE IF NOT EXISTS authors (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL);`

	_, err = db.Exec(createAuthorsTable)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Create table autors successfully")

	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}

	logDir := "/home/ryazanov/Development/TestTasks/LibraryCompose/logs"
	err = os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	logFile := filepath.Join(logDir, time.Now().Format("20060102_150405")+".log")
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)
	log.Println("Successfully connected to the database!")
	log.Println("Create table authors successfully")
}
