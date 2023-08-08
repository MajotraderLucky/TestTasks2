package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Ping(ctx context.Context, request *PingRequest) (*PingResponse, error) {
	// Подключение к базе данных
	db, err := sql.Open("mysql", "myuser:mypassword@tcp(db:3306)/mydb")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Пинг базы данных
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	return &PingResponse{Message: "Ping successful"}, nil
}

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

	// Создание gRPC сервера
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterPingServiceServer(s, &server{})

	// Запуск gRPC сервера
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
