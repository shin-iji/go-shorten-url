package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/shin-iji/go-shorten-url/database"

	_ "github.com/lib/pq"
)

var (
	storeService = &StorageService{}
	ctx          = context.Background()
	db           = database.OpenConnection()
)

const CacheDuration = 6 * time.Hour

type StorageService struct {
	redisClient *redis.Client
}

func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}

func SaveURLMapping(shortURL string, originalURL string) {
	//defer db.Close()
	sqlStatement := `INSERT INTO Shorten_URL (shortURL, originalURL, count) VALUES ($1, $2, 0)`
	_, err := db.Query(sqlStatement, shortURL, originalURL)
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortURL, originalURL))
	}
	fmt.Printf("Saved shortUrl: %s - originalUrl: %s\n", shortURL, originalURL)
}

func RetrieveInitialURL(shortURL string) string {
	//defer db.Close()
	var result string
	result, err := storeService.redisClient.Get(ctx, shortURL).Result()
	if err != nil {
		sqlStatement := `SELECT originalURL FROM Shorten_URL WHERE shorturl = $1;`
		row := db.QueryRow(sqlStatement, shortURL)
		err = row.Scan(&result)
		if err != nil {
			panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortURL))
		}
		err := storeService.redisClient.Set(ctx, shortURL, result, CacheDuration).Err()
		if err != nil {
			panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortURL, result))
		}
	}
	sqlStatement := `UPDATE Shorten_URL SET count = count + 1 WHERE shorturl = $1;`
	_, err = db.Query(sqlStatement, shortURL)
	if err != nil {
		panic(fmt.Sprintf("Failed IncreaseViewPage url | Error: %v - shortUrl: %s\n", err, shortURL))
	}
	return result
}

func GetLinkCount(shortURL string) int {
	//defer db.Close()
	var count int
	sqlStatement := `SELECT count FROM Shorten_URL WHERE shorturl = $1;`
	row := db.QueryRow(sqlStatement, shortURL)
	err := row.Scan(&count)
	if err != nil {
		panic(fmt.Sprintf("Failed GetLinkCount url | Error: %v - shortUrl: %s\n", err, shortURL))
	}
	return count
}
