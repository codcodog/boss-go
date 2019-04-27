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

const (
	areaKey        = "BOSS:AREA"
	businessKeyTpl = "BOSS:BUSINESS:%s"
	taskKey        = "BOSS:TASK"
	blockKey       = "BOSS:BLOCK"
)

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
				salary int not null,
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

func setAreaCache(area string) {
	redisClient.SAdd(areaKey, area)
}

func getAreaCache() []string {
	return getSmembers(areaKey)
}

func isCachedArea() bool {
	return isKeyExists(areaKey)
}

func setBusinessCache(area, business string) {
	redisClient.SAdd(getBusinessKey(area), business)
}

func getBusinessCache(area string) []string {
	return getSmembers(getBusinessKey(area))
}

func isCachedBusiness(area string) bool {
	return isKeyExists(getBusinessKey(area))
}

func getBusinessKey(area string) string {
	return fmt.Sprintf(businessKeyTpl, area)
}

func setTask(url string) {
	redisClient.LPush(taskKey, url)
}

func getTask() string {
	url, err := redisClient.RPop(taskKey).Result()
	checkErr(err)

	return url
}

func isCreatedTaskQueue() bool {
	return isKeyExists(taskKey)
}

func isEmptyTaskQueue() bool {
	return getTaskLen() == 0
}

func getTaskLen() int64 {
	len, err := redisClient.LLen(taskKey).Result()
	checkErr(err)

	return len
}

func setBlockRecord() int64 {
	record, err := redisClient.Incr(blockKey).Result()
	checkErr(err)

	return record
}

func restoreTask(url string) {
	redisClient.RPush(taskKey, url)
}

func getSmembers(key string) []string {
	members, err := redisClient.SMembers(key).Result()
	checkErr(err)

	return members
}

func isKeyExists(key string) bool {
	exists, err := redisClient.Exists(key).Result()
	checkErr(err)

	return exists != 0
}

func saveJD(area string, business string, salary int, experience string, industry string) {
	sqlTpl := "insert into jobs (area, business, salary, age, type) values (?, ?, ?, ?, ?)"
	stmt, err := sqlite.Prepare(sqlTpl)
	checkErr(err)

	_, err = stmt.Exec(area, business, salary, experience, industry)
	checkErr(err)
}
