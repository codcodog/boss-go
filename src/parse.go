package boss

import (
	"net/http"

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

func parseBusiness() {

}

func parseJobList() {

}

func parseJD() {

}

func parseDoc(resp *http.Response) *goquery.Document {
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)

	return doc
}
