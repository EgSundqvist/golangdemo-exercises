package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func Pause() {
	fmt.Println("Press Enter to return to the menu...")
	fmt.Scanln()
}

func ClearTerminal() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func PrintDecorativeLine() {
	fmt.Println("****************************************")
}
