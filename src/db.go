package boss

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

type redisOptions struct {
	addr     string
	port     int
	password string
	db       int
}

var redisClient *redis.Client
var sqlite *sql.DB

func init() {
	var err error
	redisClient, err = redisInit()
	checkErr(err)

	sqlite, err = dbInit()
	checkErr(err)

	tableInit()
}

func redisInit() (*redis.Client, error) {
	redisOptions := loadRedisConfig()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisOptions.addr, redisOptions.port),
		Password: fmt.Sprintf("%s", redisOptions.password),
		DB:       redisOptions.db,
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		return nil, err
	}

	return redisClient, nil
}

func loadRedisConfig() *redisOptions {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	addr := os.Getenv("REDIS_ADDR")
	port, _ := strconv.Atoi(os.Getenv("REDIS_PORT"))
	password := os.Getenv("REDIS_PASSWD")
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	redisOptions := &redisOptions{
		addr,
		port,
		password,
		db,
	}

	return redisOptions
}

func dbInit() (*sql.DB, error) {
	sqlite, err := sql.Open("sqlite3", "./boss.db")
	if err != nil {
		return nil, err
	}

	return sqlite, nil
}

func tableInit() {
	sql := `
	CREATE TABLE IF NOT EXISTS jobs(
				id integer primary key not null,
				area varchar(25) not null,
				business varchar(25) not null,
				salary varchar(25) not null,
				age varchar(25) not null,
				type varchar(25) not null
    );
	`
	sqlite.Exec(sql)
}

func redisClose() {
	redisClient.Close()
}

func sqliteClose() {
	sqlite.Close()
}

func cacheArea() {

}

func cacheBusiness() {

}

func saveJD() {
}
