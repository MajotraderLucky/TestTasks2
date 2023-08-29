package main

import (
	"context"
	"log"
	"os"

	pb "github.com/MajotraderLucky/TestTasks2/Repo/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

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
	log.Println("---------------------------------------------------")
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

	setLogger(logFile)

	logLine()

	log.Println("Starting grpc client")

	// Establish a connection to the gRPC server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client.
	client := pb.NewQueryServiceClient(conn)

	// Call the remote getData method.
	resp, err := client.GetData(context.Background(), &pb.GetDataRequest{DataId: "your_id"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// Response processing
	log.Printf("Response: %v", resp.GetTableNames())
}
