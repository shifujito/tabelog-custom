package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://tabelog.com/"
	// url := "https://qiita.com/Yaruki00/items/b50e346551690b158a79"
	getTabelogScrape(url)
}

func getTabelogScrape(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	//Bodyを読み込む
	doc, err := goquery.NewDocumentFromReader(res.Body)
	doc.Find(".list-rst__body").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})
	fmt.Println(doc.Text())
}
