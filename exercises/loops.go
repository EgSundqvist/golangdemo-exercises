package exercises

import (
	"fmt"
	"math"
	"strings"
)

func Loops1() {
	// Print numbers from 1 to 10
	for i := 1; i <= 10; i++ {
		println(i)
	}
}

func Loops2() {
	var num1, num2 int
	fmt.Print("Enter the first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter the second number: ")
	fmt.Scanln(&num2)

	start := int(math.Min(float64(num1), float64(num2)))
	end := int(math.Max(float64(num1), float64(num2)))

	fmt.Println("Printing numbers between the chosen numbers")
	for i := start + 1; i < end; i++ {
		fmt.Println(i)
	}
}

func Loops3() {
	var num1, num2 int
	var continueInput string

	for {
		fmt.Print("Enter the first number: ")
		fmt.Scanln(&num1)

		fmt.Print("Enter the second number: ")
		fmt.Scanln(&num2)

		sum := num1 + num2
		fmt.Printf("The sum of %d and %d is %d\n", num1, num2, sum)

		fmt.Print("Do you want to continue? (Y/N)?: ")
		fmt.Scanln(&continueInput)

		continueInput = strings.ToUpper(continueInput)

		if continueInput == "N" {
			break
		}

	}
}
