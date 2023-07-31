package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello, world!")

	// Создание директории logs, если она не существует
	err := os.MkdirAll("logs", 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Создание файла log.txt и запись в него текста
	err = ioutil.WriteFile("logs/log.txt", []byte("Hello, logs!"), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File log.txt created and text written successfully.")
}
