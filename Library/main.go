package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(localhost:3306)/librarydb")
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to ping database:", err)
		return
	}

	fmt.Println("Database connection successful!")
}
