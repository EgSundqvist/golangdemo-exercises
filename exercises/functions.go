package exercises

import (
	"errors"
	"fmt"
)

func ConcStrings(str1, str2 string) string {
	return fmt.Sprintf("%s %s", str1, str2)
}

func CalculateVAT(price float64) float64 {
	return price * 0.25
}

func FindTheLongestWord(words []string) string {
	var longestWord string

	for _, word := range words {
		if len(word) > len(longestWord) {
			longestWord = word
		}
	}
	return longestWord
}

func CalculateTaxsOnSalary(salary float64) (float64, string, error) {
	if salary < 0 {
		return 0, "", errors.New("negative salary, really?")
	}

	var taxRate float64
	var taxDesc string

	if salary < 15000 {
		taxRate = 0.12
		taxDesc = "12%"
	} else if salary >= 15000 && salary < 30000 {
		taxRate = 0.28
		taxDesc = "28%"
	} else {
		taxRate = 0.50
		taxDesc = "50%"
	}

	tax := salary * taxRate
	return tax, taxDesc, nil

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

func Functions2() {
	var price float64

	fmt.Println("Enter the price: ")
	fmt.Scanln(&price)

	vat := CalculateVAT(price)
	fmt.Println("The VAT is: ", vat)

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

func Functions4() {
	var salary float64

	fmt.Println("Enter your salary: ")
	fmt.Scanln(&salary)

	tax, taxDesc, err := CalculateTaxsOnSalary(salary)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The tax on your salary is %.2f (%s)\n", tax, taxDesc)
	}

}
