package main

import (
	"database/sql"
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

func createLogsDirectory() error {
	err := os.MkdirAll("logs", 0755)
	if err != nil {
		return err
	}
	return nil
}

func openLogFile() (*os.File, error) {
	logFile, err := os.OpenFile("logs/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}

func setLogger(logFile *os.File) {
	log.SetOutput(logFile)
}

func connectToDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", "myuser:mypassword@tcp(db:3306)/mydb")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func pingDatabase(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func logLine() {
	log.Println("-------------------------------------------------")
}

func cleanLog() {
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
}

func takeTables() {
	db, err := sql.Open("mysql", "myuser:mypassword@tcp(db:3306)/mydb")
	if err != nil {
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
}

func readTableAuthors() {
	db, err := sql.Open("mysql", "myuser:mypassword@tcp(db:3306)/mydb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Выполнение запроса SELECT * FROM authors
	query := "SELECT * FROM authors"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Проверка наличия данных в таблице
	if !rows.Next() {
		log.Println("Таблица authors пустая")
		return
	}

	// Вывод авторов в лог
	log.Println("Список авторов:")
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

func cleanBooksAndAuthors(authorID int) error {
	db, err := sql.Open("mysql", "myuser:mypassword@tcp(db:3306)/mydb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM books WHERE author_id = $1", authorID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM authors WHERE id = $1", authorID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func checkAuthors() bool {
	db, err := sql.Open("mysql", "myuser:mypassword@tcp(db:3306)/mydb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Проверка наличия данных в таблице
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

	// Выполнение запроса SELECT * FROM books
	query := "SELECT * FROM books"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Проверка наличия данных в таблице
	if !rows.Next() {
		log.Println("Таблица books пустая")
		return
	}

	// Вывод книг в лог
	log.Println("Список книг:")
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

	// Проверка наличия данных в таблице
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

func main() {
	err := createLogsDirectory()
	if err != nil {
		log.Fatal(err)
	}

	logFile, err := openLogFile()
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	setLogger(logFile)

	logLine()
	log.Println("Start aplication")

	db, err := connectToDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = pingDatabase(db)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection successful")

	takeTables()

	readTableAuthors()

	readTableBooks()

	if !checkAuthors() && !checkBooks() {
		log.Println("The base is empty")
	}

	cleanLog()
}
