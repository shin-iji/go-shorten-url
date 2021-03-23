package store

import (
	"database/sql"
	"fmt"

	"github.com/shin-iji/go-shorten-url/database"

	_ "github.com/lib/pq"
)

var db *sql.DB = database.OpenConnection()

func SaveURLMapping(shortURL string, originalURL string) {
	sqlStatement := `SELECT originalURL FROM Shorten_URL WHERE shortURL = $1;`
	row := db.QueryRow(sqlStatement, shortURL)
	err := row.Scan(&originalURL)
	if err != nil {
		if err == sql.ErrNoRows {
			sqlStatement := `INSERT INTO Shorten_URL (shortURL, originalURL, count) VALUES ($1, $2, 0)`
			_, err := db.Exec(sqlStatement, shortURL, originalURL)
			checkError(err)
		} else {
			panic(err)
		}
	}
	fmt.Printf("Saved shortUrl: %s - originalUrl: %s\n", shortURL, originalURL)
}

func RetrieveInitialURL(shortURL string) string {
	var originalURL string

	sqlStatement := `UPDATE Shorten_URL SET count = count + 1 WHERE shorturl = $1 RETURNING originalURL;`
	row := db.QueryRow(sqlStatement, shortURL)
	err := row.Scan(&originalURL)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(err)
		}
	}

	return originalURL
}

func GetLinkCount(shortURL string) int {
	var count int
	sqlStatement := `SELECT count FROM Shorten_URL WHERE shorturl = $1;`
	row := db.QueryRow(sqlStatement, shortURL)
	err := row.Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(err)
		}
	}

	return count
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
