package main

import (
	"fmt"
	"web-crawler/pkg/utils"
)

func main() {
	wordFilePath := "words.txt"
	validWords, err := utils.LoadValidWords(wordFilePath)

	if err != nil {
		fmt.Println("Error loading valid words:", err)
		return
	}

	fmt.Println("Valid words loaded: ", len(validWords))
}
