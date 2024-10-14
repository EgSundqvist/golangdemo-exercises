package exercises

import (
	"fmt"
	"strings"
)

func Structures1() {
	var numbers [4]int
	var num int

	for i := 0; i < 4; i++ {
		println("Enter a number: ")
		fmt.Scanln(&num)
		numbers[i] = num
	}

	max := numbers[0]
	for _, value := range numbers {
		if value > max {
			max = value
		}
	}

	fmt.Printf("The highest number is %d\n", max)
}

func Structures2() {
	var names []string
	var name, response string

	for {
		fmt.Print("Enter a name: ")
		fmt.Scanln(&name)
		names = append(names, name)

		fmt.Print("Do you want to continue? (Y/N): ")
		fmt.Scanln(&response)
		response = strings.ToUpper(response)

		if response == "N" {
			break
		}
	}

	var longestName string
	for _, n := range names {
		if len(n) > len(longestName) {
			longestName = n
		}
	}

	fmt.Printf("The longest name is %s\n", longestName)
}

func Structures3() {
	fmt.Println("1. Create account")
	fmt.Println("2. Delete account")
	fmt.Println("3. List all account numbers")
	fmt.Println("4. List total balance")
	fmt.Println("5. List all account numbers and balances")
	fmt.Println("6. Exit")
	fmt.Print("Choose an option: ")

	for {
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Create account - Functionality not implemented yet.")
		case 2:
			fmt.Println("Delete account - Functionality not implemented yet.")
		case 3:
			fmt.Println("List all account numbers - Functionality not implemented yet.")
		case 4:
			fmt.Println("List total balance - Functionality not implemented yet.")
		case 5:
			fmt.Println("List all account numbers and balances - Functionality not implemented yet.")
		case 6:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}

}
