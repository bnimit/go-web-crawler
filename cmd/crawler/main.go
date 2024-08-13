package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"runtime"
	"sync"
	"time"
	"web-crawler/internal/datafetcher"
	"web-crawler/pkg/utils"
)

func main() {
	maxProcs := runtime.GOMAXPROCS(0)
	fmt.Printf("Current GOMAXPROCS value: %d\n", maxProcs)

	wordFilePath := flag.String("wordFile", "words.txt", "Path to the file containing the valid words")
	urlFilePath := flag.String("urlFile", "urls.txt", "Path to the file containing the URLs to be crawled")

	start := time.Now()
	var wg sync.WaitGroup
	wordCountMap := make(map[string]int)
	mu := sync.Mutex{}

	// load the list of valid words from the word bank
	validWords, err := utils.LoadValidWords(*wordFilePath)

	if err != nil {
		fmt.Println("Error loading valid words:", err)
		return
	}

	fmt.Println("Number of Valid words loaded: ", len(validWords))

	urls, err := datafetcher.ReadURLs(*urlFilePath)

	if err != nil {
		fmt.Println("Error parsingh the list of URLs:", err)
		return
	}

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			htmlToString, err := datafetcher.FetchHtmlContent(url)

			if err != nil {
				fmt.Println("Error fetching content for the page: ", url)
				return
			}

			measuredCounts := datafetcher.CountValidWords(htmlToString, validWords)

			mu.Lock()
			for word, count := range measuredCounts {
				wordCountMap[word] += count
			}
			mu.Unlock()
		}(url)

	}
	wg.Wait()
	topTen := datafetcher.TopTenWords(wordCountMap)

	// Convert the map to a JSON object
	jsonData, err := json.MarshalIndent(topTen, "", "  ")
	if err != nil {
		fmt.Println("failed to convert the response to json", err)
		return
	}
	fmt.Println("The top 10 words and their counts are as follows: ")
	fmt.Println(string(jsonData))

	// Calculate and print the elapsed time
	elapsed := time.Since(start)
	fmt.Printf("Time taken to parse all URLs: %s\n", elapsed)
}
