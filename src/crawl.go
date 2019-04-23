package boss

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var options *Options

type Options struct {
	city  string
	job   string
	sleep int
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	city := os.Getenv("BOSS_CITY")
	job := os.Getenv("BOSS_JOB")
	sleep, _ := strconv.Atoi(os.Getenv("BOSS_SLEEP"))

	options = &Options{
		city,
		job,
		sleep,
	}
}

// 获取区域，深圳 -> 南山区
func getArea() {
	fmt.Println(options.city)
}

// 获取商圈，南山区 -> 科技园
func getBusiness() {
}

func getJobList() {

}

func getJD() {

}

func fakeBrowser() string {
	return "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) "
}
