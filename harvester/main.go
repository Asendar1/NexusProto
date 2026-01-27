package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	Intro()
	for {
		URL, err := TakeInput()
		if err != nil {
			fmt.Printf("\nExiting due to: %v\n", err)
			break
		}

		// start the download process
		totalSize := atomic.Int64{}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		err = StartRecursiveDownload(ctx, URL, &totalSize)
		cancel()
		fmt.Printf("\nDownload completed! Total size: %.2f MB\n\n", float64(totalSize.Load())/(1024*1024))
	}
	// If programs enter this point, either user exited or an error occurred
}
