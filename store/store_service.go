package store

import (
	"database/sql"
	"fmt"

	"github.com/shin-iji/go-shorten-url/database"

	_ "github.com/lib/pq"
)

var db *sql.DB = database.OpenConnection()

func SaveURLMapping(shortURL string, originalURL string) {
	sqlStatement := `INSERT INTO Shorten_URL (shortURL, originalURL, count) VALUES ($1, $2, 0)`
	_, err := db.Exec(sqlStatement, shortURL, originalURL)
	checkError(err)

	fmt.Printf("Saved shortUrl: %s - originalUrl: %s\n", shortURL, originalURL)
}

func RetrieveInitialURL(shortURL string) string {
	var originalURL string

	sqlStatement := `UPDATE Shorten_URL SET count = count + 1 WHERE shorturl = $1 RETURNING originalURL;`
	err := db.QueryRow(sqlStatement, shortURL).Scan(&originalURL)
	checkError(err)

	return originalURL
}

func GetLinkCount(shortURL string) int {
	var count int
	sqlStatement := `SELECT count FROM Shorten_URL WHERE shorturl = $1;`
	err := db.QueryRow(sqlStatement, shortURL).Scan(&count)
	checkError(err)

	return count
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
