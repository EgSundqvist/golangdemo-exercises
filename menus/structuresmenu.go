package menus

import (
	"fmt"

	"github.com/EgSundqvist/golangdemo-exercises/exercises"
	"github.com/EgSundqvist/golangdemo-exercises/utils"
)

func showStructuresMenu() {
	for {
		fmt.Println("Choose a Structures exercise:")
		fmt.Println("1. Structures1")
		fmt.Println("2. Structures2")
		fmt.Println("3. Structures3")
		fmt.Println("0. Back to main menu")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			utils.ClearTerminal()
			exercises.Structures1()
			utils.Pause()
			utils.ClearTerminal()
		case 2:
			utils.ClearTerminal()
			exercises.Structures2()
			utils.Pause()
			utils.ClearTerminal()
		case 3:
			utils.ClearTerminal()
			exercises.Structures3()
			utils.Pause()
			utils.ClearTerminal()
		case 0:
			utils.ClearTerminal()
			return
		default:
			utils.ClearTerminal()
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
