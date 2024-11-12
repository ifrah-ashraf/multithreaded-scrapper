package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	var urls = []string{"https://news.ycombinator.com/"}

	linkCH := make(chan string)

	for _, url := range urls {

		go func(url string) {

			linkCH <- url
		}(url)
	}

	go func() {

		for link := range linkCH {
			FetchLink(link)
		}

	}()

	defer close(linkCH)

	time.Sleep(2 * time.Second)
}

func FetchLink(link string) {

	url := link

	var urlArr []string

	go func(url string) {

		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Failed to fetch url: ", url, err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Fatalf("Failed to fetch %s: %d %s", url, resp.StatusCode, resp.Status)
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			fmt.Println("failed while reading the body in goquery: ", err.Error())
		}

		doc.Find("a").Each(func(i int, elem *goquery.Selection) {

			//atag := elem.Find("a").First()
			link, ok := elem.Attr("href")
			if ok {
				urlArr = append(urlArr, link)
			}

		})

		fmt.Println("received link is ", url)
		fmt.Println()
		UrlValidator(urlArr)
		/* for i, url := range urlArr {
			fmt.Printf("URL %d is %s\n", i, url)
		} */
		fmt.Printf("\nURL BREAK\n")
		fmt.Println()
	}(url)

}

func UrlValidator(urlArr []string) {
	var urlRegexp = regexp.MustCompile(`^(http|https)://[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}(/[a-zA-Z0-9-._~:?#@!$&'()*+,;=]*)*$`)
	// urlRegexp.MatchString(url)
	count := 0
	for i, url := range urlArr {
		res := urlRegexp.MatchString(url)

		if res {
			fmt.Printf("URL %d us %s\n", i, url)
			count++
		}
	}
	fmt.Printf("\nTotal valid URL count is %d\n", count)
}
