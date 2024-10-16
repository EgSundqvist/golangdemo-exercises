package exercises

import (
	"fmt"
)

func ConcStrings(str1, str2 string) string {
	return fmt.Sprintf("%s %s", str1, str2)
}

func Functions1() {
	var str1, str2 string

	fmt.Println("Enter string 1: ")
	fmt.Scanln(&str1)

	fmt.Println("Enter string 2: ")
	fmt.Scanln(&str2)

	concStrings := ConcStrings(str1, str2)
	fmt.Println("The concatenated string is: ", concStrings)
}
