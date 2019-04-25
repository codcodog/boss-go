package boss

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var opts *Options

type Options struct {
	city     string
	cityCode int
	job      string
	sleep    int
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	city := os.Getenv("BOSS_CITY")
	cityCode, _ := strconv.Atoi(os.Getenv("BOSS_CITY_CODE"))
	job := os.Getenv("BOSS_JOB")
	sleep, _ := strconv.Atoi(os.Getenv("BOSS_SLEEP"))

	opts = &Options{
		city,
		cityCode,
		job,
		sleep,
	}
}

// 获取区域，深圳 -> 南山区
func getArea() {
	if isCachedArea() {
		return
	}
	crawlArea()
}

func crawlArea() {
	urlTpl := "https://www.zhipin.com/job_detail/?query=%s&scity=%s&source=2"
	areaUrl := fmt.Sprintf(urlTpl, opts.job, opts.city)
	encodeUrl := getEncodeUrl(areaUrl)
	resp := request(encodeUrl)

	parseArea(resp)
}

// 获取商圈，南山区 -> 科技园
func getBusiness() {
	areas := getAreaCache()
	for _, area := range areas {
		if !isCachedBusiness(area) {
			getBusinessByArea(area)
		}
	}
}

func getBusinessByArea(area string) {
	urlTpl := "https://www.zhipin.com/c%d/b_%s-h_%s/?query=%s&ka="
	businessUrl := fmt.Sprintf(urlTpl, opts.cityCode, area, opts.city, opts.job)
	encodeUrl := getEncodeUrl(businessUrl)
	resp := request(encodeUrl)

	parseBusiness(area, resp)
}

func getJobList() {

}

func getJD() {

}

func getEncodeUrl(req string) string {
	u, _ := url.Parse(req)
	query := u.Query()
	u.RawQuery = query.Encode()

	return u.String()
}

func request(url string) *http.Response {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", fakeBrowser())
	resp, err := client.Do(req)
	checkErr(err)

	return resp
}

func fakeBrowser() string {
	return "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) "
}
