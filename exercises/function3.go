package exercises

import (
	"fmt"
)

func FindTheLongestWord(words []string) string {
	var longestWord string

	for _, word := range words {
		if len(word) > len(longestWord) {
			longestWord = word
		}
	}
	return longestWord
}

func Functions3() {
	var words []string
	var word string

	fmt.Println("Enter first word: ")
	fmt.Scanln(&word)
	words = append(words, word)

	fmt.Println("Enter second word: ")
	fmt.Scanln(&word)
	words = append(words, word)

	fmt.Println("Enter third word: ")
	fmt.Scanln(&word)
	words = append(words, word)

	longestWord := FindTheLongestWord(words)

	fmt.Println("The longest word is: ", longestWord)

}
