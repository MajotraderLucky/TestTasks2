package main

import (
	"database/sql"
	"fmt"
	"log"

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

	// Здесь вы можете добавить свой код для управления базой данных

	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
