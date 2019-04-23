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

var redisClient *redis.Client
var sqlite *sql.DB

type redisOptions struct {
	addr     string
	port     int
	password string
	db       int
}

func init() {
	redisInit()
	dbInit()
}

func redisInit() {
	redisOptions := loadRedisConfig()

	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisOptions.addr, redisOptions.port),
		Password: fmt.Sprintf("%s", redisOptions.password),
		DB:       redisOptions.db,
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatalf("redis ping error: %v \n", err)
	}
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

func dbInit() {
	var err error

	sqlite, err = sql.Open("sqlite3", "./boss.db")
	if err != nil {
		log.Fatal(err)
	}

	tableInit()
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
