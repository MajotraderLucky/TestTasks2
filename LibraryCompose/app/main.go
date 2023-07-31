package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	pb "github.com/MajotraderLucky/TestTasks2/Repo/protobuf"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	db *sql.DB
}

func (s *server) FindBooksByAuthor(ctx context.Context, req *pb.AuthorRequest) (*pb.BooksResponse, error) {
	rows, err := s.db.Query("SELECT book_name FROM books WHERE author_name = ?", req.AuthorName)
	if err != nil {
		log.Printf("Error querying books: %v", err)
		return nil, err
	}
	defer rows.Close()

	var books []string
	for rows.Next() {
		var book string
		err := rows.Scan(&book)
		if err != nil {
			log.Printf("Error scanning book: %v", err)
			continue
		}
		books = append(books, book)
	}

	return &pb.BooksResponse{Books: books}, nil
}

func (s *server) FindAuthorsByBook(ctx context.Context, req *pb.BookRequest) (*pb.AuthorsResponse, error) {
	rows, err := s.db.Query("SELECT author_name FROM authors WHERE book_name = ?", req.BookName)
	if err != nil {
		log.Printf("Error querying authors: %v", err)
		return nil, err
	}
	defer rows.Close()

	var authors []string
	for rows.Next() {
		var author string
		err := rows.Scan(&author)
		if err != nil {
			log.Printf("Error scanning author: %v", err)
			continue
		}
		authors = append(authors, author)
	}

	return &pb.AuthorsResponse{Authors: authors}, nil
}

func main() {
	db, err := sql.Open("mysql", "myuser:mypassword@tcp(172.24.0.2:3306)/mydb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLibraryServiceServer(s, &server{db: db})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
