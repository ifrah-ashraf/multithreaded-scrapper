package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

type UserData struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

func main() {
	router := gin.Default()

	router.GET("/getpage", func(c *gin.Context) {
		client := &http.Client{}

		req, err := http.NewRequest("GET", "https://news.ycombinator.com/", nil)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		resp, err := client.Do(req)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		doc.Find("span.titleline").Each(func(i int, elem *goquery.Selection) {

			text := elem.Text()
			fmt.Println("Found ", text)
		})

		c.Header("Content-Type", "application/json")
		c.String(resp.StatusCode, string(body))
	})

	router.Run(":8080")
}
