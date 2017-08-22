package main

import (
	"fmt"

	"regexp"

	"strconv"

	"flag"

	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/gommon/log"
)

var (
	flagURL = flag.String("url", "https://play.google.com/store/apps/category/GAME/collection/topselling_free", "URL to scrap")
)

var (
	rankRegexp    = regexp.MustCompile(`(\d+)`)
	packageRegexp = regexp.MustCompile(`\?id=(.+)`)
)

func main() {
	flag.Parse()

	doc, err := goquery.NewDocument(*flagURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("No, Package, Title, URL")
	doc.Find("a.title").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		pkg := packageRegexp.FindStringSubmatch(url)[1]
		title, _ := s.Attr("title")
		rankString := s.Text()

		rank, err := strconv.ParseInt(rankRegexp.FindString(rankString), 0, 32)
		if err != nil {
			rank = 0
		}

		fmt.Printf("%d, %s, %s, %s\n", rank, pkg, title, url)
	})
}
