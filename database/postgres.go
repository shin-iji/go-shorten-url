package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// const (
// 	// Initialize connection constants.
// 	HOST     = "localhost"
// 	DATABASE = "postgres"
// 	USER     = "postgres"
// 	PASSWORD = "12345"
// )

const (
	// Initialize connection constants.
	HOST     = "13.76.167.208"
	DATABASE = "postgres"
	USER     = "postgres"
	PASSWORD = "Azureuser123"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func OpenConnection() *sql.DB {
	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE)

	// Initialize connection object.
	db, err := sql.Open("postgres", connectionString)
	checkError(err)

	err = db.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database")

	return db
}
