let one  = "reponse body can only be read once unless you store it like this"
// body, err := io.ReadAll(resp.Body)

let two = "basic REQUEST , RESPONSE through gin framework"

/* router := gin.New()

	err := router.SetTrustedProxies(nil)
	if err != nil {
		panic(err)
	}

	router.GET("/ping/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"message": "hi " + name,
		})
	})

	router.POST("/user/:name/*action", func(c *gin.Context) {
		b := c.FullPath() == "/user/:name/*action" // true
		c.String(http.StatusOK, "%t", b)
	})

	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.GET("/welcome", func(c *gin.Context) {
		fistname := c.DefaultQuery("firstname", "geanie")
		lastname := c.Query("lastname")
		c.JSON(http.StatusOK, gin.H{
			"message": "firstname is " + fistname + " and lastname is " + lastname,
		})
	})

	router.POST("/formpost", func(c *gin.Context) {
		var u UserData

		if err := c.BindJSON(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid json provided"})
		}

		c.JSON(http.StatusOK, u)
	})

	router.Run(":8080") */

/* 
}) */

let three = "using gin is not favourable first learn the http basics" 

/* /*
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

		link := doc.Find("a").First()
		hr, ok := link.Attr("href")
		if ok {
			fmt.Println(hr)
		}

		firstelem := doc.Find("span.titleline").First()
		firstLink, ok := firstelem.Find("a").Attr("href")
		if ok {
			fmt.Println("first link : ", firstLink)
		}

		doc.Find("span.titleline").Each(func(i int, elem *goquery.Selection) {
			atag := elem.Find("a").First()
			link, ok := atag.Attr("href")
			if ok {
				fmt.Println("Link ", i+1, " : ", link)
			}
		})

		/* doc.Find("span.titleline").Each(func(i int, elem *goquery.Selection) {

		}) */

/* c.Header("Content-Type", "application/json")
	c.String(resp.StatusCode, string(body))
})

router.Run(":8080") */ 


let four =  "no need to read the body two times as resp.body is streamed only once"
/* body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("failed while reading the body : ", err)
} */

