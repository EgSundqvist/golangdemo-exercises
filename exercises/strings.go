package exercises

import (
	"fmt"
	"strings"
)

func Strings1() {
	var str1, str2, str3, concStrings string

	fmt.Println("Enter string 1: ")
	fmt.Scanln(&str1)

	fmt.Println("Enter string 2: ")
	fmt.Scanln(&str2)

	fmt.Println("Enter string 3: ")
	fmt.Scanln(&str3)

	concStrings = fmt.Sprintf("%s %s %s", str1, str2, str3)
	fmt.Println(concStrings)

}

func Strings2() {
	strSentance := "Hello world"

	pos1 := strings.Index(strSentance, "w")
	pos2 := strings.LastIndex(strSentance, "o")

	fmt.Println("The first position of 'w' in 'Hello world' is: ", pos1)
	fmt.Println("The last position of 'o' in 'Hello world' is: ", pos2)

}

func Strings3() {
	name := "knut andersson svensson johansdotter"
	words := strings.Fields(name)

	for i, word := range words {
		words[i] = strings.ToUpper(string(word[0])) + word[1:]
	}

	titleName := strings.Join(words, " ")

	fmt.Println(titleName)
}

func Strings4() {
	strSentance := "This is a string that should be modified"

	strSentance = strings.ReplaceAll(strSentance, " ", "*")
	count := strings.Count(strSentance, "*")

	fmt.Println(strSentance)
	fmt.Println("The number of '*' in the string is: ", count)
}

func Strings5() {
	var email string

	fmt.Println("Enter your email: ")
	fmt.Scanln(&email)

	dotIndex := strings.LastIndex(email, ".")
	afterDot := email[dotIndex+1:]

	if len(afterDot) >= 2 && len(afterDot) <= 3 && strings.Contains(email, "@") && strings.Contains(email, ".") {
		fmt.Println("Valid email")
	} else {
		fmt.Println("Invalid email")
	}

}

func Strings6() {
	var input string

	fmt.Println("Enter a word or sentence: ")
	fmt.Scanln(&input)

	noBlanksInput := strings.ReplaceAll(input, " ", "")

	runes := []rune(noBlanksInput)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	reversedInput := string(runes)

	if noBlanksInput == reversedInput {
		fmt.Println("The input is a palindrome")
	} else {
		fmt.Println("The input is not a palindrome")
	}

}
