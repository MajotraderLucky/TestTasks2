package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"strings"

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
	// Создание директории logs, если она не существует
	err := os.MkdirAll("logs", 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Открытие файла log.txt в режиме добавления и запись в него текста
	logFile, err := os.OpenFile("logs/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	// Установка логгера для вывода в файл
	log.SetOutput(logFile)

	db, err := sql.Open("mysql", "myuser:mypassword@tcp(db:3306)/mydb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Пинг базы данных
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection successful")

	// Чтение содержимого файла log.txt
	data, err := os.ReadFile("logs/log.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Разделение содержимого на строки
	lines := strings.Split(string(data), "\n")

	// Проверка количества строк
	if len(lines) > 50 {
		// Открытие файла log.txt в режиме перезаписи
		logFile, err := os.OpenFile("logs/log.txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer logFile.Close()

		// Запись последних 50 строк в файл
		for _, line := range lines[len(lines)-50:] {
			logFile.WriteString(line + "\n")
		}
	}

	// Получение списка таблиц базы данных
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Чтение названий таблиц и запись их в лог
	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("There are tables in the mydb databases:", tableName)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Выполнение запроса SELECT * FROM authors
	query := "SELECT * FROM authors"
	rows, err = db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Проверка наличия данных в таблице
	if !rows.Next() {
		log.Println("Таблица authors пустая")
	}

	log.Println("Data added")

	// Чтение данных и запись их в лог
	for rows.Next() {
		// Чтение значений строки
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}

		// Запись значений в лог
		log.Println("Data:", id, name)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Выполнение запроса SELECT * FROM books
	query = "SELECT * FROM books"
	rows, err = db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Проверка наличия данных в таблице
	if !rows.Next() {
		log.Println("Таблица books пустая")
	}

	// Чтение данных и запись их в лог
	for rows.Next() {
		// Чтение значений строки
		var id int
		var title string
		var authorID int
		err := rows.Scan(&id, &title, &authorID)
		if err != nil {
			log.Fatal(err)
		}

		// Запись значений в лог
		log.Println("Data:", id, title, authorID)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Чтение содержимого файла
	file, err := os.ReadFile("books.json")
	if err != nil {
		log.Fatal(err)
	}

	// Преобразование содержимого файла в массив структур Author
	var authors []Author
	err = json.Unmarshal(file, &authors)
	if err != nil {
		log.Fatal(err)
	}

	for _, author := range authors {
		// Выполнение операции вставки записи в таблицу authors
		_, err := db.Exec("INSERT INTO authors (name) VALUES (?)", author.Name)
		if err != nil {
			log.Fatal(err)
		}

		// Получение ID последней вставленной записи
		var authorID int64
		err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&authorID)
		if err != nil {
			log.Fatal(err)
		}

		// Вставка данных в таблицу books
		for _, book := range author.Books {
			// Выполнение операции вставки записи в таблицу books
			_, err := db.Exec("INSERT INTO books (title, author_id) VALUES (?, ?)", book, authorID)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
