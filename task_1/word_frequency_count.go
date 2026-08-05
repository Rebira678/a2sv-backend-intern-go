package main

import (
	"fmt"
	"strings"
	"unicode"
)

func counter(word string) map[string]int {
	//convert everthing into lowercase
	word = strings.ToLower(word)

	//remove punctuation and keep only letter,number and ...
	cleanText := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsSpace(r) {
			return r
		}
		return -1
	}, word)

	//split the clean text into words by spaces
	words := strings.Fields(cleanText)

	//count frequency using a map
	frequencies := make(map[string]int)
	for _, word := range words {
		frequencies[word]++
	}

	return frequencies

}

func main() {
	result := counter("Hello world! Hello Go, hello world.")
	fmt.Println(result)
}
