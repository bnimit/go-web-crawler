package datafetcher

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strings"
	"web-crawler/pkg/utils"
)

type wordMapEntry struct {
	Word  string
	Count int
}

func ReadURLs(filePath string) ([]string, error) {
	var urls []string

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filePath, err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := strings.TrimSpace(scanner.Text())
		if url != "" {
			urls = append(urls, url)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return urls, nil
}

func FetchHtmlContent(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch HTML content for the url: %s", url)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func CountValidWords(content string, validWords map[string]bool) map[string]int {
	wordCount := make(map[string]int)
	words := strings.Fields(content)

	for _, word := range words {
		word = utils.NormalizeWord(word)

		if _, valid := validWords[word]; valid {
			wordCount[word]++
		}
	}

	return wordCount
}

func TopTenWords(wordCount map[string]int) []wordMapEntry {
	var sortedWords []wordMapEntry

	// create a slice from the word count map
	for key, value := range wordCount {
		sortedWords = append(sortedWords, wordMapEntry{key, value})
	}

	//sort the slice with a custom comparator
	sort.Slice(sortedWords, func(i, j int) bool {
		if sortedWords[i].Count == sortedWords[j].Count {
			return sortedWords[i].Word < sortedWords[j].Word
		}
		return sortedWords[i].Count > sortedWords[j].Count
	})

	topTen := sortedWords
	if len(topTen) > 10 {
		topTen = sortedWords[:10]
	}

	return topTen
}
