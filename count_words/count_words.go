package countwords

import (
	"fmt"
	"strconv"
	"strings"
)

func CountWords(s string) []string {
	// words := strings.Split(strings.ToLower(s), " ")
	words := strings.Fields(strings.ToLower(s))

	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	fmt.Println(wordCount)
	var repeatedWords []string
	for word, count := range wordCount {
		if count > 1 {
			repeatedWords = append(repeatedWords, word+"|"+strconv.Itoa(count))
		}
	}
	return repeatedWords
}

func Init() {
	// words := CountWords("La vida es breve pero es mejor ser mejor ante la breve vida ser vida ser vida")
	words := CountWords("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum")
	fmt.Println(words)
}
