# Multithreaded ~~Scraper~~ Crawler in Go

> **Note:** Despite the name, this project is technically a **web crawler**, not a scraper (ironic, right?). It doesn’t extract structured data like product info or text content — it simply discovers and collects **all the hyperlinks** found on a given website.

This project is a **multithreaded web crawler** built with Go, designed to traverse and collect all links present on a main website (e.g., `wikipedia.org`). It leverages **goroutines** and **channels** for concurrent processing, ensuring efficient and fast link discovery across multiple pages.

## Features

- **Concurrent Scraping**: Uses goroutines to scrape multiple URLs simultaneously.
- **Channel Communication**: Synchronizes goroutines efficiently to manage URL scraping tasks.
- **Customizable URL List**: Easily modify the URLs to scrape by updating the `urls` array.

## Prerequisites

Ensure that **Go** is installed on your machine. [Download Go here](https://golang.org/dl/) if needed.

## Setup Instructions

1. Clone the repository:

   ```bash
   git clone https://github.com/ifrah-ashraf/multithreaded-scrapper.git
   cd multithreaded-scrapper

2. Open main.go and modify the URL list in the urls array:

   ```bash
   var urls = []string{"https://news.ycombinator.com/"}

3. Save the file and run it using

   ```bash
   go run main.go

**Made with ❤️ by ifrah ashraf**


