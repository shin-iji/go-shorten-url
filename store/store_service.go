package store

import (
	"fmt"

	"github.com/shin-iji/go-shorten-url/database"

	_ "github.com/lib/pq"
)

func SaveURLMapping(shortURL string, originalURL string) {
	db := database.OpenConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO Shorten_URL (shortURL, originalURL, count) VALUES ($1, $2, 0)`
	_, err := db.Query(sqlStatement, shortURL, originalURL)
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortURL, originalURL))
	}
	fmt.Printf("Saved shortUrl: %s - originalUrl: %s\n", shortURL, originalURL)
}

func RetrieveInitialURL(shortURL string) string {
	var result string
	db := database.OpenConnection()
	defer db.Close()
	sqlStatement := `SELECT originalURL FROM Shorten_URL WHERE shorturl = $1;`
	row := db.QueryRow(sqlStatement, shortURL)
	err := row.Scan(&result)
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortURL))
	}
	sqlStatement = `UPDATE Shorten_URL SET count = count + 1 WHERE shorturl = $1;`
	_, err = db.Query(sqlStatement, shortURL)
	if err != nil {
		panic(fmt.Sprintf("Failed IncreaseViewPage url | Error: %v - shortUrl: %s\n", err, shortURL))
	}
	return result
}

func GetLinkCount(shortURL string) int {
	var count int
	db := database.OpenConnection()
	sqlStatement := `SELECT count FROM Shorten_URL WHERE shorturl = $1;`
	row := db.QueryRow(sqlStatement, shortURL)
	err := row.Scan(&count)
	if err != nil {
		panic(fmt.Sprintf("Failed GetLinkCount url | Error: %v - shortUrl: %s\n", err, shortURL))
	}
	defer db.Close()
	return count
}
