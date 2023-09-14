package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() (*sql.DB, error) {
	// Replace with your MySQL database credentials
	username := "root"
	password := ""
	dbname := "chob-backend"
	host := "localhost" // or your MySQL server's address

	// Create a MySQL DSN (Data Source Name) string
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, dbname)

	// Open a connection to the MySQL database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Test the database connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to the MySQL database")
	return db, nil
}
