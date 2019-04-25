package boss

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func parseArea(resp *http.Response) {
	doc := parseDoc(resp)
	doc.Find("dl.condition-district").Find("a").Each(func(i int, selector *goquery.Selection) {
		// 跳过第一个 -> 不限
		if i != 0 {
			setAreaCache(selector.Text())
		}
	})
}

func parseBusiness(area string, resp *http.Response) {
	doc := parseDoc(resp)
	doc.Find("dl.condition-area").Find("a").Each(func(i int, selector *goquery.Selection) {
		if i != 0 {
			setBusinessCache(area, selector.Text())
		}
	})
}

func parseJobList(resp *http.Response) {
	doc := parseDoc(resp)
	if !hasJobs(doc) {
		return
	}

	area, business := parseLocation(doc)
	doc.Find("div.company-list").Siblings().Find("ul").Find("li").Each(func(i int, selector *goquery.Selection) {
		salary := parseSalary(selector)
		experience := parseExperience(selector)
		industry := parseIndustry(selector)

		saveJD(area, business, salary, experience, industry)
	})
}

func parseLocation(doc *goquery.Document) (string, string) {
	data, _ := doc.Find("div.job-tab").First().Attr("data-filter")
	res := strings.Split(data, `|`)
	area := res[0][1:]
	business := res[1][1:]

	return area, business
}

func parseSalary(selector *goquery.Selection) int {
	content := selector.Find("span.red").First().Text()

	return getStartNumberFromString(content)
}

func getStartNumberFromString(str string) int {
	re := regexp.MustCompile("^[0-9]+")
	numbers := re.FindAllString(str, 1)
	salary, _ := strconv.Atoi(numbers[0])

	return salary
}

func parseExperience(selector *goquery.Selection) string {
	pText := selector.Find("div.info-primary").First().Find("p").First().Text()
	sep := `<em class="vline"></em>`
	res := strings.Split(pText, sep)

	return res[1]
}

func parseIndustry(selector *goquery.Selection) string {
	pText := selector.Find("div.company-text").First().Find("p").First().Text()
	sep := `<em class="vline"></em>`
	res := strings.Split(pText, sep)

	return res[0]
}

func hasJobs(doc *goquery.Document) bool {
	return doc.HasClass("div.company-list")
}

func isBlocked(resp *http.Response) bool {
	return false
}

func parseDoc(resp *http.Response) *goquery.Document {
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)

	return doc
}
