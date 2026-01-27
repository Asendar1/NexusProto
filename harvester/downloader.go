package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync/atomic"

	"github.com/PuerkitoBio/goquery"
)

func StartRecursiveDownload(ctx context.Context, URL string, totalSize *atomic.Int64) error {

	req, err := http.NewRequestWithContext(ctx, "GET", URL, nil)
	if err != nil {
		return fmt.Errorf("Error: %v\n\n", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("Error fetching URL: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return fmt.Errorf("Error loading HTTP response body: %v", err)
	}

	doc.Find("style, nav, footer, script, img, video, header, aside").Remove()

	doc.Find("body").Each(func(i int, s *goquery.Selection) {
		downloadAndSave(i, s, totalSize)
	})
	return nil
}

func downloadAndSave(i int, s *goquery.Selection, totalSize *atomic.Int64) {

	var sb strings.Builder

	// Iterate over semantic elements to preserve structure
	s.Find("h1, h2, h3, p, pre, li, blockquote").Each(func(j int, el *goquery.Selection) {
		cleanText := strings.TrimSpace(el.Text())
		if len(cleanText) > 0 {
			sb.WriteString(cleanText)
			sb.WriteString("\n")
		}
	})

	text := sb.String()
	if len(text) == 0 {
		return
	}

	hashName := sha256.Sum256([]byte(text))
	hashString := hex.EncodeToString(hashName[:])

	if _, err := os.Stat("data"); os.IsNotExist(err) {
		err := os.Mkdir("data", 0755)
		if err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			return
		}
	}

	file, err := os.Create(fmt.Sprintf("data/page_%s.txt", hashString))
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
