package main

import (
	"database/sql"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

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
}
