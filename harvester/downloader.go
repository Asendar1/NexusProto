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

func StartRecrusiveDownload(URL string, totalSize *atomic.Int64) error {
	res, err := http.Get(URL)
	if err != nil {
		return fmt.Errorf("Error: %v\n\n", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return fmt.Errorf("Error loading HTTP response body: %v", err)
	}

	// TODO : add number total number of matches found
	doc.Find("pre").Each(func(i int, s *goquery.Selection) {
		downloadAndSave(i, s, totalSize)
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
