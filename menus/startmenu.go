package menus

import (
	"fmt"

	"github.com/EgSundqvist/golangdemo-exercises/utils"
)

func ShowMenu() {
	for {
		fmt.Println("Choose exercise type:")
		fmt.Println("1. Variables")
		fmt.Println("2. Ifs")
		fmt.Println("3. Loops")
		fmt.Println("4. Structures")
		fmt.Println("5. Strings")
		fmt.Println("6. Functions")
		fmt.Println("7. Structs")
		fmt.Println("0. Exit")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			utils.ClearTerminal()
			showVariablesMenu()
		case 2:
			utils.ClearTerminal()
			showIfsMenu()
		case 3:
			utils.ClearTerminal()
			showLoopsMenu()
		case 4:
			utils.ClearTerminal()
			showStructuresMenu()
		case 5:
			utils.ClearTerminal()
			showStringsMenu()
		case 6:
			utils.ClearTerminal()
			showFunctionsMenu()
		case 7:
			utils.ClearTerminal()
			showStructsMenu()
		case 0:
			fmt.Println("Exiting...")
			return
		default:
			utils.ClearTerminal()
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
