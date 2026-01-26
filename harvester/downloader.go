package main

import (
	"fmt"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/PuerkitoBio/goquery"
)

type Worker struct {
}

// TEST map to keep track of visited URLs
var visitedURLs = make(map[string]bool)

func StartRecursiveDownload(URL string, totalSize *atomic.Int64) error {
    if visitedURLs[URL] {
        return nil // already visited
    }
    visitedURLs[URL] = true
    fmt.Println(URL)
	res, err := http.Get(URL)
	if err != nil {
		return fmt.Errorf("Error: %v\n\n", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	baseHost := res.Request.URL.Host

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return fmt.Errorf("Error loading HTTP response body: %v", err)
	}

	// TODO : add number total number of matches found (also respect robots.txt)
	doc.Find("style, nav, footer, script, img, video, header, aside").Each(func(i int, s *goquery.Selection) {
		s.Remove()
	})

	doc.Find("div").Each(func(i int, s *goquery.Selection) {
		downloadAndSave(i, s, totalSize)
	})

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		// First check domain
		href, exists := s.Attr("href")
		if exists {
			link, err := http.NewRequest("GET", href, nil)
			if err == nil && link.URL.Host == baseHost {
				StartRecursiveDownload(link.URL.String(), totalSize)
			}
		} else {
			// TODO: i want to add rejected counter here later
			return
		}
	})


	return nil
}

func downloadAndSave(i int, s *goquery.Selection, totalSize *atomic.Int64) {
	text := s.Text()
	file, err := os.Create(fmt.Sprintf("data/page_%d.txt", i))
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}

	totalSize.Add(int64(len(text)))
}
