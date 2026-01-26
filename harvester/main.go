package main

import (
	"fmt"
	"net/http"
)

func main() {
	Intro()
	for {
		URL, err := TakeInput()
		if err != nil {
			fmt.Printf("\nExiting due to: %v\n", err)
			break
		}

		// validate the url
		_, err = http.Get(URL)
		if err != nil {
			fmt.Printf("Invalid URL: %v\n\n", err)
			fmt.Println("Make sure to copy the WHOLE URL, even with the https")
			continue
		}
		// start the download process
		// ? Start with seeing how much there is (number of pages)
		totalSize, sizeOfPage, err := GetDocSize(URL)
		if err != nil {
			fmt.Printf("Error getting document size: %v\n\n", err)
			continue
		}
		// ! DEBUG PRINT
		fmt.Printf("Document Size: %d bytes\nPage Size: %d bytes\n", totalSize, sizeOfPage)

	}
	// If programs enter this point, either user exited or an error occurred
}
