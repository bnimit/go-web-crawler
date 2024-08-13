package main

import (
	"encoding/json"
	"fmt"
	"time"
	"web-crawler/internal/datafetcher"
	"web-crawler/pkg/utils"
)

func main() {
	wordFilePath := "words.txt"
	urlFilePath := "urls.txt"

	start := time.Now()

	// load the list of valid words from the word bank
	validWords, err := utils.LoadValidWords(wordFilePath)

	if err != nil {
		fmt.Println("Error loading valid words:", err)
		return
	}

	fmt.Println("Number of Valid words loaded: ", len(validWords))

	urls, err := datafetcher.ReadURLs(urlFilePath)

	if err != nil {
		fmt.Println("Error parsingh the list of URLs:", err)
		return
	}

	wordCountMap := make(map[string]int)
	for _, url := range urls {
		htmlToString, err := datafetcher.FetchHtmlContent(url)

		if err != nil {
			fmt.Println("Error fetch content for the page: %s", url)
		}

		wordCountMap = datafetcher.CountValidWords(htmlToString, validWords)
	}

	topTen := datafetcher.TopTenWords(wordCountMap)

	// Convert the map to a JSON object
	jsonData, err := json.MarshalIndent(topTen, "", "  ")
	if err != nil {
		fmt.Println("failed to convert the response to json", err)
	}
	fmt.Println("The top 10 words and their counts are as follows: ")
	fmt.Println(string(jsonData))

	// Calculate and print the elapsed time
	elapsed := time.Since(start)
	fmt.Printf("Time taken to parse all URLs: %s\n", elapsed)
}
