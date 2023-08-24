package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

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

	// Execute a mysql query
	rows, err := db.Query(req.Sql)
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

	// Get table names from the database and write them
	// to the slice of table names
	logLine()
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
	logLine()
	log.Println("Got table names: ", tableNames)

	pb.RegisterQueryServiceServer(s, &MyQueryServiceServer{})

	// Print in the log file to the slice of table name
	logLine()
	tableName := tableNames[0]
	log.Println("Getting data from table: ", tableName)
	rows, err = db.Query("SELECT * FROM " + tableName)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var data [][]string
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		data = append(data, []string{strconv.Itoa(id), name})
	}
	logLine()
	log.Println("Got data from table: ", tableName)
	logLine()
	log.Println(data)

	// Start server
	logLine()

	logLine()
	log.Println("Starting gRPC server...")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("Server shutting")
}
