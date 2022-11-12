package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ShopInfo struct {
	shopName          string
	url               string
	rate              float64
	Price             string
	stationName       string
	distanceToStation string
	category          string
	address           string
	tel               string
	homepage          string
}

func (si *ShopInfo) getShopInfo(url string) {
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
	doc.Find(".list-rst__rst-data").Each(func(i int, s *goquery.Selection) {
		// url取得有無フラグ
		var urlExists bool
		// 店名を取得
		si.shopName = s.Find("a").First().Text()
		// 店情報のURLを取得
		si.url, urlExists = s.Find("a").First().Attr("href")
		// 取得できなかった際に、パニックを発生
		if !urlExists {
			panic(si.shopName + "のURLが取得できませんでした")
		}
		// 評価を取得
		si.rate, err = strconv.ParseFloat(s.Find(".c-rating__val").Text(), 64)
		if err != nil {
			panic(err)
		}
		// 金額取得
		si.Price = s.Find(".c-rating-v3__val").First().Text()
		// 駅名、距離、カテゴリを取得
		stationNameAndDistanceAndCategorySlice := strings.Split(s.Find(".list-rst__area-genre").Text(), " / ")
		si.stationName = strings.Split(stationNameAndDistanceAndCategorySlice[0], " ")[0]
		si.distanceToStation = strings.Split(stationNameAndDistanceAndCategorySlice[0], " ")[1]
		si.category = stationNameAndDistanceAndCategorySlice[1]
	})
}

func main() {
	url := "https://tabelog.com/tokyo/rstLst/?SrtT=rt&sort_mode=1&LstCosT=5"
	shopInfoSlice := ShopInfo{}
	shopInfoSlice.getShopInfo(url)
}
