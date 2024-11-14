# Multithreaded Scraper in Go

This project is a **multithreaded web scraper** built with Go, designed to scrape all links present on a main website (e.g., `wikipedia.org`). It leverages **goroutines** and **channels** for concurrent processing, along with proper synchronization mechanisms.

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


