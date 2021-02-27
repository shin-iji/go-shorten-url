package store

import (
	//"database/sql"
	"fmt"

	"github.com/shin-iji/go-shorten-url/database"

	_ "github.com/lib/pq"
)

func SaveURLMapping(shortURL string, originalURL string) {
	db := database.OpenConnection()
	sqlStatement := `INSERT INTO Shorten_URL (shortURL, originalURL) VALUES ($1, $2)`
	_, err := db.Query(sqlStatement, shortURL, originalURL)
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortURL, originalURL))
	}
	fmt.Printf("Saved shortUrl: %s - originalUrl: %s\n", shortURL, originalURL)
	defer db.Close()
}

func RetrieveInitialURL(shortURL string) string {
	var result string
	db := database.OpenConnection()
	sqlStatement := `SELECT originalURL FROM Shorten_URL WHERE shorturl = $1;`
	row := db.QueryRow(sqlStatement, shortURL)
	err := row.Scan(&result)
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortURL))
	}
	defer db.Close()
	return result
}
