package main

import (
	"fmt"
	"strings"
)

func RunExercises() {

	Exercise1()
}

func Exercise1() {
	stringarna := []string{"Erik", "Sundqvist", "är", "bäst"}
	for i, s := range stringarna { // i är index, s är värdet
		fmt.Printf("Index: %d, Värde: %s\n", i, s)
	}

	var age int = 39
	var name string = "Erik Sundqvist"
	if strings.Contains(name, "Erik") {
		fmt.Printf("Hej Erik!\n")
	}

	fmt.Printf("Hello, World!\n")
	fmt.Printf("Jag är %d år gammal.\n", age)
	fmt.Printf("Mitt namn är %s.\n", name)

}
