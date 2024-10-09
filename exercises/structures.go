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
