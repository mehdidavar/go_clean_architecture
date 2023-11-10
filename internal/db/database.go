package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (database *sql.DB, err error) {
	// Replace the connection string with your actual database connection details
	connectionString := "username:password@tcp(localhost:3306)/your_database"

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
