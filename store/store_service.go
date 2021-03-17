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
	var result string

	sqlStatement := `SELECT originalURL FROM Shorten_URL WHERE shorturl = $1;`
	row, err := db.Query(sqlStatement, shortURL)
	checkError(err)
	defer row.Close()
	row.Next()
	err = row.Scan(&result)
	checkError(err)

	sqlStatement = `UPDATE Shorten_URL SET count = count + 1 WHERE shorturl = $1;`
	_, err = db.Exec(sqlStatement, shortURL)
	checkError(err)
	return result
}

func GetLinkCount(shortURL string) int {
	var count int
	sqlStatement := `SELECT count FROM Shorten_URL WHERE shorturl = $1;`
	row, err := db.Query(sqlStatement, shortURL)
	checkError(err)
	defer row.Close()

	row.Next()
	err = row.Scan(&count)
	checkError(err)

	return count
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
