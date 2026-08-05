package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(word string) bool {
	// 1. Filter out non-alphanumeric characters and convert to lowercase
	var cleanrunes []rune
	for _, r := range word {
		if unicode.IsNumber(r) || unicode.IsLetter(r) {
			cleanrunes = append(cleanrunes, unicode.ToLower(r))
		}
	}
	// 2. Check if the filtered runes read the same forwards and backwards
	length := len(cleanrunes)

	for i := 0; i < length/2; i++ {
		if cleanrunes[i] != cleanrunes[length-i-1] {
			return false
		}
	}
	return true
}

func main() {
	// Example usage
	testString := "A man, a plan, a canal: Panama"
	answer := isPalindrome(testString)
	fmt.Printf("%v", answer)
}
