package menus

import (
	"fmt"

	"github.com/EgSundqvist/golangdemo-exercises/exercises"
	"github.com/EgSundqvist/golangdemo-exercises/utils"
)

func showStringsMenu() {
	for {
		fmt.Println("Choose a Strings exercise:")
		fmt.Println("1. Strings1")
		fmt.Println("2. Strings2")
		fmt.Println("3. Strings3")
		fmt.Println("4. Strings4")
		fmt.Println("5. Strings5")
		fmt.Println("6. Strings6")
		fmt.Println("0. Back to main menu")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			utils.ClearTerminal()
			exercises.Strings1()
			utils.Pause()
			utils.ClearTerminal()
		case 2:
			utils.ClearTerminal()
			exercises.Strings2()
			utils.Pause()
			utils.ClearTerminal()
		case 3:
			utils.ClearTerminal()
			exercises.Strings3()
			utils.Pause()
			utils.ClearTerminal()
		case 4:
			utils.ClearTerminal()
			exercises.Strings4()
			utils.Pause()
			utils.ClearTerminal()
		case 5:
			utils.ClearTerminal()
			exercises.Strings5()
			utils.Pause()
			utils.ClearTerminal()
		case 6:
			utils.ClearTerminal()
			exercises.Strings6()
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
