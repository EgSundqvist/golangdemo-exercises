package exercises

import (
	"fmt"
	"math"
	"strings"
)

func Variables1() {
	fmt.Printf("Hello, World!")
}

func Variables2() {
	print("2021")
	print("\n")
	print(2021)
	print("\n")
	print("2021-12-24")
	print("\n")
	print(2021 - 12 - 24)
	print("\n")
}

func Variables3() {

	var age int = 39
	var name string = "Erik Sundqvist"
	if strings.Contains(name, "Erik") {
		fmt.Printf("Hej Erik!\n")
	}
	fmt.Printf("Jag heter %s, och är %d gammal", name, age)

}

func Variables4() {
	var firstName string
	var lastName string
	fmt.Printf("Skriv in ditt förnamn:")
	fmt.Scanln(&firstName)
	fmt.Printf("Skriv in ditt efternamn:")
	fmt.Scanln(&lastName)
	fmt.Printf("Du heter %s %s", lastName, firstName)
}

func Variables5() {
	var num1 int
	var num2 int
	var sum int
	fmt.Printf("Mata in tal 1: ")
	fmt.Scanln(&num1)
	fmt.Printf("Mata in tal 2: ")
	fmt.Scanln(&num2)

	sum = int(math.Pow(float64(num1), float64(num2)))
	fmt.Printf("%d upphöjt till %d är %d\n", num1, num2, sum)

	sum = num1 * num2
	fmt.Printf("num1 * num2 = %d\n", sum)

	sum = num1 / num2
	fmt.Printf("num1 / num2 = %d\n", sum)

	sum = num1 + num2
	fmt.Printf("Summan av %d och %d är %d\n", num1, num2, sum)

	sum = num1 - num2
	fmt.Printf("num1 - num2 = %d\n", sum)

}

func Variables6() {
	var num1 float64
	var num2 float64
	var sum float64

	fmt.Printf("Mata in tal 1: ")
	fmt.Scanln(&num1)
	fmt.Printf("Mata in tal 2: ")
	fmt.Scanln(&num2)

	sum = math.Pow(num1, num2)
	fmt.Printf("%f upphöjt till %f är %f\n", num1, num2, sum)

	sum = num1 * num2
	fmt.Printf("num1 * num2 = %f\n", sum)

	sum = num1 / num2
	fmt.Printf("num1 / num2 = %f\n", sum)

	sum = num1 + num2
	fmt.Printf("Summan av %f och %f är %f\n", num1, num2, sum)

	sum = num1 - num2
	fmt.Printf("num1 - num2 = %f\n", sum)

}

func Variables7() {
	var firstName string
	var firstNameRepeat string
	fmt.Printf("Skriv in ditt förnamn: ")
	fmt.Scanln(&firstName)

	firstNameRepeat = strings.Repeat(firstName, 5)
	fmt.Printf("%s\n", firstNameRepeat)

}

func Variables8() {
	var num int
	fmt.Printf("Mata in ett heltal: ")
	fmt.Scanln(&num)
	fmt.Println(num >= 100)
}
