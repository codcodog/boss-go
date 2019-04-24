package boss

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func parseArea(contents []byte) {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	fmt.Println(doc)
}

func parseBusiness() {

}

func parseJobList() {

}

func parseJD() {

}
