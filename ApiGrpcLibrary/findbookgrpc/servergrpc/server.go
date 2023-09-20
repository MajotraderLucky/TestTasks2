package main

import (
	"context"
	"database/sql"
	"log"

	pb "github.com/MajotraderLucky/TestTasks2/Repo/protobuf"
	"github.com/MajotraderLucky/Utils/logger"
	_ "github.com/go-sql-driver/mysql"
)

type MyQueryServiceServer struct {
	pb.QueryServiceServer
}

func (s *MyQueryServiceServer) GetData(ctx context.Context, request *pb.GetDataRequest) (*pb.GetDataResponse, error) {
	// Open a connection to the database.
	db, err := sql.Open("mysql", "myuser:mypassword@tcp(db:3306)/mydb")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Get table names from the database and write them to the response.
	logger := logger.Logger{}
	logger.LogLine()
	log.Println("Getting table names from the database...")
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tableNames []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			log.Fatal(err)
		}
		tableNames = append(tableNames, tableName)
	}
	logger.LogLine()
	log.Println("Got table names:", tableNames)

	// Create and return the response.
	response := &pb.GetDataResponse{
		TableNames: tableNames,
	}
	return response, nil
}

func main() {
	// Connect the logger package and make an entry in the log.
	logger := logger.Logger{}
	err := logger.CreateLogsDir()
	if err != nil {
		log.Println(err)
	}
	err = logger.OpenLogFile()
	if err != nil {
		log.Fatal(err)
	}
	logger.SetLogger()
	logger.LogLine()

	log.Println("Hello, Server gRPC!")

	// Connect to the database.
	db, err := sql.Open("mysql", "myuser:mypassword@tcp(db:3306)/mydb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
