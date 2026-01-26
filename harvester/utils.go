package main

import (
	"fmt"
	"net/http"
)

func Intro() {
	fmt.Println("Welcome to the Harvester")
	fmt.Println("Enter any DOC URL and i'll download any page under it")
	fmt.Println("All data will be saved into a .txt files in /data/\n")
}

func TakeInput() (string, error) {
	var URL string
	fmt.Print("Enter DOC URL: ")
	_, err := fmt.Scanln(&URL)
	return URL, err
}

func GetDocSize(URL string) (int, int, error) {
	resp, _ := http.Get(URL)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return 0, 0, fmt.Errorf("failed to fetch URL: %s", resp.Status)
	}
	baseHost := resp.Request.URL.Host
	// recursive and fetch size (Also respect robots.txt later)
	totalSize := 0
	sizeOfPage := []int{}

}
