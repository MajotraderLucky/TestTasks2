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

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Чтение содержимого файла log.txt
	content, err := ioutil.ReadFile("logs/log.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Добавление новой строки к содержимому файла
	content = append(content, []byte("Successfully connected to the database!\n")...)

	// Запись обновленного содержимого обратно в файл
	err = ioutil.WriteFile("logs/log.txt", content, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")
}
