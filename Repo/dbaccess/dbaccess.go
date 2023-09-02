package dbaccess

import "database/sql"

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
