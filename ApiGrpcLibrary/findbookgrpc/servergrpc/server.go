package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"

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

func (s *MyQueryServiceServer) Query(ctx context.Context, request *pb.QueryRequest) (*pb.QueryResponse, error) {
	// Open a connection to the database.
	db, err := sql.Open("mysql", "myuser:mypassword@tcp(db:3306)/mydb")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Execute a mysql query.
	rows, err := db.Query(request.Sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Forming a mysql response
	result := make(map[string]*pb.Value)
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	for i := range columns {
		valuePtrs[i] = &values[i]
	}
	for rows.Next() {
		err := rows.Scan(valuePtrs...)
		if err != nil {
			return nil, err
		}
		rowResult := make(map[string]*pb.Value)
		for i, col := range columns {
			rowResult[col] = &pb.Value{
				Value: &pb.Value_StringValue{
					StringValue: fmt.Sprintf("%v", values[i])},
			}
		}
		result := make(map[string]*pb.Value)
		for rows.Next() {
			var rowResult string
			if err := rows.Scan(&rowResult); err != nil {
				return nil, err
			}
			result[strconv.Itoa(len(result))] = &pb.Value{
				Value: &pb.Value_StringValue{
					StringValue: rowResult,
				},
			}
		}
		if err := rows.Err(); err != nil {
			return nil, err
		}
	}
	// Creating an instance of pb.QueryResponse and filling it with data
	response := &pb.QueryResponse{
		Results: result,
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
