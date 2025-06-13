package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Ik-cyber/concurrent-downloader/internal/downloader"
	"github.com/vbauerster/mpb/v8"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <url1> <url2> ... <urlN>")
		return
	}

	urls := os.Args[1:]

	var wg sync.WaitGroup
	p := mpb.New(mpb.WithWaitGroup(&wg))

	d := downloader.NewDownloader(p, 3)

	// Setup cancellation context
	ctx, cancel := context.WithCancel(context.Background())

	// Handle Ctrl+C or kill signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalChan
		fmt.Println("\nCancellation signal received. Stopping downloads...")
		cancel()
	}()

	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			fmt.Printf("Starting download: %s\n", url)

			err := d.DownloadFile(ctx, url)
			if err != nil {
				fmt.Printf("Error downloading %s: %v\n", url, err)
				return
			}
		}(url)
	}

	wg.Wait()
	p.Wait()

	fmt.Println("All downloads completed or cancelled!")
}
