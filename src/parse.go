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

func parseBusiness(area string, resp *http.Response) {
	doc := parseDoc(resp)
	doc.Find("dl.condition-area").Find("a").Each(func(i int, selector *goquery.Selection) {
		if i != 0 {
			setBusinessCache(area, selector.Text())
		}
	})
}

func parseJobList(resp *http.Response) {
}

func parseJD(resp *http.Response) {
}

func parseDoc(resp *http.Response) *goquery.Document {
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)

	return doc
}
