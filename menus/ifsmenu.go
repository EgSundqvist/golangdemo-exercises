package menus

import (
	"fmt"

	"github.com/EgSundqvist/golangdemo-exercises/exercises"
	"github.com/EgSundqvist/golangdemo-exercises/utils"
)

func showIfsMenu() {
	for {
		fmt.Println("Choose an Ifs exercise:")
		fmt.Println("1. Ifs1")
		fmt.Println("2. Ifs2")
		fmt.Println("3. Ifs3")
		fmt.Println("4. Ifs4")
		fmt.Println("5. Ifs5")
		fmt.Println("6. Ifs6")
		fmt.Println("7. Ifs7")
		fmt.Println("8. Ifs8")
		fmt.Println("0. Back to main menu")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			utils.ClearTerminal()
			exercises.Ifs1()
			utils.Pause()
			utils.ClearTerminal()
		case 2:
			utils.ClearTerminal()
			exercises.Ifs2()
			utils.Pause()
			utils.ClearTerminal()
		case 3:
			utils.ClearTerminal()
			exercises.Ifs3()
			utils.Pause()
			utils.ClearTerminal()
		case 4:
			utils.ClearTerminal()
			exercises.Ifs4()
			utils.Pause()
			utils.ClearTerminal()
		case 5:
			utils.ClearTerminal()
			exercises.Ifs5()
			utils.Pause()
			utils.ClearTerminal()
		case 6:
			utils.ClearTerminal()
			exercises.Ifs6()
			utils.Pause()
			utils.ClearTerminal()
		case 7:
			utils.ClearTerminal()
			exercises.Ifs7()
			utils.Pause()
			utils.ClearTerminal()
		case 8:
			utils.ClearTerminal()
			exercises.Ifs8()
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
