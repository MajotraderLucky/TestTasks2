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

func (d *Database) takeTablesNames() error {
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
