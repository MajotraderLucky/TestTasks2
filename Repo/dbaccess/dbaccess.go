package dbaccess

import (
	"database/sql"
	"log"
)

type Database struct {
	db *sql.DB
}

func (d *Database) Connect() error {
	db, err := sql.Open("mysql", "myuser:mypassword@tcp(db:3306)/mydb")
	if err != nil {
		return err
	}
	d.db = db
	return nil
}

func (d *Database) Ping() error {
	err := d.db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) TakeTablesNames() error {
	// Get a list database tables
	rows, err := d.db.Query("SHOW TABLES")
	if err != nil {
		return err
	}
	defer rows.Close()

	// Read table names and write them to the log
	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			return err
		}
		log.Println("There are tables in the database:", tableName)
	}

	if err = rows.Err(); err != nil {
		return err
	}
	return nil
}

func (d *Database) CheckAuthors() bool {
	// Check the available data in the authors table
	query := "SELECT COUNT(*) FROM authors"
	var count int
	err := d.db.QueryRow(query).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		return false
	} else {
		return true
	}
}

func (d *Database) ReadTableAuthors() error {
	// Output header for the start function
	log.Println("-----------Starting the function ReadTableAuthors-----------")
	// Execute the SELECT * FROM authors request
	query := "SELECT * FROM authors"
	rows, err := d.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	// The output of the authors in the log
	log.Println("The authors in the database:")
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			return err
		}
		log.Printf("ID: %d, Name: %s\n", id, name)
	}
	if err = rows.Err(); err != nil {
		return err
	}
	// Output footer for the finished function
	log.Println("-----------Finished the function ReadTableAuthors-----------")
	return nil
}
