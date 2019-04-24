package boss

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var opts *Options

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

	opts = &Options{
		city,
		job,
		sleep,
	}
}

// 获取区域，深圳 -> 南山区
func getArea() {
	urlTpl := "https://www.zhipin.com/job_detail/?query=%s&scity=%s&source=2"
	url := fmt.Sprintf(urlTpl, opts.job, opts.city)
	contents := request(url)

	parseArea(contents)
}

// 获取商圈，南山区 -> 科技园
func getBusiness() {
}

func getJobList() {

}

func getJD() {

}

func request(url string) []byte {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", fakeBrowser())
	resp, err := client.Do(req)
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	return contents
}

func fakeBrowser() string {
	return "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) "
}
