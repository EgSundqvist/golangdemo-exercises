package menus

import (
	"fmt"

	"github.com/EgSundqvist/golangdemo-exercises/exercises"
	"github.com/EgSundqvist/golangdemo-exercises/utils"
)

func showFunctionsMenu() {
	for {
		fmt.Println("Choose a Functions exercise:")
		fmt.Println("1. Functions1")
		fmt.Println("2. Functions2")
		fmt.Println("3. Functions3")
		fmt.Println("4. Functions4")
		fmt.Println("0. Back to main menu")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			utils.ClearTerminal()
			exercises.Functions1()
			utils.Pause()
			utils.ClearTerminal()
		case 2:
			utils.ClearTerminal()
			exercises.Functions2()
			utils.Pause()
			utils.ClearTerminal()
		case 3:
			utils.ClearTerminal()
			exercises.Functions3()
			utils.Pause()
			utils.ClearTerminal()
		case 4:
			utils.ClearTerminal()
			exercises.Functions4()
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
