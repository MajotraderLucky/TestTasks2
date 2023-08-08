package main

import (
	"log"
	"os"
)

func createLogsDirectory() error {
	err := os.MkdirAll("logs", 0755)
	if err != nil {
		return err
	}
	return nil
}

func openLogFile() (*os.File, error) {
	logFile, err := os.OpenFile("logs/serverlog.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}

func main() {
	// Создание директории для логов
	err := createLogsDirectory()
	if err != nil {
		log.Fatal(err)
	}

	// Открытие файла для записи логов
	logFile, err := openLogFile()
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	// Настройка вывода логов в файл
	log.SetOutput(logFile)

	// Запись в лог
	log.Println("Hello, GRPC Server!")
}
