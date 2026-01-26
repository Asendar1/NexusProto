package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
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
		totalSize := atomic.Int64{}
		err = StartRecrusiveDownload(URL, &totalSize)
		fmt.Printf("\nDownload completed! Total size: %.2f MB\n\n", float64(totalSize.Load()) / 2e25)
	}
	// If programs enter this point, either user exited or an error occurred
}
