package menus

import (
	"fmt"

	"github.com/EgSundqvist/golangdemo-exercises/exercises"
	"github.com/EgSundqvist/golangdemo-exercises/utils"
)

func showVariablesMenu() {
	for {
		fmt.Println("Choose a Variables exercise:")
		fmt.Println("1. Variables1")
		fmt.Println("2. Variables2")
		fmt.Println("3. Variables3")
		fmt.Println("4. Variables4")
		fmt.Println("5. Variables5")
		fmt.Println("6. Variables6")
		fmt.Println("7. Variables7")
		fmt.Println("8. Variables8")
		fmt.Println("0. Back to main menu")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			utils.ClearTerminal()
			exercises.Variables1()
			utils.Pause()
			utils.ClearTerminal()
		case 2:
			utils.ClearTerminal()
			exercises.Variables2()
			utils.Pause()
			utils.ClearTerminal()
		case 3:
			utils.ClearTerminal()
			exercises.Variables3()
			utils.Pause()
			utils.ClearTerminal()
		case 4:
			utils.ClearTerminal()
			exercises.Variables4()
			utils.Pause()
			utils.ClearTerminal()
		case 5:
			utils.ClearTerminal()
			exercises.Variables5()
			utils.Pause()
			utils.ClearTerminal()
		case 6:
			utils.ClearTerminal()
			exercises.Variables6()
			utils.Pause()
			utils.ClearTerminal()
		case 7:
			utils.ClearTerminal()
			exercises.Variables7()
			utils.Pause()
			utils.ClearTerminal()
		case 8:
			utils.ClearTerminal()
			exercises.Variables8()
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
