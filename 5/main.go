package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	res, err := http.Get("https://www.yna.co.kr/safe/news")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	doc.Find("#container > div > div.content01 > div > ul > li > article").Each(func(i int, s *goquery.Selection) {
		// title := s.Find("a").Text()
		link, _ := s.Find("a").Attr("href")
		// time := s.Find("span").Text()
		// fmt.Println(title, "https:"+link, time)

		wg.Add(1)
		go (func() {
			res, err := http.Get("https:" + link)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()

			doc, err := goquery.NewDocumentFromReader(res.Body)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(doc.Find("#articleWrap > div > div.writer-zone > a > div > strong").Text())
			wg.Done()
		})()
	})
	wg.Wait()
}
