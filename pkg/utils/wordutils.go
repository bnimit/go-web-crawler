package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func IsValidWord(word string) bool {
	if len(word) < 3 {
		return false
	}

	for _, r := range word {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func LoadValidWords(filePath string) (map[string]bool, error) {
	validWordsMap := make(map[string]bool)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalWordCount := 0
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		word = strings.ToLower(word) // Normalize to lowercase
		totalWordCount++

		if IsValidWord(word) {
			validWordsMap[word] = true // Added to the valid word map
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	fmt.Println("Total word count: ", totalWordCount)
	return validWordsMap, nil
}
