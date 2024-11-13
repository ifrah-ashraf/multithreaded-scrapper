package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	t1 := time.Now()

	var urls = []string{"https://news.ycombinator.com/", "https://monkeytype.com/", "https://www.amazon.in/"}
	var wgSend sync.WaitGroup
	var wgRec sync.WaitGroup
	linkCH := make(chan string)

	for _, url := range urls {
		wgSend.Add(1)
		go func(url string) {
			defer wgSend.Done()
			linkCH <- url
		}(url)
	}

	go func() {
		wgSend.Wait()
		close(linkCH)
	}()

	for link := range linkCH {
		wgRec.Add(1)
		go func(link string) {
			defer wgRec.Done()
			FetchLink(link)
		}(link)
	}

	wgRec.Wait()
	t2 := time.Now()
	diff := t2.Sub(t1)
	fmt.Println("total Time taken is : ", diff)

}

func FetchLink(link string) {

	url := link

	var urlArr []string

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

	fmt.Printf("\nURL BREAK\n")
	fmt.Println()

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
