package exercises

import "fmt"

func Ifs1() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)
	if num > 10 {
		fmt.Println("The number is greater than 10")
	}
	if num < 10 {
		fmt.Println("The number is less than 10")
	}

}

func Ifs2() {
	var milkPackages int
	fmt.Print("Enter milk packages in stock: ")
	fmt.Scanln(&milkPackages)

	if milkPackages < 10 {
		fmt.Println("Order 30 milk packages")
	} else if milkPackages <= 20 {
		fmt.Println("Order 20 milk packages")
	} else {
		fmt.Println("Do not order any milk packages")
	}
}

func Ifs3() {
	var num1 int
	var num2 int
	var highestNum int
	fmt.Print("Enter number1: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter number2: ")
	fmt.Scanln(&num2)

	if num1 > num2 {
		highestNum = num1
		fmt.Printf("The highest number is number1 %d\n", highestNum)
	} else {
		highestNum = num2
		fmt.Printf("The highest number is number2 %d\n", highestNum)
	}
}

func Ifs4() {
	var num1 int
	var num2 int
	var num3 int
	var highestNum int
	fmt.Print("Enter number1: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter number2: ")
	fmt.Scanln(&num2)
	fmt.Print("Enter number3: ")
	fmt.Scanln(&num3)

	if num1 > num2 && num1 > num3 {
		highestNum = num1
		fmt.Printf("The highest number is number1 %d\n", highestNum)
	} else if num2 > num1 && num2 > num3 {
		highestNum = num2
		fmt.Printf("The highest number is number2 %d\n", highestNum)
	} else {
		highestNum = num3
		fmt.Printf("The highest number is number3 %d\n", highestNum)
	}

}

func Ifs5() {
	fmt.Print("Enter your age category, student adult or senior: ")
	var ageCategory string
	fmt.Scanln(&ageCategory)
	if ageCategory == "student" || ageCategory == "senior" {
		fmt.Println("You pay 20kr")
	} else if ageCategory == "adult" {
		fmt.Println("You pay 30kr")
	} else {
		fmt.Println("Invalid age category")
	}
}

func Ifs6() {
	var birthYear int
	fmt.Print("Enter your birth year: ")
	fmt.Scanln(&birthYear)

	if birthYear >= 1980 && birthYear < 1990 {
		fmt.Println("You are a 80's child")
	} else if birthYear >= 1990 && birthYear < 2000 {
		fmt.Println("You are a 90's child")
	} else {
		fmt.Println("You are neither a 80's or 90's child")

	}
}
