package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Hello, world!")

	// Создание директории logs, если она не существует
	err := os.MkdirAll("logs", 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Создание файла log.txt и запись в него текста
	err = ioutil.WriteFile("logs/log.txt", []byte("Hello, logs!\n"), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File log.txt created and text written successfully.")

	db, err := sql.Open("mysql", "myuser:mypassword@tcp(db:3306)/mydb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tableName string
	for rows.Next() {
		err := rows.Scan(&tableName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(tableName)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
