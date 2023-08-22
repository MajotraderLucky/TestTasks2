package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"os"

	pb "github.com/MajotraderLucky/TestTasks2/Repo/protobuf"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

type MyQueryServiceServer struct {
	pb.QueryServiceServer
}

func (s MyQueryServiceServer) Query(ctx context.Context, req *pb.QueryRequest) (*pb.QueryResponse, error) {
	// Open a connection to the database
	db, err := sql.Open("mysql", "myuser:mypassword@tcp(db:3306)/mydb")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return &pb.QueryResponse{}, nil
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

func logLine() {
	log.Println("----------------------------------------")
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
	log.Println("Starting grpc server...")

	// Connect to the database
	db, err := sql.Open("mysql", "myuser:mypassword@tcp(db:3306)/mydb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Connecting to database")

	// create a new server gRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Server listening on port 50051")

	s := grpc.NewServer()

	// Ping to the database
	logLine()
	log.Println("Pinging to the database...")
	_, err = db.Exec("SELECT 1")
	if err != nil {
		log.Fatal(err)
	}
	logLine()
	log.Println("Database pinged successfully!")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("Server shutting")
}
