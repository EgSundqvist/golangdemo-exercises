package menus

import (
	"fmt"

	"github.com/EgSundqvist/golangdemo-exercises/exercises"
	"github.com/EgSundqvist/golangdemo-exercises/utils"
)

func showLoopsMenu() {
	for {
		utils.PrintDecorativeLine()
		fmt.Println("Choose a Loops exercise:")
		fmt.Println("1. Loops1")
		fmt.Println("2. Loops2")
		fmt.Println("3. Loops3")
		fmt.Println("4. Loops4")
		fmt.Println("5. Loops5")
		fmt.Println("0. Back to main menu")
		utils.PrintDecorativeLine()

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			utils.ClearTerminal()
			exercises.Loops1()
			utils.Pause()
			utils.ClearTerminal()
		case 2:
			utils.ClearTerminal()
			exercises.Loops2()
			utils.Pause()
			utils.ClearTerminal()
		case 3:
			utils.ClearTerminal()
			exercises.Loops3()
			utils.Pause()
			utils.ClearTerminal()
		case 4:
			utils.ClearTerminal()
			exercises.Loops4()
			utils.Pause()
			utils.ClearTerminal()
		case 5:
			utils.ClearTerminal()
			exercises.Loops5()
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
