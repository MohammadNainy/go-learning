package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Result struct {
	URL  string // URL identifier
	TIME int    // Response time in ms
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Initialize random seed

	urls := []string{"url-1", "url-2", "url-3", "url-4", "url-5"} // List of URLs to fetch

	results := make(chan Result, len(urls)) // Buffered channel for results
	var wg sync.WaitGroup                   // WaitGroup for goroutine sync

	for _, val := range urls {
		wg.Add(1)                    // Increment counter for each goroutine
		go fetch(val, results, &wg)  // Start concurrent fetch
	}

	go func() {
		wg.Wait()        // Wait for all fetches to complete
		close(results)   // Close channel when done
	}()

	for res := range results {
		fmt.Printf("%s done in %dms\n", res.URL, res.TIME) // Print results
	}
}

func fetch(url string, ch chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done() // Signal goroutine completion
	start := time.Now() // Record start time

	delay := rand.Intn(401) + 1000 // Random delay: 1000-1400ms
	time.Sleep(time.Duration(delay) * time.Millisecond) // Simulate network delay

	elapsed := int(time.Since(start).Milliseconds()) // Calculate elapsed time
	ch <- Result{URL: url, TIME: elapsed}           // Send result to channel
}