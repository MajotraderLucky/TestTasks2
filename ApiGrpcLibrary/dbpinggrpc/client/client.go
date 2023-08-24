package main

import (
	"context"
	"log"
	"os"

	pb "github.com/MajotraderLucky/TestTasks2/Repo/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

	// Creat unprotected credentials
	creds := credentials.NewTLS(nil)

	// Establish a connection to the gRPC server.
	conn, err := grpc.Dial("grpc:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client.
	client := pb.NewQueryServiceClient(conn)

	// Call the remote getData method.
	resp, err := client.GetData(context.Background(), &pb.GetDataRequest{Id: "your_id"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// Response processing
	log.Printf("Response: %s", resp.Message)

}
